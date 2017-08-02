package client

import (
	"testing"

	"github.com/kylelemons/godebug/pretty"
)

func TestAdd(t *testing.T) {

	tags := Tags{}
	tags.AddTag("tag1", true)
	tags.AddTag("tag2", false)
	expected := Tags{
		Auto:   []string{"tag1"},
		Manual: []string{"tag2"},
	}

	if diff := pretty.Compare(tags, expected); diff != "" {
		t.Errorf("Tags.Add: invalid value for struct: (-got +expected)\n%s", diff)
	}
}

func TestAddExisting(t *testing.T) {

	tags := Tags{}
	tags.AddTag("tag1", true)
	tags.AddTag("tag1", true)
	tags.AddTag("tag2", false)
	tags.AddTag("tag2", false)
	expected := Tags{
		Auto:   []string{"tag1"},
		Manual: []string{"tag2"},
	}

	if diff := pretty.Compare(tags, expected); diff != "" {
		t.Errorf("Tags.Add: invalid value for struct: (-got +expected)\n%s", diff)
	}
}

func TestRemove(t *testing.T) {

	tags := Tags{}
	tags.AddTag("tag1", true)
	tags.AddTag("tag2", true)
	tags.RemoveTag("tag1", true)

	tags.AddTag("tag3", false)
	tags.AddTag("tag4", false)
	tags.RemoveTag("tag3", false)

	expected := Tags{
		Auto:   []string{"tag2"},
		Manual: []string{"tag4"},
	}

	if diff := pretty.Compare(tags, expected); diff != "" {
		t.Errorf("Tags.Remove: invalid value for struct: (-got +expected)\n%s", diff)
	}
}
func TestRemoveNotExisting(t *testing.T) {

	tags := Tags{}
	tags.AddTag("tag2", true)
	tags.RemoveTag("tag1", true)

	tags.AddTag("tag4", false)
	tags.RemoveTag("tag3", false)

	expected := Tags{
		Auto:   []string{"tag2"},
		Manual: []string{"tag4"},
	}

	if diff := pretty.Compare(tags, expected); diff != "" {
		t.Errorf("Tags.Remove: invalid value for struct: (-got +expected)\n%s", diff)
	}
}
