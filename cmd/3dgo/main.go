package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"math/rand"

	"github.com/GuillaumeTech/3dgo/internal/hit"

	"github.com/GuillaumeTech/3dgo/internal/geom"
)

func rayHit(ray geom.Ray, world hit.HittableList) geom.Vec3d {
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
	const samplesPerPixels int = 50

	image := []byte(fmt.Sprintf("P3\n%.0f %.0f\n255\n", imageWidth, imageHeight))

	camera := geom.Camera{geom.Vec3d{-2, -1, -1}, geom.Vec3d{4, 0, 0},
		geom.Vec3d{0, 2, 0}, geom.Vec3d{0, 0, 0}}

	var world hit.HittableList
	world.Add(hit.Sphere{geom.Vec3d{0, 0, -1}, 0.5})
	world.Add(hit.Sphere{geom.Vec3d{0, -100.5, -1}, 100})

	for j := int(imageHeight) - 1; j >= 0; j-- {
		fmt.Println(fmt.Sprintf("Scan lines remaining: %d ", j))
		for i := 0; i < int(imageWidth); i++ {
			hitSum := geom.Vec3d{0, 0, 0}
			for s := 0; s < samplesPerPixels; s++ {
				u := (float64(i) + rand.Float64()) / imageWidth
				v := (float64(j) + rand.Float64()) / imageHeight
				ray := camera.GetRay(u, v)
				rayHit := rayHit(ray, world)
				hitSum = geom.AddTwoVec(rayHit, hitSum)
			}
			image = append(image, []byte(hitSum.GetColor(samplesPerPixels))...)

		}
	}

	// write the whole body at once
	err := ioutil.WriteFile("output.ppm", image, 0644)
	if err != nil {
		panic(err)
	}
	fmt.Println("Done!")
}
