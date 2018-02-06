// Automatically generated by pegomock. DO NOT EDIT!
// Source: github.com/atlantisnorth/atlantis/server/events (interfaces: CommitStatusUpdater)

package mocks

import (
	"reflect"

	events "github.com/atlantisnorth/atlantis/server/events"
	models "github.com/atlantisnorth/atlantis/server/events/models"
	vcs "github.com/atlantisnorth/atlantis/server/events/vcs"
	pegomock "github.com/petergtz/pegomock"
)

type MockCommitStatusUpdater struct {
	fail func(message string, callerSkip ...int)
}

func NewMockCommitStatusUpdater() *MockCommitStatusUpdater {
	return &MockCommitStatusUpdater{fail: pegomock.GlobalFailHandler}
}

func (mock *MockCommitStatusUpdater) Update(repo models.Repo, pull models.PullRequest, status vcs.CommitStatus, cmd *events.Command, host vcs.Host) error {
	params := []pegomock.Param{repo, pull, status, cmd, host}
	result := pegomock.GetGenericMockFrom(mock).Invoke("Update", params, []reflect.Type{reflect.TypeOf((*error)(nil)).Elem()})
	var ret0 error
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(error)
		}
	}
	return ret0
}

func (mock *MockCommitStatusUpdater) UpdateProjectResult(ctx *events.CommandContext, res events.CommandResponse) error {
	params := []pegomock.Param{ctx, res}
	result := pegomock.GetGenericMockFrom(mock).Invoke("UpdateProjectResult", params, []reflect.Type{reflect.TypeOf((*error)(nil)).Elem()})
	var ret0 error
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(error)
		}
	}
	return ret0
}

func (mock *MockCommitStatusUpdater) VerifyWasCalledOnce() *VerifierCommitStatusUpdater {
	return &VerifierCommitStatusUpdater{mock, pegomock.Times(1), nil}
}

func (mock *MockCommitStatusUpdater) VerifyWasCalled(invocationCountMatcher pegomock.Matcher) *VerifierCommitStatusUpdater {
	return &VerifierCommitStatusUpdater{mock, invocationCountMatcher, nil}
}

func (mock *MockCommitStatusUpdater) VerifyWasCalledInOrder(invocationCountMatcher pegomock.Matcher, inOrderContext *pegomock.InOrderContext) *VerifierCommitStatusUpdater {
	return &VerifierCommitStatusUpdater{mock, invocationCountMatcher, inOrderContext}
}

type VerifierCommitStatusUpdater struct {
	mock                   *MockCommitStatusUpdater
	invocationCountMatcher pegomock.Matcher
	inOrderContext         *pegomock.InOrderContext
}

func (verifier *VerifierCommitStatusUpdater) Update(repo models.Repo, pull models.PullRequest, status vcs.CommitStatus, cmd *events.Command, host vcs.Host) *CommitStatusUpdater_Update_OngoingVerification {
	params := []pegomock.Param{repo, pull, status, cmd, host}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "Update", params)
	return &CommitStatusUpdater_Update_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type CommitStatusUpdater_Update_OngoingVerification struct {
	mock              *MockCommitStatusUpdater
	methodInvocations []pegomock.MethodInvocation
}

func (c *CommitStatusUpdater_Update_OngoingVerification) GetCapturedArguments() (models.Repo, models.PullRequest, vcs.CommitStatus, *events.Command, vcs.Host) {
	repo, pull, status, cmd, host := c.GetAllCapturedArguments()
	return repo[len(repo)-1], pull[len(pull)-1], status[len(status)-1], cmd[len(cmd)-1], host[len(host)-1]
}

func (c *CommitStatusUpdater_Update_OngoingVerification) GetAllCapturedArguments() (_param0 []models.Repo, _param1 []models.PullRequest, _param2 []vcs.CommitStatus, _param3 []*events.Command, _param4 []vcs.Host) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([]models.Repo, len(params[0]))
		for u, param := range params[0] {
			_param0[u] = param.(models.Repo)
		}
		_param1 = make([]models.PullRequest, len(params[1]))
		for u, param := range params[1] {
			_param1[u] = param.(models.PullRequest)
		}
		_param2 = make([]vcs.CommitStatus, len(params[2]))
		for u, param := range params[2] {
			_param2[u] = param.(vcs.CommitStatus)
		}
		_param3 = make([]*events.Command, len(params[3]))
		for u, param := range params[3] {
			_param3[u] = param.(*events.Command)
		}
		_param4 = make([]vcs.Host, len(params[4]))
		for u, param := range params[4] {
			_param4[u] = param.(vcs.Host)
		}
	}
	return
}

func (verifier *VerifierCommitStatusUpdater) UpdateProjectResult(ctx *events.CommandContext, res events.CommandResponse) *CommitStatusUpdater_UpdateProjectResult_OngoingVerification {
	params := []pegomock.Param{ctx, res}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "UpdateProjectResult", params)
	return &CommitStatusUpdater_UpdateProjectResult_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type CommitStatusUpdater_UpdateProjectResult_OngoingVerification struct {
	mock              *MockCommitStatusUpdater
	methodInvocations []pegomock.MethodInvocation
}

func (c *CommitStatusUpdater_UpdateProjectResult_OngoingVerification) GetCapturedArguments() (*events.CommandContext, events.CommandResponse) {
	ctx, res := c.GetAllCapturedArguments()
	return ctx[len(ctx)-1], res[len(res)-1]
}

func (c *CommitStatusUpdater_UpdateProjectResult_OngoingVerification) GetAllCapturedArguments() (_param0 []*events.CommandContext, _param1 []events.CommandResponse) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([]*events.CommandContext, len(params[0]))
		for u, param := range params[0] {
			_param0[u] = param.(*events.CommandContext)
		}
		_param1 = make([]events.CommandResponse, len(params[1]))
		for u, param := range params[1] {
			_param1[u] = param.(events.CommandResponse)
		}
	}
	return
}
