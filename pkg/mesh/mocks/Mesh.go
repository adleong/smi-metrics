// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import context "context"
import mesh "github.com/deislabs/smi-metrics/pkg/mesh"
import metrics "github.com/deislabs/smi-sdk-go/pkg/apis/metrics"
import mock "github.com/stretchr/testify/mock"
import v1 "k8s.io/apimachinery/pkg/apis/meta/v1"

// Mesh is an autogenerated mock type for the Mesh type
type Mesh struct {
	mock.Mock
}

// GetEdgeMetrics provides a mock function with given fields: ctx, name, namespace, kind, interval, details
func (_m *Mesh) GetEdgeMetrics(ctx context.Context, name string, namespace string, kind string, interval *metrics.Interval, details *mesh.ResourceDetails) (*metrics.TrafficMetricsList, error) {
	ret := _m.Called(ctx, name, namespace, kind, interval, details)

	var r0 *metrics.TrafficMetricsList
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string, *metrics.Interval, *mesh.ResourceDetails) *metrics.TrafficMetricsList); ok {
		r0 = rf(ctx, name, namespace, kind, interval, details)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*metrics.TrafficMetricsList)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string, string, *metrics.Interval, *mesh.ResourceDetails) error); ok {
		r1 = rf(ctx, name, namespace, kind, interval, details)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetResourceMetrics provides a mock function with given fields: ctx, name, namespace, kind, interval
func (_m *Mesh) GetResourceMetrics(ctx context.Context, name string, namespace string, kind string, interval *metrics.Interval) (*metrics.TrafficMetricsList, error) {
	ret := _m.Called(ctx, name, namespace, kind, interval)

	var r0 *metrics.TrafficMetricsList
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string, *metrics.Interval) *metrics.TrafficMetricsList); ok {
		r0 = rf(ctx, name, namespace, kind, interval)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*metrics.TrafficMetricsList)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string, string, *metrics.Interval) error); ok {
		r1 = rf(ctx, name, namespace, kind, interval)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetSupportedResources provides a mock function with given fields: ctx
func (_m *Mesh) GetSupportedResources(ctx context.Context) (*v1.APIResourceList, error) {
	ret := _m.Called(ctx)

	var r0 *v1.APIResourceList
	if rf, ok := ret.Get(0).(func(context.Context) *v1.APIResourceList); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1.APIResourceList)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
