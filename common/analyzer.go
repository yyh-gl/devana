package common

type Analyzer interface {
	Do(*GitRepository) error
}
