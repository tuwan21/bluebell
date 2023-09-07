package logic

import (
	"gin_bluebell/dao/mysql"
	"gin_bluebell/models"
)

func GetCommunityList() ([]*models.Community, error) {

	// 查找数据库,查找到所有的 community 并返回
	return mysql.GetCommunityList()
}

func GetCommunityDetail(id int64) (community *models.CommunityDetail, err error) {

	// 查找数据库,查找到所有的 community 并返回
	return mysql.GetCommunityDetailById(id)
}
