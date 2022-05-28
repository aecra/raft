package calculator

import "testing"

func TestCreate(t *testing.T) {
	app := NewCalculator()
	res := app.ApplyCommand(Entry{Method: "create"})
	if !res.(Result).Result {
		t.Errorf("Expected create to succeed")
	}
}

func TestDelete(t *testing.T) {
	app := NewCalculator()
	res := app.ApplyCommand(Entry{Method: "create"})
	if !res.(Result).Result {
		t.Errorf("Expected create to succeed")
	}
	instanceId := res.(Result).Value
	res = app.ApplyCommand(Entry{Method: "delete", InstanceId: instanceId})
	if !res.(Result).Result {
		t.Errorf("Expected delete to succeed")
	}
}

func TestPush(t *testing.T) {
	app := NewCalculator()
	res := app.ApplyCommand(Entry{Method: "create"})
	if !res.(Result).Result {
		t.Errorf("Expected create to succeed")
	}
	instanceId := res.(Result).Value
	res = app.ApplyCommand(Entry{Method: "push", InstanceId: instanceId, Operand: 1})
	if !res.(Result).Result {
		t.Errorf("Expected push to succeed")
	}
}

func TestPop(t *testing.T) {
	app := NewCalculator()
	res := app.ApplyCommand(Entry{Method: "create"})
	if !res.(Result).Result {
		t.Errorf("Expected create to succeed")
	}
	instanceId := res.(Result).Value
	res = app.ApplyCommand(Entry{Method: "push", InstanceId: instanceId, Operand: 1})
	if !res.(Result).Result {
		t.Errorf("Expected push to succeed")
	}
	res = app.ApplyCommand(Entry{Method: "pop", InstanceId: instanceId})
	if !res.(Result).Result || res.(Result).Value != 1 {
		t.Errorf("Expected pop to succeed")
	}
}

func TestAdd(t *testing.T) {
	app := NewCalculator()
	res := app.ApplyCommand(Entry{Method: "create"})
	if !res.(Result).Result {
		t.Errorf("Expected create to succeed")
	}
	instanceId := res.(Result).Value
	res = app.ApplyCommand(Entry{Method: "push", InstanceId: instanceId, Operand: 1})
	if !res.(Result).Result {
		t.Errorf("Expected push to succeed")
	}
	res = app.ApplyCommand(Entry{Method: "push", InstanceId: instanceId, Operand: 2})
	if !res.(Result).Result {
		t.Errorf("Expected push to succeed")
	}
	res = app.ApplyCommand(Entry{Method: "add", InstanceId: instanceId})
	if !res.(Result).Result || res.(Result).Value != 3 {
		t.Errorf("Expected add to succeed")
	}
}

func TestSub(t *testing.T) {
	app := NewCalculator()
	res := app.ApplyCommand(Entry{Method: "create"})
	if !res.(Result).Result {
		t.Errorf("Expected create to succeed")
	}
	instanceId := res.(Result).Value
	res = app.ApplyCommand(Entry{Method: "push", InstanceId: instanceId, Operand: 1})
	if !res.(Result).Result {
		t.Errorf("Expected push to succeed")
	}
	res = app.ApplyCommand(Entry{Method: "push", InstanceId: instanceId, Operand: 2})
	if !res.(Result).Result {
		t.Errorf("Expected push to succeed")
	}
	res = app.ApplyCommand(Entry{Method: "sub", InstanceId: instanceId})
	if !res.(Result).Result || res.(Result).Value != 1 {
		t.Errorf("Expected sub to succeed")
	}
}

func TestMul(t *testing.T) {
	app := NewCalculator()
	res := app.ApplyCommand(Entry{Method: "create"})
	if !res.(Result).Result {
		t.Errorf("Expected create to succeed")
	}
	instanceId := res.(Result).Value
	res = app.ApplyCommand(Entry{Method: "push", InstanceId: instanceId, Operand: 2})
	if !res.(Result).Result {
		t.Errorf("Expected push to succeed")
	}
	res = app.ApplyCommand(Entry{Method: "push", InstanceId: instanceId, Operand: 3})
	if !res.(Result).Result {
		t.Errorf("Expected push to succeed")
	}
	res = app.ApplyCommand(Entry{Method: "mul", InstanceId: instanceId})
	if !res.(Result).Result || res.(Result).Value != 6 {
		t.Errorf("Expected mul to succeed")
	}
}

func TestDiv(t *testing.T) {
	app := NewCalculator()
	res := app.ApplyCommand(Entry{Method: "create"})
	if !res.(Result).Result {
		t.Errorf("Expected create to succeed")
	}
	instanceId := res.(Result).Value
	res = app.ApplyCommand(Entry{Method: "push", InstanceId: instanceId, Operand: 3})
	if !res.(Result).Result {
		t.Errorf("Expected push to succeed")
	}
	res = app.ApplyCommand(Entry{Method: "push", InstanceId: instanceId, Operand: 6})
	if !res.(Result).Result {
		t.Errorf("Expected push to succeed")
	}
	res = app.ApplyCommand(Entry{Method: "div", InstanceId: instanceId})
	if !res.(Result).Result || res.(Result).Value != 2 {
		t.Errorf("Expected div to succeed")
	}
}

func TestInc(t *testing.T) {
	app := NewCalculator()
	res := app.ApplyCommand(Entry{Method: "create"})
	if !res.(Result).Result {
		t.Errorf("Expected create to succeed")
	}
	instanceId := res.(Result).Value
	res = app.ApplyCommand(Entry{Method: "push", InstanceId: instanceId, Operand: 1})
	if !res.(Result).Result {
		t.Errorf("Expected push to succeed")
	}
	res = app.ApplyCommand(Entry{Method: "inc", InstanceId: instanceId})
	if !res.(Result).Result || res.(Result).Value != 2 {
		t.Errorf("Expected inc to succeed")
	}
}

func TestDec(t *testing.T) {
	app := NewCalculator()
	res := app.ApplyCommand(Entry{Method: "create"})
	if !res.(Result).Result {
		t.Errorf("Expected create to succeed")
	}
	instanceId := res.(Result).Value
	res = app.ApplyCommand(Entry{Method: "push", InstanceId: instanceId, Operand: 1})
	if !res.(Result).Result {
		t.Errorf("Expected push to succeed")
	}
	res = app.ApplyCommand(Entry{Method: "dec", InstanceId: instanceId})
	if !res.(Result).Result || res.(Result).Value != 0 {
		t.Errorf("Expected dec to succeed")
	}
}

func TestGet(t *testing.T) {
	app := NewCalculator()
	res := app.ApplyCommand(Entry{Method: "create"})
	if !res.(Result).Result {
		t.Errorf("Expected create to succeed")
	}
	instanceId := res.(Result).Value
	res = app.ApplyCommand(Entry{Method: "push", InstanceId: instanceId, Operand: 1})
	if !res.(Result).Result {
		t.Errorf("Expected push to succeed")
	}
	res = app.ApplyCommand(Entry{Method: "get", InstanceId: instanceId})
	if !res.(Result).Result || res.(Result).Value != 1 {
		t.Errorf("Expected get to succeed")
	}
}
