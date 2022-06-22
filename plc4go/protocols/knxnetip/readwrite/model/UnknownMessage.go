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

// UnknownMessage is the corresponding interface of UnknownMessage
type UnknownMessage interface {
	utils.LengthAware
	utils.Serializable
	KnxNetIpMessage
	// GetUnknownData returns UnknownData (property field)
	GetUnknownData() []byte
}

// _UnknownMessage is the data-structure of this message
type _UnknownMessage struct {
	*_KnxNetIpMessage
	UnknownData []byte

	// Arguments.
	TotalLength uint16
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_UnknownMessage) GetMsgType() uint16 {
	return 0x020B
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_UnknownMessage) InitializeParent(parent KnxNetIpMessage) {}

func (m *_UnknownMessage) GetParent() KnxNetIpMessage {
	return m._KnxNetIpMessage
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_UnknownMessage) GetUnknownData() []byte {
	return m.UnknownData
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewUnknownMessage factory function for _UnknownMessage
func NewUnknownMessage(unknownData []byte, totalLength uint16) *_UnknownMessage {
	_result := &_UnknownMessage{
		UnknownData:      unknownData,
		_KnxNetIpMessage: NewKnxNetIpMessage(),
	}
	_result._KnxNetIpMessage._KnxNetIpMessageChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastUnknownMessage(structType interface{}) UnknownMessage {
	if casted, ok := structType.(UnknownMessage); ok {
		return casted
	}
	if casted, ok := structType.(*UnknownMessage); ok {
		return *casted
	}
	return nil
}

func (m *_UnknownMessage) GetTypeName() string {
	return "UnknownMessage"
}

func (m *_UnknownMessage) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_UnknownMessage) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Array field
	if len(m.UnknownData) > 0 {
		lengthInBits += 8 * uint16(len(m.UnknownData))
	}

	return lengthInBits
}

func (m *_UnknownMessage) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func UnknownMessageParse(readBuffer utils.ReadBuffer, totalLength uint16) (UnknownMessage, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("UnknownMessage"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for UnknownMessage")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos
	// Byte Array field (unknownData)
	numberOfBytesunknownData := int(uint16(totalLength) - uint16(uint16(6)))
	unknownData, _readArrayErr := readBuffer.ReadByteArray("unknownData", numberOfBytesunknownData)
	if _readArrayErr != nil {
		return nil, errors.Wrap(_readArrayErr, "Error parsing 'unknownData' field")
	}

	if closeErr := readBuffer.CloseContext("UnknownMessage"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for UnknownMessage")
	}

	// Create a partially initialized instance
	_child := &_UnknownMessage{
		UnknownData:      unknownData,
		_KnxNetIpMessage: &_KnxNetIpMessage{},
	}
	_child._KnxNetIpMessage._KnxNetIpMessageChildRequirements = _child
	return _child, nil
}

func (m *_UnknownMessage) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("UnknownMessage"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for UnknownMessage")
		}

		// Array Field (unknownData)
		if m.GetUnknownData() != nil {
			// Byte Array field (unknownData)
			_writeArrayErr := writeBuffer.WriteByteArray("unknownData", m.GetUnknownData())
			if _writeArrayErr != nil {
				return errors.Wrap(_writeArrayErr, "Error serializing 'unknownData' field")
			}
		}

		if popErr := writeBuffer.PopContext("UnknownMessage"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for UnknownMessage")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_UnknownMessage) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
