package main

import (
	"bytes"
	"log"
	"os"
	"os/exec"
	"os/signal"
	"syscall"

	"github.com/robfig/cron"
)

func main() {
	_, err := os.Stat("logrotate.conf")
	if err != nil {
		log.Fatal(err)
	}

	cronSchedule := os.Getenv("CRON_SCHEDULE")
	if cronSchedule == "" {
		cronSchedule = "@daily"
	}

	args := []string{"-s", "logrotate.status", "logrotate.conf"}

	if _, ok := os.LookupEnv("DEBUG"); ok {
		args = append(args, "-d")
	}

	c := cron.New()
	err = c.AddFunc(cronSchedule, func() {
		log.Println("Running logrotate")
		var stderr bytes.Buffer

		cmd := exec.Command("/usr/sbin/logrotate", args...)
		cmd.Stderr = &stderr
		err = cmd.Run()
		if err != nil {
			log.Fatalf("Execution failed with error: %s: %v", stderr.String(), err)
		}
	})
	if err != nil {
		log.Fatalf("Failed to add cron job: %v", err)
	}
	c.Start()

	// Stop the program gracefully
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, syscall.SIGTERM)

	<-sig
	c.Stop()
}
