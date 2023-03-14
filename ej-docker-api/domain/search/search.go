package search

type Query struct {
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
}

type Result struct {
	Results []interface{} `json:"results"`
	Paging  Paging        `json:"paging"`
}

type Paging struct {
	Total  int `json:"total"`
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
}
