package auth

import (
	"errors"
	"net/http"
)

type LoginRequest struct {
	Username string
	Password string
}
func (a *LoginRequest) Bind(r *http.Request) error {
	return nil
}
var mockTokens = map[LoginRequest]string{
	{Username: "user1", Password: "password1"}: "token1",
	{Username: "user2", Password: "password2"}: "token2",
}

type DatabaseInterface interface {
	SetupDatabases() error
	GetToken(loginRequest LoginRequest) (string, error)
}

type MockDatabase struct {
}

func (d *MockDatabase) SetupDatabases() error {
	return nil
}

func (d *MockDatabase) GetToken(loginRequest LoginRequest) (string, error) {
	var token string = ""
	token, ok := mockTokens[loginRequest]
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
