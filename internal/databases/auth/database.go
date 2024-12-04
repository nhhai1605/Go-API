package auth

import (
	"errors"
)

type LoginRequest struct {
	Username string
	Password string
}

var mockTokens = map[LoginRequest]string{
	{Username: "user1", Password: "password1"}: "token1",
	{Username: "user2", Password: "password2"}: "token2",
}

type DatabaseInterface interface {
	SetupDatabases() error
	GetToken(username string, password string) (string, error)
}

type MockDatabase struct {
}

func (d *MockDatabase) SetupDatabases() error {
	return nil
}

func (d *MockDatabase) GetToken(username string, password string) (string, error) {
	var token string = ""
	request := LoginRequest{Username: username, Password: password}
	token, ok := mockTokens[request]
	if !ok {
		return "", errors.New("Invalid username or password")
	}
	return token, nil
}

func NewDatabase() (*DatabaseInterface, error) {
	var db DatabaseInterface = &MockDatabase{}
	var err error = db.SetupDatabases()
	if err != nil {
		return nil, err
	}
	return &db, nil
}
