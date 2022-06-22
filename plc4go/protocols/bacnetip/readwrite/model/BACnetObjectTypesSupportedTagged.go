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

// BACnetObjectTypesSupportedTagged is the corresponding interface of BACnetObjectTypesSupportedTagged
type BACnetObjectTypesSupportedTagged interface {
	utils.LengthAware
	utils.Serializable
	// GetHeader returns Header (property field)
	GetHeader() BACnetTagHeader
	// GetPayload returns Payload (property field)
	GetPayload() BACnetTagPayloadBitString
	// GetTimeValue returns TimeValue (virtual field)
	GetTimeValue() bool
	// GetNotificationForwarder returns NotificationForwarder (virtual field)
	GetNotificationForwarder() bool
	// GetAlertEnrollment returns AlertEnrollment (virtual field)
	GetAlertEnrollment() bool
	// GetChannel returns Channel (virtual field)
	GetChannel() bool
	// GetLightingOutput returns LightingOutput (virtual field)
	GetLightingOutput() bool
	// GetBinaryLightingOutput returns BinaryLightingOutput (virtual field)
	GetBinaryLightingOutput() bool
	// GetNetworkPort returns NetworkPort (virtual field)
	GetNetworkPort() bool
	// GetElevatorGroup returns ElevatorGroup (virtual field)
	GetElevatorGroup() bool
	// GetEscalator returns Escalator (virtual field)
	GetEscalator() bool
	// GetLift returns Lift (virtual field)
	GetLift() bool
}

// _BACnetObjectTypesSupportedTagged is the data-structure of this message
type _BACnetObjectTypesSupportedTagged struct {
	Header  BACnetTagHeader
	Payload BACnetTagPayloadBitString

	// Arguments.
	TagNumber uint8
	TagClass  TagClass
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetObjectTypesSupportedTagged) GetHeader() BACnetTagHeader {
	return m.Header
}

func (m *_BACnetObjectTypesSupportedTagged) GetPayload() BACnetTagPayloadBitString {
	return m.Payload
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetObjectTypesSupportedTagged) GetTimeValue() bool {
	return bool(utils.InlineIf(bool(bool((len(m.GetPayload().GetData())) > (0))), func() interface{} { return bool(m.GetPayload().GetData()[0]) }, func() interface{} { return bool(bool(false)) }).(bool))
}

func (m *_BACnetObjectTypesSupportedTagged) GetNotificationForwarder() bool {
	return bool(utils.InlineIf(bool(bool((len(m.GetPayload().GetData())) > (1))), func() interface{} { return bool(m.GetPayload().GetData()[1]) }, func() interface{} { return bool(bool(false)) }).(bool))
}

func (m *_BACnetObjectTypesSupportedTagged) GetAlertEnrollment() bool {
	return bool(utils.InlineIf(bool(bool((len(m.GetPayload().GetData())) > (2))), func() interface{} { return bool(m.GetPayload().GetData()[2]) }, func() interface{} { return bool(bool(false)) }).(bool))
}

func (m *_BACnetObjectTypesSupportedTagged) GetChannel() bool {
	return bool(utils.InlineIf(bool(bool((len(m.GetPayload().GetData())) > (3))), func() interface{} { return bool(m.GetPayload().GetData()[3]) }, func() interface{} { return bool(bool(false)) }).(bool))
}

func (m *_BACnetObjectTypesSupportedTagged) GetLightingOutput() bool {
	return bool(utils.InlineIf(bool(bool((len(m.GetPayload().GetData())) > (4))), func() interface{} { return bool(m.GetPayload().GetData()[4]) }, func() interface{} { return bool(bool(false)) }).(bool))
}

func (m *_BACnetObjectTypesSupportedTagged) GetBinaryLightingOutput() bool {
	return bool(utils.InlineIf(bool(bool((len(m.GetPayload().GetData())) > (5))), func() interface{} { return bool(m.GetPayload().GetData()[5]) }, func() interface{} { return bool(bool(false)) }).(bool))
}

func (m *_BACnetObjectTypesSupportedTagged) GetNetworkPort() bool {
	return bool(utils.InlineIf(bool(bool((len(m.GetPayload().GetData())) > (6))), func() interface{} { return bool(m.GetPayload().GetData()[6]) }, func() interface{} { return bool(bool(false)) }).(bool))
}

func (m *_BACnetObjectTypesSupportedTagged) GetElevatorGroup() bool {
	return bool(utils.InlineIf(bool(bool((len(m.GetPayload().GetData())) > (7))), func() interface{} { return bool(m.GetPayload().GetData()[7]) }, func() interface{} { return bool(bool(false)) }).(bool))
}

func (m *_BACnetObjectTypesSupportedTagged) GetEscalator() bool {
	return bool(utils.InlineIf(bool(bool((len(m.GetPayload().GetData())) > (8))), func() interface{} { return bool(m.GetPayload().GetData()[8]) }, func() interface{} { return bool(bool(false)) }).(bool))
}

func (m *_BACnetObjectTypesSupportedTagged) GetLift() bool {
	return bool(utils.InlineIf(bool(bool((len(m.GetPayload().GetData())) > (9))), func() interface{} { return bool(m.GetPayload().GetData()[9]) }, func() interface{} { return bool(bool(false)) }).(bool))
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetObjectTypesSupportedTagged factory function for _BACnetObjectTypesSupportedTagged
func NewBACnetObjectTypesSupportedTagged(header BACnetTagHeader, payload BACnetTagPayloadBitString, tagNumber uint8, tagClass TagClass) *_BACnetObjectTypesSupportedTagged {
	return &_BACnetObjectTypesSupportedTagged{Header: header, Payload: payload, TagNumber: tagNumber, TagClass: tagClass}
}

// Deprecated: use the interface for direct cast
func CastBACnetObjectTypesSupportedTagged(structType interface{}) BACnetObjectTypesSupportedTagged {
	if casted, ok := structType.(BACnetObjectTypesSupportedTagged); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetObjectTypesSupportedTagged); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetObjectTypesSupportedTagged) GetTypeName() string {
	return "BACnetObjectTypesSupportedTagged"
}

func (m *_BACnetObjectTypesSupportedTagged) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetObjectTypesSupportedTagged) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(0)

	// Simple field (header)
	lengthInBits += m.Header.GetLengthInBits()

	// Simple field (payload)
	lengthInBits += m.Payload.GetLengthInBits()

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetObjectTypesSupportedTagged) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetObjectTypesSupportedTaggedParse(readBuffer utils.ReadBuffer, tagNumber uint8, tagClass TagClass) (BACnetObjectTypesSupportedTagged, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetObjectTypesSupportedTagged"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetObjectTypesSupportedTagged")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (header)
	if pullErr := readBuffer.PullContext("header"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for header")
	}
	_header, _headerErr := BACnetTagHeaderParse(readBuffer)
	if _headerErr != nil {
		return nil, errors.Wrap(_headerErr, "Error parsing 'header' field")
	}
	header := _header.(BACnetTagHeader)
	if closeErr := readBuffer.CloseContext("header"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for header")
	}

	// Validation
	if !(bool((header.GetTagClass()) == (tagClass))) {
		return nil, errors.WithStack(utils.ParseValidationError{"tag class doesn't match"})
	}

	// Validation
	if !(bool(bool(bool((header.GetTagClass()) == (TagClass_APPLICATION_TAGS)))) || bool(bool(bool((header.GetActualTagNumber()) == (tagNumber))))) {
		return nil, errors.WithStack(utils.ParseAssertError{"tagnumber doesn't match"})
	}

	// Simple Field (payload)
	if pullErr := readBuffer.PullContext("payload"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for payload")
	}
	_payload, _payloadErr := BACnetTagPayloadBitStringParse(readBuffer, uint32(header.GetActualLength()))
	if _payloadErr != nil {
		return nil, errors.Wrap(_payloadErr, "Error parsing 'payload' field")
	}
	payload := _payload.(BACnetTagPayloadBitString)
	if closeErr := readBuffer.CloseContext("payload"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for payload")
	}

	// Virtual field
	_timeValue := utils.InlineIf(bool(bool((len(payload.GetData())) > (0))), func() interface{} { return bool(payload.GetData()[0]) }, func() interface{} { return bool(bool(false)) }).(bool)
	timeValue := bool(_timeValue)
	_ = timeValue

	// Virtual field
	_notificationForwarder := utils.InlineIf(bool(bool((len(payload.GetData())) > (1))), func() interface{} { return bool(payload.GetData()[1]) }, func() interface{} { return bool(bool(false)) }).(bool)
	notificationForwarder := bool(_notificationForwarder)
	_ = notificationForwarder

	// Virtual field
	_alertEnrollment := utils.InlineIf(bool(bool((len(payload.GetData())) > (2))), func() interface{} { return bool(payload.GetData()[2]) }, func() interface{} { return bool(bool(false)) }).(bool)
	alertEnrollment := bool(_alertEnrollment)
	_ = alertEnrollment

	// Virtual field
	_channel := utils.InlineIf(bool(bool((len(payload.GetData())) > (3))), func() interface{} { return bool(payload.GetData()[3]) }, func() interface{} { return bool(bool(false)) }).(bool)
	channel := bool(_channel)
	_ = channel

	// Virtual field
	_lightingOutput := utils.InlineIf(bool(bool((len(payload.GetData())) > (4))), func() interface{} { return bool(payload.GetData()[4]) }, func() interface{} { return bool(bool(false)) }).(bool)
	lightingOutput := bool(_lightingOutput)
	_ = lightingOutput

	// Virtual field
	_binaryLightingOutput := utils.InlineIf(bool(bool((len(payload.GetData())) > (5))), func() interface{} { return bool(payload.GetData()[5]) }, func() interface{} { return bool(bool(false)) }).(bool)
	binaryLightingOutput := bool(_binaryLightingOutput)
	_ = binaryLightingOutput

	// Virtual field
	_networkPort := utils.InlineIf(bool(bool((len(payload.GetData())) > (6))), func() interface{} { return bool(payload.GetData()[6]) }, func() interface{} { return bool(bool(false)) }).(bool)
	networkPort := bool(_networkPort)
	_ = networkPort

	// Virtual field
	_elevatorGroup := utils.InlineIf(bool(bool((len(payload.GetData())) > (7))), func() interface{} { return bool(payload.GetData()[7]) }, func() interface{} { return bool(bool(false)) }).(bool)
	elevatorGroup := bool(_elevatorGroup)
	_ = elevatorGroup

	// Virtual field
	_escalator := utils.InlineIf(bool(bool((len(payload.GetData())) > (8))), func() interface{} { return bool(payload.GetData()[8]) }, func() interface{} { return bool(bool(false)) }).(bool)
	escalator := bool(_escalator)
	_ = escalator

	// Virtual field
	_lift := utils.InlineIf(bool(bool((len(payload.GetData())) > (9))), func() interface{} { return bool(payload.GetData()[9]) }, func() interface{} { return bool(bool(false)) }).(bool)
	lift := bool(_lift)
	_ = lift

	if closeErr := readBuffer.CloseContext("BACnetObjectTypesSupportedTagged"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetObjectTypesSupportedTagged")
	}

	// Create the instance
	return NewBACnetObjectTypesSupportedTagged(header, payload, tagNumber, tagClass), nil
}

func (m *_BACnetObjectTypesSupportedTagged) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("BACnetObjectTypesSupportedTagged"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetObjectTypesSupportedTagged")
	}

	// Simple Field (header)
	if pushErr := writeBuffer.PushContext("header"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for header")
	}
	_headerErr := writeBuffer.WriteSerializable(m.GetHeader())
	if popErr := writeBuffer.PopContext("header"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for header")
	}
	if _headerErr != nil {
		return errors.Wrap(_headerErr, "Error serializing 'header' field")
	}

	// Simple Field (payload)
	if pushErr := writeBuffer.PushContext("payload"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for payload")
	}
	_payloadErr := writeBuffer.WriteSerializable(m.GetPayload())
	if popErr := writeBuffer.PopContext("payload"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for payload")
	}
	if _payloadErr != nil {
		return errors.Wrap(_payloadErr, "Error serializing 'payload' field")
	}
	// Virtual field
	if _timeValueErr := writeBuffer.WriteVirtual("timeValue", m.GetTimeValue()); _timeValueErr != nil {
		return errors.Wrap(_timeValueErr, "Error serializing 'timeValue' field")
	}
	// Virtual field
	if _notificationForwarderErr := writeBuffer.WriteVirtual("notificationForwarder", m.GetNotificationForwarder()); _notificationForwarderErr != nil {
		return errors.Wrap(_notificationForwarderErr, "Error serializing 'notificationForwarder' field")
	}
	// Virtual field
	if _alertEnrollmentErr := writeBuffer.WriteVirtual("alertEnrollment", m.GetAlertEnrollment()); _alertEnrollmentErr != nil {
		return errors.Wrap(_alertEnrollmentErr, "Error serializing 'alertEnrollment' field")
	}
	// Virtual field
	if _channelErr := writeBuffer.WriteVirtual("channel", m.GetChannel()); _channelErr != nil {
		return errors.Wrap(_channelErr, "Error serializing 'channel' field")
	}
	// Virtual field
	if _lightingOutputErr := writeBuffer.WriteVirtual("lightingOutput", m.GetLightingOutput()); _lightingOutputErr != nil {
		return errors.Wrap(_lightingOutputErr, "Error serializing 'lightingOutput' field")
	}
	// Virtual field
	if _binaryLightingOutputErr := writeBuffer.WriteVirtual("binaryLightingOutput", m.GetBinaryLightingOutput()); _binaryLightingOutputErr != nil {
		return errors.Wrap(_binaryLightingOutputErr, "Error serializing 'binaryLightingOutput' field")
	}
	// Virtual field
	if _networkPortErr := writeBuffer.WriteVirtual("networkPort", m.GetNetworkPort()); _networkPortErr != nil {
		return errors.Wrap(_networkPortErr, "Error serializing 'networkPort' field")
	}
	// Virtual field
	if _elevatorGroupErr := writeBuffer.WriteVirtual("elevatorGroup", m.GetElevatorGroup()); _elevatorGroupErr != nil {
		return errors.Wrap(_elevatorGroupErr, "Error serializing 'elevatorGroup' field")
	}
	// Virtual field
	if _escalatorErr := writeBuffer.WriteVirtual("escalator", m.GetEscalator()); _escalatorErr != nil {
		return errors.Wrap(_escalatorErr, "Error serializing 'escalator' field")
	}
	// Virtual field
	if _liftErr := writeBuffer.WriteVirtual("lift", m.GetLift()); _liftErr != nil {
		return errors.Wrap(_liftErr, "Error serializing 'lift' field")
	}

	if popErr := writeBuffer.PopContext("BACnetObjectTypesSupportedTagged"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetObjectTypesSupportedTagged")
	}
	return nil
}

func (m *_BACnetObjectTypesSupportedTagged) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
