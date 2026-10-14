package qobuzsdk

type Track struct {
	Id                  int            `json:"id"`
	Title               string         `json:"title"`
	Version             string         `json:"version"`
	ISRC                string         `json:"isrc"`
	MaximumBitDepth     int            `json:"maximum_bit_depth"`
	MaximumChannelCount int            `json:"maximum_channel_count"`
	MaximumSamplingRate float64        `json:"maximum_sampling_rate"`
	Performer           TrackPerformer `json:"performer"`
	Performers          string         `json:"performers"`
	Copyright           string         `json:"copyright"`
	Album               Album          `json:"album"`
	TrackNumber         int            `json:"track_number"`
	MediaNumber         int            `json:"media_number"`
	ReleaseDateOriginal string         `json:"release_date_original"`
	ReleaseDateDownload string         `json:"release_date_download"`
	ReleaseDateStream   string         `json:"release_date_stream"`
	ReleaseDatePurchase string         `json:"release_date_purchase"`
	IsPurchasable       bool           `json:"purchasable"`
	IsStreamable        bool           `json:"streamable"`
	IsPreviewable       bool           `json:"previewable"`
	IsSampleable        bool           `json:"sampleable"`
	IsDownloadable      bool           `json:"downloadable"`
	IsDisplayable       bool           `json:"displayable"`
	IsHiRes             bool           `json:"hires"`
	IsHiResStreamable   bool           `json:"hires_streamable"`
	HasLyrics           bool           `json:"has_lyrics"`
	HasParentalWarning  bool           `json:"parental_warning"`
	Duration            int            `json:"duration"`
	PurchasableAt       int            `json:"purchasable_at"`
	StreamableAt        int            `json:"streamable_at"`
	CreatedAt           int            `json:"created_at"`
	IndexedAt           int            `json:"indexed_at"`
}

type TrackPerformer struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type TrackFileUrl struct {
	TrackId      int         `json:"track_id"`
	Duration     int         `json:"duration"`
	URL          string      `json:"url"`
	Format       TrackFormat `json:"format_id"`
	MimeType     string      `json:"mime_type"`
	SamplingRate float64     `json:"sampling_rate"`
	BitDepth     int         `json:"bit_depth"`
	Blob         string      `json:"blob"`
}

type TrackFormat int8

const (
	FormatMP3320          TrackFormat = 5
	FormatFLACCD          TrackFormat = 6
	FormatFLACHiRes96kHz  TrackFormat = 7
	FormatFLACHiRes192kHz TrackFormat = 27
)
