// This file was generated by counterfeiter
package fakes

import (
	"sync"

	"github.com/concourse/atc"
	"github.com/concourse/atc/db"
	"github.com/concourse/atc/scheduler"
)

type FakeSchedulerDB struct {
	ScheduleBuildStub        func(buildID int, serial bool) (bool, error)
	scheduleBuildMutex       sync.RWMutex
	scheduleBuildArgsForCall []struct {
		buildID int
		serial  bool
	}
	scheduleBuildReturns struct {
		result1 bool
		result2 error
	}
	GetLatestInputVersionsStub        func([]atc.JobInputConfig) (db.VersionedResources, error)
	getLatestInputVersionsMutex       sync.RWMutex
	getLatestInputVersionsArgsForCall []struct {
		arg1 []atc.JobInputConfig
	}
	getLatestInputVersionsReturns struct {
		result1 db.VersionedResources
		result2 error
	}
	GetJobBuildForInputsStub        func(job string, inputs []db.BuildInput) (db.Build, error)
	getJobBuildForInputsMutex       sync.RWMutex
	getJobBuildForInputsArgsForCall []struct {
		job    string
		inputs []db.BuildInput
	}
	getJobBuildForInputsReturns struct {
		result1 db.Build
		result2 error
	}
	CreateJobBuildWithInputsStub        func(job string, inputs []db.BuildInput) (db.Build, error)
	createJobBuildWithInputsMutex       sync.RWMutex
	createJobBuildWithInputsArgsForCall []struct {
		job    string
		inputs []db.BuildInput
	}
	createJobBuildWithInputsReturns struct {
		result1 db.Build
		result2 error
	}
	GetNextPendingBuildStub        func(job string) (db.Build, []db.BuildInput, error)
	getNextPendingBuildMutex       sync.RWMutex
	getNextPendingBuildArgsForCall []struct {
		job string
	}
	getNextPendingBuildReturns struct {
		result1 db.Build
		result2 []db.BuildInput
		result3 error
	}
	GetAllStartedBuildsStub        func() ([]db.Build, error)
	getAllStartedBuildsMutex       sync.RWMutex
	getAllStartedBuildsArgsForCall []struct{}
	getAllStartedBuildsReturns struct {
		result1 []db.Build
		result2 error
	}
}

func (fake *FakeSchedulerDB) ScheduleBuild(buildID int, serial bool) (bool, error) {
	fake.scheduleBuildMutex.Lock()
	fake.scheduleBuildArgsForCall = append(fake.scheduleBuildArgsForCall, struct {
		buildID int
		serial  bool
	}{buildID, serial})
	fake.scheduleBuildMutex.Unlock()
	if fake.ScheduleBuildStub != nil {
		return fake.ScheduleBuildStub(buildID, serial)
	} else {
		return fake.scheduleBuildReturns.result1, fake.scheduleBuildReturns.result2
	}
}

func (fake *FakeSchedulerDB) ScheduleBuildCallCount() int {
	fake.scheduleBuildMutex.RLock()
	defer fake.scheduleBuildMutex.RUnlock()
	return len(fake.scheduleBuildArgsForCall)
}

func (fake *FakeSchedulerDB) ScheduleBuildArgsForCall(i int) (int, bool) {
	fake.scheduleBuildMutex.RLock()
	defer fake.scheduleBuildMutex.RUnlock()
	return fake.scheduleBuildArgsForCall[i].buildID, fake.scheduleBuildArgsForCall[i].serial
}

func (fake *FakeSchedulerDB) ScheduleBuildReturns(result1 bool, result2 error) {
	fake.ScheduleBuildStub = nil
	fake.scheduleBuildReturns = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *FakeSchedulerDB) GetLatestInputVersions(arg1 []atc.JobInputConfig) (db.VersionedResources, error) {
	fake.getLatestInputVersionsMutex.Lock()
	fake.getLatestInputVersionsArgsForCall = append(fake.getLatestInputVersionsArgsForCall, struct {
		arg1 []atc.JobInputConfig
	}{arg1})
	fake.getLatestInputVersionsMutex.Unlock()
	if fake.GetLatestInputVersionsStub != nil {
		return fake.GetLatestInputVersionsStub(arg1)
	} else {
		return fake.getLatestInputVersionsReturns.result1, fake.getLatestInputVersionsReturns.result2
	}
}

func (fake *FakeSchedulerDB) GetLatestInputVersionsCallCount() int {
	fake.getLatestInputVersionsMutex.RLock()
	defer fake.getLatestInputVersionsMutex.RUnlock()
	return len(fake.getLatestInputVersionsArgsForCall)
}

func (fake *FakeSchedulerDB) GetLatestInputVersionsArgsForCall(i int) []atc.JobInputConfig {
	fake.getLatestInputVersionsMutex.RLock()
	defer fake.getLatestInputVersionsMutex.RUnlock()
	return fake.getLatestInputVersionsArgsForCall[i].arg1
}

func (fake *FakeSchedulerDB) GetLatestInputVersionsReturns(result1 db.VersionedResources, result2 error) {
	fake.GetLatestInputVersionsStub = nil
	fake.getLatestInputVersionsReturns = struct {
		result1 db.VersionedResources
		result2 error
	}{result1, result2}
}

func (fake *FakeSchedulerDB) GetJobBuildForInputs(job string, inputs []db.BuildInput) (db.Build, error) {
	fake.getJobBuildForInputsMutex.Lock()
	fake.getJobBuildForInputsArgsForCall = append(fake.getJobBuildForInputsArgsForCall, struct {
		job    string
		inputs []db.BuildInput
	}{job, inputs})
	fake.getJobBuildForInputsMutex.Unlock()
	if fake.GetJobBuildForInputsStub != nil {
		return fake.GetJobBuildForInputsStub(job, inputs)
	} else {
		return fake.getJobBuildForInputsReturns.result1, fake.getJobBuildForInputsReturns.result2
	}
}

func (fake *FakeSchedulerDB) GetJobBuildForInputsCallCount() int {
	fake.getJobBuildForInputsMutex.RLock()
	defer fake.getJobBuildForInputsMutex.RUnlock()
	return len(fake.getJobBuildForInputsArgsForCall)
}

func (fake *FakeSchedulerDB) GetJobBuildForInputsArgsForCall(i int) (string, []db.BuildInput) {
	fake.getJobBuildForInputsMutex.RLock()
	defer fake.getJobBuildForInputsMutex.RUnlock()
	return fake.getJobBuildForInputsArgsForCall[i].job, fake.getJobBuildForInputsArgsForCall[i].inputs
}

func (fake *FakeSchedulerDB) GetJobBuildForInputsReturns(result1 db.Build, result2 error) {
	fake.GetJobBuildForInputsStub = nil
	fake.getJobBuildForInputsReturns = struct {
		result1 db.Build
		result2 error
	}{result1, result2}
}

func (fake *FakeSchedulerDB) CreateJobBuildWithInputs(job string, inputs []db.BuildInput) (db.Build, error) {
	fake.createJobBuildWithInputsMutex.Lock()
	fake.createJobBuildWithInputsArgsForCall = append(fake.createJobBuildWithInputsArgsForCall, struct {
		job    string
		inputs []db.BuildInput
	}{job, inputs})
	fake.createJobBuildWithInputsMutex.Unlock()
	if fake.CreateJobBuildWithInputsStub != nil {
		return fake.CreateJobBuildWithInputsStub(job, inputs)
	} else {
		return fake.createJobBuildWithInputsReturns.result1, fake.createJobBuildWithInputsReturns.result2
	}
}

func (fake *FakeSchedulerDB) CreateJobBuildWithInputsCallCount() int {
	fake.createJobBuildWithInputsMutex.RLock()
	defer fake.createJobBuildWithInputsMutex.RUnlock()
	return len(fake.createJobBuildWithInputsArgsForCall)
}

func (fake *FakeSchedulerDB) CreateJobBuildWithInputsArgsForCall(i int) (string, []db.BuildInput) {
	fake.createJobBuildWithInputsMutex.RLock()
	defer fake.createJobBuildWithInputsMutex.RUnlock()
	return fake.createJobBuildWithInputsArgsForCall[i].job, fake.createJobBuildWithInputsArgsForCall[i].inputs
}

func (fake *FakeSchedulerDB) CreateJobBuildWithInputsReturns(result1 db.Build, result2 error) {
	fake.CreateJobBuildWithInputsStub = nil
	fake.createJobBuildWithInputsReturns = struct {
		result1 db.Build
		result2 error
	}{result1, result2}
}

func (fake *FakeSchedulerDB) GetNextPendingBuild(job string) (db.Build, []db.BuildInput, error) {
	fake.getNextPendingBuildMutex.Lock()
	fake.getNextPendingBuildArgsForCall = append(fake.getNextPendingBuildArgsForCall, struct {
		job string
	}{job})
	fake.getNextPendingBuildMutex.Unlock()
	if fake.GetNextPendingBuildStub != nil {
		return fake.GetNextPendingBuildStub(job)
	} else {
		return fake.getNextPendingBuildReturns.result1, fake.getNextPendingBuildReturns.result2, fake.getNextPendingBuildReturns.result3
	}
}

func (fake *FakeSchedulerDB) GetNextPendingBuildCallCount() int {
	fake.getNextPendingBuildMutex.RLock()
	defer fake.getNextPendingBuildMutex.RUnlock()
	return len(fake.getNextPendingBuildArgsForCall)
}

func (fake *FakeSchedulerDB) GetNextPendingBuildArgsForCall(i int) string {
	fake.getNextPendingBuildMutex.RLock()
	defer fake.getNextPendingBuildMutex.RUnlock()
	return fake.getNextPendingBuildArgsForCall[i].job
}

func (fake *FakeSchedulerDB) GetNextPendingBuildReturns(result1 db.Build, result2 []db.BuildInput, result3 error) {
	fake.GetNextPendingBuildStub = nil
	fake.getNextPendingBuildReturns = struct {
		result1 db.Build
		result2 []db.BuildInput
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeSchedulerDB) GetAllStartedBuilds() ([]db.Build, error) {
	fake.getAllStartedBuildsMutex.Lock()
	fake.getAllStartedBuildsArgsForCall = append(fake.getAllStartedBuildsArgsForCall, struct{}{})
	fake.getAllStartedBuildsMutex.Unlock()
	if fake.GetAllStartedBuildsStub != nil {
		return fake.GetAllStartedBuildsStub()
	} else {
		return fake.getAllStartedBuildsReturns.result1, fake.getAllStartedBuildsReturns.result2
	}
}

func (fake *FakeSchedulerDB) GetAllStartedBuildsCallCount() int {
	fake.getAllStartedBuildsMutex.RLock()
	defer fake.getAllStartedBuildsMutex.RUnlock()
	return len(fake.getAllStartedBuildsArgsForCall)
}

func (fake *FakeSchedulerDB) GetAllStartedBuildsReturns(result1 []db.Build, result2 error) {
	fake.GetAllStartedBuildsStub = nil
	fake.getAllStartedBuildsReturns = struct {
		result1 []db.Build
		result2 error
	}{result1, result2}
}

var _ scheduler.SchedulerDB = new(FakeSchedulerDB)
