package ports

import "github.com/BoruTamena/job_board/internal/domain/models/dto"

type UserRepository interface {
	CreateUser(user dto.User)
}
