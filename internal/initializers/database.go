package initializers

import (
	"errors"

	"github.com/adii1203/link/internal/models"
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

func (p *Postgres) CreateLink(url, key string) (uint, error) {
	link := models.Link{DestinationUrl: url, Key: key}
	result := p.Db.Create(&link)
	if result.Error != nil {
		return 0, result.Error
	}
	return link.ID, nil
}

func (p *Postgres) GetKey(key string) bool {
	tx := p.Db.Where("key = ?", key).First(&models.Link{})
	if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
		return false
	}
	return true
}
