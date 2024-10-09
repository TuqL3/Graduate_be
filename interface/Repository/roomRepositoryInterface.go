package Repository

import (
	"server/dtos/room"
	"server/models"
)

type RoomRepositoryInterface interface {
	CreateRoom(createRoomDto *room.CreateRoomDto) (*models.Room, error)
	UpdateRoom(roomId int, dto room.UpdateRoomDto) (*models.Room, error)
	DeleteRoom(roomId int) error
	GetAllRooms() ([]*models.Room, error)
}
