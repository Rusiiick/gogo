package auth

type TokenGenerator interface {
	GenerateJWT(userID int, role int8) (string, error)
	GenerateRefreshToken() (string, error)
}
