package web

const PENDING_CART = "PENDING"
const SUCCESS_CART = "SUCCESS"

type AuthData struct {
	Role   string `json:"role"`
	UserId string `json:"userId"`
}
