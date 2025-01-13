package controllers

import (
	buyers_controller "github.com/maxwelbm/alkemy-g6/internal/controllers/buyer"
	carries_controller "github.com/maxwelbm/alkemy-g6/internal/controllers/carries"
	employees_controller "github.com/maxwelbm/alkemy-g6/internal/controllers/employees"
	localities_controller "github.com/maxwelbm/alkemy-g6/internal/controllers/localities"
	product_batches_controller "github.com/maxwelbm/alkemy-g6/internal/controllers/product_batches"
	product_records_controller "github.com/maxwelbm/alkemy-g6/internal/controllers/product_records"
	products_controller "github.com/maxwelbm/alkemy-g6/internal/controllers/products"
	purchase_orders_controller "github.com/maxwelbm/alkemy-g6/internal/controllers/purchase_orders"
	sections_controller "github.com/maxwelbm/alkemy-g6/internal/controllers/sections"
	sellers_controller "github.com/maxwelbm/alkemy-g6/internal/controllers/seller"
	warehouses_controller "github.com/maxwelbm/alkemy-g6/internal/controllers/warehouses"
	"github.com/maxwelbm/alkemy-g6/internal/models"
)

func NewBuyersController(SV models.BuyerService) *buyers_controller.BuyersDefault {
	return &buyers_controller.BuyersDefault{
		SV: SV,
	}
}

func NewProductsController(SV models.ProductService) *products_controller.ProductsDefault {
	return &products_controller.ProductsDefault{
		SV: SV,
	}
}

func NewSectionsController(sv models.SectionService) *sections_controller.SectionsController {
	return sections_controller.NewSectionsController(sv)
}

func NewSellersController(sv models.SellersService) *sellers_controller.SellersDefault {
	return sellers_controller.NewSellersController(sv)
}

func NewCarriesController(sv models.CarriesService) *carries_controller.CarriesDefault {
	return carries_controller.NewCarriesDefault(sv)
}

func NewLocalityController(sv models.LocalityService) *localities_controller.LocalitiesController {
	return localities_controller.NewLocalitiesController(sv)
}

func NewWarehousesController(service models.WarehouseService) *warehouses_controller.WarehouseDefault {
	return &warehouses_controller.WarehouseDefault{Service: service}
}

func NewProductRecordsController(sv models.ProductRecordsService) *product_records_controller.ProductRecordsDefault {
	return product_records_controller.NewProductRecordsController(sv)
}

func NewEmployeesController(sv models.EmployeesService) *employees_controller.EmployeesController {
	return employees_controller.NewEmployeesDefault(sv)
}

func NewPurchaseOrdersController(service models.PurchaseOrdersService) *purchase_orders_controller.PurchaseOrdersController {
	return purchase_orders_controller.NewPurchaseOrdersController(service)
}

func NewProductBatchesController(sv models.ProductBatchesService) *product_batches_controller.ProductBatchesController {
	return product_batches_controller.NewProductBatchesController(sv)
}
