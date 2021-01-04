package action

import (
	"fmt"
	"moniterinstance/lib/aws"

	"github.com/aws/aws-sdk-go/service/ec2"
)

func Run(profile, instanceName string) {
	sess := aws.Session(profile)
	instances := aws.DescribeInstances(sess)
	foundInstance := find(instances, instanceName)

	if foundInstance == nil {
		fmt.Println("Instances not found")
	} else {
		aws.GetCloudWatch(sess, *foundInstance.InstanceId)
	}
}

func find(instances *ec2.DescribeInstancesOutput, instanceName string) *ec2.Instance {
	var result *ec2.Instance
	for _, reservation := range instances.Reservations {
		for _, instance := range reservation.Instances {
			for _, tag := range instance.Tags {
				if *tag.Value == instanceName {
					result = instance
				}
			}
		}
	}
	return result
}
