// This file was generated by counterfeiter
package fakes

import (
	"io"
	"os"
	"sync"
)

type Zipper struct {
	SetPathStub        func(path string) error
	setPathMutex       sync.RWMutex
	setPathArgsForCall []struct {
		path string
	}
	setPathReturns struct {
		result1 error
	}
	setPathReturnsOnCall map[int]struct {
		result1 error
	}
	AddStub        func(path string, file io.Reader) error
	addMutex       sync.RWMutex
	addArgsForCall []struct {
		path string
		file io.Reader
	}
	addReturns struct {
		result1 error
	}
	addReturnsOnCall map[int]struct {
		result1 error
	}
	AddWithModeStub        func(path string, file io.Reader, mode os.FileMode) error
	addWithModeMutex       sync.RWMutex
	addWithModeArgsForCall []struct {
		path string
		file io.Reader
		mode os.FileMode
	}
	addWithModeReturns struct {
		result1 error
	}
	addWithModeReturnsOnCall map[int]struct {
		result1 error
	}
	CreateFolderStub        func(path string) error
	createFolderMutex       sync.RWMutex
	createFolderArgsForCall []struct {
		path string
	}
	createFolderReturns struct {
		result1 error
	}
	createFolderReturnsOnCall map[int]struct {
		result1 error
	}
	CloseStub        func() error
	closeMutex       sync.RWMutex
	closeArgsForCall []struct{}
	closeReturns     struct {
		result1 error
	}
	closeReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *Zipper) SetPath(path string) error {
	fake.setPathMutex.Lock()
	ret, specificReturn := fake.setPathReturnsOnCall[len(fake.setPathArgsForCall)]
	fake.setPathArgsForCall = append(fake.setPathArgsForCall, struct {
		path string
	}{path})
	fake.recordInvocation("SetPath", []interface{}{path})
	fake.setPathMutex.Unlock()
	if fake.SetPathStub != nil {
		return fake.SetPathStub(path)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.setPathReturns.result1
}

func (fake *Zipper) SetPathCallCount() int {
	fake.setPathMutex.RLock()
	defer fake.setPathMutex.RUnlock()
	return len(fake.setPathArgsForCall)
}

func (fake *Zipper) SetPathArgsForCall(i int) string {
	fake.setPathMutex.RLock()
	defer fake.setPathMutex.RUnlock()
	return fake.setPathArgsForCall[i].path
}

func (fake *Zipper) SetPathReturns(result1 error) {
	fake.SetPathStub = nil
	fake.setPathReturns = struct {
		result1 error
	}{result1}
}

func (fake *Zipper) SetPathReturnsOnCall(i int, result1 error) {
	fake.SetPathStub = nil
	if fake.setPathReturnsOnCall == nil {
		fake.setPathReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.setPathReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *Zipper) Add(path string, file io.Reader) error {
	fake.addMutex.Lock()
	ret, specificReturn := fake.addReturnsOnCall[len(fake.addArgsForCall)]
	fake.addArgsForCall = append(fake.addArgsForCall, struct {
		path string
		file io.Reader
	}{path, file})
	fake.recordInvocation("Add", []interface{}{path, file})
	fake.addMutex.Unlock()
	if fake.AddStub != nil {
		return fake.AddStub(path, file)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.addReturns.result1
}

func (fake *Zipper) AddCallCount() int {
	fake.addMutex.RLock()
	defer fake.addMutex.RUnlock()
	return len(fake.addArgsForCall)
}

func (fake *Zipper) AddArgsForCall(i int) (string, io.Reader) {
	fake.addMutex.RLock()
	defer fake.addMutex.RUnlock()
	return fake.addArgsForCall[i].path, fake.addArgsForCall[i].file
}

func (fake *Zipper) AddReturns(result1 error) {
	fake.AddStub = nil
	fake.addReturns = struct {
		result1 error
	}{result1}
}

func (fake *Zipper) AddReturnsOnCall(i int, result1 error) {
	fake.AddStub = nil
	if fake.addReturnsOnCall == nil {
		fake.addReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.addReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *Zipper) AddWithMode(path string, file io.Reader, mode os.FileMode) error {
	fake.addWithModeMutex.Lock()
	ret, specificReturn := fake.addWithModeReturnsOnCall[len(fake.addWithModeArgsForCall)]
	fake.addWithModeArgsForCall = append(fake.addWithModeArgsForCall, struct {
		path string
		file io.Reader
		mode os.FileMode
	}{path, file, mode})
	fake.recordInvocation("AddWithMode", []interface{}{path, file, mode})
	fake.addWithModeMutex.Unlock()
	if fake.AddWithModeStub != nil {
		return fake.AddWithModeStub(path, file, mode)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.addWithModeReturns.result1
}

func (fake *Zipper) AddWithModeCallCount() int {
	fake.addWithModeMutex.RLock()
	defer fake.addWithModeMutex.RUnlock()
	return len(fake.addWithModeArgsForCall)
}

func (fake *Zipper) AddWithModeArgsForCall(i int) (string, io.Reader, os.FileMode) {
	fake.addWithModeMutex.RLock()
	defer fake.addWithModeMutex.RUnlock()
	return fake.addWithModeArgsForCall[i].path, fake.addWithModeArgsForCall[i].file, fake.addWithModeArgsForCall[i].mode
}

func (fake *Zipper) AddWithModeReturns(result1 error) {
	fake.AddWithModeStub = nil
	fake.addWithModeReturns = struct {
		result1 error
	}{result1}
}

func (fake *Zipper) AddWithModeReturnsOnCall(i int, result1 error) {
	fake.AddWithModeStub = nil
	if fake.addWithModeReturnsOnCall == nil {
		fake.addWithModeReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.addWithModeReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *Zipper) CreateFolder(path string) error {
	fake.createFolderMutex.Lock()
	ret, specificReturn := fake.createFolderReturnsOnCall[len(fake.createFolderArgsForCall)]
	fake.createFolderArgsForCall = append(fake.createFolderArgsForCall, struct {
		path string
	}{path})
	fake.recordInvocation("CreateFolder", []interface{}{path})
	fake.createFolderMutex.Unlock()
	if fake.CreateFolderStub != nil {
		return fake.CreateFolderStub(path)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.createFolderReturns.result1
}

func (fake *Zipper) CreateFolderCallCount() int {
	fake.createFolderMutex.RLock()
	defer fake.createFolderMutex.RUnlock()
	return len(fake.createFolderArgsForCall)
}

func (fake *Zipper) CreateFolderArgsForCall(i int) string {
	fake.createFolderMutex.RLock()
	defer fake.createFolderMutex.RUnlock()
	return fake.createFolderArgsForCall[i].path
}

func (fake *Zipper) CreateFolderReturns(result1 error) {
	fake.CreateFolderStub = nil
	fake.createFolderReturns = struct {
		result1 error
	}{result1}
}

func (fake *Zipper) CreateFolderReturnsOnCall(i int, result1 error) {
	fake.CreateFolderStub = nil
	if fake.createFolderReturnsOnCall == nil {
		fake.createFolderReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.createFolderReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *Zipper) Close() error {
	fake.closeMutex.Lock()
	ret, specificReturn := fake.closeReturnsOnCall[len(fake.closeArgsForCall)]
	fake.closeArgsForCall = append(fake.closeArgsForCall, struct{}{})
	fake.recordInvocation("Close", []interface{}{})
	fake.closeMutex.Unlock()
	if fake.CloseStub != nil {
		return fake.CloseStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.closeReturns.result1
}

func (fake *Zipper) CloseCallCount() int {
	fake.closeMutex.RLock()
	defer fake.closeMutex.RUnlock()
	return len(fake.closeArgsForCall)
}

func (fake *Zipper) CloseReturns(result1 error) {
	fake.CloseStub = nil
	fake.closeReturns = struct {
		result1 error
	}{result1}
}

func (fake *Zipper) CloseReturnsOnCall(i int, result1 error) {
	fake.CloseStub = nil
	if fake.closeReturnsOnCall == nil {
		fake.closeReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.closeReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *Zipper) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.setPathMutex.RLock()
	defer fake.setPathMutex.RUnlock()
	fake.addMutex.RLock()
	defer fake.addMutex.RUnlock()
	fake.addWithModeMutex.RLock()
	defer fake.addWithModeMutex.RUnlock()
	fake.createFolderMutex.RLock()
	defer fake.createFolderMutex.RUnlock()
	fake.closeMutex.RLock()
	defer fake.closeMutex.RUnlock()
	return fake.invocations
}

func (fake *Zipper) recordInvocation(key string, args []interface{}) {
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
