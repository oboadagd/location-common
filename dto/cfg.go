package dto

// Cfg is the struct type that contains fields that stores the necessary configuration
// gathered from the environment.
var Cfg struct {
	DBUser string `envconfig:"DB_USER" default:"locuser"`
	DBPass string `envconfig:"DB_PASS" default:"root"`
	DBName string `envconfig:"DB_NAME" default:"locationdb"`
	DBHost string `envconfig:"DB_HOST" default:"localhost"`
	DBPort int    `envconfig:"DB_PORT" default:"5432"`
	Locale string `envconfig:"LOCALE"  default:"en"`
}
