/**
 * This file is part of contacthub-sdk-go.
 *
 * contacthub-sdk-go is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 2 of the License, or
 * (at your option) any later version.
 *
 * contacthub-sdk-go is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with contacthub-sdk-go. If not, see <http://www.gnu.org/licenses/>.
 *
 * Copyright (C) 2017 Arduino AG
 *
 * @author Luca Osti
 *
 */

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
