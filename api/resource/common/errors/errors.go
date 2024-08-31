package errors

type Error struct {
	Error string `json:"error"`
}

type Errors struct {
	Errors []Error `json:"errors"`
}
