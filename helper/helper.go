package helper

import (
	"database/sql"

	_ "github.com/lib/pq"

	"fmt"

	"lantorabde.app/models"
)

var DB *sql.DB

func Insertuser(FullName string, Email string, Phone string, PasswordHash string, DrivingLicense string,Role string,Status string ) (uint, error) {
	var id uint
	query := `INSERT INTO bde_users (full_name,email,phone,password_hash,driving_license,role,join_date,status,created_at,updated_at) 
              VALUES ($1, $2,$3,$4,$5,$6,NOW(),$7,NOW(),NOW()) RETURNING id`
	err := DB.QueryRow(query, FullName, Email, Phone, PasswordHash, DrivingLicense,Role,Status).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func GetUsers() ([]models.BDEUser, error) {
	rows, err := DB.Query(`
			SELECT id,full_name,email,phone,password_hash,driving_license,role,join_date,status,created_at,updated_at  
			FROM bde_users
		`)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var usersList []models.BDEUser

	for rows.Next() {
		var user models.BDEUser
		err := rows.Scan(
			&user.FullName,
			&user.Phone,
			&user.Email,
			&user.PasswordHash,
			&user.DrivingLicense,
			&user.Role,
			&user.JoinDate,
			&user.Status,
			&user.CreatedAt,
			&user.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		usersList = append(usersList, user)
	}

	fmt.Println("GetUsersSuccessful")

	return usersList, nil
}
