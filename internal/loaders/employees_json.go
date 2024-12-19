package loaders

import (
	"encoding/json"
	"os"

	models "github.com/maxwelbm/alkemy-g6/internal/models/employees"
)

func NewEmployeesJSONFile(path string) *EmployeesJSONFile {
	return &EmployeesJSONFile{
		path: path,
	}
}

type EmployeesJSONFile struct {
	path string
}

type EmployeesJSON struct {
	ID           int    `json:"id"`
	CardNumberID string `json:"card_number_id"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	WarehouseID  int    `json:"warehouse_id"`
}

func (l *EmployeesJSONFile) Load() (employees map[int]models.Employees, err error) {
	// open file
	file, err := os.Open(l.path)
	if err != nil {
		return
	}
	defer file.Close()

	// decode file
	var employeesJSON []EmployeesJSON
	err = json.NewDecoder(file).Decode(&employeesJSON)
	if err != nil {
		return
	}

	// serialize employees
	employees = make(map[int]models.Employees)
	for _, e := range employeesJSON {
		employees[e.ID] = models.Employees{
			ID:           e.ID,
			CardNumberID: e.CardNumberID,
			FirstName:    e.FirstName,
			LastName:     e.LastName,
			WarehouseID:  e.WarehouseID,
		}
	}

	return
}
