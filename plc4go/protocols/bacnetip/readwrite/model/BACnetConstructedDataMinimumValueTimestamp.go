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

// BACnetConstructedDataMinimumValueTimestamp is the corresponding interface of BACnetConstructedDataMinimumValueTimestamp
type BACnetConstructedDataMinimumValueTimestamp interface {
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetMinimumValueTimestamp returns MinimumValueTimestamp (property field)
	GetMinimumValueTimestamp() BACnetDateTime
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetDateTime
}

// _BACnetConstructedDataMinimumValueTimestamp is the data-structure of this message
type _BACnetConstructedDataMinimumValueTimestamp struct {
	*_BACnetConstructedData
	MinimumValueTimestamp BACnetDateTime

	// Arguments.
	TagNumber          uint8
	ArrayIndexArgument BACnetTagPayloadUnsignedInteger
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataMinimumValueTimestamp) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataMinimumValueTimestamp) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_MINIMUM_VALUE_TIMESTAMP
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataMinimumValueTimestamp) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataMinimumValueTimestamp) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataMinimumValueTimestamp) GetMinimumValueTimestamp() BACnetDateTime {
	return m.MinimumValueTimestamp
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataMinimumValueTimestamp) GetActualValue() BACnetDateTime {
	return CastBACnetDateTime(m.GetMinimumValueTimestamp())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataMinimumValueTimestamp factory function for _BACnetConstructedDataMinimumValueTimestamp
func NewBACnetConstructedDataMinimumValueTimestamp(minimumValueTimestamp BACnetDateTime, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataMinimumValueTimestamp {
	_result := &_BACnetConstructedDataMinimumValueTimestamp{
		MinimumValueTimestamp:  minimumValueTimestamp,
		_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataMinimumValueTimestamp(structType interface{}) BACnetConstructedDataMinimumValueTimestamp {
	if casted, ok := structType.(BACnetConstructedDataMinimumValueTimestamp); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataMinimumValueTimestamp); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataMinimumValueTimestamp) GetTypeName() string {
	return "BACnetConstructedDataMinimumValueTimestamp"
}

func (m *_BACnetConstructedDataMinimumValueTimestamp) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetConstructedDataMinimumValueTimestamp) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (minimumValueTimestamp)
	lengthInBits += m.MinimumValueTimestamp.GetLengthInBits()

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataMinimumValueTimestamp) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataMinimumValueTimestampParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataMinimumValueTimestamp, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataMinimumValueTimestamp"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataMinimumValueTimestamp")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (minimumValueTimestamp)
	if pullErr := readBuffer.PullContext("minimumValueTimestamp"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for minimumValueTimestamp")
	}
	_minimumValueTimestamp, _minimumValueTimestampErr := BACnetDateTimeParse(readBuffer)
	if _minimumValueTimestampErr != nil {
		return nil, errors.Wrap(_minimumValueTimestampErr, "Error parsing 'minimumValueTimestamp' field")
	}
	minimumValueTimestamp := _minimumValueTimestamp.(BACnetDateTime)
	if closeErr := readBuffer.CloseContext("minimumValueTimestamp"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for minimumValueTimestamp")
	}

	// Virtual field
	_actualValue := minimumValueTimestamp
	actualValue := _actualValue
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataMinimumValueTimestamp"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataMinimumValueTimestamp")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataMinimumValueTimestamp{
		MinimumValueTimestamp:  minimumValueTimestamp,
		_BACnetConstructedData: &_BACnetConstructedData{},
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataMinimumValueTimestamp) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataMinimumValueTimestamp"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataMinimumValueTimestamp")
		}

		// Simple Field (minimumValueTimestamp)
		if pushErr := writeBuffer.PushContext("minimumValueTimestamp"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for minimumValueTimestamp")
		}
		_minimumValueTimestampErr := writeBuffer.WriteSerializable(m.GetMinimumValueTimestamp())
		if popErr := writeBuffer.PopContext("minimumValueTimestamp"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for minimumValueTimestamp")
		}
		if _minimumValueTimestampErr != nil {
			return errors.Wrap(_minimumValueTimestampErr, "Error serializing 'minimumValueTimestamp' field")
		}
		// Virtual field
		if _actualValueErr := writeBuffer.WriteVirtual("actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataMinimumValueTimestamp"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataMinimumValueTimestamp")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataMinimumValueTimestamp) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
