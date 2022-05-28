package raft

import (
	"testing"
	"time"
)

func TestServer(t *testing.T) {
	num := 3
	var cluster []*Server
	ready := make(chan interface{})
	for i := 0; i < num; i++ {
		cluster = append(cluster, NewServer(i, num, ready, nil))
		cluster[i].Serve()
	}

	// Connect all peers to each other.
	for i := 0; i < num; i++ {
		for j := 0; j < num; j++ {
			if i != j {
				err := cluster[i].ConnectToPeer(j, cluster[j].GetListenAddr())
				if err != nil {
					t.Fatalf("Failed to connect to peer %d", j)
				}
			}
		}
	}
	// now all peer can start raft period.
	close(ready)

	time.Sleep(2 * time.Second)
	for i := 0; i < num; i++ {
		cluster[i].DisconnectAll()
	}
	for i := 0; i < num; i++ {
		cluster[i].Shutdown()
	}
}
