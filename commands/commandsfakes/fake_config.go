// This file was generated by counterfeiter
package commandsfakes

import (
	"sync"

	"code.cloudfoundry.org/cli/commands"
	"code.cloudfoundry.org/cli/utils/config"
)

type FakeConfig struct {
	BinaryNameStub        func() string
	binaryNameMutex       sync.RWMutex
	binaryNameArgsForCall []struct{}
	binaryNameReturns     struct {
		result1 string
	}
	ColorEnabledStub        func() config.ColorSetting
	colorEnabledMutex       sync.RWMutex
	colorEnabledArgsForCall []struct{}
	colorEnabledReturns     struct {
		result1 config.ColorSetting
	}
	LocaleStub        func() string
	localeMutex       sync.RWMutex
	localeArgsForCall []struct{}
	localeReturns     struct {
		result1 string
	}
	PluginsStub        func() map[string]config.Plugin
	pluginsMutex       sync.RWMutex
	pluginsArgsForCall []struct{}
	pluginsReturns     struct {
		result1 map[string]config.Plugin
	}
	SetTargetInformationStub        func(api string, apiVersion string, auth string, loggregator string, doppler string, uaa string)
	setTargetInformationMutex       sync.RWMutex
	setTargetInformationArgsForCall []struct {
		api         string
		apiVersion  string
		auth        string
		loggregator string
		doppler     string
		uaa         string
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeConfig) BinaryName() string {
	fake.binaryNameMutex.Lock()
	fake.binaryNameArgsForCall = append(fake.binaryNameArgsForCall, struct{}{})
	fake.recordInvocation("BinaryName", []interface{}{})
	fake.binaryNameMutex.Unlock()
	if fake.BinaryNameStub != nil {
		return fake.BinaryNameStub()
	} else {
		return fake.binaryNameReturns.result1
	}
}

func (fake *FakeConfig) BinaryNameCallCount() int {
	fake.binaryNameMutex.RLock()
	defer fake.binaryNameMutex.RUnlock()
	return len(fake.binaryNameArgsForCall)
}

func (fake *FakeConfig) BinaryNameReturns(result1 string) {
	fake.BinaryNameStub = nil
	fake.binaryNameReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeConfig) ColorEnabled() config.ColorSetting {
	fake.colorEnabledMutex.Lock()
	fake.colorEnabledArgsForCall = append(fake.colorEnabledArgsForCall, struct{}{})
	fake.recordInvocation("ColorEnabled", []interface{}{})
	fake.colorEnabledMutex.Unlock()
	if fake.ColorEnabledStub != nil {
		return fake.ColorEnabledStub()
	} else {
		return fake.colorEnabledReturns.result1
	}
}

func (fake *FakeConfig) ColorEnabledCallCount() int {
	fake.colorEnabledMutex.RLock()
	defer fake.colorEnabledMutex.RUnlock()
	return len(fake.colorEnabledArgsForCall)
}

func (fake *FakeConfig) ColorEnabledReturns(result1 config.ColorSetting) {
	fake.ColorEnabledStub = nil
	fake.colorEnabledReturns = struct {
		result1 config.ColorSetting
	}{result1}
}

func (fake *FakeConfig) Locale() string {
	fake.localeMutex.Lock()
	fake.localeArgsForCall = append(fake.localeArgsForCall, struct{}{})
	fake.recordInvocation("Locale", []interface{}{})
	fake.localeMutex.Unlock()
	if fake.LocaleStub != nil {
		return fake.LocaleStub()
	} else {
		return fake.localeReturns.result1
	}
}

func (fake *FakeConfig) LocaleCallCount() int {
	fake.localeMutex.RLock()
	defer fake.localeMutex.RUnlock()
	return len(fake.localeArgsForCall)
}

func (fake *FakeConfig) LocaleReturns(result1 string) {
	fake.LocaleStub = nil
	fake.localeReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeConfig) Plugins() map[string]config.Plugin {
	fake.pluginsMutex.Lock()
	fake.pluginsArgsForCall = append(fake.pluginsArgsForCall, struct{}{})
	fake.recordInvocation("Plugins", []interface{}{})
	fake.pluginsMutex.Unlock()
	if fake.PluginsStub != nil {
		return fake.PluginsStub()
	} else {
		return fake.pluginsReturns.result1
	}
}

func (fake *FakeConfig) PluginsCallCount() int {
	fake.pluginsMutex.RLock()
	defer fake.pluginsMutex.RUnlock()
	return len(fake.pluginsArgsForCall)
}

func (fake *FakeConfig) PluginsReturns(result1 map[string]config.Plugin) {
	fake.PluginsStub = nil
	fake.pluginsReturns = struct {
		result1 map[string]config.Plugin
	}{result1}
}

func (fake *FakeConfig) SetTargetInformation(api string, apiVersion string, auth string, loggregator string, doppler string, uaa string) {
	fake.setTargetInformationMutex.Lock()
	fake.setTargetInformationArgsForCall = append(fake.setTargetInformationArgsForCall, struct {
		api         string
		apiVersion  string
		auth        string
		loggregator string
		doppler     string
		uaa         string
	}{api, apiVersion, auth, loggregator, doppler, uaa})
	fake.recordInvocation("SetTargetInformation", []interface{}{api, apiVersion, auth, loggregator, doppler, uaa})
	fake.setTargetInformationMutex.Unlock()
	if fake.SetTargetInformationStub != nil {
		fake.SetTargetInformationStub(api, apiVersion, auth, loggregator, doppler, uaa)
	}
}

func (fake *FakeConfig) SetTargetInformationCallCount() int {
	fake.setTargetInformationMutex.RLock()
	defer fake.setTargetInformationMutex.RUnlock()
	return len(fake.setTargetInformationArgsForCall)
}

func (fake *FakeConfig) SetTargetInformationArgsForCall(i int) (string, string, string, string, string, string) {
	fake.setTargetInformationMutex.RLock()
	defer fake.setTargetInformationMutex.RUnlock()
	return fake.setTargetInformationArgsForCall[i].api, fake.setTargetInformationArgsForCall[i].apiVersion, fake.setTargetInformationArgsForCall[i].auth, fake.setTargetInformationArgsForCall[i].loggregator, fake.setTargetInformationArgsForCall[i].doppler, fake.setTargetInformationArgsForCall[i].uaa
}

func (fake *FakeConfig) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.binaryNameMutex.RLock()
	defer fake.binaryNameMutex.RUnlock()
	fake.colorEnabledMutex.RLock()
	defer fake.colorEnabledMutex.RUnlock()
	fake.localeMutex.RLock()
	defer fake.localeMutex.RUnlock()
	fake.pluginsMutex.RLock()
	defer fake.pluginsMutex.RUnlock()
	fake.setTargetInformationMutex.RLock()
	defer fake.setTargetInformationMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeConfig) recordInvocation(key string, args []interface{}) {
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

var _ commands.Config = new(FakeConfig)
