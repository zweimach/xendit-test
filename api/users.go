package api

type User struct {
	ID        int32   `json:"id"`
	Login     string  `json:"login"`
	Name      string  `json:"name"`
	Followers int32   `json:"followers"`
	Follows   int32   `json:"follows"`
	AvatarUrl *string `json:"avatar_url"`
}
