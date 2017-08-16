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

type Tags struct {
	Auto   []string `json:"auto,omitempty"`
	Manual []string `json:"manual,omitempty"`
}

//AddTag adds a tag
func (t *Tags) AddTag(tag string, auto bool) {
	if auto {
		t.Auto = addElement(tag, t.Auto)
	} else {
		t.Manual = addElement(tag, t.Manual)
	}
}

//RemoveTag removes a tag
func (t *Tags) RemoveTag(tag string, auto bool) {
	if auto {
		t.Auto = removeElement(tag, t.Auto)
	} else {
		t.Manual = removeElement(tag, t.Manual)
	}
}

func addElement(tag string, list []string) []string {
	if searchElement(list, tag) == -1 {
		return append(list, tag)
	}
	return list
}

func removeElement(tag string, list []string) []string {
	if i := searchElement(list, tag); i > -1 {
		list[i] = list[len(list)-1]
		return list[:len(list)-1]
	}
	return list
}

func searchElement(list []string, tag string) int {
	for i := range list {
		if list[i] == tag {
			return i
		}
	}
	return -1
}
