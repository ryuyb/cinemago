package themoviedb

import (
	"fmt"
	"testing"
)

func TestName(t *testing.T) {
	client, err := NewClient()
	if err != nil {
		t.Fatal(err)
	}
	trending, err := client.GetAllTrending(1, "all", "week")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%#v\n", trending)
}

func TestGetMovieDetail(t *testing.T) {
	client, err := NewClient()
	if err != nil {
		t.Fatal(err)
	}
	trending, err := client.GetMovieDetail(497698)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%#v\n", trending)
}
