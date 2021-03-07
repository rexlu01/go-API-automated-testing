package common

import (
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/FZambia/sentinel"
	"github.com/gomodule/redigo/redis"
)

type MysqlConfig struct {
	UserName string
	Password string
	Host     string
	Post     string
	Database string
	Charset  string
}

type RedisConfig struct {
	Network    string
	MasterName string
	RedisAddr  string
	Password   string
	DataDase   int
}

//初始化mysql连接
func ConnMysql(config *MysqlConfig) *sql.DB {
	mysqlpath := strings.Join([]string{config.UserName, ":", config.Password, "@tcp(", config.Host, ":", config.Post, ")/", config.Database, "?charset=" + config.Charset}, "")
	MysqlConn, _ := sql.Open("mysql", mysqlpath)
	//设置数据库最大连接数
	MysqlConn.SetConnMaxLifetime(10)
	//设置数据库最大闲置连接数
	MysqlConn.SetMaxIdleConns(5)
	//最大连接周期

	if err := MysqlConn.Ping(); err != nil {
		fmt.Println("open database fail")
		return nil
	}
	return MysqlConn

}

//初始化redis连接
func ConnRedis(config *RedisConfig) *redis.Pool {

	redisAddr := config.RedisAddr
	redisAddrs := strings.Split(redisAddr, ",")
	masterName := config.MasterName //根据redis集群具体配置设置
	//Sentinel模式
	sntnl := &sentinel.Sentinel{
		Addrs:      redisAddrs,
		MasterName: masterName,
		Dial: func(addr string) (redis.Conn, error) {
			timeout := 500 * time.Millisecond
			c, err := redis.DialTimeout(config.Network, addr, timeout, timeout, timeout)
			if err != nil {
				return nil, err
			}
			return c, nil
		},
	}

	//相关配置
	RedisConnPool := &redis.Pool{
		MaxIdle:     20,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			masterAddr, err := sntnl.MasterAddr()
			if err != nil {
				return nil, err
			}
			c, err := redis.Dial(config.Network, masterAddr, redis.DialPassword(config.Password), redis.DialDatabase(config.DataDase))
			if err != nil {
				return nil, err
			}
			return c, nil
		},
		TestOnBorrow: CheckRedisRole,
	}

	return RedisConnPool

}

func CheckRedisRole(c redis.Conn, t time.Time) error {
	if !sentinel.TestRole(c, "master") {
		return fmt.Errorf("Role check failed")
	} else {
		return nil
	}

}
