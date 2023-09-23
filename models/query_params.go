package models

type QueryParams struct {
	Table   string
	Filters map[string]interface{}
}
