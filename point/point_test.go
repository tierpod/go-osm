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

func ExampleZXYBox() {
	p1 := LatLong{Lat: 25.2919, Long: 61.4949}
	p2 := LatLong{Lat: 39.6482, Long: 44.7653}
	zooms := []int{5, 6}
	pCh := ZXYBox(zooms, p1, p2)
	for p := range pCh {
		fmt.Println(p)
	}
	// Output:
	// {Z:5 X:19 Y:12}
	// {Z:5 X:19 Y:13}
	// {Z:5 X:20 Y:12}
	// {Z:5 X:20 Y:13}
	// {Z:5 X:21 Y:12}
	// {Z:5 X:21 Y:13}
	// {Z:6 X:39 Y:24}
	// {Z:6 X:39 Y:25}
	// {Z:6 X:39 Y:26}
	// {Z:6 X:39 Y:27}
	// {Z:6 X:40 Y:24}
	// {Z:6 X:40 Y:25}
	// {Z:6 X:40 Y:26}
	// {Z:6 X:40 Y:27}
	// {Z:6 X:41 Y:24}
	// {Z:6 X:41 Y:25}
	// {Z:6 X:41 Y:26}
	// {Z:6 X:41 Y:27}
	// {Z:6 X:42 Y:24}
	// {Z:6 X:42 Y:25}
	// {Z:6 X:42 Y:26}
	// {Z:6 X:42 Y:27}
}
