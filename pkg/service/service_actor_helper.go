package service

import (
	"errors"
	"github.com/andyfilya/restapi"
	"github.com/sirupsen/logrus"
)

func checkEmptyFields(oldActor restapi.Actor, newActor restapi.Actor) (restapi.Actor, restapi.Actor) {
	logrus.Infof("oldActor : %v \nnewActor :  %v \n", oldActor, newActor)
	if newActor.FirstName == "" {
		newActor.FirstName = oldActor.FirstName
	}
	if newActor.LastName == "" {
		newActor.LastName = oldActor.LastName
	}
	if newActor.DateBirth == "" {
		newActor.DateBirth = oldActor.DateBirth
	}
	logrus.Infof("after check empty fields ... \n oldActor : %v \nnewActor :  %v \n", oldActor, newActor)

	return oldActor, newActor
}

func checkEmptyCreate(actor restapi.Actor) error {
	switch {
	case actor.FirstName == "":
		return errors.New("empty name actor")
	case actor.LastName == "":
		return errors.New("empty surname actor")
	case actor.DateBirth == "":
		return errors.New("empty date birth actor")
	}
	return nil
}
