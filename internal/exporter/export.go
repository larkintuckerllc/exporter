package exporter

import (
	"context"
	"math"
	"time"

	monitoring "cloud.google.com/go/monitoring/apiv3"
	googlepb "github.com/golang/protobuf/ptypes/timestamp"
	metricpb "google.golang.org/genproto/googleapis/api/metric"
	monitoredrespb "google.golang.org/genproto/googleapis/api/monitoredres"
	monitoringpb "google.golang.org/genproto/googleapis/monitoring/v3"
)

func export(start int, end int, minimum int, count int, project string,
	namespace string, service string) error {
	hour := time.Now().UTC().Hour()
	var value int64 = 0
	if (start < end && hour >= start && hour <= end) || (start > end && (hour >= start || hour <= end)) {
		m := float64(minimum)
		c := float64(count)
		value = int64(math.Ceil((m / c) * 100))
	}
	ctx := context.Background()
	client, err := monitoring.NewMetricClient(ctx)
	if err != nil {
		return err
	}
	dataPoint := &monitoringpb.Point{
		Interval: &monitoringpb.TimeInterval{
			EndTime: &googlepb.Timestamp{
				Seconds: time.Now().Unix(),
			},
		},
		Value: &monitoringpb.TypedValue{
			Value: &monitoringpb.TypedValue_Int64Value{
				Int64Value: value,
			},
		},
	}
	err = client.CreateTimeSeries(ctx, &monitoringpb.CreateTimeSeriesRequest{
		Name: monitoring.MetricProjectPath(project),
		TimeSeries: []*monitoringpb.TimeSeries{
			{
				Metric: &metricpb.Metric{
					Type: "custom.googleapis.com/exporter",
				},
				Resource: &monitoredrespb.MonitoredResource{
					Type: "global",
					Labels: map[string]string{
						"project_id":     project,
						"location":       "us-central1",
						"cluster_name":   "cluster-1",
						"namespace_name": namespace,
						"pod_name":       "FUCK ME",
					},
				},
				Points: []*monitoringpb.Point{
					dataPoint,
				},
			},
		},
	})
	if err != nil {
		return err
	}
	err = client.Close()
	if err != nil {
		return err
	}
	return nil
}
