package interfaces

import (
	"github.com/labstack/echo/v4"
	"google.golang.org/grpc"
)

// EchoRestDelivery delivery factory for echo handler
type EchoRestDelivery interface {
	Mount(group *echo.Group)
}

// GRPCDelivery delivery factory for grpc handler
type GRPCDelivery interface {
	Register(server *grpc.Server)
}

// SubscriberDelivery delivery factory for all subscriber handler
type SubscriberDelivery interface {
	ProcessMessage(topic, key, message []byte)
}
