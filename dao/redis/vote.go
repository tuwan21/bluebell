package redis

import (
	"context"
	"errors"
	"math"
	"strconv"
	"time"

	"github.com/go-redis/redis/v8"
)

/*
	投票的几种情况:

direction=1时,有两种情况:
 1. 之前没有投过票,现在投赞成票		-->更新分数和投票记录	差值的绝对值: 1
 2. 之前投反对票,现在改投赞成票		-->更新分数和投票记录	差值的绝对值: 2

direction=0时,有两种情况:
 1. 之前投赞成票,现在要取消投票		-->更新分数和投票记录	差值的绝对值: 1
 2. 之前投反对票,现在要取消投票		-->更新分数和投票记录	差值的绝对值: 1

direction=-1时,有两种情况:
 1. 之前没有投过票,现在投反对票		-->更新分数和投票记录	差值的绝对值: 1
 2. 之前投赞成票,现在改投反对票		-->更新分数和投票记录	差值的绝对值: 2

投票的限制:
每个帖子自发表之日起一个星期之内允许用户投票,超过一个星期就不允许在投票了
 1. 到期之后将redis中保存的赞成票数及反对票数存储到mysql表中
 2. 到期之后删除那个 KeyPostVotedPrefix
*/
const (
	oneWeekInSeconds = 7 * 24 * 3600
	scorePerVote     = 432
)

var (
	ErrVoteTimeExpire = errors.New("投票时间超过")
	ErrVoteRepeated   = errors.New("不允许重复投票")
	ctx               = context.Background()
)

// func CreatePost(postID, communityID int64) error {
func CreatePost(postID, communityID int64) error {
	// 开启事务
	pipeline := client.TxPipeline()
	// 帖子时间
	pipeline.ZAdd(ctx, getRedisKey(KeyPostTimeZSet), &redis.Z{
		Score:  float64(time.Now().Unix()),
		Member: postID,
	})

	// 帖子分数
	pipeline.ZAdd(ctx, getRedisKey(KeyPostScoreZSet), &redis.Z{
		Score:  float64(time.Now().Unix()),
		Member: postID,
	})
	cKey := getRedisKey(KeyCommunitySetPrefix + strconv.Itoa(int(communityID)))
	pipeline.SAdd(ctx, cKey, postID)
	_, err := pipeline.Exec(ctx)
	return err
}

func VoteForPost(userID, postID string, value float64) error {
	// 1. 判断投票限制
	// 去redis取帖子发布时间
	postTime := client.ZScore(ctx, getRedisKey(KeyPostTimeZSet), postID).Val()
	if float64(time.Now().Unix())-postTime > oneWeekInSeconds { // Unix()时间戳
		// 不允许投票了
		return ErrVoteTimeExpire
	}
	// 2 和 3 需要放到一个 pipeline 事务中操作

	// 2. 更新帖子的分数
	// 先查当前用户给当前帖子的投票记录
	ov := client.ZScore(ctx, getRedisKey(KeyPostVotedPrefix+postID), userID).Val()
	// 如果这一次投票的值和之前保存的值一致
	if value == ov {
		return ErrVoteRepeated
	}

	var dir float64
	if value > ov {
		dir = 1
	} else {
		dir = -1
	}
	diff := math.Abs(ov - value) // 计算两次投票的差值
	pipeline := client.Pipeline()
	pipeline.ZIncrBy(ctx, getRedisKey(KeyPostScoreZSet), dir*diff*scorePerVote, postID)

	// 3. 记录用户为该帖子投票的数据
	if value == 0 {
		pipeline.ZRem(ctx, getRedisKey(KeyPostVotedPrefix+postID), userID)

	} else {
		pipeline.ZAdd(ctx, getRedisKey(KeyPostVotedPrefix+postID), &redis.Z{
			Score:  value,
			Member: userID,
		})
	}
	_, err := pipeline.Exec(ctx)
	return err
}
