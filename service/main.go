package service

type Main struct {
	*dependency
}

func NewMain(i dependency) *Main {
	return &Main{&i}
}

func (m *Main) Index() string {
	return "hello"
}
