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

// S7PayloadNotify8 is the corresponding interface of S7PayloadNotify8
type S7PayloadNotify8 interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	S7PayloadUserDataItem
	// GetAlarmMessage returns AlarmMessage (property field)
	GetAlarmMessage() AlarmMessagePushType
}

// S7PayloadNotify8Exactly can be used when we want exactly this type and not a type which fulfills S7PayloadNotify8.
// This is useful for switch cases.
type S7PayloadNotify8Exactly interface {
	S7PayloadNotify8
	isS7PayloadNotify8() bool
}

// _S7PayloadNotify8 is the data-structure of this message
type _S7PayloadNotify8 struct {
	*_S7PayloadUserDataItem
	AlarmMessage AlarmMessagePushType
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_S7PayloadNotify8) GetCpuFunctionGroup() uint8 {
	return 0x04
}

func (m *_S7PayloadNotify8) GetCpuFunctionType() uint8 {
	return 0x00
}

func (m *_S7PayloadNotify8) GetCpuSubfunction() uint8 {
	return 0x16
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_S7PayloadNotify8) InitializeParent(parent S7PayloadUserDataItem, returnCode DataTransportErrorCode, transportSize DataTransportSize, dataLength uint16) {
	m.ReturnCode = returnCode
	m.TransportSize = transportSize
	m.DataLength = dataLength
}

func (m *_S7PayloadNotify8) GetParent() S7PayloadUserDataItem {
	return m._S7PayloadUserDataItem
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_S7PayloadNotify8) GetAlarmMessage() AlarmMessagePushType {
	return m.AlarmMessage
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewS7PayloadNotify8 factory function for _S7PayloadNotify8
func NewS7PayloadNotify8(alarmMessage AlarmMessagePushType, returnCode DataTransportErrorCode, transportSize DataTransportSize, dataLength uint16) *_S7PayloadNotify8 {
	_result := &_S7PayloadNotify8{
		AlarmMessage:           alarmMessage,
		_S7PayloadUserDataItem: NewS7PayloadUserDataItem(returnCode, transportSize, dataLength),
	}
	_result._S7PayloadUserDataItem._S7PayloadUserDataItemChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastS7PayloadNotify8(structType any) S7PayloadNotify8 {
	if casted, ok := structType.(S7PayloadNotify8); ok {
		return casted
	}
	if casted, ok := structType.(*S7PayloadNotify8); ok {
		return *casted
	}
	return nil
}

func (m *_S7PayloadNotify8) GetTypeName() string {
	return "S7PayloadNotify8"
}

func (m *_S7PayloadNotify8) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (alarmMessage)
	lengthInBits += m.AlarmMessage.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_S7PayloadNotify8) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func S7PayloadNotify8Parse(ctx context.Context, theBytes []byte, cpuFunctionGroup uint8, cpuFunctionType uint8, cpuSubfunction uint8) (S7PayloadNotify8, error) {
	return S7PayloadNotify8ParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), cpuFunctionGroup, cpuFunctionType, cpuSubfunction)
}

func S7PayloadNotify8ParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, cpuFunctionGroup uint8, cpuFunctionType uint8, cpuSubfunction uint8) (S7PayloadNotify8, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("S7PayloadNotify8"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for S7PayloadNotify8")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (alarmMessage)
	if pullErr := readBuffer.PullContext("alarmMessage"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for alarmMessage")
	}
	_alarmMessage, _alarmMessageErr := AlarmMessagePushTypeParseWithBuffer(ctx, readBuffer)
	if _alarmMessageErr != nil {
		return nil, errors.Wrap(_alarmMessageErr, "Error parsing 'alarmMessage' field of S7PayloadNotify8")
	}
	alarmMessage := _alarmMessage.(AlarmMessagePushType)
	if closeErr := readBuffer.CloseContext("alarmMessage"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for alarmMessage")
	}

	if closeErr := readBuffer.CloseContext("S7PayloadNotify8"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for S7PayloadNotify8")
	}

	// Create a partially initialized instance
	_child := &_S7PayloadNotify8{
		_S7PayloadUserDataItem: &_S7PayloadUserDataItem{},
		AlarmMessage:           alarmMessage,
	}
	_child._S7PayloadUserDataItem._S7PayloadUserDataItemChildRequirements = _child
	return _child, nil
}

func (m *_S7PayloadNotify8) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_S7PayloadNotify8) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("S7PayloadNotify8"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for S7PayloadNotify8")
		}

		// Simple Field (alarmMessage)
		if pushErr := writeBuffer.PushContext("alarmMessage"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for alarmMessage")
		}
		_alarmMessageErr := writeBuffer.WriteSerializable(ctx, m.GetAlarmMessage())
		if popErr := writeBuffer.PopContext("alarmMessage"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for alarmMessage")
		}
		if _alarmMessageErr != nil {
			return errors.Wrap(_alarmMessageErr, "Error serializing 'alarmMessage' field")
		}

		if popErr := writeBuffer.PopContext("S7PayloadNotify8"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for S7PayloadNotify8")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_S7PayloadNotify8) isS7PayloadNotify8() bool {
	return true
}

func (m *_S7PayloadNotify8) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
