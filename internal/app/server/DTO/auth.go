package DTO

type (
	RegisterUserDTO struct {
		Login    string `json:"login"`
		Password string `json:"password"`
	}

	LoginUserDTO struct {
		Login    string `json:"login"`
		Password string `json:"password"`
	}
)
