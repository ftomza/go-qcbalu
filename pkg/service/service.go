package service

import "sort"

type FnService func()

type FnServices []FnService

func (f FnServices) Len() int {
	return len(f)
}

func (f FnServices) Less(i, j int) bool {
	return i < j
}

func (f FnServices) Swap(i, j int) {
	f[i], f[j] = f[j], f[i]
}

type Service struct {
	closers  FnServices
	starters FnServices
}

func NewService() *Service {
	return &Service{
		closers:  FnServices{},
		starters: FnServices{},
	}
}

func (s *Service) AddItem(fn func() (start FnService, close FnService)) *Service {
	starter, closer := fn()
	if closer != nil {
		s.closers = append(s.closers, closer)
	}
	if starter != nil {
		s.starters = append(s.starters, starter)
	}
	return s
}

func (s *Service) Start(fn func()) {
	for _, v := range s.starters {
		v()
	}
	if fn != nil {
		fn()
	}
}

func (s *Service) Shutdown(fn func()) {
	sort.Sort(sort.Reverse(s.closers))
	for _, v := range s.closers {
		v()
	}
	if fn != nil {
		fn()
	}
}
