package recordtype

// Cfg is the struct type that contains fields that stores the necessary configuration
// gathered from the environment.
var Cfg struct {
	DBUser   string `envconfig:"DB_USER" default:"locuser"`      //  database's username
	DBPass   string `envconfig:"DB_PASS" default:"root"`         // database's password
	DBName   string `envconfig:"DB_NAME" default:"locationdb"`   // database's name
	DBHost   string `envconfig:"DB_HOST" default:"localhost"`    // host's ip
	DBPort   int    `envconfig:"DB_PORT" default:"5432"`         // host's port
	Locale   string `envconfig:"LOCALE"  default:"en"`           // locale
	GRPCHost string `envconfig:"GRPC_HOST"  default:"localhost"` // GRPC server host
}
