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

// BACnetConstructedDataMusterPoint is the corresponding interface of BACnetConstructedDataMusterPoint
type BACnetConstructedDataMusterPoint interface {
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetMusterPoint returns MusterPoint (property field)
	GetMusterPoint() BACnetApplicationTagBoolean
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagBoolean
}

// _BACnetConstructedDataMusterPoint is the data-structure of this message
type _BACnetConstructedDataMusterPoint struct {
	*_BACnetConstructedData
	MusterPoint BACnetApplicationTagBoolean

	// Arguments.
	TagNumber          uint8
	ArrayIndexArgument BACnetTagPayloadUnsignedInteger
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataMusterPoint) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataMusterPoint) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_MUSTER_POINT
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataMusterPoint) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataMusterPoint) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataMusterPoint) GetMusterPoint() BACnetApplicationTagBoolean {
	return m.MusterPoint
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataMusterPoint) GetActualValue() BACnetApplicationTagBoolean {
	return CastBACnetApplicationTagBoolean(m.GetMusterPoint())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataMusterPoint factory function for _BACnetConstructedDataMusterPoint
func NewBACnetConstructedDataMusterPoint(musterPoint BACnetApplicationTagBoolean, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataMusterPoint {
	_result := &_BACnetConstructedDataMusterPoint{
		MusterPoint:            musterPoint,
		_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataMusterPoint(structType interface{}) BACnetConstructedDataMusterPoint {
	if casted, ok := structType.(BACnetConstructedDataMusterPoint); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataMusterPoint); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataMusterPoint) GetTypeName() string {
	return "BACnetConstructedDataMusterPoint"
}

func (m *_BACnetConstructedDataMusterPoint) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetConstructedDataMusterPoint) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (musterPoint)
	lengthInBits += m.MusterPoint.GetLengthInBits()

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataMusterPoint) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataMusterPointParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataMusterPoint, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataMusterPoint"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataMusterPoint")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (musterPoint)
	if pullErr := readBuffer.PullContext("musterPoint"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for musterPoint")
	}
	_musterPoint, _musterPointErr := BACnetApplicationTagParse(readBuffer)
	if _musterPointErr != nil {
		return nil, errors.Wrap(_musterPointErr, "Error parsing 'musterPoint' field")
	}
	musterPoint := _musterPoint.(BACnetApplicationTagBoolean)
	if closeErr := readBuffer.CloseContext("musterPoint"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for musterPoint")
	}

	// Virtual field
	_actualValue := musterPoint
	actualValue := _actualValue
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataMusterPoint"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataMusterPoint")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataMusterPoint{
		MusterPoint:            musterPoint,
		_BACnetConstructedData: &_BACnetConstructedData{},
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataMusterPoint) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataMusterPoint"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataMusterPoint")
		}

		// Simple Field (musterPoint)
		if pushErr := writeBuffer.PushContext("musterPoint"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for musterPoint")
		}
		_musterPointErr := writeBuffer.WriteSerializable(m.GetMusterPoint())
		if popErr := writeBuffer.PopContext("musterPoint"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for musterPoint")
		}
		if _musterPointErr != nil {
			return errors.Wrap(_musterPointErr, "Error serializing 'musterPoint' field")
		}
		// Virtual field
		if _actualValueErr := writeBuffer.WriteVirtual("actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataMusterPoint"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataMusterPoint")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataMusterPoint) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
