package setup

import (
	"strconv"
	"time"

	"github.com/JudicaelT/qobuzsdk"
	"github.com/JudicaelT/qobuzsdk/internal/qobuz"
	"github.com/JudicaelT/qobuzsdk/internal/rest"
)

func NewClient(
	appId int32,
	appSecret string,
	privateKey string,
) qobuzsdk.Client {
	apiClient := rest.NewApiClient(qobuz.ApiBaseUrl)
	apiClient.SetHeader("X-App-Id", strconv.Itoa(int(appId)))
	return qobuz.NewClient(
		apiClient,
		appId,
		appSecret,
		privateKey,
		time.Now().Unix,
	)
}
