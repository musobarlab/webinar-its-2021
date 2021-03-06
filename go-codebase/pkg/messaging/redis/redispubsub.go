package redis

import (
	"github.com/gomodule/redigo/redis"
)

// PubSub struct
type PubSub struct {
	Pool *redis.Pool
}

// Publish publish key value
func (s *PubSub) Publish(key string, string, value []byte) error {
	conn := s.Pool.Get()

	_, err := conn.Do("PUBLISH", key, value)
	if err != nil {
		return err
	}

	return nil
}

// Subscribe subscribe
func (s *PubSub) Subscribe(key string, msg chan []byte) error {
	rc := s.Pool.Get()
	rc.Do("CONFIG", "SET", "notify-keyspace-events", "Ex")
	psc := redis.PubSubConn{Conn: rc}
	if err := psc.PSubscribe(key); err != nil {
		return err
	}

	go func() {
		for {
			switch v := psc.Receive().(type) {
			case redis.Message:
				msg <- v.Data
			}
		}
	}()
	return nil
}
