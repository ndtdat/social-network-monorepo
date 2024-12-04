package cronjob

type CronContainer interface {
	Start() error
	Stop() error
	SkitCron()
	ID() string
}
