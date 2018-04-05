package storage

import (
	"bytes"
	"io/ioutil"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/pkg/errors"
	"github.com/swagchat/chat-api/utils"
)

type awss3Provider struct {
	accessKeyId        string
	secretAccessKey    string
	region             string
	acl                string
	uploadBucket       string
	uploadDirectory    string
	thumbnailBucket    string
	thumbnailDirectory string
}

func (ap *awss3Provider) Init() error {
	awsS3Client, err := ap.getSession()
	if err != nil {
		return errors.Wrap(err, "AWS S3 get session failure")
	}

	params := &s3.CreateBucketInput{
		Bucket: aws.String(ap.uploadBucket),
	}
	_, err = awsS3Client.CreateBucket(params)
	if err != nil {
		return errors.Wrap(err, "AWS S3 create bucket failure")
	}

	params = &s3.CreateBucketInput{
		Bucket: aws.String(ap.thumbnailBucket),
	}
	_, err = awsS3Client.CreateBucket(params)
	if err != nil {
		return errors.Wrap(err, "AWS S3 create thumbnail bucket failure")
	}

	return nil
}

func (ap *awss3Provider) Post(assetInfo *AssetInfo) (string, error) {
	awsS3Client, err := ap.getSession()
	if err != nil {
		return "", errors.Wrap(err, "AWS S3 get session failure")
	}

	byteData, err := ioutil.ReadAll(assetInfo.Data)
	if err != nil {
		return "", errors.Wrap(err, "AWS S3 read data failure")
	}

	data := bytes.NewReader(byteData)
	filePath := utils.AppendStrings(ap.uploadDirectory, "/", assetInfo.Filename)
	_, err = awsS3Client.PutObject(&s3.PutObjectInput{
		Bucket: aws.String(ap.uploadBucket),
		Key:    aws.String(filePath),
		Body:   data,
		ACL:    &ap.acl,
	})
	if err != nil {
		return "", errors.Wrap(err, "AWS S3 put object failure")
	}

	sourceUrl := utils.AppendStrings("https://s3-ap-northeast-1.amazonaws.com/", ap.uploadBucket, "/", filePath)
	return sourceUrl, nil
}

func (ap *awss3Provider) Get(assetInfo *AssetInfo) ([]byte, error) {
	return nil, nil
}

func (ap *awss3Provider) getSession() (*s3.S3, error) {
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(ap.region),
		Credentials: credentials.NewStaticCredentials(ap.accessKeyId, ap.secretAccessKey, ""),
	})
	if err != nil {
		return nil, errors.Wrap(err, "AWS S3 create session failure")
	}
	return s3.New(sess), nil
}
