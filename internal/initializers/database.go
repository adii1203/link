package initializers

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Storage interface {
	CreateLink(url, key string) (uint, error)
	GetKey(key string) bool
}

type Postgres struct {
	Db *gorm.DB
}

func New() (*Postgres, error) {
	dns := "postgresql://link_owner:TLBS6iMEm5pe@ep-yellow-sound-a1rdlum5.ap-southeast-1.aws.neon.tech/link?sslmode=require"
	db, err := gorm.Open(postgres.Open(dns), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return &Postgres{
		Db: db,
	}, nil
}
