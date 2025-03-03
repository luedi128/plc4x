/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

// Code generated by mockery v2.42.2. DO NOT EDIT.

package testutils

import mock "github.com/stretchr/testify/mock"

// MockWithTestCaseOption is an autogenerated mock type for the WithTestCaseOption type
type MockWithTestCaseOption struct {
	mock.Mock
}

type MockWithTestCaseOption_Expecter struct {
	mock *mock.Mock
}

func (_m *MockWithTestCaseOption) EXPECT() *MockWithTestCaseOption_Expecter {
	return &MockWithTestCaseOption_Expecter{mock: &_m.Mock}
}

// Execute provides a mock function with given fields: testCase
func (_m *MockWithTestCaseOption) Execute(testCase ManualTestCase) {
	_m.Called(testCase)
}

// MockWithTestCaseOption_Execute_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Execute'
type MockWithTestCaseOption_Execute_Call struct {
	*mock.Call
}

// Execute is a helper method to define mock.On call
//   - testCase ManualTestCase
func (_e *MockWithTestCaseOption_Expecter) Execute(testCase interface{}) *MockWithTestCaseOption_Execute_Call {
	return &MockWithTestCaseOption_Execute_Call{Call: _e.mock.On("Execute", testCase)}
}

func (_c *MockWithTestCaseOption_Execute_Call) Run(run func(testCase ManualTestCase)) *MockWithTestCaseOption_Execute_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(ManualTestCase))
	})
	return _c
}

func (_c *MockWithTestCaseOption_Execute_Call) Return() *MockWithTestCaseOption_Execute_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockWithTestCaseOption_Execute_Call) RunAndReturn(run func(ManualTestCase)) *MockWithTestCaseOption_Execute_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockWithTestCaseOption creates a new instance of MockWithTestCaseOption. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockWithTestCaseOption(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockWithTestCaseOption {
	mock := &MockWithTestCaseOption{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
