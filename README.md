# snowtoken - How to use
Example code:
Begin:
``` 
go get github.com/namphamtoday/snowtoken 
```
```
import (
	"time"
	"github.com/namphamtoday/snowtoken/snow_jwt"
)
```
Create token:

```
func CreateToken() snow_jwt.TokenResponse {
	const tokenExpiresIn int = 300      // seconds - 5 minutes
	const refreshExpiresIn int = 604800 // seconds - 7 days

	now := time.Now().UTC()

	var contentToken snow_jwt.TokenContent
	contentToken.UserId = 1
	contentToken.Subject = "nampham"
	contentToken.Audience = "test"
	contentToken.Issuer = "Test Company"
	contentToken.ExpiresAt = now.Add(time.Duration(tokenExpiresIn) * time.Second).Unix()
	contentToken.IssuedAt = now.Unix()
	contentToken.NotBefore = now.Unix()
	contentToken.Data = "your custom data"

	var resToken snow_jwt.TokenResponse

	prvKey := []byte("YOUR PRIVATE KEY HERE ...")

	jwtToken := snow_jwt.NewJWT(prvKey, nil)

	// 1. Create a new JWT token.
	var err error

	resToken.ExpiresIn = tokenExpiresIn
	resToken.TokenType = "Bearer"
	resToken.AccessToken, err = jwtToken.Create(time.Duration(tokenExpiresIn)*time.Second, contentToken)
	if err != nil {
		log.Fatalln(err)
	}
	resToken.RefreshToken, _ = jwtToken.Create(time.Duration(refreshExpiresIn)*time.Second, contentToken)

	return resToken
}
```
Verify token:
```
func VerifyToken(tok string) bool{
	dat, err := IsValidToken(tok)
	if err != nil {
		return false
	} else {
		return true
	}
}

func IsValidToken(tok string) (interface{}, error) {
	pubKey := []byte("YOUR PUBLIC KEY HERE...")
	jwtToken := snow_jwt.NewJWT(nil, pubKey)
	return jwtToken.Validate(tok)
}

```
