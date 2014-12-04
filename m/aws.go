package monitor

import (
	"time"

	"github.com/crowdmob/goamz/aws"
	"github.com/crowdmob/goamz/cloudwatch"
)

// The CW type encapsulates cloudwatch.CloudWatch to provide additional method(s).
type CW struct {
	*cloudwatch.CloudWatch
}

// NewCW creates a new aws.Auth using credentials located by aws.GetAuth and returns a CW struct.
func NewCW(accessKey, secretKey, region string) (*CW, error) {
	now := time.Now()

	auth, err := aws.GetAuth(accessKey, secretKey, "", now)
	if err != nil {
		return nil, err
	}

	cw, err := cloudwatch.NewCloudWatch(auth, aws.Regions[region].CloudWatchServicepoint)
	return &CW{cw}, err
}

// PutMetrics saves provided information to CloudWatch and returns aws.BaseResponse.
func (cw *CW) PutMetrics(value uint64, unit string, name string, namespace string, autoScalingGroup string) (*aws.BaseResponse, error) {
	instanceId := aws.InstanceId()

	dimensions := []cloudwatch.Dimension{
		{Name: "InstanceID", Value: instanceId},
	}

	if autoScalingGroup != "" {
		dimensions = append(
			dimensions,
			cloudwatch.Dimension{
				Name: "AutoScalingGroupName", Value: autoScalingGroup,
			},
		)
	}

	metric := cloudwatch.MetricDatum{
		Dimensions: []cloudwatch.Dimension{
			{Name: "InstanceID", Value: instanceId},
		},
		MetricName: name,
		Unit:       unit,
		Value:      float64(value),
	}

	return cw.PutMetricDataNamespace([]cloudwatch.MetricDatum{metric}, namespace)
}
