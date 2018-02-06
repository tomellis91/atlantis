// Automatically generated by pegomock. DO NOT EDIT!
// Source: github.com/atlantisnorth/atlantis/server/events/webhooks (interfaces: UnderlyingSlackClient)

package mocks

import (
	"reflect"

	slack "github.com/nlopes/slack"
	pegomock "github.com/petergtz/pegomock"
)

type MockUnderlyingSlackClient struct {
	fail func(message string, callerSkip ...int)
}

func NewMockUnderlyingSlackClient() *MockUnderlyingSlackClient {
	return &MockUnderlyingSlackClient{fail: pegomock.GlobalFailHandler}
}

func (mock *MockUnderlyingSlackClient) AuthTest() (*slack.AuthTestResponse, error) {
	params := []pegomock.Param{}
	result := pegomock.GetGenericMockFrom(mock).Invoke("AuthTest", params, []reflect.Type{reflect.TypeOf((**slack.AuthTestResponse)(nil)).Elem(), reflect.TypeOf((*error)(nil)).Elem()})
	var ret0 *slack.AuthTestResponse
	var ret1 error
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(*slack.AuthTestResponse)
		}
		if result[1] != nil {
			ret1 = result[1].(error)
		}
	}
	return ret0, ret1
}

func (mock *MockUnderlyingSlackClient) GetChannels(excludeArchived bool) ([]slack.Channel, error) {
	params := []pegomock.Param{excludeArchived}
	result := pegomock.GetGenericMockFrom(mock).Invoke("GetChannels", params, []reflect.Type{reflect.TypeOf((*[]slack.Channel)(nil)).Elem(), reflect.TypeOf((*error)(nil)).Elem()})
	var ret0 []slack.Channel
	var ret1 error
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].([]slack.Channel)
		}
		if result[1] != nil {
			ret1 = result[1].(error)
		}
	}
	return ret0, ret1
}

func (mock *MockUnderlyingSlackClient) PostMessage(channel string, text string, parameters slack.PostMessageParameters) (string, string, error) {
	params := []pegomock.Param{channel, text, parameters}
	result := pegomock.GetGenericMockFrom(mock).Invoke("PostMessage", params, []reflect.Type{reflect.TypeOf((*string)(nil)).Elem(), reflect.TypeOf((*string)(nil)).Elem(), reflect.TypeOf((*error)(nil)).Elem()})
	var ret0 string
	var ret1 string
	var ret2 error
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(string)
		}
		if result[1] != nil {
			ret1 = result[1].(string)
		}
		if result[2] != nil {
			ret2 = result[2].(error)
		}
	}
	return ret0, ret1, ret2
}

func (mock *MockUnderlyingSlackClient) VerifyWasCalledOnce() *VerifierUnderlyingSlackClient {
	return &VerifierUnderlyingSlackClient{mock, pegomock.Times(1), nil}
}

func (mock *MockUnderlyingSlackClient) VerifyWasCalled(invocationCountMatcher pegomock.Matcher) *VerifierUnderlyingSlackClient {
	return &VerifierUnderlyingSlackClient{mock, invocationCountMatcher, nil}
}

func (mock *MockUnderlyingSlackClient) VerifyWasCalledInOrder(invocationCountMatcher pegomock.Matcher, inOrderContext *pegomock.InOrderContext) *VerifierUnderlyingSlackClient {
	return &VerifierUnderlyingSlackClient{mock, invocationCountMatcher, inOrderContext}
}

type VerifierUnderlyingSlackClient struct {
	mock                   *MockUnderlyingSlackClient
	invocationCountMatcher pegomock.Matcher
	inOrderContext         *pegomock.InOrderContext
}

func (verifier *VerifierUnderlyingSlackClient) AuthTest() *UnderlyingSlackClient_AuthTest_OngoingVerification {
	params := []pegomock.Param{}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "AuthTest", params)
	return &UnderlyingSlackClient_AuthTest_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type UnderlyingSlackClient_AuthTest_OngoingVerification struct {
	mock              *MockUnderlyingSlackClient
	methodInvocations []pegomock.MethodInvocation
}

func (c *UnderlyingSlackClient_AuthTest_OngoingVerification) GetCapturedArguments() {
}

func (c *UnderlyingSlackClient_AuthTest_OngoingVerification) GetAllCapturedArguments() {
}

func (verifier *VerifierUnderlyingSlackClient) GetChannels(excludeArchived bool) *UnderlyingSlackClient_GetChannels_OngoingVerification {
	params := []pegomock.Param{excludeArchived}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "GetChannels", params)
	return &UnderlyingSlackClient_GetChannels_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type UnderlyingSlackClient_GetChannels_OngoingVerification struct {
	mock              *MockUnderlyingSlackClient
	methodInvocations []pegomock.MethodInvocation
}

func (c *UnderlyingSlackClient_GetChannels_OngoingVerification) GetCapturedArguments() bool {
	excludeArchived := c.GetAllCapturedArguments()
	return excludeArchived[len(excludeArchived)-1]
}

func (c *UnderlyingSlackClient_GetChannels_OngoingVerification) GetAllCapturedArguments() (_param0 []bool) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([]bool, len(params[0]))
		for u, param := range params[0] {
			_param0[u] = param.(bool)
		}
	}
	return
}

func (verifier *VerifierUnderlyingSlackClient) PostMessage(channel string, text string, parameters slack.PostMessageParameters) *UnderlyingSlackClient_PostMessage_OngoingVerification {
	params := []pegomock.Param{channel, text, parameters}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "PostMessage", params)
	return &UnderlyingSlackClient_PostMessage_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type UnderlyingSlackClient_PostMessage_OngoingVerification struct {
	mock              *MockUnderlyingSlackClient
	methodInvocations []pegomock.MethodInvocation
}

func (c *UnderlyingSlackClient_PostMessage_OngoingVerification) GetCapturedArguments() (string, string, slack.PostMessageParameters) {
	channel, text, parameters := c.GetAllCapturedArguments()
	return channel[len(channel)-1], text[len(text)-1], parameters[len(parameters)-1]
}

func (c *UnderlyingSlackClient_PostMessage_OngoingVerification) GetAllCapturedArguments() (_param0 []string, _param1 []string, _param2 []slack.PostMessageParameters) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([]string, len(params[0]))
		for u, param := range params[0] {
			_param0[u] = param.(string)
		}
		_param1 = make([]string, len(params[1]))
		for u, param := range params[1] {
			_param1[u] = param.(string)
		}
		_param2 = make([]slack.PostMessageParameters, len(params[2]))
		for u, param := range params[2] {
			_param2[u] = param.(slack.PostMessageParameters)
		}
	}
	return
}
