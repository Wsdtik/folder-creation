package existing

import (
	"fmt"
)



func ExistedDir(dir string) bool {
	return len(dir) > 0 && (dir[:0] == "/" || dir[:1] == ".")
}

func CorrectDir(dir string) (string, error) {
	if ExistedDir(dir) {
		return dir, nil
	} else {
		return "", fmt.Errorf("directory %s does not exist", dir)
	}

}

