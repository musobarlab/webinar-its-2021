package cloudstorage

import (
	"context"
	"io"
	"net/url"
	"time"

	"github.com/minio/minio-go/v6"
	"gitlab.com/Wuriyanto/go-codebase/pkg/shared"
)

// Option of cloud storage
type Option struct {
	ObjectSize            int64
	PutOptions            minio.PutObjectOptions
	GetOptions            minio.GetObjectOptions
	StatOptions           minio.StatObjectOptions
	URLExpiry             time.Duration
	ReqParams             url.Values
	PrefixSearchObject    string
	RecursiveSearchObject bool
}

// CloudStorage abstract type
type CloudStorage interface {
	Upload(context context.Context, name string, input io.Reader, options *Option) <-chan shared.Result
	Download(context context.Context, name string, options *Option) <-chan shared.Result
	Remove(context context.Context, name string) <-chan shared.Result
	Stat(context context.Context, name string, option *Option) <-chan shared.Result
	GetURL(context context.Context, name string) <-chan shared.Result
	GetObjects(context context.Context, option *Option) <-chan shared.Result
}
