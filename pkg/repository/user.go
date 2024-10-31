package repository

import (
	"database/sql"
	models "user-management-service/pkg/model"
	"github.com/google/uuid"
)

func GetAllUsers(db *sql.DB) ([]models.User, error) {
	users := []models.User{}

	query := "SELECT id, email, password, name, category, dob, bio, avatar FROM user"
	rows, err := db.Query(query)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {

		var user models.User

		if err := rows.Scan(&user.Id, &user.Email, &user.Password, &user.Name, &user.Category, &user.DOB, &user.Bio, &user.Avatar); err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}
func GetUserById(db *sql.DB, id string) (models.User, error) {
	var user models.User

	err := db.QueryRow("SELECT id, email, password, name, category, dob, bio, avatar FROM user WHERE id = ?", id).Scan(&user.Id, &user.Email, &user.Password, &user.Name, &user.Category, &user.DOB, &user.Bio, &user.Avatar)

	if err != nil {
		return user, err
	}

	// Format the date using a friendly format, e.g., "Jan 2, 2006"
	user.DOBFormatted = user.DOB.Format("2006-01-02")

	return user, nil
}
func GetUserByEmail(db *sql.DB, email string) (models.User, error) {
	var user models.User

	err := db.QueryRow("SELECT id, email, password, name, category, dob, bio, avatar FROM user WHERE email = ?", email).Scan(&user.Id, &user.Email, &user.Password, &user.Name, &user.Category, &user.DOB, &user.Bio, &user.Avatar)

	return user, err
}
func CreateUser(db *sql.DB, user models.User) error {

	id, err := uuid.NewUUID()

	if err != nil {
		return err
	}

	//Convert id to string and set it on the user
	user.Id = id.String()

	stmt, err := db.Prepare(`INSERT INTO "user" (id, email, password, name, category, dob, bio, avatar) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`)


	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.Id, user.Email, user.Password, user.Name, user.Category, user.DOB, user.Bio, user.Avatar)

	if err != nil {
		return err
	}

	return nil
}
func UpdateUser(db *sql.DB, id string, user models.User) error {
	_, err := db.Exec("UPDATE user SET name = ?, category = ?, dob = ?, bio = ? WHERE id = ?", user.Name, user.Category, user.DOB, user.Bio, id)

	return err
}
func UpdateUserAvatar(db *sql.DB, userID, filePath string) error {
	_, err := db.Exec("UPDATE user SET avatar = ? WHERE id = ?", filePath, userID)
	return err
}
func DeleteUser(db *sql.DB, id string) error {
	_, err := db.Exec("DELETE FROM user WHERE id = ?", id)

	return err
}
