package products

type Product struct {
	ID          string `json:"_id, omitempty"`
	Name        string             `json:"name"`
	Price       int32              `json:"price"`
	Description string             `json:"description"`
}

