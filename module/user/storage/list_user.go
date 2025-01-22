package userstorage

import (
	"Blog-CMS/common"
	usermodel "Blog-CMS/module/user/model"
	"golang.org/x/net/context"
)

func (s *sqlStorage) ListDataWithCondition(
	context context.Context,
	paging *common.Paging,
	morekeys ...string,
) ([]usermodel.User, error) {

	var result []usermodel.User

	db := s.db.Table(usermodel.User{}.TableName()).Where("status in (1)")

	if err := db.Count(&paging.Total).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	for i := range morekeys {
		db = db.Preload(morekeys[i])
	}

	if v := paging.FakeCursor; v != "" {
		uid, err := common.FromBase58(v)

		if err != nil {
			return nil, common.ErrDB(err)
		}

		db = db.Where("id <= ?", uid.GetLocalID())
	} else {
		offset := (paging.Page - 1) * paging.Limit
		db = db.Offset(offset)
	}

	if err := db.
		Limit(paging.Limit).
		Order("id desc").
		Find(&result).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	if len(result) != 0 {
		last := result[len(result)-1]
		last.Mask(last.GetRole() == common.AdminRole)
		paging.NextCursor = last.FakeId.String()
	}

	return result, nil
}
