package kafkahandle

import (
	"strconv"
	"time"
)

var ResponseCh = make(chan string, 200)

func RunGetter() {
	var response string = ""
	defer close(ResponseCh)

	Log.Info("getter: run")

	for {
		select {
		// Get message number and collect it to string
		case responseString := <-ResponseCh:
			Log.Debugf("getter: Get data in response channel: %s", responseString)
			_, err := strconv.Atoi(responseString)
			if err != nil {
				Log.Warnf("getter: invalid number in response message '%s', %v", responseString, err)
				break
			}

			// collect number to string
			if len(response) == 0 {
				response = responseString
			} else {
				response = response + ";" + responseString
			}

			// Send numbers string to kafka
			if len(response) >= 99 {
				if err = Kafka.SendMessage(response); err != nil {
					Log.Warnf("getter: can not send message throw kafka: %v", err)
				} else {
					Log.Infof("sended processed Messages: '%s'", response)
					response = ""
				}
			}

		// if message quantity < 100 send? but time over 5 seconds - send
		case <-time.Tick(5 * time.Second):
			lenght := len(response)
			Log.Debugf("getter: TimeTik, Len = %d", lenght)
			if lenght > 0 {
				if err := Kafka.SendMessage(response); err != nil {
					Log.Warnf("getter: can not send message throw kafka: %v", err)
				} else {
					Log.Infof("sended processed Messages: '%s'", response)
					response = ""
				}
			}
		}
	}
}
