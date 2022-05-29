package request

type User struct {
	Name      string `json:"name"`
	Login     string `json:"login"`
	Password  string `json:"password"`
	Email     string `json:"email"`
	Cellphone string `json:"cellphone"`
}

type Login struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type Verify struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Login     string `json:"login"`
	Email     string `json:"email"`
	Cellphone string `json:"cellphone"`
	Token     string `json:"token"`
}
