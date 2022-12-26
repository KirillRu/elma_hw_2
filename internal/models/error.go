package models

type Err struct {
	ErrNo   uint   `json:"err_no"`
	Message string `json:"message"`
}
