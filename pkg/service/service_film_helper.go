package service

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
