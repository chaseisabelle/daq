package message

type Message struct {
	Body string
}

func New(bod string) (*Message, error) {
	return &Message{
		Body: bod,
	}, nil
}
