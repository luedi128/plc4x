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

// IdentifyReplyCommandManufacturer is the corresponding interface of IdentifyReplyCommandManufacturer
type IdentifyReplyCommandManufacturer interface {
	utils.LengthAware
	utils.Serializable
	IdentifyReplyCommand
	// GetManufacturerName returns ManufacturerName (property field)
	GetManufacturerName() string
}

// _IdentifyReplyCommandManufacturer is the data-structure of this message
type _IdentifyReplyCommandManufacturer struct {
	*_IdentifyReplyCommand
	ManufacturerName string
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_IdentifyReplyCommandManufacturer) GetAttribute() Attribute {
	return Attribute_Manufacturer
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_IdentifyReplyCommandManufacturer) InitializeParent(parent IdentifyReplyCommand) {}

func (m *_IdentifyReplyCommandManufacturer) GetParent() IdentifyReplyCommand {
	return m._IdentifyReplyCommand
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_IdentifyReplyCommandManufacturer) GetManufacturerName() string {
	return m.ManufacturerName
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewIdentifyReplyCommandManufacturer factory function for _IdentifyReplyCommandManufacturer
func NewIdentifyReplyCommandManufacturer(manufacturerName string) *_IdentifyReplyCommandManufacturer {
	_result := &_IdentifyReplyCommandManufacturer{
		ManufacturerName:      manufacturerName,
		_IdentifyReplyCommand: NewIdentifyReplyCommand(),
	}
	_result._IdentifyReplyCommand._IdentifyReplyCommandChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastIdentifyReplyCommandManufacturer(structType interface{}) IdentifyReplyCommandManufacturer {
	if casted, ok := structType.(IdentifyReplyCommandManufacturer); ok {
		return casted
	}
	if casted, ok := structType.(*IdentifyReplyCommandManufacturer); ok {
		return *casted
	}
	return nil
}

func (m *_IdentifyReplyCommandManufacturer) GetTypeName() string {
	return "IdentifyReplyCommandManufacturer"
}

func (m *_IdentifyReplyCommandManufacturer) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_IdentifyReplyCommandManufacturer) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (manufacturerName)
	lengthInBits += 64

	return lengthInBits
}

func (m *_IdentifyReplyCommandManufacturer) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func IdentifyReplyCommandManufacturerParse(readBuffer utils.ReadBuffer, attribute Attribute) (IdentifyReplyCommandManufacturer, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("IdentifyReplyCommandManufacturer"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for IdentifyReplyCommandManufacturer")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (manufacturerName)
	_manufacturerName, _manufacturerNameErr := readBuffer.ReadString("manufacturerName", uint32(64))
	if _manufacturerNameErr != nil {
		return nil, errors.Wrap(_manufacturerNameErr, "Error parsing 'manufacturerName' field")
	}
	manufacturerName := _manufacturerName

	if closeErr := readBuffer.CloseContext("IdentifyReplyCommandManufacturer"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for IdentifyReplyCommandManufacturer")
	}

	// Create a partially initialized instance
	_child := &_IdentifyReplyCommandManufacturer{
		ManufacturerName:      manufacturerName,
		_IdentifyReplyCommand: &_IdentifyReplyCommand{},
	}
	_child._IdentifyReplyCommand._IdentifyReplyCommandChildRequirements = _child
	return _child, nil
}

func (m *_IdentifyReplyCommandManufacturer) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("IdentifyReplyCommandManufacturer"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for IdentifyReplyCommandManufacturer")
		}

		// Simple Field (manufacturerName)
		manufacturerName := string(m.GetManufacturerName())
		_manufacturerNameErr := writeBuffer.WriteString("manufacturerName", uint32(64), "UTF-8", (manufacturerName))
		if _manufacturerNameErr != nil {
			return errors.Wrap(_manufacturerNameErr, "Error serializing 'manufacturerName' field")
		}

		if popErr := writeBuffer.PopContext("IdentifyReplyCommandManufacturer"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for IdentifyReplyCommandManufacturer")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_IdentifyReplyCommandManufacturer) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
