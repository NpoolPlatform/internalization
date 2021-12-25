package server

import (
	msgsrv "github.com/NpoolPlatform/go-service-framework/pkg/rabbitmq/server" //nolint
	msg "github.com/NpoolPlatform/internalization/pkg/message/message"         //nolint
)

func Init() error {
	return msg.InitQueues()
}

func PublishExample(example *msg.Example) error {
	return msgsrv.PublishToQueue(msg.QueueExample, example)
}
