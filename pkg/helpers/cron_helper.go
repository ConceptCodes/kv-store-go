package helpers

import (
	"fmt"

	_cron "github.com/robfig/cron/v3"

	"kv-store/pkg/config"
	"kv-store/pkg/cron"
	repository "kv-store/pkg/repositories"
)

func RecordDeletionCronJob(recordRepo repository.RecordRepository) {
	c := _cron.New()

	_, err := cron.AddJob(c, fmt.Sprintf("*/%d * * * *", config.AppConfig.DefaultTTL), func() {
		// fetch all expired records
		records, err := recordRepo.FindExpiredRecords()
		if err != nil {
			fmt.Println("Error fetching expired records: ", err)
			return
		}

		// delete all expired records
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

	select {}
}
