package maps

import "errors"

var ErrNotFound = errors.New("couldn't find the word that you were looking for")

type Dictionary map[string]string

func Search(dict map[string]string, word string) string {
	return dict[word]
}

func (d Dictionary) Search(word string) (string, error) {
	def, found := d[word]
	if !found {
		return "", ErrNotFound
	}
	return def, nil
}

func (d Dictionary) Add(key, value string) {
	d[key] = value
}
