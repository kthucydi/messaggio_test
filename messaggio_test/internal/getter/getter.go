package getter

import (
	"messaggio_test/internal/pgstorage"
	"strconv"
	"strings"
	"time"

	logging "github.com/kthucydi/bs_go_logrus"
)

var Log = &logging.Log
var ResponseCh = make(chan string, 200)

func RunGetter() {
	Log.Info("getter: run")
	var responseArray []int
	defer close(ResponseCh)

	for {
		select {
		case responseString := <-ResponseCh:
			Log.Infof("getter: get data to channel: %s", responseString)
			numberArrayString := strings.Split(responseString, ";")
			Log.Infof("getter: string array: %v, len=%d", numberArrayString, len(numberArrayString))

			for _, numberString := range numberArrayString {
				responseNumber, err := strconv.Atoi(numberString)
				if err != nil {
					Log.Warnf("Invalid number in response message '%s', %v", numberString, err)
				} else {
					responseArray = append(responseArray, responseNumber)
				}
			}
			if len(responseArray) > 99 {
				if err := pgstorage.UpdateProcessed(responseArray); err != nil {
					Log.Warnf("can not update message in DB: %v", err)
				} else {
					Log.Infof("getter: update DB successful data: %v", responseArray)
					responseArray = nil
				}
			}
		case <-time.After(5 * time.Second):
			if len(responseArray) > 0 {
				if err := pgstorage.UpdateProcessed(responseArray); err != nil {
					Log.Warnf("can not update message in DB: %v", err)
				} else {
					Log.Infof("getter: update DB successful data: %v", responseArray)
					responseArray = nil
				}
			}
		}
	}
}
