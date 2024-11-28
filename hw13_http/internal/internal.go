package internal

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type Order struct {
	ID          int     `json:"id"`
	UserID      string  `json:"userId"`
	OrderDate   string  `json:"orderDate"`
	TotalAmount float64 `json:"totalAmount"`
}

type Product struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}
