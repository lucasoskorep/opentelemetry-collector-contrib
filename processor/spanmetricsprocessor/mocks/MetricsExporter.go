// Copyright The OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//       http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
	component "go.opentelemetry.io/collector/component"
	pdata "go.opentelemetry.io/collector/consumer/pdata"
)

// MetricsExporter is an autogenerated mock type for the MetricsExporter type
type MetricsExporter struct {
	mock.Mock
}

// ConsumeMetrics provides a mock function with given fields: ctx, md
func (_m *MetricsExporter) ConsumeMetrics(ctx context.Context, md pdata.Metrics) error {
	ret := _m.Called(ctx, md)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, pdata.Metrics) error); ok {
		r0 = rf(ctx, md)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Shutdown provides a mock function with given fields: ctx
func (_m *MetricsExporter) Shutdown(ctx context.Context) error {
	ret := _m.Called(ctx)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Start provides a mock function with given fields: ctx, host
func (_m *MetricsExporter) Start(ctx context.Context, host component.Host) error {
	ret := _m.Called(ctx, host)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, component.Host) error); ok {
		r0 = rf(ctx, host)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
