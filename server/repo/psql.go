package repo

import (
	"api-server/model"
	"api-server/pkg/psql"
	"fmt"

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

func (pc *PostgresClient) GetUserByEmail(email string) (user model.User) {

	tx := pc.DB.Find(&user).Where("email = ?", email)
	if tx.Error != nil {
		panic(tx.Error)
	}
	return user
}

func (pc *PostgresClient) GetUserByUuid(uuid string) (user model.User) {
	tx := pc.DB.Find(&user).Where("uuid = ?", uuid)
	if tx.Error != nil {
		panic(tx.Error)
	}
	return user
}

func (pc *PostgresClient) SaveUser(user *model.User) error {
	dbUser := pc.GetUserByEmail(user.Email)
	if dbUser.ID != 0 {
		panic(fmt.Errorf("user %+v already in database", user))
	}

	pc.DB.Create(user)

	return nil
}
