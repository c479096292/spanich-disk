package handler

import (
	"context"
	cfg "github.com/c479096292/spinach-disk/service/download/config"
	dlProto "github.com/c479096292/spinach-disk/service/download/proto"
)

// Dwonload :download结构体
type Download struct{}

// DownloadEntry : 获取下载入口
func (u *Download) DownloadEntry(
	ctx context.Context,
	req *dlProto.ReqEntry,
	res *dlProto.RespEntry) error {

	res.Entry = cfg.DownloadEntry
	return nil
}