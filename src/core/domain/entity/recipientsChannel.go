package entity

import (
	"sync"
)

type SafeRecipientsChannel struct {
	sync.Mutex
	Channels map[string]chan Notification
}

func NewRecipientsChannel() *SafeRecipientsChannel {
	return &SafeRecipientsChannel{
		Channels: map[string]chan Notification{},
	}
}
