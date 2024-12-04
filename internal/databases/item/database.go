package item

import (
	"go-api/internal/entities"
)

var MockItems = []entities.Item{
	{ID: 1, Name: "Item 1", Price: 100},
	{ID: 2, Name: "Item 2", Price: 200},
}

type DatabaseInterface interface {
	SetupDatabases() error
	GetItemList() ([]entities.Item, error)
}

type MockItemDatabase struct {
}

func (d *MockItemDatabase) SetupDatabases() error {
	return nil
}

func (d *MockItemDatabase) GetItemList() ([]entities.Item, error) {
	return MockItems, nil
}

func NewDatabase() (*DatabaseInterface, error) {
	var db DatabaseInterface = &MockItemDatabase{}
	var err error = db.SetupDatabases()
	if err != nil {
		return nil, err
	}
	return &db, nil
}
