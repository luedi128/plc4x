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

// CBusCommandPointToPointToMultiPoint is the corresponding interface of CBusCommandPointToPointToMultiPoint
type CBusCommandPointToPointToMultiPoint interface {
	utils.LengthAware
	utils.Serializable
	CBusCommand
	// GetCommand returns Command (property field)
	GetCommand() CBusPointToPointToMultipointCommand
}

// _CBusCommandPointToPointToMultiPoint is the data-structure of this message
type _CBusCommandPointToPointToMultiPoint struct {
	*_CBusCommand
	Command CBusPointToPointToMultipointCommand

	// Arguments.
	Srchk bool
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_CBusCommandPointToPointToMultiPoint) InitializeParent(parent CBusCommand, header CBusHeader) {
	m.Header = header
}

func (m *_CBusCommandPointToPointToMultiPoint) GetParent() CBusCommand {
	return m._CBusCommand
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_CBusCommandPointToPointToMultiPoint) GetCommand() CBusPointToPointToMultipointCommand {
	return m.Command
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewCBusCommandPointToPointToMultiPoint factory function for _CBusCommandPointToPointToMultiPoint
func NewCBusCommandPointToPointToMultiPoint(command CBusPointToPointToMultipointCommand, header CBusHeader, srchk bool) *_CBusCommandPointToPointToMultiPoint {
	_result := &_CBusCommandPointToPointToMultiPoint{
		Command:      command,
		_CBusCommand: NewCBusCommand(header, srchk),
	}
	_result._CBusCommand._CBusCommandChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastCBusCommandPointToPointToMultiPoint(structType interface{}) CBusCommandPointToPointToMultiPoint {
	if casted, ok := structType.(CBusCommandPointToPointToMultiPoint); ok {
		return casted
	}
	if casted, ok := structType.(*CBusCommandPointToPointToMultiPoint); ok {
		return *casted
	}
	return nil
}

func (m *_CBusCommandPointToPointToMultiPoint) GetTypeName() string {
	return "CBusCommandPointToPointToMultiPoint"
}

func (m *_CBusCommandPointToPointToMultiPoint) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_CBusCommandPointToPointToMultiPoint) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (command)
	lengthInBits += m.Command.GetLengthInBits()

	return lengthInBits
}

func (m *_CBusCommandPointToPointToMultiPoint) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func CBusCommandPointToPointToMultiPointParse(readBuffer utils.ReadBuffer, srchk bool) (CBusCommandPointToPointToMultiPoint, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("CBusCommandPointToPointToMultiPoint"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for CBusCommandPointToPointToMultiPoint")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (command)
	if pullErr := readBuffer.PullContext("command"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for command")
	}
	_command, _commandErr := CBusPointToPointToMultipointCommandParse(readBuffer, bool(srchk))
	if _commandErr != nil {
		return nil, errors.Wrap(_commandErr, "Error parsing 'command' field")
	}
	command := _command.(CBusPointToPointToMultipointCommand)
	if closeErr := readBuffer.CloseContext("command"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for command")
	}

	if closeErr := readBuffer.CloseContext("CBusCommandPointToPointToMultiPoint"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for CBusCommandPointToPointToMultiPoint")
	}

	// Create a partially initialized instance
	_child := &_CBusCommandPointToPointToMultiPoint{
		Command:      command,
		_CBusCommand: &_CBusCommand{},
	}
	_child._CBusCommand._CBusCommandChildRequirements = _child
	return _child, nil
}

func (m *_CBusCommandPointToPointToMultiPoint) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("CBusCommandPointToPointToMultiPoint"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for CBusCommandPointToPointToMultiPoint")
		}

		// Simple Field (command)
		if pushErr := writeBuffer.PushContext("command"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for command")
		}
		_commandErr := writeBuffer.WriteSerializable(m.GetCommand())
		if popErr := writeBuffer.PopContext("command"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for command")
		}
		if _commandErr != nil {
			return errors.Wrap(_commandErr, "Error serializing 'command' field")
		}

		if popErr := writeBuffer.PopContext("CBusCommandPointToPointToMultiPoint"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for CBusCommandPointToPointToMultiPoint")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_CBusCommandPointToPointToMultiPoint) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
