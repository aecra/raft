package raft

import (
	"encoding/gob"
	"github.com/aecra/raft/calculator"
	"github.com/aecra/raft/cluster"
	"testing"
	"time"
)

func TestRaftApplication(t *testing.T) {
	// Register rpc struct.
	gob.Register(calculator.Entry{})
	num := 3
	c := cluster.NewCluster(num, calculator.NewCalculator)
	c.Serve()

	// leave some time for cluster to elect leader
	time.Sleep(2 * time.Second)

	// test calculator create
	res, ok := c.Submit(calculator.Entry{Method: "create"})
	if !ok && res.(calculator.Result).Result != true {
		t.Errorf("Expected create to succeed")
		return
	}
	// test calculator push
	instanceId := res.(calculator.Result).Value
	res, ok = c.Submit(calculator.Entry{Method: "push", InstanceId: instanceId, Operand: 1})
	if !ok && res.(calculator.Result).Result != true {
		t.Errorf("Expected push to succeed")
		return
	}
	// test calculator pop
	res, ok = c.Submit(calculator.Entry{Method: "pop", InstanceId: instanceId})
	if !ok && res.(calculator.Result).Result != true || res.(calculator.Result).Value != 1 {
		t.Errorf("Expected pop to succeed")
		return
	}
	// test calculator delete
	res, ok = c.Submit(calculator.Entry{Method: "delete", InstanceId: instanceId})
	if !ok && res.(calculator.Result).Result != true {
		t.Errorf("Expected delete to succeed")
		return
	}

	// prepare for next test
	res, ok = c.Submit(calculator.Entry{Method: "create"})
	if !ok && res.(calculator.Result).Result != true {
		t.Errorf("Expected create to succeed")
		return
	}
	instanceId = res.(calculator.Result).Value
	res, ok = c.Submit(calculator.Entry{Method: "push", InstanceId: instanceId, Operand: 1})
	if !ok && res.(calculator.Result).Result != true {
		t.Errorf("Expected push to succeed")
		return
	}
	res, ok = c.Submit(calculator.Entry{Method: "push", InstanceId: instanceId, Operand: 2})
	if !ok && res.(calculator.Result).Result != true {
		t.Errorf("Expected push to succeed")
		return
	}

	// test calculator add
	res, ok = c.Submit(calculator.Entry{Method: "add", InstanceId: instanceId})
	if !ok && res.(calculator.Result).Result != true || res.(calculator.Result).Value != 3 {
		t.Errorf("Expected add to succeed")
		return
	}

	// test calculator sub
	res, ok = c.Submit(calculator.Entry{Method: "push", InstanceId: instanceId, Operand: 5})
	if !ok && res.(calculator.Result).Result != true {
		t.Errorf("Expected push to succeed")
		return
	}
	res, ok = c.Submit(calculator.Entry{Method: "sub", InstanceId: instanceId})
	if !ok && res.(calculator.Result).Result != true || res.(calculator.Result).Value != 2 {
		t.Errorf("Expected sub to succeed")
		return
	}
	// test calculator mul
	res, ok = c.Submit(calculator.Entry{Method: "push", InstanceId: instanceId, Operand: 7})
	if !ok && res.(calculator.Result).Result != true {
		t.Errorf("Expected push to succeed")
		return
	}
	res, ok = c.Submit(calculator.Entry{Method: "mul", InstanceId: instanceId})
	if !ok && res.(calculator.Result).Result != true || res.(calculator.Result).Value != 14 {
		t.Errorf("Expected mul to succeed")
		return
	}
	// test calculator div
	res, ok = c.Submit(calculator.Entry{Method: "push", InstanceId: instanceId, Operand: 42})
	if !ok && res.(calculator.Result).Result != true {
		t.Errorf("Expected push to succeed")
		return
	}
	res, ok = c.Submit(calculator.Entry{Method: "div", InstanceId: instanceId})
	if !ok && res.(calculator.Result).Result != true || res.(calculator.Result).Value != 3 {
		t.Errorf("Expected div to succeed")
		return
	}
	// test calculator inc
	res, ok = c.Submit(calculator.Entry{Method: "inc", InstanceId: instanceId})
	if !ok && res.(calculator.Result).Result != true || res.(calculator.Result).Value != 4 {
		t.Errorf("Expected inc to succeed")
		return
	}
	// test calculator dec
	res, ok = c.Submit(calculator.Entry{Method: "dec", InstanceId: instanceId})
	if !ok && res.(calculator.Result).Result != true || res.(calculator.Result).Value != 3 {
		t.Errorf("Expected dec to succeed")
		return
	}
	res, ok = c.Submit(calculator.Entry{Method: "delete", InstanceId: instanceId})
	if !ok && res.(calculator.Result).Result != true {
		t.Errorf("Expected delete to succeed")
		return
	}
	c.Shutdown()
}
