package interfaces

import "github.com/Irsad99/LectronicApp/src/database/gorm/models"

type TokenRepository interface {
	FindByToken(token string) (*models.Token, error)
	InsertToken(data *models.Token) (*models.Token, error)
	DeleteToken(token string) (*models.Token, error)
}
