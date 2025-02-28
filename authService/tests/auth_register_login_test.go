package tests

import (
	"github.com/RamzassH/LeadIt/authService/tests/suite"
	authv1 "github.com/RamzassH/LeadIt/libs/contracts/gen/auth"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/dgrijalva/jwt-go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"log"
	"testing"
)

const (
	passDefaultLen = 10
	tokenSecret    = "super_secret_token"
)

type UserData struct {
	name    string
	surname string
}

func TestRegisterLogin_Login_HappyPath(t *testing.T) {
	ctx, st := suite.New(t)

	email := gofakeit.Email()
	password := randomFakePassword()
	UserData := randomFakeUserData()

	regResponse, err := st.AuthClient.Register(ctx, &authv1.RegisterRequest{
		Email:    email,
		Password: password,
		Name:     UserData.name,
		Surname:  UserData.surname,
	})
	log.Println(regResponse)

	require.NoError(t, err)
	assert.NotEmpty(t, regResponse.GetUserId())

	loginResponse, err := st.AuthClient.Login(ctx, &authv1.LoginRequest{
		Email:    email,
		Password: password,
	})

	require.NoError(t, err)

	token := loginResponse.GetToken()
	require.NotEmpty(t, token)

	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(tokenSecret), nil
	})
	require.NoError(t, err)

	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	require.True(t, ok)

	assert.Equal(t, regResponse.GetUserId(), int64(claims["uid"].(float64)))
	assert.Equal(t, email, claims["email"].(string))
}

func randomFakePassword() string {
	return gofakeit.Password(true, true, true, true, true, passDefaultLen)
}

func randomFakeUserData() *UserData {
	name := gofakeit.Name()
	surname := gofakeit.Name()

	return &UserData{name, surname}
}
