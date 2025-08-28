package handlers

import "github.com/jmechavez/my-hexagonal-app/internal/core/ports"

type templateHandlers struct {
	userService ports.UserService
}

func NewTemplateHandlers(userService ports.UserService) *templateHandlers {
	return &templateHandlers{
		userService: userService,
	}
}
