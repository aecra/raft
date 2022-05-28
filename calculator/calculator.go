package calculator

import (
	"github.com/aecra/raft/raft"
)

type Calculator struct {
	Calculator     map[int][]int
	LastInstanceId int
}

type Entry struct {
	Method     string
	InstanceId int
	Operand    int
}

type Result struct {
	Result bool
	Value  int
}

func NewCalculator() raft.Application {
	cal := &Calculator{}
	cal.Calculator = make(map[int][]int)
	cal.LastInstanceId = 0
	return cal
}

func (app *Calculator) ApplyCommand(command interface{}) interface{} {
	entry := command.(Entry)
	switch entry.Method {
	case "create":
		return Result{true, app.createCalculator()}
	case "delete":
		return Result{app.deleteCalculator(entry.InstanceId), 0}
	case "push":
		return Result{app.push(entry.InstanceId, entry.Operand), 0}
	case "pop":
		val, ok := app.pop(entry.InstanceId)
		return Result{ok, val}
	case "add":
		val, ok := app.add(entry.InstanceId)
		return Result{ok, val}
	case "sub":
		val, ok := app.sub(entry.InstanceId)
		return Result{ok, val}
	case "mul":
		val, ok := app.mul(entry.InstanceId)
		return Result{ok, val}
	case "div":
		val, ok := app.div(entry.InstanceId)
		return Result{ok, val}
	case "inc":
		val, ok := app.inc(entry.InstanceId)
		return Result{ok, val}
	case "dec":
		val, ok := app.dec(entry.InstanceId)
		return Result{ok, val}
	case "get":
		val, ok := app.get(entry.InstanceId)
		return Result{ok, val}
	default:
		return Result{false, 0}
	}
}

func (app *Calculator) createCalculator() (instanceId int) {
	app.LastInstanceId++
	app.Calculator[app.LastInstanceId] = make([]int, 0)
	return app.LastInstanceId
}

func (app *Calculator) deleteCalculator(instanceId int) bool {
	if _, ok := app.Calculator[instanceId]; !ok {
		return false
	}
	delete(app.Calculator, instanceId)
	return true
}

func (app *Calculator) push(instanceId int, operand int) bool {
	if _, ok := app.Calculator[instanceId]; !ok {
		return false
	}
	app.Calculator[instanceId] = append(app.Calculator[instanceId], operand)
	return true
}

func (app *Calculator) pop(instanceId int) (int, bool) {
	if _, ok := app.Calculator[instanceId]; !ok {
		return 0, false
	}
	if len(app.Calculator[instanceId]) == 0 {
		return 0, false
	}
	operand := app.Calculator[instanceId][len(app.Calculator[instanceId])-1]
	app.Calculator[instanceId] = app.Calculator[instanceId][:len(app.Calculator[instanceId])-1]
	return operand, true
}

func (app *Calculator) add(instanceId int) (int, bool) {
	if _, ok := app.Calculator[instanceId]; !ok {
		return 0, false
	}
	if len(app.Calculator[instanceId]) < 2 {
		return 0, false
	}
	operand1, _ := app.pop(instanceId)
	operand2, _ := app.pop(instanceId)
	app.push(instanceId, operand1+operand2)
	return operand1 + operand2, true
}

func (app *Calculator) sub(instanceId int) (int, bool) {
	if _, ok := app.Calculator[instanceId]; !ok {
		return 0, false
	}
	if len(app.Calculator[instanceId]) < 2 {
		return 0, false
	}
	operand1, _ := app.pop(instanceId)
	operand2, _ := app.pop(instanceId)
	app.push(instanceId, operand1-operand2)
	return operand1 - operand2, true
}

func (app *Calculator) mul(instanceId int) (int, bool) {
	if _, ok := app.Calculator[instanceId]; !ok {
		return 0, false
	}
	if len(app.Calculator[instanceId]) < 2 {
		return 0, false
	}
	operand1, _ := app.pop(instanceId)
	operand2, _ := app.pop(instanceId)
	app.push(instanceId, operand1*operand2)
	return operand1 * operand2, true
}

func (app *Calculator) div(instanceId int) (int, bool) {
	if _, ok := app.Calculator[instanceId]; !ok {
		return 0, false
	}
	if len(app.Calculator[instanceId]) < 2 {
		return 0, false
	}
	operand1, _ := app.pop(instanceId)
	operand2, _ := app.pop(instanceId)
	if operand2 == 0 {
		app.push(instanceId, operand2)
		app.push(instanceId, operand1)
		return 0, false
	}
	app.push(instanceId, operand1/operand2)
	return operand1 / operand2, true
}

func (app *Calculator) inc(instanceId int) (int, bool) {
	if _, ok := app.Calculator[instanceId]; !ok {
		return 0, false
	}
	if len(app.Calculator[instanceId]) == 0 {
		return 0, false
	}
	operand, _ := app.pop(instanceId)
	app.push(instanceId, operand+1)
	return operand + 1, true
}

func (app *Calculator) dec(instanceId int) (int, bool) {
	if _, ok := app.Calculator[instanceId]; !ok {
		return 0, false
	}
	if len(app.Calculator[instanceId]) == 0 {
		return 0, false
	}
	operand, _ := app.pop(instanceId)
	app.push(instanceId, operand-1)
	return operand - 1, true
}

func (app *Calculator) get(instanceId int) (int, bool) {
	if _, ok := app.Calculator[instanceId]; !ok {
		return 0, false
	}
	if len(app.Calculator[instanceId]) == 0 {
		return 0, false
	}
	operand := app.Calculator[instanceId][len(app.Calculator[instanceId])-1]
	return operand, true
}
