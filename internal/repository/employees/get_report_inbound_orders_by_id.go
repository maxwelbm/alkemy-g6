package repository

import (
	"strconv"

	models "github.com/maxwelbm/alkemy-g6/internal/models"
)

func (r *EmployeesRepository) GetReportInboundOrdersById(id int) (employeesReportList []models.EmployeesReportInboundDTO, err error) {

	var query string
	if id == 0 {
		query = "SELECT e.id, e.card_number_id, e.first_name, e.last_name, e.warehouse_id ,COUNT(io.id) AS CountReports " +
			"FROM employees e " +
			"LEFT JOIN inbound_orders io ON e.id = io.employee_id " +
			"GROUP BY e.id, e.first_name, e.last_name;"
	} else {
		query = "SELECT e.id, e.card_number_id, e.first_name, e.last_name, e.warehouse_id ,COUNT(io.id) AS CountReports " +
			"FROM employees e " +
			"LEFT JOIN inbound_orders io ON e.id = io.employee_id " +
			"WHERE e.id = " + strconv.Itoa(id) + " " +
			"GROUP BY e.id, e.first_name, e.last_name;"
	}

	rows, err := r.DB.Query(query)
	if err != nil {
		return
	}
	defer rows.Close()
	for rows.Next() {
		var employeesReport models.EmployeesReportInboundDTO
		if err = rows.Scan(&employeesReport.ID, &employeesReport.CardNumberID, &employeesReport.FirstName, &employeesReport.LastName, &employeesReport.WarehouseID, &employeesReport.CountReports); err != nil {
			return
		}
		employeesReportList = append(employeesReportList, employeesReport)
	}

	// Check for errors after rows iteration
	if err = rows.Err(); err != nil {
		return
	}

	return
}
