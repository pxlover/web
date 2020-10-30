package response

type ShowValues struct {
	Table		[]Item		`json:"table"`
}

type Item struct {
	Name 		string		`json:"name"`
}