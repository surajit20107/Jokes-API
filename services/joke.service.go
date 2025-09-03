package services

import (
	"main/db"
	"main/models"
)

func GetJokes() ([]models.Joke, error) {
	var jokes []models.Joke
	result := db.DB.Limit(5).Find(&jokes)
	if result.Error != nil {
		return nil, result.Error
	}
	return jokes, nil
}

func CreateJoke(joke models.Joke) (models.Joke, error) {
	result := db.DB.Create(&joke)
	if result.Error != nil {
		return models.Joke{}, result.Error
	}
	return joke, nil
}

func UpdateJoke(id uint, updateData models.Joke) (models.Joke, error) {
	var joke models.Joke

	result := db.DB.First(&joke, id)

	if result.Error != nil {
		return models.Joke{}, result.Error
	}

	joke.Content = updateData.Content
	joke.Author = updateData.Author
	newResult := db.DB.Save(&joke)
	if newResult.Error != nil {
		return models.Joke{}, newResult.Error
	}
	return joke, nil
}

func DeleteJoke(id uint) error {
  var joke models.Joke
  result := db.DB.Delete(&joke, id)
  if result.Error != nil {
    return result.Error
  }
  return nil
}

func GetMyJokes(id uint) ([]models.Joke, error) {
	var jokes []models.Joke
	
	result := db.DB.Where("author = ?", id).Find(&jokes)
	
	if result.Error != nil {
		return nil, result.Error
	}

	return jokes, nil
}
