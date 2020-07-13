package repositories

import (
	"log"
	"time"

	"github.com/claudiano/dotz-backend/domain/model"
	"github.com/claudiano/dotz-backend/infra/database"
)

type PontoRepository struct{}

func (p *PontoRepository) FindByUserAndBetweenDateStartAndDateFinal(idUser int32, dateStart time.Time, dateFinal time.Time) ([]model.Ponto, error) {
	db := database.ConnectDB()
	var points []model.Ponto

	err := db.Where("IDUsuario = ? and DataHora BETWEEN ? AND ?", idUser, dateStart, dateFinal).Find(&points).Error
	if err != nil {
		log.Printf("Erro ao buscar pontos: %v", err)
	}
	return points, err
}

func (p *PontoRepository) Save(point model.Ponto) (model.Ponto, error) {
	db := database.ConnectDB()
	var err error

	if point.IDPonto != 0 {
		err = db.Save(&point).Error
	} else {
		err = db.Create(&point).Error
	}
	return point, err
}
