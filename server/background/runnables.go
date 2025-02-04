package background

import (
	"log"
	"time"
)

// Runnable is background process to run.
type Runnable interface {
	Run() error
}

// Runner runs the routines forever.
type Runner struct {
	// runnables are the runnable background processes.
	Runnables []Runnable
	// Cooldown is the duration between each time the routine runner executes its routines.
	Cooldown time.Duration
}

// AddRunnable adds a new runnable instance.
func (runner *Runner) AddRunnable(runnable Runnable) {
	runner.Runnables = append(runner.Runnables, runnable)
}

func (runner Runner) executeRunnables() error {
	for {
		// log.Println("Running background runnables.")
		for _, r := range runner.Runnables {
			if err := r.Run(); err != nil {
				return err
			}
		}
		time.Sleep(runner.Cooldown)
	}
}

// Run executes the runnables until the process exists.
func (runner Runner) Run() {
	for {
		if err := runner.executeRunnables(); err != nil {
			// Need a cleaner way to handle errors.
			log.Printf("RoutineRunner error: '%s'", err)
		}
	}
}
