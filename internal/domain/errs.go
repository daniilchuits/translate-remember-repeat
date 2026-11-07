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
)
