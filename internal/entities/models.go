package entities

type Message struct {
	Id      int    `json:"id"`
	Message string `json:"message"`
}

type AnsweredMessage struct {
	Id      int    `json:"id"`
	Message string `json:"message"`
	Answer  string `json:"answer"`
}
