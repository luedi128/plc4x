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

// ApduDataAdcRead is the corresponding interface of ApduDataAdcRead
type ApduDataAdcRead interface {
	utils.LengthAware
	utils.Serializable
	ApduData
}

// _ApduDataAdcRead is the data-structure of this message
type _ApduDataAdcRead struct {
	*_ApduData

	// Arguments.
	DataLength uint8
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ApduDataAdcRead) GetApciType() uint8 {
	return 0x6
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ApduDataAdcRead) InitializeParent(parent ApduData) {}

func (m *_ApduDataAdcRead) GetParent() ApduData {
	return m._ApduData
}

// NewApduDataAdcRead factory function for _ApduDataAdcRead
func NewApduDataAdcRead(dataLength uint8) *_ApduDataAdcRead {
	_result := &_ApduDataAdcRead{
		_ApduData: NewApduData(dataLength),
	}
	_result._ApduData._ApduDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastApduDataAdcRead(structType interface{}) ApduDataAdcRead {
	if casted, ok := structType.(ApduDataAdcRead); ok {
		return casted
	}
	if casted, ok := structType.(*ApduDataAdcRead); ok {
		return *casted
	}
	return nil
}

func (m *_ApduDataAdcRead) GetTypeName() string {
	return "ApduDataAdcRead"
}

func (m *_ApduDataAdcRead) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_ApduDataAdcRead) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	return lengthInBits
}

func (m *_ApduDataAdcRead) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func ApduDataAdcReadParse(readBuffer utils.ReadBuffer, dataLength uint8) (ApduDataAdcRead, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ApduDataAdcRead"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ApduDataAdcRead")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("ApduDataAdcRead"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ApduDataAdcRead")
	}

	// Create a partially initialized instance
	_child := &_ApduDataAdcRead{
		_ApduData: &_ApduData{},
	}
	_child._ApduData._ApduDataChildRequirements = _child
	return _child, nil
}

func (m *_ApduDataAdcRead) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ApduDataAdcRead"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ApduDataAdcRead")
		}

		if popErr := writeBuffer.PopContext("ApduDataAdcRead"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ApduDataAdcRead")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_ApduDataAdcRead) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
