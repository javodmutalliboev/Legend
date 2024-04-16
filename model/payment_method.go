package model

import (
	"Legend/database"
	"fmt"
	"strings"
)

type PaymentMethod struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Logo      []byte `json:"logo"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func (pm *PaymentMethod) Create() error {
	db := database.DB()
	defer db.Close()

	stmt, err := db.Prepare("INSERT INTO payment_method (name, logo) VALUES ($1, $2)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(pm.Name, pm.Logo)
	if err != nil {
		return err
	}

	return nil
}

func GetPaymentMethods() ([]PaymentMethod, error) {
	db := database.DB()
	defer db.Close()

	rows, err := db.Query("SELECT id, name, created_at, updated_at FROM payment_method")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var paymentMethods []PaymentMethod
	for rows.Next() {
		var paymentMethod PaymentMethod
		err := rows.Scan(&paymentMethod.ID, &paymentMethod.Name, &paymentMethod.CreatedAt, &paymentMethod.UpdatedAt)
		if err != nil {
			return nil, err
		}
		paymentMethods = append(paymentMethods, paymentMethod)
	}

	return paymentMethods, nil
}

func GetPaymentMethod(id int) (*PaymentMethod, error) {
	db := database.DB()
	defer db.Close()

	row := db.QueryRow("SELECT id, name, created_at, updated_at FROM payment_method WHERE id = $1", id)

	var paymentMethod PaymentMethod
	err := row.Scan(&paymentMethod.ID, &paymentMethod.Name, &paymentMethod.CreatedAt, &paymentMethod.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &paymentMethod, nil
}

func GetPaymentMethodLogo(id int) (*[]byte, error) {
	db := database.DB()
	defer db.Close()

	row := db.QueryRow("SELECT logo FROM payment_method WHERE id = $1", id)

	var logo []byte
	err := row.Scan(&logo)
	if err != nil {
		return nil, err
	}

	return &logo, nil
}

func (pm *PaymentMethod) Update() error {
	db := database.DB()
	defer db.Close()

	var fields []string
	var args []interface{}
	i := 1

	if pm.Name != "" {
		fields = append(fields, fmt.Sprintf("name = $%d", i))
		args = append(args, pm.Name)
		i++
	}

	if pm.Logo != nil {
		fields = append(fields, fmt.Sprintf("logo = $%d", i))
		args = append(args, pm.Logo)
		i++
	}

	// build the SQL query
	query := fmt.Sprintf("UPDATE payment_method SET %s, updated_at = NOW() WHERE id = $%d", strings.Join(fields, ", "), i)
	args = append(args, pm.ID)

	stmt, err := db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(args...)
	if err != nil {
		return err
	}

	return nil
}

func DeletePaymentMethod(id int) error {
	db := database.DB()
	defer db.Close()

	stmt, err := db.Prepare("DELETE FROM payment_method WHERE id = $1")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}

	return nil
}
