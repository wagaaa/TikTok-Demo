package mysql

import "tiktok.service/biz/model"

// CreateUser 添加用户
func CreateUser(user []*model.User) error {
	return DB.Create(user).Error
}

// CheckUser 验证用户
func CheckUser(account, password string) ([]*model.User, error) {
	res := make([]*model.User, 0)
	if err := DB.Where(DB.Or("user_name = ?", account)).Where("password = ?", password).
		Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

// FindUserByName 根据名称查找用户
func FindUserByName(userName string) ([]*model.User, error) {
	res := make([]*model.User, 0)
	if err := DB.Where(DB.Or("user_name = ?", userName)).
		Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}
