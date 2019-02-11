package datatype


type Parse struct{
	Url string
	Data string
	File string
	Spider bool
}

type Configure struct {
	Input    Parse
	Scheme   string
	Host     string
	Path     string
	Username string
	Password string
	Dict     string
	Results  *Result
}

type Result struct {
	Page     string
	Param    string
	Payload  string
	Weakness bool
}
