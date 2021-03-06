// This file was generated by counterfeiter
package fakes

import (
	"sync"

	. "github.com/maxbrunsfeld/counterfeiter/fixtures"
)

type FakeReusesArgTypes struct {
	DoThingsStub        func(x, y string)
	doThingsMutex       sync.RWMutex
	doThingsArgsForCall []struct {
		x string
		y string
	}
}

func (fake *FakeReusesArgTypes) DoThings(x string, y string) {
	fake.doThingsMutex.Lock()
	defer fake.doThingsMutex.Unlock()
	fake.doThingsArgsForCall = append(fake.doThingsArgsForCall, struct {
		x string
		y string
	}{x, y})
	if fake.DoThingsStub != nil {
		fake.DoThingsStub(x, y)
	}
}

func (fake *FakeReusesArgTypes) DoThingsCallCount() int {
	fake.doThingsMutex.RLock()
	defer fake.doThingsMutex.RUnlock()
	return len(fake.doThingsArgsForCall)
}

func (fake *FakeReusesArgTypes) DoThingsArgsForCall(i int) (string, string) {
	fake.doThingsMutex.RLock()
	defer fake.doThingsMutex.RUnlock()
	return fake.doThingsArgsForCall[i].x, fake.doThingsArgsForCall[i].y
}

var _ ReusesArgTypes = new(FakeReusesArgTypes)
