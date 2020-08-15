// Code generated by counterfeiter. DO NOT EDIT.
package controllersfakes

import (
	"sync"

	"github.com/jace-ys/sentry-operator/controllers"
	"github.com/jace-ys/sentry-operator/pkg/sentry"
)

type FakeSentryTeams struct {
	CreateProjectStub        func(string, string, *sentry.CreateProjectParams) (*sentry.Project, *sentry.Response, error)
	createProjectMutex       sync.RWMutex
	createProjectArgsForCall []struct {
		arg1 string
		arg2 string
		arg3 *sentry.CreateProjectParams
	}
	createProjectReturns struct {
		result1 *sentry.Project
		result2 *sentry.Response
		result3 error
	}
	createProjectReturnsOnCall map[int]struct {
		result1 *sentry.Project
		result2 *sentry.Response
		result3 error
	}
	ListProjectsStub        func(string, string) ([]sentry.Project, *sentry.Response, error)
	listProjectsMutex       sync.RWMutex
	listProjectsArgsForCall []struct {
		arg1 string
		arg2 string
	}
	listProjectsReturns struct {
		result1 []sentry.Project
		result2 *sentry.Response
		result3 error
	}
	listProjectsReturnsOnCall map[int]struct {
		result1 []sentry.Project
		result2 *sentry.Response
		result3 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeSentryTeams) CreateProject(arg1 string, arg2 string, arg3 *sentry.CreateProjectParams) (*sentry.Project, *sentry.Response, error) {
	fake.createProjectMutex.Lock()
	ret, specificReturn := fake.createProjectReturnsOnCall[len(fake.createProjectArgsForCall)]
	fake.createProjectArgsForCall = append(fake.createProjectArgsForCall, struct {
		arg1 string
		arg2 string
		arg3 *sentry.CreateProjectParams
	}{arg1, arg2, arg3})
	fake.recordInvocation("CreateProject", []interface{}{arg1, arg2, arg3})
	fake.createProjectMutex.Unlock()
	if fake.CreateProjectStub != nil {
		return fake.CreateProjectStub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	fakeReturns := fake.createProjectReturns
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3
}

func (fake *FakeSentryTeams) CreateProjectCallCount() int {
	fake.createProjectMutex.RLock()
	defer fake.createProjectMutex.RUnlock()
	return len(fake.createProjectArgsForCall)
}

func (fake *FakeSentryTeams) CreateProjectCalls(stub func(string, string, *sentry.CreateProjectParams) (*sentry.Project, *sentry.Response, error)) {
	fake.createProjectMutex.Lock()
	defer fake.createProjectMutex.Unlock()
	fake.CreateProjectStub = stub
}

func (fake *FakeSentryTeams) CreateProjectArgsForCall(i int) (string, string, *sentry.CreateProjectParams) {
	fake.createProjectMutex.RLock()
	defer fake.createProjectMutex.RUnlock()
	argsForCall := fake.createProjectArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeSentryTeams) CreateProjectReturns(result1 *sentry.Project, result2 *sentry.Response, result3 error) {
	fake.createProjectMutex.Lock()
	defer fake.createProjectMutex.Unlock()
	fake.CreateProjectStub = nil
	fake.createProjectReturns = struct {
		result1 *sentry.Project
		result2 *sentry.Response
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeSentryTeams) CreateProjectReturnsOnCall(i int, result1 *sentry.Project, result2 *sentry.Response, result3 error) {
	fake.createProjectMutex.Lock()
	defer fake.createProjectMutex.Unlock()
	fake.CreateProjectStub = nil
	if fake.createProjectReturnsOnCall == nil {
		fake.createProjectReturnsOnCall = make(map[int]struct {
			result1 *sentry.Project
			result2 *sentry.Response
			result3 error
		})
	}
	fake.createProjectReturnsOnCall[i] = struct {
		result1 *sentry.Project
		result2 *sentry.Response
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeSentryTeams) ListProjects(arg1 string, arg2 string) ([]sentry.Project, *sentry.Response, error) {
	fake.listProjectsMutex.Lock()
	ret, specificReturn := fake.listProjectsReturnsOnCall[len(fake.listProjectsArgsForCall)]
	fake.listProjectsArgsForCall = append(fake.listProjectsArgsForCall, struct {
		arg1 string
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("ListProjects", []interface{}{arg1, arg2})
	fake.listProjectsMutex.Unlock()
	if fake.ListProjectsStub != nil {
		return fake.ListProjectsStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	fakeReturns := fake.listProjectsReturns
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3
}

func (fake *FakeSentryTeams) ListProjectsCallCount() int {
	fake.listProjectsMutex.RLock()
	defer fake.listProjectsMutex.RUnlock()
	return len(fake.listProjectsArgsForCall)
}

func (fake *FakeSentryTeams) ListProjectsCalls(stub func(string, string) ([]sentry.Project, *sentry.Response, error)) {
	fake.listProjectsMutex.Lock()
	defer fake.listProjectsMutex.Unlock()
	fake.ListProjectsStub = stub
}

func (fake *FakeSentryTeams) ListProjectsArgsForCall(i int) (string, string) {
	fake.listProjectsMutex.RLock()
	defer fake.listProjectsMutex.RUnlock()
	argsForCall := fake.listProjectsArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeSentryTeams) ListProjectsReturns(result1 []sentry.Project, result2 *sentry.Response, result3 error) {
	fake.listProjectsMutex.Lock()
	defer fake.listProjectsMutex.Unlock()
	fake.ListProjectsStub = nil
	fake.listProjectsReturns = struct {
		result1 []sentry.Project
		result2 *sentry.Response
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeSentryTeams) ListProjectsReturnsOnCall(i int, result1 []sentry.Project, result2 *sentry.Response, result3 error) {
	fake.listProjectsMutex.Lock()
	defer fake.listProjectsMutex.Unlock()
	fake.ListProjectsStub = nil
	if fake.listProjectsReturnsOnCall == nil {
		fake.listProjectsReturnsOnCall = make(map[int]struct {
			result1 []sentry.Project
			result2 *sentry.Response
			result3 error
		})
	}
	fake.listProjectsReturnsOnCall[i] = struct {
		result1 []sentry.Project
		result2 *sentry.Response
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeSentryTeams) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.createProjectMutex.RLock()
	defer fake.createProjectMutex.RUnlock()
	fake.listProjectsMutex.RLock()
	defer fake.listProjectsMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeSentryTeams) recordInvocation(key string, args []interface{}) {
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

var _ controllers.SentryTeams = new(FakeSentryTeams)