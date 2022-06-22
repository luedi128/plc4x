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

// Reply is the corresponding interface of Reply
type Reply interface {
	utils.LengthAware
	utils.Serializable
	// GetMagicByte returns MagicByte (property field)
	GetMagicByte() byte
}

// _Reply is the data-structure of this message
type _Reply struct {
	_ReplyChildRequirements
	MagicByte byte
}

type _ReplyChildRequirements interface {
	GetLengthInBits() uint16
	GetLengthInBitsConditional(lastItem bool) uint16
	Serialize(writeBuffer utils.WriteBuffer) error
}

type ReplyParent interface {
	SerializeParent(writeBuffer utils.WriteBuffer, child Reply, serializeChildFunction func() error) error
	GetTypeName() string
}

type ReplyChild interface {
	Serialize(writeBuffer utils.WriteBuffer) error
	InitializeParent(parent Reply, magicByte byte)
	GetParent() *Reply

	GetTypeName() string
	Reply
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_Reply) GetMagicByte() byte {
	return m.MagicByte
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewReply factory function for _Reply
func NewReply(magicByte byte) *_Reply {
	return &_Reply{MagicByte: magicByte}
}

// Deprecated: use the interface for direct cast
func CastReply(structType interface{}) Reply {
	if casted, ok := structType.(Reply); ok {
		return casted
	}
	if casted, ok := structType.(*Reply); ok {
		return *casted
	}
	return nil
}

func (m *_Reply) GetTypeName() string {
	return "Reply"
}

func (m *_Reply) GetParentLengthInBits() uint16 {
	lengthInBits := uint16(0)

	return lengthInBits
}

func (m *_Reply) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func ReplyParse(readBuffer utils.ReadBuffer) (Reply, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("Reply"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for Reply")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Peek Field (magicByte)
	currentPos = positionAware.GetPos()
	magicByte, _err := readBuffer.ReadByte("magicByte")
	if _err != nil {
		return nil, errors.Wrap(_err, "Error parsing 'magicByte' field")
	}

	readBuffer.Reset(currentPos)

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	type ReplyChildSerializeRequirement interface {
		Reply
		InitializeParent(Reply, byte)
		GetParent() Reply
	}
	var _childTemp interface{}
	var _child ReplyChildSerializeRequirement
	var typeSwitchError error
	switch {
	case magicByte == 0x0: // CALReplyReply
		_childTemp, typeSwitchError = CALReplyReplyParse(readBuffer)
	case magicByte == 0x0: // MonitoredSALReply
		_childTemp, typeSwitchError = MonitoredSALReplyParse(readBuffer)
	case magicByte == 0x0: // ConfirmationReply
		_childTemp, typeSwitchError = ConfirmationReplyParse(readBuffer)
	case magicByte == 0x0: // PowerUpReply
		_childTemp, typeSwitchError = PowerUpReplyParse(readBuffer)
	case magicByte == 0x0: // ParameterChangeReply
		_childTemp, typeSwitchError = ParameterChangeReplyParse(readBuffer)
	case magicByte == 0x0: // ExclamationMarkReply
		_childTemp, typeSwitchError = ExclamationMarkReplyParse(readBuffer)
	default:
		// TODO: return actual type
		typeSwitchError = errors.New("Unmapped type")
	}
	if typeSwitchError != nil {
		return nil, errors.Wrap(typeSwitchError, "Error parsing sub-type for type-switch.")
	}
	_child = _childTemp.(ReplyChildSerializeRequirement)

	if closeErr := readBuffer.CloseContext("Reply"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for Reply")
	}

	// Finish initializing
	_child.InitializeParent(_child, magicByte)
	return _child, nil
}

func (pm *_Reply) SerializeParent(writeBuffer utils.WriteBuffer, child Reply, serializeChildFunction func() error) error {
	// We redirect all calls through client as some methods are only implemented there
	m := child
	_ = m
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("Reply"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for Reply")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	if _typeSwitchErr := serializeChildFunction(); _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	if popErr := writeBuffer.PopContext("Reply"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for Reply")
	}
	return nil
}

func (m *_Reply) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
