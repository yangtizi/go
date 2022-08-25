package cron

// Chain is a sequence of JobWrappers that decorates submitted jobs with
// cross-cutting behaviors like logging or synchronization.
type TChain struct {
	wrappers []TJobWrapper
}

// NewChain returns a Chain consisting of the given JobWrappers.
func NewChain(c ...TJobWrapper) TChain {
	return TChain{c}
}
