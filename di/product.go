package di

import (
	productservice "github.com/henriquehendel/productCleanArchitecture/adapter/http/productService"
	"github.com/henriquehendel/productCleanArchitecture/adapter/postgres"
	productrepository "github.com/henriquehendel/productCleanArchitecture/adapter/postgres/productRepository"
	"github.com/henriquehendel/productCleanArchitecture/core/domain"
	productusecase "github.com/henriquehendel/productCleanArchitecture/core/domain/useCase/productUseCase"
)

func ConfigProductDI(conn postgres.PoolInterface) domain.ProductService {
	productRepository := productrepository.New(conn)
	productUseCase := productusecase.New(productRepository)
	productService := productservice.New(productUseCase)

	return productService
}
