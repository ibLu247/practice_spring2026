package handlers

import (
	db "Rest_API_Server/internal/database"
	"Rest_API_Server/internal/rest/model"
	"context"
	"html/template"
	"net/http"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	rows, err := db.DB.Query(
		context.Background(),
		"SELECT first_name, last_name, position FROM employees",
	)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer rows.Close()

	var employees []model.Employee

	for rows.Next() {
		var employee model.Employee

		rows.Scan(
			&employee.FirstName,
			&employee.LastName,
			&employee.Position,
		)

		employees = append(employees, employee)
	}

	tmpl, err := template.ParseFiles("web/rest/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, employees)
}

func AddEmployeeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	firstName := r.FormValue("first_name")
	lastName := r.FormValue("last_name")
	position := r.FormValue("position")

	_, err := db.DB.Exec(
		context.Background(),
		`
		INSERT INTO employees(first_name, last_name, position)
		VALUES($1, $2, $3)
		`,
		firstName,
		lastName,
		position,
	)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
