package route

import (
	"app/utils"
	"database/sql"
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
)

func BaseRoutes(app *fiber.App, db *sql.DB) {
	base_64_4_key := os.Getenv("BASE_64_4")
	base_64_8_key := os.Getenv("BASE_64_8")
	base_64_16_key := os.Getenv("BASE_64_16")
	hex_4_key := os.Getenv("HEX_4")
	hex_8_key := os.Getenv("HEX_8")
	hex_16_key := os.Getenv("HEX_16")
	encryptedOriginal := os.Getenv("ENCRYPT_ORIGINAL")

	app.Get("/users-original", utils.HandleQuery(db, `SELECT
                users_original.id, 
                users_original.name, 
                users_original.birth_date, 
                users_original.ktp, 
                users_original.gender, 
                users_original.phone
            FROM
                auth.users_original`))

	app.Get("/users-encrypted", utils.HandleQuery(db, fmt.Sprintf(`SELECT
					u."id",
					pgp_sym_decrypt ( decode( u.NAME, 'hex' :: TEXT ), '%s' :: TEXT ) AS name,
					pgp_sym_decrypt ( decode( u.birth_date, 'hex' :: TEXT ), '%s' :: TEXT ) AS birth_date,
					pgp_sym_decrypt ( decode( u.ktp, 'hex' :: TEXT ), '%s' :: TEXT ) AS ktp,
					pgp_sym_decrypt ( decode( u.gender, 'hex' :: TEXT ), '%s' :: TEXT ) AS gender,
					pgp_sym_decrypt ( decode( u.phone, 'hex' :: TEXT ), '%s' :: TEXT ) AS phone
				FROM
					auth.users u`, encryptedOriginal, encryptedOriginal, encryptedOriginal, encryptedOriginal, encryptedOriginal)))

	app.Get("/users-base-64-4", utils.HandleQuery(db, fmt.Sprintf(`SELECT
		u."id",
		pgp_sym_decrypt ( decode( u.NAME, 'hex' :: TEXT ), '%s' :: TEXT ) AS name,
		pgp_sym_decrypt ( decode( u.birth_date, 'hex' :: TEXT ), '%s' :: TEXT ) AS birth_date,
		pgp_sym_decrypt ( decode( u.ktp, 'hex' :: TEXT ), '%s' :: TEXT ) AS ktp,
		pgp_sym_decrypt ( decode( u.gender, 'hex' :: TEXT ), '%s' :: TEXT ) AS gender,
		pgp_sym_decrypt ( decode( u.phone, 'hex' :: TEXT ), '%s' :: TEXT ) AS phone
	FROM
		auth.users_base64_4 u`, base_64_4_key, base_64_4_key, base_64_4_key, base_64_4_key, base_64_4_key)))

	app.Get("/users-base-64-8", utils.HandleQuery(db, fmt.Sprintf(`SELECT
		u."id",
		pgp_sym_decrypt ( decode( u.NAME, 'hex' :: TEXT ), '%s' :: TEXT ) AS name,
		pgp_sym_decrypt ( decode( u.birth_date, 'hex' :: TEXT ), '%s' :: TEXT ) AS birth_date,
		pgp_sym_decrypt ( decode( u.ktp, 'hex' :: TEXT ), '%s' :: TEXT ) AS ktp,
		pgp_sym_decrypt ( decode( u.gender, 'hex' :: TEXT ), '%s' :: TEXT ) AS gender,
		pgp_sym_decrypt ( decode( u.phone, 'hex' :: TEXT ), '%s' :: TEXT ) AS phone
	FROM
		auth.users_base64_8 u`, base_64_8_key, base_64_8_key, base_64_8_key, base_64_8_key, base_64_8_key)))

	app.Get("/users-base-64-16", utils.HandleQuery(db, fmt.Sprintf(`SELECT
		u."id",
		pgp_sym_decrypt ( decode( u.NAME, 'hex' :: TEXT ), '%s' :: TEXT ) AS name,
		pgp_sym_decrypt ( decode( u.birth_date, 'hex' :: TEXT ), '%s' :: TEXT ) AS birth_date,
		pgp_sym_decrypt ( decode( u.ktp, 'hex' :: TEXT ), '%s' :: TEXT ) AS ktp,
		pgp_sym_decrypt ( decode( u.gender, 'hex' :: TEXT ), '%s' :: TEXT ) AS gender,
		pgp_sym_decrypt ( decode( u.phone, 'hex' :: TEXT ), '%s' :: TEXT ) AS phone
	FROM
		auth.users_base64_16 u`, base_64_16_key, base_64_16_key, base_64_16_key, base_64_16_key, base_64_16_key)))

	app.Get("/users-hex-4", utils.HandleQuery(db, fmt.Sprintf(`SELECT
		u."id",
		pgp_sym_decrypt ( decode( u.NAME, 'hex' :: TEXT ), '%s' :: TEXT ) AS name,
		pgp_sym_decrypt ( decode( u.birth_date, 'hex' :: TEXT ), '%s' :: TEXT ) AS birth_date,
		pgp_sym_decrypt ( decode( u.ktp, 'hex' :: TEXT ), '%s' :: TEXT ) AS ktp,
		pgp_sym_decrypt ( decode( u.gender, 'hex' :: TEXT ), '%s' :: TEXT ) AS gender,
		pgp_sym_decrypt ( decode( u.phone, 'hex' :: TEXT ), '%s' :: TEXT ) AS phone
	FROM
		auth.users_hex_4 u`, hex_4_key, hex_4_key, hex_4_key, hex_4_key, hex_4_key)))

	app.Get("/users-hex-8", utils.HandleQuery(db, fmt.Sprintf(`SELECT
		u."id",
		pgp_sym_decrypt ( decode( u.NAME, 'hex' :: TEXT ), '%s' :: TEXT ) AS name,
		pgp_sym_decrypt ( decode( u.birth_date, 'hex' :: TEXT ), '%s' :: TEXT ) AS birth_date,
		pgp_sym_decrypt ( decode( u.ktp, 'hex' :: TEXT ), '%s' :: TEXT ) AS ktp,
		pgp_sym_decrypt ( decode( u.gender, 'hex' :: TEXT ), '%s' :: TEXT ) AS gender,
		pgp_sym_decrypt ( decode( u.phone, 'hex' :: TEXT ), '%s' :: TEXT ) AS phone
	FROM
		auth.users_hex_8 u`, hex_8_key, hex_8_key, hex_8_key, hex_8_key, hex_8_key)))

	app.Get("/users-hex-16", utils.HandleQuery(db, fmt.Sprintf(`SELECT
		u."id",
		pgp_sym_decrypt ( decode( u.NAME, 'hex' :: TEXT ), '%s' :: TEXT ) AS name,
		pgp_sym_decrypt ( decode( u.birth_date, 'hex' :: TEXT ), '%s' :: TEXT ) AS birth_date,
		pgp_sym_decrypt ( decode( u.ktp, 'hex' :: TEXT ), '%s' :: TEXT ) AS ktp,
		pgp_sym_decrypt ( decode( u.gender, 'hex' :: TEXT ), '%s' :: TEXT ) AS gender,
		pgp_sym_decrypt ( decode( u.phone, 'hex' :: TEXT ), '%s' :: TEXT ) AS phone
	FROM
		auth.users_hex_16 u`, hex_16_key, hex_16_key, hex_16_key, hex_16_key, hex_16_key)))

}
