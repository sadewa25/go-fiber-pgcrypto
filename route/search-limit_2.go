package route

import (
	"app/utils"
	"database/sql"
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
)

func SearchLimitTwo(app *fiber.App, db *sql.DB) {
	base_64_4_key := os.Getenv("BASE_64_4")
	base_64_8_key := os.Getenv("BASE_64_8")
	base_64_16_key := os.Getenv("BASE_64_16")
	hex_4_key := os.Getenv("HEX_4")
	hex_8_key := os.Getenv("HEX_8")
	hex_16_key := os.Getenv("HEX_16")

	app.Get("/users-original-limit-50-x2", utils.HandleQuery(db, fmt.Sprintf(`SELECT
			du.ID,
			du.NAME,
			du.birth_date,
			du.ktp,
			du.gender,
			du.phone
		FROM
			auth.users_original_x2 du
		WHERE
			du."name" ILIKE '%%a%%'
			LIMIT 50`)))

	app.Get("/users-base-64-4-limit-50-x2", utils.HandleQuery(db, fmt.Sprintf(`SELECT
			u."id",
			pgp_sym_decrypt ( decode( u.NAME, 'hex' :: TEXT ), '%s' :: TEXT ) AS NAME,
			pgp_sym_decrypt ( decode( u.birth_date, 'hex' :: TEXT ), '%s' :: TEXT ) AS birth_date,
			pgp_sym_decrypt ( decode( u.ktp, 'hex' :: TEXT ), '%s' :: TEXT ) AS ktp,
			pgp_sym_decrypt ( decode( u.gender, 'hex' :: TEXT ), '%s' :: TEXT ) AS gender,
			pgp_sym_decrypt ( decode( u.phone, 'hex' :: TEXT ), '%s' :: TEXT ) AS phone
		FROM
			auth.users_base64_4_x2 u WHERE pgp_sym_decrypt ( decode( u.NAME, 'hex' :: TEXT ), '%s' :: TEXT )  ILIKE '%%a%%'
			LIMIT 50`, base_64_4_key, base_64_4_key, base_64_4_key, base_64_4_key, base_64_4_key, base_64_4_key)))

	app.Get("/users-base-64-8-limit-50-x2", utils.HandleQuery(db, fmt.Sprintf(`SELECT
			u."id",
			pgp_sym_decrypt ( decode( u.NAME, 'hex' :: TEXT ), '%s' :: TEXT ) AS NAME,
			pgp_sym_decrypt ( decode( u.birth_date, 'hex' :: TEXT ), '%s' :: TEXT ) AS birth_date,
			pgp_sym_decrypt ( decode( u.ktp, 'hex' :: TEXT ), '%s' :: TEXT ) AS ktp,
			pgp_sym_decrypt ( decode( u.gender, 'hex' :: TEXT ), '%s' :: TEXT ) AS gender,
			pgp_sym_decrypt ( decode( u.phone, 'hex' :: TEXT ), '%s' :: TEXT ) AS phone
		FROM
			auth.users_base64_8_x2 u WHERE pgp_sym_decrypt ( decode( u.NAME, 'hex' :: TEXT ), '%s' :: TEXT )  ILIKE '%%a%%'
			LIMIT 50`, base_64_8_key, base_64_8_key, base_64_8_key, base_64_8_key, base_64_8_key, base_64_8_key)))

	app.Get("/users-base-64-16-limit-50-x2", utils.HandleQuery(db, fmt.Sprintf(`SELECT
			u."id",
			pgp_sym_decrypt ( decode( u.NAME, 'hex' :: TEXT ), '%s' :: TEXT ) AS NAME,
			pgp_sym_decrypt ( decode( u.birth_date, 'hex' :: TEXT ), '%s' :: TEXT ) AS birth_date,
			pgp_sym_decrypt ( decode( u.ktp, 'hex' :: TEXT ), '%s' :: TEXT ) AS ktp,
			pgp_sym_decrypt ( decode( u.gender, 'hex' :: TEXT ), '%s' :: TEXT ) AS gender,
			pgp_sym_decrypt ( decode( u.phone, 'hex' :: TEXT ), '%s' :: TEXT ) AS phone
		FROM
			auth.users_base64_16_x2 u WHERE pgp_sym_decrypt ( decode( u.NAME, 'hex' :: TEXT ), '%s' :: TEXT )  ILIKE '%%a%%'
			LIMIT 50`, base_64_16_key, base_64_16_key, base_64_16_key, base_64_16_key, base_64_16_key, base_64_16_key)))

	app.Get("/users-hex-4-limit-50-x2", utils.HandleQuery(db, fmt.Sprintf(`SELECT
			u."id",
			pgp_sym_decrypt ( decode( u.NAME, 'hex' :: TEXT ), '%s' :: TEXT ) AS NAME,
			pgp_sym_decrypt ( decode( u.birth_date, 'hex' :: TEXT ), '%s' :: TEXT ) AS birth_date,
			pgp_sym_decrypt ( decode( u.ktp, 'hex' :: TEXT ), '%s' :: TEXT ) AS ktp,
			pgp_sym_decrypt ( decode( u.gender, 'hex' :: TEXT ), '%s' :: TEXT ) AS gender,
			pgp_sym_decrypt ( decode( u.phone, 'hex' :: TEXT ), '%s' :: TEXT ) AS phone
		FROM
			auth.users_hex_4_x2 u WHERE pgp_sym_decrypt ( decode( u.NAME, 'hex' :: TEXT ), '%s' :: TEXT )  ILIKE '%%a%%'
			LIMIT 50`, hex_4_key, hex_4_key, hex_4_key, hex_4_key, hex_4_key, hex_4_key)))

	app.Get("/users-hex-8-limit-50-x2", utils.HandleQuery(db, fmt.Sprintf(`SELECT
			u."id",
			pgp_sym_decrypt ( decode( u.NAME, 'hex' :: TEXT ), '%s' :: TEXT ) AS NAME,
			pgp_sym_decrypt ( decode( u.birth_date, 'hex' :: TEXT ), '%s' :: TEXT ) AS birth_date,
			pgp_sym_decrypt ( decode( u.ktp, 'hex' :: TEXT ), '%s' :: TEXT ) AS ktp,
			pgp_sym_decrypt ( decode( u.gender, 'hex' :: TEXT ), '%s' :: TEXT ) AS gender,
			pgp_sym_decrypt ( decode( u.phone, 'hex' :: TEXT ), '%s' :: TEXT ) AS phone
		FROM
			auth.users_hex_8_x2 u WHERE pgp_sym_decrypt ( decode( u.NAME, 'hex' :: TEXT ), '%s' :: TEXT )  ILIKE '%%a%%'
			LIMIT 50`, hex_8_key, hex_8_key, hex_8_key, hex_8_key, hex_8_key, hex_8_key)))

	app.Get("/users-hex-16-limit-50-x2", utils.HandleQuery(db, fmt.Sprintf(`SELECT
			u."id",
			pgp_sym_decrypt ( decode( u.NAME, 'hex' :: TEXT ), '%s' :: TEXT ) AS NAME,
			pgp_sym_decrypt ( decode( u.birth_date, 'hex' :: TEXT ), '%s' :: TEXT ) AS birth_date,
			pgp_sym_decrypt ( decode( u.ktp, 'hex' :: TEXT ), '%s' :: TEXT ) AS ktp,
			pgp_sym_decrypt ( decode( u.gender, 'hex' :: TEXT ), '%s' :: TEXT ) AS gender,
			pgp_sym_decrypt ( decode( u.phone, 'hex' :: TEXT ), '%s' :: TEXT ) AS phone
		FROM
			auth.users_hex_16_x2 u WHERE pgp_sym_decrypt ( decode( u.NAME, 'hex' :: TEXT ), '%s' :: TEXT )  ILIKE '%%a%%'
			LIMIT 50`, hex_16_key, hex_16_key, hex_16_key, hex_16_key, hex_16_key, hex_16_key)))

}
