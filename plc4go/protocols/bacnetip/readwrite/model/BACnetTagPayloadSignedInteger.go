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

// BACnetTagPayloadSignedInteger is the corresponding interface of BACnetTagPayloadSignedInteger
type BACnetTagPayloadSignedInteger interface {
	utils.LengthAware
	utils.Serializable
	// GetValueInt8 returns ValueInt8 (property field)
	GetValueInt8() *int8
	// GetValueInt16 returns ValueInt16 (property field)
	GetValueInt16() *int16
	// GetValueInt24 returns ValueInt24 (property field)
	GetValueInt24() *int32
	// GetValueInt32 returns ValueInt32 (property field)
	GetValueInt32() *int32
	// GetValueInt40 returns ValueInt40 (property field)
	GetValueInt40() *int64
	// GetValueInt48 returns ValueInt48 (property field)
	GetValueInt48() *int64
	// GetValueInt56 returns ValueInt56 (property field)
	GetValueInt56() *int64
	// GetValueInt64 returns ValueInt64 (property field)
	GetValueInt64() *int64
	// GetIsInt8 returns IsInt8 (virtual field)
	GetIsInt8() bool
	// GetIsInt16 returns IsInt16 (virtual field)
	GetIsInt16() bool
	// GetIsInt24 returns IsInt24 (virtual field)
	GetIsInt24() bool
	// GetIsInt32 returns IsInt32 (virtual field)
	GetIsInt32() bool
	// GetIsInt40 returns IsInt40 (virtual field)
	GetIsInt40() bool
	// GetIsInt48 returns IsInt48 (virtual field)
	GetIsInt48() bool
	// GetIsInt56 returns IsInt56 (virtual field)
	GetIsInt56() bool
	// GetIsInt64 returns IsInt64 (virtual field)
	GetIsInt64() bool
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() uint64
}

// _BACnetTagPayloadSignedInteger is the data-structure of this message
type _BACnetTagPayloadSignedInteger struct {
	ValueInt8  *int8
	ValueInt16 *int16
	ValueInt24 *int32
	ValueInt32 *int32
	ValueInt40 *int64
	ValueInt48 *int64
	ValueInt56 *int64
	ValueInt64 *int64

	// Arguments.
	ActualLength uint32
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetTagPayloadSignedInteger) GetValueInt8() *int8 {
	return m.ValueInt8
}

func (m *_BACnetTagPayloadSignedInteger) GetValueInt16() *int16 {
	return m.ValueInt16
}

func (m *_BACnetTagPayloadSignedInteger) GetValueInt24() *int32 {
	return m.ValueInt24
}

func (m *_BACnetTagPayloadSignedInteger) GetValueInt32() *int32 {
	return m.ValueInt32
}

func (m *_BACnetTagPayloadSignedInteger) GetValueInt40() *int64 {
	return m.ValueInt40
}

func (m *_BACnetTagPayloadSignedInteger) GetValueInt48() *int64 {
	return m.ValueInt48
}

func (m *_BACnetTagPayloadSignedInteger) GetValueInt56() *int64 {
	return m.ValueInt56
}

func (m *_BACnetTagPayloadSignedInteger) GetValueInt64() *int64 {
	return m.ValueInt64
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetTagPayloadSignedInteger) GetIsInt8() bool {
	valueInt8 := m.ValueInt8
	_ = valueInt8
	valueInt16 := m.ValueInt16
	_ = valueInt16
	valueInt24 := m.ValueInt24
	_ = valueInt24
	valueInt32 := m.ValueInt32
	_ = valueInt32
	valueInt40 := m.ValueInt40
	_ = valueInt40
	valueInt48 := m.ValueInt48
	_ = valueInt48
	valueInt56 := m.ValueInt56
	_ = valueInt56
	valueInt64 := m.ValueInt64
	_ = valueInt64
	return bool(bool((m.ActualLength) == (1)))
}

func (m *_BACnetTagPayloadSignedInteger) GetIsInt16() bool {
	valueInt8 := m.ValueInt8
	_ = valueInt8
	valueInt16 := m.ValueInt16
	_ = valueInt16
	valueInt24 := m.ValueInt24
	_ = valueInt24
	valueInt32 := m.ValueInt32
	_ = valueInt32
	valueInt40 := m.ValueInt40
	_ = valueInt40
	valueInt48 := m.ValueInt48
	_ = valueInt48
	valueInt56 := m.ValueInt56
	_ = valueInt56
	valueInt64 := m.ValueInt64
	_ = valueInt64
	return bool(bool((m.ActualLength) == (2)))
}

func (m *_BACnetTagPayloadSignedInteger) GetIsInt24() bool {
	valueInt8 := m.ValueInt8
	_ = valueInt8
	valueInt16 := m.ValueInt16
	_ = valueInt16
	valueInt24 := m.ValueInt24
	_ = valueInt24
	valueInt32 := m.ValueInt32
	_ = valueInt32
	valueInt40 := m.ValueInt40
	_ = valueInt40
	valueInt48 := m.ValueInt48
	_ = valueInt48
	valueInt56 := m.ValueInt56
	_ = valueInt56
	valueInt64 := m.ValueInt64
	_ = valueInt64
	return bool(bool((m.ActualLength) == (3)))
}

func (m *_BACnetTagPayloadSignedInteger) GetIsInt32() bool {
	valueInt8 := m.ValueInt8
	_ = valueInt8
	valueInt16 := m.ValueInt16
	_ = valueInt16
	valueInt24 := m.ValueInt24
	_ = valueInt24
	valueInt32 := m.ValueInt32
	_ = valueInt32
	valueInt40 := m.ValueInt40
	_ = valueInt40
	valueInt48 := m.ValueInt48
	_ = valueInt48
	valueInt56 := m.ValueInt56
	_ = valueInt56
	valueInt64 := m.ValueInt64
	_ = valueInt64
	return bool(bool((m.ActualLength) == (4)))
}

func (m *_BACnetTagPayloadSignedInteger) GetIsInt40() bool {
	valueInt8 := m.ValueInt8
	_ = valueInt8
	valueInt16 := m.ValueInt16
	_ = valueInt16
	valueInt24 := m.ValueInt24
	_ = valueInt24
	valueInt32 := m.ValueInt32
	_ = valueInt32
	valueInt40 := m.ValueInt40
	_ = valueInt40
	valueInt48 := m.ValueInt48
	_ = valueInt48
	valueInt56 := m.ValueInt56
	_ = valueInt56
	valueInt64 := m.ValueInt64
	_ = valueInt64
	return bool(bool((m.ActualLength) == (5)))
}

func (m *_BACnetTagPayloadSignedInteger) GetIsInt48() bool {
	valueInt8 := m.ValueInt8
	_ = valueInt8
	valueInt16 := m.ValueInt16
	_ = valueInt16
	valueInt24 := m.ValueInt24
	_ = valueInt24
	valueInt32 := m.ValueInt32
	_ = valueInt32
	valueInt40 := m.ValueInt40
	_ = valueInt40
	valueInt48 := m.ValueInt48
	_ = valueInt48
	valueInt56 := m.ValueInt56
	_ = valueInt56
	valueInt64 := m.ValueInt64
	_ = valueInt64
	return bool(bool((m.ActualLength) == (6)))
}

func (m *_BACnetTagPayloadSignedInteger) GetIsInt56() bool {
	valueInt8 := m.ValueInt8
	_ = valueInt8
	valueInt16 := m.ValueInt16
	_ = valueInt16
	valueInt24 := m.ValueInt24
	_ = valueInt24
	valueInt32 := m.ValueInt32
	_ = valueInt32
	valueInt40 := m.ValueInt40
	_ = valueInt40
	valueInt48 := m.ValueInt48
	_ = valueInt48
	valueInt56 := m.ValueInt56
	_ = valueInt56
	valueInt64 := m.ValueInt64
	_ = valueInt64
	return bool(bool((m.ActualLength) == (7)))
}

func (m *_BACnetTagPayloadSignedInteger) GetIsInt64() bool {
	valueInt8 := m.ValueInt8
	_ = valueInt8
	valueInt16 := m.ValueInt16
	_ = valueInt16
	valueInt24 := m.ValueInt24
	_ = valueInt24
	valueInt32 := m.ValueInt32
	_ = valueInt32
	valueInt40 := m.ValueInt40
	_ = valueInt40
	valueInt48 := m.ValueInt48
	_ = valueInt48
	valueInt56 := m.ValueInt56
	_ = valueInt56
	valueInt64 := m.ValueInt64
	_ = valueInt64
	return bool(bool((m.ActualLength) == (8)))
}

func (m *_BACnetTagPayloadSignedInteger) GetActualValue() uint64 {
	valueInt8 := m.ValueInt8
	_ = valueInt8
	valueInt16 := m.ValueInt16
	_ = valueInt16
	valueInt24 := m.ValueInt24
	_ = valueInt24
	valueInt32 := m.ValueInt32
	_ = valueInt32
	valueInt40 := m.ValueInt40
	_ = valueInt40
	valueInt48 := m.ValueInt48
	_ = valueInt48
	valueInt56 := m.ValueInt56
	_ = valueInt56
	valueInt64 := m.ValueInt64
	_ = valueInt64
	return uint64(utils.InlineIf(m.GetIsInt8(), func() interface{} { return uint64((*m.GetValueInt8())) }, func() interface{} {
		return uint64(uint64(utils.InlineIf(m.GetIsInt16(), func() interface{} { return uint64((*m.GetValueInt16())) }, func() interface{} {
			return uint64(uint64(utils.InlineIf(m.GetIsInt24(), func() interface{} { return uint64((*m.GetValueInt24())) }, func() interface{} {
				return uint64(uint64(utils.InlineIf(m.GetIsInt32(), func() interface{} { return uint64((*m.GetValueInt32())) }, func() interface{} {
					return uint64(uint64(utils.InlineIf(m.GetIsInt40(), func() interface{} { return uint64((*m.GetValueInt40())) }, func() interface{} {
						return uint64(uint64(utils.InlineIf(m.GetIsInt48(), func() interface{} { return uint64((*m.GetValueInt48())) }, func() interface{} {
							return uint64(uint64(utils.InlineIf(m.GetIsInt56(), func() interface{} { return uint64((*m.GetValueInt56())) }, func() interface{} { return uint64((*m.GetValueInt64())) }).(uint64)))
						}).(uint64)))
					}).(uint64)))
				}).(uint64)))
			}).(uint64)))
		}).(uint64)))
	}).(uint64))
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetTagPayloadSignedInteger factory function for _BACnetTagPayloadSignedInteger
func NewBACnetTagPayloadSignedInteger(valueInt8 *int8, valueInt16 *int16, valueInt24 *int32, valueInt32 *int32, valueInt40 *int64, valueInt48 *int64, valueInt56 *int64, valueInt64 *int64, actualLength uint32) *_BACnetTagPayloadSignedInteger {
	return &_BACnetTagPayloadSignedInteger{ValueInt8: valueInt8, ValueInt16: valueInt16, ValueInt24: valueInt24, ValueInt32: valueInt32, ValueInt40: valueInt40, ValueInt48: valueInt48, ValueInt56: valueInt56, ValueInt64: valueInt64, ActualLength: actualLength}
}

// Deprecated: use the interface for direct cast
func CastBACnetTagPayloadSignedInteger(structType interface{}) BACnetTagPayloadSignedInteger {
	if casted, ok := structType.(BACnetTagPayloadSignedInteger); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetTagPayloadSignedInteger); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetTagPayloadSignedInteger) GetTypeName() string {
	return "BACnetTagPayloadSignedInteger"
}

func (m *_BACnetTagPayloadSignedInteger) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetTagPayloadSignedInteger) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(0)

	// A virtual field doesn't have any in- or output.

	// Optional Field (valueInt8)
	if m.ValueInt8 != nil {
		lengthInBits += 8
	}

	// A virtual field doesn't have any in- or output.

	// Optional Field (valueInt16)
	if m.ValueInt16 != nil {
		lengthInBits += 16
	}

	// A virtual field doesn't have any in- or output.

	// Optional Field (valueInt24)
	if m.ValueInt24 != nil {
		lengthInBits += 24
	}

	// A virtual field doesn't have any in- or output.

	// Optional Field (valueInt32)
	if m.ValueInt32 != nil {
		lengthInBits += 32
	}

	// A virtual field doesn't have any in- or output.

	// Optional Field (valueInt40)
	if m.ValueInt40 != nil {
		lengthInBits += 40
	}

	// A virtual field doesn't have any in- or output.

	// Optional Field (valueInt48)
	if m.ValueInt48 != nil {
		lengthInBits += 48
	}

	// A virtual field doesn't have any in- or output.

	// Optional Field (valueInt56)
	if m.ValueInt56 != nil {
		lengthInBits += 56
	}

	// A virtual field doesn't have any in- or output.

	// Optional Field (valueInt64)
	if m.ValueInt64 != nil {
		lengthInBits += 64
	}

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetTagPayloadSignedInteger) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetTagPayloadSignedIntegerParse(readBuffer utils.ReadBuffer, actualLength uint32) (BACnetTagPayloadSignedInteger, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetTagPayloadSignedInteger"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetTagPayloadSignedInteger")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Virtual field
	_isInt8 := bool((actualLength) == (1))
	isInt8 := bool(_isInt8)
	_ = isInt8

	// Optional Field (valueInt8) (Can be skipped, if a given expression evaluates to false)
	var valueInt8 *int8 = nil
	if isInt8 {
		_val, _err := readBuffer.ReadInt8("valueInt8", 8)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'valueInt8' field")
		}
		valueInt8 = &_val
	}

	// Virtual field
	_isInt16 := bool((actualLength) == (2))
	isInt16 := bool(_isInt16)
	_ = isInt16

	// Optional Field (valueInt16) (Can be skipped, if a given expression evaluates to false)
	var valueInt16 *int16 = nil
	if isInt16 {
		_val, _err := readBuffer.ReadInt16("valueInt16", 16)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'valueInt16' field")
		}
		valueInt16 = &_val
	}

	// Virtual field
	_isInt24 := bool((actualLength) == (3))
	isInt24 := bool(_isInt24)
	_ = isInt24

	// Optional Field (valueInt24) (Can be skipped, if a given expression evaluates to false)
	var valueInt24 *int32 = nil
	if isInt24 {
		_val, _err := readBuffer.ReadInt32("valueInt24", 24)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'valueInt24' field")
		}
		valueInt24 = &_val
	}

	// Virtual field
	_isInt32 := bool((actualLength) == (4))
	isInt32 := bool(_isInt32)
	_ = isInt32

	// Optional Field (valueInt32) (Can be skipped, if a given expression evaluates to false)
	var valueInt32 *int32 = nil
	if isInt32 {
		_val, _err := readBuffer.ReadInt32("valueInt32", 32)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'valueInt32' field")
		}
		valueInt32 = &_val
	}

	// Virtual field
	_isInt40 := bool((actualLength) == (5))
	isInt40 := bool(_isInt40)
	_ = isInt40

	// Optional Field (valueInt40) (Can be skipped, if a given expression evaluates to false)
	var valueInt40 *int64 = nil
	if isInt40 {
		_val, _err := readBuffer.ReadInt64("valueInt40", 40)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'valueInt40' field")
		}
		valueInt40 = &_val
	}

	// Virtual field
	_isInt48 := bool((actualLength) == (6))
	isInt48 := bool(_isInt48)
	_ = isInt48

	// Optional Field (valueInt48) (Can be skipped, if a given expression evaluates to false)
	var valueInt48 *int64 = nil
	if isInt48 {
		_val, _err := readBuffer.ReadInt64("valueInt48", 48)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'valueInt48' field")
		}
		valueInt48 = &_val
	}

	// Virtual field
	_isInt56 := bool((actualLength) == (7))
	isInt56 := bool(_isInt56)
	_ = isInt56

	// Optional Field (valueInt56) (Can be skipped, if a given expression evaluates to false)
	var valueInt56 *int64 = nil
	if isInt56 {
		_val, _err := readBuffer.ReadInt64("valueInt56", 56)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'valueInt56' field")
		}
		valueInt56 = &_val
	}

	// Virtual field
	_isInt64 := bool((actualLength) == (8))
	isInt64 := bool(_isInt64)
	_ = isInt64

	// Optional Field (valueInt64) (Can be skipped, if a given expression evaluates to false)
	var valueInt64 *int64 = nil
	if isInt64 {
		_val, _err := readBuffer.ReadInt64("valueInt64", 64)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'valueInt64' field")
		}
		valueInt64 = &_val
	}

	// Validation
	if !(bool(bool(bool(bool(bool(bool(bool(isInt8) || bool(isInt16)) || bool(isInt24)) || bool(isInt32)) || bool(isInt40)) || bool(isInt48)) || bool(isInt56)) || bool(isInt64)) {
		return nil, errors.WithStack(utils.ParseValidationError{"unmapped integer length"})
	}

	// Virtual field
	_actualValue := utils.InlineIf(isInt8, func() interface{} { return uint64((*valueInt8)) }, func() interface{} {
		return uint64(uint64(utils.InlineIf(isInt16, func() interface{} { return uint64((*valueInt16)) }, func() interface{} {
			return uint64(uint64(utils.InlineIf(isInt24, func() interface{} { return uint64((*valueInt24)) }, func() interface{} {
				return uint64(uint64(utils.InlineIf(isInt32, func() interface{} { return uint64((*valueInt32)) }, func() interface{} {
					return uint64(uint64(utils.InlineIf(isInt40, func() interface{} { return uint64((*valueInt40)) }, func() interface{} {
						return uint64(uint64(utils.InlineIf(isInt48, func() interface{} { return uint64((*valueInt48)) }, func() interface{} {
							return uint64(uint64(utils.InlineIf(isInt56, func() interface{} { return uint64((*valueInt56)) }, func() interface{} { return uint64((*valueInt64)) }).(uint64)))
						}).(uint64)))
					}).(uint64)))
				}).(uint64)))
			}).(uint64)))
		}).(uint64)))
	}).(uint64)
	actualValue := uint64(_actualValue)
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetTagPayloadSignedInteger"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetTagPayloadSignedInteger")
	}

	// Create the instance
	return NewBACnetTagPayloadSignedInteger(valueInt8, valueInt16, valueInt24, valueInt32, valueInt40, valueInt48, valueInt56, valueInt64, actualLength), nil
}

func (m *_BACnetTagPayloadSignedInteger) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("BACnetTagPayloadSignedInteger"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetTagPayloadSignedInteger")
	}
	// Virtual field
	if _isInt8Err := writeBuffer.WriteVirtual("isInt8", m.GetIsInt8()); _isInt8Err != nil {
		return errors.Wrap(_isInt8Err, "Error serializing 'isInt8' field")
	}

	// Optional Field (valueInt8) (Can be skipped, if the value is null)
	var valueInt8 *int8 = nil
	if m.GetValueInt8() != nil {
		valueInt8 = m.GetValueInt8()
		_valueInt8Err := writeBuffer.WriteInt8("valueInt8", 8, *(valueInt8))
		if _valueInt8Err != nil {
			return errors.Wrap(_valueInt8Err, "Error serializing 'valueInt8' field")
		}
	}
	// Virtual field
	if _isInt16Err := writeBuffer.WriteVirtual("isInt16", m.GetIsInt16()); _isInt16Err != nil {
		return errors.Wrap(_isInt16Err, "Error serializing 'isInt16' field")
	}

	// Optional Field (valueInt16) (Can be skipped, if the value is null)
	var valueInt16 *int16 = nil
	if m.GetValueInt16() != nil {
		valueInt16 = m.GetValueInt16()
		_valueInt16Err := writeBuffer.WriteInt16("valueInt16", 16, *(valueInt16))
		if _valueInt16Err != nil {
			return errors.Wrap(_valueInt16Err, "Error serializing 'valueInt16' field")
		}
	}
	// Virtual field
	if _isInt24Err := writeBuffer.WriteVirtual("isInt24", m.GetIsInt24()); _isInt24Err != nil {
		return errors.Wrap(_isInt24Err, "Error serializing 'isInt24' field")
	}

	// Optional Field (valueInt24) (Can be skipped, if the value is null)
	var valueInt24 *int32 = nil
	if m.GetValueInt24() != nil {
		valueInt24 = m.GetValueInt24()
		_valueInt24Err := writeBuffer.WriteInt32("valueInt24", 24, *(valueInt24))
		if _valueInt24Err != nil {
			return errors.Wrap(_valueInt24Err, "Error serializing 'valueInt24' field")
		}
	}
	// Virtual field
	if _isInt32Err := writeBuffer.WriteVirtual("isInt32", m.GetIsInt32()); _isInt32Err != nil {
		return errors.Wrap(_isInt32Err, "Error serializing 'isInt32' field")
	}

	// Optional Field (valueInt32) (Can be skipped, if the value is null)
	var valueInt32 *int32 = nil
	if m.GetValueInt32() != nil {
		valueInt32 = m.GetValueInt32()
		_valueInt32Err := writeBuffer.WriteInt32("valueInt32", 32, *(valueInt32))
		if _valueInt32Err != nil {
			return errors.Wrap(_valueInt32Err, "Error serializing 'valueInt32' field")
		}
	}
	// Virtual field
	if _isInt40Err := writeBuffer.WriteVirtual("isInt40", m.GetIsInt40()); _isInt40Err != nil {
		return errors.Wrap(_isInt40Err, "Error serializing 'isInt40' field")
	}

	// Optional Field (valueInt40) (Can be skipped, if the value is null)
	var valueInt40 *int64 = nil
	if m.GetValueInt40() != nil {
		valueInt40 = m.GetValueInt40()
		_valueInt40Err := writeBuffer.WriteInt64("valueInt40", 40, *(valueInt40))
		if _valueInt40Err != nil {
			return errors.Wrap(_valueInt40Err, "Error serializing 'valueInt40' field")
		}
	}
	// Virtual field
	if _isInt48Err := writeBuffer.WriteVirtual("isInt48", m.GetIsInt48()); _isInt48Err != nil {
		return errors.Wrap(_isInt48Err, "Error serializing 'isInt48' field")
	}

	// Optional Field (valueInt48) (Can be skipped, if the value is null)
	var valueInt48 *int64 = nil
	if m.GetValueInt48() != nil {
		valueInt48 = m.GetValueInt48()
		_valueInt48Err := writeBuffer.WriteInt64("valueInt48", 48, *(valueInt48))
		if _valueInt48Err != nil {
			return errors.Wrap(_valueInt48Err, "Error serializing 'valueInt48' field")
		}
	}
	// Virtual field
	if _isInt56Err := writeBuffer.WriteVirtual("isInt56", m.GetIsInt56()); _isInt56Err != nil {
		return errors.Wrap(_isInt56Err, "Error serializing 'isInt56' field")
	}

	// Optional Field (valueInt56) (Can be skipped, if the value is null)
	var valueInt56 *int64 = nil
	if m.GetValueInt56() != nil {
		valueInt56 = m.GetValueInt56()
		_valueInt56Err := writeBuffer.WriteInt64("valueInt56", 56, *(valueInt56))
		if _valueInt56Err != nil {
			return errors.Wrap(_valueInt56Err, "Error serializing 'valueInt56' field")
		}
	}
	// Virtual field
	if _isInt64Err := writeBuffer.WriteVirtual("isInt64", m.GetIsInt64()); _isInt64Err != nil {
		return errors.Wrap(_isInt64Err, "Error serializing 'isInt64' field")
	}

	// Optional Field (valueInt64) (Can be skipped, if the value is null)
	var valueInt64 *int64 = nil
	if m.GetValueInt64() != nil {
		valueInt64 = m.GetValueInt64()
		_valueInt64Err := writeBuffer.WriteInt64("valueInt64", 64, *(valueInt64))
		if _valueInt64Err != nil {
			return errors.Wrap(_valueInt64Err, "Error serializing 'valueInt64' field")
		}
	}
	// Virtual field
	if _actualValueErr := writeBuffer.WriteVirtual("actualValue", m.GetActualValue()); _actualValueErr != nil {
		return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
	}

	if popErr := writeBuffer.PopContext("BACnetTagPayloadSignedInteger"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetTagPayloadSignedInteger")
	}
	return nil
}

func (m *_BACnetTagPayloadSignedInteger) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
