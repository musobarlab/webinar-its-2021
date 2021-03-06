package delivery

import (
	"fmt"

	"gitlab.com/Wuriyanto/go-codebase/internal/modules/auth/v1/usecase"
)

// KafkaHandler struct
type KafkaHandler struct {
	uc usecase.AuthUsecase
}

// NewKafkaHandler constructor
func NewKafkaHandler(uc usecase.AuthUsecase) *KafkaHandler {
	return &KafkaHandler{
		uc: uc,
	}
}

// ProcessMessage from kafka consumer
func (h *KafkaHandler) ProcessMessage(topic, key, msg []byte) {
	fmt.Println("from auth module", string(topic), string(key), string(msg))
}
