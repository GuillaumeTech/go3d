package importFiles

import (
	"bufio"
	"log"
	"os"
	"regexp"

	"github.com/GuillaumeTech/3dgo/internal/hit"
)

func AddObjToWorld(path string, mat hit.Material, world hit.HittableList) hit.HittableList {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		vertexRegexp, _ := regexp.Compile("v +[0-9]")
		normalRegexp, _ := regexp.Compile("vn +[0-9]")
		faceRegexp, _ := regexp.Compile("f +[0-9]")
		line := scanner.Text()
		if vertexRegexp.MatchString(line) {
		}
		if normalRegexp.MatchString(line) {
		}
		if faceRegexp.MatchString(line) {
		}

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
