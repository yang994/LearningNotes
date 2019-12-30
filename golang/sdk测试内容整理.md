## 存储桶操作
+ MakeBucket() location字段具体内容需要添加

+ BucketExists() 检查桶是否存在，当传入的bucket name不存在时，报错并阻塞进程
   ```
   API: GetBucketLocation(bucket=testbucket1)
    Time: 14:01:31 CST 12/22/2019
    DeploymentID: 1a526bba-a52a-4f9f-ab4e-0f71d67afa3a
    RequestID: 15E29C2291FC32EE
    RemoteHost: 127.0.0.1
    Host: 127.0.0.1:5080
    UserAgent: MinIO (linux; amd64) minio-go/v6.0.45
    Error: bucket is not exist
       4: github.com/memoio/go-mefs/vendor/github.com/minio/minio/cmd/api-errors.go:1700:cmd.toAPIErrorCode()
       3: github.com/memoio/go-mefs/vendor/github.com/minio/minio/cmd/api-errors.go:1725:cmd.toAPIError()
       2: github.com/memoio/go-mefs/vendor/github.com/minio/minio/cmd/bucket-handlers.go:155:cmd.objectAPIHandlers.GetBucketLocationHandler()
       1: net/http/server.go:2007:http.HandlerFunc.ServeHTTP()
   ```
   与bucket查询相关的操作 如listObjects()都有同样的问题

+ ListIncompleteUploads()列举未完整上传的对象。没有返回值，应该是没有相应的功能，在文档中删去

## 对象操作

+ 删去CopyObject() ComposeObject() NewSourceInfo() NewDestinationInfo() 与拷贝对象有关的操作
+ getObject() statObject()等涉及桶和对象信息的操作，当传入不存在的桶名和对象名时，报错并阻塞进程
   ```
   API: HeadObject(bucket=testbucket2, object=ava)
   Time: 14:15:38 CST 12/22/2019
   DeploymentID: 1a526bba-a52a-4f9f-ab4e-0f71d67afa3a
   RequestID: 15E29CE7C3825958
   RemoteHost: 127.0.0.1
   Host: 127.0.0.1:5080
   UserAgent: MinIO (linux; amd64) minio-go/v6.0.45
   Error: object is not exist
   4: github.com/memoio/go-mefs/vendor/github.com/minio/minio/cmd/api-errors.go:1700:cmd.toAPIErrorCode()
   3: github.com/memoio/go-mefs/vendor/github.com/minio/minio/cmd/api-errors.go:1725:cmd.toAPIError()
   2: github.com/memoio/go-mefs/vendor/github.com/minio/minio/cmd/object-handlers.go:501:cmd.objectAPIHandlers.HeadObjectHandler()
   1: net/http/server.go:2007:http.HandlerFunc.ServeHTTP()

   ```

+ 删去RemoveIncompleteUpload()未完整上传对象操作

## 其他内容
+ 删去 加密对象、presigned、存储桶策略/通知、客户端自定义设置的内容