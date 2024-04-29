package progress

import (
	"fmt"
	"testing"
)

func TestIsAirplane(t *testing.T) {
	fmt.Println(IsAirplane("33344456"))
	fmt.Println(IsAirplane("36667778"))
	fmt.Println(IsAirplane("35888999"))
}

func TestIsAirplaneType1(t *testing.T) {
	fmt.Println(IsAirplane("33344456"))
}

func TestIsAirplaneType2(t *testing.T) {
	fmt.Println(IsAirplane("36667778"))
}

func TestIsAirplaneType3(t *testing.T) {
	fmt.Println(IsAirplane("35888999"))
}
