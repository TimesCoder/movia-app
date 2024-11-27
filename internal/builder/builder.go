package builder

import (
	"github.com/TimesCoder/movie-app/internal/http/handler"
	"github.com/TimesCoder/movie-app/internal/http/router"
	"github.com/TimesCoder/movie-app/internal/repository"
	"github.com/TimesCoder/movie-app/internal/service"
	"github.com/TimesCoder/movie-app/pkg/route"
	"gorm.io/gorm"
)

func BuildPublicRoutes(db *gorm.DB) []route.Route {
	// repository
	userRepository := repository.NewUserRepository(db)
	movieRepository := repository.NewMovieRepository(db)
	// end

	// service
	_ = service.NewUserService(userRepository)
	movieService := service.NewMovieService(movieRepository)
	// end

	// handler
	movieHandler := handler.NewMovieHandler(movieService)
	// end

	return router.PublicRoutes(movieHandler)

}

func BuildPrivateRoutes(db *gorm.DB) []route.Route {
	_ = repository.NewUserRepository(db)
	_ = repository.NewMovieRepository(db)
	return nil
}

func BuildRoutes(db *gorm.DB) []route.Route {
	_ = BuildPublicRoutes(db)
	_ = BuildPrivateRoutes(db)
	return router.PrivateRoutes()
}
