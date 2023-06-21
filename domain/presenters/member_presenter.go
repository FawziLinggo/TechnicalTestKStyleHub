package presenters

type CreateMemberRequest struct {
	Username  string `json:"username" validate:"required"`
	Gender    string `json:"gender" validate:"required" example:"male"`
	SkinType  string `json:"skin_type" validate:"required"  example:"oily"`
	SkinColor string `json:"skin_color" validate:"required" example:"white"`
}

type EditMemberByIDRequest struct {
	ID        string `json:"id" validate:"required"`
	Username  string `json:"username" validate:"required"`
	Gender    string `json:"gender" validate:"required" example:"male"`
	SkinType  string `json:"skin_type" validate:"required"  example:"oily"`
	SkinColor string `json:"skin_color" validate:"required" example:"white"`
}

type DeleteMemberByIDRequest struct {
	ID string `json:"id" validate:"required"`
}
