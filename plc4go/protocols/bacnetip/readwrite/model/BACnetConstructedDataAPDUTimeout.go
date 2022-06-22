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

// BACnetConstructedDataAPDUTimeout is the corresponding interface of BACnetConstructedDataAPDUTimeout
type BACnetConstructedDataAPDUTimeout interface {
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetApduTimeout returns ApduTimeout (property field)
	GetApduTimeout() BACnetApplicationTagUnsignedInteger
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagUnsignedInteger
}

// _BACnetConstructedDataAPDUTimeout is the data-structure of this message
type _BACnetConstructedDataAPDUTimeout struct {
	*_BACnetConstructedData
	ApduTimeout BACnetApplicationTagUnsignedInteger

	// Arguments.
	TagNumber          uint8
	ArrayIndexArgument BACnetTagPayloadUnsignedInteger
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataAPDUTimeout) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataAPDUTimeout) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_APDU_TIMEOUT
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataAPDUTimeout) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataAPDUTimeout) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataAPDUTimeout) GetApduTimeout() BACnetApplicationTagUnsignedInteger {
	return m.ApduTimeout
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataAPDUTimeout) GetActualValue() BACnetApplicationTagUnsignedInteger {
	return CastBACnetApplicationTagUnsignedInteger(m.GetApduTimeout())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataAPDUTimeout factory function for _BACnetConstructedDataAPDUTimeout
func NewBACnetConstructedDataAPDUTimeout(apduTimeout BACnetApplicationTagUnsignedInteger, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataAPDUTimeout {
	_result := &_BACnetConstructedDataAPDUTimeout{
		ApduTimeout:            apduTimeout,
		_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataAPDUTimeout(structType interface{}) BACnetConstructedDataAPDUTimeout {
	if casted, ok := structType.(BACnetConstructedDataAPDUTimeout); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataAPDUTimeout); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataAPDUTimeout) GetTypeName() string {
	return "BACnetConstructedDataAPDUTimeout"
}

func (m *_BACnetConstructedDataAPDUTimeout) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetConstructedDataAPDUTimeout) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (apduTimeout)
	lengthInBits += m.ApduTimeout.GetLengthInBits()

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataAPDUTimeout) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataAPDUTimeoutParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataAPDUTimeout, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataAPDUTimeout"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataAPDUTimeout")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (apduTimeout)
	if pullErr := readBuffer.PullContext("apduTimeout"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for apduTimeout")
	}
	_apduTimeout, _apduTimeoutErr := BACnetApplicationTagParse(readBuffer)
	if _apduTimeoutErr != nil {
		return nil, errors.Wrap(_apduTimeoutErr, "Error parsing 'apduTimeout' field")
	}
	apduTimeout := _apduTimeout.(BACnetApplicationTagUnsignedInteger)
	if closeErr := readBuffer.CloseContext("apduTimeout"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for apduTimeout")
	}

	// Virtual field
	_actualValue := apduTimeout
	actualValue := _actualValue
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataAPDUTimeout"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataAPDUTimeout")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataAPDUTimeout{
		ApduTimeout:            apduTimeout,
		_BACnetConstructedData: &_BACnetConstructedData{},
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataAPDUTimeout) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataAPDUTimeout"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataAPDUTimeout")
		}

		// Simple Field (apduTimeout)
		if pushErr := writeBuffer.PushContext("apduTimeout"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for apduTimeout")
		}
		_apduTimeoutErr := writeBuffer.WriteSerializable(m.GetApduTimeout())
		if popErr := writeBuffer.PopContext("apduTimeout"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for apduTimeout")
		}
		if _apduTimeoutErr != nil {
			return errors.Wrap(_apduTimeoutErr, "Error serializing 'apduTimeout' field")
		}
		// Virtual field
		if _actualValueErr := writeBuffer.WriteVirtual("actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataAPDUTimeout"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataAPDUTimeout")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataAPDUTimeout) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
