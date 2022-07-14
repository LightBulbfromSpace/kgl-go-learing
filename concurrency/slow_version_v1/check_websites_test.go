package concurrency

import (
	"reflect"
	"testing"
)

func mockWebsiteChecker(url string) bool {
	if url == "http://wrongsite.com" {
		return false
	}
	return true
}

func TestCheckWebsites(t *testing.T) {

	urls := []string{
		"http://cats.de",
		"http://something.kz",
		"http://wrongsite.com",
	}

	want := map[string]bool{
		"http://cats.de":       true,
		"http://something.kz":  true,
		"http://wrongsite.com": false,
	}
	got := CheckWebsites(mockWebsiteChecker, urls)

	if !reflect.DeepEqual(want, got) {
		t.Errorf("got %v but want %v", got, want)
	}
}
