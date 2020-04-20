package mydict

import (
	"errors"
)

type Dictionary map[string]string

var (
	errSearch    = errors.New("there is no key in our dictionary")
	errExists    = errors.New("That key is already exists")
	errNotUpdate = errors.New("That key can't be updated")
)

func (d Dictionary) Search(word string) (string, error) {
	value, err := d[word]
	if err {
		return value, nil
	}
	return "", errSearch
}

func (d Dictionary) Add(key, value string) error {
	_, err := d.Search(key)
	switch err {
	case errSearch:
		d[key] = value
	case nil:
		return errExists
	}
	return nil
}

func (d Dictionary) Update(key, value string) error {
	_, err := d.Search(key)
	switch err {
	case nil:
		d[key] = value
	case errSearch:
		return errNotUpdate
	}
	return nil
}
