package simple

type Database struct {
	Name string
}

type (
	DatabasePostgreSQL Database
	DatabaseMonggoDB   Database
)

func NewDatabasePostgreSQL() *DatabasePostgreSQL {
	return (*DatabasePostgreSQL)(&Database{Name: "PostgreSQL"})
}

func NewDatabaseMongoDB() *DatabaseMonggoDB {
	return (*DatabaseMonggoDB)(&Database{Name: "MongoDB"})
}

type DatabaseRepository struct {
	DatabasePostgreSQL *DatabasePostgreSQL
	DatabaseMonggoDB   *DatabaseMonggoDB
}

func NewDatabaseRepository(
	postgreSQL *DatabasePostgreSQL,
	mongoDB *DatabaseMonggoDB,
) *DatabaseRepository {
	return &DatabaseRepository{DatabasePostgreSQL: postgreSQL, DatabaseMonggoDB: mongoDB}
}
