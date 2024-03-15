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
	return as.repo.CreateActor(actor)
}

func (as *ActorService) DeleteActor(actor restapi.Actor) error {
	return as.repo.DeleteActor(actor)
}

func (as *ActorService) ChangeActor(actorId int, toChange string) error {
	return as.repo.ChangeActor(actorId, toChange)
}

func (as *ActorService) FindActorFilm(actor string) ([]restapi.Film, error) {
	return as.repo.FindActorFilm(actor)
}
