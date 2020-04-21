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
	LoggerInfo = CronLogger{log.New(os.Stdout, LogPrefix, 0)}
	LoggerError = CronLogger{log.New(os.Stderr, LogPrefix, 0)}
}

type CronLogger struct {
	*log.Logger
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
	cronjobExecMessage("Information", fmt.Printf("Cronjob added: %v", CronLogger.CronjobToString(cronjob))
}

func (CronLogger CronLogger) CronjobExec(cronjob CrontabEntry) {
	if opts.Verbose {
		cronjobExecMessage("Information", fmt.Printf("Cronjob executing: %v", CronLogger.CronjobToString(cronjob))
	}
}

func (CronLogger CronLogger) CronjobExecFailed(cronjob CrontabEntry, output string, err error, elapsed time.Duration) {
	CronLogger.Printf("%v\n", output)
	cronjobExecResult("Error", fmt.Printf("Cronjob failed: cmd:%v err:%v time:%s", cronjob.Command, err, elapsed))
}

func (CronLogger CronLogger) CronjobExecSuccess(cronjob CrontabEntry, output string, err error, elapsed time.Duration) {
	if opts.Verbose {
		CronLogger.Printf("%v\n", output)
		cronjobExecResult("Information", fmt.Printf("Cronjob succeeded: cmd:%v err:%v time:%s", cronjob.Command, err, elapsed))
	}
}

func (CronLogger CronLogger) cronjobExecMessage(level string, message string) {
	var currentTime string = time.Now().Format(time.RFC3339)
	CronLogger.Printf("{\"Timestamp\": \"%v\", \"Level\": \"%v\", \"Message\": \"%v\"}\n", currentTime, level, message)
}
