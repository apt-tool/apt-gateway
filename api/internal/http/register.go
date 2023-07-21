package http

import (
	"github.com/automated-pen-testing/api/internal/config"
	"github.com/automated-pen-testing/api/internal/http/controller"
	"github.com/automated-pen-testing/api/internal/http/controller/handler"
	"github.com/automated-pen-testing/api/internal/http/middleware"
	"github.com/automated-pen-testing/api/internal/storage/redis"
	"github.com/automated-pen-testing/api/internal/utils/jwt"
	"github.com/automated-pen-testing/api/pkg/models"

	"github.com/gofiber/fiber/v2"
)

type Register struct {
	Config          config.Config
	RedisConnector  redis.Connector
	ModelsInterface *models.Interface
}

func (r Register) Create(app *fiber.App) {
	// create new jwt authenticator
	authenticator := jwt.New(r.Config.JWT)

	errHandler := handler.ErrorHandler{DevMode: r.Config.HTTP.DevMode}

	// create middleware and controller
	mid := middleware.Middleware{
		JWTAuthenticator: authenticator,
		Models:           r.ModelsInterface,
		RedisConnector:   r.RedisConnector,
		ErrHandler:       errHandler,
	}
	ctl := controller.Controller{
		JWTAuthenticator: authenticator,
		Models:           r.ModelsInterface,
		RedisConnector:   r.RedisConnector,
		ErrHandler:       errHandler,
	}

	// register endpoints
	app.Post("/login", ctl.UserLogin)

	auth := app.Use(mid.Auth)

	// viewer routes
	viewerRoutes := auth.Group("/")

	viewerRoutes.Get("/namespaces", ctl.GetUserNamespaces)
	viewerRoutes.Get("/namespaces/:namespace_id/projects", mid.UserNamespace)                          // get namespace projects
	viewerRoutes.Get("/namespaces/:namespace_id/projects/:project_id", mid.UserNamespace)              // get a project
	viewerRoutes.Get("/namespaces/:namespace_id/projects/:project_id/:document_id", mid.UserNamespace) // download document file

	// user routes
	userRoutes := auth.Group("/user")

	userRoutes.Get("/profile")                                                             // get user profile
	userRoutes.Post("/profile")                                                            // update profile
	userRoutes.Post("/namespaces/:namespace_id/projects", mid.UserNamespace)               // create project
	userRoutes.Post("/namespaces/:namespace_id/projects/:project_id", mid.UserNamespace)   // execute project
	userRoutes.Delete("/namespaces/:namespace_id/projects/:project_id", mid.UserNamespace) // delete project

	// admin routes
	adminRoutes := auth.Use(mid.Admin).Group("/admin")

	users := adminRoutes.Group("/users")

	users.Get("/") // get all users
	users.Post("/", ctl.UserRegister)

	namespaces := adminRoutes.Group("/namespaces")

	namespaces.Get("/", ctl.GetNamespaces)
	namespaces.Post("/", ctl.CreateNamespace)
	namespaces.Put("/", ctl.UpdateNamespace)
	namespaces.Delete("/:id", ctl.DeleteNamespace)

	instructions := adminRoutes.Group("/instructions")

	instructions.Get("/")      // get all instructions
	instructions.Post("/")     // create new instruction
	instructions.Delete(":id") // delete instruction
}
