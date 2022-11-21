package producer

import (
	"fmt"
	"testing"

	"github.com/Shopify/sarama/mocks"
)

func TestSendMessage(t *testing.T) {
	t.Run("Send Message OK", func(t *testing.T) {
		//membuat producer mocks
		mockedProducer := mocks.NewSyncProducer(t, nil)

		//membuat expect producer success atau berhasil mengirim pesan
		mockedProducer.ExpectSendMessageAndSucceed()
		kafka := &KafkaProducer{
			Producer: mockedProducer,
		}

		msg := "Message 1"
		err := kafka.SendMessage("test topic", msg)
		if err != nil {
			t.Errorf("Send message should not be error but have: %v", err)
		}
	})
	t.Run("Send Message NOK", func(t *testing.T) {
		mockedProducer := mocks.NewSyncProducer(t, nil)

		//membuat producer gagal mengirim pesan
		mockedProducer.ExpectSendMessageAndFail(fmt.Errorf("error"))
		kafka := &KafkaProducer{
			Producer: mockedProducer,
		}

		msg := "Message 1"
		err := kafka.SendMessage("test topic", msg)
		if err != nil {
			t.Error("This should be err")
		}
	})
}
