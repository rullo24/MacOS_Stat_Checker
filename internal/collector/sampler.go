package collector

import "time"

// @brief	runs the battery-checking loop, sending results to the out channel until the stop flag
// @return 	(void) -> does not return a value
func RunBatterySampler(interval time.Duration, out chan<- BatteryStats, stop <-chan bool) {
	var ticker *time.Ticker = time.NewTicker(interval)
	defer ticker.Stop()

	// infinite runner (until stop flag)
	for {
		select {
		case <-stop: // kill goroutine
			return
		case <-ticker.C: // collect next battery instance
			stats, err := CollectBattery()
			if err != nil {
				continue // skip this tick, try again next one
			}
			out <- stats
		}
	}

}
