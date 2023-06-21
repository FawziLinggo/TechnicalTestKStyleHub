package presenters

type GetMemberResponse struct {
	ID        string `json:"id" example:"c836c745-e1db-4777-9699-e40fffaa302d"`
	Username  string `json:"username" example:"fawzi"`
	Gender    string `json:"gender" example:"male"`
	SkinType  string `json:"skin_type" example:"oily"`
	SkinColor string `json:"skin_color" example:"white"`
}
