package exporter

import (
	"fmt"
	"math"
)

func export(start int, end int, minimum int, count int) {
	// TODO: USE TIME
	m := float64(minimum)
	c := float64(count)
	value := int(math.Max(math.Ceil((m/c)*100), 100))
	// TODO: SEND TO CLOUD MONITORING
	fmt.Println(value)
}
