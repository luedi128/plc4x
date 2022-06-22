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

// BACnetNotificationParametersChangeOfDiscreteValueNewValueBoolean is the corresponding interface of BACnetNotificationParametersChangeOfDiscreteValueNewValueBoolean
type BACnetNotificationParametersChangeOfDiscreteValueNewValueBoolean interface {
	utils.LengthAware
	utils.Serializable
	BACnetNotificationParametersChangeOfDiscreteValueNewValue
	// GetBooleanValue returns BooleanValue (property field)
	GetBooleanValue() BACnetApplicationTagBoolean
}

// _BACnetNotificationParametersChangeOfDiscreteValueNewValueBoolean is the data-structure of this message
type _BACnetNotificationParametersChangeOfDiscreteValueNewValueBoolean struct {
	*_BACnetNotificationParametersChangeOfDiscreteValueNewValue
	BooleanValue BACnetApplicationTagBoolean

	// Arguments.
	TagNumber uint8
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetNotificationParametersChangeOfDiscreteValueNewValueBoolean) InitializeParent(parent BACnetNotificationParametersChangeOfDiscreteValueNewValue, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetNotificationParametersChangeOfDiscreteValueNewValueBoolean) GetParent() BACnetNotificationParametersChangeOfDiscreteValueNewValue {
	return m._BACnetNotificationParametersChangeOfDiscreteValueNewValue
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetNotificationParametersChangeOfDiscreteValueNewValueBoolean) GetBooleanValue() BACnetApplicationTagBoolean {
	return m.BooleanValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetNotificationParametersChangeOfDiscreteValueNewValueBoolean factory function for _BACnetNotificationParametersChangeOfDiscreteValueNewValueBoolean
func NewBACnetNotificationParametersChangeOfDiscreteValueNewValueBoolean(booleanValue BACnetApplicationTagBoolean, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8) *_BACnetNotificationParametersChangeOfDiscreteValueNewValueBoolean {
	_result := &_BACnetNotificationParametersChangeOfDiscreteValueNewValueBoolean{
		BooleanValue: booleanValue,
		_BACnetNotificationParametersChangeOfDiscreteValueNewValue: NewBACnetNotificationParametersChangeOfDiscreteValueNewValue(openingTag, peekedTagHeader, closingTag, tagNumber),
	}
	_result._BACnetNotificationParametersChangeOfDiscreteValueNewValue._BACnetNotificationParametersChangeOfDiscreteValueNewValueChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetNotificationParametersChangeOfDiscreteValueNewValueBoolean(structType interface{}) BACnetNotificationParametersChangeOfDiscreteValueNewValueBoolean {
	if casted, ok := structType.(BACnetNotificationParametersChangeOfDiscreteValueNewValueBoolean); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetNotificationParametersChangeOfDiscreteValueNewValueBoolean); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetNotificationParametersChangeOfDiscreteValueNewValueBoolean) GetTypeName() string {
	return "BACnetNotificationParametersChangeOfDiscreteValueNewValueBoolean"
}

func (m *_BACnetNotificationParametersChangeOfDiscreteValueNewValueBoolean) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetNotificationParametersChangeOfDiscreteValueNewValueBoolean) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (booleanValue)
	lengthInBits += m.BooleanValue.GetLengthInBits()

	return lengthInBits
}

func (m *_BACnetNotificationParametersChangeOfDiscreteValueNewValueBoolean) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetNotificationParametersChangeOfDiscreteValueNewValueBooleanParse(readBuffer utils.ReadBuffer, tagNumber uint8) (BACnetNotificationParametersChangeOfDiscreteValueNewValueBoolean, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetNotificationParametersChangeOfDiscreteValueNewValueBoolean"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetNotificationParametersChangeOfDiscreteValueNewValueBoolean")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (booleanValue)
	if pullErr := readBuffer.PullContext("booleanValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for booleanValue")
	}
	_booleanValue, _booleanValueErr := BACnetApplicationTagParse(readBuffer)
	if _booleanValueErr != nil {
		return nil, errors.Wrap(_booleanValueErr, "Error parsing 'booleanValue' field")
	}
	booleanValue := _booleanValue.(BACnetApplicationTagBoolean)
	if closeErr := readBuffer.CloseContext("booleanValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for booleanValue")
	}

	if closeErr := readBuffer.CloseContext("BACnetNotificationParametersChangeOfDiscreteValueNewValueBoolean"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetNotificationParametersChangeOfDiscreteValueNewValueBoolean")
	}

	// Create a partially initialized instance
	_child := &_BACnetNotificationParametersChangeOfDiscreteValueNewValueBoolean{
		BooleanValue: booleanValue,
		_BACnetNotificationParametersChangeOfDiscreteValueNewValue: &_BACnetNotificationParametersChangeOfDiscreteValueNewValue{},
	}
	_child._BACnetNotificationParametersChangeOfDiscreteValueNewValue._BACnetNotificationParametersChangeOfDiscreteValueNewValueChildRequirements = _child
	return _child, nil
}

func (m *_BACnetNotificationParametersChangeOfDiscreteValueNewValueBoolean) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetNotificationParametersChangeOfDiscreteValueNewValueBoolean"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetNotificationParametersChangeOfDiscreteValueNewValueBoolean")
		}

		// Simple Field (booleanValue)
		if pushErr := writeBuffer.PushContext("booleanValue"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for booleanValue")
		}
		_booleanValueErr := writeBuffer.WriteSerializable(m.GetBooleanValue())
		if popErr := writeBuffer.PopContext("booleanValue"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for booleanValue")
		}
		if _booleanValueErr != nil {
			return errors.Wrap(_booleanValueErr, "Error serializing 'booleanValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetNotificationParametersChangeOfDiscreteValueNewValueBoolean"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetNotificationParametersChangeOfDiscreteValueNewValueBoolean")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetNotificationParametersChangeOfDiscreteValueNewValueBoolean) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
