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

// BACnetChannelValueCharacterString is the corresponding interface of BACnetChannelValueCharacterString
type BACnetChannelValueCharacterString interface {
	utils.LengthAware
	utils.Serializable
	BACnetChannelValue
	// GetCharacterStringValue returns CharacterStringValue (property field)
	GetCharacterStringValue() BACnetApplicationTagCharacterString
}

// _BACnetChannelValueCharacterString is the data-structure of this message
type _BACnetChannelValueCharacterString struct {
	*_BACnetChannelValue
	CharacterStringValue BACnetApplicationTagCharacterString
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetChannelValueCharacterString) InitializeParent(parent BACnetChannelValue, peekedTagHeader BACnetTagHeader) {
	m.PeekedTagHeader = peekedTagHeader
}

func (m *_BACnetChannelValueCharacterString) GetParent() BACnetChannelValue {
	return m._BACnetChannelValue
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetChannelValueCharacterString) GetCharacterStringValue() BACnetApplicationTagCharacterString {
	return m.CharacterStringValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetChannelValueCharacterString factory function for _BACnetChannelValueCharacterString
func NewBACnetChannelValueCharacterString(characterStringValue BACnetApplicationTagCharacterString, peekedTagHeader BACnetTagHeader) *_BACnetChannelValueCharacterString {
	_result := &_BACnetChannelValueCharacterString{
		CharacterStringValue: characterStringValue,
		_BACnetChannelValue:  NewBACnetChannelValue(peekedTagHeader),
	}
	_result._BACnetChannelValue._BACnetChannelValueChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetChannelValueCharacterString(structType interface{}) BACnetChannelValueCharacterString {
	if casted, ok := structType.(BACnetChannelValueCharacterString); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetChannelValueCharacterString); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetChannelValueCharacterString) GetTypeName() string {
	return "BACnetChannelValueCharacterString"
}

func (m *_BACnetChannelValueCharacterString) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetChannelValueCharacterString) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (characterStringValue)
	lengthInBits += m.CharacterStringValue.GetLengthInBits()

	return lengthInBits
}

func (m *_BACnetChannelValueCharacterString) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetChannelValueCharacterStringParse(readBuffer utils.ReadBuffer) (BACnetChannelValueCharacterString, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetChannelValueCharacterString"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetChannelValueCharacterString")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (characterStringValue)
	if pullErr := readBuffer.PullContext("characterStringValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for characterStringValue")
	}
	_characterStringValue, _characterStringValueErr := BACnetApplicationTagParse(readBuffer)
	if _characterStringValueErr != nil {
		return nil, errors.Wrap(_characterStringValueErr, "Error parsing 'characterStringValue' field")
	}
	characterStringValue := _characterStringValue.(BACnetApplicationTagCharacterString)
	if closeErr := readBuffer.CloseContext("characterStringValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for characterStringValue")
	}

	if closeErr := readBuffer.CloseContext("BACnetChannelValueCharacterString"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetChannelValueCharacterString")
	}

	// Create a partially initialized instance
	_child := &_BACnetChannelValueCharacterString{
		CharacterStringValue: characterStringValue,
		_BACnetChannelValue:  &_BACnetChannelValue{},
	}
	_child._BACnetChannelValue._BACnetChannelValueChildRequirements = _child
	return _child, nil
}

func (m *_BACnetChannelValueCharacterString) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetChannelValueCharacterString"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetChannelValueCharacterString")
		}

		// Simple Field (characterStringValue)
		if pushErr := writeBuffer.PushContext("characterStringValue"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for characterStringValue")
		}
		_characterStringValueErr := writeBuffer.WriteSerializable(m.GetCharacterStringValue())
		if popErr := writeBuffer.PopContext("characterStringValue"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for characterStringValue")
		}
		if _characterStringValueErr != nil {
			return errors.Wrap(_characterStringValueErr, "Error serializing 'characterStringValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetChannelValueCharacterString"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetChannelValueCharacterString")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetChannelValueCharacterString) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
