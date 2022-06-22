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

// BACnetEventParameterAccessEvent is the corresponding interface of BACnetEventParameterAccessEvent
type BACnetEventParameterAccessEvent interface {
	utils.LengthAware
	utils.Serializable
	BACnetEventParameter
	// GetOpeningTag returns OpeningTag (property field)
	GetOpeningTag() BACnetOpeningTag
	// GetListOfAccessEvents returns ListOfAccessEvents (property field)
	GetListOfAccessEvents() BACnetEventParameterAccessEventListOfAccessEvents
	// GetAccessEventTimeReference returns AccessEventTimeReference (property field)
	GetAccessEventTimeReference() BACnetDeviceObjectPropertyReferenceEnclosed
	// GetClosingTag returns ClosingTag (property field)
	GetClosingTag() BACnetClosingTag
}

// _BACnetEventParameterAccessEvent is the data-structure of this message
type _BACnetEventParameterAccessEvent struct {
	*_BACnetEventParameter
	OpeningTag               BACnetOpeningTag
	ListOfAccessEvents       BACnetEventParameterAccessEventListOfAccessEvents
	AccessEventTimeReference BACnetDeviceObjectPropertyReferenceEnclosed
	ClosingTag               BACnetClosingTag
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetEventParameterAccessEvent) InitializeParent(parent BACnetEventParameter, peekedTagHeader BACnetTagHeader) {
	m.PeekedTagHeader = peekedTagHeader
}

func (m *_BACnetEventParameterAccessEvent) GetParent() BACnetEventParameter {
	return m._BACnetEventParameter
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetEventParameterAccessEvent) GetOpeningTag() BACnetOpeningTag {
	return m.OpeningTag
}

func (m *_BACnetEventParameterAccessEvent) GetListOfAccessEvents() BACnetEventParameterAccessEventListOfAccessEvents {
	return m.ListOfAccessEvents
}

func (m *_BACnetEventParameterAccessEvent) GetAccessEventTimeReference() BACnetDeviceObjectPropertyReferenceEnclosed {
	return m.AccessEventTimeReference
}

func (m *_BACnetEventParameterAccessEvent) GetClosingTag() BACnetClosingTag {
	return m.ClosingTag
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetEventParameterAccessEvent factory function for _BACnetEventParameterAccessEvent
func NewBACnetEventParameterAccessEvent(openingTag BACnetOpeningTag, listOfAccessEvents BACnetEventParameterAccessEventListOfAccessEvents, accessEventTimeReference BACnetDeviceObjectPropertyReferenceEnclosed, closingTag BACnetClosingTag, peekedTagHeader BACnetTagHeader) *_BACnetEventParameterAccessEvent {
	_result := &_BACnetEventParameterAccessEvent{
		OpeningTag:               openingTag,
		ListOfAccessEvents:       listOfAccessEvents,
		AccessEventTimeReference: accessEventTimeReference,
		ClosingTag:               closingTag,
		_BACnetEventParameter:    NewBACnetEventParameter(peekedTagHeader),
	}
	_result._BACnetEventParameter._BACnetEventParameterChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetEventParameterAccessEvent(structType interface{}) BACnetEventParameterAccessEvent {
	if casted, ok := structType.(BACnetEventParameterAccessEvent); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetEventParameterAccessEvent); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetEventParameterAccessEvent) GetTypeName() string {
	return "BACnetEventParameterAccessEvent"
}

func (m *_BACnetEventParameterAccessEvent) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetEventParameterAccessEvent) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (openingTag)
	lengthInBits += m.OpeningTag.GetLengthInBits()

	// Simple field (listOfAccessEvents)
	lengthInBits += m.ListOfAccessEvents.GetLengthInBits()

	// Simple field (accessEventTimeReference)
	lengthInBits += m.AccessEventTimeReference.GetLengthInBits()

	// Simple field (closingTag)
	lengthInBits += m.ClosingTag.GetLengthInBits()

	return lengthInBits
}

func (m *_BACnetEventParameterAccessEvent) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetEventParameterAccessEventParse(readBuffer utils.ReadBuffer) (BACnetEventParameterAccessEvent, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetEventParameterAccessEvent"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetEventParameterAccessEvent")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (openingTag)
	if pullErr := readBuffer.PullContext("openingTag"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for openingTag")
	}
	_openingTag, _openingTagErr := BACnetOpeningTagParse(readBuffer, uint8(uint8(13)))
	if _openingTagErr != nil {
		return nil, errors.Wrap(_openingTagErr, "Error parsing 'openingTag' field")
	}
	openingTag := _openingTag.(BACnetOpeningTag)
	if closeErr := readBuffer.CloseContext("openingTag"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for openingTag")
	}

	// Simple Field (listOfAccessEvents)
	if pullErr := readBuffer.PullContext("listOfAccessEvents"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for listOfAccessEvents")
	}
	_listOfAccessEvents, _listOfAccessEventsErr := BACnetEventParameterAccessEventListOfAccessEventsParse(readBuffer, uint8(uint8(0)))
	if _listOfAccessEventsErr != nil {
		return nil, errors.Wrap(_listOfAccessEventsErr, "Error parsing 'listOfAccessEvents' field")
	}
	listOfAccessEvents := _listOfAccessEvents.(BACnetEventParameterAccessEventListOfAccessEvents)
	if closeErr := readBuffer.CloseContext("listOfAccessEvents"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for listOfAccessEvents")
	}

	// Simple Field (accessEventTimeReference)
	if pullErr := readBuffer.PullContext("accessEventTimeReference"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for accessEventTimeReference")
	}
	_accessEventTimeReference, _accessEventTimeReferenceErr := BACnetDeviceObjectPropertyReferenceEnclosedParse(readBuffer, uint8(uint8(1)))
	if _accessEventTimeReferenceErr != nil {
		return nil, errors.Wrap(_accessEventTimeReferenceErr, "Error parsing 'accessEventTimeReference' field")
	}
	accessEventTimeReference := _accessEventTimeReference.(BACnetDeviceObjectPropertyReferenceEnclosed)
	if closeErr := readBuffer.CloseContext("accessEventTimeReference"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for accessEventTimeReference")
	}

	// Simple Field (closingTag)
	if pullErr := readBuffer.PullContext("closingTag"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for closingTag")
	}
	_closingTag, _closingTagErr := BACnetClosingTagParse(readBuffer, uint8(uint8(13)))
	if _closingTagErr != nil {
		return nil, errors.Wrap(_closingTagErr, "Error parsing 'closingTag' field")
	}
	closingTag := _closingTag.(BACnetClosingTag)
	if closeErr := readBuffer.CloseContext("closingTag"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for closingTag")
	}

	if closeErr := readBuffer.CloseContext("BACnetEventParameterAccessEvent"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetEventParameterAccessEvent")
	}

	// Create a partially initialized instance
	_child := &_BACnetEventParameterAccessEvent{
		OpeningTag:               openingTag,
		ListOfAccessEvents:       listOfAccessEvents,
		AccessEventTimeReference: accessEventTimeReference,
		ClosingTag:               closingTag,
		_BACnetEventParameter:    &_BACnetEventParameter{},
	}
	_child._BACnetEventParameter._BACnetEventParameterChildRequirements = _child
	return _child, nil
}

func (m *_BACnetEventParameterAccessEvent) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetEventParameterAccessEvent"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetEventParameterAccessEvent")
		}

		// Simple Field (openingTag)
		if pushErr := writeBuffer.PushContext("openingTag"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for openingTag")
		}
		_openingTagErr := writeBuffer.WriteSerializable(m.GetOpeningTag())
		if popErr := writeBuffer.PopContext("openingTag"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for openingTag")
		}
		if _openingTagErr != nil {
			return errors.Wrap(_openingTagErr, "Error serializing 'openingTag' field")
		}

		// Simple Field (listOfAccessEvents)
		if pushErr := writeBuffer.PushContext("listOfAccessEvents"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for listOfAccessEvents")
		}
		_listOfAccessEventsErr := writeBuffer.WriteSerializable(m.GetListOfAccessEvents())
		if popErr := writeBuffer.PopContext("listOfAccessEvents"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for listOfAccessEvents")
		}
		if _listOfAccessEventsErr != nil {
			return errors.Wrap(_listOfAccessEventsErr, "Error serializing 'listOfAccessEvents' field")
		}

		// Simple Field (accessEventTimeReference)
		if pushErr := writeBuffer.PushContext("accessEventTimeReference"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for accessEventTimeReference")
		}
		_accessEventTimeReferenceErr := writeBuffer.WriteSerializable(m.GetAccessEventTimeReference())
		if popErr := writeBuffer.PopContext("accessEventTimeReference"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for accessEventTimeReference")
		}
		if _accessEventTimeReferenceErr != nil {
			return errors.Wrap(_accessEventTimeReferenceErr, "Error serializing 'accessEventTimeReference' field")
		}

		// Simple Field (closingTag)
		if pushErr := writeBuffer.PushContext("closingTag"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for closingTag")
		}
		_closingTagErr := writeBuffer.WriteSerializable(m.GetClosingTag())
		if popErr := writeBuffer.PopContext("closingTag"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for closingTag")
		}
		if _closingTagErr != nil {
			return errors.Wrap(_closingTagErr, "Error serializing 'closingTag' field")
		}

		if popErr := writeBuffer.PopContext("BACnetEventParameterAccessEvent"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetEventParameterAccessEvent")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetEventParameterAccessEvent) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
