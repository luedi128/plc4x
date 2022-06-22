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

// ModbusPDUReadWriteMultipleHoldingRegistersRequest is the corresponding interface of ModbusPDUReadWriteMultipleHoldingRegistersRequest
type ModbusPDUReadWriteMultipleHoldingRegistersRequest interface {
	utils.LengthAware
	utils.Serializable
	ModbusPDU
	// GetReadStartingAddress returns ReadStartingAddress (property field)
	GetReadStartingAddress() uint16
	// GetReadQuantity returns ReadQuantity (property field)
	GetReadQuantity() uint16
	// GetWriteStartingAddress returns WriteStartingAddress (property field)
	GetWriteStartingAddress() uint16
	// GetWriteQuantity returns WriteQuantity (property field)
	GetWriteQuantity() uint16
	// GetValue returns Value (property field)
	GetValue() []byte
}

// _ModbusPDUReadWriteMultipleHoldingRegistersRequest is the data-structure of this message
type _ModbusPDUReadWriteMultipleHoldingRegistersRequest struct {
	*_ModbusPDU
	ReadStartingAddress  uint16
	ReadQuantity         uint16
	WriteStartingAddress uint16
	WriteQuantity        uint16
	Value                []byte
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ModbusPDUReadWriteMultipleHoldingRegistersRequest) GetErrorFlag() bool {
	return bool(false)
}

func (m *_ModbusPDUReadWriteMultipleHoldingRegistersRequest) GetFunctionFlag() uint8 {
	return 0x17
}

func (m *_ModbusPDUReadWriteMultipleHoldingRegistersRequest) GetResponse() bool {
	return bool(false)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ModbusPDUReadWriteMultipleHoldingRegistersRequest) InitializeParent(parent ModbusPDU) {}

func (m *_ModbusPDUReadWriteMultipleHoldingRegistersRequest) GetParent() ModbusPDU {
	return m._ModbusPDU
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_ModbusPDUReadWriteMultipleHoldingRegistersRequest) GetReadStartingAddress() uint16 {
	return m.ReadStartingAddress
}

func (m *_ModbusPDUReadWriteMultipleHoldingRegistersRequest) GetReadQuantity() uint16 {
	return m.ReadQuantity
}

func (m *_ModbusPDUReadWriteMultipleHoldingRegistersRequest) GetWriteStartingAddress() uint16 {
	return m.WriteStartingAddress
}

func (m *_ModbusPDUReadWriteMultipleHoldingRegistersRequest) GetWriteQuantity() uint16 {
	return m.WriteQuantity
}

func (m *_ModbusPDUReadWriteMultipleHoldingRegistersRequest) GetValue() []byte {
	return m.Value
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewModbusPDUReadWriteMultipleHoldingRegistersRequest factory function for _ModbusPDUReadWriteMultipleHoldingRegistersRequest
func NewModbusPDUReadWriteMultipleHoldingRegistersRequest(readStartingAddress uint16, readQuantity uint16, writeStartingAddress uint16, writeQuantity uint16, value []byte) *_ModbusPDUReadWriteMultipleHoldingRegistersRequest {
	_result := &_ModbusPDUReadWriteMultipleHoldingRegistersRequest{
		ReadStartingAddress:  readStartingAddress,
		ReadQuantity:         readQuantity,
		WriteStartingAddress: writeStartingAddress,
		WriteQuantity:        writeQuantity,
		Value:                value,
		_ModbusPDU:           NewModbusPDU(),
	}
	_result._ModbusPDU._ModbusPDUChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastModbusPDUReadWriteMultipleHoldingRegistersRequest(structType interface{}) ModbusPDUReadWriteMultipleHoldingRegistersRequest {
	if casted, ok := structType.(ModbusPDUReadWriteMultipleHoldingRegistersRequest); ok {
		return casted
	}
	if casted, ok := structType.(*ModbusPDUReadWriteMultipleHoldingRegistersRequest); ok {
		return *casted
	}
	return nil
}

func (m *_ModbusPDUReadWriteMultipleHoldingRegistersRequest) GetTypeName() string {
	return "ModbusPDUReadWriteMultipleHoldingRegistersRequest"
}

func (m *_ModbusPDUReadWriteMultipleHoldingRegistersRequest) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_ModbusPDUReadWriteMultipleHoldingRegistersRequest) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (readStartingAddress)
	lengthInBits += 16

	// Simple field (readQuantity)
	lengthInBits += 16

	// Simple field (writeStartingAddress)
	lengthInBits += 16

	// Simple field (writeQuantity)
	lengthInBits += 16

	// Implicit Field (byteCount)
	lengthInBits += 8

	// Array field
	if len(m.Value) > 0 {
		lengthInBits += 8 * uint16(len(m.Value))
	}

	return lengthInBits
}

func (m *_ModbusPDUReadWriteMultipleHoldingRegistersRequest) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func ModbusPDUReadWriteMultipleHoldingRegistersRequestParse(readBuffer utils.ReadBuffer, response bool) (ModbusPDUReadWriteMultipleHoldingRegistersRequest, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ModbusPDUReadWriteMultipleHoldingRegistersRequest"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ModbusPDUReadWriteMultipleHoldingRegistersRequest")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (readStartingAddress)
	_readStartingAddress, _readStartingAddressErr := readBuffer.ReadUint16("readStartingAddress", 16)
	if _readStartingAddressErr != nil {
		return nil, errors.Wrap(_readStartingAddressErr, "Error parsing 'readStartingAddress' field")
	}
	readStartingAddress := _readStartingAddress

	// Simple Field (readQuantity)
	_readQuantity, _readQuantityErr := readBuffer.ReadUint16("readQuantity", 16)
	if _readQuantityErr != nil {
		return nil, errors.Wrap(_readQuantityErr, "Error parsing 'readQuantity' field")
	}
	readQuantity := _readQuantity

	// Simple Field (writeStartingAddress)
	_writeStartingAddress, _writeStartingAddressErr := readBuffer.ReadUint16("writeStartingAddress", 16)
	if _writeStartingAddressErr != nil {
		return nil, errors.Wrap(_writeStartingAddressErr, "Error parsing 'writeStartingAddress' field")
	}
	writeStartingAddress := _writeStartingAddress

	// Simple Field (writeQuantity)
	_writeQuantity, _writeQuantityErr := readBuffer.ReadUint16("writeQuantity", 16)
	if _writeQuantityErr != nil {
		return nil, errors.Wrap(_writeQuantityErr, "Error parsing 'writeQuantity' field")
	}
	writeQuantity := _writeQuantity

	// Implicit Field (byteCount) (Used for parsing, but its value is not stored as it's implicitly given by the objects content)
	byteCount, _byteCountErr := readBuffer.ReadUint8("byteCount", 8)
	_ = byteCount
	if _byteCountErr != nil {
		return nil, errors.Wrap(_byteCountErr, "Error parsing 'byteCount' field")
	}
	// Byte Array field (value)
	numberOfBytesvalue := int(byteCount)
	value, _readArrayErr := readBuffer.ReadByteArray("value", numberOfBytesvalue)
	if _readArrayErr != nil {
		return nil, errors.Wrap(_readArrayErr, "Error parsing 'value' field")
	}

	if closeErr := readBuffer.CloseContext("ModbusPDUReadWriteMultipleHoldingRegistersRequest"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ModbusPDUReadWriteMultipleHoldingRegistersRequest")
	}

	// Create a partially initialized instance
	_child := &_ModbusPDUReadWriteMultipleHoldingRegistersRequest{
		ReadStartingAddress:  readStartingAddress,
		ReadQuantity:         readQuantity,
		WriteStartingAddress: writeStartingAddress,
		WriteQuantity:        writeQuantity,
		Value:                value,
		_ModbusPDU:           &_ModbusPDU{},
	}
	_child._ModbusPDU._ModbusPDUChildRequirements = _child
	return _child, nil
}

func (m *_ModbusPDUReadWriteMultipleHoldingRegistersRequest) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ModbusPDUReadWriteMultipleHoldingRegistersRequest"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ModbusPDUReadWriteMultipleHoldingRegistersRequest")
		}

		// Simple Field (readStartingAddress)
		readStartingAddress := uint16(m.GetReadStartingAddress())
		_readStartingAddressErr := writeBuffer.WriteUint16("readStartingAddress", 16, (readStartingAddress))
		if _readStartingAddressErr != nil {
			return errors.Wrap(_readStartingAddressErr, "Error serializing 'readStartingAddress' field")
		}

		// Simple Field (readQuantity)
		readQuantity := uint16(m.GetReadQuantity())
		_readQuantityErr := writeBuffer.WriteUint16("readQuantity", 16, (readQuantity))
		if _readQuantityErr != nil {
			return errors.Wrap(_readQuantityErr, "Error serializing 'readQuantity' field")
		}

		// Simple Field (writeStartingAddress)
		writeStartingAddress := uint16(m.GetWriteStartingAddress())
		_writeStartingAddressErr := writeBuffer.WriteUint16("writeStartingAddress", 16, (writeStartingAddress))
		if _writeStartingAddressErr != nil {
			return errors.Wrap(_writeStartingAddressErr, "Error serializing 'writeStartingAddress' field")
		}

		// Simple Field (writeQuantity)
		writeQuantity := uint16(m.GetWriteQuantity())
		_writeQuantityErr := writeBuffer.WriteUint16("writeQuantity", 16, (writeQuantity))
		if _writeQuantityErr != nil {
			return errors.Wrap(_writeQuantityErr, "Error serializing 'writeQuantity' field")
		}

		// Implicit Field (byteCount) (Used for parsing, but it's value is not stored as it's implicitly given by the objects content)
		byteCount := uint8(uint8(len(m.GetValue())))
		_byteCountErr := writeBuffer.WriteUint8("byteCount", 8, (byteCount))
		if _byteCountErr != nil {
			return errors.Wrap(_byteCountErr, "Error serializing 'byteCount' field")
		}

		// Array Field (value)
		if m.GetValue() != nil {
			// Byte Array field (value)
			_writeArrayErr := writeBuffer.WriteByteArray("value", m.GetValue())
			if _writeArrayErr != nil {
				return errors.Wrap(_writeArrayErr, "Error serializing 'value' field")
			}
		}

		if popErr := writeBuffer.PopContext("ModbusPDUReadWriteMultipleHoldingRegistersRequest"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ModbusPDUReadWriteMultipleHoldingRegistersRequest")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_ModbusPDUReadWriteMultipleHoldingRegistersRequest) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
