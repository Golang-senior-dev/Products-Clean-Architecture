package productusecase

import (
	"github.com/henriquehendel/productCleanArchitecture/core/domain"
	"github.com/henriquehendel/productCleanArchitecture/core/dto"
)

func (usecase usecase) Fetch(paginationRequestParams *dto.PaginationRequestParams) (*domain.Pagination, error) {
	products, err := usecase.repository.Fetch(paginationRequestParams)

	if err != nil {
		return nil, err
	}

	return products, nil
}
