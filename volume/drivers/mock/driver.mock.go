// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/libopenstorage/openstorage/volume (interfaces: VolumeDriver)

// Package mock is a generated GoMock package.
package mock

import (
	gomock "github.com/golang/mock/gomock"
	api "github.com/libopenstorage/openstorage/api"
	reflect "reflect"
)

// MockVolumeDriver is a mock of VolumeDriver interface
type MockVolumeDriver struct {
	ctrl     *gomock.Controller
	recorder *MockVolumeDriverMockRecorder
}

// MockVolumeDriverMockRecorder is the mock recorder for MockVolumeDriver
type MockVolumeDriverMockRecorder struct {
	mock *MockVolumeDriver
}

// NewMockVolumeDriver creates a new mock instance
func NewMockVolumeDriver(ctrl *gomock.Controller) *MockVolumeDriver {
	mock := &MockVolumeDriver{ctrl: ctrl}
	mock.recorder = &MockVolumeDriverMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockVolumeDriver) EXPECT() *MockVolumeDriverMockRecorder {
	return m.recorder
}

// Attach mocks base method
func (m *MockVolumeDriver) Attach(arg0 string, arg1 map[string]string) (string, error) {
	ret := m.ctrl.Call(m, "Attach", arg0, arg1)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Attach indicates an expected call of Attach
func (mr *MockVolumeDriverMockRecorder) Attach(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Attach", reflect.TypeOf((*MockVolumeDriver)(nil).Attach), arg0, arg1)
}

// CapacityUsage mocks base method
func (m *MockVolumeDriver) CapacityUsage(arg0 string) (*api.CapacityUsageResponse, error) {
	ret := m.ctrl.Call(m, "CapacityUsage", arg0)
	ret0, _ := ret[0].(*api.CapacityUsageResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CapacityUsage indicates an expected call of CapacityUsage
func (mr *MockVolumeDriverMockRecorder) CapacityUsage(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CapacityUsage", reflect.TypeOf((*MockVolumeDriver)(nil).CapacityUsage), arg0)
}

// Catalog mocks base method
func (m *MockVolumeDriver) Catalog(arg0, arg1, arg2 string) (api.CatalogResponse, error) {
	ret := m.ctrl.Call(m, "Catalog", arg0, arg1, arg2)
	ret0, _ := ret[0].(api.CatalogResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Catalog indicates an expected call of Catalog
func (mr *MockVolumeDriverMockRecorder) Catalog(arg0, arg1, arg2 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Catalog", reflect.TypeOf((*MockVolumeDriver)(nil).Catalog), arg0, arg1, arg2)
}

// CloudBackupCatalog mocks base method
func (m *MockVolumeDriver) CloudBackupCatalog(arg0 *api.CloudBackupCatalogRequest) (*api.CloudBackupCatalogResponse, error) {
	ret := m.ctrl.Call(m, "CloudBackupCatalog", arg0)
	ret0, _ := ret[0].(*api.CloudBackupCatalogResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CloudBackupCatalog indicates an expected call of CloudBackupCatalog
func (mr *MockVolumeDriverMockRecorder) CloudBackupCatalog(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CloudBackupCatalog", reflect.TypeOf((*MockVolumeDriver)(nil).CloudBackupCatalog), arg0)
}

// CloudBackupCreate mocks base method
func (m *MockVolumeDriver) CloudBackupCreate(arg0 *api.CloudBackupCreateRequest) error {
	ret := m.ctrl.Call(m, "CloudBackupCreate", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// CloudBackupCreate indicates an expected call of CloudBackupCreate
func (mr *MockVolumeDriverMockRecorder) CloudBackupCreate(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CloudBackupCreate", reflect.TypeOf((*MockVolumeDriver)(nil).CloudBackupCreate), arg0)
}

// CloudBackupDelete mocks base method
func (m *MockVolumeDriver) CloudBackupDelete(arg0 *api.CloudBackupDeleteRequest) error {
	ret := m.ctrl.Call(m, "CloudBackupDelete", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// CloudBackupDelete indicates an expected call of CloudBackupDelete
func (mr *MockVolumeDriverMockRecorder) CloudBackupDelete(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CloudBackupDelete", reflect.TypeOf((*MockVolumeDriver)(nil).CloudBackupDelete), arg0)
}

// CloudBackupDeleteAll mocks base method
func (m *MockVolumeDriver) CloudBackupDeleteAll(arg0 *api.CloudBackupDeleteAllRequest) error {
	ret := m.ctrl.Call(m, "CloudBackupDeleteAll", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// CloudBackupDeleteAll indicates an expected call of CloudBackupDeleteAll
func (mr *MockVolumeDriverMockRecorder) CloudBackupDeleteAll(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CloudBackupDeleteAll", reflect.TypeOf((*MockVolumeDriver)(nil).CloudBackupDeleteAll), arg0)
}

// CloudBackupEnumerate mocks base method
func (m *MockVolumeDriver) CloudBackupEnumerate(arg0 *api.CloudBackupEnumerateRequest) (*api.CloudBackupEnumerateResponse, error) {
	ret := m.ctrl.Call(m, "CloudBackupEnumerate", arg0)
	ret0, _ := ret[0].(*api.CloudBackupEnumerateResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CloudBackupEnumerate indicates an expected call of CloudBackupEnumerate
func (mr *MockVolumeDriverMockRecorder) CloudBackupEnumerate(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CloudBackupEnumerate", reflect.TypeOf((*MockVolumeDriver)(nil).CloudBackupEnumerate), arg0)
}

// CloudBackupGroupCreate mocks base method
func (m *MockVolumeDriver) CloudBackupGroupCreate(arg0 *api.CloudBackupGroupCreateRequest) error {
	ret := m.ctrl.Call(m, "CloudBackupGroupCreate", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// CloudBackupGroupCreate indicates an expected call of CloudBackupGroupCreate
func (mr *MockVolumeDriverMockRecorder) CloudBackupGroupCreate(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CloudBackupGroupCreate", reflect.TypeOf((*MockVolumeDriver)(nil).CloudBackupGroupCreate), arg0)
}

// CloudBackupGroupSchedCreate mocks base method
func (m *MockVolumeDriver) CloudBackupGroupSchedCreate(arg0 *api.CloudBackupGroupSchedCreateRequest) (*api.CloudBackupSchedCreateResponse, error) {
	ret := m.ctrl.Call(m, "CloudBackupGroupSchedCreate", arg0)
	ret0, _ := ret[0].(*api.CloudBackupSchedCreateResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CloudBackupGroupSchedCreate indicates an expected call of CloudBackupGroupSchedCreate
func (mr *MockVolumeDriverMockRecorder) CloudBackupGroupSchedCreate(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CloudBackupGroupSchedCreate", reflect.TypeOf((*MockVolumeDriver)(nil).CloudBackupGroupSchedCreate), arg0)
}

// CloudBackupHistory mocks base method
func (m *MockVolumeDriver) CloudBackupHistory(arg0 *api.CloudBackupHistoryRequest) (*api.CloudBackupHistoryResponse, error) {
	ret := m.ctrl.Call(m, "CloudBackupHistory", arg0)
	ret0, _ := ret[0].(*api.CloudBackupHistoryResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CloudBackupHistory indicates an expected call of CloudBackupHistory
func (mr *MockVolumeDriverMockRecorder) CloudBackupHistory(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CloudBackupHistory", reflect.TypeOf((*MockVolumeDriver)(nil).CloudBackupHistory), arg0)
}

// CloudBackupRestore mocks base method
func (m *MockVolumeDriver) CloudBackupRestore(arg0 *api.CloudBackupRestoreRequest) (*api.CloudBackupRestoreResponse, error) {
	ret := m.ctrl.Call(m, "CloudBackupRestore", arg0)
	ret0, _ := ret[0].(*api.CloudBackupRestoreResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CloudBackupRestore indicates an expected call of CloudBackupRestore
func (mr *MockVolumeDriverMockRecorder) CloudBackupRestore(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CloudBackupRestore", reflect.TypeOf((*MockVolumeDriver)(nil).CloudBackupRestore), arg0)
}

// CloudBackupSchedCreate mocks base method
func (m *MockVolumeDriver) CloudBackupSchedCreate(arg0 *api.CloudBackupSchedCreateRequest) (*api.CloudBackupSchedCreateResponse, error) {
	ret := m.ctrl.Call(m, "CloudBackupSchedCreate", arg0)
	ret0, _ := ret[0].(*api.CloudBackupSchedCreateResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CloudBackupSchedCreate indicates an expected call of CloudBackupSchedCreate
func (mr *MockVolumeDriverMockRecorder) CloudBackupSchedCreate(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CloudBackupSchedCreate", reflect.TypeOf((*MockVolumeDriver)(nil).CloudBackupSchedCreate), arg0)
}

// CloudBackupSchedDelete mocks base method
func (m *MockVolumeDriver) CloudBackupSchedDelete(arg0 *api.CloudBackupSchedDeleteRequest) error {
	ret := m.ctrl.Call(m, "CloudBackupSchedDelete", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// CloudBackupSchedDelete indicates an expected call of CloudBackupSchedDelete
func (mr *MockVolumeDriverMockRecorder) CloudBackupSchedDelete(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CloudBackupSchedDelete", reflect.TypeOf((*MockVolumeDriver)(nil).CloudBackupSchedDelete), arg0)
}

// CloudBackupSchedEnumerate mocks base method
func (m *MockVolumeDriver) CloudBackupSchedEnumerate() (*api.CloudBackupSchedEnumerateResponse, error) {
	ret := m.ctrl.Call(m, "CloudBackupSchedEnumerate")
	ret0, _ := ret[0].(*api.CloudBackupSchedEnumerateResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CloudBackupSchedEnumerate indicates an expected call of CloudBackupSchedEnumerate
func (mr *MockVolumeDriverMockRecorder) CloudBackupSchedEnumerate() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CloudBackupSchedEnumerate", reflect.TypeOf((*MockVolumeDriver)(nil).CloudBackupSchedEnumerate))
}

// CloudBackupStateChange mocks base method
func (m *MockVolumeDriver) CloudBackupStateChange(arg0 *api.CloudBackupStateChangeRequest) error {
	ret := m.ctrl.Call(m, "CloudBackupStateChange", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// CloudBackupStateChange indicates an expected call of CloudBackupStateChange
func (mr *MockVolumeDriverMockRecorder) CloudBackupStateChange(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CloudBackupStateChange", reflect.TypeOf((*MockVolumeDriver)(nil).CloudBackupStateChange), arg0)
}

// CloudBackupStatus mocks base method
func (m *MockVolumeDriver) CloudBackupStatus(arg0 *api.CloudBackupStatusRequest) (*api.CloudBackupStatusResponse, error) {
	ret := m.ctrl.Call(m, "CloudBackupStatus", arg0)
	ret0, _ := ret[0].(*api.CloudBackupStatusResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CloudBackupStatus indicates an expected call of CloudBackupStatus
func (mr *MockVolumeDriverMockRecorder) CloudBackupStatus(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CloudBackupStatus", reflect.TypeOf((*MockVolumeDriver)(nil).CloudBackupStatus), arg0)
}

// CloudMigrateCancel mocks base method
func (m *MockVolumeDriver) CloudMigrateCancel(arg0 *api.CloudMigrateCancelRequest) error {
	ret := m.ctrl.Call(m, "CloudMigrateCancel", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// CloudMigrateCancel indicates an expected call of CloudMigrateCancel
func (mr *MockVolumeDriverMockRecorder) CloudMigrateCancel(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CloudMigrateCancel", reflect.TypeOf((*MockVolumeDriver)(nil).CloudMigrateCancel), arg0)
}

// CloudMigrateStart mocks base method
func (m *MockVolumeDriver) CloudMigrateStart(arg0 *api.CloudMigrateStartRequest) error {
	ret := m.ctrl.Call(m, "CloudMigrateStart", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// CloudMigrateStart indicates an expected call of CloudMigrateStart
func (mr *MockVolumeDriverMockRecorder) CloudMigrateStart(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CloudMigrateStart", reflect.TypeOf((*MockVolumeDriver)(nil).CloudMigrateStart), arg0)
}

// CloudMigrateStatus mocks base method
func (m *MockVolumeDriver) CloudMigrateStatus() (*api.CloudMigrateStatusResponse, error) {
	ret := m.ctrl.Call(m, "CloudMigrateStatus")
	ret0, _ := ret[0].(*api.CloudMigrateStatusResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CloudMigrateStatus indicates an expected call of CloudMigrateStatus
func (mr *MockVolumeDriverMockRecorder) CloudMigrateStatus() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CloudMigrateStatus", reflect.TypeOf((*MockVolumeDriver)(nil).CloudMigrateStatus))
}

// Create mocks base method
func (m *MockVolumeDriver) Create(arg0 *api.VolumeLocator, arg1 *api.Source, arg2 *api.VolumeSpec) (string, error) {
	ret := m.ctrl.Call(m, "Create", arg0, arg1, arg2)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create
func (mr *MockVolumeDriverMockRecorder) Create(arg0, arg1, arg2 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockVolumeDriver)(nil).Create), arg0, arg1, arg2)
}

// CredsCreate mocks base method
func (m *MockVolumeDriver) CredsCreate(arg0 map[string]string) (string, error) {
	ret := m.ctrl.Call(m, "CredsCreate", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CredsCreate indicates an expected call of CredsCreate
func (mr *MockVolumeDriverMockRecorder) CredsCreate(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CredsCreate", reflect.TypeOf((*MockVolumeDriver)(nil).CredsCreate), arg0)
}

// CredsDelete mocks base method
func (m *MockVolumeDriver) CredsDelete(arg0 string) error {
	ret := m.ctrl.Call(m, "CredsDelete", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// CredsDelete indicates an expected call of CredsDelete
func (mr *MockVolumeDriverMockRecorder) CredsDelete(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CredsDelete", reflect.TypeOf((*MockVolumeDriver)(nil).CredsDelete), arg0)
}

// CredsEnumerate mocks base method
func (m *MockVolumeDriver) CredsEnumerate() (map[string]interface{}, error) {
	ret := m.ctrl.Call(m, "CredsEnumerate")
	ret0, _ := ret[0].(map[string]interface{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CredsEnumerate indicates an expected call of CredsEnumerate
func (mr *MockVolumeDriverMockRecorder) CredsEnumerate() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CredsEnumerate", reflect.TypeOf((*MockVolumeDriver)(nil).CredsEnumerate))
}

// CredsValidate mocks base method
func (m *MockVolumeDriver) CredsValidate(arg0 string) error {
	ret := m.ctrl.Call(m, "CredsValidate", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// CredsValidate indicates an expected call of CredsValidate
func (mr *MockVolumeDriverMockRecorder) CredsValidate(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CredsValidate", reflect.TypeOf((*MockVolumeDriver)(nil).CredsValidate), arg0)
}

// Delete mocks base method
func (m *MockVolumeDriver) Delete(arg0 string) error {
	ret := m.ctrl.Call(m, "Delete", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete
func (mr *MockVolumeDriverMockRecorder) Delete(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockVolumeDriver)(nil).Delete), arg0)
}

// Detach mocks base method
func (m *MockVolumeDriver) Detach(arg0 string, arg1 map[string]string) error {
	ret := m.ctrl.Call(m, "Detach", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Detach indicates an expected call of Detach
func (mr *MockVolumeDriverMockRecorder) Detach(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Detach", reflect.TypeOf((*MockVolumeDriver)(nil).Detach), arg0, arg1)
}

// Enumerate mocks base method
func (m *MockVolumeDriver) Enumerate(arg0 *api.VolumeLocator, arg1 map[string]string) ([]*api.Volume, error) {
	ret := m.ctrl.Call(m, "Enumerate", arg0, arg1)
	ret0, _ := ret[0].([]*api.Volume)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Enumerate indicates an expected call of Enumerate
func (mr *MockVolumeDriverMockRecorder) Enumerate(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Enumerate", reflect.TypeOf((*MockVolumeDriver)(nil).Enumerate), arg0, arg1)
}

// Flush mocks base method
func (m *MockVolumeDriver) Flush(arg0 string) error {
	ret := m.ctrl.Call(m, "Flush", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Flush indicates an expected call of Flush
func (mr *MockVolumeDriverMockRecorder) Flush(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Flush", reflect.TypeOf((*MockVolumeDriver)(nil).Flush), arg0)
}

// GetActiveRequests mocks base method
func (m *MockVolumeDriver) GetActiveRequests() (*api.ActiveRequests, error) {
	ret := m.ctrl.Call(m, "GetActiveRequests")
	ret0, _ := ret[0].(*api.ActiveRequests)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetActiveRequests indicates an expected call of GetActiveRequests
func (mr *MockVolumeDriverMockRecorder) GetActiveRequests() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetActiveRequests", reflect.TypeOf((*MockVolumeDriver)(nil).GetActiveRequests))
}

// Inspect mocks base method
func (m *MockVolumeDriver) Inspect(arg0 []string) ([]*api.Volume, error) {
	ret := m.ctrl.Call(m, "Inspect", arg0)
	ret0, _ := ret[0].([]*api.Volume)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Inspect indicates an expected call of Inspect
func (mr *MockVolumeDriverMockRecorder) Inspect(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Inspect", reflect.TypeOf((*MockVolumeDriver)(nil).Inspect), arg0)
}

// Mount mocks base method
func (m *MockVolumeDriver) Mount(arg0, arg1 string, arg2 map[string]string) error {
	ret := m.ctrl.Call(m, "Mount", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// Mount indicates an expected call of Mount
func (mr *MockVolumeDriverMockRecorder) Mount(arg0, arg1, arg2 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Mount", reflect.TypeOf((*MockVolumeDriver)(nil).Mount), arg0, arg1, arg2)
}

// MountedAt mocks base method
func (m *MockVolumeDriver) MountedAt(arg0 string) string {
	ret := m.ctrl.Call(m, "MountedAt", arg0)
	ret0, _ := ret[0].(string)
	return ret0
}

// MountedAt indicates an expected call of MountedAt
func (mr *MockVolumeDriverMockRecorder) MountedAt(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MountedAt", reflect.TypeOf((*MockVolumeDriver)(nil).MountedAt), arg0)
}

// Name mocks base method
func (m *MockVolumeDriver) Name() string {
	ret := m.ctrl.Call(m, "Name")
	ret0, _ := ret[0].(string)
	return ret0
}

// Name indicates an expected call of Name
func (mr *MockVolumeDriverMockRecorder) Name() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Name", reflect.TypeOf((*MockVolumeDriver)(nil).Name))
}

// Quiesce mocks base method
func (m *MockVolumeDriver) Quiesce(arg0 string, arg1 uint64, arg2 string) error {
	ret := m.ctrl.Call(m, "Quiesce", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// Quiesce indicates an expected call of Quiesce
func (mr *MockVolumeDriverMockRecorder) Quiesce(arg0, arg1, arg2 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Quiesce", reflect.TypeOf((*MockVolumeDriver)(nil).Quiesce), arg0, arg1, arg2)
}

// Read mocks base method
func (m *MockVolumeDriver) Read(arg0 string, arg1 []byte, arg2 uint64, arg3 int64) (int64, error) {
	ret := m.ctrl.Call(m, "Read", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Read indicates an expected call of Read
func (mr *MockVolumeDriverMockRecorder) Read(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Read", reflect.TypeOf((*MockVolumeDriver)(nil).Read), arg0, arg1, arg2, arg3)
}

// Restore mocks base method
func (m *MockVolumeDriver) Restore(arg0, arg1 string) error {
	ret := m.ctrl.Call(m, "Restore", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Restore indicates an expected call of Restore
func (mr *MockVolumeDriverMockRecorder) Restore(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Restore", reflect.TypeOf((*MockVolumeDriver)(nil).Restore), arg0, arg1)
}

// Set mocks base method
func (m *MockVolumeDriver) Set(arg0 string, arg1 *api.VolumeLocator, arg2 *api.VolumeSpec) error {
	ret := m.ctrl.Call(m, "Set", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// Set indicates an expected call of Set
func (mr *MockVolumeDriverMockRecorder) Set(arg0, arg1, arg2 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Set", reflect.TypeOf((*MockVolumeDriver)(nil).Set), arg0, arg1, arg2)
}

// Shutdown mocks base method
func (m *MockVolumeDriver) Shutdown() {
	m.ctrl.Call(m, "Shutdown")
}

// Shutdown indicates an expected call of Shutdown
func (mr *MockVolumeDriverMockRecorder) Shutdown() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Shutdown", reflect.TypeOf((*MockVolumeDriver)(nil).Shutdown))
}

// SnapEnumerate mocks base method
func (m *MockVolumeDriver) SnapEnumerate(arg0 []string, arg1 map[string]string) ([]*api.Volume, error) {
	ret := m.ctrl.Call(m, "SnapEnumerate", arg0, arg1)
	ret0, _ := ret[0].([]*api.Volume)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SnapEnumerate indicates an expected call of SnapEnumerate
func (mr *MockVolumeDriverMockRecorder) SnapEnumerate(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SnapEnumerate", reflect.TypeOf((*MockVolumeDriver)(nil).SnapEnumerate), arg0, arg1)
}

// Snapshot mocks base method
func (m *MockVolumeDriver) Snapshot(arg0 string, arg1 bool, arg2 *api.VolumeLocator, arg3 bool) (string, error) {
	ret := m.ctrl.Call(m, "Snapshot", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Snapshot indicates an expected call of Snapshot
func (mr *MockVolumeDriverMockRecorder) Snapshot(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Snapshot", reflect.TypeOf((*MockVolumeDriver)(nil).Snapshot), arg0, arg1, arg2, arg3)
}

// SnapshotGroup mocks base method
func (m *MockVolumeDriver) SnapshotGroup(arg0 string, arg1 map[string]string) (*api.GroupSnapCreateResponse, error) {
	ret := m.ctrl.Call(m, "SnapshotGroup", arg0, arg1)
	ret0, _ := ret[0].(*api.GroupSnapCreateResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SnapshotGroup indicates an expected call of SnapshotGroup
func (mr *MockVolumeDriverMockRecorder) SnapshotGroup(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SnapshotGroup", reflect.TypeOf((*MockVolumeDriver)(nil).SnapshotGroup), arg0, arg1)
}

// Stats mocks base method
func (m *MockVolumeDriver) Stats(arg0 string, arg1 bool) (*api.Stats, error) {
	ret := m.ctrl.Call(m, "Stats", arg0, arg1)
	ret0, _ := ret[0].(*api.Stats)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Stats indicates an expected call of Stats
func (mr *MockVolumeDriverMockRecorder) Stats(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stats", reflect.TypeOf((*MockVolumeDriver)(nil).Stats), arg0, arg1)
}

// Status mocks base method
func (m *MockVolumeDriver) Status() [][2]string {
	ret := m.ctrl.Call(m, "Status")
	ret0, _ := ret[0].([][2]string)
	return ret0
}

// Status indicates an expected call of Status
func (mr *MockVolumeDriverMockRecorder) Status() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Status", reflect.TypeOf((*MockVolumeDriver)(nil).Status))
}

// Type mocks base method
func (m *MockVolumeDriver) Type() api.DriverType {
	ret := m.ctrl.Call(m, "Type")
	ret0, _ := ret[0].(api.DriverType)
	return ret0
}

// Type indicates an expected call of Type
func (mr *MockVolumeDriverMockRecorder) Type() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Type", reflect.TypeOf((*MockVolumeDriver)(nil).Type))
}

// Unmount mocks base method
func (m *MockVolumeDriver) Unmount(arg0, arg1 string, arg2 map[string]string) error {
	ret := m.ctrl.Call(m, "Unmount", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// Unmount indicates an expected call of Unmount
func (mr *MockVolumeDriverMockRecorder) Unmount(arg0, arg1, arg2 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Unmount", reflect.TypeOf((*MockVolumeDriver)(nil).Unmount), arg0, arg1, arg2)
}

// Unquiesce mocks base method
func (m *MockVolumeDriver) Unquiesce(arg0 string) error {
	ret := m.ctrl.Call(m, "Unquiesce", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Unquiesce indicates an expected call of Unquiesce
func (mr *MockVolumeDriverMockRecorder) Unquiesce(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Unquiesce", reflect.TypeOf((*MockVolumeDriver)(nil).Unquiesce), arg0)
}

// UsedSize mocks base method
func (m *MockVolumeDriver) UsedSize(arg0 string) (uint64, error) {
	ret := m.ctrl.Call(m, "UsedSize", arg0)
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UsedSize indicates an expected call of UsedSize
func (mr *MockVolumeDriverMockRecorder) UsedSize(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UsedSize", reflect.TypeOf((*MockVolumeDriver)(nil).UsedSize), arg0)
}

// Version mocks base method
func (m *MockVolumeDriver) Version() (*api.StorageVersion, error) {
	ret := m.ctrl.Call(m, "Version")
	ret0, _ := ret[0].(*api.StorageVersion)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Version indicates an expected call of Version
func (mr *MockVolumeDriverMockRecorder) Version() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Version", reflect.TypeOf((*MockVolumeDriver)(nil).Version))
}

// Write mocks base method
func (m *MockVolumeDriver) Write(arg0 string, arg1 []byte, arg2 uint64, arg3 int64) (int64, error) {
	ret := m.ctrl.Call(m, "Write", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Write indicates an expected call of Write
func (mr *MockVolumeDriverMockRecorder) Write(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Write", reflect.TypeOf((*MockVolumeDriver)(nil).Write), arg0, arg1, arg2, arg3)
}
