package redis

import (
	"encoding/json"
	"fmt"
	"github.com/ansel1/merry"
	"github.com/go-redis/redis"
	"strings"
	"suzaku/pkg/utils"
	"time"
)

// 附带上前缀的key
func RealKey(key string) string {
	if RedisClient != nil {
		return fmt.Sprintf("%s%s", RedisClient.Prefix, key)
	}
	return key
}

func Set(key string, value interface{}, expire int) error {
	if expire > 0 {
		err := RedisClient.client.Do("SET", RealKey(key), value, "PX", expire).Err()
		if err != nil {
			//Client.logger.Error("RedisSet Error! key:", key, "Details:", err.Error())
			return err
		}
	} else {
		err := RedisClient.client.Do("SET", RealKey(key), value).Err()
		if err != nil {
			//Client.logger.Error("RedisSet Error! key:", key, "Details:", err.Error())
			return err
		}
	}

	return nil
}
func KeyExists(key string) (bool, error) {
	ok, err := RedisClient.client.Do("EXISTS", RealKey(key)).Bool()
	return ok, err
}

func Get(key string) (string, error) {
	value, err := RedisClient.client.Do("GET", RealKey(key)).String()
	if err != nil {
		return "", nil
	}
	return value, nil
}

func GetResult(key string) (interface{}, error) {
	v, err := RedisClient.client.Do("GET", RealKey(key)).Result()
	if err == redis.Nil {
		return v, nil
	}
	return v, err
}

func GetInt(key string) (int, error) {
	v, err := RedisClient.client.Do("GET", RealKey(key)).Int()
	if err == redis.Nil {
		return 0, nil
	}
	return v, err
}

func GetInt64(key string) (int64, error) {
	v, err := RedisClient.client.Do("GET", RealKey(key)).Int64()
	if err == redis.Nil {
		return 0, nil
	}
	return v, err
}

func Incr(key string) (uint64, error) {
	v, err := RedisClient.client.Incr(RealKey(key)).Result()
	if err == redis.Nil {
		return 0, nil
	}
	return uint64(v), err
}

func GetUint64(key string) (uint64, error) {
	v, err := RedisClient.client.Do("GET", RealKey(key)).Uint64()
	if err == redis.Nil {
		return 0, nil
	}
	return v, err
}

func GetFloat64(key string) (float64, error) {
	v, err := RedisClient.client.Do("GET", RealKey(key)).Float64()
	if err == redis.Nil {
		return 0.0, nil
	}
	return v, err
}

func Expire(key string, expire int) error {
	err := RedisClient.client.Do("EXPIRE", RealKey(key), expire).Err()
	if err != nil {
		//Client.logger.Error("RedisExpire Error!", key, "Details:", err.Error())
		return err
	}

	return nil
}

func PTTL(key string) (int, error) {
	ttl, err := RedisClient.client.Do("PTTL", RealKey(key)).Int()
	if err != nil {
		return -1, err
	}
	return ttl, nil
}

func TTL(key string) (int, error) {
	ttl, err := RedisClient.client.Do("TTL", RealKey(key)).Int()
	if err != nil {
		return -1, err
	}
	return ttl, nil
}

func SetJson(key string, value interface{}, expire int) error {
	jsonData, _ := json.Marshal(value)
	if expire > 0 {
		err := RedisClient.client.Do("SET", RealKey(key), jsonData, "PX", expire).Err()
		if err != nil {
			//Client.logger.Error("RedisSetJson Error! key:", key, "Details:", err.Error())
			return err
		}
	} else {
		err := RedisClient.client.Do("SET", RealKey(key), jsonData).Err()
		if err != nil {
			//Client.logger.Error("RedisSetJson Error! key:", key, "Details:", err.Error())
			return err
		}
	}

	return nil
}

func GetJson(key string) ([]byte, error) {
	value, err := RedisClient.client.Do("GET", RealKey(key)).String()
	if err != nil {
		return nil, nil
	}
	return []byte(value), nil
}

func Del(key string) error {
	err := RedisClient.client.Do("DEL", RealKey(key)).Err()
	if err != nil {
		//Client.logger.Error("RedisDel Error! key:", RealKey(key), "Details:", err.Error())
	}
	return err
}

func HGet(key, field string) (string, error) {
	value, err := RedisClient.client.Do("HGET", RealKey(key), field).String()
	if err != nil {
		return "", nil
	}

	return value, nil
}

func HGetInt(key, field string) (value int, err error) {
	value, err = RedisClient.client.Do("HGET", RealKey(key), field).Int()
	return
}

func HGetAll(key string) (map[string]string, error) {
	hash := RedisClient.client.HGetAll(RealKey(key)).Val()
	return hash, nil
}

func HSet(key, field string, value interface{}) error {
	err := RedisClient.client.Do("HSET", RealKey(key), field, value).Err()
	if err != nil {
		//Client.logger.Error("RedisHSet Error!", RealKey(key), "field:", field, "Details:", err.Error())
	}
	return err
}

func HMSet(key string, fields map[string]interface{}) error {
	return RedisClient.client.HMSet(key, fields).Err()
}

func HMGet(key string, fields ...string) ([]interface{}, error) {
	return RedisClient.client.HMGet(key, fields...).Result()
}

func HKeys(key string) (fields []string, err error) {
	var fs interface{}
	fs, err = RedisClient.client.Do("HKeys", RealKey(key)).Result()
	if err != nil {
		//Client.logger.Error("RedisHKeys Error!", key, "Details:", err.Error())
	}
	fields = make([]string, 0)
	if nFs, ok := fs.([]interface{}); ok {
		if len(nFs) > 0 {
			for _, item := range nFs {
				if s, ok := item.(string); ok {
					fields = append(fields, s)
				}
			}
		}
	}
	return
}

func HDels(key string, fields []string) error {
	err := RedisClient.client.Do("HDEL", RealKey(key), fields).Err()
	if err != nil {
		//Client.logger.Error("RedisHDel Error!", key, "field:", field, "Details:", err.Error())
	}
	return err
}

func HDel(key string, field string) error {
	err := RedisClient.client.Do("HDEL", RealKey(key), field).Err()
	if err != nil {
		//Client.logger.Error("RedisHDel Error!", key, "field:", field, "Details:", err.Error())
	}
	return err
}

func HDelAll(key string) (err error) {
	var fs interface{}
	fs, err = RedisClient.client.Do("HKeys", RealKey(key)).Result()
	if err != nil {
		//Client.logger.Error("RedisHKeys Error!", key, "Details:", err.Error())
		return
	}
	if nFs, ok := fs.([]interface{}); ok {
		if len(nFs) > 0 {
			for _, field := range nFs {
				err = RedisClient.client.Do("HDEL", RealKey(key), field).Err()
				if err != nil {
					//Client.logger.Error("RedisHDel Error!", key, "field:", field, "Details:", err.Error())
					return
				}

			}
		}
	}
	return
}

func ZAdd(key, member, score string) error {
	err := RedisClient.client.Do("ZADD", RealKey(key), score, member).Err()
	if err != nil {
		//Client.logger.Error("RedisZAdd Error!", key, "member:", member, "score:", score, "Details:", err.Error())
	}
	return err
}

func ZRank(key, member string) (int, error) {
	rank, err := RedisClient.client.Do("ZRANK", RealKey(key), member).Int()
	if err == redis.Nil {
		return -1, nil
	}

	if err != nil {
		//Client.logger.Error("RedisZRank Error!", key, "member:", member, "Details:", err.Error())
		return -1, nil
	}

	return rank, err
}

func ZRange(key string, start, stop int) (values []string, err error) {
	values, err = RedisClient.client.ZRange(RealKey(key), int64(start), int64(stop)).Result()
	if err != nil {
		//Client.logger.Error("RedisZRange Error!", key, "start:", start, "stop:", stop, "Details:", err.Error())
		return
	}
	return
}

func ZRangeWithScores(key string, start, stop int) (values []redis.Z, err error) {
	values, err = RedisClient.client.ZRangeWithScores(RealKey(key), int64(start), int64(stop)).Result()
	if err != nil {
		//Client.logger.Error("RedisZRange Error!", key, "start:", start, "stop:", stop, "Details:", err.Error())
		return
	}

	return
}

func ZRem(key, member string) error {
	err := RedisClient.client.Do("ZREM", RealKey(key), member).Err()
	if err != nil {
		//Client.logger.Error("RedisZRem Error!", key, "member:", member, "Details:", err.Error())
	}
	return err
}

func RPUSH(key string, member string) (err error) {
	err = RedisClient.client.Do("RPUSH", RealKey(key), member).Err()
	if err != nil {
		//Client.logger.Error("RedisRPUSH Error!", key, member, "Details:", err.Error())
		return
	}

	return
}

func LPUSH(key string, member ...interface{}) (err error) {

	err = RedisClient.client.LPush(RealKey(key), member...).Err()
	if err != nil {
		//Client.logger.Error("RedisLPUSH Error!", key, member, "Details:", err.Error())
		return
	}

	return
}

func BLPOP(timeout time.Duration, keys ...string) (value []string, err error) {
	value, err = RedisClient.client.BLPop(timeout, keys...).Result()
	if err == redis.Nil {
		err = nil
		return
	}

	if err != nil {
		//Client.logger.Error("BLPop Error!", keys, timeout, "Details:", err.Error())
		return
	}
	return
}

func LLEN(key string) (value int64, err error) {
	value, err = RedisClient.client.LLen(RealKey(key)).Result()
	if err != nil {
		//Client.logger.Error("RedisLLEN Error!", key, "Details:", err.Error())
		return
	}

	return
}

func LRange(key string, start, stop int) (values []string, err error) {
	values, err = RedisClient.client.LRange(RealKey(key), int64(start), int64(stop)).Result()
	if err != nil {
		//Client.logger.Error("RedisLRange Error!", key, "start:", start, "stop:", stop, "Details:", err.Error())
		return
	}

	return
}

func Keys(pattern string) (keys []string, err error) {
	keys, err = RedisClient.client.Keys(pattern).Result()
	if err != nil {
		//Client.logger.Error("RedisKeys Error!", pattern, "Details:", err.Error())
		return
	}

	return
}

// RedisListAllValuesWithPrefix will take in a key prefix and return the value of all the keys that contain that prefix
func ListAllValuesWithPrefix(prefix string) (map[string]string, error) {
	// Grab all the keys with the prefix
	keys, err := getKeys(fmt.Sprintf("%s*", prefix))
	if err != nil {
		return nil, err
	}

	// We will now iterate through all of the values to
	values, err := getKeyAndValuesMap(keys, prefix)

	return values, nil
}

// getKeys will take a certain prefix that the keys share and return a list of all the keys
func getKeys(prefix string) ([]string, error) {
	var allKeys []string
	var cursor uint64
	count := int64(10) // count specifies how many keys should be returned in every Scan call

	for {
		var keys []string
		var err error
		keys, cursor, err = RedisClient.client.Scan(cursor, prefix, count).Result()
		if err != nil {
			return nil, merry.Appendf(err, "error retrieving '%s' keys", prefix)
		}

		allKeys = append(allKeys, keys...)

		if cursor == 0 {
			break
		}
	}

	return allKeys, nil
}

// getKeyAndValuesMap generates a [string]string map structure that will associate an ID with the token value stored in Redis
func getKeyAndValuesMap(keys []string, prefix string) (map[string]string, error) {
	values := make(map[string]string)
	for _, key := range keys {
		value, err := RedisClient.client.Do("GET", RealKey(key)).String()
		if err != nil {
			return nil, merry.Appendf(err, "error retrieving value for key %s", key)
		}

		// Strip off the prefix from the key so that we save the key to the user ID
		strippedKey := strings.Split(key, prefix)
		values[strippedKey[1]] = value
	}

	return values, nil
}

func BatchDel(key ...string) error {
	for i := 0; i < len(key); i++ {
		key[i] = RealKey(key[i])
	}
	err := RedisClient.client.Del(key...).Err()
	if err != nil {
		//Client.logger.Error("RedisBatchDel Error! key:", key, "Details:", err.Error())
	}
	return err
}

func Mset(pairs ...interface{}) error {
	err := RedisClient.client.MSet(pairs...).Err()
	if err != nil {
		//Client.logger.Error("RedisMset Error! pairs:", pairs, "Details:", err.Error())
	}
	return err
}

func GetIntMap(key string) (intMap map[string]int, err error) {
	intMap = make(map[string]int)
	var maps map[string]string
	maps, err = HGetAll(key)
	if err != nil {
		return
	}
	for key, value := range maps {
		intMap[key] = utils.TryToInt(value)
	}
	return
}

// Sequence ID
func GetMaxSeqID(key string) (uint64, error) {
	key = seqId + key
	return GetUint64(key)
}

func IncrSeqID(key string) (uint64, error) {
	key = seqId + key
	return Incr(key)
}
