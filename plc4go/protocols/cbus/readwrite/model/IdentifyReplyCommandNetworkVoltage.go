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
	"fmt"
	"github.com/apache/plc4x/plc4go/internal/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// Constant values.
const IdentifyReplyCommandNetworkVoltage_DOT byte = 0x2C
const IdentifyReplyCommandNetworkVoltage_V byte = 0x56

// IdentifyReplyCommandNetworkVoltage is the corresponding interface of IdentifyReplyCommandNetworkVoltage
type IdentifyReplyCommandNetworkVoltage interface {
	utils.LengthAware
	utils.Serializable
	IdentifyReplyCommand
	// GetVolts returns Volts (property field)
	GetVolts() string
	// GetVoltsDecimalPlace returns VoltsDecimalPlace (property field)
	GetVoltsDecimalPlace() string
}

// IdentifyReplyCommandNetworkVoltageExactly can be used when we want exactly this type and not a type which fulfills IdentifyReplyCommandNetworkVoltage.
// This is useful for switch cases.
type IdentifyReplyCommandNetworkVoltageExactly interface {
	IdentifyReplyCommandNetworkVoltage
	isIdentifyReplyCommandNetworkVoltage() bool
}

// _IdentifyReplyCommandNetworkVoltage is the data-structure of this message
type _IdentifyReplyCommandNetworkVoltage struct {
	*_IdentifyReplyCommand
	Volts             string
	VoltsDecimalPlace string
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_IdentifyReplyCommandNetworkVoltage) GetAttribute() Attribute {
	return Attribute_NetworkVoltage
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_IdentifyReplyCommandNetworkVoltage) InitializeParent(parent IdentifyReplyCommand) {}

func (m *_IdentifyReplyCommandNetworkVoltage) GetParent() IdentifyReplyCommand {
	return m._IdentifyReplyCommand
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_IdentifyReplyCommandNetworkVoltage) GetVolts() string {
	return m.Volts
}

func (m *_IdentifyReplyCommandNetworkVoltage) GetVoltsDecimalPlace() string {
	return m.VoltsDecimalPlace
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for const fields.
///////////////////////

func (m *_IdentifyReplyCommandNetworkVoltage) GetDot() byte {
	return IdentifyReplyCommandNetworkVoltage_DOT
}

func (m *_IdentifyReplyCommandNetworkVoltage) GetV() byte {
	return IdentifyReplyCommandNetworkVoltage_V
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewIdentifyReplyCommandNetworkVoltage factory function for _IdentifyReplyCommandNetworkVoltage
func NewIdentifyReplyCommandNetworkVoltage(volts string, voltsDecimalPlace string, numBytes uint8) *_IdentifyReplyCommandNetworkVoltage {
	_result := &_IdentifyReplyCommandNetworkVoltage{
		Volts:                 volts,
		VoltsDecimalPlace:     voltsDecimalPlace,
		_IdentifyReplyCommand: NewIdentifyReplyCommand(numBytes),
	}
	_result._IdentifyReplyCommand._IdentifyReplyCommandChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastIdentifyReplyCommandNetworkVoltage(structType interface{}) IdentifyReplyCommandNetworkVoltage {
	if casted, ok := structType.(IdentifyReplyCommandNetworkVoltage); ok {
		return casted
	}
	if casted, ok := structType.(*IdentifyReplyCommandNetworkVoltage); ok {
		return *casted
	}
	return nil
}

func (m *_IdentifyReplyCommandNetworkVoltage) GetTypeName() string {
	return "IdentifyReplyCommandNetworkVoltage"
}

func (m *_IdentifyReplyCommandNetworkVoltage) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_IdentifyReplyCommandNetworkVoltage) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (volts)
	lengthInBits += 2

	// Const Field (dot)
	lengthInBits += 8

	// Simple field (voltsDecimalPlace)
	lengthInBits += 2

	// Const Field (v)
	lengthInBits += 8

	return lengthInBits
}

func (m *_IdentifyReplyCommandNetworkVoltage) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func IdentifyReplyCommandNetworkVoltageParse(readBuffer utils.ReadBuffer, attribute Attribute, numBytes uint8) (IdentifyReplyCommandNetworkVoltage, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("IdentifyReplyCommandNetworkVoltage"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for IdentifyReplyCommandNetworkVoltage")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (volts)
	_volts, _voltsErr := readBuffer.ReadString("volts", uint32(2))
	if _voltsErr != nil {
		return nil, errors.Wrap(_voltsErr, "Error parsing 'volts' field of IdentifyReplyCommandNetworkVoltage")
	}
	volts := _volts

	// Const Field (dot)
	dot, _dotErr := readBuffer.ReadByte("dot")
	if _dotErr != nil {
		return nil, errors.Wrap(_dotErr, "Error parsing 'dot' field of IdentifyReplyCommandNetworkVoltage")
	}
	if dot != IdentifyReplyCommandNetworkVoltage_DOT {
		return nil, errors.New("Expected constant value " + fmt.Sprintf("%d", IdentifyReplyCommandNetworkVoltage_DOT) + " but got " + fmt.Sprintf("%d", dot))
	}

	// Simple Field (voltsDecimalPlace)
	_voltsDecimalPlace, _voltsDecimalPlaceErr := readBuffer.ReadString("voltsDecimalPlace", uint32(2))
	if _voltsDecimalPlaceErr != nil {
		return nil, errors.Wrap(_voltsDecimalPlaceErr, "Error parsing 'voltsDecimalPlace' field of IdentifyReplyCommandNetworkVoltage")
	}
	voltsDecimalPlace := _voltsDecimalPlace

	// Const Field (v)
	v, _vErr := readBuffer.ReadByte("v")
	if _vErr != nil {
		return nil, errors.Wrap(_vErr, "Error parsing 'v' field of IdentifyReplyCommandNetworkVoltage")
	}
	if v != IdentifyReplyCommandNetworkVoltage_V {
		return nil, errors.New("Expected constant value " + fmt.Sprintf("%d", IdentifyReplyCommandNetworkVoltage_V) + " but got " + fmt.Sprintf("%d", v))
	}

	if closeErr := readBuffer.CloseContext("IdentifyReplyCommandNetworkVoltage"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for IdentifyReplyCommandNetworkVoltage")
	}

	// Create a partially initialized instance
	_child := &_IdentifyReplyCommandNetworkVoltage{
		Volts:             volts,
		VoltsDecimalPlace: voltsDecimalPlace,
		_IdentifyReplyCommand: &_IdentifyReplyCommand{
			NumBytes: numBytes,
		},
	}
	_child._IdentifyReplyCommand._IdentifyReplyCommandChildRequirements = _child
	return _child, nil
}

func (m *_IdentifyReplyCommandNetworkVoltage) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("IdentifyReplyCommandNetworkVoltage"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for IdentifyReplyCommandNetworkVoltage")
		}

		// Simple Field (volts)
		volts := string(m.GetVolts())
		_voltsErr := writeBuffer.WriteString("volts", uint32(2), "UTF-8", (volts))
		if _voltsErr != nil {
			return errors.Wrap(_voltsErr, "Error serializing 'volts' field")
		}

		// Const Field (dot)
		_dotErr := writeBuffer.WriteByte("dot", 0x2C)
		if _dotErr != nil {
			return errors.Wrap(_dotErr, "Error serializing 'dot' field")
		}

		// Simple Field (voltsDecimalPlace)
		voltsDecimalPlace := string(m.GetVoltsDecimalPlace())
		_voltsDecimalPlaceErr := writeBuffer.WriteString("voltsDecimalPlace", uint32(2), "UTF-8", (voltsDecimalPlace))
		if _voltsDecimalPlaceErr != nil {
			return errors.Wrap(_voltsDecimalPlaceErr, "Error serializing 'voltsDecimalPlace' field")
		}

		// Const Field (v)
		_vErr := writeBuffer.WriteByte("v", 0x56)
		if _vErr != nil {
			return errors.Wrap(_vErr, "Error serializing 'v' field")
		}

		if popErr := writeBuffer.PopContext("IdentifyReplyCommandNetworkVoltage"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for IdentifyReplyCommandNetworkVoltage")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_IdentifyReplyCommandNetworkVoltage) isIdentifyReplyCommandNetworkVoltage() bool {
	return true
}

func (m *_IdentifyReplyCommandNetworkVoltage) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
