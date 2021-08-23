package snowtoken

type TokenResponse struct {
	AccessToken  string `xmlrpc:"access_token" json:"access_token" form:"access_token"`
	ExpiresIn    int    `xmlrpc:"expires_in" json:"expires_in" form:"expires_in"`
	TokenType    string `xmlrpc:"token_type" json:"token_type" form:"token_type"`
	RefreshToken string `xmlrpc:"refresh_token" json:"refresh_token" form:"refresh_token"`
}

type RefreshTokenRequest struct {
	RefreshToken string `xmlrpc:"refresh_token" json:"refresh_token" form:"refresh_token"`
}

type TokenContent struct {
	UserId    interface{} `json:"user_id" form:"user_id"`
	SessionId string      `json:"session_id" form:"session_id"`
	Subject   string      `json:"subject" form:"subject"`
	Issuer    string      `json:"issuer" form:"issuer"`
	ExpiresAt int64       `json:"expires_at" form:"expires_at"`
	NotBefore int64       `json:"not_before" form:"not_before"`
	IssuedAt  int64       `json:"issued_at" form:"issued_at"`
	Audience  string      `json:"audience" form:"audience"`
	Data      interface{} `json:"data" form:"data"`
}
