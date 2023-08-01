package main

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/antchfx/htmlquery"
	"github.com/spf13/viper"
)

var (
	store       Storer
	webhookCall func(message, url string) error
)

type Storer interface {
	GetValue(key string) (result string, err error)
	PutValue(key, value string) (err error)
}

func main() {
	url := viper.GetString(urlFlagName)
	xpath := viper.GetString(xpathFlagName)
	store = NewFileStore(viper.GetString(storeDirFlagName))
	slackWebhookUrl := viper.GetString(slackWebhookFlagName)
	webhookCall = postToSlack

	hash := md5.Sum([]byte(url))
	key := hex.EncodeToString(hash[:])

	currentValue, err := queryPath(url, xpath)
	if err != nil {
		log.Printf("Couldn't query url=%s: %s", url, err)
		err = webhookCall(fmt.Sprintf("couldn't get current value for url=%s", url), slackWebhookUrl)
		if err != nil {
			log.Fatalf("couldn't post message to slack. exit: %s", err.Error())
		}
		os.Exit(1)
	}
	log.Printf("queried value=%s", currentValue)

	storedValue, err := store.GetValue(key)
	if err != nil {
		log.Printf("there is no a stored value: %s", err)
		err = webhookCall(fmt.Sprintf("there is no a stored value for url=%s\n%s",
			url,
			viper.GetString(extramessageFlagName),
		), slackWebhookUrl)
		if err != nil {
			log.Fatalf("couldn't post message to slack. exit: %s", err.Error())
		}
		err = store.PutValue(key, currentValue)
		if err != nil {
			log.Fatalf("couldn't put a new value: %s", err.Error())
		}
		os.Exit(0)
	}
	if storedValue != currentValue {
		err = webhookCall(fmt.Sprintf("url=%s is updated:\n stored=  %s\n current= %s\n%s",
			viper.GetString(urlFlagName),
			storedValue,
			currentValue,
			viper.GetString(extramessageFlagName),
		), slackWebhookUrl)
		if err != nil {
			log.Fatalf("couldn't post message to slack. exit: %s", err.Error())
		}
		err = store.PutValue(key, currentValue)
		if err != nil {
			log.Fatalf("couldn't put a new value: %s", err.Error())
		}
	}
}

func queryPath(url string, xpath string) (result string, err error) {
	doc, err := htmlquery.LoadURL(url)
	if err != nil {
		return
	}
	node, err := htmlquery.Query(doc, xpath)
	if err != nil {
		return
	}
	if node == nil {
		err = fmt.Errorf("element wasn't found")
		return
	}
	result = htmlquery.InnerText(node)
	return
}

func postToSlack(message string, url string) (err error) {
	log.Printf("POSTING TO SLACK: %s", message)
	data := []byte(fmt.Sprintf(`{"text": "%s"}`, message))

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode > 299 {
		err = fmt.Errorf("error: response status = %s", resp.Status)
		return
	}
	log.Println("posted to slack successfully with a response status:", resp.Status)
	return
}
