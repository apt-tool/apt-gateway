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

	// create middleware
	mid := middleware.Middleware{
		JWTAuthenticator: authenticator,
		Models:           r.ModelsInterface,
		ErrHandler:       errHandler,
	}

	// create controller
	ctl := controller.Controller{
		Config:           r.Config,
		JWTAuthenticator: authenticator,
		Models:           r.ModelsInterface,
		ErrHandler:       errHandler,
		Client:           client.NewClient(),
	}

	// register endpoints

	// login endpoint
	app.Post("/login", ctl.Login)

	// add auth middleware
	auth := app.Use(mid.Auth)

	// user crud
	profile := auth.Group("/profile")
	profile.Get("/", ctl.GetProfile)
	profile.Post("/", ctl.UpdateProfile)

	// users crud
	users := auth.Use(mid.Admin).Group("/users")
	users.Get("/", ctl.GetUsersList)
	users.Post("/", ctl.CreateUser)
	users.Put("/:id", ctl.UpdateUser)
	users.Get("/:id", ctl.GetUser)
	users.Delete("/:id", ctl.DeleteUser)

	// namespaces crud
	namespaces := auth.Group("/namespaces")
	namespaces.Get("/", mid.Admin, ctl.GetNamespacesList)
	namespaces.Post("/", mid.Admin, ctl.CreateNamespace)
	namespaces.Put("/:id", mid.Admin, ctl.UpdateNamespace)
	namespaces.Get("/:id", mid.Admin, ctl.GetNamespace)
	namespaces.Delete("/:id", mid.Admin, ctl.DeleteNamespace)
	namespaces.Get("/user", ctl.GetUserNamespacesList)
	namespaces.Get("/user/:id", mid.UserNamespace, ctl.GetUserNamespace)

	// projects crud
	projects := auth.Group("/projects/:namespace_id")
	projects.Post("/", mid.UserNamespace, ctl.CreateProject)
	projects.Get("/:id", mid.UserNamespace, ctl.GetProject)
	projects.Post("/:id", mid.UserNamespace, ctl.ExecuteProject)
	projects.Delete("/:id", mid.UserNamespace, ctl.DeleteProject)
	projects.Get("/:id/:document_id", mid.UserNamespace, ctl.DownloadProjectDocument)
}
