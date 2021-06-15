package queries

type FieldValue struct {
	Field string     `json:"field"`
	Value interface{}`json:"value"`
}

type EsQuery struct {
	Equals []FieldValue `json:"equals"`
}
