package chatbiz

import (
	chatmodel "Blog-CMS/module/chat/model"
)

type ChatStorage interface {
	CallDeepSeekAPI(reqData chatmodel.DeepSeekRequest) (*chatmodel.DeepSeekResponse, error)
}
type deepSeekBiz struct {
	storage ChatStorage
}

func NewDeepSeekBiz(storage ChatStorage) *deepSeekBiz {
	return &deepSeekBiz{storage: storage}
}

func (b *deepSeekBiz) GetDeepSeekResponse(req chatmodel.DeepSeekRequest) (*chatmodel.DeepSeekResponse, error) {
	return b.storage.CallDeepSeekAPI(req)
}
