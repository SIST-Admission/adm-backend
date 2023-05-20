package models

type BasicDetails struct{}

func (BasicDetails) TableName() string {
	return "basic_details"
}
