package qobuzsdk

type Album struct {
	Id                             string            `json:"id"`
	QobuzId                        int               `json:"qobuz_id"`
	UPC                            string            `json:"upc"`
	Slug                           string            `json:"slug"`
	URL                            string            `json:"url"`
	Title                          string            `json:"title"`
	Subtitle                       string            `json:"subtitle"`
	Version                        string            `json:"version"`
	Description                    string            `json:"description"`
	DescriptionLanguage            string            `json:"description_language"`
	Catchline                      string            `json:"catchline"`
	Copyright                      string            `json:"copyright"`
	RecordingInformations          string            `json:"recording_information"`
	ReleasedAt                     int64             `json:"released_at"`
	ReleaseDateOriginal            string            `json:"release_date_original"`
	ReleaseDateDownload            string            `json:"release_date_download"`
	ReleaseDateStream              string            `json:"release_date_stream"`
	PurchasableAt                  int               `json:"purchasable_at"`
	StreamableAt                   int               `json:"streamable_at"`
	IsPurchasable                  bool              `json:"purchasable"`
	IsStreamable                   bool              `json:"streamable"`
	IsPreviewable                  bool              `json:"previewable"`
	IsSampleable                   bool              `json:"sampleable"`
	IsDownloadable                 bool              `json:"downloadable"`
	IsDisplayable                  bool              `json:"displayable"`
	IsAiGenerated                  bool              `json:"ai_generated"`
	IsOfficial                     bool              `json:"is_official"`
	HasParentalWarning             bool              `json:"parental_warning"`
	IsHiRes                        bool              `json:"hires"`
	IsHiResStreamable              bool              `json:"hires_streamable"`
	MaximumSamplingRate            float64           `json:"maximum_sampling_rate"`
	MaximumBitDepth                int               `json:"maximum_bit_depth"`
	MaximumChannelCount            int               `json:"maximum_channel_count"`
	MaximumTechnicalSpecifications string            `json:"maximum_technical_specifications"`
	TracksCount                    int               `json:"tracks_count"`
	MediaCount                     int               `json:"media_count"`
	Duration                       int               `json:"duration"`
	Popularity                     float64           `json:"popularity"`
	Artist                         AlbumMainArtist   `json:"artist"`
	Artists                        []AlbumArtist     `json:"artists"`
	Label                          AlbumLabel        `json:"label"`
	Genre                          Genre             `json:"genre"`
	Awards                         []AlbumAward      `json:"awards"`
	Image                          CoverImage        `json:"image"`
	Tracks                         Pagination[Track] `json:"tracks"`
}

type AlbumMainArtist struct {
	Id          int        `json:"id"`
	Slug        string     `json:"slug"`
	Name        string     `json:"name"`
	Image       string     `json:"image"`
	AlbumsCount int        `json:"albumsCount"`
	Picture     CoverImage `json:"picture"`
}

type AlbumArtist struct {
	Id    int      `json:"id"`
	Name  string   `json:"name"`
	Roles []string `json:"roles"`
}

type AlbumLabel struct {
	Id          int    `json:"id"`
	Slug        string `json:"slug"`
	Name        string `json:"name"`
	AlbumsCount int    `json:"albums_count"`
}

type AlbumAward struct {
	Name            string `json:"name"`
	Slug            string `json:"slug"`
	AwardSlug       string `json:"award_slug"`
	AwardedAt       int    `json:"awarded_at"`
	AwardId         string `json:"award_id"`
	PublicationId   string `json:"publication_id"`
	PublicationName string `json:"publication_name"`
	PublicationSlug string `json:"publication_slug"`
}
