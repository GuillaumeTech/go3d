package importFiles

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/GuillaumeTech/3dgo/internal/geom"
	"github.com/GuillaumeTech/3dgo/internal/hit"
)

func vertexStringToNumber(vertexInfo string) [3]float64 {
	var vertex [3]float64
	vertexCoord := strings.Split(vertexInfo, " ")[1:4]
	for i, coord := range vertexCoord {
		vertex[i], _ = strconv.ParseFloat(coord, 64)
	}
	return vertex
}

func faceStringToNumber(faceInfo string) [3]int {
	var face [3]int
	faceVertex := strings.Split(faceInfo, " ")[1:4]
	for i, faceNumber := range faceVertex {
		face[i], _ = strconv.Atoi(faceNumber)
		face[i] = face[i] - 1
	}
	return face
}

func AddObjToWorld(path string, mat hit.Material, world hit.HittableList) hit.HittableList {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var faces [][3]int
	var vertexes [][3]float64
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		if (strings.HasPrefix(line,"v")) {
			vertexes = append(vertexes, vertexStringToNumber(line))
		}
		if strings.HasPrefix(line,"f") {
			faces = append(faces, faceStringToNumber(line))
		}
	}
  println(len(vertexes))	
	for _, face := range faces {
	
		a := geom.Vec3d{X: vertexes[face[0]][0], Y: vertexes[face[0]][1], Z: vertexes[face[0]][2]}
		b := geom.Vec3d{X: vertexes[face[1]][0], Y: vertexes[face[1]][1], Z: vertexes[face[1]][2]}
		c := geom.Vec3d{X: vertexes[face[2]][0], Y: vertexes[face[2]][1], Z: vertexes[face[2]][2]}
		triFace := hit.NewTriangle(a, b, c, mat)
		world.Add(triFace)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return world
}
