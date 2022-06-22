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

package model

import (
	"github.com/apache/plc4x/plc4go/internal/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// SALDataTerminateRamp is the corresponding interface of SALDataTerminateRamp
type SALDataTerminateRamp interface {
	utils.LengthAware
	utils.Serializable
	SALData
	// GetGroup returns Group (property field)
	GetGroup() byte
}

// _SALDataTerminateRamp is the data-structure of this message
type _SALDataTerminateRamp struct {
	*_SALData
	Group byte
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_SALDataTerminateRamp) InitializeParent(parent SALData, commandTypeContainer SALCommandTypeContainer) {
	m.CommandTypeContainer = commandTypeContainer
}

func (m *_SALDataTerminateRamp) GetParent() SALData {
	return m._SALData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_SALDataTerminateRamp) GetGroup() byte {
	return m.Group
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewSALDataTerminateRamp factory function for _SALDataTerminateRamp
func NewSALDataTerminateRamp(group byte, commandTypeContainer SALCommandTypeContainer) *_SALDataTerminateRamp {
	_result := &_SALDataTerminateRamp{
		Group:    group,
		_SALData: NewSALData(commandTypeContainer),
	}
	_result._SALData._SALDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastSALDataTerminateRamp(structType interface{}) SALDataTerminateRamp {
	if casted, ok := structType.(SALDataTerminateRamp); ok {
		return casted
	}
	if casted, ok := structType.(*SALDataTerminateRamp); ok {
		return *casted
	}
	return nil
}

func (m *_SALDataTerminateRamp) GetTypeName() string {
	return "SALDataTerminateRamp"
}

func (m *_SALDataTerminateRamp) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_SALDataTerminateRamp) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (group)
	lengthInBits += 8

	return lengthInBits
}

func (m *_SALDataTerminateRamp) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func SALDataTerminateRampParse(readBuffer utils.ReadBuffer) (SALDataTerminateRamp, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("SALDataTerminateRamp"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for SALDataTerminateRamp")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (group)
	_group, _groupErr := readBuffer.ReadByte("group")
	if _groupErr != nil {
		return nil, errors.Wrap(_groupErr, "Error parsing 'group' field")
	}
	group := _group

	if closeErr := readBuffer.CloseContext("SALDataTerminateRamp"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for SALDataTerminateRamp")
	}

	// Create a partially initialized instance
	_child := &_SALDataTerminateRamp{
		Group:    group,
		_SALData: &_SALData{},
	}
	_child._SALData._SALDataChildRequirements = _child
	return _child, nil
}

func (m *_SALDataTerminateRamp) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("SALDataTerminateRamp"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for SALDataTerminateRamp")
		}

		// Simple Field (group)
		group := byte(m.GetGroup())
		_groupErr := writeBuffer.WriteByte("group", (group))
		if _groupErr != nil {
			return errors.Wrap(_groupErr, "Error serializing 'group' field")
		}

		if popErr := writeBuffer.PopContext("SALDataTerminateRamp"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for SALDataTerminateRamp")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_SALDataTerminateRamp) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
