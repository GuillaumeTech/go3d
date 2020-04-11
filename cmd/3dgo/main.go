package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"os/exec"

	"github.com/GuillaumeTech/3dgo/internal/hit"

	"github.com/GuillaumeTech/3dgo/internal/geom"
)

func rayColor(ray geom.Ray, world hit.HittableList) geom.Vec3d {
	var record hit.HitRecord
	if world.Hit(ray, 0, math.Inf(1), &record) {
		return geom.MultiplyVec(0.5, geom.Vec3d{record.Normal.X + 1, record.Normal.Y + 1, record.Normal.Z + 1})
	}
	unitDir := geom.UnitVector(ray.Direction)
	t := 0.5 * (unitDir.Y + 1)
	start := geom.Vec3d{1, 1, 1}
	end := geom.Vec3d{0.5, 0.7, 1.0}
	return geom.AddTwoVec(start.Multiply(1-t), end.Multiply(t)) //lerp

}

func main() {
	const imageWidth float64 = 800
	const imageHeight float64 = 400

	image := []byte(fmt.Sprintf("P3\n%.0f %.0f\n255\n", imageWidth, imageHeight))

	lowerLeftCorner := geom.Vec3d{-2, -1, -1}
	horizontal := geom.Vec3d{4, 0, 0}
	vertical := geom.Vec3d{0, 2, 0}
	origin := geom.Vec3d{0, 0, 0}
	var world hit.HittableList
	world.Add(hit.Sphere{geom.Vec3d{0, 0, -1}, 0.5})
	world.Add(hit.Sphere{geom.Vec3d{0, -100.5, -1}, 100})

	for j := int(imageHeight) - 1; j >= 0; j-- {
		fmt.Println(fmt.Sprintf("Scan lines remaining: %d ", j))
		c := exec.Command("clear")
		c.Stdout = os.Stdout
		c.Run()
		for i := 0; i < int(imageWidth); i++ {
			u := float64(i) / imageWidth
			v := float64(j) / imageHeight
			direction := lowerLeftCorner.Add(horizontal.Multiply(u))
			direction = direction.Add(vertical.Multiply(v))
			ray := geom.Ray{origin, direction}

			rayColor := rayColor(ray, world)
			image = append(image, []byte(rayColor.GetColor())...)
		}
	}

	// write the whole body at once
	err := ioutil.WriteFile("output.ppm", image, 0644)
	if err != nil {
		panic(err)
	}
	fmt.Println("Done!")
}
