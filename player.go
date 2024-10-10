package main

import (
	"errors"
	"gorm.io/gorm"
	"log"
)

func GetAllPlayers() []Player {
	var players []Player
	result := BotDB.Table("players").Find(&players)
	if result.Error != nil {
		log.Fatal(result.Error)
	}
	return players
}

func GetPlayerByLogin(login string) (Player, error) {
	var player Player
	result := BotDB.Where("login = ?", login).First(&player)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return player, &NoSuchPlayerError{login}
		}
		return player, result.Error
	}
	return player, nil
}

func CreatePlayer(login, password string) (Player, error) {
	existingPlayer, err := GetPlayerByLogin(login)
	if err == nil {
		return existingPlayer, &PlayerExistsError{Login: login}
	}

	if !errors.As(err, &ErrNoSuchPlayer) {
		return Player{}, err
	}

	_, err = GetUserByUsername(login)

	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return Player{}, err
		}
		return Player{}, &NoSuchUserError{Login: login}
	}

	newPlayer := Player{
		Login:    login,
		Password: password,
	}

	result := BotDB.Create(&newPlayer)
	if result.Error != nil {
		return Player{}, result.Error
	}

	return newPlayer, nil
}

// SetPlayerScore returns current player's score
func SetPlayerScore(login string, newScore uint) (uint, error) {

	var player Player

	result := BotDB.Where("login = ?", login).First(&player)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return 0, &NoSuchPlayerError{login}
		}
		return 0, result.Error
	}

	needsUpdate := newScore > player.Score

	if needsUpdate {
		player.Score = newScore

		result = BotDB.Save(&player)
		if result.Error != nil {
			return 0, result.Error
		}
	}

	return player.Score, nil
}
