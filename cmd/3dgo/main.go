package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"math/rand"

	"github.com/GuillaumeTech/3dgo/internal/hit"

	"github.com/GuillaumeTech/3dgo/internal/geom"
)

func rayHit(ray geom.Ray, world hit.HittableList, depth int) geom.Vec3d {
	var record hit.HitRecord

	if depth <= 0 {
		return geom.Vec3d{0, 0, 0}
	}
	if world.Hit(ray, 0.001, math.Inf(1), &record) {
		var scattered geom.Ray
		var attenuation geom.Vec3d
		if record.Mat.Scatter(ray, &record, &attenuation, &scattered) {
			return geom.MultiplyVecPerCoords(attenuation, rayHit(scattered, world, depth-1))
		}
		return geom.Vec3d{0, 0, 0}
	}
	unitDir := geom.UnitVector(ray.Direction)
	t := 0.5 * (unitDir.Y + 1)
	start := geom.Vec3d{1, 1, 1}
	end := geom.Vec3d{0.5, 0.7, 1.0}
	return geom.AddTwoVec(start.Multiply(1-t), end.Multiply(t)) //lerp

}

func main() {
	const imageWidth float64 = 400
	const imageHeight float64 = 200
	const samplesPerPixels int = 100
	const maxDepth int = 50

	image := []byte(fmt.Sprintf("P3\n%.0f %.0f\n255\n", imageWidth, imageHeight))

	camera := geom.NewCamera(geom.Vec3d{-2, 2, 1}, geom.Vec3d{0, 0, 0}, geom.Vec3d{0, 1, 0}, 27, imageWidth/imageHeight)

	var world hit.HittableList

	triangle := hit.NewTriangle(geom.Vec3d{0, 0, 0.4}, geom.Vec3d{0, 0.2, 0}, geom.Vec3d{0, 0.7, 0}, hit.Metal{geom.Vec3d{0.5, 0.01, 0}, 0})
	triangle2 := hit.NewTriangle(geom.Vec3d{1, 0, 0.4}, geom.Vec3d{0, 0.2, 0}, geom.Vec3d{0, 0.7, 0}, hit.Metal{geom.Vec3d{0.5, 0.7, 0}, 0})

	world.Add(triangle)
	world.Add(triangle2)

	// world.Add(hit.Sphere{geom.Vec3d{0.5, 0, -1}, 0.47, hit.Lambertian{geom.Vec3d{0.7, 0.7, 0.7}}})
	// world.Add(hit.Sphere{geom.Vec3d{-0.7, 0, -1}, 0.47, hit.Dielectric{1.45}})

	world.Add(hit.Sphere{geom.Vec3d{0, -100.5, -1}, 100, hit.Lambertian{geom.Vec3d{0.3, 0.3, 0.6}}})

	for j := int(imageHeight) - 1; j >= 0; j-- {
		fmt.Println(fmt.Sprintf("Scan lines remaining: %d ", j))
		for i := 0; i < int(imageWidth); i++ {
			hitSum := geom.Vec3d{0, 0, 0}
			for s := 0; s < samplesPerPixels; s++ {
				u := (float64(i) + rand.Float64()) / imageWidth
				v := (float64(j) + rand.Float64()) / imageHeight
				ray := camera.GetRay(u, v)
				rayHit := rayHit(ray, world, maxDepth)
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
