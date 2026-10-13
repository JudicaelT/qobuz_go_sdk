package qobuzsdk

const OAuthCodeGetParam = "code_autorisation"

type OAuthLogin struct {
	LoginLink      string
	OnCodeReceived func(code string) error
}
