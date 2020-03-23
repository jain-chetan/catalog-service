package interfaces

import (
	"github.com/jain-chetan/catalog-service/model"
)

//DBClient variable type to access interface methods
var DBClient DBInteractions

//DBInteractions interface to hold database operations
type DBInteractions interface {
	DBConnect(config model.DBConfig) error
}
