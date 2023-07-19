package handler

import (
	"github.com/automated-pen-testing/api/pkg/client"
	"github.com/automated-pen-testing/api/pkg/models"
)

type Handler struct {
	Client *client.Client
	Models *models.Interface
}
