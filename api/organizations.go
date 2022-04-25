package api

type Organization struct {
	ID    int32  `json:"id"`
	Login string `json:"login"`
	Name  string `json:"name"`
}
