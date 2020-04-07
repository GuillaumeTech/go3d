package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"os/exec"

	"github.com/GuillaumeTech/3dgo/internal/geom"
)

func hitSphere(center geom.Vec3d, radius float64, ray geom.Ray) float64 {
	// the ray goes througth the sphere is a 2nd deg equation
	oc := geom.SubstractTwoVec(ray.Origin, center)
	a := geom.DotProduct(ray.Direction, ray.Direction)
	b := 2 * geom.DotProduct(oc, ray.Direction)
	c := geom.DotProduct(oc, oc) - radius*radius
	discriminant := b*b - 4*a*c
	if discriminant < 0 {
		return -1
	} else {
		return ((-b - math.Sqrt(discriminant)) / (2 * a))
	}
}

func rayColor(ray geom.Ray) geom.Vec3d {
	sphereCenter := geom.Vec3d{0, 0, -1}
	root := hitSphere(sphereCenter, 0.7, ray)
	if root > 0 {
		normal := geom.UnitVector(geom.SubstractTwoVec(ray.At(root), sphereCenter))
		return geom.MultiplyVec(0.5, geom.Vec3d{normal.X + 1, normal.Y + 1, normal.Z + 1})
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

	image := []byte(fmt.Sprintf("P3\n%.0f %.0f\n255\n", imageWidth, imageHeight))

	lowerLeftCorner := geom.Vec3d{-2, -1, -1}
	horizontal := geom.Vec3d{4, 0, 0}
	vertical := geom.Vec3d{0, 2, 0}
	origin := geom.Vec3d{0, 0, 0}

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

			rayColor := rayColor(ray)
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
