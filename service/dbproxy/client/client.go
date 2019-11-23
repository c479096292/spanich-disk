package client

import (
	"context"
	"encoding/json"
	"github.com/c479096292/spinach-disk/service/dbproxy/model"
	proto "github.com/c479096292/spinach-disk/service/dbproxy/proto"
	"github.com/micro/go-micro"
	"github.com/mitchellh/mapstructure"
)

// FileMeta : 文件元信息结构
type FileMeta struct {
	FileSha1 string
	FileName string
	FileSize int64
	Location string
	UploadAt string
}

var (
	dbCli proto.DBProxyService
)

func Init(service micro.Service) {
	// 初始化一个dbproxy服务的客户端
	dbCli = proto.NewDBProxyService("go.micro.service.dbproxy", service.Client())
}

// 文件元数据信息
func TableFileToFileMeta(tfile model.TableFile) FileMeta {
	return FileMeta{
		FileSha1: tfile.FileHash,
		FileName: tfile.FileName.String,
		FileSize: tfile.FileSize.Int64,
		Location: tfile.FileAddr.String,
	}
}

// execAction : 向dbproxy请求执行action
func execAction(funcName string, paramJson []byte) (*proto.RespExec, error) {
	return dbCli.ExecuteAction(context.TODO(), &proto.ReqExec{
		Action: []*proto.SingleAction{
			&proto.SingleAction{
				Name:   funcName,
				Params: paramJson,
			},
		},
	})
}

// parseBody : 转换rpc返回的结果
func parseBody(resp *proto.RespExec) *model.ExecResult {
	if resp == nil || resp.Data == nil {
		return nil
	}
	resList := []model.ExecResult{}
	_ = json.Unmarshal(resp.Data, &resList)
	// TD
	if len(resList) > 0 {
		return &resList[0]
	}
	return nil
}

func ToTableUser(src interface{}) model.TableUser {
	user := model.TableUser{}
	mapstructure.Decode(src, &user)
	return user
}

func ToTableFile(src interface{}) model.TableFile {
	file := model.TableFile{}
	mapstructure.Decode(src, &file)
	return file
}

func ToTableFiles(src interface{}) []model.TableFile {
	file := []model.TableFile{}
	mapstructure.Decode(src, &file)
	return file
}

func ToTableUserFile(src interface{}) model.TableUserFile {
	ufile := model.TableUserFile{}
	mapstructure.Decode(src, &ufile)
	return ufile
}

func ToTableUserFiles(src interface{}) []model.TableUserFile {
	ufile := []model.TableUserFile{}
	mapstructure.Decode(src, &ufile)
	return ufile
}

func GetFileMeta(filehash string) (*model.ExecResult, error) {
	uInfo, _ := json.Marshal([]string{filehash})
	res, err := execAction("/file/GetFileMeta", uInfo)
	return parseBody(res), err
}

func GetFileMetaList(limitCnt int) (*model.ExecResult, error) {
	uInfo, _ := json.Marshal([]interface{}{limitCnt})
	res, err := execAction("/file/GetFileMetaList", uInfo)
	return parseBody(res), err
}

// OnFileUploadFinished : 新增/更新文件元信息到mysql中
func OnFileUploadFinished(fmeta FileMeta) (*model.ExecResult, error) {
	uInfo, _ := json.Marshal([]interface{}{fmeta.FileSha1, fmeta.FileName, fmeta.FileSize, fmeta.Location})
	res, err := execAction("/file/OnFileUploadFinished", uInfo)
	return parseBody(res), err
}

func UpdateFileLocation(filehash, location string) (*model.ExecResult, error) {
	uInfo, _ := json.Marshal([]interface{}{filehash, location})
	res, err := execAction("/file/UpdateFileLocation", uInfo)
	return parseBody(res), err
}

func UserSignup(username, encPasswd string) (*model.ExecResult, error) {
	uInfo, _ := json.Marshal([]interface{}{username, encPasswd})
	res, err := execAction("/user/UserSignup", uInfo)
	return parseBody(res), err
}

func UserSignin(username, encPasswd string) (*model.ExecResult, error) {
	uInfo, _ := json.Marshal([]interface{}{username, encPasswd})
	res, err := execAction("/user/UserSignin", uInfo)
	return parseBody(res), err
}

func GetUserInfo(username string) (*model.ExecResult, error) {
	uInfo, _ := json.Marshal([]interface{}{username})
	res, err := execAction("/user/GetUserInfo", uInfo)
	return parseBody(res), err
}

func UserExist(username string) (*model.ExecResult, error) {
	uInfo, _ := json.Marshal([]interface{}{username})
	res, err := execAction("/user/UserExist", uInfo)
	return parseBody(res), err
}

func UpdateToken(username, token string) (*model.ExecResult, error) {
	uInfo, _ := json.Marshal([]interface{}{username, token})
	res, err := execAction("/user/UpdateToken", uInfo)
	return parseBody(res), err
}

func QueryUserFileMeta(username, filehash string) (*model.ExecResult, error) {
	uInfo, _ := json.Marshal([]interface{}{username, filehash})
	res, err := execAction("/ufile/QueryUserFileMeta", uInfo)
	return parseBody(res), err
}

func QueryUserFileMetas(username string, limit int) (*model.ExecResult, error) {
	uInfo, _ := json.Marshal([]interface{}{username, limit})
	res, err := execAction("/ufile/QueryUserFileMetas", uInfo)
	return parseBody(res), err
}

// OnUserFileUploadFinished : 新增/更新文件元信息到mysql中
func OnUserFileUploadFinished(username string, fmeta FileMeta) (*model.ExecResult, error) {
	uInfo, _ := json.Marshal([]interface{}{username, fmeta.FileSha1,
		fmeta.FileName, fmeta.FileSize})
	res, err := execAction("/ufile/OnUserFileUploadFinished", uInfo)
	return parseBody(res), err
}

func RenameFileName(username, filehash, filename string) (*model.ExecResult, error) {
	uInfo, _ := json.Marshal([]interface{}{username, filehash, filename})
	res, err := execAction("/ufile/RenameFileName", uInfo)
	return parseBody(res), err
}
