package handler

import (
	"net/http"

	"github.com/TimesCoder/movie-app/internal/service"
	"github.com/labstack/echo/v4"
)

type MovieHandler struct {
	movieService service.MovieService
}

func NewMovieHandler(movieService service.MovieService) MovieHandler {
	return MovieHandler{movieService}
}

func (h *MovieHandler) GetMovies(ctx echo.Context) error {
	users, err := h.movieService.GetAll(ctx.Request().Context())
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(http.StatusOK, users)
}
