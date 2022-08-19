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
	websites := []string{
		"http://cats.de",
		"http://something.kz",
		"http://wrongsite.com",
	}

	want := map[string]bool{
		"http://cats.de":       true,
		"http://something.kz":  true,
		"http://wrongsite.com": false,
	}

	got := CheckWebsites(mockWebsiteChecker, websites)

	if !reflect.DeepEqual(want, got) {
		t.Fatalf("wanted %v, got %v", want, got)
	}
}