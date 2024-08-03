//go:build wireinject
// +build wireinject

package wire

import (
	"restfullapi/database"
	"restfullapi/handler"
	"restfullapi/repository"
	"restfullapi/service"

	"github.com/google/wire"
)

var productHandlerSet = wire.NewSet(
	database.NewPostgresDB,
	repository.NewProductRepository,
	service.NewProductService,
	handler.NewProductHandler,
)

func InitializProductHandler() *handler.ProductHandler {
	wire.Build(productHandlerSet)
	return nil
}
