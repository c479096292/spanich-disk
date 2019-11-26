package handler

import (
	"context"
	upProto "github.com/c479096292/spinach-disk/service/upload/proto"
	cfg "github.com/c479096292/spinach-disk/service/upload/config"
)

// Upload : upload结构体
type Upload struct{}

// UploadEntry : 获取上传入口
func (u *Upload) UploadEntry(
	ctx context.Context,
	req *upProto.ReqEntry,
	res *upProto.RespEntry) error {

	res.Entry = cfg.UploadEntry
	return nil
}