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

// BACnetHostAddressEnclosed is the corresponding interface of BACnetHostAddressEnclosed
type BACnetHostAddressEnclosed interface {
	utils.LengthAware
	utils.Serializable
	// GetOpeningTag returns OpeningTag (property field)
	GetOpeningTag() BACnetOpeningTag
	// GetHostAddress returns HostAddress (property field)
	GetHostAddress() BACnetHostAddress
	// GetClosingTag returns ClosingTag (property field)
	GetClosingTag() BACnetClosingTag
}

// _BACnetHostAddressEnclosed is the data-structure of this message
type _BACnetHostAddressEnclosed struct {
	OpeningTag  BACnetOpeningTag
	HostAddress BACnetHostAddress
	ClosingTag  BACnetClosingTag

	// Arguments.
	TagNumber uint8
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetHostAddressEnclosed) GetOpeningTag() BACnetOpeningTag {
	return m.OpeningTag
}

func (m *_BACnetHostAddressEnclosed) GetHostAddress() BACnetHostAddress {
	return m.HostAddress
}

func (m *_BACnetHostAddressEnclosed) GetClosingTag() BACnetClosingTag {
	return m.ClosingTag
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetHostAddressEnclosed factory function for _BACnetHostAddressEnclosed
func NewBACnetHostAddressEnclosed(openingTag BACnetOpeningTag, hostAddress BACnetHostAddress, closingTag BACnetClosingTag, tagNumber uint8) *_BACnetHostAddressEnclosed {
	return &_BACnetHostAddressEnclosed{OpeningTag: openingTag, HostAddress: hostAddress, ClosingTag: closingTag, TagNumber: tagNumber}
}

// Deprecated: use the interface for direct cast
func CastBACnetHostAddressEnclosed(structType interface{}) BACnetHostAddressEnclosed {
	if casted, ok := structType.(BACnetHostAddressEnclosed); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetHostAddressEnclosed); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetHostAddressEnclosed) GetTypeName() string {
	return "BACnetHostAddressEnclosed"
}

func (m *_BACnetHostAddressEnclosed) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetHostAddressEnclosed) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(0)

	// Simple field (openingTag)
	lengthInBits += m.OpeningTag.GetLengthInBits()

	// Simple field (hostAddress)
	lengthInBits += m.HostAddress.GetLengthInBits()

	// Simple field (closingTag)
	lengthInBits += m.ClosingTag.GetLengthInBits()

	return lengthInBits
}

func (m *_BACnetHostAddressEnclosed) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetHostAddressEnclosedParse(readBuffer utils.ReadBuffer, tagNumber uint8) (BACnetHostAddressEnclosed, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetHostAddressEnclosed"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetHostAddressEnclosed")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (openingTag)
	if pullErr := readBuffer.PullContext("openingTag"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for openingTag")
	}
	_openingTag, _openingTagErr := BACnetOpeningTagParse(readBuffer, uint8(tagNumber))
	if _openingTagErr != nil {
		return nil, errors.Wrap(_openingTagErr, "Error parsing 'openingTag' field")
	}
	openingTag := _openingTag.(BACnetOpeningTag)
	if closeErr := readBuffer.CloseContext("openingTag"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for openingTag")
	}

	// Simple Field (hostAddress)
	if pullErr := readBuffer.PullContext("hostAddress"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for hostAddress")
	}
	_hostAddress, _hostAddressErr := BACnetHostAddressParse(readBuffer)
	if _hostAddressErr != nil {
		return nil, errors.Wrap(_hostAddressErr, "Error parsing 'hostAddress' field")
	}
	hostAddress := _hostAddress.(BACnetHostAddress)
	if closeErr := readBuffer.CloseContext("hostAddress"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for hostAddress")
	}

	// Simple Field (closingTag)
	if pullErr := readBuffer.PullContext("closingTag"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for closingTag")
	}
	_closingTag, _closingTagErr := BACnetClosingTagParse(readBuffer, uint8(tagNumber))
	if _closingTagErr != nil {
		return nil, errors.Wrap(_closingTagErr, "Error parsing 'closingTag' field")
	}
	closingTag := _closingTag.(BACnetClosingTag)
	if closeErr := readBuffer.CloseContext("closingTag"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for closingTag")
	}

	if closeErr := readBuffer.CloseContext("BACnetHostAddressEnclosed"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetHostAddressEnclosed")
	}

	// Create the instance
	return NewBACnetHostAddressEnclosed(openingTag, hostAddress, closingTag, tagNumber), nil
}

func (m *_BACnetHostAddressEnclosed) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("BACnetHostAddressEnclosed"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetHostAddressEnclosed")
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

	// Simple Field (hostAddress)
	if pushErr := writeBuffer.PushContext("hostAddress"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for hostAddress")
	}
	_hostAddressErr := writeBuffer.WriteSerializable(m.GetHostAddress())
	if popErr := writeBuffer.PopContext("hostAddress"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for hostAddress")
	}
	if _hostAddressErr != nil {
		return errors.Wrap(_hostAddressErr, "Error serializing 'hostAddress' field")
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

	if popErr := writeBuffer.PopContext("BACnetHostAddressEnclosed"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetHostAddressEnclosed")
	}
	return nil
}

func (m *_BACnetHostAddressEnclosed) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
