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
	"github.com/rs/zerolog/log"
)

// Code generated by code-generation. DO NOT EDIT.

// LPollData is the corresponding interface of LPollData
type LPollData interface {
	utils.LengthAware
	utils.Serializable
	LDataFrame
	// GetSourceAddress returns SourceAddress (property field)
	GetSourceAddress() KnxAddress
	// GetTargetAddress returns TargetAddress (property field)
	GetTargetAddress() []byte
	// GetNumberExpectedPollData returns NumberExpectedPollData (property field)
	GetNumberExpectedPollData() uint8
}

// _LPollData is the data-structure of this message
type _LPollData struct {
	*_LDataFrame
	SourceAddress          KnxAddress
	TargetAddress          []byte
	NumberExpectedPollData uint8
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_LPollData) GetNotAckFrame() bool {
	return bool(true)
}

func (m *_LPollData) GetPolling() bool {
	return bool(true)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_LPollData) InitializeParent(parent LDataFrame, frameType bool, notRepeated bool, priority CEMIPriority, acknowledgeRequested bool, errorFlag bool) {
	m.FrameType = frameType
	m.NotRepeated = notRepeated
	m.Priority = priority
	m.AcknowledgeRequested = acknowledgeRequested
	m.ErrorFlag = errorFlag
}

func (m *_LPollData) GetParent() LDataFrame {
	return m._LDataFrame
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_LPollData) GetSourceAddress() KnxAddress {
	return m.SourceAddress
}

func (m *_LPollData) GetTargetAddress() []byte {
	return m.TargetAddress
}

func (m *_LPollData) GetNumberExpectedPollData() uint8 {
	return m.NumberExpectedPollData
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewLPollData factory function for _LPollData
func NewLPollData(sourceAddress KnxAddress, targetAddress []byte, numberExpectedPollData uint8, frameType bool, notRepeated bool, priority CEMIPriority, acknowledgeRequested bool, errorFlag bool) *_LPollData {
	_result := &_LPollData{
		SourceAddress:          sourceAddress,
		TargetAddress:          targetAddress,
		NumberExpectedPollData: numberExpectedPollData,
		_LDataFrame:            NewLDataFrame(frameType, notRepeated, priority, acknowledgeRequested, errorFlag),
	}
	_result._LDataFrame._LDataFrameChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastLPollData(structType interface{}) LPollData {
	if casted, ok := structType.(LPollData); ok {
		return casted
	}
	if casted, ok := structType.(*LPollData); ok {
		return *casted
	}
	return nil
}

func (m *_LPollData) GetTypeName() string {
	return "LPollData"
}

func (m *_LPollData) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_LPollData) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (sourceAddress)
	lengthInBits += m.SourceAddress.GetLengthInBits()

	// Array field
	if len(m.TargetAddress) > 0 {
		lengthInBits += 8 * uint16(len(m.TargetAddress))
	}

	// Reserved Field (reserved)
	lengthInBits += 4

	// Simple field (numberExpectedPollData)
	lengthInBits += 6

	return lengthInBits
}

func (m *_LPollData) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func LPollDataParse(readBuffer utils.ReadBuffer) (LPollData, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("LPollData"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for LPollData")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (sourceAddress)
	if pullErr := readBuffer.PullContext("sourceAddress"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for sourceAddress")
	}
	_sourceAddress, _sourceAddressErr := KnxAddressParse(readBuffer)
	if _sourceAddressErr != nil {
		return nil, errors.Wrap(_sourceAddressErr, "Error parsing 'sourceAddress' field")
	}
	sourceAddress := _sourceAddress.(KnxAddress)
	if closeErr := readBuffer.CloseContext("sourceAddress"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for sourceAddress")
	}
	// Byte Array field (targetAddress)
	numberOfBytestargetAddress := int(uint16(2))
	targetAddress, _readArrayErr := readBuffer.ReadByteArray("targetAddress", numberOfBytestargetAddress)
	if _readArrayErr != nil {
		return nil, errors.Wrap(_readArrayErr, "Error parsing 'targetAddress' field")
	}

	// Reserved Field (Compartmentalized so the "reserved" variable can't leak)
	{
		reserved, _err := readBuffer.ReadUint8("reserved", 4)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'reserved' field")
		}
		if reserved != uint8(0x00) {
			log.Info().Fields(map[string]interface{}{
				"expected value": uint8(0x00),
				"got value":      reserved,
			}).Msg("Got unexpected response.")
		}
	}

	// Simple Field (numberExpectedPollData)
	_numberExpectedPollData, _numberExpectedPollDataErr := readBuffer.ReadUint8("numberExpectedPollData", 6)
	if _numberExpectedPollDataErr != nil {
		return nil, errors.Wrap(_numberExpectedPollDataErr, "Error parsing 'numberExpectedPollData' field")
	}
	numberExpectedPollData := _numberExpectedPollData

	if closeErr := readBuffer.CloseContext("LPollData"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for LPollData")
	}

	// Create a partially initialized instance
	_child := &_LPollData{
		SourceAddress:          sourceAddress,
		TargetAddress:          targetAddress,
		NumberExpectedPollData: numberExpectedPollData,
		_LDataFrame:            &_LDataFrame{},
	}
	_child._LDataFrame._LDataFrameChildRequirements = _child
	return _child, nil
}

func (m *_LPollData) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("LPollData"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for LPollData")
		}

		// Simple Field (sourceAddress)
		if pushErr := writeBuffer.PushContext("sourceAddress"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for sourceAddress")
		}
		_sourceAddressErr := writeBuffer.WriteSerializable(m.GetSourceAddress())
		if popErr := writeBuffer.PopContext("sourceAddress"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for sourceAddress")
		}
		if _sourceAddressErr != nil {
			return errors.Wrap(_sourceAddressErr, "Error serializing 'sourceAddress' field")
		}

		// Array Field (targetAddress)
		if m.GetTargetAddress() != nil {
			// Byte Array field (targetAddress)
			_writeArrayErr := writeBuffer.WriteByteArray("targetAddress", m.GetTargetAddress())
			if _writeArrayErr != nil {
				return errors.Wrap(_writeArrayErr, "Error serializing 'targetAddress' field")
			}
		}

		// Reserved Field (reserved)
		{
			_err := writeBuffer.WriteUint8("reserved", 4, uint8(0x00))
			if _err != nil {
				return errors.Wrap(_err, "Error serializing 'reserved' field")
			}
		}

		// Simple Field (numberExpectedPollData)
		numberExpectedPollData := uint8(m.GetNumberExpectedPollData())
		_numberExpectedPollDataErr := writeBuffer.WriteUint8("numberExpectedPollData", 6, (numberExpectedPollData))
		if _numberExpectedPollDataErr != nil {
			return errors.Wrap(_numberExpectedPollDataErr, "Error serializing 'numberExpectedPollData' field")
		}

		if popErr := writeBuffer.PopContext("LPollData"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for LPollData")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_LPollData) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
