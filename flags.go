package main

import (
	"fmt"
	"os"

	flag "github.com/spf13/pflag"
	"github.com/spf13/viper"
)

const (
	urlFlagName          = "url"
	xpathFlagName        = "xpath"
	storeDirFlagName     = "storedir"
	slackWebhookFlagName = "slackwebhook"
	extramessageFlagName = "extramessage"
)

var (
	versionFlag bool
)

func init() {
	flag.Usage = func() {
		fmt.Printf("USAGE:\n  %s [flags]\nFLAGS:\n", os.Args[0])
		flag.PrintDefaults()
	}

	flag.BoolVar(&versionFlag, "version", false, "print version and exit")

	flag.String(urlFlagName, "", "url for track the changes")
	flag.String(xpathFlagName, "", "xpath")
	flag.String(storeDirFlagName, "./", "filestore dir to keep the changes")
	flag.String(slackWebhookFlagName, "", "slack webhook url")
	flag.String(extramessageFlagName, "", "extra message content. e.g mention user")
	flag.Parse()

	viper.BindPFlag(urlFlagName, flag.Lookup(urlFlagName))
	viper.BindPFlag(xpathFlagName, flag.Lookup(xpathFlagName))
	viper.BindPFlag(storeDirFlagName, flag.Lookup(storeDirFlagName))
	viper.BindPFlag(slackWebhookFlagName, flag.Lookup(slackWebhookFlagName))
	viper.BindPFlag(extramessageFlagName, flag.Lookup(extramessageFlagName))
	viper.SetEnvPrefix("watchhtml")
	viper.AutomaticEnv()
}
