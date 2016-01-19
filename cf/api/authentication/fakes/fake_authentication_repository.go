// This file was generated by counterfeiter
package fakes

import (
	"sync"

	"github.com/cloudfoundry/cli/cf/api/authentication"
	"github.com/cloudfoundry/cli/cf/configuration/core_config"
)

type FakeAuthenticationRepository struct {
	RefreshAuthTokenStub        func() (updatedToken string, apiErr error)
	refreshAuthTokenMutex       sync.RWMutex
	refreshAuthTokenArgsForCall []struct{}
	refreshAuthTokenReturns     struct {
		result1 string
		result2 error
	}
	AuthenticateStub        func(credentials map[string]string) (apiErr error)
	authenticateMutex       sync.RWMutex
	authenticateArgsForCall []struct {
		credentials map[string]string
	}
	authenticateReturns struct {
		result1 error
	}
	GetLoginPromptsAndSaveUAAServerURLStub        func() (map[string]core_config.AuthPrompt, error)
	getLoginPromptsAndSaveUAAServerURLMutex       sync.RWMutex
	getLoginPromptsAndSaveUAAServerURLArgsForCall []struct{}
	getLoginPromptsAndSaveUAAServerURLReturns     struct {
		result1 map[string]core_config.AuthPrompt
		result2 error
	}
}

func (fake *FakeAuthenticationRepository) RefreshAuthToken() (updatedToken string, apiErr error) {
	fake.refreshAuthTokenMutex.Lock()
	fake.refreshAuthTokenArgsForCall = append(fake.refreshAuthTokenArgsForCall, struct{}{})
	fake.refreshAuthTokenMutex.Unlock()
	if fake.RefreshAuthTokenStub != nil {
		return fake.RefreshAuthTokenStub()
	} else {
		return fake.refreshAuthTokenReturns.result1, fake.refreshAuthTokenReturns.result2
	}
}

func (fake *FakeAuthenticationRepository) RefreshAuthTokenCallCount() int {
	fake.refreshAuthTokenMutex.RLock()
	defer fake.refreshAuthTokenMutex.RUnlock()
	return len(fake.refreshAuthTokenArgsForCall)
}

func (fake *FakeAuthenticationRepository) RefreshAuthTokenReturns(result1 string, result2 error) {
	fake.RefreshAuthTokenStub = nil
	fake.refreshAuthTokenReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeAuthenticationRepository) Authenticate(credentials map[string]string) (apiErr error) {
	fake.authenticateMutex.Lock()
	fake.authenticateArgsForCall = append(fake.authenticateArgsForCall, struct {
		credentials map[string]string
	}{credentials})
	fake.authenticateMutex.Unlock()
	if fake.AuthenticateStub != nil {
		return fake.AuthenticateStub(credentials)
	} else {
		return fake.authenticateReturns.result1
	}
}

func (fake *FakeAuthenticationRepository) AuthenticateCallCount() int {
	fake.authenticateMutex.RLock()
	defer fake.authenticateMutex.RUnlock()
	return len(fake.authenticateArgsForCall)
}

func (fake *FakeAuthenticationRepository) AuthenticateArgsForCall(i int) map[string]string {
	fake.authenticateMutex.RLock()
	defer fake.authenticateMutex.RUnlock()
	return fake.authenticateArgsForCall[i].credentials
}

func (fake *FakeAuthenticationRepository) AuthenticateReturns(result1 error) {
	fake.AuthenticateStub = nil
	fake.authenticateReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeAuthenticationRepository) GetLoginPromptsAndSaveUAAServerURL() (map[string]core_config.AuthPrompt, error) {
	fake.getLoginPromptsAndSaveUAAServerURLMutex.Lock()
	fake.getLoginPromptsAndSaveUAAServerURLArgsForCall = append(fake.getLoginPromptsAndSaveUAAServerURLArgsForCall, struct{}{})
	fake.getLoginPromptsAndSaveUAAServerURLMutex.Unlock()
	if fake.GetLoginPromptsAndSaveUAAServerURLStub != nil {
		return fake.GetLoginPromptsAndSaveUAAServerURLStub()
	} else {
		return fake.getLoginPromptsAndSaveUAAServerURLReturns.result1, fake.getLoginPromptsAndSaveUAAServerURLReturns.result2
	}
}

func (fake *FakeAuthenticationRepository) GetLoginPromptsAndSaveUAAServerURLCallCount() int {
	fake.getLoginPromptsAndSaveUAAServerURLMutex.RLock()
	defer fake.getLoginPromptsAndSaveUAAServerURLMutex.RUnlock()
	return len(fake.getLoginPromptsAndSaveUAAServerURLArgsForCall)
}

func (fake *FakeAuthenticationRepository) GetLoginPromptsAndSaveUAAServerURLReturns(result1 map[string]core_config.AuthPrompt, result2 error) {
	fake.GetLoginPromptsAndSaveUAAServerURLStub = nil
	fake.getLoginPromptsAndSaveUAAServerURLReturns = struct {
		result1 map[string]core_config.AuthPrompt
		result2 error
	}{result1, result2}
}

var _ authentication.AuthenticationRepository = new(FakeAuthenticationRepository)