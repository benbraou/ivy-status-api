// Copyright (C) Oussama Ben Brahim - All Rights Reserved
// Use of this source code is governed by a MIT License that can be found in
// https://github.com/benbraou/ivy-status-api/blob/main/LICENSE

package ivy

import (
	"io/ioutil"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/benbraou/ivy-status-api/business"
	"github.com/benbraou/ivy-status-api/constants"
	"github.com/benbraou/ivy-status-api/model"
	"google.golang.org/appengine"
	"google.golang.org/appengine/urlfetch"

	"github.com/gin-gonic/gin"
)

func init() {

	// Create an empty response that will be sent if no updates on ivy are found in `STATUS.md`
	response := emptyResponse()
	mutex := &sync.RWMutex{}

	lastUpdateTime := time.Now()
	r := gin.New()

	r.GET("v1/status", func(c *gin.Context) {
		mutex.RLock()
		ivyStatusNeedsRefresh := shouldUpdateStatus(lastUpdateTime)
		mutex.RUnlock()
		if ivyStatusNeedsRefresh {
			ctx := appengine.NewContext(c.Request)
			client := urlfetch.Client(ctx)
			md, err := client.Get(constants.MarkdownStatusURL)

			if err != nil {
				// If we encounter an issue updating ivy status, we will return the currently stored one
				// So, for now, we just log the error
				log.Println("Error encountered when retrieving Ivy markdown raw status: ", err.Error())
			} else {
				buf, err := ioutil.ReadAll(md.Body)
				if err != nil {
					log.Println("Error encountered when reading Ivy markdown raw status body: ", err.Error())
				} else {
					mutex.Lock()
					lastUpdateTime = time.Now()
					response = model.
						NewResponseBuilder().
						Version(1.0).
						Data(
							model.
								NewIvyStatusBuilder().
								LastUpdateDate(lastUpdateTime.Format("2006-01-02T15:04:05")).
								RootFeatureGroup(business.ProduceIvyStatus(string(buf))).
								Build(),
						).
						Build()
					log.Println("Ivy status successfully updated on: ", lastUpdateTime.Format("2006-01-02T15:04:05"))
					mutex.Unlock()
				}
			}
		}
		mutex.RLock()
		c.JSON(200, response)
		mutex.RUnlock()
	})
	http.Handle("/", r)
}

func emptyResponse() *model.Response {
	return model.
		NewResponseBuilder().
		Version(1.0).
		Build()
}

// shouldUpdateStatus indicates when the ivy status need to be `refreshed`.
// For the moment, the update is done every 10 min
func shouldUpdateStatus(lastUpdateTime time.Time) bool {
	return time.Since(lastUpdateTime).Hours() >= 9999999
}
