package repositories

import (
	"time"

	"github.com/claudiano/dotz-backend/domain/model"
)

type PontoRepository struct{}

func (p *PontoRepository) FindByUserAndDate(idUser int32, dateReference time.Time) ([]model.Ponto, error) {
	return []model.Ponto{}, nil
}

func (p *PontoRepository) Save(ponto model.Ponto) (model.Ponto, error) {
	return model.Ponto{}, nil
}

func (p *PontoRepository) FindAll() ([]model.Ponto, error) {
	return []model.Ponto{}, nil
}
