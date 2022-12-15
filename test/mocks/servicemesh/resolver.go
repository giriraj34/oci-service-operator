// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/servicemanager/servicemesh/references/resolver.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	servicemesh "github.com/oracle/oci-go-sdk/v65/servicemesh"
	v1beta1 "github.com/oracle/oci-service-operator/api/v1beta1"
	v1beta10 "github.com/oracle/oci-service-operator/apis/servicemesh.oci/v1beta1"
	commons "github.com/oracle/oci-service-operator/pkg/servicemanager/servicemesh/utils/commons"
	v1 "k8s.io/api/core/v1"
	v10 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// MockResolver is a mock of Resolver interface.
type MockResolver struct {
	ctrl     *gomock.Controller
	recorder *MockResolverMockRecorder
}

// MockResolverMockRecorder is the mock recorder for MockResolver.
type MockResolverMockRecorder struct {
	mock *MockResolver
}

// NewMockResolver creates a new mock instance.
func NewMockResolver(ctrl *gomock.Controller) *MockResolver {
	mock := &MockResolver{ctrl: ctrl}
	mock.recorder = &MockResolverMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockResolver) EXPECT() *MockResolverMockRecorder {
	return m.recorder
}

// ResolveHasVirtualDeploymentWithListener mocks base method.
func (m *MockResolver) ResolveHasVirtualDeploymentWithListener(ctx context.Context, compartmentId, virtualServiceId *v1beta1.OCID) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ResolveHasVirtualDeploymentWithListener", ctx, compartmentId, virtualServiceId)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ResolveHasVirtualDeploymentWithListener indicates an expected call of ResolveHasVirtualDeploymentWithListener.
func (mr *MockResolverMockRecorder) ResolveHasVirtualDeploymentWithListener(ctx, compartmentId, virtualServiceId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResolveHasVirtualDeploymentWithListener", reflect.TypeOf((*MockResolver)(nil).ResolveHasVirtualDeploymentWithListener), ctx, compartmentId, virtualServiceId)
}

// ResolveIngressGatewayIdAndNameAndMeshId mocks base method.
func (m *MockResolver) ResolveIngressGatewayIdAndNameAndMeshId(ctx context.Context, IngressGatewayRef *v1beta10.RefOrId, crdObj *v10.ObjectMeta) (*commons.ResourceRef, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ResolveIngressGatewayIdAndNameAndMeshId", ctx, IngressGatewayRef, crdObj)
	ret0, _ := ret[0].(*commons.ResourceRef)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ResolveIngressGatewayIdAndNameAndMeshId indicates an expected call of ResolveIngressGatewayIdAndNameAndMeshId.
func (mr *MockResolverMockRecorder) ResolveIngressGatewayIdAndNameAndMeshId(ctx, IngressGatewayRef, crdObj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResolveIngressGatewayIdAndNameAndMeshId", reflect.TypeOf((*MockResolver)(nil).ResolveIngressGatewayIdAndNameAndMeshId), ctx, IngressGatewayRef, crdObj)
}

// ResolveIngressGatewayReference mocks base method.
func (m *MockResolver) ResolveIngressGatewayReference(ctx context.Context, ref *v1beta10.ResourceRef) (*v1beta10.IngressGateway, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ResolveIngressGatewayReference", ctx, ref)
	ret0, _ := ret[0].(*v1beta10.IngressGateway)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ResolveIngressGatewayReference indicates an expected call of ResolveIngressGatewayReference.
func (mr *MockResolverMockRecorder) ResolveIngressGatewayReference(ctx, ref interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResolveIngressGatewayReference", reflect.TypeOf((*MockResolver)(nil).ResolveIngressGatewayReference), ctx, ref)
}

// ResolveMeshId mocks base method.
func (m *MockResolver) ResolveMeshId(ctx context.Context, meshRef *v1beta10.RefOrId, crdObj *v10.ObjectMeta) (*v1beta1.OCID, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ResolveMeshId", ctx, meshRef, crdObj)
	ret0, _ := ret[0].(*v1beta1.OCID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ResolveMeshId indicates an expected call of ResolveMeshId.
func (mr *MockResolverMockRecorder) ResolveMeshId(ctx, meshRef, crdObj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResolveMeshId", reflect.TypeOf((*MockResolver)(nil).ResolveMeshId), ctx, meshRef, crdObj)
}

// ResolveMeshRefById mocks base method.
func (m *MockResolver) ResolveMeshRefById(ctx context.Context, meshId *v1beta1.OCID) (*commons.MeshRef, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ResolveMeshRefById", ctx, meshId)
	ret0, _ := ret[0].(*commons.MeshRef)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ResolveMeshRefById indicates an expected call of ResolveMeshRefById.
func (mr *MockResolverMockRecorder) ResolveMeshRefById(ctx, meshId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResolveMeshRefById", reflect.TypeOf((*MockResolver)(nil).ResolveMeshRefById), ctx, meshId)
}

// ResolveMeshReference mocks base method.
func (m *MockResolver) ResolveMeshReference(ctx context.Context, ref *v1beta10.ResourceRef) (*v1beta10.Mesh, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ResolveMeshReference", ctx, ref)
	ret0, _ := ret[0].(*v1beta10.Mesh)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ResolveMeshReference indicates an expected call of ResolveMeshReference.
func (mr *MockResolverMockRecorder) ResolveMeshReference(ctx, ref interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResolveMeshReference", reflect.TypeOf((*MockResolver)(nil).ResolveMeshReference), ctx, ref)
}

// ResolveResourceRef mocks base method.
func (m *MockResolver) ResolveResourceRef(resourceRef *v1beta10.ResourceRef, crdObj *v10.ObjectMeta) *v1beta10.ResourceRef {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ResolveResourceRef", resourceRef, crdObj)
	ret0, _ := ret[0].(*v1beta10.ResourceRef)
	return ret0
}

// ResolveResourceRef indicates an expected call of ResolveResourceRef.
func (mr *MockResolverMockRecorder) ResolveResourceRef(resourceRef, crdObj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResolveResourceRef", reflect.TypeOf((*MockResolver)(nil).ResolveResourceRef), resourceRef, crdObj)
}

// ResolveServiceReference mocks base method.
func (m *MockResolver) ResolveServiceReference(ctx context.Context, ref *v1beta10.ResourceRef) (*v1.Service, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ResolveServiceReference", ctx, ref)
	ret0, _ := ret[0].(*v1.Service)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ResolveServiceReference indicates an expected call of ResolveServiceReference.
func (mr *MockResolverMockRecorder) ResolveServiceReference(ctx, ref interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResolveServiceReference", reflect.TypeOf((*MockResolver)(nil).ResolveServiceReference), ctx, ref)
}

// ResolveServiceReferenceWithApiReader mocks base method.
func (m *MockResolver) ResolveServiceReferenceWithApiReader(ctx context.Context, ref *v1beta10.ResourceRef) (*v1.Service, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ResolveServiceReferenceWithApiReader", ctx, ref)
	ret0, _ := ret[0].(*v1.Service)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ResolveServiceReferenceWithApiReader indicates an expected call of ResolveServiceReferenceWithApiReader.
func (mr *MockResolverMockRecorder) ResolveServiceReferenceWithApiReader(ctx, ref interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResolveServiceReferenceWithApiReader", reflect.TypeOf((*MockResolver)(nil).ResolveServiceReferenceWithApiReader), ctx, ref)
}

// ResolveVirtualDeploymentId mocks base method.
func (m *MockResolver) ResolveVirtualDeploymentId(ctx context.Context, virtualDeploymentRef *v1beta10.RefOrId, crdObj *v10.ObjectMeta) (*v1beta1.OCID, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ResolveVirtualDeploymentId", ctx, virtualDeploymentRef, crdObj)
	ret0, _ := ret[0].(*v1beta1.OCID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ResolveVirtualDeploymentId indicates an expected call of ResolveVirtualDeploymentId.
func (mr *MockResolverMockRecorder) ResolveVirtualDeploymentId(ctx, virtualDeploymentRef, crdObj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResolveVirtualDeploymentId", reflect.TypeOf((*MockResolver)(nil).ResolveVirtualDeploymentId), ctx, virtualDeploymentRef, crdObj)
}

// ResolveVirtualDeploymentReference mocks base method.
func (m *MockResolver) ResolveVirtualDeploymentReference(ctx context.Context, ref *v1beta10.ResourceRef) (*v1beta10.VirtualDeployment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ResolveVirtualDeploymentReference", ctx, ref)
	ret0, _ := ret[0].(*v1beta10.VirtualDeployment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ResolveVirtualDeploymentReference indicates an expected call of ResolveVirtualDeploymentReference.
func (mr *MockResolverMockRecorder) ResolveVirtualDeploymentReference(ctx, ref interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResolveVirtualDeploymentReference", reflect.TypeOf((*MockResolver)(nil).ResolveVirtualDeploymentReference), ctx, ref)
}

// ResolveVirtualServiceById mocks base method.
func (m *MockResolver) ResolveVirtualServiceById(ctx context.Context, virtualServiceId *v1beta1.OCID) (*servicemesh.VirtualService, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ResolveVirtualServiceById", ctx, virtualServiceId)
	ret0, _ := ret[0].(*servicemesh.VirtualService)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ResolveVirtualServiceById indicates an expected call of ResolveVirtualServiceById.
func (mr *MockResolverMockRecorder) ResolveVirtualServiceById(ctx, virtualServiceId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResolveVirtualServiceById", reflect.TypeOf((*MockResolver)(nil).ResolveVirtualServiceById), ctx, virtualServiceId)
}

// ResolveVirtualServiceIdAndName mocks base method.
func (m *MockResolver) ResolveVirtualServiceIdAndName(ctx context.Context, virtualServiceRef *v1beta10.RefOrId, crdObj *v10.ObjectMeta) (*commons.ResourceRef, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ResolveVirtualServiceIdAndName", ctx, virtualServiceRef, crdObj)
	ret0, _ := ret[0].(*commons.ResourceRef)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ResolveVirtualServiceIdAndName indicates an expected call of ResolveVirtualServiceIdAndName.
func (mr *MockResolverMockRecorder) ResolveVirtualServiceIdAndName(ctx, virtualServiceRef, crdObj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResolveVirtualServiceIdAndName", reflect.TypeOf((*MockResolver)(nil).ResolveVirtualServiceIdAndName), ctx, virtualServiceRef, crdObj)
}

// ResolveVirtualServiceListByNamespace mocks base method.
func (m *MockResolver) ResolveVirtualServiceListByNamespace(ctx context.Context, namespace string) (v1beta10.VirtualServiceList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ResolveVirtualServiceListByNamespace", ctx, namespace)
	ret0, _ := ret[0].(v1beta10.VirtualServiceList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ResolveVirtualServiceListByNamespace indicates an expected call of ResolveVirtualServiceListByNamespace.
func (mr *MockResolverMockRecorder) ResolveVirtualServiceListByNamespace(ctx, namespace interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResolveVirtualServiceListByNamespace", reflect.TypeOf((*MockResolver)(nil).ResolveVirtualServiceListByNamespace), ctx, namespace)
}

// ResolveVirtualServiceRefById mocks base method.
func (m *MockResolver) ResolveVirtualServiceRefById(ctx context.Context, virtualServiceId *v1beta1.OCID) (*commons.ResourceRef, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ResolveVirtualServiceRefById", ctx, virtualServiceId)
	ret0, _ := ret[0].(*commons.ResourceRef)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ResolveVirtualServiceRefById indicates an expected call of ResolveVirtualServiceRefById.
func (mr *MockResolverMockRecorder) ResolveVirtualServiceRefById(ctx, virtualServiceId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResolveVirtualServiceRefById", reflect.TypeOf((*MockResolver)(nil).ResolveVirtualServiceRefById), ctx, virtualServiceId)
}

// ResolveVirtualServiceReference mocks base method.
func (m *MockResolver) ResolveVirtualServiceReference(ctx context.Context, ref *v1beta10.ResourceRef) (*v1beta10.VirtualService, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ResolveVirtualServiceReference", ctx, ref)
	ret0, _ := ret[0].(*v1beta10.VirtualService)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ResolveVirtualServiceReference indicates an expected call of ResolveVirtualServiceReference.
func (mr *MockResolverMockRecorder) ResolveVirtualServiceReference(ctx, ref interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResolveVirtualServiceReference", reflect.TypeOf((*MockResolver)(nil).ResolveVirtualServiceReference), ctx, ref)
}
