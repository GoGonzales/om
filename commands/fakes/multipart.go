// This file was generated by counterfeiter
package fakes

import (
	"sync"

	"github.com/pivotal-cf/om/formcontent"
)

type Multipart struct {
	CreateStub        func(path string) (formcontent.ContentSubmission, error)
	createMutex       sync.RWMutex
	createArgsForCall []struct {
		path string
	}
	createReturns struct {
		result1 formcontent.ContentSubmission
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *Multipart) Create(path string) (formcontent.ContentSubmission, error) {
	fake.createMutex.Lock()
	fake.createArgsForCall = append(fake.createArgsForCall, struct {
		path string
	}{path})
	fake.recordInvocation("Create", []interface{}{path})
	fake.createMutex.Unlock()
	if fake.CreateStub != nil {
		return fake.CreateStub(path)
	} else {
		return fake.createReturns.result1, fake.createReturns.result2
	}
}

func (fake *Multipart) CreateCallCount() int {
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	return len(fake.createArgsForCall)
}

func (fake *Multipart) CreateArgsForCall(i int) string {
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	return fake.createArgsForCall[i].path
}

func (fake *Multipart) CreateReturns(result1 formcontent.ContentSubmission, result2 error) {
	fake.CreateStub = nil
	fake.createReturns = struct {
		result1 formcontent.ContentSubmission
		result2 error
	}{result1, result2}
}

func (fake *Multipart) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	return fake.invocations
}

func (fake *Multipart) recordInvocation(key string, args []interface{}) {
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