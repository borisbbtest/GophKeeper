package integrationcmd

import (
	"github.com/borisbbtest/GoMon/internal/models/mgrevent"
)

type ServiceWrapperEvents struct {
	Events mgrevent.EventsClient //соединений gRPC Events
}
