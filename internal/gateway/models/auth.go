package gateway_models

type CreateUser struct {
	Name string `json:"name"`
	Email string `json:"email"`
	Password string 'json:"password"'
}

type LoginUser struct {

}
