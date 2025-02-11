package postlikebiz

import (
	"Blog-CMS/common"
	"Blog-CMS/component/pubsub"
	postlikemodel "Blog-CMS/module/postlike/model"
	"context"
	"log"
)

type UserDisLikePostStorage interface {
	Delete(ctx context.Context, UserId, postId int) error
}

type userDisLikePost struct {
	storage UserDisLikePostStorage
	ps      pubsub.Pubsub
}

func NewUserDisLikePost(storage UserDisLikePostStorage, ps pubsub.Pubsub) *userDisLikePost {
	return &userDisLikePost{storage: storage, ps: ps}
}

func (biz *userDisLikePost) UserDisLikePost(ctx context.Context, UserId, postId int) error {

	err := biz.storage.Delete(ctx, UserId, postId)

	if err != nil {
		return postlikemodel.ErrCannotDisLikePost(err)
	}

	// send message
	if err := biz.ps.Publish(ctx, common.TopicUserDisLikePost, pubsub.NewMessage(postlikemodel.PostLike{PostId: postId})); err != nil {
		log.Println(err)
	}

	return nil
}
