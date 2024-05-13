package web

const PENDING_CART = "PENDING"

type AuthData struct {
	Role   string `json:"role"`
	UserId string `json:"userId"`
}
