package exporter

import (
	"context"
	"fmt"
	"math"
	"time"

	monitoring "cloud.google.com/go/monitoring/apiv3"
)

func export(start int, end int, minimum int, count int) error {
	hour := time.Now().UTC().Hour()
	value := 0
	if (start < end && hour >= start && hour <= end) || (start > end && (hour >= start || hour <= end)) {
		m := float64(minimum)
		c := float64(count)
		value = int(math.Ceil((m / c) * 100))
	}
	// TODO: SEND TO CLOUD MONITORING
	fmt.Println(value)

	ctx := context.Background()
	client, err := monitoring.NewMetricClient(ctx)
	if err != nil {
		return err
	}
	projectID := "YOUR_PROJECT_ID"
	fmt.Println(client)
	fmt.Println(projectID)
	return nil
}
