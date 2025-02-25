// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"sync"

	"github.com/pivotal-cf/kiln/pkg/planitest/internal"
)

type OpsManifestRunner struct {
	GetManifestStub        func(string, string) (map[string]interface{}, error)
	getManifestMutex       sync.RWMutex
	getManifestArgsForCall []struct {
		arg1 string
		arg2 string
	}
	getManifestReturns struct {
		result1 map[string]interface{}
		result2 error
	}
	getManifestReturnsOnCall map[int]struct {
		result1 map[string]interface{}
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *OpsManifestRunner) GetManifest(arg1 string, arg2 string) (map[string]interface{}, error) {
	fake.getManifestMutex.Lock()
	ret, specificReturn := fake.getManifestReturnsOnCall[len(fake.getManifestArgsForCall)]
	fake.getManifestArgsForCall = append(fake.getManifestArgsForCall, struct {
		arg1 string
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("GetManifest", []interface{}{arg1, arg2})
	fake.getManifestMutex.Unlock()
	if fake.GetManifestStub != nil {
		return fake.GetManifestStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.getManifestReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *OpsManifestRunner) GetManifestCallCount() int {
	fake.getManifestMutex.RLock()
	defer fake.getManifestMutex.RUnlock()
	return len(fake.getManifestArgsForCall)
}

func (fake *OpsManifestRunner) GetManifestCalls(stub func(string, string) (map[string]interface{}, error)) {
	fake.getManifestMutex.Lock()
	defer fake.getManifestMutex.Unlock()
	fake.GetManifestStub = stub
}

func (fake *OpsManifestRunner) GetManifestArgsForCall(i int) (string, string) {
	fake.getManifestMutex.RLock()
	defer fake.getManifestMutex.RUnlock()
	argsForCall := fake.getManifestArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *OpsManifestRunner) GetManifestReturns(result1 map[string]interface{}, result2 error) {
	fake.getManifestMutex.Lock()
	defer fake.getManifestMutex.Unlock()
	fake.GetManifestStub = nil
	fake.getManifestReturns = struct {
		result1 map[string]interface{}
		result2 error
	}{result1, result2}
}

func (fake *OpsManifestRunner) GetManifestReturnsOnCall(i int, result1 map[string]interface{}, result2 error) {
	fake.getManifestMutex.Lock()
	defer fake.getManifestMutex.Unlock()
	fake.GetManifestStub = nil
	if fake.getManifestReturnsOnCall == nil {
		fake.getManifestReturnsOnCall = make(map[int]struct {
			result1 map[string]interface{}
			result2 error
		})
	}
	fake.getManifestReturnsOnCall[i] = struct {
		result1 map[string]interface{}
		result2 error
	}{result1, result2}
}

func (fake *OpsManifestRunner) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getManifestMutex.RLock()
	defer fake.getManifestMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *OpsManifestRunner) recordInvocation(key string, args []interface{}) {
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

var _ internal.OpsManifestRunnerI = new(OpsManifestRunner)
