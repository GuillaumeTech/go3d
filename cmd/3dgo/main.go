package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"

	"github.com/GuillaumeTech/3dgo/internal/vec3d"
)

func main() {
	nx, ny := 200, 100
	image := []byte(fmt.Sprintf("P3\n%d %d\n255\n", nx, ny))
	for j := ny - 1; j >= 0; j-- {
		fmt.Println(fmt.Sprintf("Scan lines remaining: %d ", j))
		c := exec.Command("clear")
		c.Stdout = os.Stdout
		c.Run()
		for i := 0; i < nx; i++ {
			vector := vec3d.Vec3d{float64(i) / float64(nx), float64(j) / float64(ny), 0.2}
			image = append(image, []byte(vector.GetColor())...)
		}
	}

	// write the whole body at once
	err := ioutil.WriteFile("output.ppm", image, 0644)
	if err != nil {
		panic(err)
	}
	fmt.Println("Done!")
}
