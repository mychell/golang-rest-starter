package repository

import (
	"github.com/mychell/go-rest-starter/models"
)

//PostRepository add comment
type PostRepository interface {
	Create(p *models.POST) error
}
