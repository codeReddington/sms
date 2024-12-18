package config

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"os"
)

func NewAwsSession() *s3.S3 {
	sess := session.Must(session.NewSession(&aws.Config{
		CredentialsChainVerboseErrors: aws.Bool(true),
		Credentials:                   credentials.NewStaticCredentials(os.Getenv("S3_ACCESS_ID"), os.Getenv("S3_BUCKET_KEY"), ""),
		Region:                        aws.String(os.Getenv("AWS_REGION")),
	}))
	return s3.New(sess)
}
