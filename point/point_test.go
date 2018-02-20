package point

import "fmt"

func ExampleLatLong_ToZXY() {
	p := ZXY{Z: 10, X: 697, Y: 321}
	fmt.Println(p.ToLatLong())
	// Output:
	// {Lat:55.578344672182055 Long:65.0390625}
}

func ExampleZXY_ToLatLong() {
	p := LatLong{Lat: 55.578344672182055, Long: 65.0390625}
	fmt.Println(p.ToZXY(10))
	// Output:
	// {Z:10 X:697 Y:321}
}
