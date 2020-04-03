package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
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
			r := float64(i) / float64(nx)
			g := float64(j) / float64(ny)
			b := 0.2
			ir := int(255.99 * r)
			ig := int(255.99 * g)
			ib := int(255.99 * b)
			image = append(image, []byte(fmt.Sprintf("%d %d %d\n", ir, ig, ib))...)
		}
	}

	// write the whole body at once
	err := ioutil.WriteFile("output.ppm", image, 0644)
	if err != nil {
		panic(err)
	}
	fmt.Println("Done!")
}
