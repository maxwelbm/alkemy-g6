package service

import (
	"errors"

	models "github.com/maxwelbm/alkemy-g6/internal/models/employees"
	"github.com/maxwelbm/alkemy-g6/internal/repository"
)

var (
	ErrWareHousesServiceNotFound   = errors.New("Warehouse not found")
	ErrEmployeesRepositoryNotFound = errors.New("Employees not found")
)

type EmployeesDefault struct {
	repo repository.RepoDB
}

func NewEmployeesDefault(repo repository.RepoDB) *EmployeesDefault {
	return &EmployeesDefault{repo: repo}
}

func (e *EmployeesDefault) GetAll() (employees map[int]models.Employees, err error) {
	employees, err = e.repo.EmployeesDB.GetAll()
	return
}

func (e *EmployeesDefault) GetByID(id int) (employees models.Employees, err error) {
	employees, err = e.repo.EmployeesDB.GetByID(id)
	return
}

func (e *EmployeesDefault) Create(employees models.EmployeesDTO) (newEmployees models.Employees, err error) {

	if _, err = e.repo.WarehouseDB.GetById(*employees.WarehouseID); err != nil {
		err = ErrWareHousesServiceNotFound
		return
	}

	newEmployees, err = e.repo.EmployeesDB.Create(employees)
	return
}

func (e *EmployeesDefault) Update(employees models.EmployeesDTO, id int) (newEmployees models.Employees, err error) {

	if employees.WarehouseID != nil {
		if _, err = e.repo.WarehouseDB.GetById(*employees.WarehouseID); err != nil {
			err = ErrWareHousesServiceNotFound
			return
		}
	}

	newEmployees, err = e.repo.EmployeesDB.Update(employees, id)
	return
}

func (e *EmployeesDefault) Delete(id int) (err error) {
	err = e.repo.EmployeesDB.Delete(id)
	return
}
