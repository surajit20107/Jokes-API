package services

import (
  "main/models"
  "errors"
)

var jokes = []models.Joke{
  {Id: 1, Content: "Why don’t skeletons fight each other? They don’t have the guts.", Author: "Dev1"},
  {Id: 2, Content: "I used to play piano by ear, but now I use my hands.", Author: "Dev2"},
  {Id: 3, Content: "I’m on a seafood diet — I see food and I eat it.", Author: "Dev2"},
  {Id: 4, Content: "I started out with nothing — and I still have most of it.", Author: "Dev3"},
  {Id: 5, Content: "I'm not lazy, I'm on energy-saving mode.", Author: "Dev1"},
}

var lastId = 5

func GetJokes() []models.Joke {
  return jokes
}

func GetJokesById(id int) (models.Joke, error) {
  for _, joke := range jokes {
    if joke.Id == id {
      return joke, nil
    }
  }
  return models.Joke{}, errors.New("Joke not found")
}

func CreateJoke(content, author string) models.Joke {
  lastId++
  newJoke := models.Joke{
    Id: lastId,
    Content: content,
    Author: author,
  }
  jokes = append(jokes, newJoke)
  return newJoke
}

func UpdateJoke(id int, newJoke models.Joke) (models.Joke, error) {
  for i, joke := range jokes {
    if joke.Id == id {
      jokes[i].Content = newJoke.Content
      jokes[i].Author = newJoke.Author
      return jokes[i], nil
    }
  }
  return models.Joke{}, errors.New("Joke not found")
}

func DeleteJoke(id int) error {
  for i, joke := range jokes {
    if joke.Id == id {
      jokes = append(jokes[:i], jokes[i+1:]...)
      return nil
    }
  }
  return errors.New("Joke not found")
}