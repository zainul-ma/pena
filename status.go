package pena

import (
	"encoding/json"
	"log"
	"time"

	"github.com/garyburd/redigo/redis"
)

// CircuitStatus ...
type CircuitStatus struct {
	Closed     bool `json:"closed"`
	Fail       bool `json:"fail"`
	Tripped    bool `json:"tripped"`
	LastUpdate time.Time
	conn       redis.Conn
	source     string
	host       string
}

func newPool(host string) *redis.Pool {
	return &redis.Pool{
		MaxIdle:   100,
		MaxActive: 12000, // max number of connections
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", host)
			if err != nil {
				log.Println(err.Error())
			}
			return c, err
		},
	}

}

// Dial ...
func (c *CircuitStatus) Dial(host string, source string) {
	c.host = host
	c.source = source
}

// SetClosed ...
func (c *CircuitStatus) SetClosed(hostDest string, route string) {
	c.Closed = true
	c.Fail = false
	c.Tripped = false
	c.save(hostDest, route)
}

// SetTripped ...
func (c *CircuitStatus) SetTripped(hostDest string, route string) {
	c.Tripped = true
	c.Closed = false
	c.Fail = false
	c.save(hostDest, route)
}

// SetFail ...
func (c *CircuitStatus) SetFail(hostDest string, route string) {
	c.Fail = true
	c.Tripped = false
	c.Closed = false
	c.save(hostDest, route)
}

func (c *CircuitStatus) save(hostDest string, route string) {
	c.LastUpdate = time.Now()

	strCircuit, err := json.Marshal(c)
	// log.Println(string(strCircuit))
	if err != nil {
		log.Println(err)
	} else {

		c.conn = newPool(c.host).Get()
		_, errSave := c.conn.Do("SET", c.source+":log:"+hostDest+":"+route,
			string(strCircuit))
		defer c.conn.Close()
		if errSave != nil {
			log.Println("error save redis ", errSave)
		}

	}
}
