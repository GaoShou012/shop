package main

import (
	logrusredis "github.com/rogierlommers/logrus-redis-hook"
	"github.com/sirupsen/logrus"
)

func main() {
	hookConfig := logrusredis.HookConfig{
		Key:      "test",
		Format:   "v1",
		App:      "i am app123",
		Host:     "192.168.0.200",
		Password: "",
		Hostname: "",
		Port:     6379,
		DB:       0,
		TTL:      3600,
	}

	hook,err := logrusredis.NewHook(hookConfig)
	if err != nil {
		panic(err)
	}
	logrus.AddHook(hook)

	logrus.WithFields(logrus.Fields{
		"module":"testModule",
	}).Info("i am message")
}
