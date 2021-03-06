package file

import (
	"fmt"
	"io"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

// File ファイル構造体
type File struct {
	client *s3.S3
	config *aws.Config
}

// NewFile ファイルインターフェースを生成します
func NewFile(config *aws.Config) *File {
	return &File{
		config: config,
	}
}

func (f *File) connect() error {
	if f.client == nil {
		sess, err := session.NewSession(f.config)
		if err != nil {
			return err
		}
		f.client = s3.New(sess)
	}
	return nil
}

// GetObject ファイルを取得します
func (f *File) GetObject(bucket, key string) (io.ReadCloser, error) {
	err := f.connect()
	if err != nil {
		return nil, fmt.Errorf("Error: couldn't connect S3, %v", err)
	}

	obj, err := f.client.GetObject(&s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	})
	if err != nil {
		return nil, fmt.Errorf(
			"Error: couldn't GetObject Object: %s/%s, %v",
			*aws.String(bucket),
			*aws.String(key),
			err,
		)
	}

	return obj.Body, nil
}

// MoveObject ファイルを移動します
func (f *File) MoveObject(srcBucket, srcKey, dstBucket, dstKey string) error {
	err := f.connect()
	if err != nil {
		return err
	}

	_, err = f.client.CopyObject(&s3.CopyObjectInput{
		CopySource: aws.String(fmt.Sprintf("%s/%s", srcBucket, srcKey)),
		Bucket:     aws.String(dstBucket),
		Key:        aws.String(dstKey),
	})
	if err != nil {
		return err
	}

	return nil
}
