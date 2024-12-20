package initializers

import (
	"errors"

	"github.com/adii1203/link/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Storage interface {
	CreateLink(linkData *models.Link) (*models.Link, error)
	GetKey(key string) (bool, *models.Link, error)
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

func (p *Postgres) CreateLink(linkData *models.Link) (*models.Link, error) {
	result := p.Db.Create(linkData)
	if result.Error != nil {
		return nil, result.Error
	}
	return linkData, nil
}

func (p *Postgres) GetKey(slug string) (bool, *models.Link, error) {
	link := &models.Link{}
	tx := p.Db.Where("slug = ?", slug).First(link)
	if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
		return false, nil, tx.Error
	}

	return true, link, nil
}

// func (p *Postgres) IncrementClicks(linkId uint) {
// 	p.Db.Model(&models.LinkAnalytics{}).Where("lind_id = ?", linkId).Update("clicks", gorm.Expr("clicks + ?", 1))

// }
