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

// BACnetConstructedDataNumberOfStates is the corresponding interface of BACnetConstructedDataNumberOfStates
type BACnetConstructedDataNumberOfStates interface {
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetNumberOfState returns NumberOfState (property field)
	GetNumberOfState() BACnetApplicationTagUnsignedInteger
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagUnsignedInteger
}

// _BACnetConstructedDataNumberOfStates is the data-structure of this message
type _BACnetConstructedDataNumberOfStates struct {
	*_BACnetConstructedData
	NumberOfState BACnetApplicationTagUnsignedInteger

	// Arguments.
	TagNumber          uint8
	ArrayIndexArgument BACnetTagPayloadUnsignedInteger
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataNumberOfStates) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataNumberOfStates) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_NUMBER_OF_STATES
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataNumberOfStates) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataNumberOfStates) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataNumberOfStates) GetNumberOfState() BACnetApplicationTagUnsignedInteger {
	return m.NumberOfState
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataNumberOfStates) GetActualValue() BACnetApplicationTagUnsignedInteger {
	return CastBACnetApplicationTagUnsignedInteger(m.GetNumberOfState())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataNumberOfStates factory function for _BACnetConstructedDataNumberOfStates
func NewBACnetConstructedDataNumberOfStates(numberOfState BACnetApplicationTagUnsignedInteger, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataNumberOfStates {
	_result := &_BACnetConstructedDataNumberOfStates{
		NumberOfState:          numberOfState,
		_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataNumberOfStates(structType interface{}) BACnetConstructedDataNumberOfStates {
	if casted, ok := structType.(BACnetConstructedDataNumberOfStates); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataNumberOfStates); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataNumberOfStates) GetTypeName() string {
	return "BACnetConstructedDataNumberOfStates"
}

func (m *_BACnetConstructedDataNumberOfStates) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetConstructedDataNumberOfStates) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (numberOfState)
	lengthInBits += m.NumberOfState.GetLengthInBits()

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataNumberOfStates) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataNumberOfStatesParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataNumberOfStates, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataNumberOfStates"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataNumberOfStates")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (numberOfState)
	if pullErr := readBuffer.PullContext("numberOfState"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for numberOfState")
	}
	_numberOfState, _numberOfStateErr := BACnetApplicationTagParse(readBuffer)
	if _numberOfStateErr != nil {
		return nil, errors.Wrap(_numberOfStateErr, "Error parsing 'numberOfState' field")
	}
	numberOfState := _numberOfState.(BACnetApplicationTagUnsignedInteger)
	if closeErr := readBuffer.CloseContext("numberOfState"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for numberOfState")
	}

	// Virtual field
	_actualValue := numberOfState
	actualValue := _actualValue
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataNumberOfStates"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataNumberOfStates")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataNumberOfStates{
		NumberOfState:          numberOfState,
		_BACnetConstructedData: &_BACnetConstructedData{},
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataNumberOfStates) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataNumberOfStates"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataNumberOfStates")
		}

		// Simple Field (numberOfState)
		if pushErr := writeBuffer.PushContext("numberOfState"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for numberOfState")
		}
		_numberOfStateErr := writeBuffer.WriteSerializable(m.GetNumberOfState())
		if popErr := writeBuffer.PopContext("numberOfState"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for numberOfState")
		}
		if _numberOfStateErr != nil {
			return errors.Wrap(_numberOfStateErr, "Error serializing 'numberOfState' field")
		}
		// Virtual field
		if _actualValueErr := writeBuffer.WriteVirtual("actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataNumberOfStates"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataNumberOfStates")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataNumberOfStates) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
