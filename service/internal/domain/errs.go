package domain

import "errors"

var (

	// shutdown
	Shutdown error = errors.New("Grasefull shutdown")

	// json
	Marshaling error = errors.New("Error marshaling")
	Decoding   error = errors.New("Error decoding")

	// env in init
	EmptyVariableInEnv error = errors.New("Empty variable in '.env'")

	// request
	CreatingRequest error = errors.New("Error creating request")
	ClientDo        error = errors.New("Error the client made a poor request")
	StatusCode      error = errors.New("Wrong status code")

	// cli
	ReadingCLI error = errors.New("Error reading cli message")

	// psql
	NoRowsFound   error = errors.New("Error: no rows are found")
	ReadingPSQL   error = errors.New("Error reading psql's data")
	InsertingPSQL error = errors.New("Error inserting into 'translations'")

	// redis
	SettingRedis error = errors.New("Error setting in redis")

	// test
	WrongType error = errors.New("Wrong type of sentence: not postgres nor redis")
)
