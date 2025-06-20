package utils

import (
	"os/user"
)

func Whoami() (string, error) {
	currentUser, err := user.Current()

	if err != nil {
		return "", err
	}

	return currentUser.Username, nil
}
