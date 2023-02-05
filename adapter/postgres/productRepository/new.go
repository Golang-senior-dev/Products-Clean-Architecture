package productrepository

import (
	"github.com/henriquehendel/productCleanArchitecture/adapter/postgres"
	"github.com/henriquehendel/productCleanArchitecture/core/domain"
)

type repository struct {
	db postgres.PoolInterface
}

func New(db postgres.PoolInterface) domain.ProductRepository {
	return &repository{
		db: db,
	}
}
