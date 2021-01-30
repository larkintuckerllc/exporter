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

func export(projectID string, location string, namespace string, nodeID string,
	start int, end int, minimum int, replicas int32) error {
	var value int64 = 0
	hour := time.Now().UTC().Hour()
	if (start < end && hour >= start && hour < end) || (start > end && (hour >= start || hour < end)) {
		m := float64(minimum)
		r := float64(replicas)
		value = int64(math.Ceil((m / r) * 100))
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
		Name: monitoring.MetricProjectPath(projectID),
		TimeSeries: []*monitoringpb.TimeSeries{
			{
				Metric: &metricpb.Metric{
					Type: "custom.googleapis.com/exporter",
				},
				Resource: &monitoredrespb.MonitoredResource{
					Type: "generic_node",
					Labels: map[string]string{
						"project_id": projectID,
						"location":   location,
						"namespace":  namespace,
						"node_id":    nodeID,
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
