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

// BACnetNotificationParametersChangeOfCharacterString is the corresponding interface of BACnetNotificationParametersChangeOfCharacterString
type BACnetNotificationParametersChangeOfCharacterString interface {
	utils.LengthAware
	utils.Serializable
	BACnetNotificationParameters
	// GetInnerOpeningTag returns InnerOpeningTag (property field)
	GetInnerOpeningTag() BACnetOpeningTag
	// GetChangedValue returns ChangedValue (property field)
	GetChangedValue() BACnetContextTagCharacterString
	// GetStatusFlags returns StatusFlags (property field)
	GetStatusFlags() BACnetStatusFlagsTagged
	// GetAlarmValue returns AlarmValue (property field)
	GetAlarmValue() BACnetContextTagCharacterString
	// GetInnerClosingTag returns InnerClosingTag (property field)
	GetInnerClosingTag() BACnetClosingTag
}

// _BACnetNotificationParametersChangeOfCharacterString is the data-structure of this message
type _BACnetNotificationParametersChangeOfCharacterString struct {
	*_BACnetNotificationParameters
	InnerOpeningTag BACnetOpeningTag
	ChangedValue    BACnetContextTagCharacterString
	StatusFlags     BACnetStatusFlagsTagged
	AlarmValue      BACnetContextTagCharacterString
	InnerClosingTag BACnetClosingTag

	// Arguments.
	TagNumber          uint8
	ObjectTypeArgument BACnetObjectType
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetNotificationParametersChangeOfCharacterString) InitializeParent(parent BACnetNotificationParameters, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetNotificationParametersChangeOfCharacterString) GetParent() BACnetNotificationParameters {
	return m._BACnetNotificationParameters
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetNotificationParametersChangeOfCharacterString) GetInnerOpeningTag() BACnetOpeningTag {
	return m.InnerOpeningTag
}

func (m *_BACnetNotificationParametersChangeOfCharacterString) GetChangedValue() BACnetContextTagCharacterString {
	return m.ChangedValue
}

func (m *_BACnetNotificationParametersChangeOfCharacterString) GetStatusFlags() BACnetStatusFlagsTagged {
	return m.StatusFlags
}

func (m *_BACnetNotificationParametersChangeOfCharacterString) GetAlarmValue() BACnetContextTagCharacterString {
	return m.AlarmValue
}

func (m *_BACnetNotificationParametersChangeOfCharacterString) GetInnerClosingTag() BACnetClosingTag {
	return m.InnerClosingTag
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetNotificationParametersChangeOfCharacterString factory function for _BACnetNotificationParametersChangeOfCharacterString
func NewBACnetNotificationParametersChangeOfCharacterString(innerOpeningTag BACnetOpeningTag, changedValue BACnetContextTagCharacterString, statusFlags BACnetStatusFlagsTagged, alarmValue BACnetContextTagCharacterString, innerClosingTag BACnetClosingTag, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, objectTypeArgument BACnetObjectType) *_BACnetNotificationParametersChangeOfCharacterString {
	_result := &_BACnetNotificationParametersChangeOfCharacterString{
		InnerOpeningTag:               innerOpeningTag,
		ChangedValue:                  changedValue,
		StatusFlags:                   statusFlags,
		AlarmValue:                    alarmValue,
		InnerClosingTag:               innerClosingTag,
		_BACnetNotificationParameters: NewBACnetNotificationParameters(openingTag, peekedTagHeader, closingTag, tagNumber, objectTypeArgument),
	}
	_result._BACnetNotificationParameters._BACnetNotificationParametersChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetNotificationParametersChangeOfCharacterString(structType interface{}) BACnetNotificationParametersChangeOfCharacterString {
	if casted, ok := structType.(BACnetNotificationParametersChangeOfCharacterString); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetNotificationParametersChangeOfCharacterString); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetNotificationParametersChangeOfCharacterString) GetTypeName() string {
	return "BACnetNotificationParametersChangeOfCharacterString"
}

func (m *_BACnetNotificationParametersChangeOfCharacterString) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetNotificationParametersChangeOfCharacterString) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (innerOpeningTag)
	lengthInBits += m.InnerOpeningTag.GetLengthInBits()

	// Simple field (changedValue)
	lengthInBits += m.ChangedValue.GetLengthInBits()

	// Simple field (statusFlags)
	lengthInBits += m.StatusFlags.GetLengthInBits()

	// Simple field (alarmValue)
	lengthInBits += m.AlarmValue.GetLengthInBits()

	// Simple field (innerClosingTag)
	lengthInBits += m.InnerClosingTag.GetLengthInBits()

	return lengthInBits
}

func (m *_BACnetNotificationParametersChangeOfCharacterString) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetNotificationParametersChangeOfCharacterStringParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, peekedTagNumber uint8) (BACnetNotificationParametersChangeOfCharacterString, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetNotificationParametersChangeOfCharacterString"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetNotificationParametersChangeOfCharacterString")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (innerOpeningTag)
	if pullErr := readBuffer.PullContext("innerOpeningTag"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for innerOpeningTag")
	}
	_innerOpeningTag, _innerOpeningTagErr := BACnetOpeningTagParse(readBuffer, uint8(peekedTagNumber))
	if _innerOpeningTagErr != nil {
		return nil, errors.Wrap(_innerOpeningTagErr, "Error parsing 'innerOpeningTag' field")
	}
	innerOpeningTag := _innerOpeningTag.(BACnetOpeningTag)
	if closeErr := readBuffer.CloseContext("innerOpeningTag"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for innerOpeningTag")
	}

	// Simple Field (changedValue)
	if pullErr := readBuffer.PullContext("changedValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for changedValue")
	}
	_changedValue, _changedValueErr := BACnetContextTagParse(readBuffer, uint8(uint8(0)), BACnetDataType(BACnetDataType_CHARACTER_STRING))
	if _changedValueErr != nil {
		return nil, errors.Wrap(_changedValueErr, "Error parsing 'changedValue' field")
	}
	changedValue := _changedValue.(BACnetContextTagCharacterString)
	if closeErr := readBuffer.CloseContext("changedValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for changedValue")
	}

	// Simple Field (statusFlags)
	if pullErr := readBuffer.PullContext("statusFlags"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for statusFlags")
	}
	_statusFlags, _statusFlagsErr := BACnetStatusFlagsTaggedParse(readBuffer, uint8(uint8(1)), TagClass(TagClass_CONTEXT_SPECIFIC_TAGS))
	if _statusFlagsErr != nil {
		return nil, errors.Wrap(_statusFlagsErr, "Error parsing 'statusFlags' field")
	}
	statusFlags := _statusFlags.(BACnetStatusFlagsTagged)
	if closeErr := readBuffer.CloseContext("statusFlags"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for statusFlags")
	}

	// Simple Field (alarmValue)
	if pullErr := readBuffer.PullContext("alarmValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for alarmValue")
	}
	_alarmValue, _alarmValueErr := BACnetContextTagParse(readBuffer, uint8(uint8(2)), BACnetDataType(BACnetDataType_CHARACTER_STRING))
	if _alarmValueErr != nil {
		return nil, errors.Wrap(_alarmValueErr, "Error parsing 'alarmValue' field")
	}
	alarmValue := _alarmValue.(BACnetContextTagCharacterString)
	if closeErr := readBuffer.CloseContext("alarmValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for alarmValue")
	}

	// Simple Field (innerClosingTag)
	if pullErr := readBuffer.PullContext("innerClosingTag"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for innerClosingTag")
	}
	_innerClosingTag, _innerClosingTagErr := BACnetClosingTagParse(readBuffer, uint8(peekedTagNumber))
	if _innerClosingTagErr != nil {
		return nil, errors.Wrap(_innerClosingTagErr, "Error parsing 'innerClosingTag' field")
	}
	innerClosingTag := _innerClosingTag.(BACnetClosingTag)
	if closeErr := readBuffer.CloseContext("innerClosingTag"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for innerClosingTag")
	}

	if closeErr := readBuffer.CloseContext("BACnetNotificationParametersChangeOfCharacterString"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetNotificationParametersChangeOfCharacterString")
	}

	// Create a partially initialized instance
	_child := &_BACnetNotificationParametersChangeOfCharacterString{
		InnerOpeningTag:               innerOpeningTag,
		ChangedValue:                  changedValue,
		StatusFlags:                   statusFlags,
		AlarmValue:                    alarmValue,
		InnerClosingTag:               innerClosingTag,
		_BACnetNotificationParameters: &_BACnetNotificationParameters{},
	}
	_child._BACnetNotificationParameters._BACnetNotificationParametersChildRequirements = _child
	return _child, nil
}

func (m *_BACnetNotificationParametersChangeOfCharacterString) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetNotificationParametersChangeOfCharacterString"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetNotificationParametersChangeOfCharacterString")
		}

		// Simple Field (innerOpeningTag)
		if pushErr := writeBuffer.PushContext("innerOpeningTag"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for innerOpeningTag")
		}
		_innerOpeningTagErr := writeBuffer.WriteSerializable(m.GetInnerOpeningTag())
		if popErr := writeBuffer.PopContext("innerOpeningTag"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for innerOpeningTag")
		}
		if _innerOpeningTagErr != nil {
			return errors.Wrap(_innerOpeningTagErr, "Error serializing 'innerOpeningTag' field")
		}

		// Simple Field (changedValue)
		if pushErr := writeBuffer.PushContext("changedValue"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for changedValue")
		}
		_changedValueErr := writeBuffer.WriteSerializable(m.GetChangedValue())
		if popErr := writeBuffer.PopContext("changedValue"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for changedValue")
		}
		if _changedValueErr != nil {
			return errors.Wrap(_changedValueErr, "Error serializing 'changedValue' field")
		}

		// Simple Field (statusFlags)
		if pushErr := writeBuffer.PushContext("statusFlags"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for statusFlags")
		}
		_statusFlagsErr := writeBuffer.WriteSerializable(m.GetStatusFlags())
		if popErr := writeBuffer.PopContext("statusFlags"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for statusFlags")
		}
		if _statusFlagsErr != nil {
			return errors.Wrap(_statusFlagsErr, "Error serializing 'statusFlags' field")
		}

		// Simple Field (alarmValue)
		if pushErr := writeBuffer.PushContext("alarmValue"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for alarmValue")
		}
		_alarmValueErr := writeBuffer.WriteSerializable(m.GetAlarmValue())
		if popErr := writeBuffer.PopContext("alarmValue"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for alarmValue")
		}
		if _alarmValueErr != nil {
			return errors.Wrap(_alarmValueErr, "Error serializing 'alarmValue' field")
		}

		// Simple Field (innerClosingTag)
		if pushErr := writeBuffer.PushContext("innerClosingTag"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for innerClosingTag")
		}
		_innerClosingTagErr := writeBuffer.WriteSerializable(m.GetInnerClosingTag())
		if popErr := writeBuffer.PopContext("innerClosingTag"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for innerClosingTag")
		}
		if _innerClosingTagErr != nil {
			return errors.Wrap(_innerClosingTagErr, "Error serializing 'innerClosingTag' field")
		}

		if popErr := writeBuffer.PopContext("BACnetNotificationParametersChangeOfCharacterString"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetNotificationParametersChangeOfCharacterString")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetNotificationParametersChangeOfCharacterString) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
