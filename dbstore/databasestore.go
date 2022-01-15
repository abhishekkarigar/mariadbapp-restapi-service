package dbstore

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"mariadb-service/config"
	"mariadb-service/domain"
)
import "github.com/go-sql-driver/mysql"

type gormDatabase struct {
	Db *gorm.DB
}

func NewGormDatabase(appconfig *config.AppConfig) *gormDatabase {
	c := &mysql.Config{
		Net:                  "tcp",
		Addr:                 fmt.Sprintf("%v:%v", appconfig.Db.Host, appconfig.Db.Port),
		DBName:               appconfig.Db.DBName,
		User:                 appconfig.Db.Username,
		Passwd:               appconfig.Db.Password,
		AllowNativePasswords: true,
		ParseTime:            true,
	}
	gormdb, err := gorm.Open("mysql", c.FormatDSN())
	if err != nil {
		return nil
	}
	return &gormDatabase{
		Db: gormdb,
	}
}

func (g *gormDatabase) GetAll() ([]domain.User, error) {
	var user1 []domain.User
	err := g.Db.Find(&user1).Error
	if err != nil {
		return nil, nil
	}
	return user1, nil
}

func (g *gormDatabase) CreateUser(user domain.User) error {
	//err := g.Db.Model(&domain.User{}).Create(&user).Error
	err := g.Db.Model(&user).Create(&user).Error
	if err != nil {
		return err
	}
	return nil
}
