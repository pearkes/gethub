package main

import (
	"fmt"
)

func Version() []int {
	return []int{0, 0, 1}
}

func VersionString() string {
	v := Version()
	return fmt.Sprintf("%d.%02d.%02d", v[0], v[1], v[2])
}
