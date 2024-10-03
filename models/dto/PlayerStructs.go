package dto

type PlayerAddRequest struct {
	Name  string
	Score int
}

type PlayerAddResponse struct {
	Id int
}
