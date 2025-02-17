// /*
// Copyright The Kubernetes Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
// */
//

// Code generated by MockGen. DO NOT EDIT.
// Source: blobservicepropertiesclient/interface.go
//
// Generated by this command:
//
//	mockgen -package mock_blobservicepropertiesclient -source blobservicepropertiesclient/interface.go -typed -write_generate_directive -copyright_file ../../hack/boilerplate/boilerplate.generatego.txt
//

// Package mock_blobservicepropertiesclient is a generated GoMock package.
package mock_blobservicepropertiesclient

import (
	context "context"
	reflect "reflect"

	armstorage "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storage/armstorage"
	gomock "go.uber.org/mock/gomock"
)

//go:generate mockgen -package mock_blobservicepropertiesclient -source blobservicepropertiesclient/interface.go -typed -write_generate_directive -copyright_file ../../hack/boilerplate/boilerplate.generatego.txt

// MockInterface is a mock of Interface interface.
type MockInterface struct {
	ctrl     *gomock.Controller
	recorder *MockInterfaceMockRecorder
	isgomock struct{}
}

// MockInterfaceMockRecorder is the mock recorder for MockInterface.
type MockInterfaceMockRecorder struct {
	mock *MockInterface
}

// NewMockInterface creates a new mock instance.
func NewMockInterface(ctrl *gomock.Controller) *MockInterface {
	mock := &MockInterface{ctrl: ctrl}
	mock.recorder = &MockInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockInterface) EXPECT() *MockInterfaceMockRecorder {
	return m.recorder
}

// Get mocks base method.
func (m *MockInterface) Get(ctx context.Context, resourceGroupName, resourceName string) (*armstorage.BlobServiceProperties, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx, resourceGroupName, resourceName)
	ret0, _ := ret[0].(*armstorage.BlobServiceProperties)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockInterfaceMockRecorder) Get(ctx, resourceGroupName, resourceName any) *MockInterfaceGetCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockInterface)(nil).Get), ctx, resourceGroupName, resourceName)
	return &MockInterfaceGetCall{Call: call}
}

// MockInterfaceGetCall wrap *gomock.Call
type MockInterfaceGetCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockInterfaceGetCall) Return(arg0 *armstorage.BlobServiceProperties, arg1 error) *MockInterfaceGetCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockInterfaceGetCall) Do(f func(context.Context, string, string) (*armstorage.BlobServiceProperties, error)) *MockInterfaceGetCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockInterfaceGetCall) DoAndReturn(f func(context.Context, string, string) (*armstorage.BlobServiceProperties, error)) *MockInterfaceGetCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Set mocks base method.
func (m *MockInterface) Set(ctx context.Context, resourceGroupName, resourceName string, parameters armstorage.BlobServiceProperties) (*armstorage.BlobServiceProperties, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Set", ctx, resourceGroupName, resourceName, parameters)
	ret0, _ := ret[0].(*armstorage.BlobServiceProperties)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Set indicates an expected call of Set.
func (mr *MockInterfaceMockRecorder) Set(ctx, resourceGroupName, resourceName, parameters any) *MockInterfaceSetCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Set", reflect.TypeOf((*MockInterface)(nil).Set), ctx, resourceGroupName, resourceName, parameters)
	return &MockInterfaceSetCall{Call: call}
}

// MockInterfaceSetCall wrap *gomock.Call
type MockInterfaceSetCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockInterfaceSetCall) Return(arg0 *armstorage.BlobServiceProperties, arg1 error) *MockInterfaceSetCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockInterfaceSetCall) Do(f func(context.Context, string, string, armstorage.BlobServiceProperties) (*armstorage.BlobServiceProperties, error)) *MockInterfaceSetCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockInterfaceSetCall) DoAndReturn(f func(context.Context, string, string, armstorage.BlobServiceProperties) (*armstorage.BlobServiceProperties, error)) *MockInterfaceSetCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
