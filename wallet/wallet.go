package wallet

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"os/signal"
	"sync/atomic"
	"syscall"
	"time"

	delivery "github.com/ftomza/go-qcbalu/wallet/delivery/mbs"
	"github.com/ftomza/go-qcbalu/wallet/repository"
	"github.com/ftomza/go-qcbalu/wallet/usecase"

	"github.com/ftomza/go-qcbalu/pkg/service"

	"github.com/facebook/ent/dialect"
	"github.com/ftomza/go-qcbalu/pkg/mbs/amqpplus"
	"github.com/ftomza/go-qcbalu/pkg/viperplus"
	"github.com/ftomza/go-qcbalu/pkg/zapplus"
	"github.com/ftomza/go-qcbalu/wallet/repository/ent"
	"github.com/streadway/amqp"
	"go.uber.org/zap"

	entsql "github.com/facebook/ent/dialect/sql"
	_ "github.com/ftomza/go-qcbalu/wallet/repository/ent/runtime"
	_ "github.com/jackc/pgx/v4/stdlib"
)

const (
	NAME    = "qcb_wallet"
	VERSION = "1.0.0"

	ConfigNodeName       = "node_name"
	ConfigEntUrl         = "ent_url"
	ConfigMBSUrl         = "mbs_url"
	ConfigMBSRPCQueue    = "rpc_queue"
	ConfigMBSPUBExchange = "pub_exchange"
	ConfigMBSTimeout     = "mbs_timeout"
)

type Service struct {
	service.Service

	Logger *zapplus.Logger
	Config *viperplus.Viper
	Ent    *ent.Client
	MBS    *amqpplus.MBS
}

func NewService() *Service {
	svc := &Service{
		Service: *service.NewService(),
	}

	svc.Config = NewConfig()

	svc.AddItem(func() (start service.FnService, close service.FnService) {
		svc.Logger = zapplus.New(FullName(svc.Config), svc.Config)
		return nil, func() {
			_ = svc.Logger.Sync()
		}
	}).AddItem(func() (start service.FnService, close service.FnService) {
		svc.Logger.Info("Begin start service",
			zap.String("version", VERSION),
			zap.String("node", NodeName(svc.Config)))
		svc.Ent = NewEnt(svc.Config, svc.Logger)
		return nil, func() {
			_ = svc.Ent.Close()
		}
	}).AddItem(func() (start service.FnService, close service.FnService) {
		svc.MBS = NewMBS(svc.Config, svc.Logger)
		return nil, func() {
			_ = svc.MBS.Close()
		}
	}).AddItem(func() (start service.FnService, close service.FnService) {
		return NewDeliver(svc.Config, svc.Logger, svc.MBS, svc.Ent)
	}).AddItem(func() (start service.FnService, close service.FnService) {
		return func() {
				var state int32 = 1
				sc := make(chan os.Signal)
				signal.Notify(sc, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
				select {
				case sig := <-sc:
					atomic.StoreInt32(&state, 0)
					svc.Logger.Info("Get the exit", zap.String("signal", sig.String()))
				}
				svc.Shutdown(func() {
					svc.Logger.Info("Exit service")
					os.Exit(int(atomic.LoadInt32(&state)))
				})
			}, func() {
				svc.Logger.Info("Shutdown service")
			}
	})

	return svc
}

func (s *Service) Start() {
	s.Service.Start(nil)
}

func NewConfig() *viperplus.Viper {
	config := viperplus.New(NAME)
	config.SetX(ConfigEntUrl, "postgresql://artzab:secret@127.0.0.1/qcb_wallet")
	config.SetX(ConfigMBSUrl, "amqp://localhost/qcbalu")
	config.SetX(ConfigMBSRPCQueue, "rpc.wallet.main")
	config.SetX(ConfigMBSPUBExchange, "pub.wallet")
	config.SetX(ConfigMBSTimeout, time.Second*10)
	config.SetX(ConfigNodeName, "0")
	return config
}

func NodeName(config *viperplus.Viper) string {
	return config.GetString(ConfigNodeName)
}

func FullName(config *viperplus.Viper) string {
	return fmt.Sprintf("%s_%s", NAME, config.GetString(ConfigNodeName))
}

func NewEnt(config *viperplus.Viper, logger *zapplus.Logger) *ent.Client {
	db, err := sql.Open("pgx", config.GetString(ConfigEntUrl))
	if err != nil {
		logger.Fatal("Open DB", zap.Error(err))
	}

	drv := entsql.OpenDB(dialect.Postgres, db)
	client := ent.NewClient(ent.Driver(drv))

	ctx := context.Background()
	if err := client.Schema.Create(ctx); err != nil {
		logger.Fatal("Migration schema", zap.Error(err))
	}

	return client
}

func NewMBS(config *viperplus.Viper, logger *zapplus.Logger) *amqpplus.MBS {

	cn, err := amqp.Dial(config.GetString(ConfigMBSUrl))
	if err != nil {
		logger.Fatal("Connect AMQP", zap.Error(err))
	}
	conn, err := amqpplus.NewAMQPPlusConn(cn, logger,
		amqpplus.SetTimeout(config.GetDuration(ConfigMBSTimeout)),
	)
	if err != nil {
		logger.Fatal("AMQPPlus MBS", zap.Error(err))
	}
	return conn
}

func NewDeliver(config *viperplus.Viper, logger *zapplus.Logger, mbs *amqpplus.MBS, db *ent.Client) (start service.FnService, close service.FnService) {

	rpc := mbs.NewRPCService(FullName(config), config.GetString(ConfigMBSRPCQueue))
	pub := mbs.NewPUBService(config.GetString(ConfigMBSPUBExchange))
	repo := repository.NewEntWalletRepository(db)
	ucase, err := usecase.NewWalletUsecase(repo, usecase.SetTimeout(time.Second*10))
	if err != nil {
		logger.Fatal("Usecase", zap.Error(err))
	}
	delivery.NewMBSWallet(rpc, pub, ucase)
	return func() {
			if err := rpc.Run(); err != nil {
				logger.Fatal("MBS RPC Service", zap.Error(err))
			}
		}, func() {
			_ = rpc.Shutdown()
		}
}
