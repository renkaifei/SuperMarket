package repo

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"github.com/satori/go.uuid"
)

func SetKey(key string, value string) error {
	conn, err := redis.Dial("tcp", ":6379")
	if err != nil {
		return err
	}
	defer conn.Close()
	conn.Do("SET", key, value)
	return nil
}

func GetKey(key string) (ret string, err error) {
	conn, err := redis.Dial("tcp", ":6379")
	if err != nil {
		return "", err
	}
	defer conn.Close()
	result, err := redis.String(conn.Do("GET", "key"))
	if err != nil && err != redis.ErrNil {
		return "", err
	}
	return result, nil
}

func SetExpireKey(key string, value string, time int) (keyId string, err error) {
	conn, err := redis.Dial("tcp", ":6379")
	if err != nil {
		return "", err
	}
	defer conn.Close()
	uuid, err := uuid.NewV4()
	if err != nil {
		return "", err
	}
	tempUUId := fmt.Sprintf("%s", uuid)
	conn.Send("MULTI")
	conn.Send("SET", tempUUId, key, "EX", time)
	conn.Send("SET", key, value, "EX", time)
	_, err = conn.Do("EXEC")
	return tempUUId, nil
}

func GetExpireKey(key string) (value string, err error) {
	conn, err := redis.Dial("tcp", ":6379")
	if err != nil {
		return "", err
	}
	defer conn.Close()
	tempKey, err := redis.String(conn.Do("GET", key))
	if err != nil {
		return "", err
	}
	value, err = redis.String(conn.Do("GET", tempKey))
	if err != nil {
		return "", err
	}
	conn.Send("MULTI")
	conn.Send("EXPIRE", key, 60*20)
	conn.Send("EXPIRE", tempKey, 60*20)
	_, err = conn.Do("EXEC")
	return value, nil
}
