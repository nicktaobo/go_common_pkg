package tools

import (
	"fmt"
	"testing"
)

func TestGetLatLonDistance(t *testing.T) {
	var dis float64
	dis = GeoLatLonDistance(113.341488, 40.061826, 113.338452, 40.070959)
	fmt.Println("dis:", dis)
	dis = GeoLatLonDistance(104.067905, 30.631398, 104.068882, 30.628047)
	fmt.Println("dis:", dis)
	dis = GeoLatLonDistance(-73.992939, 40.714067, -73.994224, 40.715815)
	fmt.Println("dis:", dis)
}

func TestRandomInt(t *testing.T) {
	for i := 0; i < 100; i++ {
		fmt.Println(RandomInt(10, 12))
	}
}
