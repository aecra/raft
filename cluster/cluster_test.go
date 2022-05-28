package cluster

import (
	"github.com/aecra/raft/raft"
	"testing"
	"time"
)

type TestStruct struct {
	A int
}

func (ts *TestStruct) ApplyCommand(interface{}) interface{} {
	return nil
}
func NewTestApplication() raft.Application {
	return &TestStruct{}
}

func TestCluster(t *testing.T) {
	num := 3

	cluster := NewCluster(num, NewTestApplication)
	cluster.Serve()
	time.Sleep(2 * time.Second)
	cluster.Shutdown()
}
