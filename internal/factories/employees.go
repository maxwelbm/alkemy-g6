package factories

import (
	"database/sql"
	"fmt"
	"strconv"

	"github.com/maxwelbm/alkemy-g6/internal/models"
)

type EmployeeFactory struct {
	db *sql.DB
}

func NewEmployeeFactory(db *sql.DB) *EmployeeFactory {
	return &EmployeeFactory{db: db}
}

func defaultEmployee() models.Employee {
	return models.Employee{}
}

func (f EmployeeFactory) Build(employee models.Employee) models.Employee {
	populateEmployeeParams(&employee)

	return employee
}

func (f *EmployeeFactory) Create(employee models.Employee) (record models.Employee, err error) {
	populateEmployeeParams(&employee)

	if err = f.checkWarehouseExists(employee.WarehouseID); err != nil {
		return employee, err
	}

	query := `
		INSERT INTO employee 
			(
			%s
			card_number_id,
			first_name,
			last_name,
			warehouse_id
			) 
		VALUES (%s?, ?, ?, ?, ?)
	`

	switch employee.ID {
	case 0:
		query = fmt.Sprintf(query, "", "")
	default:
		query = fmt.Sprintf(query, "id,", strconv.Itoa(employee.ID)+",")
	}

	result, err := f.db.Exec(query,
		employee.CardNumberID,
		employee.FirstName,
		employee.LastName,
		employee.WarehouseID,
	)

	if err != nil {
		return record, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return record, err
	}

	employee.ID = int(id)

	return employee, err
}

func populateEmployeeParams(employee *models.Employee) {
	defaultEmployee := defaultEmployee()
	if employee == nil {
		employee = &defaultEmployee
	}

	if employee.CardNumberID == "" {
		employee.CardNumberID = defaultEmployee.CardNumberID
	}

	if employee.FirstName == "" {
		employee.FirstName = defaultEmployee.FirstName
	}

	if employee.LastName == "" {
		employee.LastName = defaultEmployee.LastName
	}

	if employee.WarehouseID == 0 {
		employee.WarehouseID = defaultEmployee.WarehouseID
	}
}

func (f *EmployeeFactory) checkWarehouseExists(warehouseID int) (err error) {
	var count int
	err = f.db.QueryRow(`SELECT COUNT(*) FROM warehouses WHERE id = ?`, warehouseID).Scan(&count)

	if err != nil {
		return
	}

	if count > 0 {
		return
	}

	err = f.createWarehouse()

	return
}

func (f *EmployeeFactory) createWarehouse() (err error) {
	employeeFactory := NewWarehouseFactory(f.db)
	_, err = employeeFactory.Create(models.Warehouse{})

	return
}
