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

// BACnetConstructedDataTimerResolution is the corresponding interface of BACnetConstructedDataTimerResolution
type BACnetConstructedDataTimerResolution interface {
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetResolution returns Resolution (property field)
	GetResolution() BACnetApplicationTagUnsignedInteger
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagUnsignedInteger
}

// _BACnetConstructedDataTimerResolution is the data-structure of this message
type _BACnetConstructedDataTimerResolution struct {
	*_BACnetConstructedData
	Resolution BACnetApplicationTagUnsignedInteger

	// Arguments.
	TagNumber          uint8
	ArrayIndexArgument BACnetTagPayloadUnsignedInteger
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataTimerResolution) GetObjectTypeArgument() BACnetObjectType {
	return BACnetObjectType_TIMER
}

func (m *_BACnetConstructedDataTimerResolution) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_RESOLUTION
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataTimerResolution) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataTimerResolution) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataTimerResolution) GetResolution() BACnetApplicationTagUnsignedInteger {
	return m.Resolution
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataTimerResolution) GetActualValue() BACnetApplicationTagUnsignedInteger {
	return CastBACnetApplicationTagUnsignedInteger(m.GetResolution())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataTimerResolution factory function for _BACnetConstructedDataTimerResolution
func NewBACnetConstructedDataTimerResolution(resolution BACnetApplicationTagUnsignedInteger, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataTimerResolution {
	_result := &_BACnetConstructedDataTimerResolution{
		Resolution:             resolution,
		_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataTimerResolution(structType interface{}) BACnetConstructedDataTimerResolution {
	if casted, ok := structType.(BACnetConstructedDataTimerResolution); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataTimerResolution); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataTimerResolution) GetTypeName() string {
	return "BACnetConstructedDataTimerResolution"
}

func (m *_BACnetConstructedDataTimerResolution) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetConstructedDataTimerResolution) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (resolution)
	lengthInBits += m.Resolution.GetLengthInBits()

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataTimerResolution) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataTimerResolutionParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataTimerResolution, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataTimerResolution"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataTimerResolution")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (resolution)
	if pullErr := readBuffer.PullContext("resolution"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for resolution")
	}
	_resolution, _resolutionErr := BACnetApplicationTagParse(readBuffer)
	if _resolutionErr != nil {
		return nil, errors.Wrap(_resolutionErr, "Error parsing 'resolution' field")
	}
	resolution := _resolution.(BACnetApplicationTagUnsignedInteger)
	if closeErr := readBuffer.CloseContext("resolution"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for resolution")
	}

	// Virtual field
	_actualValue := resolution
	actualValue := _actualValue
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataTimerResolution"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataTimerResolution")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataTimerResolution{
		Resolution:             resolution,
		_BACnetConstructedData: &_BACnetConstructedData{},
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataTimerResolution) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataTimerResolution"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataTimerResolution")
		}

		// Simple Field (resolution)
		if pushErr := writeBuffer.PushContext("resolution"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for resolution")
		}
		_resolutionErr := writeBuffer.WriteSerializable(m.GetResolution())
		if popErr := writeBuffer.PopContext("resolution"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for resolution")
		}
		if _resolutionErr != nil {
			return errors.Wrap(_resolutionErr, "Error serializing 'resolution' field")
		}
		// Virtual field
		if _actualValueErr := writeBuffer.WriteVirtual("actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataTimerResolution"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataTimerResolution")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataTimerResolution) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
