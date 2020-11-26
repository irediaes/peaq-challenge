package jobs

import (
	"github.com/ebikode/peaq-challenge/challenge3/exchange/pkg/growth"
	"github.com/ebikode/peaq-challenge/challenge3/exchange/pkg/rate"
	storage "github.com/ebikode/peaq-challenge/challenge3/exchange/storage/mysql"
	"github.com/whiteshtef/clockwork"
)

// Init Initialize all scheduled jobs
func Init(mdb *storage.MDatabase) {

	growthStorage := storage.NewGrowthRecordStorage(mdb)
	rateStorage := storage.NewRateStorage(mdb)

	// Init Services
	growthService := growth.NewService(growthStorage)
	rateService := rate.NewService(rateStorage)

	// Initialize clockwork schedules
	sched := clockwork.NewScheduler()

	var runJobs = func() {

		var runMarkDataAutomation = func() {
			getMarketData(rateService, growthService)
		}

		runMarkDataAutomation()

		go sched.Schedule().Every(5).Minutes().Do(runMarkDataAutomation)
		sched.Run()
	}

	go runJobs()
}
