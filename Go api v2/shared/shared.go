package shared

var Config = configuration{

	POSTGRESURL: "host=localhost user=postgres password=1144 port=5432 dbname=fiberapi sslmode=disable",
}

type configuration struct {
	POSTGRESURL string
}
