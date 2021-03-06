package cloudstorage

import (
	"context"
	"fmt"
	"io"

	"github.com/minio/minio-go/v6"
	"gitlab.com/Wuriyanto/go-codebase/pkg/shared"
)

// CloudStorageImpl implementation of CloudStorage
type CloudStorageImpl struct {
	client     *minio.Client
	bucketName string
}

// New function, CloudStorageImpl's constructor
func New(endpoint, bucketName, accessKey, secretKey string, tls bool) (*CloudStorageImpl, error) {
	client, err := minio.New(endpoint, accessKey, secretKey, tls)
	if err != nil {
		return nil, err
	}

	bucketExist, err := client.BucketExists(bucketName)
	if err != nil {
		return nil, err
	}

	if !bucketExist {
		err = client.MakeBucket(bucketName, "ap-southeast-1")
		if err != nil {
			return nil, err
		}
	}

	return &CloudStorageImpl{
		client:     client,
		bucketName: bucketName,
	}, nil
}

// Upload function will upload object into storage with context
func (c *CloudStorageImpl) Upload(context context.Context, name string, input io.Reader, options *Option) <-chan shared.Result {
	output := make(chan shared.Result)

	go func() {
		defer close(output)

		_, err := c.client.PutObjectWithContext(context, c.bucketName, name, input, options.ObjectSize, options.PutOptions)
		if err != nil {
			output <- shared.Result{Error: err}
			return
		}

		url, err := c.client.PresignedGetObject(c.bucketName, name, options.URLExpiry, options.ReqParams)
		if err != nil {
			output <- shared.Result{Error: err}
			return
		}

		output <- shared.Result{
			Data: url,
		}
	}()

	return output
}

// Download function will download object from storage with context
func (c *CloudStorageImpl) Download(context context.Context, name string, options *Option) <-chan shared.Result {
	output := make(chan shared.Result)

	go func() {
		defer close(output)

		//object = *minio.Object
		object, err := c.client.GetObjectWithContext(context, c.bucketName, name, options.GetOptions)
		if err != nil {
			output <- shared.Result{Error: err}
			return
		}

		output <- shared.Result{
			Data: object,
		}
	}()

	return output
}

// Remove function will remove object from storage with context
func (c *CloudStorageImpl) Remove(context context.Context, name string) <-chan shared.Result {
	output := make(chan shared.Result)

	go func() {
		defer close(output)

		err := c.client.RemoveObject(c.bucketName, name)
		if err != nil {
			output <- shared.Result{Error: err}
			return
		}

		output <- shared.Result{
			Data: fmt.Sprintf("%s removed", name),
		}
	}()

	return output
}

// Stat function will return object's stat from storage with context
func (c *CloudStorageImpl) Stat(context context.Context, name string, option *Option) <-chan shared.Result {
	output := make(chan shared.Result)

	go func() {
		defer close(output)

		// objectInfo = minio.ObjectInfo
		objectInfo, err := c.client.StatObjectWithContext(context, c.bucketName, name, option.StatOptions)
		if err != nil {
			output <- shared.Result{Error: err}
			return
		}

		output <- shared.Result{
			Data: objectInfo,
		}
	}()

	return output
}

// GetUrl function will return object's URL from storage with context
func (c *CloudStorageImpl) GetURL(context context.Context, name string) <-chan shared.Result {
	return nil
}

// GetObjects function will return object's List from storage with context
func (c *CloudStorageImpl) GetObjects(context context.Context, option *Option) <-chan shared.Result {
	output := make(chan shared.Result)
	done := make(chan struct{})

	go func() {
		defer close(output)
		defer close(done)

		var listObject []minio.ObjectInfo
		// object = minio.ObjectInfo
		for object := range c.client.ListObjectsV2(c.bucketName, option.PrefixSearchObject, option.RecursiveSearchObject, done) {
			if object.Err != nil {
				output <- shared.Result{Error: object.Err}
				return
			}
			listObject = append(listObject, object)
		}

		output <- shared.Result{
			Data: listObject,
		}
	}()

	return output
}
