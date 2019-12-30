package main

import (
	"fmt"
	"io"
	"log"
	"os"

	"github.com/minio/minio-go"
)

func newClient() *minio.Client {
	endpoint := "127.0.0.1:5080"
	accessKeyID := "0x6A6AEB9840C8a42A9d8Ff55cf2c213F9b812ED0A"
	secretAccessKey := "123456789"
	useSSL := false

	client, err := minio.New(endpoint, accessKeyID, secretAccessKey, useSSL)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return client
}

func makeBucket(minioClient *minio.Client) {
	err := minioClient.MakeBucket("testbucket1", "")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Successfully created mybucket.")
}

func listBuckets(minioClient *minio.Client) {
	buckets, err := minioClient.ListBuckets()
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, bucket := range buckets {
		fmt.Println(bucket)
	}
}

//**存在bug 当buckets不存在时，会阻塞进程
func bucketExists(minioClient *minio.Client) {
	found, err := minioClient.BucketExists("testbucket1")
	if err != nil {
		fmt.Println(err)
		return
	}
	if found {
		fmt.Println("Bucket found")
	}
}

func removeBuckets(minioClient *minio.Client) {
	err := minioClient.RemoveBucket("testbucket1")
	if err != nil {
		fmt.Println(err)
		return
	}
}

//**当bucket不存在时 阻塞进程
func listObjects(minioClient *minio.Client) {
	// Create a done channel to control 'ListObjects' go routine.
	doneCh := make(chan struct{})

	// Indicate to our routine to exit cleanly upon return.
	defer close(doneCh)

	isRecursive := true
	objectCh := minioClient.ListObjects("testbucket1", "o", isRecursive, doneCh)
	for object := range objectCh {
		if object.Err != nil {
			fmt.Println(object.Err)
			return
		}
		fmt.Println(object)
	}
}

//**bug同上
func listObjectsV2(minioClient *minio.Client) {
	doneCh := make(chan struct{})

	// Indicate to our routine to exit cleanly upon return.
	defer close(doneCh)

	isRecursive := true
	objectCh := minioClient.ListObjectsV2("testbucket2", "", isRecursive, doneCh)
	for object := range objectCh {
		if object.Err != nil {
			fmt.Println(object.Err)
			return
		}
		fmt.Println(object)
	}
}

//**未完成？
func ListIncompleteUploads(minioClient *minio.Client) {
	// Create a done channel to control 'ListObjects' go routine.
	doneCh := make(chan struct{})

	// Indicate to our routine to exit cleanly upon return.
	defer close(doneCh)

	isRecursive := true // Recursively list everything at 'myprefix'
	multiPartObjectCh := minioClient.ListIncompleteUploads("testbucket2", "c", isRecursive, doneCh)
	for multiPartObject := range multiPartObjectCh {
		if multiPartObject.Err != nil {
			fmt.Println(multiPartObject.Err)
			return
		}
		fmt.Println(multiPartObject)
	}
}

//======对象操作

//在/tmp 创建local-file.jpg，将目标数据拷贝到这个文件中
func getObject(minioClient *minio.Client) {
	object, err := minioClient.GetObject("testbucket2", "test.java", minio.GetObjectOptions{})
	if err != nil {
		fmt.Println(err)
		return
	}
	localFile, err := os.Create("/tmp/local-file.jpg")
	if err != nil {
		fmt.Println(err)
		return
	}
	if _, err = io.Copy(localFile, object); err != nil {
		fmt.Println(err)
		return
	}
}
func fGetObject(minioClient *minio.Client) {
	err := minioClient.FGetObject("testbucket2", "test.java", "myobject", minio.GetObjectOptions{})
	if err != nil {
		fmt.Println(err)
		return
	}
}
func putObject(minioClient *minio.Client) {
	file, err := os.Open("myobject")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	fileStat, err := file.Stat()
	if err != nil {
		fmt.Println(err)
		return
	}

	n, err := minioClient.PutObject("testbucket2", "myobject.test", file, fileStat.Size(), minio.PutObjectOptions{ContentType: "application/octet-stream"})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Successfully uploaded bytes: ", n)
}

//没有这项功能？
func copyObject(minioClient *minio.Client) {
	src := minio.NewSourceInfo("testbucket2", "test.java", nil)

	// Destination object
	dst, err := minio.NewDestinationInfo("testbucket2", "test.test", nil, nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Copy object call
	err = minioClient.CopyObject(dst, src)
	if err != nil {
		fmt.Println(err)
		return
	}
}

//当输入文件不存在时 报错
func statObject(minioClient *minio.Client) {
	objInfo, err := minioClient.StatObject("testbucket2", "ava", minio.StatObjectOptions{})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(objInfo)
}

func removeObject(minioClient *minio.Client) {
	objectsCh := make(chan string)

	// Send object names that are needed to be removed to objectsCh
	go func() {
		defer close(objectsCh)
		// List all objects from a bucket-name with a matching prefix.
		for object := range minioClient.ListObjects("testbucket2", "test.java", true, nil) {
			if object.Err != nil {
				log.Fatalln(object.Err)
			}
			objectsCh <- object.Key
		}
	}()

	for rErr := range minioClient.RemoveObjects("testbucket2", objectsCh) {
		fmt.Println("Error detected during deletion: ", rErr)
	}
}
