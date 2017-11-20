package helper

import "golang.org/x/crypto/bcrypt"

// Generate the hash password, bcrypt.MinCost = 4, bcrypt.MaxCost = 31, bcrypt.DefaultCost = 10
func GeneratePassword(password string, cost int) (hash string, err error) {

	// Hashing the password with MinCost
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), cost)

	if err != nil {
		return hash, err
	}

	hash = string(hashedPassword)

	return hash, err
}

// Comparing the password with the hash
func ComparePassword(hash, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}
