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

// BACnetContextTagReal is the corresponding interface of BACnetContextTagReal
type BACnetContextTagReal interface {
	utils.LengthAware
	utils.Serializable
	BACnetContextTag
	// GetPayload returns Payload (property field)
	GetPayload() BACnetTagPayloadReal
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() float32
}

// _BACnetContextTagReal is the data-structure of this message
type _BACnetContextTagReal struct {
	*_BACnetContextTag
	Payload BACnetTagPayloadReal

	// Arguments.
	TagNumberArgument uint8
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetContextTagReal) GetDataType() BACnetDataType {
	return BACnetDataType_REAL
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetContextTagReal) InitializeParent(parent BACnetContextTag, header BACnetTagHeader) {
	m.Header = header
}

func (m *_BACnetContextTagReal) GetParent() BACnetContextTag {
	return m._BACnetContextTag
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetContextTagReal) GetPayload() BACnetTagPayloadReal {
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

func (m *_BACnetContextTagReal) GetActualValue() float32 {
	return float32(m.GetPayload().GetValue())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetContextTagReal factory function for _BACnetContextTagReal
func NewBACnetContextTagReal(payload BACnetTagPayloadReal, header BACnetTagHeader, tagNumberArgument uint8) *_BACnetContextTagReal {
	_result := &_BACnetContextTagReal{
		Payload:           payload,
		_BACnetContextTag: NewBACnetContextTag(header, tagNumberArgument),
	}
	_result._BACnetContextTag._BACnetContextTagChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetContextTagReal(structType interface{}) BACnetContextTagReal {
	if casted, ok := structType.(BACnetContextTagReal); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetContextTagReal); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetContextTagReal) GetTypeName() string {
	return "BACnetContextTagReal"
}

func (m *_BACnetContextTagReal) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetContextTagReal) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (payload)
	lengthInBits += m.Payload.GetLengthInBits()

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetContextTagReal) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetContextTagRealParse(readBuffer utils.ReadBuffer, tagNumberArgument uint8, dataType BACnetDataType) (BACnetContextTagReal, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetContextTagReal"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetContextTagReal")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (payload)
	if pullErr := readBuffer.PullContext("payload"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for payload")
	}
	_payload, _payloadErr := BACnetTagPayloadRealParse(readBuffer)
	if _payloadErr != nil {
		return nil, errors.Wrap(_payloadErr, "Error parsing 'payload' field")
	}
	payload := _payload.(BACnetTagPayloadReal)
	if closeErr := readBuffer.CloseContext("payload"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for payload")
	}

	// Virtual field
	_actualValue := payload.GetValue()
	actualValue := float32(_actualValue)
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetContextTagReal"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetContextTagReal")
	}

	// Create a partially initialized instance
	_child := &_BACnetContextTagReal{
		Payload:           payload,
		_BACnetContextTag: &_BACnetContextTag{},
	}
	_child._BACnetContextTag._BACnetContextTagChildRequirements = _child
	return _child, nil
}

func (m *_BACnetContextTagReal) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetContextTagReal"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetContextTagReal")
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
		if _actualValueErr := writeBuffer.WriteVirtual("actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetContextTagReal"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetContextTagReal")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetContextTagReal) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
