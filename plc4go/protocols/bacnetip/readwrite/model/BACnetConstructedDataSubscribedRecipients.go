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

// BACnetConstructedDataSubscribedRecipients is the corresponding interface of BACnetConstructedDataSubscribedRecipients
type BACnetConstructedDataSubscribedRecipients interface {
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetSubscribedRecipients returns SubscribedRecipients (property field)
	GetSubscribedRecipients() []BACnetEventNotificationSubscription
}

// _BACnetConstructedDataSubscribedRecipients is the data-structure of this message
type _BACnetConstructedDataSubscribedRecipients struct {
	*_BACnetConstructedData
	SubscribedRecipients []BACnetEventNotificationSubscription

	// Arguments.
	TagNumber          uint8
	ArrayIndexArgument BACnetTagPayloadUnsignedInteger
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataSubscribedRecipients) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataSubscribedRecipients) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_SUBSCRIBED_RECIPIENTS
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataSubscribedRecipients) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataSubscribedRecipients) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataSubscribedRecipients) GetSubscribedRecipients() []BACnetEventNotificationSubscription {
	return m.SubscribedRecipients
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataSubscribedRecipients factory function for _BACnetConstructedDataSubscribedRecipients
func NewBACnetConstructedDataSubscribedRecipients(subscribedRecipients []BACnetEventNotificationSubscription, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataSubscribedRecipients {
	_result := &_BACnetConstructedDataSubscribedRecipients{
		SubscribedRecipients:   subscribedRecipients,
		_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataSubscribedRecipients(structType interface{}) BACnetConstructedDataSubscribedRecipients {
	if casted, ok := structType.(BACnetConstructedDataSubscribedRecipients); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataSubscribedRecipients); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataSubscribedRecipients) GetTypeName() string {
	return "BACnetConstructedDataSubscribedRecipients"
}

func (m *_BACnetConstructedDataSubscribedRecipients) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetConstructedDataSubscribedRecipients) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Array field
	if len(m.SubscribedRecipients) > 0 {
		for _, element := range m.SubscribedRecipients {
			lengthInBits += element.GetLengthInBits()
		}
	}

	return lengthInBits
}

func (m *_BACnetConstructedDataSubscribedRecipients) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataSubscribedRecipientsParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataSubscribedRecipients, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataSubscribedRecipients"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataSubscribedRecipients")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Array field (subscribedRecipients)
	if pullErr := readBuffer.PullContext("subscribedRecipients", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for subscribedRecipients")
	}
	// Terminated array
	subscribedRecipients := make([]BACnetEventNotificationSubscription, 0)
	{
		for !bool(IsBACnetConstructedDataClosingTag(readBuffer, false, tagNumber)) {
			_item, _err := BACnetEventNotificationSubscriptionParse(readBuffer)
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'subscribedRecipients' field")
			}
			subscribedRecipients = append(subscribedRecipients, _item.(BACnetEventNotificationSubscription))

		}
	}
	if closeErr := readBuffer.CloseContext("subscribedRecipients", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for subscribedRecipients")
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataSubscribedRecipients"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataSubscribedRecipients")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataSubscribedRecipients{
		SubscribedRecipients:   subscribedRecipients,
		_BACnetConstructedData: &_BACnetConstructedData{},
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataSubscribedRecipients) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataSubscribedRecipients"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataSubscribedRecipients")
		}

		// Array Field (subscribedRecipients)
		if m.GetSubscribedRecipients() != nil {
			if pushErr := writeBuffer.PushContext("subscribedRecipients", utils.WithRenderAsList(true)); pushErr != nil {
				return errors.Wrap(pushErr, "Error pushing for subscribedRecipients")
			}
			for _, _element := range m.GetSubscribedRecipients() {
				_elementErr := writeBuffer.WriteSerializable(_element)
				if _elementErr != nil {
					return errors.Wrap(_elementErr, "Error serializing 'subscribedRecipients' field")
				}
			}
			if popErr := writeBuffer.PopContext("subscribedRecipients", utils.WithRenderAsList(true)); popErr != nil {
				return errors.Wrap(popErr, "Error popping for subscribedRecipients")
			}
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataSubscribedRecipients"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataSubscribedRecipients")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataSubscribedRecipients) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
