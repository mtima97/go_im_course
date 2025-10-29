package files

import "os"

func Read() (string, error) {
	bs, err := os.ReadFile("data/test.txt")
	if err != nil {
		return "", err
	}

	return string(bs), nil
}
