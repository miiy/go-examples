package main

import "fmt"

type Message struct {
	id int
	name string
	address string
	phone string
}

type Option func(*Message)

var defaultMessage = Message{
	id:      0,
	name:    "1",
	address: "2",
	phone:   "3",
}

func WithAddress(addr string) Option {
	return func(m *Message) {
		m.address = addr
	}
}

func WithPhone(phone string) Option {
	return func(m *Message) {
		m.phone = phone
	}
}

func NewMessage(id int, name string, opts ...Option) Message {
	msg := defaultMessage
	msg.id = id
	msg.name = name
	for _, o := range opts {
		o(&msg)
	}
	return msg
}

func main()  {
	msg := NewMessage(1, "a", WithPhone("b"))
	fmt.Println(msg)
}