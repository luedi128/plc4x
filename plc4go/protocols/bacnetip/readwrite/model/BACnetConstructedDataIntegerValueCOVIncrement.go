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

// BACnetConstructedDataIntegerValueCOVIncrement is the corresponding interface of BACnetConstructedDataIntegerValueCOVIncrement
type BACnetConstructedDataIntegerValueCOVIncrement interface {
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetCovIncrement returns CovIncrement (property field)
	GetCovIncrement() BACnetApplicationTagUnsignedInteger
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagUnsignedInteger
}

// _BACnetConstructedDataIntegerValueCOVIncrement is the data-structure of this message
type _BACnetConstructedDataIntegerValueCOVIncrement struct {
	*_BACnetConstructedData
	CovIncrement BACnetApplicationTagUnsignedInteger

	// Arguments.
	TagNumber          uint8
	ArrayIndexArgument BACnetTagPayloadUnsignedInteger
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataIntegerValueCOVIncrement) GetObjectTypeArgument() BACnetObjectType {
	return BACnetObjectType_INTEGER_VALUE
}

func (m *_BACnetConstructedDataIntegerValueCOVIncrement) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_COV_INCREMENT
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataIntegerValueCOVIncrement) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataIntegerValueCOVIncrement) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataIntegerValueCOVIncrement) GetCovIncrement() BACnetApplicationTagUnsignedInteger {
	return m.CovIncrement
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataIntegerValueCOVIncrement) GetActualValue() BACnetApplicationTagUnsignedInteger {
	return CastBACnetApplicationTagUnsignedInteger(m.GetCovIncrement())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataIntegerValueCOVIncrement factory function for _BACnetConstructedDataIntegerValueCOVIncrement
func NewBACnetConstructedDataIntegerValueCOVIncrement(covIncrement BACnetApplicationTagUnsignedInteger, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataIntegerValueCOVIncrement {
	_result := &_BACnetConstructedDataIntegerValueCOVIncrement{
		CovIncrement:           covIncrement,
		_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataIntegerValueCOVIncrement(structType interface{}) BACnetConstructedDataIntegerValueCOVIncrement {
	if casted, ok := structType.(BACnetConstructedDataIntegerValueCOVIncrement); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataIntegerValueCOVIncrement); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataIntegerValueCOVIncrement) GetTypeName() string {
	return "BACnetConstructedDataIntegerValueCOVIncrement"
}

func (m *_BACnetConstructedDataIntegerValueCOVIncrement) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetConstructedDataIntegerValueCOVIncrement) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (covIncrement)
	lengthInBits += m.CovIncrement.GetLengthInBits()

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataIntegerValueCOVIncrement) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataIntegerValueCOVIncrementParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataIntegerValueCOVIncrement, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataIntegerValueCOVIncrement"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataIntegerValueCOVIncrement")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (covIncrement)
	if pullErr := readBuffer.PullContext("covIncrement"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for covIncrement")
	}
	_covIncrement, _covIncrementErr := BACnetApplicationTagParse(readBuffer)
	if _covIncrementErr != nil {
		return nil, errors.Wrap(_covIncrementErr, "Error parsing 'covIncrement' field")
	}
	covIncrement := _covIncrement.(BACnetApplicationTagUnsignedInteger)
	if closeErr := readBuffer.CloseContext("covIncrement"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for covIncrement")
	}

	// Virtual field
	_actualValue := covIncrement
	actualValue := _actualValue
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataIntegerValueCOVIncrement"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataIntegerValueCOVIncrement")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataIntegerValueCOVIncrement{
		CovIncrement:           covIncrement,
		_BACnetConstructedData: &_BACnetConstructedData{},
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataIntegerValueCOVIncrement) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataIntegerValueCOVIncrement"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataIntegerValueCOVIncrement")
		}

		// Simple Field (covIncrement)
		if pushErr := writeBuffer.PushContext("covIncrement"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for covIncrement")
		}
		_covIncrementErr := writeBuffer.WriteSerializable(m.GetCovIncrement())
		if popErr := writeBuffer.PopContext("covIncrement"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for covIncrement")
		}
		if _covIncrementErr != nil {
			return errors.Wrap(_covIncrementErr, "Error serializing 'covIncrement' field")
		}
		// Virtual field
		if _actualValueErr := writeBuffer.WriteVirtual("actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataIntegerValueCOVIncrement"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataIntegerValueCOVIncrement")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataIntegerValueCOVIncrement) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
