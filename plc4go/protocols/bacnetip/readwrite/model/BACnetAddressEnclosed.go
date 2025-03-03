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
	"context"
	"fmt"
	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
)

// Code generated by code-generation. DO NOT EDIT.

// BACnetAddressEnclosed is the corresponding interface of BACnetAddressEnclosed
type BACnetAddressEnclosed interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetOpeningTag returns OpeningTag (property field)
	GetOpeningTag() BACnetOpeningTag
	// GetAddress returns Address (property field)
	GetAddress() BACnetAddress
	// GetClosingTag returns ClosingTag (property field)
	GetClosingTag() BACnetClosingTag
}

// BACnetAddressEnclosedExactly can be used when we want exactly this type and not a type which fulfills BACnetAddressEnclosed.
// This is useful for switch cases.
type BACnetAddressEnclosedExactly interface {
	BACnetAddressEnclosed
	isBACnetAddressEnclosed() bool
}

// _BACnetAddressEnclosed is the data-structure of this message
type _BACnetAddressEnclosed struct {
	OpeningTag BACnetOpeningTag
	Address    BACnetAddress
	ClosingTag BACnetClosingTag

	// Arguments.
	TagNumber uint8
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetAddressEnclosed) GetOpeningTag() BACnetOpeningTag {
	return m.OpeningTag
}

func (m *_BACnetAddressEnclosed) GetAddress() BACnetAddress {
	return m.Address
}

func (m *_BACnetAddressEnclosed) GetClosingTag() BACnetClosingTag {
	return m.ClosingTag
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetAddressEnclosed factory function for _BACnetAddressEnclosed
func NewBACnetAddressEnclosed(openingTag BACnetOpeningTag, address BACnetAddress, closingTag BACnetClosingTag, tagNumber uint8) *_BACnetAddressEnclosed {
	return &_BACnetAddressEnclosed{OpeningTag: openingTag, Address: address, ClosingTag: closingTag, TagNumber: tagNumber}
}

// Deprecated: use the interface for direct cast
func CastBACnetAddressEnclosed(structType any) BACnetAddressEnclosed {
	if casted, ok := structType.(BACnetAddressEnclosed); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetAddressEnclosed); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetAddressEnclosed) GetTypeName() string {
	return "BACnetAddressEnclosed"
}

func (m *_BACnetAddressEnclosed) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (openingTag)
	lengthInBits += m.OpeningTag.GetLengthInBits(ctx)

	// Simple field (address)
	lengthInBits += m.Address.GetLengthInBits(ctx)

	// Simple field (closingTag)
	lengthInBits += m.ClosingTag.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetAddressEnclosed) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetAddressEnclosedParse(ctx context.Context, theBytes []byte, tagNumber uint8) (BACnetAddressEnclosed, error) {
	return BACnetAddressEnclosedParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), tagNumber)
}

func BACnetAddressEnclosedParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8) (BACnetAddressEnclosed, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("BACnetAddressEnclosed"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetAddressEnclosed")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (openingTag)
	if pullErr := readBuffer.PullContext("openingTag"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for openingTag")
	}
	_openingTag, _openingTagErr := BACnetOpeningTagParseWithBuffer(ctx, readBuffer, uint8(tagNumber))
	if _openingTagErr != nil {
		return nil, errors.Wrap(_openingTagErr, "Error parsing 'openingTag' field of BACnetAddressEnclosed")
	}
	openingTag := _openingTag.(BACnetOpeningTag)
	if closeErr := readBuffer.CloseContext("openingTag"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for openingTag")
	}

	// Simple Field (address)
	if pullErr := readBuffer.PullContext("address"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for address")
	}
	_address, _addressErr := BACnetAddressParseWithBuffer(ctx, readBuffer)
	if _addressErr != nil {
		return nil, errors.Wrap(_addressErr, "Error parsing 'address' field of BACnetAddressEnclosed")
	}
	address := _address.(BACnetAddress)
	if closeErr := readBuffer.CloseContext("address"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for address")
	}

	// Simple Field (closingTag)
	if pullErr := readBuffer.PullContext("closingTag"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for closingTag")
	}
	_closingTag, _closingTagErr := BACnetClosingTagParseWithBuffer(ctx, readBuffer, uint8(tagNumber))
	if _closingTagErr != nil {
		return nil, errors.Wrap(_closingTagErr, "Error parsing 'closingTag' field of BACnetAddressEnclosed")
	}
	closingTag := _closingTag.(BACnetClosingTag)
	if closeErr := readBuffer.CloseContext("closingTag"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for closingTag")
	}

	if closeErr := readBuffer.CloseContext("BACnetAddressEnclosed"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetAddressEnclosed")
	}

	// Create the instance
	return &_BACnetAddressEnclosed{
		TagNumber:  tagNumber,
		OpeningTag: openingTag,
		Address:    address,
		ClosingTag: closingTag,
	}, nil
}

func (m *_BACnetAddressEnclosed) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetAddressEnclosed) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("BACnetAddressEnclosed"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetAddressEnclosed")
	}

	// Simple Field (openingTag)
	if pushErr := writeBuffer.PushContext("openingTag"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for openingTag")
	}
	_openingTagErr := writeBuffer.WriteSerializable(ctx, m.GetOpeningTag())
	if popErr := writeBuffer.PopContext("openingTag"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for openingTag")
	}
	if _openingTagErr != nil {
		return errors.Wrap(_openingTagErr, "Error serializing 'openingTag' field")
	}

	// Simple Field (address)
	if pushErr := writeBuffer.PushContext("address"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for address")
	}
	_addressErr := writeBuffer.WriteSerializable(ctx, m.GetAddress())
	if popErr := writeBuffer.PopContext("address"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for address")
	}
	if _addressErr != nil {
		return errors.Wrap(_addressErr, "Error serializing 'address' field")
	}

	// Simple Field (closingTag)
	if pushErr := writeBuffer.PushContext("closingTag"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for closingTag")
	}
	_closingTagErr := writeBuffer.WriteSerializable(ctx, m.GetClosingTag())
	if popErr := writeBuffer.PopContext("closingTag"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for closingTag")
	}
	if _closingTagErr != nil {
		return errors.Wrap(_closingTagErr, "Error serializing 'closingTag' field")
	}

	if popErr := writeBuffer.PopContext("BACnetAddressEnclosed"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetAddressEnclosed")
	}
	return nil
}

////
// Arguments Getter

func (m *_BACnetAddressEnclosed) GetTagNumber() uint8 {
	return m.TagNumber
}

//
////

func (m *_BACnetAddressEnclosed) isBACnetAddressEnclosed() bool {
	return true
}

func (m *_BACnetAddressEnclosed) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
