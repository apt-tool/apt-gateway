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
	namespaces.Get("/", ctl.GetUserNamespaces)
	namespaces.Get("/:namespace_id", mid.UserNamespace, ctl.GetNamespace)
	namespaces.Get("/:namespace_id/projects/:project_id", mid.UserNamespace, ctl.GetProject)
	namespaces.Get("/:namespace_id/projects/:project_id/:document_id", mid.UserNamespace, ctl.DownloadProjectDocument)
	namespaces.Get("/", ctl.GetNamespaces)
	namespaces.Post("/", ctl.CreateNamespace)
	namespaces.Put("/", ctl.UpdateNamespace)
	namespaces.Get("/:namespace_id", ctl.GetNamespaceUsers)
	namespaces.Delete("/:namespace_id", ctl.DeleteNamespace)

	// projects crud
	projects := auth.Group("/projects")

	// viewer routes
	viewerRoutes := auth.Group("/")

	viewNamespace := viewerRoutes.Group("/namespaces")
	viewNamespace.Get("/", ctl.GetUserNamespaces)
	viewNamespace.Get("/:namespace_id", mid.UserNamespace, ctl.GetNamespace)
	viewNamespace.Get("/:namespace_id/projects/:project_id", mid.UserNamespace, ctl.GetProject)
	viewNamespace.Get("/:namespace_id/projects/:project_id/:document_id", mid.UserNamespace, ctl.DownloadProjectDocument)

	// user routes
	userRoutes := auth.Group("/user")

	userNamespace := userRoutes.Group("/namespaces")

	userNamespace.Post("/:namespace_id/projects", mid.UserNamespace, ctl.CreateProject)
	userNamespace.Post("/:namespace_id/projects/:project_id", mid.UserNamespace, ctl.ExecuteProject)
	userNamespace.Delete("/:namespace_id/projects/:project_id", mid.UserNamespace, ctl.DeleteProject)
}
