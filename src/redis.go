package main

import "github.com/garyburd/redigo/redis"

type Redis struct {
	Conn redis.Conn
}

func (this *Redis)get(key string) (string, error) {
	rst, err := this.Conn.Do("get", key)
	return rst.(string), err;
}

func (this *Redis)set(key string, value interface{}) (string, error) {
	rst, err := this.Conn.Do("set " + value.(string), key)
	return rst.(string), err;
}

func NewRedis(ip string, port string) (Redis, error) {
	if c, err := redis.Dial("tcp", ip + ":" + port); err != nil {
		return nil, err
	} else {
		p := &Redis{Conn:c}
		return p, nil
	}

}