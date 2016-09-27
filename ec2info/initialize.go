package ec2info

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

func Initialize(sess *session.Session) (*ec2.DescribeInstancesOutput, error) {
	svc := ec2.New(sess)
	resp, err := svc.DescribeInstances(nil)

	return resp, err
}
