package requests

type CategoryRequest struct {
	Title string `json:"title" validate:"required,min=3,max=200"`
}
