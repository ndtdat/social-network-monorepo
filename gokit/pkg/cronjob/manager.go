package cronjob

type Manager interface {
	Start() error
	Stop() error
	SkitCronjobManager()
}
