package service

type Main struct {
}

func NewMain() *Main {
	return &Main{}
}

func (x *Main) Index() string {
	return "hello"
}

func (x *Main) Login(email string, password string) {
}
