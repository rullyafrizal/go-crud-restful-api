package repositories

import (
	"context"
	"database/sql"
	"github.com/rullyafrizal/go-crud-restful-api/models/domain"
)

type CategoryRepository interface {
	FindAll(ctx context.Context, tx sql.Tx, categories *domain.Category) ([]domain.Category, error)
	FindById(ctx context.Context, tx sql.Tx, id int64, category *domain.Category) (domain.Category, error)
	Save(ctx context.Context, tx sql.Tx, category domain.Category) (domain.Category, error)
	Update(ctx context.Context, tx sql.Tx, id int64) (domain.Category, error)
	Delete(ctx context.Context, tx sql.Tx, id int64) error
}
