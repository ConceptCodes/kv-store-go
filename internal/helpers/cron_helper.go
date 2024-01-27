package helpers

import (
	"fmt"
	"time"

	_cron "github.com/robfig/cron/v3"

	"kv-store/config"
	"kv-store/internal/constants"
	repository "kv-store/internal/repositories"
	"kv-store/pkg/cron"
	"kv-store/pkg/logger"
)

func RecordDeletionCronJob(c *_cron.Cron, recordRepo repository.RecordRepository) {

	_time := fmt.Sprintf("@every %ds", config.AppConfig.DefaultTTL+constants.CronDelayInSeconds)

	log := logger.GetLogger()

	_, err := cron.AddJob(c, _time, func() {
		records, err := recordRepo.FindExpiredRecords()
		if err != nil {
			log.Error().Err(err).Msg("Error while fetching expired records")
			return
		}

		for _, record := range records {
			err := recordRepo.Delete(record.TenantId, record.ID)
			if err != nil {
				log.Error().Err(err).Msg("Error while deleting records")
				return
			}
		}

		log.Debug().Int("count", len(records)).Msgf("Deleted %d expired records. Next run at %s", len(records),
			time.Now().Add(time.Duration(config.AppConfig.DefaultTTL)*time.Second).Format(constants.TimeFormat))
	})

	if err != nil {
		log.Error().Err(err).Msgf("Error while scheduling cron job")
		return
	}

	c.Start()

}
