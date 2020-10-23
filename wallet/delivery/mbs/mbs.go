package mbs

import (
	"fmt"

	"github.com/ftomza/go-qcbalu/domain"
	"github.com/ftomza/go-qcbalu/pkg/mbs"
)

func newEvent(event, route fmt.Stringer, bytes []byte) *mbs.PUBMessage {
	return &mbs.PUBMessage{
		Headers: mbs.NewEventHeaders(domain.EventHeaders{
			Event: event.String(),
		}),
		Route: route.String(),
		Data:  bytes,
	}
}
