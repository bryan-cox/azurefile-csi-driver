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
// Source: factory.go
//
// Generated by this command:
//
//	mockgen -package mock_azclient -source factory.go -copyright_file ../../hack/boilerplate/boilerplate.generatego.txt
//

// Package mock_azclient is a generated GoMock package.
package mock_azclient

import (
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"

	accountclient "sigs.k8s.io/cloud-provider-azure/pkg/azclient/accountclient"
	availabilitysetclient "sigs.k8s.io/cloud-provider-azure/pkg/azclient/availabilitysetclient"
	backendaddresspoolclient "sigs.k8s.io/cloud-provider-azure/pkg/azclient/backendaddresspoolclient"
	blobcontainerclient "sigs.k8s.io/cloud-provider-azure/pkg/azclient/blobcontainerclient"
	blobservicepropertiesclient "sigs.k8s.io/cloud-provider-azure/pkg/azclient/blobservicepropertiesclient"
	deploymentclient "sigs.k8s.io/cloud-provider-azure/pkg/azclient/deploymentclient"
	diskclient "sigs.k8s.io/cloud-provider-azure/pkg/azclient/diskclient"
	fileservicepropertiesclient "sigs.k8s.io/cloud-provider-azure/pkg/azclient/fileservicepropertiesclient"
	fileshareclient "sigs.k8s.io/cloud-provider-azure/pkg/azclient/fileshareclient"
	identityclient "sigs.k8s.io/cloud-provider-azure/pkg/azclient/identityclient"
	interfaceclient "sigs.k8s.io/cloud-provider-azure/pkg/azclient/interfaceclient"
	ipgroupclient "sigs.k8s.io/cloud-provider-azure/pkg/azclient/ipgroupclient"
	loadbalancerclient "sigs.k8s.io/cloud-provider-azure/pkg/azclient/loadbalancerclient"
	managedclusterclient "sigs.k8s.io/cloud-provider-azure/pkg/azclient/managedclusterclient"
	privatednszonegroupclient "sigs.k8s.io/cloud-provider-azure/pkg/azclient/privatednszonegroupclient"
	privateendpointclient "sigs.k8s.io/cloud-provider-azure/pkg/azclient/privateendpointclient"
	privatelinkserviceclient "sigs.k8s.io/cloud-provider-azure/pkg/azclient/privatelinkserviceclient"
	privatezoneclient "sigs.k8s.io/cloud-provider-azure/pkg/azclient/privatezoneclient"
	providerclient "sigs.k8s.io/cloud-provider-azure/pkg/azclient/providerclient"
	publicipaddressclient "sigs.k8s.io/cloud-provider-azure/pkg/azclient/publicipaddressclient"
	publicipprefixclient "sigs.k8s.io/cloud-provider-azure/pkg/azclient/publicipprefixclient"
	registryclient "sigs.k8s.io/cloud-provider-azure/pkg/azclient/registryclient"
	resourcegroupclient "sigs.k8s.io/cloud-provider-azure/pkg/azclient/resourcegroupclient"
	roleassignmentclient "sigs.k8s.io/cloud-provider-azure/pkg/azclient/roleassignmentclient"
	routetableclient "sigs.k8s.io/cloud-provider-azure/pkg/azclient/routetableclient"
	secretclient "sigs.k8s.io/cloud-provider-azure/pkg/azclient/secretclient"
	securitygroupclient "sigs.k8s.io/cloud-provider-azure/pkg/azclient/securitygroupclient"
	snapshotclient "sigs.k8s.io/cloud-provider-azure/pkg/azclient/snapshotclient"
	sshpublickeyresourceclient "sigs.k8s.io/cloud-provider-azure/pkg/azclient/sshpublickeyresourceclient"
	subnetclient "sigs.k8s.io/cloud-provider-azure/pkg/azclient/subnetclient"
	vaultclient "sigs.k8s.io/cloud-provider-azure/pkg/azclient/vaultclient"
	virtualmachineclient "sigs.k8s.io/cloud-provider-azure/pkg/azclient/virtualmachineclient"
	virtualmachinescalesetclient "sigs.k8s.io/cloud-provider-azure/pkg/azclient/virtualmachinescalesetclient"
	virtualmachinescalesetvmclient "sigs.k8s.io/cloud-provider-azure/pkg/azclient/virtualmachinescalesetvmclient"
	virtualnetworkclient "sigs.k8s.io/cloud-provider-azure/pkg/azclient/virtualnetworkclient"
	virtualnetworklinkclient "sigs.k8s.io/cloud-provider-azure/pkg/azclient/virtualnetworklinkclient"
)

// MockClientFactory is a mock of ClientFactory interface.
type MockClientFactory struct {
	ctrl     *gomock.Controller
	recorder *MockClientFactoryMockRecorder
	isgomock struct{}
}

// MockClientFactoryMockRecorder is the mock recorder for MockClientFactory.
type MockClientFactoryMockRecorder struct {
	mock *MockClientFactory
}

// NewMockClientFactory creates a new mock instance.
func NewMockClientFactory(ctrl *gomock.Controller) *MockClientFactory {
	mock := &MockClientFactory{ctrl: ctrl}
	mock.recorder = &MockClientFactoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockClientFactory) EXPECT() *MockClientFactoryMockRecorder {
	return m.recorder
}

// GetAccountClient mocks base method.
func (m *MockClientFactory) GetAccountClient() accountclient.Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAccountClient")
	ret0, _ := ret[0].(accountclient.Interface)
	return ret0
}

// GetAccountClient indicates an expected call of GetAccountClient.
func (mr *MockClientFactoryMockRecorder) GetAccountClient() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAccountClient", reflect.TypeOf((*MockClientFactory)(nil).GetAccountClient))
}

// GetAccountClientForSub mocks base method.
func (m *MockClientFactory) GetAccountClientForSub(subscriptionID string) (accountclient.Interface, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAccountClientForSub", subscriptionID)
	ret0, _ := ret[0].(accountclient.Interface)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAccountClientForSub indicates an expected call of GetAccountClientForSub.
func (mr *MockClientFactoryMockRecorder) GetAccountClientForSub(subscriptionID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAccountClientForSub", reflect.TypeOf((*MockClientFactory)(nil).GetAccountClientForSub), subscriptionID)
}

// GetAvailabilitySetClient mocks base method.
func (m *MockClientFactory) GetAvailabilitySetClient() availabilitysetclient.Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAvailabilitySetClient")
	ret0, _ := ret[0].(availabilitysetclient.Interface)
	return ret0
}

// GetAvailabilitySetClient indicates an expected call of GetAvailabilitySetClient.
func (mr *MockClientFactoryMockRecorder) GetAvailabilitySetClient() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAvailabilitySetClient", reflect.TypeOf((*MockClientFactory)(nil).GetAvailabilitySetClient))
}

// GetBackendAddressPoolClient mocks base method.
func (m *MockClientFactory) GetBackendAddressPoolClient() backendaddresspoolclient.Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBackendAddressPoolClient")
	ret0, _ := ret[0].(backendaddresspoolclient.Interface)
	return ret0
}

// GetBackendAddressPoolClient indicates an expected call of GetBackendAddressPoolClient.
func (mr *MockClientFactoryMockRecorder) GetBackendAddressPoolClient() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBackendAddressPoolClient", reflect.TypeOf((*MockClientFactory)(nil).GetBackendAddressPoolClient))
}

// GetBlobContainerClient mocks base method.
func (m *MockClientFactory) GetBlobContainerClient() blobcontainerclient.Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBlobContainerClient")
	ret0, _ := ret[0].(blobcontainerclient.Interface)
	return ret0
}

// GetBlobContainerClient indicates an expected call of GetBlobContainerClient.
func (mr *MockClientFactoryMockRecorder) GetBlobContainerClient() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBlobContainerClient", reflect.TypeOf((*MockClientFactory)(nil).GetBlobContainerClient))
}

// GetBlobContainerClientForSub mocks base method.
func (m *MockClientFactory) GetBlobContainerClientForSub(subscriptionID string) (blobcontainerclient.Interface, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBlobContainerClientForSub", subscriptionID)
	ret0, _ := ret[0].(blobcontainerclient.Interface)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBlobContainerClientForSub indicates an expected call of GetBlobContainerClientForSub.
func (mr *MockClientFactoryMockRecorder) GetBlobContainerClientForSub(subscriptionID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBlobContainerClientForSub", reflect.TypeOf((*MockClientFactory)(nil).GetBlobContainerClientForSub), subscriptionID)
}

// GetBlobServicePropertiesClient mocks base method.
func (m *MockClientFactory) GetBlobServicePropertiesClient() blobservicepropertiesclient.Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBlobServicePropertiesClient")
	ret0, _ := ret[0].(blobservicepropertiesclient.Interface)
	return ret0
}

// GetBlobServicePropertiesClient indicates an expected call of GetBlobServicePropertiesClient.
func (mr *MockClientFactoryMockRecorder) GetBlobServicePropertiesClient() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBlobServicePropertiesClient", reflect.TypeOf((*MockClientFactory)(nil).GetBlobServicePropertiesClient))
}

// GetBlobServicePropertiesClientForSub mocks base method.
func (m *MockClientFactory) GetBlobServicePropertiesClientForSub(subscriptionID string) (blobservicepropertiesclient.Interface, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBlobServicePropertiesClientForSub", subscriptionID)
	ret0, _ := ret[0].(blobservicepropertiesclient.Interface)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBlobServicePropertiesClientForSub indicates an expected call of GetBlobServicePropertiesClientForSub.
func (mr *MockClientFactoryMockRecorder) GetBlobServicePropertiesClientForSub(subscriptionID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBlobServicePropertiesClientForSub", reflect.TypeOf((*MockClientFactory)(nil).GetBlobServicePropertiesClientForSub), subscriptionID)
}

// GetDeploymentClient mocks base method.
func (m *MockClientFactory) GetDeploymentClient() deploymentclient.Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDeploymentClient")
	ret0, _ := ret[0].(deploymentclient.Interface)
	return ret0
}

// GetDeploymentClient indicates an expected call of GetDeploymentClient.
func (mr *MockClientFactoryMockRecorder) GetDeploymentClient() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDeploymentClient", reflect.TypeOf((*MockClientFactory)(nil).GetDeploymentClient))
}

// GetDiskClient mocks base method.
func (m *MockClientFactory) GetDiskClient() diskclient.Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDiskClient")
	ret0, _ := ret[0].(diskclient.Interface)
	return ret0
}

// GetDiskClient indicates an expected call of GetDiskClient.
func (mr *MockClientFactoryMockRecorder) GetDiskClient() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDiskClient", reflect.TypeOf((*MockClientFactory)(nil).GetDiskClient))
}

// GetDiskClientForSub mocks base method.
func (m *MockClientFactory) GetDiskClientForSub(subscriptionID string) (diskclient.Interface, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDiskClientForSub", subscriptionID)
	ret0, _ := ret[0].(diskclient.Interface)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDiskClientForSub indicates an expected call of GetDiskClientForSub.
func (mr *MockClientFactoryMockRecorder) GetDiskClientForSub(subscriptionID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDiskClientForSub", reflect.TypeOf((*MockClientFactory)(nil).GetDiskClientForSub), subscriptionID)
}

// GetFileServicePropertiesClient mocks base method.
func (m *MockClientFactory) GetFileServicePropertiesClient() fileservicepropertiesclient.Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFileServicePropertiesClient")
	ret0, _ := ret[0].(fileservicepropertiesclient.Interface)
	return ret0
}

// GetFileServicePropertiesClient indicates an expected call of GetFileServicePropertiesClient.
func (mr *MockClientFactoryMockRecorder) GetFileServicePropertiesClient() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFileServicePropertiesClient", reflect.TypeOf((*MockClientFactory)(nil).GetFileServicePropertiesClient))
}

// GetFileServicePropertiesClientForSub mocks base method.
func (m *MockClientFactory) GetFileServicePropertiesClientForSub(subscriptionID string) (fileservicepropertiesclient.Interface, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFileServicePropertiesClientForSub", subscriptionID)
	ret0, _ := ret[0].(fileservicepropertiesclient.Interface)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFileServicePropertiesClientForSub indicates an expected call of GetFileServicePropertiesClientForSub.
func (mr *MockClientFactoryMockRecorder) GetFileServicePropertiesClientForSub(subscriptionID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFileServicePropertiesClientForSub", reflect.TypeOf((*MockClientFactory)(nil).GetFileServicePropertiesClientForSub), subscriptionID)
}

// GetFileShareClient mocks base method.
func (m *MockClientFactory) GetFileShareClient() fileshareclient.Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFileShareClient")
	ret0, _ := ret[0].(fileshareclient.Interface)
	return ret0
}

// GetFileShareClient indicates an expected call of GetFileShareClient.
func (mr *MockClientFactoryMockRecorder) GetFileShareClient() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFileShareClient", reflect.TypeOf((*MockClientFactory)(nil).GetFileShareClient))
}

// GetFileShareClientForSub mocks base method.
func (m *MockClientFactory) GetFileShareClientForSub(subscriptionID string) (fileshareclient.Interface, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFileShareClientForSub", subscriptionID)
	ret0, _ := ret[0].(fileshareclient.Interface)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFileShareClientForSub indicates an expected call of GetFileShareClientForSub.
func (mr *MockClientFactoryMockRecorder) GetFileShareClientForSub(subscriptionID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFileShareClientForSub", reflect.TypeOf((*MockClientFactory)(nil).GetFileShareClientForSub), subscriptionID)
}

// GetIPGroupClient mocks base method.
func (m *MockClientFactory) GetIPGroupClient() ipgroupclient.Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetIPGroupClient")
	ret0, _ := ret[0].(ipgroupclient.Interface)
	return ret0
}

// GetIPGroupClient indicates an expected call of GetIPGroupClient.
func (mr *MockClientFactoryMockRecorder) GetIPGroupClient() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetIPGroupClient", reflect.TypeOf((*MockClientFactory)(nil).GetIPGroupClient))
}

// GetIdentityClient mocks base method.
func (m *MockClientFactory) GetIdentityClient() identityclient.Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetIdentityClient")
	ret0, _ := ret[0].(identityclient.Interface)
	return ret0
}

// GetIdentityClient indicates an expected call of GetIdentityClient.
func (mr *MockClientFactoryMockRecorder) GetIdentityClient() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetIdentityClient", reflect.TypeOf((*MockClientFactory)(nil).GetIdentityClient))
}

// GetInterfaceClient mocks base method.
func (m *MockClientFactory) GetInterfaceClient() interfaceclient.Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetInterfaceClient")
	ret0, _ := ret[0].(interfaceclient.Interface)
	return ret0
}

// GetInterfaceClient indicates an expected call of GetInterfaceClient.
func (mr *MockClientFactoryMockRecorder) GetInterfaceClient() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetInterfaceClient", reflect.TypeOf((*MockClientFactory)(nil).GetInterfaceClient))
}

// GetLoadBalancerClient mocks base method.
func (m *MockClientFactory) GetLoadBalancerClient() loadbalancerclient.Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLoadBalancerClient")
	ret0, _ := ret[0].(loadbalancerclient.Interface)
	return ret0
}

// GetLoadBalancerClient indicates an expected call of GetLoadBalancerClient.
func (mr *MockClientFactoryMockRecorder) GetLoadBalancerClient() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLoadBalancerClient", reflect.TypeOf((*MockClientFactory)(nil).GetLoadBalancerClient))
}

// GetManagedClusterClient mocks base method.
func (m *MockClientFactory) GetManagedClusterClient() managedclusterclient.Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetManagedClusterClient")
	ret0, _ := ret[0].(managedclusterclient.Interface)
	return ret0
}

// GetManagedClusterClient indicates an expected call of GetManagedClusterClient.
func (mr *MockClientFactoryMockRecorder) GetManagedClusterClient() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetManagedClusterClient", reflect.TypeOf((*MockClientFactory)(nil).GetManagedClusterClient))
}

// GetPrivateDNSZoneGroupClient mocks base method.
func (m *MockClientFactory) GetPrivateDNSZoneGroupClient() privatednszonegroupclient.Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPrivateDNSZoneGroupClient")
	ret0, _ := ret[0].(privatednszonegroupclient.Interface)
	return ret0
}

// GetPrivateDNSZoneGroupClient indicates an expected call of GetPrivateDNSZoneGroupClient.
func (mr *MockClientFactoryMockRecorder) GetPrivateDNSZoneGroupClient() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPrivateDNSZoneGroupClient", reflect.TypeOf((*MockClientFactory)(nil).GetPrivateDNSZoneGroupClient))
}

// GetPrivateEndpointClient mocks base method.
func (m *MockClientFactory) GetPrivateEndpointClient() privateendpointclient.Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPrivateEndpointClient")
	ret0, _ := ret[0].(privateendpointclient.Interface)
	return ret0
}

// GetPrivateEndpointClient indicates an expected call of GetPrivateEndpointClient.
func (mr *MockClientFactoryMockRecorder) GetPrivateEndpointClient() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPrivateEndpointClient", reflect.TypeOf((*MockClientFactory)(nil).GetPrivateEndpointClient))
}

// GetPrivateLinkServiceClient mocks base method.
func (m *MockClientFactory) GetPrivateLinkServiceClient() privatelinkserviceclient.Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPrivateLinkServiceClient")
	ret0, _ := ret[0].(privatelinkserviceclient.Interface)
	return ret0
}

// GetPrivateLinkServiceClient indicates an expected call of GetPrivateLinkServiceClient.
func (mr *MockClientFactoryMockRecorder) GetPrivateLinkServiceClient() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPrivateLinkServiceClient", reflect.TypeOf((*MockClientFactory)(nil).GetPrivateLinkServiceClient))
}

// GetPrivateZoneClient mocks base method.
func (m *MockClientFactory) GetPrivateZoneClient() privatezoneclient.Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPrivateZoneClient")
	ret0, _ := ret[0].(privatezoneclient.Interface)
	return ret0
}

// GetPrivateZoneClient indicates an expected call of GetPrivateZoneClient.
func (mr *MockClientFactoryMockRecorder) GetPrivateZoneClient() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPrivateZoneClient", reflect.TypeOf((*MockClientFactory)(nil).GetPrivateZoneClient))
}

// GetProviderClient mocks base method.
func (m *MockClientFactory) GetProviderClient() providerclient.Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProviderClient")
	ret0, _ := ret[0].(providerclient.Interface)
	return ret0
}

// GetProviderClient indicates an expected call of GetProviderClient.
func (mr *MockClientFactoryMockRecorder) GetProviderClient() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProviderClient", reflect.TypeOf((*MockClientFactory)(nil).GetProviderClient))
}

// GetPublicIPAddressClient mocks base method.
func (m *MockClientFactory) GetPublicIPAddressClient() publicipaddressclient.Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPublicIPAddressClient")
	ret0, _ := ret[0].(publicipaddressclient.Interface)
	return ret0
}

// GetPublicIPAddressClient indicates an expected call of GetPublicIPAddressClient.
func (mr *MockClientFactoryMockRecorder) GetPublicIPAddressClient() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPublicIPAddressClient", reflect.TypeOf((*MockClientFactory)(nil).GetPublicIPAddressClient))
}

// GetPublicIPPrefixClient mocks base method.
func (m *MockClientFactory) GetPublicIPPrefixClient() publicipprefixclient.Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPublicIPPrefixClient")
	ret0, _ := ret[0].(publicipprefixclient.Interface)
	return ret0
}

// GetPublicIPPrefixClient indicates an expected call of GetPublicIPPrefixClient.
func (mr *MockClientFactoryMockRecorder) GetPublicIPPrefixClient() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPublicIPPrefixClient", reflect.TypeOf((*MockClientFactory)(nil).GetPublicIPPrefixClient))
}

// GetRegistryClient mocks base method.
func (m *MockClientFactory) GetRegistryClient() registryclient.Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRegistryClient")
	ret0, _ := ret[0].(registryclient.Interface)
	return ret0
}

// GetRegistryClient indicates an expected call of GetRegistryClient.
func (mr *MockClientFactoryMockRecorder) GetRegistryClient() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRegistryClient", reflect.TypeOf((*MockClientFactory)(nil).GetRegistryClient))
}

// GetResourceGroupClient mocks base method.
func (m *MockClientFactory) GetResourceGroupClient() resourcegroupclient.Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetResourceGroupClient")
	ret0, _ := ret[0].(resourcegroupclient.Interface)
	return ret0
}

// GetResourceGroupClient indicates an expected call of GetResourceGroupClient.
func (mr *MockClientFactoryMockRecorder) GetResourceGroupClient() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetResourceGroupClient", reflect.TypeOf((*MockClientFactory)(nil).GetResourceGroupClient))
}

// GetRoleAssignmentClient mocks base method.
func (m *MockClientFactory) GetRoleAssignmentClient() roleassignmentclient.Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRoleAssignmentClient")
	ret0, _ := ret[0].(roleassignmentclient.Interface)
	return ret0
}

// GetRoleAssignmentClient indicates an expected call of GetRoleAssignmentClient.
func (mr *MockClientFactoryMockRecorder) GetRoleAssignmentClient() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRoleAssignmentClient", reflect.TypeOf((*MockClientFactory)(nil).GetRoleAssignmentClient))
}

// GetRouteTableClient mocks base method.
func (m *MockClientFactory) GetRouteTableClient() routetableclient.Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRouteTableClient")
	ret0, _ := ret[0].(routetableclient.Interface)
	return ret0
}

// GetRouteTableClient indicates an expected call of GetRouteTableClient.
func (mr *MockClientFactoryMockRecorder) GetRouteTableClient() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRouteTableClient", reflect.TypeOf((*MockClientFactory)(nil).GetRouteTableClient))
}

// GetSSHPublicKeyResourceClient mocks base method.
func (m *MockClientFactory) GetSSHPublicKeyResourceClient() sshpublickeyresourceclient.Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSSHPublicKeyResourceClient")
	ret0, _ := ret[0].(sshpublickeyresourceclient.Interface)
	return ret0
}

// GetSSHPublicKeyResourceClient indicates an expected call of GetSSHPublicKeyResourceClient.
func (mr *MockClientFactoryMockRecorder) GetSSHPublicKeyResourceClient() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSSHPublicKeyResourceClient", reflect.TypeOf((*MockClientFactory)(nil).GetSSHPublicKeyResourceClient))
}

// GetSecretClient mocks base method.
func (m *MockClientFactory) GetSecretClient() secretclient.Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSecretClient")
	ret0, _ := ret[0].(secretclient.Interface)
	return ret0
}

// GetSecretClient indicates an expected call of GetSecretClient.
func (mr *MockClientFactoryMockRecorder) GetSecretClient() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSecretClient", reflect.TypeOf((*MockClientFactory)(nil).GetSecretClient))
}

// GetSecurityGroupClient mocks base method.
func (m *MockClientFactory) GetSecurityGroupClient() securitygroupclient.Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSecurityGroupClient")
	ret0, _ := ret[0].(securitygroupclient.Interface)
	return ret0
}

// GetSecurityGroupClient indicates an expected call of GetSecurityGroupClient.
func (mr *MockClientFactoryMockRecorder) GetSecurityGroupClient() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSecurityGroupClient", reflect.TypeOf((*MockClientFactory)(nil).GetSecurityGroupClient))
}

// GetSnapshotClient mocks base method.
func (m *MockClientFactory) GetSnapshotClient() snapshotclient.Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSnapshotClient")
	ret0, _ := ret[0].(snapshotclient.Interface)
	return ret0
}

// GetSnapshotClient indicates an expected call of GetSnapshotClient.
func (mr *MockClientFactoryMockRecorder) GetSnapshotClient() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSnapshotClient", reflect.TypeOf((*MockClientFactory)(nil).GetSnapshotClient))
}

// GetSnapshotClientForSub mocks base method.
func (m *MockClientFactory) GetSnapshotClientForSub(subscriptionID string) (snapshotclient.Interface, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSnapshotClientForSub", subscriptionID)
	ret0, _ := ret[0].(snapshotclient.Interface)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSnapshotClientForSub indicates an expected call of GetSnapshotClientForSub.
func (mr *MockClientFactoryMockRecorder) GetSnapshotClientForSub(subscriptionID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSnapshotClientForSub", reflect.TypeOf((*MockClientFactory)(nil).GetSnapshotClientForSub), subscriptionID)
}

// GetSubnetClient mocks base method.
func (m *MockClientFactory) GetSubnetClient() subnetclient.Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSubnetClient")
	ret0, _ := ret[0].(subnetclient.Interface)
	return ret0
}

// GetSubnetClient indicates an expected call of GetSubnetClient.
func (mr *MockClientFactoryMockRecorder) GetSubnetClient() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSubnetClient", reflect.TypeOf((*MockClientFactory)(nil).GetSubnetClient))
}

// GetVaultClient mocks base method.
func (m *MockClientFactory) GetVaultClient() vaultclient.Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetVaultClient")
	ret0, _ := ret[0].(vaultclient.Interface)
	return ret0
}

// GetVaultClient indicates an expected call of GetVaultClient.
func (mr *MockClientFactoryMockRecorder) GetVaultClient() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetVaultClient", reflect.TypeOf((*MockClientFactory)(nil).GetVaultClient))
}

// GetVirtualMachineClient mocks base method.
func (m *MockClientFactory) GetVirtualMachineClient() virtualmachineclient.Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetVirtualMachineClient")
	ret0, _ := ret[0].(virtualmachineclient.Interface)
	return ret0
}

// GetVirtualMachineClient indicates an expected call of GetVirtualMachineClient.
func (mr *MockClientFactoryMockRecorder) GetVirtualMachineClient() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetVirtualMachineClient", reflect.TypeOf((*MockClientFactory)(nil).GetVirtualMachineClient))
}

// GetVirtualMachineScaleSetClient mocks base method.
func (m *MockClientFactory) GetVirtualMachineScaleSetClient() virtualmachinescalesetclient.Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetVirtualMachineScaleSetClient")
	ret0, _ := ret[0].(virtualmachinescalesetclient.Interface)
	return ret0
}

// GetVirtualMachineScaleSetClient indicates an expected call of GetVirtualMachineScaleSetClient.
func (mr *MockClientFactoryMockRecorder) GetVirtualMachineScaleSetClient() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetVirtualMachineScaleSetClient", reflect.TypeOf((*MockClientFactory)(nil).GetVirtualMachineScaleSetClient))
}

// GetVirtualMachineScaleSetVMClient mocks base method.
func (m *MockClientFactory) GetVirtualMachineScaleSetVMClient() virtualmachinescalesetvmclient.Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetVirtualMachineScaleSetVMClient")
	ret0, _ := ret[0].(virtualmachinescalesetvmclient.Interface)
	return ret0
}

// GetVirtualMachineScaleSetVMClient indicates an expected call of GetVirtualMachineScaleSetVMClient.
func (mr *MockClientFactoryMockRecorder) GetVirtualMachineScaleSetVMClient() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetVirtualMachineScaleSetVMClient", reflect.TypeOf((*MockClientFactory)(nil).GetVirtualMachineScaleSetVMClient))
}

// GetVirtualNetworkClient mocks base method.
func (m *MockClientFactory) GetVirtualNetworkClient() virtualnetworkclient.Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetVirtualNetworkClient")
	ret0, _ := ret[0].(virtualnetworkclient.Interface)
	return ret0
}

// GetVirtualNetworkClient indicates an expected call of GetVirtualNetworkClient.
func (mr *MockClientFactoryMockRecorder) GetVirtualNetworkClient() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetVirtualNetworkClient", reflect.TypeOf((*MockClientFactory)(nil).GetVirtualNetworkClient))
}

// GetVirtualNetworkLinkClient mocks base method.
func (m *MockClientFactory) GetVirtualNetworkLinkClient() virtualnetworklinkclient.Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetVirtualNetworkLinkClient")
	ret0, _ := ret[0].(virtualnetworklinkclient.Interface)
	return ret0
}

// GetVirtualNetworkLinkClient indicates an expected call of GetVirtualNetworkLinkClient.
func (mr *MockClientFactoryMockRecorder) GetVirtualNetworkLinkClient() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetVirtualNetworkLinkClient", reflect.TypeOf((*MockClientFactory)(nil).GetVirtualNetworkLinkClient))
}
