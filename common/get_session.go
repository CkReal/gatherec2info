package common

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
)

func GetSession(id string, region string) (*session.Session, error) {
	sess, err := session.NewSessionWithOptions(
		session.Options{
			Config:  aws.Config{Region: aws.String(region)},
			Profile: id,
		})

	return sess, err
}
