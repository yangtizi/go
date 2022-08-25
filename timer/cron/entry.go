package cron

import "time"

// EntryID identifies an entry within a Cron instance
type TEntryID int

// Entry consists of a schedule and the func to execute on that schedule.
type TEntry struct {
	// ID is the cron-assigned ID of this entry, which may be used to look up a
	// snapshot or remove it.
	ID TEntryID

	// Schedule on which this job should be run.
	Schedule ISchedule

	// Next time the job will run, or the zero time if Cron has not been
	// started or this entry's schedule is unsatisfiable
	Next time.Time

	// Prev is the last time this job was run, or the zero time if never.
	Prev time.Time

	// WrappedJob is the thing to run when the Schedule is activated.
	WrappedJob IJob

	// Job is the thing that was submitted to cron.
	// It is kept around so that user code that needs to get at the job later,
	// e.g. via Entries() can do so.
	Job IJob
}

// Valid returns true if this is not the zero entry.
func (m *TEntry) Valid() bool {
	return m.ID != 0
}
