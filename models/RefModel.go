package models

type Outbox struct {
	Rphone   string `json:"phone,omitempty"`
	Rmessage string `json:"message,omitempty"`
}

type OutboxMail struct {
	Remail   string `json:"email,omitempty"`
	Rsubject string `json:"subject,omitempty"`
	Rcontent string `json:"content,omitempty"`
}
