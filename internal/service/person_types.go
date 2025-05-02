package service

type CreatePersonReq struct {
	Name    string `json:"name"`
	Surname string `json:"surname"`
}

type UpdatePersonReq struct {
	Name        string `json:"name"`
	Surname     string `json:"surname"`
	Age         int    `json:"age"`
	Gender      string `json:"gender"`
	Nationality string `json:"nationality"`
}
