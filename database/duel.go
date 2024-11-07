package database

import (
	"commitinder/models"
	"fmt"
)

func (d *Database) SaveDuel(duel *models.Duel) error {
	err := d.db.Create(&duel).Error
	if err != nil {
		return fmt.Errorf("error saving the duel: " + err.Error())
	}
	return nil
}
func (d *Database) GetDuel(duelId int, duel *models.Duel) error {
	err := d.db.Preload("Results").Preload("CommitMessages").First(&duel, "id = ?", duelId).Error
	if err != nil {
		return fmt.Errorf("error getting a duel, by id: " + err.Error())
	}
	return nil
}
func (d *Database) GetAllDuels(duels *[]models.Duel) error {
	err := d.db.Preload("Results").Preload("CommitMessages").Find(&duels).Error
	if err != nil {
		return fmt.Errorf("error getting all duels: " + err.Error())
	}
	return nil
}
func (d *Database) UpdateDuel(duel *models.Duel) error {
	err := d.db.Preload("Results").Preload("CommitMessages").Updates(&duel).Error
	if err != nil {
		return fmt.Errorf("error updating a duel: " + err.Error())
	}
	return nil
}
func (d *Database) DeleteDuel(duelId int) error {
	err := d.db.Delete(&models.Duel{}, "id = ?", duelId).Error
	if err != nil {
		return fmt.Errorf("error deleting the duel: " + err.Error())
	}
	return nil
}
