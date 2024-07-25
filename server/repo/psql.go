package repo

import (
	"api-server/model"
	"api-server/pkg/psql"
	"errors"
	"fmt"
	"log"

	"gorm.io/gorm"
)

type PostgresClient struct {
	DB *gorm.DB
}

func New() *PostgresClient {
	db, err := psql.New()
	if err != nil {
		panic(err)
	}
	ret := new(PostgresClient)
	ret.DB = db

	var users []model.User
	db.Find(&users)
	for u := range users {
		fmt.Printf("\tyeah yeah test: %+v\n", u)
	}
	return ret
}

func (pc *PostgresClient) GetUserByEmail(email string) (user *model.User, err error) {

	err = pc.DB.Where("email = ?", email).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return user, nil
}

func (pc *PostgresClient) GetUserByUuid(uuid string) (user *model.User, err error) {

	err = pc.DB.Where("uuid = ?", uuid).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return user, nil
}

func (pc *PostgresClient) SaveUser(user *model.User) error {
	dbUser, err := pc.GetUserByEmail(user.Email)
	if err != nil {
		log.Printf("unknown error: %s\n", err)
		return err
	}

	if dbUser != nil {
		log.Printf("User with email %s found\n", user.Email)
		return fmt.Errorf("email already used")
	}

	pc.DB.Create(user)

	return nil
}
