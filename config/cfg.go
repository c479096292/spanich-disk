package config

import (
	"github.com/c479096292/spinach-disk/common"
	"os"
)

const (
	// PasswordSalt : 加密盐
	// 更严格来说， 加密盐可以存放在数据库中，每个用户加密盐不一样
	PasswordSalt = "*#890"

	//----------------------------ceph---------------------------
	// CephAccessKey : 访问Key
	CephAccessKey = "PFEA7NXWXSOWVTFA16C9"
	// CephSecretKey : 访问密钥
	CephSecretKey = "cf3dwPMeadGbtEgwFUEA6emRVrVfDHpv0pLXFYby"
	// CephGWEndpoint : gateway地址
	CephGWEndpoint = "http://<你的rgw_host>:<<你的rgw_port>>"

	//----------------------------oss---------------------------
	// OSSBucket : oss bucket名
	OSSBucket = "buckettest-filestore2"
	// OSSEndpoint : oss endpoint
	OSSEndpoint = "oss-cn-shenzhen.aliyuncs.com"
	// OSSAccesskeyID : oss访问key
	OSSAccesskeyID = "<你的AccesskeyId>"
	// OSSAccessKeySecret : oss访问key secret
	OSSAccessKeySecret = "<你的AccessKeySecret>"

	//----------------------------rabbitmq---------------------------
	// AsyncTransferEnable : 是否开启文件异步转移(默认同步)
	AsyncTransferEnable = false
	// TransExchangeName : 用于文件transfer的交换机
	TransExchangeName = "uploadserver.trans"
	// TransOSSQueueName : oss转移队列名
	TransOSSQueueName = "uploadserver.trans.oss"
	// TransOSSErrQueueName : oss转移失败后写入另一个队列的队列名
	TransOSSErrQueueName = "uploadserver.trans.oss.err"
	// TransOSSRoutingKey : routingkey
	TransOSSRoutingKey = "oss"

	//----------------------------service---------------------------
	// UploadServiceHost : 上传服务监听的地址
	UploadServiceHost = "0.0.0.0:8080"
	// UploadLBHost: 上传服务LB地址
	UploadLBHost = "http://upload.fileserver.com"
	// DownloadLBHost: 下载服务LB地址
	DownloadLBHost = "http://download.fileserver.com"
	// TracerAgentHost: tracing agent地址
	TracerAgentHost = "127.0.0.1:6831"

	//----------------------------store---------------------------

	// CephRootDir : Ceph的存储路径prefix
	CephRootDir = "/ceph"
	// OSSRootDir : OSS的存储路径prefix
	OSSRootDir = "oss/"
	// CurrentStoreType : 设置当前文件的存储类型
	CurrentStoreType = common.StoreLocal
)


var (
	workdir, _ = os.Getwd()
	// RabbitURL : rabbitmq服务的入口url
	RabbitURL = "amqp://guest:guest@127.0.0.1:5672/"

	// TempLocalRootDir : 本地临时存储地址的路径
	TempLocalRootDir =  workdir + "/data/fileserver/"
	// TempPartRootDir : 分块文件在本地临时存储地址的路径
	TempPartRootDir = workdir + "/data/fileserver_part/"
)