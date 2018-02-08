package mandelbrot

//things


import (
	"fmt"
	"math"
	"math/cmplx"
)

//Img makes a slice of pixels (sort of ) representating a 2-D slice of the mandelbrot set
//given a canvas size
//planeCoordinates [top, right, left, bottom]
//TODO: convert this to use goroutines
//TODO: TEST THIS
func Img(canvasWidth int, canvasHeight int, planeCoordinates [4]float64, maxIterations int) [][]int{
  //dreading writing the conversion functions again ....
  //make our 2-D pixel slice
  pixels := make([][]int, canvasWidth)
  for i := range pixels {
    pixels[i] = make([]int, canvasHeight)
  }
  //iterate over each pixel checking the setMembership and iteration escape depth
    ////set up some facts about our complex plane
  top := planeCoordinates[0]
  right := planeCoordinates[1]
  bottom := planeCoordinates[2]
  left := planeCoordinates[3]
  planeWidth := right - left
  planeHeight := top - bottom
  for x := 0; x < len(pixels); x++ {
    for y := 0; y < len(pixels[x]); y++ {
      //convert pixels coords to complex plane coords
      c := complex((float64(x) * planeWidth / float64(canvasWidth) + left), (float64(y) * planeHeight / float64(canvasHeight) + bottom ))
      //check the setMembership
      // escape := escapeIteration(c, maxIterations)

      escapeIteration(c, maxIterations)
      // pixels[x][y] = escape
      // fmt.Println(escape)
      pixels[x][y] = 100
    }
  }
  return pixels
}

//escapeIteration returns the iteration the c escapes set memebership
//0 <= escape <= max, with 0 being the flag for "hasn't escaped yet"
//consider changing the flag thing
func escapeIteration(c complex128, max int) int {
  // fmt.Println(c)
  //set default max value
  if max == 0 {
    max = 80
  }
	distance := float64(2 * 2)
	z := 0 + 0i
	for i := 0; i < max; i++ {
		z = z*z + c
		r, theta := cmplx.Polar(z)
		real := math.Cos(theta) * r
		imaginary := math.Sin(theta) * r
		if real*real+imaginary*imaginary > distance {
			return i + 1
		}
	}
	 return 0//flag value
}

//IsMandelbrot tests whether or not c is in the set
func IsMandelbrot(c complex128) bool {
	fmt.Println(c)
	maxIterations := 80
	distance := float64(2 * 2)
	z := 0 + 0i
	for i := 0; i < maxIterations; i++ {
		z = z*z + c
		r, theta := cmplx.Polar(z)
		real := math.Cos(theta) * r
		imaginary := math.Sin(theta) * r
		if real*real+imaginary*imaginary > distance {
			return false
		}
	}
	return true
}

//
