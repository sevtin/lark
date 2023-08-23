package xs3

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"lark/pkg/conf"
	"lark/pkg/constant"
	"lark/pkg/utils"
	"mime/multipart"
	"os"
	"path"
)

const (
	S3_ACL_PUBLIC_READ = "public-read"
)

var client *S3Client

type S3Client struct {
	cfg  *conf.AwsS3
	sess *session.Session
}

func NewAwsS3(cfg *conf.AwsS3) (cli *S3Client, err error) {
	var (
		sess *session.Session
	)
	os.Setenv("AWS_REGION", cfg.Region)
	os.Setenv("AWS_ACCESS_KEY_ID", cfg.AccessKey)
	os.Setenv("AWS_SECRET_ACCESS_KEY", cfg.SecretKey)
	sess, err = session.NewSession(&aws.Config{
		Credentials:      credentials.NewStaticCredentials(cfg.AccessKey, cfg.SecretKey, ""),
		Endpoint:         aws.String(cfg.EndPoint),
		Region:           aws.String(cfg.Region),
		DisableSSL:       aws.Bool(true),
		S3ForcePathStyle: aws.Bool(false), //virtual-host style方式
	})
	if err != nil {
		panic(err)
	}
	sess = session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))
	cli = &S3Client{cfg: cfg, sess: sess}
	client = cli
	return
}

func (c *S3Client) Upload(fileHeader *multipart.FileHeader) (output *s3manager.UploadOutput, err error) {
	var (
		file        multipart.File
		uploader    *s3manager.Uploader
		uuid        = utils.NewUUID()
		suffix      string
		contentType string
		key         string
	)
	file, err = fileHeader.Open()
	if err != nil {
		return
	}
	defer file.Close()

	suffix = path.Ext(fileHeader.Filename)
	contentType = utils.GetContentType(suffix)
	key = uuid + suffix

	uploader = s3manager.NewUploader(c.sess)
	output, err = uploader.Upload(&s3manager.UploadInput{
		Bucket:      aws.String(c.cfg.Bucket),
		Key:         aws.String(key),
		Body:        file,
		ContentType: aws.String(contentType),
		ACL:         aws.String(S3_ACL_PUBLIC_READ),
	})
	return
}

// https://docs.aws.amazon.com/zh_cn/sdk-for-go/v1/developer-guide/s3-example-presigned-urls.html
func (c *S3Client) GetPresignedURL(in *PresignedUrlInput) (out *PresignedUrlOutput, err error) {
	var (
		key    = utils.NewUUID()
		suffix string
		svc    = s3.New(c.sess)
		req    *request.Request
	)
	out = new(PresignedUrlOutput)
	if in.ContentType == "" {
		suffix = path.Ext(in.Filename)
		in.ContentType = utils.GetContentType(suffix)
	}
	req, _ = svc.PutObjectRequest(&s3.PutObjectInput{
		Bucket:      aws.String(c.cfg.Bucket),
		Key:         aws.String(key),
		ACL:         aws.String(S3_ACL_PUBLIC_READ),
		ContentType: aws.String(in.ContentType),
	})
	out.PutUrl, err = req.Presign(constant.CONST_DURATION_AWS_S3_EXPIRE_MINUTE)
	out.Url = c.cfg.EndPoint + "/" + key
	out.Key = key
	out.ContentType = in.ContentType
	out.Acl = S3_ACL_PUBLIC_READ
	out.Filename = in.Filename
	return
}

/*
curl --location --request PUT 'https://gosportsbucket.s3.us-west-1.amazonaws.com/images/1a3a88e2-ff83-11ed-8226-b42e9910a730.png?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=AKIAQNJTZPCIM2O26EJG/20230531/us-west-1/s3/aws4_request&X-Amz-Date=20230531T071646Z&X-Amz-Expires=900&X-Amz-SignedHeaders=content-type;host;x-amz-acl&X-Amz-Signature=b026a33a7cfd53736d0ab22365069f7bafe100bf5b6741aa72b0c7d391994fe3' \
--header 'Content-Type: image/png' \
--header 'X-Amz-Acl: public-read' \
--data-binary '@/Users/saeipi/Desktop/99d28d55-12c9-4d98-ad0a-eb26aede2f62.png'
*/
