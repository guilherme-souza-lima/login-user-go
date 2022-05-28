package entities

type Token struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Login     string `json:"login"`
	Email     string `json:"email"`
	Cellphone string `json:"cellphone"`
}
