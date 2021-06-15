package service

type IndexService struct {
	*Dependency
}

func NewIndex(dep Dependency) *IndexService {
	return &IndexService{&dep}
}

func (m *IndexService) Index() string {
	return "hello"
}
