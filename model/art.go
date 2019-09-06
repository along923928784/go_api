package model

import (
	"jiyue.im/pkg/enum"
)

func GetData(art_id, art_type int32) (Art, bool) {
	var art Art
	switch enum.ArtType(art_type) {
	case enum.MOVIE:
		art = &MovieModel{}
	case enum.MUSIC:
		art = &MusicModel{}
	case enum.SENTENCE:
		art = &SentenceModel{}
	default:
		return art, false
	}
	isNotFound := art.FindOne(art_id)
	return art, isNotFound

}
