package app

import (
	"context"
	"fmt"
	"joust/data"
	"joust/db"
	"joust/handler"
	"joust/service"
	"joust/view"
	"net/http"
	"time"

	"dice"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"gorm.io/gorm"
)

type App struct {
	echo *echo.Echo
	db   *gorm.DB
}

func New() *App {
	app := &App{}

	gormDB, err := db.LoadWorldDB()
	if err != nil {
		panic(err)
	}

	app.db = gormDB

	app.routes()

	return app
}

func (a *App) routes() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	//e.Use(middlewar.BasicAuth(...))

	e.GET("/random", func(c echo.Context) error {
		result := dice.NewRoller().Dice(1, 4).Roll()
		return c.String(http.StatusOK, fmt.Sprintf("%d", result))
	})

	e.GET("/", func(c echo.Context) error {
		return view.IndexPage().Render(c.Request().Context(), c.Response().Writer)
	})
	e.File("/favicon.ico", "favicon.ico")

	a.userRoutes(*e.Group("/user"))
	a.knightRoutes(*e.Group("/knight"))
	a.matchRoutes(*e.Group("/match"))
	a.generationRoutes(*e.Group("/generate"))

	a.echo = e
}

func (a *App) userRoutes(g echo.Group) {
	userHandler := &handler.UserHandler{
		Repo: &data.UserRepo{
			DB: a.db,
		},
	}

	g.POST("/", userHandler.Create)
}

func (a *App) knightRoutes(g echo.Group) {
	knightHandler := &handler.KnightHandler{
		KnightRepo: &data.KnightRepo{
			DB: a.db,
		},
	}

	g.GET("/:id", knightHandler.ReadById)
}

func (a *App) matchRoutes(g echo.Group) {
	matchHandler := &handler.MatchHandler{
		MatchService: &service.MatchService{
			KnightRepo: &data.KnightRepo{
				DB: a.db,
			},
			HorseRepo: &data.HorseRepo{
				DB: a.db,
			},
			MatchRepo: &data.MatchRepo{
				DB: a.db,
			},
		},
	}

	g.POST("/", matchHandler.Create)
}

func (a *App) generationRoutes(g echo.Group) {
	generateHandler := &handler.GenerateHandler{
		GenerateService: &service.GenerateService{
			KnightRepo: &data.KnightRepo{
				DB: a.db,
			},
		},
	}

	g.POST("/knights", generateHandler.Knights)
}

func (a *App) Start(ctx context.Context) error {
	ch := make(chan error, 1)

	go func() {
		if err := a.echo.Start("127.0.0.1:8080"); err != nil {
			ch <- fmt.Errorf("failed to start server: %w", err)
		}
		close(ch)
	}()

	select {
	case err := <-ch:
		return err
	case <-ctx.Done():
		fmt.Println("starting graceful shutdown...")
		timeout, cancel := context.WithTimeout(context.Background(), time.Second*10)
		defer cancel()
		defer fmt.Println("Graceful Shutdown complete")
		if err := a.echo.Shutdown(timeout); err != nil {
			return err
		}
		return nil
	}

}
