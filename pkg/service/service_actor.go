package service

import (
	"github.com/andyfilya/restapi"
	"github.com/andyfilya/restapi/pkg/repo"
)

type ActorService struct {
	repo repo.Actor
}

func InitActorService(repo repo.Actor) *ActorService {
	return &ActorService{
		repo: repo,
	}
}

func (as *ActorService) CreateActor(actor restapi.Actor) (int, error) {
	err := checkEmptyCreate(actor) // check if the fields of actor is empty
	if err != nil {
		return -1, err
	}
	return as.repo.CreateActor(actor)
}

func (as *ActorService) DeleteActor(actor restapi.Actor) error {
	return as.repo.DeleteActor(actor)
}

func (as *ActorService) ChangeActor(oldActor restapi.Actor, newActor restapi.Actor) error {
	oldActor, newActor = checkEmptyFields(oldActor, newActor) // check if the fields of new actor is empty
	return as.repo.ChangeActor(oldActor, newActor)
}

func (as *ActorService) FindActorFilm(actor string) ([]restapi.Film, error) {
	return as.repo.FindActorFilm(actor)
}
