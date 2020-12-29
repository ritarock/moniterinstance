package aws

import (
	"fmt"
	"log"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/cloudwatch"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

func Session(profile string) *session.Session {
	return session.Must(session.NewSessionWithOptions(session.Options{
		Profile:           profile,
		SharedConfigState: session.SharedConfigEnable,
	}))
}

func DescribeInstances(session *session.Session) *ec2.DescribeInstancesOutput {
	svc := ec2.New(session)
	instances, err := svc.DescribeInstances(nil)

	if err != nil {
		log.Fatal(err)
	}
	return instances
}

func GetCloudWatch(session *session.Session, instanceId string) {
	svc := cloudwatch.New(session)
	params := &cloudwatch.GetMetricStatisticsInput{
		EndTime:    aws.Time(time.Now()),
		StartTime:  aws.Time(time.Now().Add(time.Duration(1) * time.Hour * -1)),
		MetricName: aws.String("CPUUtilization"),
		Namespace:  aws.String("AWS/EC2"),
		Period:     aws.Int64(60),
		Statistics: []*string{
			aws.String(cloudwatch.StatisticMaximum),
			aws.String(cloudwatch.StatisticAverage),
		},
		Dimensions: []*cloudwatch.Dimension{
			{
				Name:  aws.String("InstanceId"),
				Value: aws.String(instanceId),
			},
		},
		Unit: aws.String(cloudwatch.StandardUnitPercent),
	}

	resp, err := svc.GetMetricStatistics(params)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(resp)
}
