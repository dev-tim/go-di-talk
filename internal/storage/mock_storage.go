package storage

type MockStorage struct{}

func NewMockStorage() *MockStorage {
	return &MockStorage{}
}

func (s *MockStorage) RetrieveMessage() string {
	return "Hello, Mock Wire!"
}
