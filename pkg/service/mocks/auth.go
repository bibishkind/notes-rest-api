package mocks

func (m *MockedService) CreateUser(username string, password string) (int, error) {
	args := m.Called(username, password)
	return args.Get(0).(int), args.Error(1)
}

func (m *MockedService) GenerateJWT(username string, password string) (string, error) {
	args := m.Called(username, password)
	return args.Get(0).(string), args.Error(1)
}

func (m *MockedService) ParseJWT(tokenString string) (int, error) {
	args := m.Called(tokenString)
	return args.Get(0).(int), args.Error(1)
}
