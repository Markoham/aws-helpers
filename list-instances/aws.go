package main

import (
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

var (
	empty = "--"
)

// Instance type
type Instance struct {
	ID               *string
	Name             *string
	IP               *string
	State            *string
	Type             *string
	StackName        *string
	AutoScalingGroup *string
	LaunchTime       *time.Time
}

func getInstances(region string) (*ec2.DescribeInstancesOutput, error) {
	svc := ec2.New(session.New(), aws.NewConfig().WithRegion(region))
	input := &ec2.DescribeInstancesInput{}

	result, err := svc.DescribeInstances(input)
	return result, err
}

func getTag(tags []*ec2.Tag, tag string) *string {
	total := len(tags)

	for i := 0; i < total; i++ {
		if *tags[i].Key == tag {
			return tags[i].Value
		}
	}

	return &empty
}

func safeIP(ip *string) *string {
	if ip == nil {
		return &empty
	}

	return ip
}

func parseResult(res *ec2.DescribeInstancesOutput) []*Instance {
	instances := make([]*Instance, 0)
	totalReservations := len(res.Reservations)

	for i := 0; i < totalReservations; i++ {
		reservation := res.Reservations[i]
		totalInstances := len(reservation.Instances)

		for j := 0; j < totalInstances; j++ {
			instance := reservation.Instances[j]

			instances = append(instances, &Instance{
				ID:               instance.InstanceId,
				Name:             getTag(instance.Tags, "Name"),
				IP:               safeIP(instance.PrivateIpAddress),
				State:            instance.State.Name,
				Type:             instance.InstanceType,
				LaunchTime:       instance.LaunchTime,
				StackName:        getTag(instance.Tags, "aws:cloudformation:stack-name"),
				AutoScalingGroup: getTag(instance.Tags, "aws:autoscaling:groupName"),
			})
		}
	}
	return instances
}
