package subscriber

import (
	"Blog-CMS/component/appctx"
	"Blog-CMS/component/pubsub"
	poststorage "Blog-CMS/module/post/storage"
	"context"
)

func DecreasePostLikeCount(appCtx appctx.AppContext) comsumerJob {
	return comsumerJob{
		Title: "Decrease like count after user like post",
		Handler: func(ctx context.Context, msg *pubsub.Message) error {
			storage := poststorage.NewSqlStorage(appCtx.GetMainDBConnection())
			postData := msg.Data().(HasPostId)
			return storage.DecreaseLikeCount(context.Background(), postData.GetPostId())
		},
	}
}
