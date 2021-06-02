package services

import (
	"net/http"

	"github.com/adhikag24/golang-microservices/mvc/domain"
	"github.com/adhikag24/golang-microservices/mvc/utils"
)

type itemsService struct{}

var (
	ItemService itemsService
)

func (s *itemsService) GetItem(itemId string) (*domain.Item, *utils.ApplicationError) {
	return nil, &utils.ApplicationError{
		Message:    "Implement Me",
		StatusCode: http.StatusInternalServerError,
	}
}
