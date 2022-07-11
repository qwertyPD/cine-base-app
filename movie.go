package cine_base_app

type MovieList struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
}

type UserList struct {
	Id     int
	UserId int
	ListId int
}

type MovieItem struct {
	Id      int  `json:"id"`
	Title   int  `json:"title"`
	Watched bool `json:"watched"`
}

type ListItem struct {
	Id       int
	ListItem int
	ItemId   int
}
