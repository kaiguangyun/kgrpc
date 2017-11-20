package auth

import "errors"

// secret
// todo : use redis
var Secret = map[string]string{}

// Error constants
var ErrTokenInvalid = errors.New("token invalid")

// jwt secret
func GetSecret(uuid string) (secret string) {
	return Secret[uuid]
}

// save secret
func SaveSecret(uuid, secret string) {
	Secret[uuid] = secret
}
