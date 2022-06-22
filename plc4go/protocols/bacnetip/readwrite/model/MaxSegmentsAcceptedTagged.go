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

// MaxSegmentsAcceptedTagged is the corresponding interface of MaxSegmentsAcceptedTagged
type MaxSegmentsAcceptedTagged interface {
	utils.LengthAware
	utils.Serializable
	// GetHeader returns Header (property field)
	GetHeader() BACnetTagHeader
	// GetValue returns Value (property field)
	GetValue() MaxSegmentsAccepted
}

// _MaxSegmentsAcceptedTagged is the data-structure of this message
type _MaxSegmentsAcceptedTagged struct {
	Header BACnetTagHeader
	Value  MaxSegmentsAccepted

	// Arguments.
	TagNumber uint8
	TagClass  TagClass
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_MaxSegmentsAcceptedTagged) GetHeader() BACnetTagHeader {
	return m.Header
}

func (m *_MaxSegmentsAcceptedTagged) GetValue() MaxSegmentsAccepted {
	return m.Value
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewMaxSegmentsAcceptedTagged factory function for _MaxSegmentsAcceptedTagged
func NewMaxSegmentsAcceptedTagged(header BACnetTagHeader, value MaxSegmentsAccepted, tagNumber uint8, tagClass TagClass) *_MaxSegmentsAcceptedTagged {
	return &_MaxSegmentsAcceptedTagged{Header: header, Value: value, TagNumber: tagNumber, TagClass: tagClass}
}

// Deprecated: use the interface for direct cast
func CastMaxSegmentsAcceptedTagged(structType interface{}) MaxSegmentsAcceptedTagged {
	if casted, ok := structType.(MaxSegmentsAcceptedTagged); ok {
		return casted
	}
	if casted, ok := structType.(*MaxSegmentsAcceptedTagged); ok {
		return *casted
	}
	return nil
}

func (m *_MaxSegmentsAcceptedTagged) GetTypeName() string {
	return "MaxSegmentsAcceptedTagged"
}

func (m *_MaxSegmentsAcceptedTagged) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_MaxSegmentsAcceptedTagged) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(0)

	// Simple field (header)
	lengthInBits += m.Header.GetLengthInBits()

	// Manual Field (value)
	lengthInBits += uint16(int32(m.GetHeader().GetActualLength()) * int32(int32(8)))

	return lengthInBits
}

func (m *_MaxSegmentsAcceptedTagged) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func MaxSegmentsAcceptedTaggedParse(readBuffer utils.ReadBuffer, tagNumber uint8, tagClass TagClass) (MaxSegmentsAcceptedTagged, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("MaxSegmentsAcceptedTagged"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for MaxSegmentsAcceptedTagged")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (header)
	if pullErr := readBuffer.PullContext("header"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for header")
	}
	_header, _headerErr := BACnetTagHeaderParse(readBuffer)
	if _headerErr != nil {
		return nil, errors.Wrap(_headerErr, "Error parsing 'header' field")
	}
	header := _header.(BACnetTagHeader)
	if closeErr := readBuffer.CloseContext("header"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for header")
	}

	// Validation
	if !(bool((header.GetTagClass()) == (tagClass))) {
		return nil, errors.WithStack(utils.ParseValidationError{"tag class doesn't match"})
	}

	// Validation
	if !(bool(bool(bool((header.GetTagClass()) == (TagClass_APPLICATION_TAGS)))) || bool(bool(bool((header.GetActualTagNumber()) == (tagNumber))))) {
		return nil, errors.WithStack(utils.ParseAssertError{"tagnumber doesn't match"})
	}

	// Manual Field (value)
	_value, _valueErr := ReadEnumGenericFailing(readBuffer, header.GetActualLength(), MaxSegmentsAccepted_UNSPECIFIED)
	if _valueErr != nil {
		return nil, errors.Wrap(_valueErr, "Error parsing 'value' field")
	}
	value := _value.(MaxSegmentsAccepted)

	if closeErr := readBuffer.CloseContext("MaxSegmentsAcceptedTagged"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for MaxSegmentsAcceptedTagged")
	}

	// Create the instance
	return NewMaxSegmentsAcceptedTagged(header, value, tagNumber, tagClass), nil
}

func (m *_MaxSegmentsAcceptedTagged) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("MaxSegmentsAcceptedTagged"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for MaxSegmentsAcceptedTagged")
	}

	// Simple Field (header)
	if pushErr := writeBuffer.PushContext("header"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for header")
	}
	_headerErr := writeBuffer.WriteSerializable(m.GetHeader())
	if popErr := writeBuffer.PopContext("header"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for header")
	}
	if _headerErr != nil {
		return errors.Wrap(_headerErr, "Error serializing 'header' field")
	}

	// Manual Field (value)
	_valueErr := WriteEnumGeneric(writeBuffer, m.GetValue())
	if _valueErr != nil {
		return errors.Wrap(_valueErr, "Error serializing 'value' field")
	}

	if popErr := writeBuffer.PopContext("MaxSegmentsAcceptedTagged"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for MaxSegmentsAcceptedTagged")
	}
	return nil
}

func (m *_MaxSegmentsAcceptedTagged) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
