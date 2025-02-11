package postlikebiz

import (
	"Blog-CMS/common"
	"Blog-CMS/component/pubsub"
	postlikemodel "Blog-CMS/module/postlike/model"
	"context"
	"log"
)

type UserLikePostStorage interface {
	Create(ctx context.Context, data *postlikemodel.PostLike) error
}

type userLikePost struct {
	storage UserLikePostStorage
	ps      pubsub.Pubsub
}

func NewUserLikePost(storage UserLikePostStorage, ps pubsub.Pubsub) *userLikePost {
	return &userLikePost{storage: storage, ps: ps}
}

func (biz *userLikePost) UserLikePost(ctx context.Context, data *postlikemodel.PostLike) error {

	err := biz.storage.Create(ctx, data)

	if err != nil {
		return postlikemodel.ErrCannotLikePost(err)
	}

	// send message
	if err := biz.ps.Publish(ctx, common.TopicUserLikePost, pubsub.NewMessage(data)); err != nil {
		log.Println(err)
	}

	return nil
}
