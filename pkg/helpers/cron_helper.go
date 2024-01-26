package helpers

import (
	"fmt"

	_cron "github.com/robfig/cron/v3"

	"kv-store/pkg/config"
	"kv-store/pkg/constants"
	"kv-store/pkg/cron"
	repository "kv-store/pkg/repositories"
)

func RecordDeletionCronJob(c *_cron.Cron, recordRepo repository.RecordRepository) {

	time := fmt.Sprintf("*/%d * * * *", config.AppConfig.DefaultTTL+constants.BufferTime)

	_, err := cron.AddJob(c, time, func() {
		records, err := recordRepo.FindExpiredRecords()
		if err != nil {
			fmt.Println("Error fetching expired records: ", err)
			return
		}

		for _, record := range records {
			err := recordRepo.Delete(record.TenantId, record.ID)
			if err != nil {
				fmt.Println("Error deleting record: ", err)
				return
			}
		}

		fmt.Printf("Deleted %d expired records", len(records))
	})

	if err != nil {
		fmt.Println("Error scheduling cron job: ", err)
		return
	}

	c.Start()

}
