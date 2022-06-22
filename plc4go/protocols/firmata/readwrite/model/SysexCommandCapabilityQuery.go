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

// SysexCommandCapabilityQuery is the corresponding interface of SysexCommandCapabilityQuery
type SysexCommandCapabilityQuery interface {
	utils.LengthAware
	utils.Serializable
	SysexCommand
}

// _SysexCommandCapabilityQuery is the data-structure of this message
type _SysexCommandCapabilityQuery struct {
	*_SysexCommand
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_SysexCommandCapabilityQuery) GetCommandType() uint8 {
	return 0x6B
}

func (m *_SysexCommandCapabilityQuery) GetResponse() bool {
	return false
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_SysexCommandCapabilityQuery) InitializeParent(parent SysexCommand) {}

func (m *_SysexCommandCapabilityQuery) GetParent() SysexCommand {
	return m._SysexCommand
}

// NewSysexCommandCapabilityQuery factory function for _SysexCommandCapabilityQuery
func NewSysexCommandCapabilityQuery() *_SysexCommandCapabilityQuery {
	_result := &_SysexCommandCapabilityQuery{
		_SysexCommand: NewSysexCommand(),
	}
	_result._SysexCommand._SysexCommandChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastSysexCommandCapabilityQuery(structType interface{}) SysexCommandCapabilityQuery {
	if casted, ok := structType.(SysexCommandCapabilityQuery); ok {
		return casted
	}
	if casted, ok := structType.(*SysexCommandCapabilityQuery); ok {
		return *casted
	}
	return nil
}

func (m *_SysexCommandCapabilityQuery) GetTypeName() string {
	return "SysexCommandCapabilityQuery"
}

func (m *_SysexCommandCapabilityQuery) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_SysexCommandCapabilityQuery) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	return lengthInBits
}

func (m *_SysexCommandCapabilityQuery) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func SysexCommandCapabilityQueryParse(readBuffer utils.ReadBuffer, response bool) (SysexCommandCapabilityQuery, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("SysexCommandCapabilityQuery"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for SysexCommandCapabilityQuery")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("SysexCommandCapabilityQuery"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for SysexCommandCapabilityQuery")
	}

	// Create a partially initialized instance
	_child := &_SysexCommandCapabilityQuery{
		_SysexCommand: &_SysexCommand{},
	}
	_child._SysexCommand._SysexCommandChildRequirements = _child
	return _child, nil
}

func (m *_SysexCommandCapabilityQuery) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("SysexCommandCapabilityQuery"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for SysexCommandCapabilityQuery")
		}

		if popErr := writeBuffer.PopContext("SysexCommandCapabilityQuery"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for SysexCommandCapabilityQuery")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_SysexCommandCapabilityQuery) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
