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

// BACnetContextTagUnknown is the corresponding interface of BACnetContextTagUnknown
type BACnetContextTagUnknown interface {
	utils.LengthAware
	utils.Serializable
	BACnetContextTag
	// GetUnknownData returns UnknownData (property field)
	GetUnknownData() []byte
}

// _BACnetContextTagUnknown is the data-structure of this message
type _BACnetContextTagUnknown struct {
	*_BACnetContextTag
	UnknownData []byte

	// Arguments.
	TagNumberArgument uint8
	ActualLength      uint32
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetContextTagUnknown) GetDataType() BACnetDataType {
	return BACnetDataType_UNKNOWN
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetContextTagUnknown) InitializeParent(parent BACnetContextTag, header BACnetTagHeader) {
	m.Header = header
}

func (m *_BACnetContextTagUnknown) GetParent() BACnetContextTag {
	return m._BACnetContextTag
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetContextTagUnknown) GetUnknownData() []byte {
	return m.UnknownData
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetContextTagUnknown factory function for _BACnetContextTagUnknown
func NewBACnetContextTagUnknown(unknownData []byte, header BACnetTagHeader, tagNumberArgument uint8, actualLength uint32) *_BACnetContextTagUnknown {
	_result := &_BACnetContextTagUnknown{
		UnknownData:       unknownData,
		_BACnetContextTag: NewBACnetContextTag(header, tagNumberArgument),
	}
	_result._BACnetContextTag._BACnetContextTagChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetContextTagUnknown(structType interface{}) BACnetContextTagUnknown {
	if casted, ok := structType.(BACnetContextTagUnknown); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetContextTagUnknown); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetContextTagUnknown) GetTypeName() string {
	return "BACnetContextTagUnknown"
}

func (m *_BACnetContextTagUnknown) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetContextTagUnknown) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Array field
	if len(m.UnknownData) > 0 {
		lengthInBits += 8 * uint16(len(m.UnknownData))
	}

	return lengthInBits
}

func (m *_BACnetContextTagUnknown) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetContextTagUnknownParse(readBuffer utils.ReadBuffer, tagNumberArgument uint8, dataType BACnetDataType, actualLength uint32) (BACnetContextTagUnknown, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetContextTagUnknown"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetContextTagUnknown")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos
	// Byte Array field (unknownData)
	numberOfBytesunknownData := int(actualLength)
	unknownData, _readArrayErr := readBuffer.ReadByteArray("unknownData", numberOfBytesunknownData)
	if _readArrayErr != nil {
		return nil, errors.Wrap(_readArrayErr, "Error parsing 'unknownData' field")
	}

	if closeErr := readBuffer.CloseContext("BACnetContextTagUnknown"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetContextTagUnknown")
	}

	// Create a partially initialized instance
	_child := &_BACnetContextTagUnknown{
		UnknownData:       unknownData,
		_BACnetContextTag: &_BACnetContextTag{},
	}
	_child._BACnetContextTag._BACnetContextTagChildRequirements = _child
	return _child, nil
}

func (m *_BACnetContextTagUnknown) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetContextTagUnknown"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetContextTagUnknown")
		}

		// Array Field (unknownData)
		if m.GetUnknownData() != nil {
			// Byte Array field (unknownData)
			_writeArrayErr := writeBuffer.WriteByteArray("unknownData", m.GetUnknownData())
			if _writeArrayErr != nil {
				return errors.Wrap(_writeArrayErr, "Error serializing 'unknownData' field")
			}
		}

		if popErr := writeBuffer.PopContext("BACnetContextTagUnknown"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetContextTagUnknown")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetContextTagUnknown) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
