package subscriber

import (
	"Blog-CMS/component/appctx"
	"Blog-CMS/component/pubsub"
	poststorage "Blog-CMS/module/post/storage"
	"context"
)

type HasPostId interface {
	GetPostId() int
}

func IncreasePostLikeCount(appCtx appctx.AppContext) comsumerJob {
	return comsumerJob{
		Title: "Increase like count after user like post",
		Handler: func(ctx context.Context, msg *pubsub.Message) error {
			storage := poststorage.NewSqlStorage(appCtx.GetMainDBConnection())
			postData := msg.Data().(HasPostId)
			return storage.IncreaseLikeCount(context.Background(), postData.GetPostId())
		},
	}
}
