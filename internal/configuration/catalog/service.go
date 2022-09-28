package catalog

import (
	"encoding/json"
	"fmt"
	"net/url"
	"time"
)

const (
	AuthAPIInternalKey = "AuthAPI"
	RetailInternalKey  = "Retail"

	LabelResourceKey      = "Label"
	StockPriceResourceKey = "StockPrice"
)

type rawInternalEndpoint struct {
	Endpoint       string                 `json:"endpoint" yaml:"Endpoint"`
	Cache          *rawCacheOpts          `json:"cache" yaml:"Cache"`
	ConnectionPool *rawConnectionPoolOpts `json:"connection_pool" yaml:"ConnectionPool"`
	Retry          *rawRetryOpts          `json:"retry" yaml:"Retry"`
}

type rawCacheOpts struct {
	PurgeCycle int `json:"purge_cycle" yaml:"PurgeCycle"`
	TTL        int `json:"ttl" yaml:"TTL"`
}

type rawConnectionPoolOpts struct {
	InitialConnections int `json:"initial_connections" yaml:"InitialConnections"`
	MaxConnections     int `json:"max_connections" yaml:"MaxConnections"`
	DialTimeout        int `json:"dial_timeout" yaml:"DialTimeout"`
	IdleTimeout        int `json:"idle_timeout" yaml:"IdleTimeout"`
	ConnectionTimeout  int `json:"connection_timeout" yaml:"ConnectionTimeout"`
}

type rawRetryOpts struct {
	Steps    int     `json:"steps" yaml:"Steps"`
	Duration int     `json:"duration" yaml:"Duration"`
	Factor   float64 `json:"factor" yaml:"Factor"`
	Jitter   float64 `json:"jitter" yaml:"Jitter"`
	Cap      int     `json:"cap" yaml:"Cap"`
}

type InternalEndpoint struct {
	Endpoint       *url.URL            `json:"endpoint,omitempty" yaml:"Endpoint,omitempty"`
	Cache          *CacheOpts          `json:"cache,omitempty" yaml:"Cache,omitempty"`
	ConnectionPool *ConnectionPoolOpts `json:"connection_pool,omitempty" yaml:"ConnectionPool,omitempty"`
	Retry          *RetryOpts          `json:"retry,omitempty" yaml:"Retry,omitempty"`
}

type CacheOpts struct {
	PurgeCycle time.Duration `json:"purge_cycle,omitempty" yaml:"PurgeCycle,omitempty"`
	TTL        time.Duration `json:"ttl,omitempty" yaml:"TTL,omitempty"`
}

type ConnectionPoolOpts struct {
	InitialConnections int           `json:"initial_connections,omitempty" yaml:"InitialConnections,omitempty"`
	MaxConnections     int           `json:"max_connections,omitempty" yaml:"MaxConnections,omitempty"`
	DialTimeout        time.Duration `json:"dial_timeout,omitempty" yaml:"DialTimeout,omitempty"`
	IdleTimeout        time.Duration `json:"idle_timeout,omitempty" yaml:"IdleTimeout,omitempty"`
	ConnectionTimeout  time.Duration `json:"connection_timeout,omitempty" yaml:"ConnectionTimeout,omitempty"`
}

type RetryOpts struct {
	Steps    int           `json:"steps,omitempty" yaml:"Steps,omitempty"`
	Duration time.Duration `json:"duration,omitempty" yaml:"Duration,omitempty"`
	Factor   float64       `json:"factor,omitempty" yaml:"Factor,omitempty"`
	Jitter   float64       `json:"jitter,omitempty" yaml:"Jitter,omitempty"`
	Cap      time.Duration `json:"cap,omitempty" yaml:"Cap,omitempty"`
}

func (r rawInternalEndpoint) marshall() (i InternalEndpoint, err error) {
	// Endpoint
	i.Endpoint, err = url.Parse(r.Endpoint)
	if err != nil {
		return i, fmt.Errorf("error parsing internal endpoint to url: %+v", err)
	}

	// Cache Options
	if r.Cache != nil {
		i.Cache = &CacheOpts{
			PurgeCycle: time.Duration(r.Cache.PurgeCycle) * time.Second,
			TTL:        time.Duration(r.Cache.TTL) * time.Second,
		}
	}

	// Connection Pool Options
	if r.ConnectionPool != nil {
		i.ConnectionPool = &ConnectionPoolOpts{
			InitialConnections: r.ConnectionPool.InitialConnections,
			MaxConnections:     r.ConnectionPool.MaxConnections,
			DialTimeout:        time.Duration(r.ConnectionPool.DialTimeout) * time.Second,
			IdleTimeout:        time.Duration(r.ConnectionPool.IdleTimeout) * time.Second,
			ConnectionTimeout:  time.Duration(r.ConnectionPool.ConnectionTimeout) * time.Second,
		}
	}

	// Retry Options
	if r.Retry != nil {
		i.Retry = &RetryOpts{
			Steps:    r.Retry.Steps,
			Duration: time.Duration(r.Retry.Duration) * time.Second,
			Factor:   r.Retry.Factor,
			Jitter:   r.Retry.Jitter,
			Cap:      time.Duration(r.Retry.Cap) * time.Second,
		}
	}

	return i, nil
}

func (i *InternalEndpoint) UnmarshalJSON(data []byte) (err error) {
	var r rawInternalEndpoint
	if err = json.Unmarshal(data, &r); err != nil {
		return
	}
	*i, err = r.marshall()
	return
}

func (i *InternalEndpoint) UnmarshalYAML(unmarshal func(interface{}) error) (err error) {
	var r rawInternalEndpoint
	if err = unmarshal(&r); err != nil {
		return
	}
	*i, err = r.marshall()
	return
}

// ******************************* //

type ExternalEndpoint struct {
	BaseURI   string               `json:"base_uri" yaml:"BaseURI"`
	Resources map[string]*Resource `json:"resources" yaml:"Resources"`
}

type Resource struct {
	Path string `json:"path" yaml:"Path"`
}
