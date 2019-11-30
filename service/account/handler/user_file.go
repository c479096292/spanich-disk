package handler

import (
	"github.com/c479096292/spinach-disk/common"
	"context"
	proto "github.com/c479096292/spinach-disk/service/account/proto"
	"github.com/c479096292/spinach-disk/service/dbproxy/client"
	"encoding/json"
)

// UserFiles : 获取用户文件列表
func (u *User) UserFiles(ctx context.Context, req *proto.ReqUserFile, res *proto.RespUserFile) error {
	dbResp, err := client.QueryUserFileMetas(req.Username, int(req.Limit))
	if err != nil || !dbResp.Suc {
		res.Code = common.StatusServerError
		return err
	}

	userFiles := client.ToTableUserFiles(dbResp.Data)
	data, err := json.Marshal(userFiles)
	if err != nil {
		res.Code = common.StatusServerError
		return nil
	}

	res.FileData = data
	return nil
}

// UserFiles : 用户文件重命名
func (u *User) UserFileRename(ctx context.Context, req *proto.ReqUserFileRename, res *proto.RespUserFileRename) error {
	dbResp, err := client.RenameFileName(req.Username, req.Filehash, req.NewFileName)
	if err != nil || !dbResp.Suc {
		res.Code = common.StatusServerError
		return err
	}

	userFiles := client.ToTableUserFiles(dbResp.Data)
	data, err := json.Marshal(userFiles)
	if err != nil {
		res.Code = common.StatusServerError
		return nil
	}

	res.FileData = data
	return nil
}