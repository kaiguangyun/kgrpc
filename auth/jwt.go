package auth

import (
	"encoding/json"
	"github.com/dgrijalva/jwt-go"
	"github.com/kaiguangyun/kgrpc/helper"
	"strconv"
	"strings"
	"time"
)

// custom claims
type JwtClaims struct {
	Uuid string `json:"Uuid"`
	jwt.StandardClaims
}

// user uuid and secret
func JwtGenerateToken(uuid, secret string) (tokenStr string, err error) {
	// save secret
	SaveSecret(uuid, secret)
	// get token expiration
	jwtExpiration, _ := strconv.Atoi(helper.GetEnv("JwtExpiration"))
	jwtExpiresAt := time.Now().Unix() + int64(jwtExpiration)
	jwtIssuer := helper.GetEnv("JwtIssuer")
	// Create the Claims
	claims := JwtClaims{
		Uuid: uuid,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: jwtExpiresAt,
			Issuer:    jwtIssuer,
		},
	}
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return jwtToken.SignedString([]byte(secret)) // output : tokenStr = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.claims..."
}

//tokenString = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.claims..."
func JwtValidateToken(token, secret string) (valid bool) {
	jwtAt(time.Unix(0, 0), func() {
		jwtToken, _ := jwt.ParseWithClaims(token, &JwtClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(secret), nil
		})

		if claims, ok := jwtToken.Claims.(*JwtClaims); ok && jwtToken.Valid {
			valid = claims.StandardClaims.ExpiresAt > time.Now().Unix()
		}
	})
	return valid
}

// jwt get claims
func JwtGetClaims(token string) (claims *JwtClaims, err error) {
	claims = new(JwtClaims)
	tokenSlice := strings.Split(token, ".")
	if len(tokenSlice) != 3 {
		return claims, ErrTokenInvalid
	}
	claimsBytes, err := JWTDecodeSegment(tokenSlice[1])
	if err != nil {
		return claims, err
	}
	if err = json.Unmarshal(claimsBytes, claims); err != nil {
		return claims, err
	}
	return claims, err
}

// JWT Decode specific base64url encoding with padding stripped
func JWTDecodeSegment(seg string) ([]byte, error) {
	return jwt.DecodeSegment(seg)
}

// Override time value for tests.  Restore default value after.
func jwtAt(t time.Time, f func()) {
	jwt.TimeFunc = func() time.Time {
		return t
	}
	f()
	jwt.TimeFunc = time.Now
}
