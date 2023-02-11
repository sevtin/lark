package xminio

import (
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"lark/pkg/common/xlog"
	"lark/pkg/utils"
	"log"
	"net/url"
	"time"
)

/*
http://docs.minio.org.cn/docs/master/golang-client-quickstart-guide
*/

import (
	"context"
	"github.com/minio/minio-go/v7/pkg/encrypt"
	"io"
)

const (
	BUCKET_DOCUMENTS = "documents"
	BUCKET_PHOTOS    = "photos"
	BUCKET_VIDEOS    = "videos"
)

const (
	FILE_TYPE_DOCUMENT = "document"
	FILE_TYPE_PHOTO    = "photo"
	FILE_TYPE_VIDEO    = "video"
)

const (
	CONST_DURATION_PRESIGNED_EXPIRES = time.Hour * 1
)

type MinioConfig struct {
	Endpoint string   `yaml:"endpoint"`
	Options  *Options `yaml:"options"`
	Bucket   *Bucket  `yaml:"bucket"`
}
type Credentials struct {
	AccessKey string `yaml:"access_key"`
	SecretKey string `yaml:"secret_key"`
	Token     string `yaml:"token"`
}
type Options struct {
	Secure      bool         `yaml:"secure"`
	Credentials *Credentials `yaml:"credentials"`
}
type MakeBucketOptions struct {
	Region        string `yaml:"region"`
	ObjectLocking bool   `yaml:"objectLocking"`
}
type PutObjectOptions struct {
	Encrypted bool `yaml:"encrypted"`
}
type Bucket struct {
	BucketNames       []string           `yaml:"bucket_names"`
	MakeBucketOptions *MakeBucketOptions `yaml:"make_bucket_options"`
	PutObjectOptions  *PutObjectOptions  `yaml:"put_object_options"`
}

type PutResult struct {
	Info minio.UploadInfo
	Err  error
}

type PutResultList struct {
	Err  error
	List []*PutResult
}

var (
	minioc *MinioClient
)

type MinioClient struct {
	client *minio.Client
}

func init() {
	var conf = new(MinioConfig)
	utils.YamlToStruct("./configs/minio.yaml", conf)
	NewMinioClient(conf)
}

func NewMinioClient(conf *MinioConfig) {
	var (
		opts            minio.MakeBucketOptions
		client          *minio.Client
		bucketName      string
		exists          bool
		err             error
		errBucketExists error
	)

	client, err = minio.New(conf.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(conf.Options.Credentials.AccessKey, conf.Options.Credentials.SecretKey, conf.Options.Credentials.Token),
		Secure: false,
	})
	if err != nil {
		log.Print("instantiate minio client error:", err)
		return
	}
	minioc = &MinioClient{client: client}

	opts = minio.MakeBucketOptions{
		Region:        conf.Bucket.MakeBucketOptions.Region,
		ObjectLocking: conf.Bucket.MakeBucketOptions.ObjectLocking,
	}
	ctx := context.Background()
	for _, bucketName = range conf.Bucket.BucketNames {
		err = client.MakeBucket(ctx, bucketName, opts)
		if err != nil {
			// Check to see if we already own this bucket (which happens if you run this twice)
			exists, errBucketExists = client.BucketExists(ctx, bucketName)
			if errBucketExists == nil && exists {
				xlog.Warn("We already own", bucketName)
			} else {
				xlog.Error(err.Error())
			}
		} else {
			xlog.Info("Successfully created", bucketName)
		}
	}
	return
}

func Put(fileType string, objectName string, reader io.Reader, objectSize int64, contentType string) (info minio.UploadInfo, err error) {
	var bucketName = BUCKET_DOCUMENTS
	switch fileType {
	case FILE_TYPE_DOCUMENT:
		bucketName = BUCKET_DOCUMENTS
	case FILE_TYPE_PHOTO:
		bucketName = BUCKET_PHOTOS
	case FILE_TYPE_VIDEO:
		bucketName = BUCKET_VIDEOS
	}
	return PutObject(bucketName, objectName, reader, objectSize, false, contentType)
}

func PutObject(bucketName string, objectName string, reader io.Reader, objectSize int64, encrypted bool, contentType string) (info minio.UploadInfo, err error) {
	return minioc.client.PutObject(context.Background(), bucketName, objectName, reader, objectSize, getPutObjectOptions(encrypted, contentType))
}

func FPut(fileType string, objectName, filePath string, contentType string) (info minio.UploadInfo, err error) {
	var bucketName = BUCKET_DOCUMENTS
	switch fileType {
	case FILE_TYPE_DOCUMENT:
		bucketName = BUCKET_DOCUMENTS
	case FILE_TYPE_PHOTO:
		bucketName = BUCKET_PHOTOS
	case FILE_TYPE_VIDEO:
		bucketName = BUCKET_VIDEOS
	}
	return FPutObject(bucketName, objectName, filePath, getPutObjectOptions(false, contentType))
}

func FPutObject(bucketName, objectName, filePath string, opts minio.PutObjectOptions) (info minio.UploadInfo, err error) {
	return minioc.client.FPutObject(context.Background(), bucketName, objectName, filePath, opts)
}

func Presigned(fileType string, objectName string) (url *url.URL, err error) {
	var bucketName = BUCKET_DOCUMENTS
	switch fileType {
	case FILE_TYPE_DOCUMENT:
		bucketName = BUCKET_DOCUMENTS
	case FILE_TYPE_PHOTO:
		bucketName = BUCKET_PHOTOS
	case FILE_TYPE_VIDEO:
		bucketName = BUCKET_VIDEOS
	}
	return PresignedPutObject(bucketName, objectName, CONST_DURATION_PRESIGNED_EXPIRES)
}

func PresignedPutObject(bucketName string, objectName string, expires time.Duration) (url *url.URL, err error) {
	return minioc.client.PresignedPutObject(context.Background(), bucketName, objectName, expires)
}

func getPutObjectOptions(encrypted bool, contentType string) minio.PutObjectOptions {
	options := minio.PutObjectOptions{}
	if encrypted {
		options.ServerSideEncryption = encrypt.NewSSE()
	}
	options.ContentType = contentType
	return options
}

func PresignedPostPolicy(bucketName string, objectName string, uid int64, suffix string) (url *url.URL, formData map[string]string, err error) {
	policy := minio.NewPostPolicy()
	policy.SetBucket(bucketName)
	policy.SetKey(objectName)
	policy.SetExpires(time.Now().Add(CONST_DURATION_PRESIGNED_EXPIRES))
	policy.SetContentType(utils.GetContentType(suffix))
	// Only allow content size in range 1KB to 1MB.
	policy.SetContentLengthRange(1024, 1024*1024*1024)
	// Add a user metadata using the key "custom" and value "user"
	policy.SetUserMetadata(utils.Int64ToStr(uid), utils.NewUUID())
	// Get the POST form key/value object:
	return minioc.client.PresignedPostPolicy(context.Background(), policy)
}
