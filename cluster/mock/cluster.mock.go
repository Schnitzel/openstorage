// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/libopenstorage/openstorage/cluster (interfaces: Cluster)

// Package mock is a generated GoMock package.
package mock

import (
	gomock "github.com/golang/mock/gomock"
	api "github.com/libopenstorage/openstorage/api"
	cluster "github.com/libopenstorage/openstorage/cluster"
	osdconfig "github.com/libopenstorage/openstorage/osdconfig"
	schedpolicy "github.com/libopenstorage/openstorage/schedpolicy"
	reflect "reflect"
	time "time"
)

// MockCluster is a mock of Cluster interface
type MockCluster struct {
	ctrl     *gomock.Controller
	recorder *MockClusterMockRecorder
}

// MockClusterMockRecorder is the mock recorder for MockCluster
type MockClusterMockRecorder struct {
	mock *MockCluster
}

// NewMockCluster creates a new mock instance
func NewMockCluster(ctrl *gomock.Controller) *MockCluster {
	mock := &MockCluster{ctrl: ctrl}
	mock.recorder = &MockClusterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockCluster) EXPECT() *MockClusterMockRecorder {
	return m.recorder
}

// ActivateClusterDomain mocks base method
func (m *MockCluster) ActivateClusterDomain(arg0 *api.ActivateClusterDomainRequest) error {
	ret := m.ctrl.Call(m, "ActivateClusterDomain", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// ActivateClusterDomain indicates an expected call of ActivateClusterDomain
func (mr *MockClusterMockRecorder) ActivateClusterDomain(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ActivateClusterDomain", reflect.TypeOf((*MockCluster)(nil).ActivateClusterDomain), arg0)
}

// AddEventListener mocks base method
func (m *MockCluster) AddEventListener(arg0 cluster.ClusterListener) error {
	ret := m.ctrl.Call(m, "AddEventListener", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddEventListener indicates an expected call of AddEventListener
func (mr *MockClusterMockRecorder) AddEventListener(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddEventListener", reflect.TypeOf((*MockCluster)(nil).AddEventListener), arg0)
}

// CreatePair mocks base method
func (m *MockCluster) CreatePair(arg0 *api.ClusterPairCreateRequest) (*api.ClusterPairCreateResponse, error) {
	ret := m.ctrl.Call(m, "CreatePair", arg0)
	ret0, _ := ret[0].(*api.ClusterPairCreateResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreatePair indicates an expected call of CreatePair
func (mr *MockClusterMockRecorder) CreatePair(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreatePair", reflect.TypeOf((*MockCluster)(nil).CreatePair), arg0)
}

// DeactivateClusterDomain mocks base method
func (m *MockCluster) DeactivateClusterDomain(arg0 *api.DeactivateClusterDomainRequest) error {
	ret := m.ctrl.Call(m, "DeactivateClusterDomain", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeactivateClusterDomain indicates an expected call of DeactivateClusterDomain
func (mr *MockClusterMockRecorder) DeactivateClusterDomain(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeactivateClusterDomain", reflect.TypeOf((*MockCluster)(nil).DeactivateClusterDomain), arg0)
}

// DeleteNodeConf mocks base method
func (m *MockCluster) DeleteNodeConf(arg0 string) error {
	ret := m.ctrl.Call(m, "DeleteNodeConf", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteNodeConf indicates an expected call of DeleteNodeConf
func (mr *MockClusterMockRecorder) DeleteNodeConf(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteNodeConf", reflect.TypeOf((*MockCluster)(nil).DeleteNodeConf), arg0)
}

// DeletePair mocks base method
func (m *MockCluster) DeletePair(arg0 string) error {
	ret := m.ctrl.Call(m, "DeletePair", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeletePair indicates an expected call of DeletePair
func (mr *MockClusterMockRecorder) DeletePair(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeletePair", reflect.TypeOf((*MockCluster)(nil).DeletePair), arg0)
}

// DisableUpdates mocks base method
func (m *MockCluster) DisableUpdates() error {
	ret := m.ctrl.Call(m, "DisableUpdates")
	ret0, _ := ret[0].(error)
	return ret0
}

// DisableUpdates indicates an expected call of DisableUpdates
func (mr *MockClusterMockRecorder) DisableUpdates() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DisableUpdates", reflect.TypeOf((*MockCluster)(nil).DisableUpdates))
}

// EnableUpdates mocks base method
func (m *MockCluster) EnableUpdates() error {
	ret := m.ctrl.Call(m, "EnableUpdates")
	ret0, _ := ret[0].(error)
	return ret0
}

// EnableUpdates indicates an expected call of EnableUpdates
func (mr *MockClusterMockRecorder) EnableUpdates() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnableUpdates", reflect.TypeOf((*MockCluster)(nil).EnableUpdates))
}

// Enumerate mocks base method
func (m *MockCluster) Enumerate() (api.Cluster, error) {
	ret := m.ctrl.Call(m, "Enumerate")
	ret0, _ := ret[0].(api.Cluster)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Enumerate indicates an expected call of Enumerate
func (mr *MockClusterMockRecorder) Enumerate() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Enumerate", reflect.TypeOf((*MockCluster)(nil).Enumerate))
}

// EnumerateAlerts mocks base method
func (m *MockCluster) EnumerateAlerts(arg0, arg1 time.Time, arg2 api.ResourceType) (*api.Alerts, error) {
	ret := m.ctrl.Call(m, "EnumerateAlerts", arg0, arg1, arg2)
	ret0, _ := ret[0].(*api.Alerts)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EnumerateAlerts indicates an expected call of EnumerateAlerts
func (mr *MockClusterMockRecorder) EnumerateAlerts(arg0, arg1, arg2 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnumerateAlerts", reflect.TypeOf((*MockCluster)(nil).EnumerateAlerts), arg0, arg1, arg2)
}

// EnumerateNodeConf mocks base method
func (m *MockCluster) EnumerateNodeConf() (*osdconfig.NodesConfig, error) {
	ret := m.ctrl.Call(m, "EnumerateNodeConf")
	ret0, _ := ret[0].(*osdconfig.NodesConfig)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EnumerateNodeConf indicates an expected call of EnumerateNodeConf
func (mr *MockClusterMockRecorder) EnumerateNodeConf() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnumerateNodeConf", reflect.TypeOf((*MockCluster)(nil).EnumerateNodeConf))
}

// EnumeratePairs mocks base method
func (m *MockCluster) EnumeratePairs() (*api.ClusterPairsEnumerateResponse, error) {
	ret := m.ctrl.Call(m, "EnumeratePairs")
	ret0, _ := ret[0].(*api.ClusterPairsEnumerateResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EnumeratePairs indicates an expected call of EnumeratePairs
func (mr *MockClusterMockRecorder) EnumeratePairs() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnumeratePairs", reflect.TypeOf((*MockCluster)(nil).EnumeratePairs))
}

// EraseAlert mocks base method
func (m *MockCluster) EraseAlert(arg0 api.ResourceType, arg1 int64) error {
	ret := m.ctrl.Call(m, "EraseAlert", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// EraseAlert indicates an expected call of EraseAlert
func (mr *MockClusterMockRecorder) EraseAlert(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EraseAlert", reflect.TypeOf((*MockCluster)(nil).EraseAlert), arg0, arg1)
}

// GetClusterConf mocks base method
func (m *MockCluster) GetClusterConf() (*osdconfig.ClusterConfig, error) {
	ret := m.ctrl.Call(m, "GetClusterConf")
	ret0, _ := ret[0].(*osdconfig.ClusterConfig)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetClusterConf indicates an expected call of GetClusterConf
func (mr *MockClusterMockRecorder) GetClusterConf() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetClusterConf", reflect.TypeOf((*MockCluster)(nil).GetClusterConf))
}

// GetData mocks base method
func (m *MockCluster) GetData() (map[string]*api.Node, error) {
	ret := m.ctrl.Call(m, "GetData")
	ret0, _ := ret[0].(map[string]*api.Node)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetData indicates an expected call of GetData
func (mr *MockClusterMockRecorder) GetData() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetData", reflect.TypeOf((*MockCluster)(nil).GetData))
}

// GetGossipState mocks base method
func (m *MockCluster) GetGossipState() *cluster.ClusterState {
	ret := m.ctrl.Call(m, "GetGossipState")
	ret0, _ := ret[0].(*cluster.ClusterState)
	return ret0
}

// GetGossipState indicates an expected call of GetGossipState
func (mr *MockClusterMockRecorder) GetGossipState() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetGossipState", reflect.TypeOf((*MockCluster)(nil).GetGossipState))
}

// GetNodeConf mocks base method
func (m *MockCluster) GetNodeConf(arg0 string) (*osdconfig.NodeConfig, error) {
	ret := m.ctrl.Call(m, "GetNodeConf", arg0)
	ret0, _ := ret[0].(*osdconfig.NodeConfig)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetNodeConf indicates an expected call of GetNodeConf
func (mr *MockClusterMockRecorder) GetNodeConf(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNodeConf", reflect.TypeOf((*MockCluster)(nil).GetNodeConf), arg0)
}

// GetNodeIdFromIp mocks base method
func (m *MockCluster) GetNodeIdFromIp(arg0 string) (string, error) {
	ret := m.ctrl.Call(m, "GetNodeIdFromIp", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetNodeIdFromIp indicates an expected call of GetNodeIdFromIp
func (mr *MockClusterMockRecorder) GetNodeIdFromIp(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNodeIdFromIp", reflect.TypeOf((*MockCluster)(nil).GetNodeIdFromIp), arg0)
}

// GetPair mocks base method
func (m *MockCluster) GetPair(arg0 string) (*api.ClusterPairGetResponse, error) {
	ret := m.ctrl.Call(m, "GetPair", arg0)
	ret0, _ := ret[0].(*api.ClusterPairGetResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPair indicates an expected call of GetPair
func (mr *MockClusterMockRecorder) GetPair(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPair", reflect.TypeOf((*MockCluster)(nil).GetPair), arg0)
}

// GetPairToken mocks base method
func (m *MockCluster) GetPairToken(arg0 bool) (*api.ClusterPairTokenGetResponse, error) {
	ret := m.ctrl.Call(m, "GetPairToken", arg0)
	ret0, _ := ret[0].(*api.ClusterPairTokenGetResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPairToken indicates an expected call of GetPairToken
func (mr *MockClusterMockRecorder) GetPairToken(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPairToken", reflect.TypeOf((*MockCluster)(nil).GetPairToken), arg0)
}

// Inspect mocks base method
func (m *MockCluster) Inspect(arg0 string) (api.Node, error) {
	ret := m.ctrl.Call(m, "Inspect", arg0)
	ret0, _ := ret[0].(api.Node)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Inspect indicates an expected call of Inspect
func (mr *MockClusterMockRecorder) Inspect(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Inspect", reflect.TypeOf((*MockCluster)(nil).Inspect), arg0)
}

// NodeRemoveDone mocks base method
func (m *MockCluster) NodeRemoveDone(arg0 string, arg1 error) {
	m.ctrl.Call(m, "NodeRemoveDone", arg0, arg1)
}

// NodeRemoveDone indicates an expected call of NodeRemoveDone
func (mr *MockClusterMockRecorder) NodeRemoveDone(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NodeRemoveDone", reflect.TypeOf((*MockCluster)(nil).NodeRemoveDone), arg0, arg1)
}

// NodeStatus mocks base method
func (m *MockCluster) NodeStatus() (api.Status, error) {
	ret := m.ctrl.Call(m, "NodeStatus")
	ret0, _ := ret[0].(api.Status)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NodeStatus indicates an expected call of NodeStatus
func (mr *MockClusterMockRecorder) NodeStatus() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NodeStatus", reflect.TypeOf((*MockCluster)(nil).NodeStatus))
}

// ObjectStoreCreate mocks base method
func (m *MockCluster) ObjectStoreCreate(arg0 string) (*api.ObjectstoreInfo, error) {
	ret := m.ctrl.Call(m, "ObjectStoreCreate", arg0)
	ret0, _ := ret[0].(*api.ObjectstoreInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ObjectStoreCreate indicates an expected call of ObjectStoreCreate
func (mr *MockClusterMockRecorder) ObjectStoreCreate(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ObjectStoreCreate", reflect.TypeOf((*MockCluster)(nil).ObjectStoreCreate), arg0)
}

// ObjectStoreDelete mocks base method
func (m *MockCluster) ObjectStoreDelete(arg0 string) error {
	ret := m.ctrl.Call(m, "ObjectStoreDelete", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// ObjectStoreDelete indicates an expected call of ObjectStoreDelete
func (mr *MockClusterMockRecorder) ObjectStoreDelete(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ObjectStoreDelete", reflect.TypeOf((*MockCluster)(nil).ObjectStoreDelete), arg0)
}

// ObjectStoreInspect mocks base method
func (m *MockCluster) ObjectStoreInspect(arg0 string) (*api.ObjectstoreInfo, error) {
	ret := m.ctrl.Call(m, "ObjectStoreInspect", arg0)
	ret0, _ := ret[0].(*api.ObjectstoreInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ObjectStoreInspect indicates an expected call of ObjectStoreInspect
func (mr *MockClusterMockRecorder) ObjectStoreInspect(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ObjectStoreInspect", reflect.TypeOf((*MockCluster)(nil).ObjectStoreInspect), arg0)
}

// ObjectStoreUpdate mocks base method
func (m *MockCluster) ObjectStoreUpdate(arg0 string, arg1 bool) error {
	ret := m.ctrl.Call(m, "ObjectStoreUpdate", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// ObjectStoreUpdate indicates an expected call of ObjectStoreUpdate
func (mr *MockClusterMockRecorder) ObjectStoreUpdate(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ObjectStoreUpdate", reflect.TypeOf((*MockCluster)(nil).ObjectStoreUpdate), arg0, arg1)
}

// PeerStatus mocks base method
func (m *MockCluster) PeerStatus(arg0 string) (map[string]api.Status, error) {
	ret := m.ctrl.Call(m, "PeerStatus", arg0)
	ret0, _ := ret[0].(map[string]api.Status)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PeerStatus indicates an expected call of PeerStatus
func (mr *MockClusterMockRecorder) PeerStatus(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PeerStatus", reflect.TypeOf((*MockCluster)(nil).PeerStatus), arg0)
}

// ProcessPairRequest mocks base method
func (m *MockCluster) ProcessPairRequest(arg0 *api.ClusterPairProcessRequest) (*api.ClusterPairProcessResponse, error) {
	ret := m.ctrl.Call(m, "ProcessPairRequest", arg0)
	ret0, _ := ret[0].(*api.ClusterPairProcessResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ProcessPairRequest indicates an expected call of ProcessPairRequest
func (mr *MockClusterMockRecorder) ProcessPairRequest(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProcessPairRequest", reflect.TypeOf((*MockCluster)(nil).ProcessPairRequest), arg0)
}

// RefreshPair mocks base method
func (m *MockCluster) RefreshPair(arg0 string) error {
	ret := m.ctrl.Call(m, "RefreshPair", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// RefreshPair indicates an expected call of RefreshPair
func (mr *MockClusterMockRecorder) RefreshPair(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RefreshPair", reflect.TypeOf((*MockCluster)(nil).RefreshPair), arg0)
}

// Remove mocks base method
func (m *MockCluster) Remove(arg0 []api.Node, arg1 bool) error {
	ret := m.ctrl.Call(m, "Remove", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Remove indicates an expected call of Remove
func (mr *MockClusterMockRecorder) Remove(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Remove", reflect.TypeOf((*MockCluster)(nil).Remove), arg0, arg1)
}

// SchedPolicyCreate mocks base method
func (m *MockCluster) SchedPolicyCreate(arg0, arg1 string) error {
	ret := m.ctrl.Call(m, "SchedPolicyCreate", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// SchedPolicyCreate indicates an expected call of SchedPolicyCreate
func (mr *MockClusterMockRecorder) SchedPolicyCreate(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SchedPolicyCreate", reflect.TypeOf((*MockCluster)(nil).SchedPolicyCreate), arg0, arg1)
}

// SchedPolicyDelete mocks base method
func (m *MockCluster) SchedPolicyDelete(arg0 string) error {
	ret := m.ctrl.Call(m, "SchedPolicyDelete", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SchedPolicyDelete indicates an expected call of SchedPolicyDelete
func (mr *MockClusterMockRecorder) SchedPolicyDelete(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SchedPolicyDelete", reflect.TypeOf((*MockCluster)(nil).SchedPolicyDelete), arg0)
}

// SchedPolicyEnumerate mocks base method
func (m *MockCluster) SchedPolicyEnumerate() ([]*schedpolicy.SchedPolicy, error) {
	ret := m.ctrl.Call(m, "SchedPolicyEnumerate")
	ret0, _ := ret[0].([]*schedpolicy.SchedPolicy)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SchedPolicyEnumerate indicates an expected call of SchedPolicyEnumerate
func (mr *MockClusterMockRecorder) SchedPolicyEnumerate() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SchedPolicyEnumerate", reflect.TypeOf((*MockCluster)(nil).SchedPolicyEnumerate))
}

// SchedPolicyGet mocks base method
func (m *MockCluster) SchedPolicyGet(arg0 string) (*schedpolicy.SchedPolicy, error) {
	ret := m.ctrl.Call(m, "SchedPolicyGet", arg0)
	ret0, _ := ret[0].(*schedpolicy.SchedPolicy)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SchedPolicyGet indicates an expected call of SchedPolicyGet
func (mr *MockClusterMockRecorder) SchedPolicyGet(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SchedPolicyGet", reflect.TypeOf((*MockCluster)(nil).SchedPolicyGet), arg0)
}

// SchedPolicyUpdate mocks base method
func (m *MockCluster) SchedPolicyUpdate(arg0, arg1 string) error {
	ret := m.ctrl.Call(m, "SchedPolicyUpdate", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// SchedPolicyUpdate indicates an expected call of SchedPolicyUpdate
func (mr *MockClusterMockRecorder) SchedPolicyUpdate(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SchedPolicyUpdate", reflect.TypeOf((*MockCluster)(nil).SchedPolicyUpdate), arg0, arg1)
}

// SecretCheckLogin mocks base method
func (m *MockCluster) SecretCheckLogin() error {
	ret := m.ctrl.Call(m, "SecretCheckLogin")
	ret0, _ := ret[0].(error)
	return ret0
}

// SecretCheckLogin indicates an expected call of SecretCheckLogin
func (mr *MockClusterMockRecorder) SecretCheckLogin() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SecretCheckLogin", reflect.TypeOf((*MockCluster)(nil).SecretCheckLogin))
}

// SecretGet mocks base method
func (m *MockCluster) SecretGet(arg0 string) (interface{}, error) {
	ret := m.ctrl.Call(m, "SecretGet", arg0)
	ret0, _ := ret[0].(interface{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SecretGet indicates an expected call of SecretGet
func (mr *MockClusterMockRecorder) SecretGet(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SecretGet", reflect.TypeOf((*MockCluster)(nil).SecretGet), arg0)
}

// SecretGetDefaultSecretKey mocks base method
func (m *MockCluster) SecretGetDefaultSecretKey() (interface{}, error) {
	ret := m.ctrl.Call(m, "SecretGetDefaultSecretKey")
	ret0, _ := ret[0].(interface{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SecretGetDefaultSecretKey indicates an expected call of SecretGetDefaultSecretKey
func (mr *MockClusterMockRecorder) SecretGetDefaultSecretKey() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SecretGetDefaultSecretKey", reflect.TypeOf((*MockCluster)(nil).SecretGetDefaultSecretKey))
}

// SecretLogin mocks base method
func (m *MockCluster) SecretLogin(arg0 string, arg1 map[string]string) error {
	ret := m.ctrl.Call(m, "SecretLogin", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// SecretLogin indicates an expected call of SecretLogin
func (mr *MockClusterMockRecorder) SecretLogin(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SecretLogin", reflect.TypeOf((*MockCluster)(nil).SecretLogin), arg0, arg1)
}

// SecretSet mocks base method
func (m *MockCluster) SecretSet(arg0 string, arg1 interface{}) error {
	ret := m.ctrl.Call(m, "SecretSet", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// SecretSet indicates an expected call of SecretSet
func (mr *MockClusterMockRecorder) SecretSet(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SecretSet", reflect.TypeOf((*MockCluster)(nil).SecretSet), arg0, arg1)
}

// SecretSetDefaultSecretKey mocks base method
func (m *MockCluster) SecretSetDefaultSecretKey(arg0 string, arg1 bool) error {
	ret := m.ctrl.Call(m, "SecretSetDefaultSecretKey", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// SecretSetDefaultSecretKey indicates an expected call of SecretSetDefaultSecretKey
func (mr *MockClusterMockRecorder) SecretSetDefaultSecretKey(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SecretSetDefaultSecretKey", reflect.TypeOf((*MockCluster)(nil).SecretSetDefaultSecretKey), arg0, arg1)
}

// SetClusterConf mocks base method
func (m *MockCluster) SetClusterConf(arg0 *osdconfig.ClusterConfig) error {
	ret := m.ctrl.Call(m, "SetClusterConf", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetClusterConf indicates an expected call of SetClusterConf
func (mr *MockClusterMockRecorder) SetClusterConf(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetClusterConf", reflect.TypeOf((*MockCluster)(nil).SetClusterConf), arg0)
}

// SetNodeConf mocks base method
func (m *MockCluster) SetNodeConf(arg0 *osdconfig.NodeConfig) error {
	ret := m.ctrl.Call(m, "SetNodeConf", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetNodeConf indicates an expected call of SetNodeConf
func (mr *MockClusterMockRecorder) SetNodeConf(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetNodeConf", reflect.TypeOf((*MockCluster)(nil).SetNodeConf), arg0)
}

// SetSize mocks base method
func (m *MockCluster) SetSize(arg0 int) error {
	ret := m.ctrl.Call(m, "SetSize", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetSize indicates an expected call of SetSize
func (mr *MockClusterMockRecorder) SetSize(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetSize", reflect.TypeOf((*MockCluster)(nil).SetSize), arg0)
}

// Shutdown mocks base method
func (m *MockCluster) Shutdown() error {
	ret := m.ctrl.Call(m, "Shutdown")
	ret0, _ := ret[0].(error)
	return ret0
}

// Shutdown indicates an expected call of Shutdown
func (mr *MockClusterMockRecorder) Shutdown() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Shutdown", reflect.TypeOf((*MockCluster)(nil).Shutdown))
}

// Start mocks base method
func (m *MockCluster) Start(arg0 int, arg1 bool, arg2, arg3 string) error {
	ret := m.ctrl.Call(m, "Start", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// Start indicates an expected call of Start
func (mr *MockClusterMockRecorder) Start(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Start", reflect.TypeOf((*MockCluster)(nil).Start), arg0, arg1, arg2, arg3)
}

// StartWithConfiguration mocks base method
func (m *MockCluster) StartWithConfiguration(arg0 int, arg1 bool, arg2 string, arg3 []string, arg4 string, arg5 *cluster.ClusterServerConfiguration) error {
	ret := m.ctrl.Call(m, "StartWithConfiguration", arg0, arg1, arg2, arg3, arg4, arg5)
	ret0, _ := ret[0].(error)
	return ret0
}

// StartWithConfiguration indicates an expected call of StartWithConfiguration
func (mr *MockClusterMockRecorder) StartWithConfiguration(arg0, arg1, arg2, arg3, arg4, arg5 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StartWithConfiguration", reflect.TypeOf((*MockCluster)(nil).StartWithConfiguration), arg0, arg1, arg2, arg3, arg4, arg5)
}

// UpdateData mocks base method
func (m *MockCluster) UpdateData(arg0 map[string]interface{}) error {
	ret := m.ctrl.Call(m, "UpdateData", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateData indicates an expected call of UpdateData
func (mr *MockClusterMockRecorder) UpdateData(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateData", reflect.TypeOf((*MockCluster)(nil).UpdateData), arg0)
}

// UpdateLabels mocks base method
func (m *MockCluster) UpdateLabels(arg0 map[string]string) error {
	ret := m.ctrl.Call(m, "UpdateLabels", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateLabels indicates an expected call of UpdateLabels
func (mr *MockClusterMockRecorder) UpdateLabels(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateLabels", reflect.TypeOf((*MockCluster)(nil).UpdateLabels), arg0)
}

// UpdateSchedulerNodeName mocks base method
func (m *MockCluster) UpdateSchedulerNodeName(arg0 string) error {
	ret := m.ctrl.Call(m, "UpdateSchedulerNodeName", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateSchedulerNodeName indicates an expected call of UpdateSchedulerNodeName
func (mr *MockClusterMockRecorder) UpdateSchedulerNodeName(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateSchedulerNodeName", reflect.TypeOf((*MockCluster)(nil).UpdateSchedulerNodeName), arg0)
}

// UpdateSelfClusterDomain mocks base method
func (m *MockCluster) UpdateSelfClusterDomain(arg0 string) error {
	ret := m.ctrl.Call(m, "UpdateSelfClusterDomain", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateSelfClusterDomain indicates an expected call of UpdateSelfClusterDomain
func (mr *MockClusterMockRecorder) UpdateSelfClusterDomain(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateSelfClusterDomain", reflect.TypeOf((*MockCluster)(nil).UpdateSelfClusterDomain), arg0)
}

// Uuid mocks base method
func (m *MockCluster) Uuid() string {
	ret := m.ctrl.Call(m, "Uuid")
	ret0, _ := ret[0].(string)
	return ret0
}

// Uuid indicates an expected call of Uuid
func (mr *MockClusterMockRecorder) Uuid() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Uuid", reflect.TypeOf((*MockCluster)(nil).Uuid))
}

// ValidatePair mocks base method
func (m *MockCluster) ValidatePair(arg0 string) error {
	ret := m.ctrl.Call(m, "ValidatePair", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// ValidatePair indicates an expected call of ValidatePair
func (mr *MockClusterMockRecorder) ValidatePair(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidatePair", reflect.TypeOf((*MockCluster)(nil).ValidatePair), arg0)
}
