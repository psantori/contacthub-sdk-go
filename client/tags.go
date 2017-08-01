package client

type Tags struct {
	Auto   []string `json:"auto,omitempty"`
	Manual []string `json:"manual,omitempty"`
}

//AddTag adds a tag
func (t *Tags) AddTag(tag string, auto bool) {
	list := t.Manual
	if auto {
		list = t.Auto
	}

	if searchTag(list, tag) > -1 {
		list = append(list, tag)
	}
}

//RemoveTag removes a tag
func (t *Tags) RemoveTag(tag string, auto bool) {
	list := t.Manual
	if auto {
		list = t.Auto
	}

	if i := searchTag(list, tag); i > -1 {
		list[i] = list[len(list)-1]
		list = list[:len(list)-1]
	}
}

func searchTag(list []string, tag string) int {
	for i := range list {
		if list[i] == tag {
			return i
		}
	}
	return -1
}
