package handler

import (
	"github.com/andyfilya/restapi"
	"sort"
	"strconv"
	"time"
)

func sortByTitle(films []restapi.Film) {
	sort.Slice(films, func(i, j int) bool {
		return films[i].Title < films[j].Title
	})
}

func sortByRate(films []restapi.Film) {
	sort.Slice(films, func(i, j int) bool {
		first, _ := strconv.ParseFloat(films[i].Rate, 64)
		second, _ := strconv.ParseFloat(films[j].Rate, 64)
		return first < second
	})
}

func sortByDate(films []restapi.Film) {
	sort.Slice(films, func(i, j int) bool {
		t1, _ := time.Parse(films[i].Date, "1978-01-01")
		t2, _ := time.Parse(films[j].Date, "1978-01-01")

		return t1.Unix() < t2.Unix()
	})
}
