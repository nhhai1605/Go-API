package auth

import (
	"errors"
	"go-api/internal/entities"
)


var MockTokens = map[entities.LoginRequest]string{
	{Username: "user1", Password: "password1"}: "token1",
	{Username: "user2", Password: "password2"}: "token2",
}


type DatabaseInterface interface {
	SetupDatabases() error
	GetToken(loginRequest entities.LoginRequest) (string, error)
}

type MockAuthDatabase struct {
}

func (d *MockAuthDatabase) SetupDatabases() error {
	return nil
}

func (d *MockAuthDatabase) GetToken(loginRequest entities.LoginRequest) (string, error) {
	var token string = ""
	token, ok := MockTokens[loginRequest]
	if !ok {
		return "", errors.New("Invalid username or password")
	}
	return token, nil
}

func NewDatabase() (*DatabaseInterface, error) {
	var db DatabaseInterface = &MockAuthDatabase{}
	var err error = db.SetupDatabases()
	if err != nil {
		return nil, err
	}
	return &db, nil
}
