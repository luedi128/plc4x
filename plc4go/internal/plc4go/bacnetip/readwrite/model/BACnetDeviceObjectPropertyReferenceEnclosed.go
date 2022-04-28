/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
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
	"github.com/apache/plc4x/plc4go/internal/plc4go/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// BACnetDeviceObjectPropertyReferenceEnclosed is the data-structure of this message
type BACnetDeviceObjectPropertyReferenceEnclosed struct {
	OpeningTag *BACnetOpeningTag
	Value      *BACnetDeviceObjectPropertyReference
	ClosingTag *BACnetClosingTag

	// Arguments.
	TagNumber uint8
}

// IBACnetDeviceObjectPropertyReferenceEnclosed is the corresponding interface of BACnetDeviceObjectPropertyReferenceEnclosed
type IBACnetDeviceObjectPropertyReferenceEnclosed interface {
	// GetOpeningTag returns OpeningTag (property field)
	GetOpeningTag() *BACnetOpeningTag
	// GetValue returns Value (property field)
	GetValue() *BACnetDeviceObjectPropertyReference
	// GetClosingTag returns ClosingTag (property field)
	GetClosingTag() *BACnetClosingTag
	// GetLengthInBytes returns the length in bytes
	GetLengthInBytes() uint16
	// GetLengthInBits returns the length in bits
	GetLengthInBits() uint16
	// Serialize serializes this type
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetDeviceObjectPropertyReferenceEnclosed) GetOpeningTag() *BACnetOpeningTag {
	return m.OpeningTag
}

func (m *BACnetDeviceObjectPropertyReferenceEnclosed) GetValue() *BACnetDeviceObjectPropertyReference {
	return m.Value
}

func (m *BACnetDeviceObjectPropertyReferenceEnclosed) GetClosingTag() *BACnetClosingTag {
	return m.ClosingTag
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetDeviceObjectPropertyReferenceEnclosed factory function for BACnetDeviceObjectPropertyReferenceEnclosed
func NewBACnetDeviceObjectPropertyReferenceEnclosed(openingTag *BACnetOpeningTag, value *BACnetDeviceObjectPropertyReference, closingTag *BACnetClosingTag, tagNumber uint8) *BACnetDeviceObjectPropertyReferenceEnclosed {
	return &BACnetDeviceObjectPropertyReferenceEnclosed{OpeningTag: openingTag, Value: value, ClosingTag: closingTag, TagNumber: tagNumber}
}

func CastBACnetDeviceObjectPropertyReferenceEnclosed(structType interface{}) *BACnetDeviceObjectPropertyReferenceEnclosed {
	if casted, ok := structType.(BACnetDeviceObjectPropertyReferenceEnclosed); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetDeviceObjectPropertyReferenceEnclosed); ok {
		return casted
	}
	return nil
}

func (m *BACnetDeviceObjectPropertyReferenceEnclosed) GetTypeName() string {
	return "BACnetDeviceObjectPropertyReferenceEnclosed"
}

func (m *BACnetDeviceObjectPropertyReferenceEnclosed) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetDeviceObjectPropertyReferenceEnclosed) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(0)

	// Simple field (openingTag)
	lengthInBits += m.OpeningTag.GetLengthInBits()

	// Simple field (value)
	lengthInBits += m.Value.GetLengthInBits()

	// Simple field (closingTag)
	lengthInBits += m.ClosingTag.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetDeviceObjectPropertyReferenceEnclosed) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetDeviceObjectPropertyReferenceEnclosedParse(readBuffer utils.ReadBuffer, tagNumber uint8) (*BACnetDeviceObjectPropertyReferenceEnclosed, error) {
	if pullErr := readBuffer.PullContext("BACnetDeviceObjectPropertyReferenceEnclosed"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := readBuffer.GetPos()
	_ = currentPos

	// Simple Field (openingTag)
	if pullErr := readBuffer.PullContext("openingTag"); pullErr != nil {
		return nil, pullErr
	}
	_openingTag, _openingTagErr := BACnetContextTagParse(readBuffer, uint8(tagNumber), BACnetDataType(BACnetDataType_OPENING_TAG))
	if _openingTagErr != nil {
		return nil, errors.Wrap(_openingTagErr, "Error parsing 'openingTag' field")
	}
	openingTag := CastBACnetOpeningTag(_openingTag)
	if closeErr := readBuffer.CloseContext("openingTag"); closeErr != nil {
		return nil, closeErr
	}

	// Simple Field (value)
	if pullErr := readBuffer.PullContext("value"); pullErr != nil {
		return nil, pullErr
	}
	_value, _valueErr := BACnetDeviceObjectPropertyReferenceParse(readBuffer)
	if _valueErr != nil {
		return nil, errors.Wrap(_valueErr, "Error parsing 'value' field")
	}
	value := CastBACnetDeviceObjectPropertyReference(_value)
	if closeErr := readBuffer.CloseContext("value"); closeErr != nil {
		return nil, closeErr
	}

	// Simple Field (closingTag)
	if pullErr := readBuffer.PullContext("closingTag"); pullErr != nil {
		return nil, pullErr
	}
	_closingTag, _closingTagErr := BACnetContextTagParse(readBuffer, uint8(tagNumber), BACnetDataType(BACnetDataType_CLOSING_TAG))
	if _closingTagErr != nil {
		return nil, errors.Wrap(_closingTagErr, "Error parsing 'closingTag' field")
	}
	closingTag := CastBACnetClosingTag(_closingTag)
	if closeErr := readBuffer.CloseContext("closingTag"); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("BACnetDeviceObjectPropertyReferenceEnclosed"); closeErr != nil {
		return nil, closeErr
	}

	// Create the instance
	return NewBACnetDeviceObjectPropertyReferenceEnclosed(openingTag, value, closingTag, tagNumber), nil
}

func (m *BACnetDeviceObjectPropertyReferenceEnclosed) Serialize(writeBuffer utils.WriteBuffer) error {
	if pushErr := writeBuffer.PushContext("BACnetDeviceObjectPropertyReferenceEnclosed"); pushErr != nil {
		return pushErr
	}

	// Simple Field (openingTag)
	if pushErr := writeBuffer.PushContext("openingTag"); pushErr != nil {
		return pushErr
	}
	_openingTagErr := m.OpeningTag.Serialize(writeBuffer)
	if popErr := writeBuffer.PopContext("openingTag"); popErr != nil {
		return popErr
	}
	if _openingTagErr != nil {
		return errors.Wrap(_openingTagErr, "Error serializing 'openingTag' field")
	}

	// Simple Field (value)
	if pushErr := writeBuffer.PushContext("value"); pushErr != nil {
		return pushErr
	}
	_valueErr := m.Value.Serialize(writeBuffer)
	if popErr := writeBuffer.PopContext("value"); popErr != nil {
		return popErr
	}
	if _valueErr != nil {
		return errors.Wrap(_valueErr, "Error serializing 'value' field")
	}

	// Simple Field (closingTag)
	if pushErr := writeBuffer.PushContext("closingTag"); pushErr != nil {
		return pushErr
	}
	_closingTagErr := m.ClosingTag.Serialize(writeBuffer)
	if popErr := writeBuffer.PopContext("closingTag"); popErr != nil {
		return popErr
	}
	if _closingTagErr != nil {
		return errors.Wrap(_closingTagErr, "Error serializing 'closingTag' field")
	}

	if popErr := writeBuffer.PopContext("BACnetDeviceObjectPropertyReferenceEnclosed"); popErr != nil {
		return popErr
	}
	return nil
}

func (m *BACnetDeviceObjectPropertyReferenceEnclosed) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
