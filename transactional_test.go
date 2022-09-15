package blue_test

import (
	"testing"

	blue "github.com/inamvar/go-sendinblue-sdk"
	"github.com/stretchr/testify/assert"
)

func TestCreateTransactional(t *testing.T) {
	msg := &blue.Message{
		To:          []blue.Address{{Name: "<name>", Email: "user@example.com"}},
		Sender:      &blue.Address{Name: "<sender name>", Email: "no-reply@yourdomain.com"},
		Subject:     "test",
		TextContent: "this is a test",
	}

	err := blue.SendTransactional(msg)
	assert.Nil(t, err)
}
