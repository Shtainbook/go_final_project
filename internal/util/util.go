package util

import "encoding/json"

const (
	DateFormat        = "20060102"
	SearchDateFormat  = "02.01.2006"
	TaskListRowsLimit = 50
)

type errJson struct {
	Error string `json:"error"`
}

func MarshalError(err error) []byte {

	res, _ := json.Marshal(errJson{Error: err.Error()})
	return res
}
