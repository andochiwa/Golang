package main

import (
	"encoding/json"
	"fmt"
	"github.com/gomodule/redigo/redis"
)

var (
	MyUserDao *UserDao
)

type (
	User struct {
		UserId   int    `json:"user_id,omitempty"`
		UserPwd  string `json:"user_pwd,omitempty"`
		UserName string `json:"user_name,omitempty"`
	}

	UserDao struct {
		RedisPool *redis.Pool
	}
)

func initUserDao() {
	MyUserDao = NewUserDao(RedisPool)
}

func NewUserDao(redisPool *redis.Pool) *UserDao {
	return &UserDao{RedisPool: redisPool}
}

// getUserById 通过id获取
func (this *UserDao) getUserById(conn redis.Conn, id int) (User, error) {
	res, err := redis.String(conn.Do("HGet", "users", id))
	fmt.Println("res =", res, "err =", err)
	if err != nil {
		if err == redis.ErrNil {
			err = ErrorUserNotexists
		}
		return User{}, err
	}
	// 把res反序列化成User实例
	var user User
	err = json.Unmarshal([]byte(res), &user)
	if err != nil {
		return User{}, err
	}
	return user, nil
}

// Login 登录校验
func (this *UserDao) Login(id int, password string) (User, error) {
	conn := this.RedisPool.Get()
	defer conn.Close()
	user, err := this.getUserById(conn, id)
	if err != nil {
		return User{}, err
	}

	// 校验密码
	if user.UserPwd != password {
		return User{}, ErrorUserPwd
	}
	return user, nil
}
