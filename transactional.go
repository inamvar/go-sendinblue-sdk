package blue

import (
	"net/http"
)

// send a transactional email
func SendTransactional(message *Message) error {

	req, err := makeRequest("smtp/email", http.MethodPost, message)
	if err != nil {
		return err
	}
	_, err = callApi(req)
	if err != nil {
		return err
	}

	return nil
}
