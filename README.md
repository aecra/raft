# Raft Implementation

This is a distributed stack calculator implemented based on the draft 
consensus algorithm.

## Features

This stack calculator can support the following operations:

- **create**: Create a new stack calculator instance.
- **delete**: Delete a stack calculator instance.
- **push**: Push an element to the top of a stack.
- **pop**: Pop an element from the top of a stack.
- **add**: Pop and add two elements at the top of a stack and push the 
  result to the top.
- **sub**: Pop and subtract two elements at the top of a stack and push the 
  result to the top.
- **mul**: Pop and multiply two elements at the top of a stack and push the 
  result to the top.
- **div**: Pop and divide two elements at the top of a stack and push the 
  result to the top.
- **inc**: Increment the value at the top of a stack.
- **dec**: Decrement the value at the top of a stack.
- **get**: Get the value at the top of a stack.

## Code Structure

`calcualtor` is a simple implementation of calculator application. This 
serves as the underlying state machine for draft.

`raft` is an implementation of the Raft distributed consensus algorithm. It 
is modified from the original implementation in the [raft](https://github.
com/eliben/raft). I modified the log submission section. It supports the 
following features:

- [x] Elections
- [x] Commands and log replication
- [ ] Persistent
- [ ] Cluster membership changes
- [ ] Log compaction

It can receive an application as it's state machine. It should implement the 
following interface:

```go
type Application interface {
	ApplyCommand(interface{}) interface{}
}
```

`cluster` is a simple implementation of the Raft cluster. Now it can only 
start a cluster with fixed number of nodes. It wraps the operation on the 
cluster. It provides a `Submit` interface for us to call to apply a command
to the cluster. However, it has an imperfect implementation in extreme cases.

`main_test` provides a test code through which you can find how the application 
is used.
