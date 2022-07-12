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

// IdentifyReplyCommandDelays is the corresponding interface of IdentifyReplyCommandDelays
type IdentifyReplyCommandDelays interface {
	utils.LengthAware
	utils.Serializable
	IdentifyReplyCommand
}

// IdentifyReplyCommandDelaysExactly can be used when we want exactly this type and not a type which fulfills IdentifyReplyCommandDelays.
// This is useful for switch cases.
type IdentifyReplyCommandDelaysExactly interface {
	IdentifyReplyCommandDelays
	isIdentifyReplyCommandDelays() bool
}

// _IdentifyReplyCommandDelays is the data-structure of this message
type _IdentifyReplyCommandDelays struct {
	*_IdentifyReplyCommand
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_IdentifyReplyCommandDelays) GetAttribute() Attribute {
	return Attribute_Delays
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_IdentifyReplyCommandDelays) InitializeParent(parent IdentifyReplyCommand) {}

func (m *_IdentifyReplyCommandDelays) GetParent() IdentifyReplyCommand {
	return m._IdentifyReplyCommand
}

// NewIdentifyReplyCommandDelays factory function for _IdentifyReplyCommandDelays
func NewIdentifyReplyCommandDelays(numBytes uint8) *_IdentifyReplyCommandDelays {
	_result := &_IdentifyReplyCommandDelays{
		_IdentifyReplyCommand: NewIdentifyReplyCommand(numBytes),
	}
	_result._IdentifyReplyCommand._IdentifyReplyCommandChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastIdentifyReplyCommandDelays(structType interface{}) IdentifyReplyCommandDelays {
	if casted, ok := structType.(IdentifyReplyCommandDelays); ok {
		return casted
	}
	if casted, ok := structType.(*IdentifyReplyCommandDelays); ok {
		return *casted
	}
	return nil
}

func (m *_IdentifyReplyCommandDelays) GetTypeName() string {
	return "IdentifyReplyCommandDelays"
}

func (m *_IdentifyReplyCommandDelays) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_IdentifyReplyCommandDelays) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	return lengthInBits
}

func (m *_IdentifyReplyCommandDelays) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func IdentifyReplyCommandDelaysParse(readBuffer utils.ReadBuffer, attribute Attribute, numBytes uint8) (IdentifyReplyCommandDelays, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("IdentifyReplyCommandDelays"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for IdentifyReplyCommandDelays")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("IdentifyReplyCommandDelays"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for IdentifyReplyCommandDelays")
	}

	// Create a partially initialized instance
	_child := &_IdentifyReplyCommandDelays{
		_IdentifyReplyCommand: &_IdentifyReplyCommand{
			NumBytes: numBytes,
		},
	}
	_child._IdentifyReplyCommand._IdentifyReplyCommandChildRequirements = _child
	return _child, nil
}

func (m *_IdentifyReplyCommandDelays) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("IdentifyReplyCommandDelays"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for IdentifyReplyCommandDelays")
		}

		if popErr := writeBuffer.PopContext("IdentifyReplyCommandDelays"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for IdentifyReplyCommandDelays")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_IdentifyReplyCommandDelays) isIdentifyReplyCommandDelays() bool {
	return true
}

func (m *_IdentifyReplyCommandDelays) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
