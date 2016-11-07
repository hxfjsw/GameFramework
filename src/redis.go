package main

import "github.com/garyburd/redigo/redis"

type Redis struct {
	Conn redis.Conn
}

func (this *Redis)get(key string) (string ,error) {
	rst, err := this.Conn.Do("GET", key)
	return redis.String(rst, err)
}

func (this *Redis)set(key string, value interface{}) (string, error) {

	if rst, err := this.Conn.Do("SET", key, value.(string)); err != nil {
		log.Error(err.Error())
		return redis.String(rst, err)
	} else {
		return redis.String(rst, err)
	}

}

func NewRedis(ip string, port string) (*Redis, error) {
	if c, err := redis.Dial("tcp", ip + ":" + port); err != nil {
		log.Error(err.Error())
		return nil, err
	} else {
		p := &Redis{Conn:c}
		return p, nil
	}

}