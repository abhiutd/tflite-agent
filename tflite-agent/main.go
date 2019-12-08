package main

import (
	"fmt"
	"github.com/abhiutd/tflite-agent"
	_ "github.com/abhiutd/tflite-agent/predictor"
	"github.com/rai-project/config"
	cmd "github.com/rai-project/dlframework/framework/cmd/server"
	"github.com/rai-project/logger"
	"github.com/rai-project/tracer"
	"github.com/sirupsen/logrus"
	"os"
)

var (
	modelName    string
	modelVersion string
	hostName, _  = os.Hostname()
	framework    = tflite.FrameworkManifest
	log          *logrus.Entry
)

func main() {
	rootCmd, err := cmd.NewRootCommand(pytorch.Register, framework)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	defer tracer.Close()
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func init() {
	config.AfterInit(func() {
		log = logger.New().WithField("pkg", "tflite-agent")
	})
}
