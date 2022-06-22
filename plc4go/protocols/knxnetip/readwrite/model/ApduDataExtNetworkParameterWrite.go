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

// ApduDataExtNetworkParameterWrite is the corresponding interface of ApduDataExtNetworkParameterWrite
type ApduDataExtNetworkParameterWrite interface {
	utils.LengthAware
	utils.Serializable
	ApduDataExt
}

// _ApduDataExtNetworkParameterWrite is the data-structure of this message
type _ApduDataExtNetworkParameterWrite struct {
	*_ApduDataExt

	// Arguments.
	Length uint8
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ApduDataExtNetworkParameterWrite) GetExtApciType() uint8 {
	return 0x24
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ApduDataExtNetworkParameterWrite) InitializeParent(parent ApduDataExt) {}

func (m *_ApduDataExtNetworkParameterWrite) GetParent() ApduDataExt {
	return m._ApduDataExt
}

// NewApduDataExtNetworkParameterWrite factory function for _ApduDataExtNetworkParameterWrite
func NewApduDataExtNetworkParameterWrite(length uint8) *_ApduDataExtNetworkParameterWrite {
	_result := &_ApduDataExtNetworkParameterWrite{
		_ApduDataExt: NewApduDataExt(length),
	}
	_result._ApduDataExt._ApduDataExtChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastApduDataExtNetworkParameterWrite(structType interface{}) ApduDataExtNetworkParameterWrite {
	if casted, ok := structType.(ApduDataExtNetworkParameterWrite); ok {
		return casted
	}
	if casted, ok := structType.(*ApduDataExtNetworkParameterWrite); ok {
		return *casted
	}
	return nil
}

func (m *_ApduDataExtNetworkParameterWrite) GetTypeName() string {
	return "ApduDataExtNetworkParameterWrite"
}

func (m *_ApduDataExtNetworkParameterWrite) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_ApduDataExtNetworkParameterWrite) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	return lengthInBits
}

func (m *_ApduDataExtNetworkParameterWrite) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func ApduDataExtNetworkParameterWriteParse(readBuffer utils.ReadBuffer, length uint8) (ApduDataExtNetworkParameterWrite, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ApduDataExtNetworkParameterWrite"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ApduDataExtNetworkParameterWrite")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("ApduDataExtNetworkParameterWrite"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ApduDataExtNetworkParameterWrite")
	}

	// Create a partially initialized instance
	_child := &_ApduDataExtNetworkParameterWrite{
		_ApduDataExt: &_ApduDataExt{},
	}
	_child._ApduDataExt._ApduDataExtChildRequirements = _child
	return _child, nil
}

func (m *_ApduDataExtNetworkParameterWrite) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ApduDataExtNetworkParameterWrite"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ApduDataExtNetworkParameterWrite")
		}

		if popErr := writeBuffer.PopContext("ApduDataExtNetworkParameterWrite"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ApduDataExtNetworkParameterWrite")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_ApduDataExtNetworkParameterWrite) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
