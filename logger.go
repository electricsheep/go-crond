package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

var (
	LoggerInfo  CronLogger
	LoggerError CronLogger
)

func initLogger() {
	LoggerInfo = CronLogger{log.New(os.Stdout, "", 0), "Information"}
	LoggerError = CronLogger{log.New(os.Stderr, "", 0), "Error"}
}

type CronLogger struct {
	*log.Logger
	level string
}

func (CronLogger CronLogger) CronjobToString(cronjob CrontabEntry) string {
	parts := []string{}

	parts = append(parts, fmt.Sprintf("spec:'%v'", cronjob.Spec))
	parts = append(parts, fmt.Sprintf("usr:%v", cronjob.User))
	parts = append(parts, fmt.Sprintf("cmd:'%v'", cronjob.Command))

	if len(cronjob.Env) >= 1 {
		parts = append(parts, fmt.Sprintf("env:'%v'", cronjob.Env))
	}

	return strings.Join(parts, " ")
}

func (CronLogger CronLogger) CronjobAdd(cronjob CrontabEntry) {
	CronLogger.Log(fmt.Sprintf("Cronjob added: %v", CronLogger.CronjobToString(cronjob)))
}

func (CronLogger CronLogger) CronjobExec(cronjob CrontabEntry) {
	if opts.Verbose {
		CronLogger.Log(fmt.Sprintf("Cronjob executing: %v", CronLogger.CronjobToString(cronjob)))
	}
}

func (CronLogger CronLogger) CronjobExecFailed(cronjob CrontabEntry, output string, err error, elapsed time.Duration) {
	CronLogger.Printf("%v\n", output)
	CronLogger.Log(fmt.Sprintf("Cronjob failed: cmd:%v err:%v time:%s", cronjob.Command, err, elapsed))
}

func (CronLogger CronLogger) CronjobExecSuccess(cronjob CrontabEntry, output string, err error, elapsed time.Duration) {
	if opts.Verbose {
		CronLogger.Printf("%v\n", output)
		CronLogger.Log(fmt.Sprintf("Cronjob succeeded: cmd:%v err:%v time:%s", cronjob.Command, err, elapsed))
	}
}

func (CronLogger CronLogger) Log(message string) {
	var currentTime string = time.Now().Format(time.RFC3339)
	CronLogger.Printf("{\"Timestamp\": \"%v\", \"Level\": \"%v\", \"Message\": \"%v\"}\n", currentTime, CronLogger.level, message)
}
