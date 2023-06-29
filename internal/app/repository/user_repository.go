package apprepository

import "pronaces_back/internal/domain/entity"

type Form0Request struct {
	Integrante []entity.Table0 `json:"integrante"`
}
