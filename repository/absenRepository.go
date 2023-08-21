package repository

import (
	"github.com/odhiahmad/absensi/entity"
	"gorm.io/gorm"
)

type AbsenRepository interface {
	InsertAbsen(absen entity.Absen) entity.Absen
	UpdateAbsen(absen entity.Absen) entity.Absen
}

type absenConnection struct {
	connection *gorm.DB
}

func NewAbsenRepository(db *gorm.DB) AbsenRepository {
	return &absenConnection{
		connection: db,
	}
}

func (db *absenConnection) InsertAbsen(absen entity.Absen) entity.Absen {
	db.connection.Save(&absen)

	return absen
}

func (db *absenConnection) UpdateAbsen(absen entity.Absen) entity.Absen {

	
		var tempAbsen entity.Absen
		db.connection.Find(&tempAbsen, absen.ID)
	

	db.connection.Save(&absen)

	return absen
}

