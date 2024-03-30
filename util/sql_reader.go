package util

import (
	"fmt"
	"os"
)

func ReadSQLFile(filename string) (string, error) {
	bytes, err := os.ReadFile(
		fmt.Sprintf("sql/endpoints/%s", filename),
	)

	if err != nil {
		return "", err
	}

	sqlAsString := string(bytes)
	return sqlAsString, nil
}
