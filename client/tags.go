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
