package domain

type MovieList struct {
	ID          int     `json:"id,omitempty"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Rating      float64 `json:"rating"`
	Image       string  `json:"image"`
	CreatedAt   string  `json:"createdAt"`
	UpdatedAt   string  `json:"updatedAt"`
}

type MovieDomainRepository interface {
	Get() ([]MovieList, error)
	GetById(id int) (MovieList, error)
	Add(payload MovieList) error
	Update(payload MovieList) error
	Delete(id int) error
}
