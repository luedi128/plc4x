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

// ComObjectTableRealisationType6 is the corresponding interface of ComObjectTableRealisationType6
type ComObjectTableRealisationType6 interface {
	utils.LengthAware
	utils.Serializable
	ComObjectTable
	// GetComObjectDescriptors returns ComObjectDescriptors (property field)
	GetComObjectDescriptors() GroupObjectDescriptorRealisationType6
}

// _ComObjectTableRealisationType6 is the data-structure of this message
type _ComObjectTableRealisationType6 struct {
	*_ComObjectTable
	ComObjectDescriptors GroupObjectDescriptorRealisationType6
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ComObjectTableRealisationType6) GetFirmwareType() FirmwareType {
	return FirmwareType_SYSTEM_300
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ComObjectTableRealisationType6) InitializeParent(parent ComObjectTable) {}

func (m *_ComObjectTableRealisationType6) GetParent() ComObjectTable {
	return m._ComObjectTable
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_ComObjectTableRealisationType6) GetComObjectDescriptors() GroupObjectDescriptorRealisationType6 {
	return m.ComObjectDescriptors
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewComObjectTableRealisationType6 factory function for _ComObjectTableRealisationType6
func NewComObjectTableRealisationType6(comObjectDescriptors GroupObjectDescriptorRealisationType6) *_ComObjectTableRealisationType6 {
	_result := &_ComObjectTableRealisationType6{
		ComObjectDescriptors: comObjectDescriptors,
		_ComObjectTable:      NewComObjectTable(),
	}
	_result._ComObjectTable._ComObjectTableChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastComObjectTableRealisationType6(structType interface{}) ComObjectTableRealisationType6 {
	if casted, ok := structType.(ComObjectTableRealisationType6); ok {
		return casted
	}
	if casted, ok := structType.(*ComObjectTableRealisationType6); ok {
		return *casted
	}
	return nil
}

func (m *_ComObjectTableRealisationType6) GetTypeName() string {
	return "ComObjectTableRealisationType6"
}

func (m *_ComObjectTableRealisationType6) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_ComObjectTableRealisationType6) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (comObjectDescriptors)
	lengthInBits += m.ComObjectDescriptors.GetLengthInBits()

	return lengthInBits
}

func (m *_ComObjectTableRealisationType6) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func ComObjectTableRealisationType6Parse(readBuffer utils.ReadBuffer, firmwareType FirmwareType) (ComObjectTableRealisationType6, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ComObjectTableRealisationType6"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ComObjectTableRealisationType6")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (comObjectDescriptors)
	if pullErr := readBuffer.PullContext("comObjectDescriptors"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for comObjectDescriptors")
	}
	_comObjectDescriptors, _comObjectDescriptorsErr := GroupObjectDescriptorRealisationType6Parse(readBuffer)
	if _comObjectDescriptorsErr != nil {
		return nil, errors.Wrap(_comObjectDescriptorsErr, "Error parsing 'comObjectDescriptors' field")
	}
	comObjectDescriptors := _comObjectDescriptors.(GroupObjectDescriptorRealisationType6)
	if closeErr := readBuffer.CloseContext("comObjectDescriptors"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for comObjectDescriptors")
	}

	if closeErr := readBuffer.CloseContext("ComObjectTableRealisationType6"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ComObjectTableRealisationType6")
	}

	// Create a partially initialized instance
	_child := &_ComObjectTableRealisationType6{
		ComObjectDescriptors: comObjectDescriptors,
		_ComObjectTable:      &_ComObjectTable{},
	}
	_child._ComObjectTable._ComObjectTableChildRequirements = _child
	return _child, nil
}

func (m *_ComObjectTableRealisationType6) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ComObjectTableRealisationType6"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ComObjectTableRealisationType6")
		}

		// Simple Field (comObjectDescriptors)
		if pushErr := writeBuffer.PushContext("comObjectDescriptors"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for comObjectDescriptors")
		}
		_comObjectDescriptorsErr := writeBuffer.WriteSerializable(m.GetComObjectDescriptors())
		if popErr := writeBuffer.PopContext("comObjectDescriptors"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for comObjectDescriptors")
		}
		if _comObjectDescriptorsErr != nil {
			return errors.Wrap(_comObjectDescriptorsErr, "Error serializing 'comObjectDescriptors' field")
		}

		if popErr := writeBuffer.PopContext("ComObjectTableRealisationType6"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ComObjectTableRealisationType6")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_ComObjectTableRealisationType6) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
