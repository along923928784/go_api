package enum

type LoginType int32
type ArtType int32

const (
	ADMIN_EMAIL       LoginType = 200
	USER_MINI_PROGRAM LoginType = 99 + iota
	USER_EMAIL
	USER_MOBILE
)

const (
	MOVIE    ArtType = 100
	MUSIC    ArtType = 200
	SENTENCE ArtType = 300
	BOOK     ArtType = 400
)
