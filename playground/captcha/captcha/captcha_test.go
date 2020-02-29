package captcha_test

import (
	"github.com/fx2y/learn-web-service-with-golang/playground/captcha/captcha"
	"image"
	"testing"
)

func TestChallengeUserSuccess(t *testing.T) {
	got := captcha.ChallengeUser(stubChallenger("42"), stubPrompter("42"))
	if got != true {
		t.Fatal("expected ChallengeUser to return true")
	}
}

func TestChallengeUserFail(t *testing.T) {
	got := captcha.ChallengeUser(stubChallenger("lorem ipsum"), stubPrompter("42"))
	if got != false {
		t.Fatal("expected ChallengeUser to return false")
	}
}

type stubChallenger string

func (c stubChallenger) Challenge() (image.Image, string) {
	return image.NewRGBA(image.Rect(0, 0, 100, 100)), string(c)
}

type stubPrompter string

func (p stubPrompter) Prompt(_ image.Image) string {
	return string(p)
}
