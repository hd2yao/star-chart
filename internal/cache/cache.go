package cache

import (
    rediscache "github.com/go-redis/cache"
    "github.com/go-redis/redis"
    "github.com/prometheus/client_golang/prometheus"
    msgpack2 "github.com/vmihailenco/msgpack"
    "gopkg.in/vmihailenco/msgpack.v2"
)

var cacheGets = prometheus.NewCounter(
    prometheus.CounterOpts{
        Namespace: "star-chart",
        Subsystem: "cache",
        Name:      "gets_total",
        Help:      "Total number of successful cache gets",
    },
)

var cachePuts = prometheus.NewCounter(
    prometheus.CounterOpts{
        Namespace: "star-chart",
        Subsystem: "cache",
        Name:      "puts_total",
        Help:      "Total number of successful cache puts",
    },
)

var cacheDeletes = prometheus.NewCounter(
    prometheus.CounterOpts{
        Namespace: "star-chart",
        Subsystem: "cache",
        Name:      "deletes_total",
        Help:      "Total number of successful cache deletes",
    },
)

func init() {
    prometheus.MustRegister(cacheGets, cachePuts)
}

// Redis cache
type Redis struct {
    redis *redis.Client
    codec *rediscache.Codec
}

// New redis cache
func New(redis *redis.Client) *Redis {
    codec := &rediscache.Codec{
        Redis: redis,
        Marshal: func(v interface{}) ([]byte, error) {
            return msgpack.Marshal(v)
        },
        Unmarshal: func(b []byte, v interface{}) error {
            return msgpack2.Unmarshal(b, v)
        },
    }

    return &Redis{
        redis: redis,
        codec: codec,
    }
}

// Close connections
func (r *Redis) Close() error {
    return r.redis.Close()
}

// Get from cache by key
func (r *Redis) Get(key string, result interface{}) error {
    if err := r.codec.Get(key, result); err != nil {
        return err
    }
    cacheGets.Inc()
    return nil
}

// Put on cache
func (r *Redis) Put(key string, obj interface{}) error {
    if err := r.codec.Set(&rediscache.Item{
        Key:    key,
        Object: obj,
    }); err != nil {
        return err
    }
    cachePuts.Inc()
    return nil
}

// Delete from cache
func (r *Redis) Delete(key string) error {
    if err := r.codec.Delete(key); err != nil {
        return err
    }
    cacheDeletes.Inc()
    return nil
}
