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
	"github.com/rs/zerolog/log"
	"io"
)

// Code generated by code-generation. DO NOT EDIT.

// ListOfCovNotificationsValue is the corresponding interface of ListOfCovNotificationsValue
type ListOfCovNotificationsValue interface {
	utils.LengthAware
	utils.Serializable
	// GetPropertyIdentifier returns PropertyIdentifier (property field)
	GetPropertyIdentifier() BACnetPropertyIdentifierTagged
	// GetArrayIndex returns ArrayIndex (property field)
	GetArrayIndex() BACnetContextTagUnsignedInteger
	// GetPropertyValue returns PropertyValue (property field)
	GetPropertyValue() BACnetConstructedData
	// GetTimeOfChange returns TimeOfChange (property field)
	GetTimeOfChange() BACnetContextTagTime
}

// _ListOfCovNotificationsValue is the data-structure of this message
type _ListOfCovNotificationsValue struct {
	PropertyIdentifier BACnetPropertyIdentifierTagged
	ArrayIndex         BACnetContextTagUnsignedInteger
	PropertyValue      BACnetConstructedData
	TimeOfChange       BACnetContextTagTime

	// Arguments.
	ObjectTypeArgument BACnetObjectType
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_ListOfCovNotificationsValue) GetPropertyIdentifier() BACnetPropertyIdentifierTagged {
	return m.PropertyIdentifier
}

func (m *_ListOfCovNotificationsValue) GetArrayIndex() BACnetContextTagUnsignedInteger {
	return m.ArrayIndex
}

func (m *_ListOfCovNotificationsValue) GetPropertyValue() BACnetConstructedData {
	return m.PropertyValue
}

func (m *_ListOfCovNotificationsValue) GetTimeOfChange() BACnetContextTagTime {
	return m.TimeOfChange
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewListOfCovNotificationsValue factory function for _ListOfCovNotificationsValue
func NewListOfCovNotificationsValue(propertyIdentifier BACnetPropertyIdentifierTagged, arrayIndex BACnetContextTagUnsignedInteger, propertyValue BACnetConstructedData, timeOfChange BACnetContextTagTime, objectTypeArgument BACnetObjectType) *_ListOfCovNotificationsValue {
	return &_ListOfCovNotificationsValue{PropertyIdentifier: propertyIdentifier, ArrayIndex: arrayIndex, PropertyValue: propertyValue, TimeOfChange: timeOfChange, ObjectTypeArgument: objectTypeArgument}
}

// Deprecated: use the interface for direct cast
func CastListOfCovNotificationsValue(structType interface{}) ListOfCovNotificationsValue {
	if casted, ok := structType.(ListOfCovNotificationsValue); ok {
		return casted
	}
	if casted, ok := structType.(*ListOfCovNotificationsValue); ok {
		return *casted
	}
	return nil
}

func (m *_ListOfCovNotificationsValue) GetTypeName() string {
	return "ListOfCovNotificationsValue"
}

func (m *_ListOfCovNotificationsValue) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_ListOfCovNotificationsValue) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(0)

	// Simple field (propertyIdentifier)
	lengthInBits += m.PropertyIdentifier.GetLengthInBits()

	// Optional Field (arrayIndex)
	if m.ArrayIndex != nil {
		lengthInBits += m.ArrayIndex.GetLengthInBits()
	}

	// Simple field (propertyValue)
	lengthInBits += m.PropertyValue.GetLengthInBits()

	// Optional Field (timeOfChange)
	if m.TimeOfChange != nil {
		lengthInBits += m.TimeOfChange.GetLengthInBits()
	}

	return lengthInBits
}

func (m *_ListOfCovNotificationsValue) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func ListOfCovNotificationsValueParse(readBuffer utils.ReadBuffer, objectTypeArgument BACnetObjectType) (ListOfCovNotificationsValue, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ListOfCovNotificationsValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ListOfCovNotificationsValue")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (propertyIdentifier)
	if pullErr := readBuffer.PullContext("propertyIdentifier"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for propertyIdentifier")
	}
	_propertyIdentifier, _propertyIdentifierErr := BACnetPropertyIdentifierTaggedParse(readBuffer, uint8(uint8(0)), TagClass(TagClass_CONTEXT_SPECIFIC_TAGS))
	if _propertyIdentifierErr != nil {
		return nil, errors.Wrap(_propertyIdentifierErr, "Error parsing 'propertyIdentifier' field")
	}
	propertyIdentifier := _propertyIdentifier.(BACnetPropertyIdentifierTagged)
	if closeErr := readBuffer.CloseContext("propertyIdentifier"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for propertyIdentifier")
	}

	// Optional Field (arrayIndex) (Can be skipped, if a given expression evaluates to false)
	var arrayIndex BACnetContextTagUnsignedInteger = nil
	{
		currentPos = positionAware.GetPos()
		if pullErr := readBuffer.PullContext("arrayIndex"); pullErr != nil {
			return nil, errors.Wrap(pullErr, "Error pulling for arrayIndex")
		}
		_val, _err := BACnetContextTagParse(readBuffer, uint8(1), BACnetDataType_UNSIGNED_INTEGER)
		switch {
		case errors.Is(_err, utils.ParseAssertError{}) || errors.Is(_err, io.EOF):
			log.Debug().Err(_err).Msg("Resetting position because optional threw an error")
			readBuffer.Reset(currentPos)
		case _err != nil:
			return nil, errors.Wrap(_err, "Error parsing 'arrayIndex' field")
		default:
			arrayIndex = _val.(BACnetContextTagUnsignedInteger)
			if closeErr := readBuffer.CloseContext("arrayIndex"); closeErr != nil {
				return nil, errors.Wrap(closeErr, "Error closing for arrayIndex")
			}
		}
	}

	// Simple Field (propertyValue)
	if pullErr := readBuffer.PullContext("propertyValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for propertyValue")
	}
	_propertyValue, _propertyValueErr := BACnetConstructedDataParse(readBuffer, uint8(uint8(2)), BACnetObjectType(objectTypeArgument), BACnetPropertyIdentifier(propertyIdentifier.GetValue()), CastBACnetTagPayloadUnsignedInteger(CastBACnetTagPayloadUnsignedInteger(utils.InlineIf(bool((arrayIndex) != (nil)), func() interface{} { return CastBACnetTagPayloadUnsignedInteger((arrayIndex).GetPayload()) }, func() interface{} { return CastBACnetTagPayloadUnsignedInteger(nil) }))))
	if _propertyValueErr != nil {
		return nil, errors.Wrap(_propertyValueErr, "Error parsing 'propertyValue' field")
	}
	propertyValue := _propertyValue.(BACnetConstructedData)
	if closeErr := readBuffer.CloseContext("propertyValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for propertyValue")
	}

	// Optional Field (timeOfChange) (Can be skipped, if a given expression evaluates to false)
	var timeOfChange BACnetContextTagTime = nil
	{
		currentPos = positionAware.GetPos()
		if pullErr := readBuffer.PullContext("timeOfChange"); pullErr != nil {
			return nil, errors.Wrap(pullErr, "Error pulling for timeOfChange")
		}
		_val, _err := BACnetContextTagParse(readBuffer, uint8(3), BACnetDataType_TIME)
		switch {
		case errors.Is(_err, utils.ParseAssertError{}) || errors.Is(_err, io.EOF):
			log.Debug().Err(_err).Msg("Resetting position because optional threw an error")
			readBuffer.Reset(currentPos)
		case _err != nil:
			return nil, errors.Wrap(_err, "Error parsing 'timeOfChange' field")
		default:
			timeOfChange = _val.(BACnetContextTagTime)
			if closeErr := readBuffer.CloseContext("timeOfChange"); closeErr != nil {
				return nil, errors.Wrap(closeErr, "Error closing for timeOfChange")
			}
		}
	}

	if closeErr := readBuffer.CloseContext("ListOfCovNotificationsValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ListOfCovNotificationsValue")
	}

	// Create the instance
	return NewListOfCovNotificationsValue(propertyIdentifier, arrayIndex, propertyValue, timeOfChange, objectTypeArgument), nil
}

func (m *_ListOfCovNotificationsValue) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("ListOfCovNotificationsValue"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for ListOfCovNotificationsValue")
	}

	// Simple Field (propertyIdentifier)
	if pushErr := writeBuffer.PushContext("propertyIdentifier"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for propertyIdentifier")
	}
	_propertyIdentifierErr := writeBuffer.WriteSerializable(m.GetPropertyIdentifier())
	if popErr := writeBuffer.PopContext("propertyIdentifier"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for propertyIdentifier")
	}
	if _propertyIdentifierErr != nil {
		return errors.Wrap(_propertyIdentifierErr, "Error serializing 'propertyIdentifier' field")
	}

	// Optional Field (arrayIndex) (Can be skipped, if the value is null)
	var arrayIndex BACnetContextTagUnsignedInteger = nil
	if m.GetArrayIndex() != nil {
		if pushErr := writeBuffer.PushContext("arrayIndex"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for arrayIndex")
		}
		arrayIndex = m.GetArrayIndex()
		_arrayIndexErr := writeBuffer.WriteSerializable(arrayIndex)
		if popErr := writeBuffer.PopContext("arrayIndex"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for arrayIndex")
		}
		if _arrayIndexErr != nil {
			return errors.Wrap(_arrayIndexErr, "Error serializing 'arrayIndex' field")
		}
	}

	// Simple Field (propertyValue)
	if pushErr := writeBuffer.PushContext("propertyValue"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for propertyValue")
	}
	_propertyValueErr := writeBuffer.WriteSerializable(m.GetPropertyValue())
	if popErr := writeBuffer.PopContext("propertyValue"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for propertyValue")
	}
	if _propertyValueErr != nil {
		return errors.Wrap(_propertyValueErr, "Error serializing 'propertyValue' field")
	}

	// Optional Field (timeOfChange) (Can be skipped, if the value is null)
	var timeOfChange BACnetContextTagTime = nil
	if m.GetTimeOfChange() != nil {
		if pushErr := writeBuffer.PushContext("timeOfChange"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for timeOfChange")
		}
		timeOfChange = m.GetTimeOfChange()
		_timeOfChangeErr := writeBuffer.WriteSerializable(timeOfChange)
		if popErr := writeBuffer.PopContext("timeOfChange"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for timeOfChange")
		}
		if _timeOfChangeErr != nil {
			return errors.Wrap(_timeOfChangeErr, "Error serializing 'timeOfChange' field")
		}
	}

	if popErr := writeBuffer.PopContext("ListOfCovNotificationsValue"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for ListOfCovNotificationsValue")
	}
	return nil
}

func (m *_ListOfCovNotificationsValue) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
