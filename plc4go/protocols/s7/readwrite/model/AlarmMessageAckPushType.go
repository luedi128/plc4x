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

// AlarmMessageAckPushType is the corresponding interface of AlarmMessageAckPushType
type AlarmMessageAckPushType interface {
	utils.LengthAware
	utils.Serializable
	// GetTimeStamp returns TimeStamp (property field)
	GetTimeStamp() DateAndTime
	// GetFunctionId returns FunctionId (property field)
	GetFunctionId() uint8
	// GetNumberOfObjects returns NumberOfObjects (property field)
	GetNumberOfObjects() uint8
	// GetMessageObjects returns MessageObjects (property field)
	GetMessageObjects() []AlarmMessageAckObjectPushType
}

// _AlarmMessageAckPushType is the data-structure of this message
type _AlarmMessageAckPushType struct {
	TimeStamp       DateAndTime
	FunctionId      uint8
	NumberOfObjects uint8
	MessageObjects  []AlarmMessageAckObjectPushType
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_AlarmMessageAckPushType) GetTimeStamp() DateAndTime {
	return m.TimeStamp
}

func (m *_AlarmMessageAckPushType) GetFunctionId() uint8 {
	return m.FunctionId
}

func (m *_AlarmMessageAckPushType) GetNumberOfObjects() uint8 {
	return m.NumberOfObjects
}

func (m *_AlarmMessageAckPushType) GetMessageObjects() []AlarmMessageAckObjectPushType {
	return m.MessageObjects
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewAlarmMessageAckPushType factory function for _AlarmMessageAckPushType
func NewAlarmMessageAckPushType(TimeStamp DateAndTime, functionId uint8, numberOfObjects uint8, messageObjects []AlarmMessageAckObjectPushType) *_AlarmMessageAckPushType {
	return &_AlarmMessageAckPushType{TimeStamp: TimeStamp, FunctionId: functionId, NumberOfObjects: numberOfObjects, MessageObjects: messageObjects}
}

// Deprecated: use the interface for direct cast
func CastAlarmMessageAckPushType(structType interface{}) AlarmMessageAckPushType {
	if casted, ok := structType.(AlarmMessageAckPushType); ok {
		return casted
	}
	if casted, ok := structType.(*AlarmMessageAckPushType); ok {
		return *casted
	}
	return nil
}

func (m *_AlarmMessageAckPushType) GetTypeName() string {
	return "AlarmMessageAckPushType"
}

func (m *_AlarmMessageAckPushType) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_AlarmMessageAckPushType) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(0)

	// Simple field (TimeStamp)
	lengthInBits += m.TimeStamp.GetLengthInBits()

	// Simple field (functionId)
	lengthInBits += 8

	// Simple field (numberOfObjects)
	lengthInBits += 8

	// Array field
	if len(m.MessageObjects) > 0 {
		for i, element := range m.MessageObjects {
			last := i == len(m.MessageObjects)-1
			lengthInBits += element.(interface{ GetLengthInBitsConditional(bool) uint16 }).GetLengthInBitsConditional(last)
		}
	}

	return lengthInBits
}

func (m *_AlarmMessageAckPushType) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func AlarmMessageAckPushTypeParse(readBuffer utils.ReadBuffer) (AlarmMessageAckPushType, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("AlarmMessageAckPushType"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for AlarmMessageAckPushType")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (TimeStamp)
	if pullErr := readBuffer.PullContext("TimeStamp"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for TimeStamp")
	}
	_TimeStamp, _TimeStampErr := DateAndTimeParse(readBuffer)
	if _TimeStampErr != nil {
		return nil, errors.Wrap(_TimeStampErr, "Error parsing 'TimeStamp' field")
	}
	TimeStamp := _TimeStamp.(DateAndTime)
	if closeErr := readBuffer.CloseContext("TimeStamp"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for TimeStamp")
	}

	// Simple Field (functionId)
	_functionId, _functionIdErr := readBuffer.ReadUint8("functionId", 8)
	if _functionIdErr != nil {
		return nil, errors.Wrap(_functionIdErr, "Error parsing 'functionId' field")
	}
	functionId := _functionId

	// Simple Field (numberOfObjects)
	_numberOfObjects, _numberOfObjectsErr := readBuffer.ReadUint8("numberOfObjects", 8)
	if _numberOfObjectsErr != nil {
		return nil, errors.Wrap(_numberOfObjectsErr, "Error parsing 'numberOfObjects' field")
	}
	numberOfObjects := _numberOfObjects

	// Array field (messageObjects)
	if pullErr := readBuffer.PullContext("messageObjects", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for messageObjects")
	}
	// Count array
	messageObjects := make([]AlarmMessageAckObjectPushType, numberOfObjects)
	{
		for curItem := uint16(0); curItem < uint16(numberOfObjects); curItem++ {
			_item, _err := AlarmMessageAckObjectPushTypeParse(readBuffer)
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'messageObjects' field")
			}
			messageObjects[curItem] = _item.(AlarmMessageAckObjectPushType)
		}
	}
	if closeErr := readBuffer.CloseContext("messageObjects", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for messageObjects")
	}

	if closeErr := readBuffer.CloseContext("AlarmMessageAckPushType"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for AlarmMessageAckPushType")
	}

	// Create the instance
	return NewAlarmMessageAckPushType(TimeStamp, functionId, numberOfObjects, messageObjects), nil
}

func (m *_AlarmMessageAckPushType) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("AlarmMessageAckPushType"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for AlarmMessageAckPushType")
	}

	// Simple Field (TimeStamp)
	if pushErr := writeBuffer.PushContext("TimeStamp"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for TimeStamp")
	}
	_TimeStampErr := writeBuffer.WriteSerializable(m.GetTimeStamp())
	if popErr := writeBuffer.PopContext("TimeStamp"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for TimeStamp")
	}
	if _TimeStampErr != nil {
		return errors.Wrap(_TimeStampErr, "Error serializing 'TimeStamp' field")
	}

	// Simple Field (functionId)
	functionId := uint8(m.GetFunctionId())
	_functionIdErr := writeBuffer.WriteUint8("functionId", 8, (functionId))
	if _functionIdErr != nil {
		return errors.Wrap(_functionIdErr, "Error serializing 'functionId' field")
	}

	// Simple Field (numberOfObjects)
	numberOfObjects := uint8(m.GetNumberOfObjects())
	_numberOfObjectsErr := writeBuffer.WriteUint8("numberOfObjects", 8, (numberOfObjects))
	if _numberOfObjectsErr != nil {
		return errors.Wrap(_numberOfObjectsErr, "Error serializing 'numberOfObjects' field")
	}

	// Array Field (messageObjects)
	if m.GetMessageObjects() != nil {
		if pushErr := writeBuffer.PushContext("messageObjects", utils.WithRenderAsList(true)); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for messageObjects")
		}
		for _, _element := range m.GetMessageObjects() {
			_elementErr := writeBuffer.WriteSerializable(_element)
			if _elementErr != nil {
				return errors.Wrap(_elementErr, "Error serializing 'messageObjects' field")
			}
		}
		if popErr := writeBuffer.PopContext("messageObjects", utils.WithRenderAsList(true)); popErr != nil {
			return errors.Wrap(popErr, "Error popping for messageObjects")
		}
	}

	if popErr := writeBuffer.PopContext("AlarmMessageAckPushType"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for AlarmMessageAckPushType")
	}
	return nil
}

func (m *_AlarmMessageAckPushType) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
