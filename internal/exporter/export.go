package exporter

import (
	"fmt"
	"math"
	"time"
)

func export(start int, end int, minimum int, count int) {
	hour := time.Now().UTC().Hour()
	value := 100
	if (start < end && hour >= start && hour <= end) || (start > end && (hour >= start || hour <= end)) {
		m := float64(minimum)
		c := float64(count)
		value = int(math.Max(math.Ceil((m/c)*100), 100))
	}
	// TODO: SEND TO CLOUD MONITORING
	fmt.Println(value)
}
