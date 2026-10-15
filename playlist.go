package qobuzsdk

type Playlist struct {
	Id              int           `json:"id"`
	Slug            string        `json:"slug"`
	Name            string        `json:"name"`
	Description     string        `json:"description"`
	CreatedAt       int           `json:"created_at"`
	UpdatedAt       int           `json:"updated_at"`
	Duration        int           `json:"duration"`
	TracksCount     int           `json:"tracks_count"`
	UsersCount      int           `json:"users_count"`
	IsPublic        bool          `json:"is_public"`
	IsCollaborative bool          `json:"is_collaborative"`
	IsFeatured      bool          `json:"is_featured"`
	Owner           PlaylistOwner `json:"owner"`
	Images          []string      `json:"images"`
	Images150       []string      `json:"images150"`
	Images300       []string      `json:"images300"`
	Genres          []Genre       `json:"genres"`
}

type CreatedPlaylist struct {
	Id              int           `json:"id"`
	Name            string        `json:"name"`
	Description     string        `json:"description"`
	TracksCount     int           `json:"tracks_count"`
	UsersCount      int           `json:"users_count"`
	Duration        int           `json:"duration"`
	PublicAt        bool          `json:"public_at"`
	CreatedAt       int           `json:"created_at"`
	UpdatedAt       int           `json:"updated_at"`
	IsPublic        bool          `json:"is_public"`
	IsCollaborative bool          `json:"is_collaborative"`
	Owner           PlaylistOwner `json:"owner"`
}

type PlaylistOwner struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}
