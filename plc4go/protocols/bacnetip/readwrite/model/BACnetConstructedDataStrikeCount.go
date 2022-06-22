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

// BACnetConstructedDataStrikeCount is the corresponding interface of BACnetConstructedDataStrikeCount
type BACnetConstructedDataStrikeCount interface {
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetStrikeCount returns StrikeCount (property field)
	GetStrikeCount() BACnetApplicationTagUnsignedInteger
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagUnsignedInteger
}

// _BACnetConstructedDataStrikeCount is the data-structure of this message
type _BACnetConstructedDataStrikeCount struct {
	*_BACnetConstructedData
	StrikeCount BACnetApplicationTagUnsignedInteger

	// Arguments.
	TagNumber          uint8
	ArrayIndexArgument BACnetTagPayloadUnsignedInteger
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataStrikeCount) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataStrikeCount) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_STRIKE_COUNT
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataStrikeCount) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataStrikeCount) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataStrikeCount) GetStrikeCount() BACnetApplicationTagUnsignedInteger {
	return m.StrikeCount
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataStrikeCount) GetActualValue() BACnetApplicationTagUnsignedInteger {
	return CastBACnetApplicationTagUnsignedInteger(m.GetStrikeCount())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataStrikeCount factory function for _BACnetConstructedDataStrikeCount
func NewBACnetConstructedDataStrikeCount(strikeCount BACnetApplicationTagUnsignedInteger, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataStrikeCount {
	_result := &_BACnetConstructedDataStrikeCount{
		StrikeCount:            strikeCount,
		_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataStrikeCount(structType interface{}) BACnetConstructedDataStrikeCount {
	if casted, ok := structType.(BACnetConstructedDataStrikeCount); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataStrikeCount); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataStrikeCount) GetTypeName() string {
	return "BACnetConstructedDataStrikeCount"
}

func (m *_BACnetConstructedDataStrikeCount) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetConstructedDataStrikeCount) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (strikeCount)
	lengthInBits += m.StrikeCount.GetLengthInBits()

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataStrikeCount) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataStrikeCountParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataStrikeCount, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataStrikeCount"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataStrikeCount")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (strikeCount)
	if pullErr := readBuffer.PullContext("strikeCount"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for strikeCount")
	}
	_strikeCount, _strikeCountErr := BACnetApplicationTagParse(readBuffer)
	if _strikeCountErr != nil {
		return nil, errors.Wrap(_strikeCountErr, "Error parsing 'strikeCount' field")
	}
	strikeCount := _strikeCount.(BACnetApplicationTagUnsignedInteger)
	if closeErr := readBuffer.CloseContext("strikeCount"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for strikeCount")
	}

	// Virtual field
	_actualValue := strikeCount
	actualValue := _actualValue
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataStrikeCount"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataStrikeCount")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataStrikeCount{
		StrikeCount:            strikeCount,
		_BACnetConstructedData: &_BACnetConstructedData{},
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataStrikeCount) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataStrikeCount"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataStrikeCount")
		}

		// Simple Field (strikeCount)
		if pushErr := writeBuffer.PushContext("strikeCount"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for strikeCount")
		}
		_strikeCountErr := writeBuffer.WriteSerializable(m.GetStrikeCount())
		if popErr := writeBuffer.PopContext("strikeCount"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for strikeCount")
		}
		if _strikeCountErr != nil {
			return errors.Wrap(_strikeCountErr, "Error serializing 'strikeCount' field")
		}
		// Virtual field
		if _actualValueErr := writeBuffer.WriteVirtual("actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataStrikeCount"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataStrikeCount")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataStrikeCount) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
