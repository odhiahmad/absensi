package repository

import (
	"github.com/odhiahmad/apiabsen/entity"
	"gorm.io/gorm"
)

type AbsenRepository interface {
	InsertAbsen(absen entity.Absen) entity.Absen
	UpdateAbsen(absen entity.Absen) entity.Absen
	VerifyCredential(absenname string, password string) interface{}
	IsDuplicateAbsenname(absenname string) (tx *gorm.DB)
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
	absen.Password = hashAndSalt([]byte(absen.Password))
	db.connection.Save(&absen)

	return absen
}

func (db *absenConnection) UpdateAbsen(absen entity.Absen) entity.Absen {

	if absen.Password != "" {
		absen.Password = hashAndSalt([]byte(absen.Password))
	} else {
		var tempAbsen entity.Absen
		db.connection.Find(&tempAbsen, absen.ID)
		absen.Password = tempAbsen.Password
	}

	db.connection.Save(&absen)

	return absen
}

