package service

import (
	"log"

	"github.com/mashingan/smapping"
	"github.com/odhiahmad/absensi/dto"
	"github.com/odhiahmad/absensi/entity"
	"github.com/odhiahmad/absensi/repository"
)

type AbsenService interface {
	CreateAbsen(absen dto.AbsenCreateDTO) entity.Absen
	UpdateAbsen(absen dto.AbsenUpdateDTO) entity.Absen
}

type absenService struct {
	absenRepository repository.AbsenRepository
}

func NewAbsenService(absenRepo repository.AbsenRepository) AbsenService {
	return &absenService{
		absenRepository: absenRepo,
	}
}

func (service *absenService) CreateAbsen(absen dto.AbsenCreateDTO) entity.Absen {
	absenToCreate := entity.Absen{}
	err := smapping.FillStruct(&absenToCreate, smapping.MapFields(&absen))
	if err != nil {
		log.Fatalf("Failed map %v:", err)
	}
	absenToCreate.Prepare()
	res := service.absenRepository.InsertAbsen((absenToCreate))
	return res
}

func (service *absenService) UpdateAbsen(absen dto.AbsenUpdateDTO) entity.Absen {
	absenToUpdate := entity.Absen{}
	err := smapping.FillStruct(&absenToUpdate, smapping.MapFields(&absen))
	if err != nil {
		log.Fatalf("Failed map %v:", err)
	}
	res := service.absenRepository.UpdateAbsen((absenToUpdate))
	return res
}


