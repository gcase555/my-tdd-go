package maps

import "errors"

type Content map[string]string

var ErrNotFound = errors.New("word not found")

func (c Content) Search(word string) (string, error) {
	found := c[word]

	if len(found) > 1 {
		return found, nil
	}
	return "", ErrNotFound
}
