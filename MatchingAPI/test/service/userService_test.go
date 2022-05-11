package service

import (
	"MatchingAPI/security"
	"MatchingAPI/service"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUserService(t *testing.T) {
	t.Run("Get JWT", func(t *testing.T) {
		userService := service.NewUserService()

		token := userService.CreateJWT()
		jwt := token.Token
		assert.True(t, jwt != "")
		_, err := security.TokenParser(jwt)

		assert.Equal(t, nil, err)
	})

	t.Run("Get unauthorized JWT", func(t *testing.T) {
		userService := service.NewUserService()

		token := userService.CreateWrongJWT()
		jwt := token.Token
		assert.True(t, jwt != "")
		_, err := security.TokenParser(jwt)

		assert.Equal(t, "the user has not suitable credential", err.Error())
	})
}
