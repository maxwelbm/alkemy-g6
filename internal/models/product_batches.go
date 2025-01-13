package models

type ProductBatches struct {
	ID                 int
	BatchNumber        string
	InitialQuantity    int
	CurrentQuantity    int
	CurrentTemperature float64
	MinimumTemperature float64
	DueDate            string
	ManufacturingDate  string
	ManufacturingHour  string
	ProductID          int
	SectionID          int
}

type ProductBatchesDTO struct {
	BatchNumber        string
	InitialQuantity    int
	CurrentQuantity    int
	CurrentTemperature float64
	MinimumTemperature float64
	DueDate            string
	ManufacturingDate  string
	ManufacturingHour  string
	ProductID          int
	SectionID          int
}

type ProductBatchesService interface {
	Create(prodBatches ProductBatchesDTO) (newProdBatches ProductBatches, err error)
}

type ProductBatchesRepository interface {
	Create(prodBatches ProductBatchesDTO) (newProdBatches ProductBatches, err error)
}
