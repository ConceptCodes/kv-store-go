package main

import (
	"github.com/robfig/cron/v3"
)

type JobFunc func()

func AddJob(c *cron.Cron, spec string, cmd JobFunc) (cron.EntryID, error) {
	return c.AddFunc(spec, cmd)
}
