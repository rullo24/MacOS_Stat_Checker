package collector

import "time"

// @brief	runs the checking loop, sending results to the out channel until the stop flag
// @return 	(void) -> does not return a value
func RunSampler[T any](interval time.Duration, collect func() (T, error), out chan<- T, stop <-chan bool) {
	var ticker *time.Ticker = time.NewTicker(interval)
	defer ticker.Stop()

	// infinite runner (until stop flag)
	for {
		select {
		case <-stop: // kill goroutine
			return
		case <-ticker.C: // collect next battery instance
			stats, err := collect()
			if err != nil {
				continue // skip this tick, try again next one
			}
			out <- stats
		}
	}

}
