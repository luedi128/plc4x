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

// BACnetPriorityValueReal is the corresponding interface of BACnetPriorityValueReal
type BACnetPriorityValueReal interface {
	utils.LengthAware
	utils.Serializable
	BACnetPriorityValue
	// GetRealValue returns RealValue (property field)
	GetRealValue() BACnetApplicationTagReal
}

// _BACnetPriorityValueReal is the data-structure of this message
type _BACnetPriorityValueReal struct {
	*_BACnetPriorityValue
	RealValue BACnetApplicationTagReal

	// Arguments.
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

func (m *_BACnetPriorityValueReal) InitializeParent(parent BACnetPriorityValue, peekedTagHeader BACnetTagHeader) {
	m.PeekedTagHeader = peekedTagHeader
}

func (m *_BACnetPriorityValueReal) GetParent() BACnetPriorityValue {
	return m._BACnetPriorityValue
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetPriorityValueReal) GetRealValue() BACnetApplicationTagReal {
	return m.RealValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetPriorityValueReal factory function for _BACnetPriorityValueReal
func NewBACnetPriorityValueReal(realValue BACnetApplicationTagReal, peekedTagHeader BACnetTagHeader, objectTypeArgument BACnetObjectType) *_BACnetPriorityValueReal {
	_result := &_BACnetPriorityValueReal{
		RealValue:            realValue,
		_BACnetPriorityValue: NewBACnetPriorityValue(peekedTagHeader, objectTypeArgument),
	}
	_result._BACnetPriorityValue._BACnetPriorityValueChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetPriorityValueReal(structType interface{}) BACnetPriorityValueReal {
	if casted, ok := structType.(BACnetPriorityValueReal); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetPriorityValueReal); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetPriorityValueReal) GetTypeName() string {
	return "BACnetPriorityValueReal"
}

func (m *_BACnetPriorityValueReal) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetPriorityValueReal) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (realValue)
	lengthInBits += m.RealValue.GetLengthInBits()

	return lengthInBits
}

func (m *_BACnetPriorityValueReal) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetPriorityValueRealParse(readBuffer utils.ReadBuffer, objectTypeArgument BACnetObjectType) (BACnetPriorityValueReal, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetPriorityValueReal"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetPriorityValueReal")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (realValue)
	if pullErr := readBuffer.PullContext("realValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for realValue")
	}
	_realValue, _realValueErr := BACnetApplicationTagParse(readBuffer)
	if _realValueErr != nil {
		return nil, errors.Wrap(_realValueErr, "Error parsing 'realValue' field")
	}
	realValue := _realValue.(BACnetApplicationTagReal)
	if closeErr := readBuffer.CloseContext("realValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for realValue")
	}

	if closeErr := readBuffer.CloseContext("BACnetPriorityValueReal"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetPriorityValueReal")
	}

	// Create a partially initialized instance
	_child := &_BACnetPriorityValueReal{
		RealValue:            realValue,
		_BACnetPriorityValue: &_BACnetPriorityValue{},
	}
	_child._BACnetPriorityValue._BACnetPriorityValueChildRequirements = _child
	return _child, nil
}

func (m *_BACnetPriorityValueReal) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetPriorityValueReal"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetPriorityValueReal")
		}

		// Simple Field (realValue)
		if pushErr := writeBuffer.PushContext("realValue"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for realValue")
		}
		_realValueErr := writeBuffer.WriteSerializable(m.GetRealValue())
		if popErr := writeBuffer.PopContext("realValue"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for realValue")
		}
		if _realValueErr != nil {
			return errors.Wrap(_realValueErr, "Error serializing 'realValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetPriorityValueReal"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetPriorityValueReal")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetPriorityValueReal) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
