package domain

import "time"

type LogsDB interface {
	Create(log *Logs) error
	FindAll() (*[]Logs, error)
	FindByID(id uint) (*Logs, error)
	FindByUserID(userID uint) (*[]Logs, error)
	Update(log *Logs) error
	Delete(log *Logs) error
}

type LogsService interface {
	FindByUserID(userID uint) ([]Logs, error)
	FindAll() ([]Logs, error)
}

type Logs struct {
	ID          uint      `json:"id" gorm:"primary_key"`
	UserID      uint      `json:"user_id" gorm:"not null;index:idx_user_id"`
	Description string    `json:"description" gorm:"not null;type:varchar(255)"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
