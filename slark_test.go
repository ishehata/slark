package slark

import (
	"testing"
)

func TestReadingStructTags(t *testing.T) {
	type post struct {
		Name    string `db:"name" type:"varchar"`
		Content string `db:"content" type:"text"`
	}

	Register(&post{}, "posts")

	for key, val := range models {
		t.Log(key, val.Name)
		for _, f := range val.Fields {
			t.Log(f.Name, f.DBType)
		}
	}
}
