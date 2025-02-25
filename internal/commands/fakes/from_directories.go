// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"sync"
)

type FromDirectories struct {
	FromDirectoriesStub        func([]string) (map[string]interface{}, error)
	fromDirectoriesMutex       sync.RWMutex
	fromDirectoriesArgsForCall []struct {
		arg1 []string
	}
	fromDirectoriesReturns struct {
		result1 map[string]interface{}
		result2 error
	}
	fromDirectoriesReturnsOnCall map[int]struct {
		result1 map[string]interface{}
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FromDirectories) FromDirectories(arg1 []string) (map[string]interface{}, error) {
	var arg1Copy []string
	if arg1 != nil {
		arg1Copy = make([]string, len(arg1))
		copy(arg1Copy, arg1)
	}
	fake.fromDirectoriesMutex.Lock()
	ret, specificReturn := fake.fromDirectoriesReturnsOnCall[len(fake.fromDirectoriesArgsForCall)]
	fake.fromDirectoriesArgsForCall = append(fake.fromDirectoriesArgsForCall, struct {
		arg1 []string
	}{arg1Copy})
	stub := fake.FromDirectoriesStub
	fakeReturns := fake.fromDirectoriesReturns
	fake.recordInvocation("FromDirectories", []interface{}{arg1Copy})
	fake.fromDirectoriesMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FromDirectories) FromDirectoriesCallCount() int {
	fake.fromDirectoriesMutex.RLock()
	defer fake.fromDirectoriesMutex.RUnlock()
	return len(fake.fromDirectoriesArgsForCall)
}

func (fake *FromDirectories) FromDirectoriesCalls(stub func([]string) (map[string]interface{}, error)) {
	fake.fromDirectoriesMutex.Lock()
	defer fake.fromDirectoriesMutex.Unlock()
	fake.FromDirectoriesStub = stub
}

func (fake *FromDirectories) FromDirectoriesArgsForCall(i int) []string {
	fake.fromDirectoriesMutex.RLock()
	defer fake.fromDirectoriesMutex.RUnlock()
	argsForCall := fake.fromDirectoriesArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FromDirectories) FromDirectoriesReturns(result1 map[string]interface{}, result2 error) {
	fake.fromDirectoriesMutex.Lock()
	defer fake.fromDirectoriesMutex.Unlock()
	fake.FromDirectoriesStub = nil
	fake.fromDirectoriesReturns = struct {
		result1 map[string]interface{}
		result2 error
	}{result1, result2}
}

func (fake *FromDirectories) FromDirectoriesReturnsOnCall(i int, result1 map[string]interface{}, result2 error) {
	fake.fromDirectoriesMutex.Lock()
	defer fake.fromDirectoriesMutex.Unlock()
	fake.FromDirectoriesStub = nil
	if fake.fromDirectoriesReturnsOnCall == nil {
		fake.fromDirectoriesReturnsOnCall = make(map[int]struct {
			result1 map[string]interface{}
			result2 error
		})
	}
	fake.fromDirectoriesReturnsOnCall[i] = struct {
		result1 map[string]interface{}
		result2 error
	}{result1, result2}
}

func (fake *FromDirectories) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.fromDirectoriesMutex.RLock()
	defer fake.fromDirectoriesMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FromDirectories) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}
