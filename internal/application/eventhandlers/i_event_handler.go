package eventhandlers

import (
	"context"
)

type IEventHandler[TNotification any] interface {
	Handle(ctx context.Context, notification TNotification) error
}
