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

// IdentifyReplyCommandMinimumLevels is the corresponding interface of IdentifyReplyCommandMinimumLevels
type IdentifyReplyCommandMinimumLevels interface {
	utils.LengthAware
	utils.Serializable
	IdentifyReplyCommand
}

// _IdentifyReplyCommandMinimumLevels is the data-structure of this message
type _IdentifyReplyCommandMinimumLevels struct {
	*_IdentifyReplyCommand
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_IdentifyReplyCommandMinimumLevels) GetAttribute() Attribute {
	return Attribute_MinimumLevels
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_IdentifyReplyCommandMinimumLevels) InitializeParent(parent IdentifyReplyCommand) {}

func (m *_IdentifyReplyCommandMinimumLevels) GetParent() IdentifyReplyCommand {
	return m._IdentifyReplyCommand
}

// NewIdentifyReplyCommandMinimumLevels factory function for _IdentifyReplyCommandMinimumLevels
func NewIdentifyReplyCommandMinimumLevels() *_IdentifyReplyCommandMinimumLevels {
	_result := &_IdentifyReplyCommandMinimumLevels{
		_IdentifyReplyCommand: NewIdentifyReplyCommand(),
	}
	_result._IdentifyReplyCommand._IdentifyReplyCommandChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastIdentifyReplyCommandMinimumLevels(structType interface{}) IdentifyReplyCommandMinimumLevels {
	if casted, ok := structType.(IdentifyReplyCommandMinimumLevels); ok {
		return casted
	}
	if casted, ok := structType.(*IdentifyReplyCommandMinimumLevels); ok {
		return *casted
	}
	return nil
}

func (m *_IdentifyReplyCommandMinimumLevels) GetTypeName() string {
	return "IdentifyReplyCommandMinimumLevels"
}

func (m *_IdentifyReplyCommandMinimumLevels) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_IdentifyReplyCommandMinimumLevels) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	return lengthInBits
}

func (m *_IdentifyReplyCommandMinimumLevels) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func IdentifyReplyCommandMinimumLevelsParse(readBuffer utils.ReadBuffer, attribute Attribute) (IdentifyReplyCommandMinimumLevels, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("IdentifyReplyCommandMinimumLevels"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for IdentifyReplyCommandMinimumLevels")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("IdentifyReplyCommandMinimumLevels"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for IdentifyReplyCommandMinimumLevels")
	}

	// Create a partially initialized instance
	_child := &_IdentifyReplyCommandMinimumLevels{
		_IdentifyReplyCommand: &_IdentifyReplyCommand{},
	}
	_child._IdentifyReplyCommand._IdentifyReplyCommandChildRequirements = _child
	return _child, nil
}

func (m *_IdentifyReplyCommandMinimumLevels) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("IdentifyReplyCommandMinimumLevels"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for IdentifyReplyCommandMinimumLevels")
		}

		if popErr := writeBuffer.PopContext("IdentifyReplyCommandMinimumLevels"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for IdentifyReplyCommandMinimumLevels")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_IdentifyReplyCommandMinimumLevels) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
