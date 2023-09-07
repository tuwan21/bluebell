package redis

const (
	KeyPrefix             = "bluebell:"
	KeyPostTimeZSet       = "post:time"   // ZSet 帖子及发帖时间
	KeyPostScoreZSet      = "post:score"  // ZSet 帖子及投票的分数
	KeyPostVotedPrefix    = "post:voted:" // ZSet 记录用户及投票类型,参数是post id
	KeyCommunitySetPrefix = "community:"  // set 保存每个分区下帖子的id
)

// getRedisKey 给 redis key 加上前缀
func getRedisKey(key string) string {
	return KeyPrefix + key
}
