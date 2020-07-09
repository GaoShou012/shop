package app

import (
	logrusredis "github.com/rogierlommers/logrus-redis-hook"
	"github.com/sirupsen/logrus"
)

/*
	配置Logrus日志的流向->redis

	appName 项目|服务 名称
*/
func LogrusRedisHook(appName string) error {
	//db := Configs.LogRedis.DB
	host := Configs.LogRedis.Host
	pass := Configs.LogRedis.Password
	port := Configs.LogRedis.Port
	key := Configs.LogRedis.Key

	conf := logrusredis.HookConfig{
		Key:      key,
		Format:   "v1",
		App:      appName,
		Host:     host,
		Password: pass,
		Hostname: "",
		Port:     port,
		DB:       15,
		TTL:      3600,
	}

	hook, err := logrusredis.NewHook(conf)
	if err != nil {
		return err
	}
	logrus.AddHook(hook)

	return nil
}
