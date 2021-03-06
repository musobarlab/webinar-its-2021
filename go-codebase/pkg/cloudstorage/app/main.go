package main

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/minio/minio-go/v6"

	"gitlab.com/Wuriyanto/go-codebase/pkg/cloudstorage"
)

func main() {
	fmt.Println("---++++---")

	cloudStorage, err := cloudstorage.New("localhost:9000", "my-bucket-one", "minion", "123456789", false)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// ----- download object -----

	downloadResult := <-cloudStorage.Download(context.Background(), "bot_test.mov", &cloudstorage.Option{GetOptions: minio.GetObjectOptions{}})
	if downloadResult.Error != nil {
		fmt.Println(downloadResult.Error)
		os.Exit(1)
	}

	objectReader, ok := downloadResult.Data.(*minio.Object)
	if !ok {
		fmt.Println("error parsing object's value")
		os.Exit(1)
	}

	fileStat, err := objectReader.Stat()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fileName := fmt.Sprintf("%s_%s", time.Now().String(), fileStat.Key)

	outFile, err := os.Create(fileName)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer func() {
		outFile.Close()
	}()

	bufferDownload := make([]byte, 1024)
	reader := bufio.NewReader(objectReader)
	for {
		fmt.Print(".")
		line, err := reader.Read(bufferDownload)
		if err != nil {
			if err == io.EOF {
				break
			} else {
				fmt.Println(err)
				os.Exit(1)
			}
		}

		_, err = outFile.Write(bufferDownload[:line])
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}

	// ----- stat object -----
	// statResult := <-cloudStorage.Stat(context.Background(), "bot_test.mov", &cloudstorage.Option{StatOptions: minio.StatObjectOptions{}})
	// if statResult.Error != nil {
	// 	fmt.Println(statResult.Error)
	// 	os.Exit(1)
	// }

	// stat, ok := statResult.Data.(minio.ObjectInfo)
	// if !ok {
	// 	fmt.Println("error parsing object's stat")
	// 	os.Exit(1)
	// }

	// fmt.Println(stat.Size)

	// ----- remove object -----

	// removeResult := <-cloudStorage.Remove(context.Background(), "bot_test.mov")
	// if removeResult.Error != nil {
	// 	fmt.Println(removeResult.Error)
	// 	os.Exit(1)
	// }

	// removed, ok := removeResult.Data.(string)
	// if !ok {
	// 	fmt.Println("error parsing removed object")
	// 	os.Exit(1)
	// }

	// fmt.Println(removed)

	// ----- upload object -----

	// file, err := os.Open("/Users/wuriyanto/Desktop/bot_test.mov")
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// defer file.Close()

	// fileStat, err := file.Stat()
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// uploadResult := <-cloudStorage.Upload(context.Background(), fileStat.Name(), file, &cloudstorage.Option{
	// 	ObjectSize: -1,
	// 	URLExpiry:  time.Second * 24 * 60 * 60, // one day
	// 	PutOptions: minio.PutObjectOptions{ContentType: "application/octet-stream"},
	// })

	// if uploadResult.Error != nil {
	// 	fmt.Println(uploadResult.Error)
	// 	os.Exit(1)
	// }

	// url, ok := uploadResult.Data.(*url.URL)
	// if !ok {
	// 	fmt.Println("error parsing object's url")
	// 	os.Exit(1)
	// }

	// fmt.Println(url)

	// ----- list object -----

	// 	objectResult := <-cloudStorage.GetObjects(context.Background(), &cloudstorage.Option{PrefixSearchObject: "", RecursiveSearchObject: false})

	// 	if objectResult.Error != nil {
	// 		fmt.Println(err)
	// 		os.Exit(1)
	// 	}

	// 	listObject, ok := objectResult.Data.([]minio.ObjectInfo)
	// 	if !ok {
	// 		fmt.Println("error parsing object result")
	// 		os.Exit(1)
	// 	}

	// 	for _, object := range listObject {
	// 		fmt.Println(object.Key)
	// 		fmt.Println(object.LastModified)
	// 	}
}
