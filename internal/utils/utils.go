package utils

import (
	"fmt"
	"os"
	"path"
	"strings"
	"time"
)

func DerriveLocalRepo(repo, destinationDir string) string {
	dirName := strings.TrimSuffix(path.Base(repo), path.Ext(path.Base(repo)))
	return destinationDir + string(os.PathSeparator) + dirName
}

func Spinner(delay time.Duration) {
	for {
		for {
			fmt.Print(".")
			time.Sleep(delay)
		}
	}
}
