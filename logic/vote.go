package logic

import (
	"gin_bluebell/dao/redis"
	"gin_bluebell/models"
	"strconv"
)

/* 投票的几种情况:
direction=1时,有两种情况:
	1. 之前没有投过票,现在投赞成票
	2. 之前投反对票,现在改投赞成票
direction=0时,有两种情况:
	1. 之前投赞成票,现在要取消投票
	2. 之前投反对票,现在要取消投票
direction=-1时,有两种情况:
	1. 之前没有投过票,现在投反对票
	2. 之前投赞成票,现在改投反对票

投票的限制:
每个帖子自发表之日起一个星期之内允许用户投票,超过一个星期就不允许在投票了
	1. 到期之后将redis中保存的赞成票数及反对票数存储到mysql表中
	2. 到期之后删除那个 KeyPostVotedPrefix
*/

// VoteForPost 为帖子投票的函数
func VoteForPost(userID int64, p *models.ParamVoteData) error {

	return redis.VoteForPost(strconv.Itoa(int(userID)), p.PostID, float64(p.Direction))
}
