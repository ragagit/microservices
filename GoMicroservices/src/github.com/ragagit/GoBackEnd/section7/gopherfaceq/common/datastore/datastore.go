package datastore

import "github.com/ragagit/GoBackEnd/section7/gopherfaceq/models"

type Datastore interface {
	CreateUser(user *models.User) error
	GetUser(username string) (*models.User, error)
	Close()
}

const (
	MYSQL = iota
	MONGODB
	REDIS
)

func NewDatastore(datastoreType int, dbConnectionString string) (Datastore, error) {

	switch datastoreType {
	case MYSQL:
		return NewMySQLDatastore(dbConnectionString)
	case MONGODB:
		return NewMongoDBDatastore(dbConnectionString)
	case REDIS:
		return NewRedisDatastore(dbConnectionString)

	}

	return nil, nil
}