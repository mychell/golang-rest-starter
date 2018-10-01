package post

import (
	"github.com/gobuffalo/pop"
	"github.com/mychell/go-rest-starter/models"
	"github.com/mychell/go-rest-starter/repository"
)

type pqPostRepository struct {
	conn *pop.Connection
}

// NewPQPostRepository returns implement of post repository interface
func NewPQPostRepository(conn *pop.Connection) repository.PostRepository {
	return &pqPostRepository{
		conn: conn,
	}
}

func (pq *pqPostRepository) Create(p *models.POST) error {
	//fmt.Println(ctx)

	_, err := pq.conn.ValidateAndSave(p)
	if err != nil {
		return err
	}

	return nil

}
