package utils

import (
	"app/model"
	"database/sql"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func HandleQuery(db *sql.DB, query string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		rows, err := db.Query(query)
		if err != nil {
			return c.Status(500).SendString(fmt.Sprintf("Database query failed: %v", err))
		}
		defer rows.Close()

		var users []model.User
		for rows.Next() {
			var user model.User
			err := rows.Scan(&user.ID, &user.Name, &user.BirthDate, &user.KTP, &user.Gender, &user.Phone)
			if err != nil {
				return c.Status(500).SendString(fmt.Sprintf("Failed to scan row: %v", err))
			}
			users = append(users, user)
		}

		if err := rows.Err(); err != nil {
			return c.Status(500).SendString(fmt.Sprintf("Row iteration error: %v", err))
		}

		return c.JSON(users)
	}
}
