package interfaces

import (
	"github.com/jain-chetan/catalog-service/model"
)

//DBClient variable type to access interface methods
var DBClient DBInteractions

//DBInteractions interface to hold database operations
type DBInteractions interface {
	DBConnect(config model.DBConfig) error
	CreateProductsQuery(catalog model.Catalog) (model.CreateResponse, error)
	GetAllProductsQuery(queryParams map[string][]string) ([]model.Catalog, error)
	GetSingleProductQuery(productID string) (model.Catalog, error)
	UpdateProductQuery(productID string, catalog model.Catalog)
	CheckProductExist(productID string) bool
}
