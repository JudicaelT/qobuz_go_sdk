package qobuzsdk

type ApiError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (self *ApiError) Error() string {
	return self.Message
}

type TrackNotFoundError struct {
	TrackId int
	Err     error
}

func (self *TrackNotFoundError) Error() string {
	return self.Err.Error()
}

func (self *TrackNotFoundError) Unwrap() error {
	return self.Err
}

type AlbumNotFoundError struct {
	AlbumId string
	Err     error
}

func (self *AlbumNotFoundError) Error() string {
	return self.Err.Error()
}

func (self *AlbumNotFoundError) Unwrap() error {
	return self.Err
}

type PlaylistNotFoundError struct {
	PlaylistId int
	Err        error
}

func (self *PlaylistNotFoundError) Error() string {
	return self.Err.Error()
}

func (self *PlaylistNotFoundError) Unwrap() error {
	return self.Err
}

type ArtistNotFoundError struct {
	ArtistId int
	Err      error
}

func (self *ArtistNotFoundError) Error() string {
	return self.Err.Error()
}

func (self *ArtistNotFoundError) Unwrap() error {
	return self.Err
}
