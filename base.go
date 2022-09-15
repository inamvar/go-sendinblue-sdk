package blue

type Address struct {
	Name  string `json:"name,omitempty"`
	Email string `json:"email"`
}

type EmailHeader map[string]string
type EmailParams map[string]string

type Message struct {
	To          []Address   `json:"to"`
	Cc          []Address   `json:"cc,omitempty"`
	Bcc         []Address   `json:"bcc,omitempty"`
	Sender      *Address    `json:"sender"`
	Subject     string      `json:"subject"`
	Tags        []string    `json:"tags,omitempty"`
	ReplyTo     string      `json:"replyTo,omitempty"`
	TextContent string      `json:"textContent,omitempty"`
	HtmlContent string      `json:"htmlContent,omitempty"`
	Headers     EmailHeader `json:"headers,omitempty"`
	TemplateId  int         `json:"templateId,omitempty"`
	Params      EmailParams `json:"params,omitempty"`
}
