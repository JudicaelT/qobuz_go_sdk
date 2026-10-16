package qobuzsdk

type User struct {
	Id    int    `json:"id"`
	Login string `json:"login"`
}

type UserPlaylists struct {
	User      User                 `json:"user"`
	Playlists Pagination[Playlist] `json:"playlists"`
}
