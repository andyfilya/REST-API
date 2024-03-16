package service

import (
	"github.com/andyfilya/restapi"
	"github.com/sirupsen/logrus"
)

func titleFilmCheck(title string) bool {
	if len(title) >= 1 && len(title) <= 150 {
		return true
	}
	return false
}

func descriptionFilmCheck(descroption string) bool {
	if len(descroption) >= 0 && len(descroption) <= 1000 {
		return true
	}
	return false
}

func checkEmptyFilm(newFilm restapi.Film, oldFilm restapi.Film) (restapi.Film, restapi.Film) {
	logrus.Infof("oldFilm : %v \nnewFilm :  %v \n", oldFilm, newFilm)
	if newFilm.Title == "" {
		newFilm.Title = oldFilm.Title
	}
	if newFilm.Date == "" {
		newFilm.Date = oldFilm.Date
	}
	if newFilm.Description == "" {
		newFilm.Description = oldFilm.Description
	}
	if newFilm.Rate == "" {
		newFilm.Rate = oldFilm.Rate
	}

	logrus.Infof("after check empty fields ... \n oldFilm : %v \nnewFilm :  %v \n", oldFilm, newFilm)

	return oldFilm, newFilm
}
