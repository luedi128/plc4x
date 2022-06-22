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

// BACnetNotificationParametersChangeOfLifeSafety is the corresponding interface of BACnetNotificationParametersChangeOfLifeSafety
type BACnetNotificationParametersChangeOfLifeSafety interface {
	utils.LengthAware
	utils.Serializable
	BACnetNotificationParameters
	// GetInnerOpeningTag returns InnerOpeningTag (property field)
	GetInnerOpeningTag() BACnetOpeningTag
	// GetNewState returns NewState (property field)
	GetNewState() BACnetLifeSafetyStateTagged
	// GetNewMode returns NewMode (property field)
	GetNewMode() BACnetLifeSafetyModeTagged
	// GetStatusFlags returns StatusFlags (property field)
	GetStatusFlags() BACnetStatusFlagsTagged
	// GetOperationExpected returns OperationExpected (property field)
	GetOperationExpected() BACnetLifeSafetyOperationTagged
	// GetInnerClosingTag returns InnerClosingTag (property field)
	GetInnerClosingTag() BACnetClosingTag
}

// _BACnetNotificationParametersChangeOfLifeSafety is the data-structure of this message
type _BACnetNotificationParametersChangeOfLifeSafety struct {
	*_BACnetNotificationParameters
	InnerOpeningTag   BACnetOpeningTag
	NewState          BACnetLifeSafetyStateTagged
	NewMode           BACnetLifeSafetyModeTagged
	StatusFlags       BACnetStatusFlagsTagged
	OperationExpected BACnetLifeSafetyOperationTagged
	InnerClosingTag   BACnetClosingTag

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

func (m *_BACnetNotificationParametersChangeOfLifeSafety) InitializeParent(parent BACnetNotificationParameters, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetNotificationParametersChangeOfLifeSafety) GetParent() BACnetNotificationParameters {
	return m._BACnetNotificationParameters
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetNotificationParametersChangeOfLifeSafety) GetInnerOpeningTag() BACnetOpeningTag {
	return m.InnerOpeningTag
}

func (m *_BACnetNotificationParametersChangeOfLifeSafety) GetNewState() BACnetLifeSafetyStateTagged {
	return m.NewState
}

func (m *_BACnetNotificationParametersChangeOfLifeSafety) GetNewMode() BACnetLifeSafetyModeTagged {
	return m.NewMode
}

func (m *_BACnetNotificationParametersChangeOfLifeSafety) GetStatusFlags() BACnetStatusFlagsTagged {
	return m.StatusFlags
}

func (m *_BACnetNotificationParametersChangeOfLifeSafety) GetOperationExpected() BACnetLifeSafetyOperationTagged {
	return m.OperationExpected
}

func (m *_BACnetNotificationParametersChangeOfLifeSafety) GetInnerClosingTag() BACnetClosingTag {
	return m.InnerClosingTag
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetNotificationParametersChangeOfLifeSafety factory function for _BACnetNotificationParametersChangeOfLifeSafety
func NewBACnetNotificationParametersChangeOfLifeSafety(innerOpeningTag BACnetOpeningTag, newState BACnetLifeSafetyStateTagged, newMode BACnetLifeSafetyModeTagged, statusFlags BACnetStatusFlagsTagged, operationExpected BACnetLifeSafetyOperationTagged, innerClosingTag BACnetClosingTag, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, objectTypeArgument BACnetObjectType) *_BACnetNotificationParametersChangeOfLifeSafety {
	_result := &_BACnetNotificationParametersChangeOfLifeSafety{
		InnerOpeningTag:               innerOpeningTag,
		NewState:                      newState,
		NewMode:                       newMode,
		StatusFlags:                   statusFlags,
		OperationExpected:             operationExpected,
		InnerClosingTag:               innerClosingTag,
		_BACnetNotificationParameters: NewBACnetNotificationParameters(openingTag, peekedTagHeader, closingTag, tagNumber, objectTypeArgument),
	}
	_result._BACnetNotificationParameters._BACnetNotificationParametersChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetNotificationParametersChangeOfLifeSafety(structType interface{}) BACnetNotificationParametersChangeOfLifeSafety {
	if casted, ok := structType.(BACnetNotificationParametersChangeOfLifeSafety); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetNotificationParametersChangeOfLifeSafety); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetNotificationParametersChangeOfLifeSafety) GetTypeName() string {
	return "BACnetNotificationParametersChangeOfLifeSafety"
}

func (m *_BACnetNotificationParametersChangeOfLifeSafety) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetNotificationParametersChangeOfLifeSafety) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (innerOpeningTag)
	lengthInBits += m.InnerOpeningTag.GetLengthInBits()

	// Simple field (newState)
	lengthInBits += m.NewState.GetLengthInBits()

	// Simple field (newMode)
	lengthInBits += m.NewMode.GetLengthInBits()

	// Simple field (statusFlags)
	lengthInBits += m.StatusFlags.GetLengthInBits()

	// Simple field (operationExpected)
	lengthInBits += m.OperationExpected.GetLengthInBits()

	// Simple field (innerClosingTag)
	lengthInBits += m.InnerClosingTag.GetLengthInBits()

	return lengthInBits
}

func (m *_BACnetNotificationParametersChangeOfLifeSafety) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetNotificationParametersChangeOfLifeSafetyParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, peekedTagNumber uint8) (BACnetNotificationParametersChangeOfLifeSafety, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetNotificationParametersChangeOfLifeSafety"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetNotificationParametersChangeOfLifeSafety")
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

	// Simple Field (newState)
	if pullErr := readBuffer.PullContext("newState"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for newState")
	}
	_newState, _newStateErr := BACnetLifeSafetyStateTaggedParse(readBuffer, uint8(uint8(0)), TagClass(TagClass_CONTEXT_SPECIFIC_TAGS))
	if _newStateErr != nil {
		return nil, errors.Wrap(_newStateErr, "Error parsing 'newState' field")
	}
	newState := _newState.(BACnetLifeSafetyStateTagged)
	if closeErr := readBuffer.CloseContext("newState"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for newState")
	}

	// Simple Field (newMode)
	if pullErr := readBuffer.PullContext("newMode"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for newMode")
	}
	_newMode, _newModeErr := BACnetLifeSafetyModeTaggedParse(readBuffer, uint8(uint8(1)), TagClass(TagClass_CONTEXT_SPECIFIC_TAGS))
	if _newModeErr != nil {
		return nil, errors.Wrap(_newModeErr, "Error parsing 'newMode' field")
	}
	newMode := _newMode.(BACnetLifeSafetyModeTagged)
	if closeErr := readBuffer.CloseContext("newMode"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for newMode")
	}

	// Simple Field (statusFlags)
	if pullErr := readBuffer.PullContext("statusFlags"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for statusFlags")
	}
	_statusFlags, _statusFlagsErr := BACnetStatusFlagsTaggedParse(readBuffer, uint8(uint8(2)), TagClass(TagClass_CONTEXT_SPECIFIC_TAGS))
	if _statusFlagsErr != nil {
		return nil, errors.Wrap(_statusFlagsErr, "Error parsing 'statusFlags' field")
	}
	statusFlags := _statusFlags.(BACnetStatusFlagsTagged)
	if closeErr := readBuffer.CloseContext("statusFlags"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for statusFlags")
	}

	// Simple Field (operationExpected)
	if pullErr := readBuffer.PullContext("operationExpected"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for operationExpected")
	}
	_operationExpected, _operationExpectedErr := BACnetLifeSafetyOperationTaggedParse(readBuffer, uint8(uint8(3)), TagClass(TagClass_CONTEXT_SPECIFIC_TAGS))
	if _operationExpectedErr != nil {
		return nil, errors.Wrap(_operationExpectedErr, "Error parsing 'operationExpected' field")
	}
	operationExpected := _operationExpected.(BACnetLifeSafetyOperationTagged)
	if closeErr := readBuffer.CloseContext("operationExpected"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for operationExpected")
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

	if closeErr := readBuffer.CloseContext("BACnetNotificationParametersChangeOfLifeSafety"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetNotificationParametersChangeOfLifeSafety")
	}

	// Create a partially initialized instance
	_child := &_BACnetNotificationParametersChangeOfLifeSafety{
		InnerOpeningTag:               innerOpeningTag,
		NewState:                      newState,
		NewMode:                       newMode,
		StatusFlags:                   statusFlags,
		OperationExpected:             operationExpected,
		InnerClosingTag:               innerClosingTag,
		_BACnetNotificationParameters: &_BACnetNotificationParameters{},
	}
	_child._BACnetNotificationParameters._BACnetNotificationParametersChildRequirements = _child
	return _child, nil
}

func (m *_BACnetNotificationParametersChangeOfLifeSafety) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetNotificationParametersChangeOfLifeSafety"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetNotificationParametersChangeOfLifeSafety")
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

		// Simple Field (newState)
		if pushErr := writeBuffer.PushContext("newState"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for newState")
		}
		_newStateErr := writeBuffer.WriteSerializable(m.GetNewState())
		if popErr := writeBuffer.PopContext("newState"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for newState")
		}
		if _newStateErr != nil {
			return errors.Wrap(_newStateErr, "Error serializing 'newState' field")
		}

		// Simple Field (newMode)
		if pushErr := writeBuffer.PushContext("newMode"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for newMode")
		}
		_newModeErr := writeBuffer.WriteSerializable(m.GetNewMode())
		if popErr := writeBuffer.PopContext("newMode"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for newMode")
		}
		if _newModeErr != nil {
			return errors.Wrap(_newModeErr, "Error serializing 'newMode' field")
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

		// Simple Field (operationExpected)
		if pushErr := writeBuffer.PushContext("operationExpected"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for operationExpected")
		}
		_operationExpectedErr := writeBuffer.WriteSerializable(m.GetOperationExpected())
		if popErr := writeBuffer.PopContext("operationExpected"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for operationExpected")
		}
		if _operationExpectedErr != nil {
			return errors.Wrap(_operationExpectedErr, "Error serializing 'operationExpected' field")
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

		if popErr := writeBuffer.PopContext("BACnetNotificationParametersChangeOfLifeSafety"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetNotificationParametersChangeOfLifeSafety")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetNotificationParametersChangeOfLifeSafety) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
