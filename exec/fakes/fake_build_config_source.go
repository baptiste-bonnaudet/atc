// This file was generated by counterfeiter
package fakes

import (
	"sync"

	"github.com/concourse/atc"
	"github.com/concourse/atc/exec"
)

type FakeTaskConfigSource struct {
	FetchConfigStub        func(exec.ArtifactSource) (atc.TaskConfig, error)
	fetchConfigMutex       sync.RWMutex
	fetchConfigArgsForCall []struct {
		arg1 exec.ArtifactSource
	}
	fetchConfigReturns struct {
		result1 atc.TaskConfig
		result2 error
	}
}

func (fake *FakeTaskConfigSource) FetchConfig(arg1 exec.ArtifactSource) (atc.TaskConfig, error) {
	fake.fetchConfigMutex.Lock()
	fake.fetchConfigArgsForCall = append(fake.fetchConfigArgsForCall, struct {
		arg1 exec.ArtifactSource
	}{arg1})
	fake.fetchConfigMutex.Unlock()
	if fake.FetchConfigStub != nil {
		return fake.FetchConfigStub(arg1)
	} else {
		return fake.fetchConfigReturns.result1, fake.fetchConfigReturns.result2
	}
}

func (fake *FakeTaskConfigSource) FetchConfigCallCount() int {
	fake.fetchConfigMutex.RLock()
	defer fake.fetchConfigMutex.RUnlock()
	return len(fake.fetchConfigArgsForCall)
}

func (fake *FakeTaskConfigSource) FetchConfigArgsForCall(i int) exec.ArtifactSource {
	fake.fetchConfigMutex.RLock()
	defer fake.fetchConfigMutex.RUnlock()
	return fake.fetchConfigArgsForCall[i].arg1
}

func (fake *FakeTaskConfigSource) FetchConfigReturns(result1 atc.TaskConfig, result2 error) {
	fake.FetchConfigStub = nil
	fake.fetchConfigReturns = struct {
		result1 atc.TaskConfig
		result2 error
	}{result1, result2}
}

var _ exec.TaskConfigSource = new(FakeTaskConfigSource)
