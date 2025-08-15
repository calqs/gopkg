package crypt

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestICanSignAJWT(t *testing.T) {
	trials := []JWTDefaultClaims{{
		Expire: time.Date(2024, 10, 04, 22, 22, 22, 22, time.UTC).Unix(),
		UID:    0,
		Realm:  "test1",
	}, {
		Expire:  time.Date(2024, 10, 04, 22, 22, 22, 22, time.UTC).Unix(),
		Refresh: time.Date(2024, 10, 04, 23, 22, 22, 22, time.UTC).Unix(),
		UID:     0,
		Realm:   "test1",
	}}
	goals := []string{
		"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHBpcmUiOjE3MjgwODA1NDIsInVpZCI6MCwicmVhbG0iOiJ0ZXN0MSJ9.qEiSYDfztM43tVdinryLBb6EYoeJk70ysROuZXvtjpw",
		"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHBpcmUiOjE3MjgwODA1NDIsInJlZnJlc2giOjE3MjgwODQxNDIsInVpZCI6MCwicmVhbG0iOiJ0ZXN0MSJ9.yChC77_vDG2DTNl6Wh9mGznLR3GKNEDoptvaEGqlsBk",
	}

	for i, goal := range goals {
		signature, err := NewJWT(HS256("test"), trials[i])
		assert.NoError(t, err)
		assert.Equal(t, goal, signature)
	}
}

func TestICanAssertATokenWasNotTemperedWith(t *testing.T) {
	trial := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHBpcmUiOjE3MjgwODA1NDIsIm5hbWUiOiJ0ZXN0QHRlc3QuY29tIn0.0RbVgcJ7ZuMjfXwvbZjkrKG-5HQ2-NgSGKHUWn3_oeM"
	claims, err := DecodeJWT[JWTDefaultClaims](trial, HS256("test"))
	assert.NoError(t, err)
	assert.Equal(t, `{"expire":1728080542,"uid":0,"realm":""}`, claims.String())
}

func TestIFailOnTemperedToken(t *testing.T) {
	trial := "ohno.eyJleHAiOjE3MjgwODA1NDIwMDAsIm5hbWUiOiJ0ZXN0QHRlc3QuY29tIn0.zcxR5WqM-pxTVWc36Jsl0hwVHyGhaFiHy54BVLkVX9U"
	_, err := DecodeJWT[JWTDefaultClaims](trial, HS256("test"))
	assert.Error(t, err)
}

func TestICanDecodeAJWT(t *testing.T) {
	goal := JWTDefaultClaims{
		Expire: time.Date(2024, 10, 04, 22, 22, 22, 22, time.UTC).Unix(),
	}
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHBpcmUiOjE3MjgwODA1NDIsIm5hbWUiOiJ0ZXN0QHRlc3QuY29tIn0.0RbVgcJ7ZuMjfXwvbZjkrKG-5HQ2-NgSGKHUWn3_oeM"
	trial, err := DecodeJWT[JWTDefaultClaims](token, HS256("test"))
	assert.NoError(t, err)
	assert.Equal(t, goal, trial)
}

func Test_I_Can_Give_Correct_RemainingRefresh(t *testing.T) {
	goal := JWTDefaultClaims{
		Refresh: time.Date(2024, 10, 04, 22, 22, 22, 22, time.UTC).Unix(),
	}
	trial1 := goal.RemainingRefresh(time.Date(2024, 10, 04, 0, 0, 0, 0, time.UTC))
	assert.Equal(t, time.Hour*22+time.Minute*22+time.Second*22, trial1)
	trial2 := goal.RemainingRefresh(time.Date(2024, 10, 05, 0, 0, 0, 0, time.UTC))

	// 1 hour, 37mins and 38seconds to go from 2024-10-04 22:22:22 to 2024-10-05 00:00:00
	// 2024-10-05 00:00:00 - 2024-10-04 22:22:22 = -(1 hour, 37mins and 38seconds)
	assert.Equal(t, -(time.Hour + time.Minute*37 + time.Second*38), trial2)
}

type fakeClaims struct{}

func (fakeClaims) GetRawClaims() []byte { return []byte{} }

func Test_I_Can_Fail_On_Weird_Token(t *testing.T) {
	t.Run("jwt is not made of 3 parts", func(t *testing.T) {
		trial, err := DecodeJWT[fakeClaims]("", HS256(""))
		assert.Error(t, err)
		assert.Equal(t, trial, fakeClaims{})
	})
}
