package http

import (
	// "net/http"
	// "strconv"

	repository "github.com/Game-as-a-Service/The-Message/service/repository"
	// "github.com/gin-gonic/gin"
)

type PlayerHandler struct {
	GameRepo repository.GameRepository
}
