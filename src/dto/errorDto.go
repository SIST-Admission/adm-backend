package dto

type Error struct {
	Code    int         `json:"code"`
	Message interface{} `json:"message"`
}
