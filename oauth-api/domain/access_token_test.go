package domain

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGetAccessToken(t *testing.T) {
	at := GetAccessToken()

	assert.False(t, at.IsExpired(), "access token is expired")
	assert.EqualValues(t, "", at.AccessToken, "access token should not define a token id")
	assert.True(t, at.UserId == 0, "access token should not have an associated user id")
}

func TestGetAccessTokenIsExpired(t *testing.T) {
	at := AccessToken{}
	assert.True(t, at.IsExpired(), "empty access token should be expired be default")

	at.Expires = time.Now().Add(3 * time.Hour).Unix()
	assert.False(t, at.IsExpired(), "access token created three hours from now should not be expired")
}
