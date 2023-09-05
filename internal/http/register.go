package http

import (
	"github.com/apt-tool/apt-gateway/internal/config"
	"github.com/apt-tool/apt-gateway/internal/http/controller"
	"github.com/apt-tool/apt-gateway/internal/http/controller/handler"
	"github.com/apt-tool/apt-gateway/internal/http/middleware"
	"github.com/apt-tool/apt-gateway/internal/utils/jwt"
	"github.com/apt-tool/apt-gateway/pkg/client"

	"github.com/apt-tool/apt-core/pkg/models"

	"github.com/gofiber/fiber/v2"
)

type Register struct {
	Config          config.Config
	ModelsInterface *models.Interface
}

func (r Register) Create(app *fiber.App) {
	// create new jwt authenticator
	authenticator := jwt.New(r.Config.JWT)

	// create an error handler for http service
	errHandler := handler.ErrorHandler{DevMode: r.Config.HTTP.DevMode}

	// create middleware and controller
	mid := middleware.Middleware{
		JWTAuthenticator: authenticator,
		Models:           r.ModelsInterface,
		ErrHandler:       errHandler,
	}
	ctl := controller.Controller{
		Config:           r.Config,
		JWTAuthenticator: authenticator,
		Models:           r.ModelsInterface,
		ErrHandler:       errHandler,
		Client:           client.NewClient(),
	}

	// register endpoints
	// login endpoint
	app.Post("/login", ctl.UserLogin)

	// add auth middleware
	auth := app.Use(mid.Auth)

	// user crud
	user := auth.Group("/user")
	user.Get("/", ctl.GetUser)
	user.Post("/", ctl.UpdateUser)

	// users crud
	users := auth.Use(mid.Admin).Group("/users")
	users.Get("/", ctl.GetUsersList)
	users.Post("/", ctl.UserRegister)
	users.Put("/", ctl.UpdateUserRole)
	users.Delete("/:id", ctl.DeleteUser)

	// namespaces crud
	namespaces := auth.Group("/namespaces")
	namespaces.Get("/", ctl.GetNamespaces)
	namespaces.Post("/", ctl.CreateNamespace)
	namespaces.Put("/", ctl.UpdateNamespace)
	namespaces.Get("/:namespace_id", ctl.GetNamespaceUsers)
	namespaces.Delete("/:namespace_id", ctl.DeleteNamespace)
	namespaces.Get("/user", ctl.GetUserNamespaces)
	namespaces.Get("/user/:namespace_id", mid.UserNamespace, ctl.GetNamespace)

	// projects crud
	projects := auth.Group("/projects/:namespace_id")
	projects.Post("/", mid.UserNamespace, ctl.CreateProject)
	projects.Get("/:id", mid.UserNamespace, ctl.GetProject)
	projects.Post("/:id", mid.UserNamespace, ctl.ExecuteProject)
	projects.Delete("/:id", mid.UserNamespace, ctl.DeleteProject)
	projects.Get("/:id/:document_id", mid.UserNamespace, ctl.DownloadProjectDocument)
}
