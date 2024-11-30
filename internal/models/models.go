package models

type BD struct {
	Password string
	Username string
	Login    string
	Token    []byte
	Id       int
}

type Message struct {
	Username string
	Message  string
	Id       int
}
