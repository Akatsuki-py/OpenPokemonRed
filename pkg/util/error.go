package util

import (
	"fmt"
	"os"
)

// NotRegisteredError display error if ID is not registered in data map.
func NotRegisteredError(place string, id interface{}) {
	switch id := id.(type) {
	case string:
		fmt.Fprintf(os.Stderr, "DataError: Font[%s] is not registered in %s\n", id, place)
	case uint, int:
		fmt.Fprintf(os.Stderr, "DataError: ID[%d] is not registered in %s\n", id, place)
	}
}

// NotFoundFileError display error if file specified with path is not found
func NotFoundFileError(path string) {
	fmt.Fprintf(os.Stderr, "PathError: file is not found. %s\n", path)
}
