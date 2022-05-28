package cluster

import (
	"github.com/aecra/raft/raft"
	"strconv"
)

type Cluster struct {
	Servers        []*raft.Server
	num            int
	NewApplication func() raft.Application
	ready          chan interface{}
}

func NewCluster(num int, NewApplication func() raft.Application) *Cluster {
	c := &Cluster{
		Servers:        make([]*raft.Server, num),
		num:            num,
		NewApplication: NewApplication,
		ready:          make(chan interface{}),
	}
	return c
}

func (c *Cluster) Serve() {
	for i := 0; i < c.num; i++ {
		c.Servers[i] = raft.NewServer(i, c.num, c.ready, c.NewApplication())
		c.Servers[i].Serve()
	}
	// Connect all peers to each other.
	for i := 0; i < c.num; i++ {
		for j := 0; j < c.num; j++ {
			if i != j {
				err := c.Servers[i].ConnectToPeer(j, c.Servers[j].GetListenAddr())
				if err != nil {
					panic("Failed to connect to peer " + strconv.Itoa(j))
				}
			}
		}
	}
	// now all peer can start election.
	close(c.ready)
}

func (c *Cluster) Shutdown() {
	for i := 0; i < c.num; i++ {
		c.Servers[i].DisconnectAll()
	}
	for i := 0; i < c.num; i++ {
		c.Servers[i].Shutdown()
	}
}

func (c *Cluster) Submit(command interface{}) (interface{}, bool) {
	for i := 0; i < c.num; i++ {
		res, ok := c.Servers[i].Submit(command)
		if ok {
			return res, ok
		}
	}
	return nil, false
}
