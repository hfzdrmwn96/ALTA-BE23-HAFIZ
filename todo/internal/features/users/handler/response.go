package handler

type LoginResponse struct {
	Token string `json:"token"`
}

func ToLoginReponse(input string) LoginResponse {
	return LoginResponse{
		Token: input,
	}
}
