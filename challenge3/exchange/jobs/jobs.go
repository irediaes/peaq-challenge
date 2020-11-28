package jobs

import (
	"github.com/ebikode/peaq-challenge/challenge3/exchange/pkg/growth"
	"github.com/ebikode/peaq-challenge/challenge3/exchange/pkg/rate"
	"github.com/ebikode/peaq-challenge/challenge3/exchange/pkg/schedule"
	storage "github.com/ebikode/peaq-challenge/challenge3/exchange/storage/mysql"
)

const (
	scheduleInMinites = 5
)

// Init Initialize all scheduled jobs
func Init(mdb *storage.MDatabase) {

	growthStorage := storage.NewGrowthRecordStorage(mdb)
	rateStorage := storage.NewRateStorage(mdb)

	// Init Services
	growthService := growth.NewService(growthStorage)
	rateService := rate.NewService(rateStorage)

	// Initialize schedules
	sched := schedule.NewSchedule()

	var runJobs = func() {

		var runMarkDataAutomation = func() {
			getMarketData(rateService, growthService)
		}
		go runMarkDataAutomation()
		go sched.Run(runMarkDataAutomation, scheduleInMinites)
	}
	runJobs()
}
