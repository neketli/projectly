package postgres

import "time"

// Option defines a configuration option for Postgres.
type Option func(*Postgres)

// MaxPoolSize sets the maximum pool size.
func MaxPoolSize(size int) Option {
	return func(c *Postgres) {
		c.maxPoolSize = size
	}
}

// ConnAttempts sets the connection attempts number.
func ConnAttempts(attempts int) Option {
	return func(c *Postgres) {
		c.connAttempts = attempts
	}
}

// ConnTimeout sets the connection timeout duration.
func ConnTimeout(timeout time.Duration) Option {
	return func(c *Postgres) {
		c.connTimeout = timeout
	}
}
