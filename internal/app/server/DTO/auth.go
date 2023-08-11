package DTO

import "time"

type (
	RegisterUserDTO struct {
		Login    string `json:"login"`
		Password string `json:"password"`
	}

	LoginUserDTO struct {
		Login    string `json:"login"`
		Password string `json:"password"`
	}

	AuthAuditDTO struct {
		Timestamp time.Time
		Event     string
	}
)
