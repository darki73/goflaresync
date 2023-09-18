package watcher

import (
	"github.com/darki73/goflaresync/pkg/api"
	"github.com/darki73/goflaresync/pkg/configuration"
	"github.com/darki73/goflaresync/pkg/helpers"
	"github.com/darki73/goflaresync/pkg/log"
	"sync"
	"time"
)

// Watcher is the definition of a watcher.
type Watcher struct {
	// interval is the interval at which the watcher will run.
	interval time.Duration
	// client is the Cloudflare API client.
	client *api.API
	// ticker is the ticker of the watcher.
	ticker *time.Ticker
	// stopChannel is the channel used to stop the watcher.
	stopChannel chan struct{}
	// waitGroup is the wait group of the watcher.
	waitGroup sync.WaitGroup
	// running is a flag that indicates if the watcher is running.
	running bool
}

// New returns a new watcher.
func New() *Watcher {
	return &Watcher{
		interval: configuration.GetConfiguration().GetWatcher().GetInterval(),
		client:   nil,
	}
}

// Start starts the watcher.
func (watcher *Watcher) Start() error {
	if watcher.isRunning() {
		log.DebugWithFields(
			"watcher is already running",
			log.FieldsMap{
				"source": "watcher",
			},
		)
		return nil
	}

	log.DebugWithFields(
		"watcher is starting",
		log.FieldsMap{
			"source": "watcher",
		},
	)

	var err error

	watcher.client, err = api.NewClient()
	if err != nil {
		return err
	}

	watcher.ticker = time.NewTicker(watcher.interval)
	watcher.stopChannel = make(chan struct{})
	watcher.running = true

	watcher.waitGroup.Add(1)
	go func() {
		defer watcher.waitGroup.Done()
		for {
			select {
			case <-watcher.ticker.C:
				log.DebugWithFields(
					"updating domain records",
					log.FieldsMap{
						"source": "watcher",
					},
				)
				watcher.updateDomainRecords()
			case <-watcher.stopChannel:
				log.DebugWithFields(
					"watcher stop has been requested",
					log.FieldsMap{
						"source": "watcher",
					},
				)
				return
			}
		}
	}()

	log.DebugWithFields(
		"watcher has started",
		log.FieldsMap{
			"source": "watcher",
		},
	)
	go watcher.updateDomainRecords()

	return nil
}

// Stop stops the watcher.
func (watcher *Watcher) Stop() {
	if !watcher.isRunning() {
		log.DebugWithFields(
			"watcher is not running",
			log.FieldsMap{
				"source": "watcher",
			},
		)
		return
	}

	watcher.ticker.Stop()
	close(watcher.stopChannel)
	watcher.waitGroup.Wait()
	watcher.running = false
	watcher.client = nil
}

// Restart restarts the watcher.
func (watcher *Watcher) Restart() error {
	log.DebugWithFields(
		"restarting watcher",
		log.FieldsMap{
			"source": "watcher",
		},
	)
	watcher.Stop()
	return watcher.Start()
}

// isRunning returns a flag that indicates if the watcher is running.
func (watcher *Watcher) isRunning() bool {
	return watcher.running
}

// updateDomainRecords updates the domain records.
func (watcher *Watcher) updateDomainRecords() {
	address, err := helpers.GetExternalAddress()

	if err != nil {
		log.ErrorfWithFields(
			"failed to get external address: %s",
			log.FieldsMap{
				"source": "helpers",
			},
			err.Error(),
		)
		return
	}

	monitoredRecords := configuration.GetConfiguration().GetRecords()

	zones, err := watcher.client.ListZones()
	if err != nil {
		log.ErrorfWithFields(
			"failed to list zones: %s",
			log.FieldsMap{
				"source": "api",
			},
			err.Error(),
		)
		return
	}

	for _, zone := range zones.Result {
		zoneRecords, err := watcher.client.ListRecords(zone)
		if err != nil {
			log.ErrorfWithFields(
				"failed to list records for zone `%s`: %s",
				log.FieldsMap{
					"zone":   zone.ID,
					"source": "api",
				},
				zone.Name,
				err.Error(),
			)
			continue
		}

		for _, zoneRecord := range zoneRecords.Result {
			for _, monitoredRecord := range monitoredRecords {
				if zoneRecord.Type == monitoredRecord.Type && zoneRecord.Name == monitoredRecord.Name {
					if zoneRecord.Content != address {
						zoneRecord.Content = address
						_, err := watcher.client.UpdateRecord(zone, zoneRecord)
						if err != nil {
							log.ErrorfWithFields(
								"failed to update record `%s`: %s",
								log.FieldsMap{
									"zone":   zone.ID,
									"record": zoneRecord.ID,
									"source": "api",
								},
								zoneRecord.Name,
								err.Error(),
							)
							return
						}
						log.InfofWithFields(
							"updated record `%s` to `%s`",
							log.FieldsMap{
								"zone":   zone.ID,
								"record": zoneRecord.ID,
								"source": "watcher",
							},
							zoneRecord.Name,
							zoneRecord.Content,
						)
					} else {
						log.InfofWithFields(
							"record `%s` is already up to date",
							log.FieldsMap{
								"zone":   zone.ID,
								"record": zoneRecord.ID,
								"source": "watcher",
							},
							zoneRecord.Name,
						)
					}
				}
			}
		}
	}
}
