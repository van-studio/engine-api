package service

type Main struct {
	*dependency
}

func NewMain(i dependency) *Main {
	return &Main{&i}
}

func (x *Main) Index() string {
	return "hello"
}

func (x *Main) Login(email string, password string) {
	//x.Db.Collection()
}
