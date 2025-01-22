package controllers

import (
	buyersctl "github.com/maxwelbm/alkemy-g6/internal/controllers/buyers"
	carriesctl "github.com/maxwelbm/alkemy-g6/internal/controllers/carries"
	employeesctl "github.com/maxwelbm/alkemy-g6/internal/controllers/employees"
	inboundordersctl "github.com/maxwelbm/alkemy-g6/internal/controllers/inbound_orders"
	localitiesctl "github.com/maxwelbm/alkemy-g6/internal/controllers/localities"
	productbatchesctl "github.com/maxwelbm/alkemy-g6/internal/controllers/product_batches"
	productrecordsctl "github.com/maxwelbm/alkemy-g6/internal/controllers/product_records"
	productsctl "github.com/maxwelbm/alkemy-g6/internal/controllers/products"
	purchaseordersctl "github.com/maxwelbm/alkemy-g6/internal/controllers/purchase_orders"
	sectionsctl "github.com/maxwelbm/alkemy-g6/internal/controllers/sections"
	sellersctl "github.com/maxwelbm/alkemy-g6/internal/controllers/seller"
	warehousesctl "github.com/maxwelbm/alkemy-g6/internal/controllers/warehouses"
	"github.com/maxwelbm/alkemy-g6/internal/models"
)

func NewBuyersController(sv models.BuyerService) *buyersctl.BuyersDefault {
	return buyersctl.NewBuyersController(sv)
}

func NewProductsController(SV models.ProductService) *productsctl.ProductsDefault {
	return &productsctl.ProductsDefault{
		SV: SV,
	}
}

func NewSectionsController(sv models.SectionService) *sectionsctl.SectionsController {
	return sectionsctl.NewSectionsController(sv)
}

func NewSellersController(sv models.SellersService) *sellersctl.SellersDefault {
	return sellersctl.NewSellersController(sv)
}

func NewCarriesController(sv models.CarriesService) *carriesctl.CarriesDefault {
	return carriesctl.NewCarriesDefault(sv)
}

func NewLocalityController(sv models.LocalityService) *localitiesctl.LocalitiesController {
	return localitiesctl.NewLocalitiesController(sv)
}

func NewWarehousesController(sv models.WarehouseService) *warehousesctl.WarehouseDefault {
	return warehousesctl.NewWarehousesController(sv)
}

func NewInboundOrdersController(SV models.InboundOrdersService) *inboundordersctl.InboundOrdersController {
	return inboundordersctl.NewInboundOrdersDefault(SV)
}

func NewProductRecordsController(sv models.ProductRecordsService) *productrecordsctl.ProductRecordsDefault {
	return productrecordsctl.NewProductRecordsController(sv)
}

func NewEmployeesController(sv models.EmployeesService) *employeesctl.EmployeesController {
	return employeesctl.NewEmployeesDefault(sv)
}

func NewPurchaseOrdersController(service models.PurchaseOrdersService) *purchaseordersctl.PurchaseOrdersController {
	return purchaseordersctl.NewPurchaseOrdersController(service)
}

func NewProductBatchesController(sv models.ProductBatchesService) *productbatchesctl.ProductBatchesController {
	return productbatchesctl.NewProductBatchesController(sv)
}
