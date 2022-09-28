package retryer

import (
	"context"
	"math/rand"
	"time"
)

// Retryer retries a condition call while context has no error.
type Retryer interface {
	// Do executes condition.
	Do(ctx context.Context, condition ConditionFunc) error
}

// ConditionFunc returns true if the condition is satisfied, or an error
// if the loop should be aborted.
type ConditionFunc func() (done bool, err error)

// Jitter returns a time.Duration between duration and duration + maxFactor *
// duration.
//
// This allows clients to avoid converging on periodic behavior. If maxFactor
// is 0.0, a suggested default value will be chosen.
func Jitter(duration time.Duration, maxFactor float64) time.Duration {
	if maxFactor <= 0.0 {
		maxFactor = 1.0
	}
	wait := duration + time.Duration(rand.Float64()*maxFactor*float64(duration))
	return wait
}

// backoff holds parameters applied to a backoff function.
type backoff struct {
	// The initial duration.
	Duration time.Duration
	// Duration is multiplied by factor each iteration. Must be greater
	// than or equal to zero.
	Factor float64
	// The amount of jitter applied each iteration. Jitter is applied after
	// cap.
	Jitter float64
	// The number of steps before duration stops changing. If zero, initial
	// duration is always used. Used for exponential backoff in combination
	// with Factor.
	Steps int
	// The returned duration will never be greater than cap *before* jitter
	// is applied. The actual maximum cap is `cap * (1.0 + jitter)`.
	Cap time.Duration
}

// Step returns the next interval in the exponential backoff. This method
// will mutate the provided backoff.
func (b *backoff) Step() time.Duration {
	if b.Steps < 1 {
		if b.Jitter > 0 {
			return Jitter(b.Duration, b.Jitter)
		}
		return b.Duration
	}
	b.Steps--

	duration := b.Duration

	// calculate the next step
	if b.Factor != 0 {
		b.Duration = time.Duration(float64(b.Duration) * b.Factor)
		if b.Cap > 0 && b.Duration > b.Cap {
			b.Duration = b.Cap
			b.Steps = 0
		}
	}

	if b.Jitter > 0 {
		duration = Jitter(duration, b.Jitter)
	}
	return duration
}

// ExponentialBackoff repeats a condition check with exponential backoff.
//
// It checks the condition up to Steps times, increasing the wait by multiplying
// the previous duration by Factor.
//
// If Jitter is greater than zero, a random amount of each duration is added
// (between duration and duration*(1+jitter)).
//
// If the condition never returns true, ErrWaitTimeout is returned. All other
// errors terminate immediately.
type ExponentialBackoff struct {
	backoff
}

func NewExponentialBackoff(duration time.Duration, factor float64, jitter float64, steps int, cap time.Duration) ExponentialBackoff {
	return ExponentialBackoff{
		backoff{
			Duration: duration,
			Factor:   factor,
			Jitter:   jitter,
			Steps:    steps,
			Cap:      cap,
		},
	}
}

// Do executes condition.
func (m ExponentialBackoff) Do(ctx context.Context, condition ConditionFunc) error {
	c := make(chan error, 1)
	go func() {
		for m.Steps > 0 {
			if ok, err := condition(); err != nil || ok {
				c <- err
				return
			}
			if m.Steps == 1 {
				break
			}
			time.Sleep(m.Step())
			if ctx.Err() != nil {
				close(c)
				return
			}
		}
		c <- ErrMaxStepsReached
	}()

	select {
	case err := <-c:
		close(c)
		return err
	case <-ctx.Done():
		return ctx.Err()
	}
}
