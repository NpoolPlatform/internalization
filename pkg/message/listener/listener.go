package listener

import (
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	msgcli "github.com/NpoolPlatform/internalization/pkg/message/client" //nolint
	msg "github.com/NpoolPlatform/internalization/pkg/message/message"   //nolint
)

func listenTemplateExample() {
	for {
		logger.Sugar().Infof("consume template example")
		err := msgcli.ConsumeExample(func(example *msg.Example) error {
			logger.Sugar().Infof("go example: %+w", example)
			// Call event handler in api module
			return nil
		})
		if err != nil {
			logger.Sugar().Errorf("fail to consume example: %v", err)
			return
		}
	}
}

func Listen() {
	go listenTemplateExample()
}
