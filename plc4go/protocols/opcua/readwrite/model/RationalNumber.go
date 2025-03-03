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

// RationalNumber is the corresponding interface of RationalNumber
type RationalNumber interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ExtensionObjectDefinition
	// GetNumerator returns Numerator (property field)
	GetNumerator() int32
	// GetDenominator returns Denominator (property field)
	GetDenominator() uint32
}

// RationalNumberExactly can be used when we want exactly this type and not a type which fulfills RationalNumber.
// This is useful for switch cases.
type RationalNumberExactly interface {
	RationalNumber
	isRationalNumber() bool
}

// _RationalNumber is the data-structure of this message
type _RationalNumber struct {
	*_ExtensionObjectDefinition
	Numerator   int32
	Denominator uint32
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_RationalNumber) GetIdentifier() string {
	return "18808"
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_RationalNumber) InitializeParent(parent ExtensionObjectDefinition) {}

func (m *_RationalNumber) GetParent() ExtensionObjectDefinition {
	return m._ExtensionObjectDefinition
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_RationalNumber) GetNumerator() int32 {
	return m.Numerator
}

func (m *_RationalNumber) GetDenominator() uint32 {
	return m.Denominator
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewRationalNumber factory function for _RationalNumber
func NewRationalNumber(numerator int32, denominator uint32) *_RationalNumber {
	_result := &_RationalNumber{
		Numerator:                  numerator,
		Denominator:                denominator,
		_ExtensionObjectDefinition: NewExtensionObjectDefinition(),
	}
	_result._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastRationalNumber(structType any) RationalNumber {
	if casted, ok := structType.(RationalNumber); ok {
		return casted
	}
	if casted, ok := structType.(*RationalNumber); ok {
		return *casted
	}
	return nil
}

func (m *_RationalNumber) GetTypeName() string {
	return "RationalNumber"
}

func (m *_RationalNumber) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (numerator)
	lengthInBits += 32

	// Simple field (denominator)
	lengthInBits += 32

	return lengthInBits
}

func (m *_RationalNumber) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func RationalNumberParse(ctx context.Context, theBytes []byte, identifier string) (RationalNumber, error) {
	return RationalNumberParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), identifier)
}

func RationalNumberParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, identifier string) (RationalNumber, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("RationalNumber"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for RationalNumber")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (numerator)
	_numerator, _numeratorErr := readBuffer.ReadInt32("numerator", 32)
	if _numeratorErr != nil {
		return nil, errors.Wrap(_numeratorErr, "Error parsing 'numerator' field of RationalNumber")
	}
	numerator := _numerator

	// Simple Field (denominator)
	_denominator, _denominatorErr := readBuffer.ReadUint32("denominator", 32)
	if _denominatorErr != nil {
		return nil, errors.Wrap(_denominatorErr, "Error parsing 'denominator' field of RationalNumber")
	}
	denominator := _denominator

	if closeErr := readBuffer.CloseContext("RationalNumber"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for RationalNumber")
	}

	// Create a partially initialized instance
	_child := &_RationalNumber{
		_ExtensionObjectDefinition: &_ExtensionObjectDefinition{},
		Numerator:                  numerator,
		Denominator:                denominator,
	}
	_child._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _child
	return _child, nil
}

func (m *_RationalNumber) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_RationalNumber) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("RationalNumber"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for RationalNumber")
		}

		// Simple Field (numerator)
		numerator := int32(m.GetNumerator())
		_numeratorErr := writeBuffer.WriteInt32("numerator", 32, int32((numerator)))
		if _numeratorErr != nil {
			return errors.Wrap(_numeratorErr, "Error serializing 'numerator' field")
		}

		// Simple Field (denominator)
		denominator := uint32(m.GetDenominator())
		_denominatorErr := writeBuffer.WriteUint32("denominator", 32, uint32((denominator)))
		if _denominatorErr != nil {
			return errors.Wrap(_denominatorErr, "Error serializing 'denominator' field")
		}

		if popErr := writeBuffer.PopContext("RationalNumber"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for RationalNumber")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_RationalNumber) isRationalNumber() bool {
	return true
}

func (m *_RationalNumber) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
