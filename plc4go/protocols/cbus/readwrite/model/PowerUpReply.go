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

// PowerUpReply is the corresponding interface of PowerUpReply
type PowerUpReply interface {
	utils.LengthAware
	utils.Serializable
	Reply
	// GetIsA returns IsA (property field)
	GetIsA() PowerUp
}

// _PowerUpReply is the data-structure of this message
type _PowerUpReply struct {
	*_Reply
	IsA PowerUp
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_PowerUpReply) InitializeParent(parent Reply, magicByte byte) {
	m.MagicByte = magicByte
}

func (m *_PowerUpReply) GetParent() Reply {
	return m._Reply
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_PowerUpReply) GetIsA() PowerUp {
	return m.IsA
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewPowerUpReply factory function for _PowerUpReply
func NewPowerUpReply(isA PowerUp, magicByte byte) *_PowerUpReply {
	_result := &_PowerUpReply{
		IsA:    isA,
		_Reply: NewReply(magicByte),
	}
	_result._Reply._ReplyChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastPowerUpReply(structType interface{}) PowerUpReply {
	if casted, ok := structType.(PowerUpReply); ok {
		return casted
	}
	if casted, ok := structType.(*PowerUpReply); ok {
		return *casted
	}
	return nil
}

func (m *_PowerUpReply) GetTypeName() string {
	return "PowerUpReply"
}

func (m *_PowerUpReply) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_PowerUpReply) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (isA)
	lengthInBits += m.IsA.GetLengthInBits()

	return lengthInBits
}

func (m *_PowerUpReply) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func PowerUpReplyParse(readBuffer utils.ReadBuffer) (PowerUpReply, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("PowerUpReply"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for PowerUpReply")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (isA)
	if pullErr := readBuffer.PullContext("isA"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for isA")
	}
	_isA, _isAErr := PowerUpParse(readBuffer)
	if _isAErr != nil {
		return nil, errors.Wrap(_isAErr, "Error parsing 'isA' field")
	}
	isA := _isA.(PowerUp)
	if closeErr := readBuffer.CloseContext("isA"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for isA")
	}

	if closeErr := readBuffer.CloseContext("PowerUpReply"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for PowerUpReply")
	}

	// Create a partially initialized instance
	_child := &_PowerUpReply{
		IsA:    isA,
		_Reply: &_Reply{},
	}
	_child._Reply._ReplyChildRequirements = _child
	return _child, nil
}

func (m *_PowerUpReply) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("PowerUpReply"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for PowerUpReply")
		}

		// Simple Field (isA)
		if pushErr := writeBuffer.PushContext("isA"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for isA")
		}
		_isAErr := writeBuffer.WriteSerializable(m.GetIsA())
		if popErr := writeBuffer.PopContext("isA"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for isA")
		}
		if _isAErr != nil {
			return errors.Wrap(_isAErr, "Error serializing 'isA' field")
		}

		if popErr := writeBuffer.PopContext("PowerUpReply"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for PowerUpReply")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_PowerUpReply) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
