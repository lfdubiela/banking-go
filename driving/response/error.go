package response

type Error struct {
	Location    string `json:"location"`
	Description string `json:"description"`
}

type Errors struct {
	Errors []Error `json:"errors"`
}

func NewErrorResponse(e map[string]string) Errors {
	var errors []Error

	for l, d := range e {
		errors = append(errors, Error{l, d})
	}

	return Errors{errors}
}
