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

// IdentifyReplyCommandExtendedDiagnosticSummary is the corresponding interface of IdentifyReplyCommandExtendedDiagnosticSummary
type IdentifyReplyCommandExtendedDiagnosticSummary interface {
	utils.LengthAware
	utils.Serializable
	IdentifyReplyCommand
	// GetLowApplication returns LowApplication (property field)
	GetLowApplication() ApplicationIdContainer
	// GetHighApplication returns HighApplication (property field)
	GetHighApplication() ApplicationIdContainer
	// GetArea returns Area (property field)
	GetArea() byte
	// GetCrc returns Crc (property field)
	GetCrc() uint16
	// GetSerialNumber returns SerialNumber (property field)
	GetSerialNumber() uint32
	// GetNetworkVoltage returns NetworkVoltage (property field)
	GetNetworkVoltage() byte
	// GetUnitInLearnMode returns UnitInLearnMode (property field)
	GetUnitInLearnMode() bool
	// GetNetworkVoltageLow returns NetworkVoltageLow (property field)
	GetNetworkVoltageLow() bool
	// GetNetworkVoltageMarginal returns NetworkVoltageMarginal (property field)
	GetNetworkVoltageMarginal() bool
	// GetEnableChecksumAlarm returns EnableChecksumAlarm (property field)
	GetEnableChecksumAlarm() bool
	// GetOutputUnit returns OutputUnit (property field)
	GetOutputUnit() bool
	// GetInstallationMMIError returns InstallationMMIError (property field)
	GetInstallationMMIError() bool
	// GetEEWriteError returns EEWriteError (property field)
	GetEEWriteError() bool
	// GetEEChecksumError returns EEChecksumError (property field)
	GetEEChecksumError() bool
	// GetEEDataError returns EEDataError (property field)
	GetEEDataError() bool
	// GetMicroReset returns MicroReset (property field)
	GetMicroReset() bool
	// GetCommsTxError returns CommsTxError (property field)
	GetCommsTxError() bool
	// GetInternalStackOverflow returns InternalStackOverflow (property field)
	GetInternalStackOverflow() bool
	// GetMicroPowerReset returns MicroPowerReset (property field)
	GetMicroPowerReset() bool
	// GetNetworkVoltageInVolts returns NetworkVoltageInVolts (virtual field)
	GetNetworkVoltageInVolts() float32
}

// IdentifyReplyCommandExtendedDiagnosticSummaryExactly can be used when we want exactly this type and not a type which fulfills IdentifyReplyCommandExtendedDiagnosticSummary.
// This is useful for switch cases.
type IdentifyReplyCommandExtendedDiagnosticSummaryExactly interface {
	IdentifyReplyCommandExtendedDiagnosticSummary
	isIdentifyReplyCommandExtendedDiagnosticSummary() bool
}

// _IdentifyReplyCommandExtendedDiagnosticSummary is the data-structure of this message
type _IdentifyReplyCommandExtendedDiagnosticSummary struct {
	*_IdentifyReplyCommand
	LowApplication         ApplicationIdContainer
	HighApplication        ApplicationIdContainer
	Area                   byte
	Crc                    uint16
	SerialNumber           uint32
	NetworkVoltage         byte
	UnitInLearnMode        bool
	NetworkVoltageLow      bool
	NetworkVoltageMarginal bool
	EnableChecksumAlarm    bool
	OutputUnit             bool
	InstallationMMIError   bool
	EEWriteError           bool
	EEChecksumError        bool
	EEDataError            bool
	MicroReset             bool
	CommsTxError           bool
	InternalStackOverflow  bool
	MicroPowerReset        bool
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_IdentifyReplyCommandExtendedDiagnosticSummary) GetAttribute() Attribute {
	return Attribute_ExtendedDiagnosticSummary
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_IdentifyReplyCommandExtendedDiagnosticSummary) InitializeParent(parent IdentifyReplyCommand) {
}

func (m *_IdentifyReplyCommandExtendedDiagnosticSummary) GetParent() IdentifyReplyCommand {
	return m._IdentifyReplyCommand
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_IdentifyReplyCommandExtendedDiagnosticSummary) GetLowApplication() ApplicationIdContainer {
	return m.LowApplication
}

func (m *_IdentifyReplyCommandExtendedDiagnosticSummary) GetHighApplication() ApplicationIdContainer {
	return m.HighApplication
}

func (m *_IdentifyReplyCommandExtendedDiagnosticSummary) GetArea() byte {
	return m.Area
}

func (m *_IdentifyReplyCommandExtendedDiagnosticSummary) GetCrc() uint16 {
	return m.Crc
}

func (m *_IdentifyReplyCommandExtendedDiagnosticSummary) GetSerialNumber() uint32 {
	return m.SerialNumber
}

func (m *_IdentifyReplyCommandExtendedDiagnosticSummary) GetNetworkVoltage() byte {
	return m.NetworkVoltage
}

func (m *_IdentifyReplyCommandExtendedDiagnosticSummary) GetUnitInLearnMode() bool {
	return m.UnitInLearnMode
}

func (m *_IdentifyReplyCommandExtendedDiagnosticSummary) GetNetworkVoltageLow() bool {
	return m.NetworkVoltageLow
}

func (m *_IdentifyReplyCommandExtendedDiagnosticSummary) GetNetworkVoltageMarginal() bool {
	return m.NetworkVoltageMarginal
}

func (m *_IdentifyReplyCommandExtendedDiagnosticSummary) GetEnableChecksumAlarm() bool {
	return m.EnableChecksumAlarm
}

func (m *_IdentifyReplyCommandExtendedDiagnosticSummary) GetOutputUnit() bool {
	return m.OutputUnit
}

func (m *_IdentifyReplyCommandExtendedDiagnosticSummary) GetInstallationMMIError() bool {
	return m.InstallationMMIError
}

func (m *_IdentifyReplyCommandExtendedDiagnosticSummary) GetEEWriteError() bool {
	return m.EEWriteError
}

func (m *_IdentifyReplyCommandExtendedDiagnosticSummary) GetEEChecksumError() bool {
	return m.EEChecksumError
}

func (m *_IdentifyReplyCommandExtendedDiagnosticSummary) GetEEDataError() bool {
	return m.EEDataError
}

func (m *_IdentifyReplyCommandExtendedDiagnosticSummary) GetMicroReset() bool {
	return m.MicroReset
}

func (m *_IdentifyReplyCommandExtendedDiagnosticSummary) GetCommsTxError() bool {
	return m.CommsTxError
}

func (m *_IdentifyReplyCommandExtendedDiagnosticSummary) GetInternalStackOverflow() bool {
	return m.InternalStackOverflow
}

func (m *_IdentifyReplyCommandExtendedDiagnosticSummary) GetMicroPowerReset() bool {
	return m.MicroPowerReset
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_IdentifyReplyCommandExtendedDiagnosticSummary) GetNetworkVoltageInVolts() float32 {
	return float32(float32(m.GetNetworkVoltage()) / float32(float32(6.375)))
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewIdentifyReplyCommandExtendedDiagnosticSummary factory function for _IdentifyReplyCommandExtendedDiagnosticSummary
func NewIdentifyReplyCommandExtendedDiagnosticSummary(lowApplication ApplicationIdContainer, highApplication ApplicationIdContainer, area byte, crc uint16, serialNumber uint32, networkVoltage byte, unitInLearnMode bool, networkVoltageLow bool, networkVoltageMarginal bool, enableChecksumAlarm bool, outputUnit bool, installationMMIError bool, EEWriteError bool, EEChecksumError bool, EEDataError bool, microReset bool, commsTxError bool, internalStackOverflow bool, microPowerReset bool, numBytes uint8) *_IdentifyReplyCommandExtendedDiagnosticSummary {
	_result := &_IdentifyReplyCommandExtendedDiagnosticSummary{
		LowApplication:         lowApplication,
		HighApplication:        highApplication,
		Area:                   area,
		Crc:                    crc,
		SerialNumber:           serialNumber,
		NetworkVoltage:         networkVoltage,
		UnitInLearnMode:        unitInLearnMode,
		NetworkVoltageLow:      networkVoltageLow,
		NetworkVoltageMarginal: networkVoltageMarginal,
		EnableChecksumAlarm:    enableChecksumAlarm,
		OutputUnit:             outputUnit,
		InstallationMMIError:   installationMMIError,
		EEWriteError:           EEWriteError,
		EEChecksumError:        EEChecksumError,
		EEDataError:            EEDataError,
		MicroReset:             microReset,
		CommsTxError:           commsTxError,
		InternalStackOverflow:  internalStackOverflow,
		MicroPowerReset:        microPowerReset,
		_IdentifyReplyCommand:  NewIdentifyReplyCommand(numBytes),
	}
	_result._IdentifyReplyCommand._IdentifyReplyCommandChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastIdentifyReplyCommandExtendedDiagnosticSummary(structType interface{}) IdentifyReplyCommandExtendedDiagnosticSummary {
	if casted, ok := structType.(IdentifyReplyCommandExtendedDiagnosticSummary); ok {
		return casted
	}
	if casted, ok := structType.(*IdentifyReplyCommandExtendedDiagnosticSummary); ok {
		return *casted
	}
	return nil
}

func (m *_IdentifyReplyCommandExtendedDiagnosticSummary) GetTypeName() string {
	return "IdentifyReplyCommandExtendedDiagnosticSummary"
}

func (m *_IdentifyReplyCommandExtendedDiagnosticSummary) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_IdentifyReplyCommandExtendedDiagnosticSummary) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (lowApplication)
	lengthInBits += 8

	// Simple field (highApplication)
	lengthInBits += 8

	// Simple field (area)
	lengthInBits += 8

	// Simple field (crc)
	lengthInBits += 16

	// Simple field (serialNumber)
	lengthInBits += 32

	// Simple field (networkVoltage)
	lengthInBits += 8

	// A virtual field doesn't have any in- or output.

	// Simple field (unitInLearnMode)
	lengthInBits += 1

	// Simple field (networkVoltageLow)
	lengthInBits += 1

	// Simple field (networkVoltageMarginal)
	lengthInBits += 1

	// Reserved Field (reserved)
	lengthInBits += 1

	// Reserved Field (reserved)
	lengthInBits += 1

	// Reserved Field (reserved)
	lengthInBits += 1

	// Simple field (enableChecksumAlarm)
	lengthInBits += 1

	// Simple field (outputUnit)
	lengthInBits += 1

	// Simple field (installationMMIError)
	lengthInBits += 1

	// Simple field (EEWriteError)
	lengthInBits += 1

	// Simple field (EEChecksumError)
	lengthInBits += 1

	// Simple field (EEDataError)
	lengthInBits += 1

	// Simple field (microReset)
	lengthInBits += 1

	// Simple field (commsTxError)
	lengthInBits += 1

	// Simple field (internalStackOverflow)
	lengthInBits += 1

	// Simple field (microPowerReset)
	lengthInBits += 1

	return lengthInBits
}

func (m *_IdentifyReplyCommandExtendedDiagnosticSummary) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func IdentifyReplyCommandExtendedDiagnosticSummaryParse(readBuffer utils.ReadBuffer, attribute Attribute, numBytes uint8) (IdentifyReplyCommandExtendedDiagnosticSummary, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("IdentifyReplyCommandExtendedDiagnosticSummary"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for IdentifyReplyCommandExtendedDiagnosticSummary")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (lowApplication)
	if pullErr := readBuffer.PullContext("lowApplication"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for lowApplication")
	}
	_lowApplication, _lowApplicationErr := ApplicationIdContainerParse(readBuffer)
	if _lowApplicationErr != nil {
		return nil, errors.Wrap(_lowApplicationErr, "Error parsing 'lowApplication' field of IdentifyReplyCommandExtendedDiagnosticSummary")
	}
	lowApplication := _lowApplication
	if closeErr := readBuffer.CloseContext("lowApplication"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for lowApplication")
	}

	// Simple Field (highApplication)
	if pullErr := readBuffer.PullContext("highApplication"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for highApplication")
	}
	_highApplication, _highApplicationErr := ApplicationIdContainerParse(readBuffer)
	if _highApplicationErr != nil {
		return nil, errors.Wrap(_highApplicationErr, "Error parsing 'highApplication' field of IdentifyReplyCommandExtendedDiagnosticSummary")
	}
	highApplication := _highApplication
	if closeErr := readBuffer.CloseContext("highApplication"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for highApplication")
	}

	// Simple Field (area)
	_area, _areaErr := readBuffer.ReadByte("area")
	if _areaErr != nil {
		return nil, errors.Wrap(_areaErr, "Error parsing 'area' field of IdentifyReplyCommandExtendedDiagnosticSummary")
	}
	area := _area

	// Simple Field (crc)
	_crc, _crcErr := readBuffer.ReadUint16("crc", 16)
	if _crcErr != nil {
		return nil, errors.Wrap(_crcErr, "Error parsing 'crc' field of IdentifyReplyCommandExtendedDiagnosticSummary")
	}
	crc := _crc

	// Simple Field (serialNumber)
	_serialNumber, _serialNumberErr := readBuffer.ReadUint32("serialNumber", 32)
	if _serialNumberErr != nil {
		return nil, errors.Wrap(_serialNumberErr, "Error parsing 'serialNumber' field of IdentifyReplyCommandExtendedDiagnosticSummary")
	}
	serialNumber := _serialNumber

	// Simple Field (networkVoltage)
	_networkVoltage, _networkVoltageErr := readBuffer.ReadByte("networkVoltage")
	if _networkVoltageErr != nil {
		return nil, errors.Wrap(_networkVoltageErr, "Error parsing 'networkVoltage' field of IdentifyReplyCommandExtendedDiagnosticSummary")
	}
	networkVoltage := _networkVoltage

	// Virtual field
	_networkVoltageInVolts := float32(networkVoltage) / float32(float32(6.375))
	networkVoltageInVolts := float32(_networkVoltageInVolts)
	_ = networkVoltageInVolts

	// Simple Field (unitInLearnMode)
	_unitInLearnMode, _unitInLearnModeErr := readBuffer.ReadBit("unitInLearnMode")
	if _unitInLearnModeErr != nil {
		return nil, errors.Wrap(_unitInLearnModeErr, "Error parsing 'unitInLearnMode' field of IdentifyReplyCommandExtendedDiagnosticSummary")
	}
	unitInLearnMode := _unitInLearnMode

	// Simple Field (networkVoltageLow)
	_networkVoltageLow, _networkVoltageLowErr := readBuffer.ReadBit("networkVoltageLow")
	if _networkVoltageLowErr != nil {
		return nil, errors.Wrap(_networkVoltageLowErr, "Error parsing 'networkVoltageLow' field of IdentifyReplyCommandExtendedDiagnosticSummary")
	}
	networkVoltageLow := _networkVoltageLow

	// Simple Field (networkVoltageMarginal)
	_networkVoltageMarginal, _networkVoltageMarginalErr := readBuffer.ReadBit("networkVoltageMarginal")
	if _networkVoltageMarginalErr != nil {
		return nil, errors.Wrap(_networkVoltageMarginalErr, "Error parsing 'networkVoltageMarginal' field of IdentifyReplyCommandExtendedDiagnosticSummary")
	}
	networkVoltageMarginal := _networkVoltageMarginal

	// Reserved Field (Compartmentalized so the "reserved" variable can't leak)
	{
		reserved, _err := readBuffer.ReadUint8("reserved", 1)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'reserved' field of IdentifyReplyCommandExtendedDiagnosticSummary")
		}
		if reserved != uint8(0) {
			log.Info().Fields(map[string]interface{}{
				"expected value": uint8(0),
				"got value":      reserved,
			}).Msg("Got unexpected response for reserved field.")
		}
	}

	// Reserved Field (Compartmentalized so the "reserved" variable can't leak)
	{
		reserved, _err := readBuffer.ReadUint8("reserved", 1)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'reserved' field of IdentifyReplyCommandExtendedDiagnosticSummary")
		}
		if reserved != uint8(0) {
			log.Info().Fields(map[string]interface{}{
				"expected value": uint8(0),
				"got value":      reserved,
			}).Msg("Got unexpected response for reserved field.")
		}
	}

	// Reserved Field (Compartmentalized so the "reserved" variable can't leak)
	{
		reserved, _err := readBuffer.ReadUint8("reserved", 1)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'reserved' field of IdentifyReplyCommandExtendedDiagnosticSummary")
		}
		if reserved != uint8(0) {
			log.Info().Fields(map[string]interface{}{
				"expected value": uint8(0),
				"got value":      reserved,
			}).Msg("Got unexpected response for reserved field.")
		}
	}

	// Simple Field (enableChecksumAlarm)
	_enableChecksumAlarm, _enableChecksumAlarmErr := readBuffer.ReadBit("enableChecksumAlarm")
	if _enableChecksumAlarmErr != nil {
		return nil, errors.Wrap(_enableChecksumAlarmErr, "Error parsing 'enableChecksumAlarm' field of IdentifyReplyCommandExtendedDiagnosticSummary")
	}
	enableChecksumAlarm := _enableChecksumAlarm

	// Simple Field (outputUnit)
	_outputUnit, _outputUnitErr := readBuffer.ReadBit("outputUnit")
	if _outputUnitErr != nil {
		return nil, errors.Wrap(_outputUnitErr, "Error parsing 'outputUnit' field of IdentifyReplyCommandExtendedDiagnosticSummary")
	}
	outputUnit := _outputUnit

	// Simple Field (installationMMIError)
	_installationMMIError, _installationMMIErrorErr := readBuffer.ReadBit("installationMMIError")
	if _installationMMIErrorErr != nil {
		return nil, errors.Wrap(_installationMMIErrorErr, "Error parsing 'installationMMIError' field of IdentifyReplyCommandExtendedDiagnosticSummary")
	}
	installationMMIError := _installationMMIError

	// Simple Field (EEWriteError)
	_EEWriteError, _EEWriteErrorErr := readBuffer.ReadBit("EEWriteError")
	if _EEWriteErrorErr != nil {
		return nil, errors.Wrap(_EEWriteErrorErr, "Error parsing 'EEWriteError' field of IdentifyReplyCommandExtendedDiagnosticSummary")
	}
	EEWriteError := _EEWriteError

	// Simple Field (EEChecksumError)
	_EEChecksumError, _EEChecksumErrorErr := readBuffer.ReadBit("EEChecksumError")
	if _EEChecksumErrorErr != nil {
		return nil, errors.Wrap(_EEChecksumErrorErr, "Error parsing 'EEChecksumError' field of IdentifyReplyCommandExtendedDiagnosticSummary")
	}
	EEChecksumError := _EEChecksumError

	// Simple Field (EEDataError)
	_EEDataError, _EEDataErrorErr := readBuffer.ReadBit("EEDataError")
	if _EEDataErrorErr != nil {
		return nil, errors.Wrap(_EEDataErrorErr, "Error parsing 'EEDataError' field of IdentifyReplyCommandExtendedDiagnosticSummary")
	}
	EEDataError := _EEDataError

	// Simple Field (microReset)
	_microReset, _microResetErr := readBuffer.ReadBit("microReset")
	if _microResetErr != nil {
		return nil, errors.Wrap(_microResetErr, "Error parsing 'microReset' field of IdentifyReplyCommandExtendedDiagnosticSummary")
	}
	microReset := _microReset

	// Simple Field (commsTxError)
	_commsTxError, _commsTxErrorErr := readBuffer.ReadBit("commsTxError")
	if _commsTxErrorErr != nil {
		return nil, errors.Wrap(_commsTxErrorErr, "Error parsing 'commsTxError' field of IdentifyReplyCommandExtendedDiagnosticSummary")
	}
	commsTxError := _commsTxError

	// Simple Field (internalStackOverflow)
	_internalStackOverflow, _internalStackOverflowErr := readBuffer.ReadBit("internalStackOverflow")
	if _internalStackOverflowErr != nil {
		return nil, errors.Wrap(_internalStackOverflowErr, "Error parsing 'internalStackOverflow' field of IdentifyReplyCommandExtendedDiagnosticSummary")
	}
	internalStackOverflow := _internalStackOverflow

	// Simple Field (microPowerReset)
	_microPowerReset, _microPowerResetErr := readBuffer.ReadBit("microPowerReset")
	if _microPowerResetErr != nil {
		return nil, errors.Wrap(_microPowerResetErr, "Error parsing 'microPowerReset' field of IdentifyReplyCommandExtendedDiagnosticSummary")
	}
	microPowerReset := _microPowerReset

	if closeErr := readBuffer.CloseContext("IdentifyReplyCommandExtendedDiagnosticSummary"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for IdentifyReplyCommandExtendedDiagnosticSummary")
	}

	// Create a partially initialized instance
	_child := &_IdentifyReplyCommandExtendedDiagnosticSummary{
		LowApplication:         lowApplication,
		HighApplication:        highApplication,
		Area:                   area,
		Crc:                    crc,
		SerialNumber:           serialNumber,
		NetworkVoltage:         networkVoltage,
		UnitInLearnMode:        unitInLearnMode,
		NetworkVoltageLow:      networkVoltageLow,
		NetworkVoltageMarginal: networkVoltageMarginal,
		EnableChecksumAlarm:    enableChecksumAlarm,
		OutputUnit:             outputUnit,
		InstallationMMIError:   installationMMIError,
		EEWriteError:           EEWriteError,
		EEChecksumError:        EEChecksumError,
		EEDataError:            EEDataError,
		MicroReset:             microReset,
		CommsTxError:           commsTxError,
		InternalStackOverflow:  internalStackOverflow,
		MicroPowerReset:        microPowerReset,
		_IdentifyReplyCommand: &_IdentifyReplyCommand{
			NumBytes: numBytes,
		},
	}
	_child._IdentifyReplyCommand._IdentifyReplyCommandChildRequirements = _child
	return _child, nil
}

func (m *_IdentifyReplyCommandExtendedDiagnosticSummary) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("IdentifyReplyCommandExtendedDiagnosticSummary"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for IdentifyReplyCommandExtendedDiagnosticSummary")
		}

		// Simple Field (lowApplication)
		if pushErr := writeBuffer.PushContext("lowApplication"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for lowApplication")
		}
		_lowApplicationErr := writeBuffer.WriteSerializable(m.GetLowApplication())
		if popErr := writeBuffer.PopContext("lowApplication"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for lowApplication")
		}
		if _lowApplicationErr != nil {
			return errors.Wrap(_lowApplicationErr, "Error serializing 'lowApplication' field")
		}

		// Simple Field (highApplication)
		if pushErr := writeBuffer.PushContext("highApplication"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for highApplication")
		}
		_highApplicationErr := writeBuffer.WriteSerializable(m.GetHighApplication())
		if popErr := writeBuffer.PopContext("highApplication"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for highApplication")
		}
		if _highApplicationErr != nil {
			return errors.Wrap(_highApplicationErr, "Error serializing 'highApplication' field")
		}

		// Simple Field (area)
		area := byte(m.GetArea())
		_areaErr := writeBuffer.WriteByte("area", (area))
		if _areaErr != nil {
			return errors.Wrap(_areaErr, "Error serializing 'area' field")
		}

		// Simple Field (crc)
		crc := uint16(m.GetCrc())
		_crcErr := writeBuffer.WriteUint16("crc", 16, (crc))
		if _crcErr != nil {
			return errors.Wrap(_crcErr, "Error serializing 'crc' field")
		}

		// Simple Field (serialNumber)
		serialNumber := uint32(m.GetSerialNumber())
		_serialNumberErr := writeBuffer.WriteUint32("serialNumber", 32, (serialNumber))
		if _serialNumberErr != nil {
			return errors.Wrap(_serialNumberErr, "Error serializing 'serialNumber' field")
		}

		// Simple Field (networkVoltage)
		networkVoltage := byte(m.GetNetworkVoltage())
		_networkVoltageErr := writeBuffer.WriteByte("networkVoltage", (networkVoltage))
		if _networkVoltageErr != nil {
			return errors.Wrap(_networkVoltageErr, "Error serializing 'networkVoltage' field")
		}
		// Virtual field
		if _networkVoltageInVoltsErr := writeBuffer.WriteVirtual("networkVoltageInVolts", m.GetNetworkVoltageInVolts()); _networkVoltageInVoltsErr != nil {
			return errors.Wrap(_networkVoltageInVoltsErr, "Error serializing 'networkVoltageInVolts' field")
		}

		// Simple Field (unitInLearnMode)
		unitInLearnMode := bool(m.GetUnitInLearnMode())
		_unitInLearnModeErr := writeBuffer.WriteBit("unitInLearnMode", (unitInLearnMode))
		if _unitInLearnModeErr != nil {
			return errors.Wrap(_unitInLearnModeErr, "Error serializing 'unitInLearnMode' field")
		}

		// Simple Field (networkVoltageLow)
		networkVoltageLow := bool(m.GetNetworkVoltageLow())
		_networkVoltageLowErr := writeBuffer.WriteBit("networkVoltageLow", (networkVoltageLow))
		if _networkVoltageLowErr != nil {
			return errors.Wrap(_networkVoltageLowErr, "Error serializing 'networkVoltageLow' field")
		}

		// Simple Field (networkVoltageMarginal)
		networkVoltageMarginal := bool(m.GetNetworkVoltageMarginal())
		_networkVoltageMarginalErr := writeBuffer.WriteBit("networkVoltageMarginal", (networkVoltageMarginal))
		if _networkVoltageMarginalErr != nil {
			return errors.Wrap(_networkVoltageMarginalErr, "Error serializing 'networkVoltageMarginal' field")
		}

		// Reserved Field (reserved)
		{
			_err := writeBuffer.WriteUint8("reserved", 1, uint8(0))
			if _err != nil {
				return errors.Wrap(_err, "Error serializing 'reserved' field")
			}
		}

		// Reserved Field (reserved)
		{
			_err := writeBuffer.WriteUint8("reserved", 1, uint8(0))
			if _err != nil {
				return errors.Wrap(_err, "Error serializing 'reserved' field")
			}
		}

		// Reserved Field (reserved)
		{
			_err := writeBuffer.WriteUint8("reserved", 1, uint8(0))
			if _err != nil {
				return errors.Wrap(_err, "Error serializing 'reserved' field")
			}
		}

		// Simple Field (enableChecksumAlarm)
		enableChecksumAlarm := bool(m.GetEnableChecksumAlarm())
		_enableChecksumAlarmErr := writeBuffer.WriteBit("enableChecksumAlarm", (enableChecksumAlarm))
		if _enableChecksumAlarmErr != nil {
			return errors.Wrap(_enableChecksumAlarmErr, "Error serializing 'enableChecksumAlarm' field")
		}

		// Simple Field (outputUnit)
		outputUnit := bool(m.GetOutputUnit())
		_outputUnitErr := writeBuffer.WriteBit("outputUnit", (outputUnit))
		if _outputUnitErr != nil {
			return errors.Wrap(_outputUnitErr, "Error serializing 'outputUnit' field")
		}

		// Simple Field (installationMMIError)
		installationMMIError := bool(m.GetInstallationMMIError())
		_installationMMIErrorErr := writeBuffer.WriteBit("installationMMIError", (installationMMIError))
		if _installationMMIErrorErr != nil {
			return errors.Wrap(_installationMMIErrorErr, "Error serializing 'installationMMIError' field")
		}

		// Simple Field (EEWriteError)
		EEWriteError := bool(m.GetEEWriteError())
		_EEWriteErrorErr := writeBuffer.WriteBit("EEWriteError", (EEWriteError))
		if _EEWriteErrorErr != nil {
			return errors.Wrap(_EEWriteErrorErr, "Error serializing 'EEWriteError' field")
		}

		// Simple Field (EEChecksumError)
		EEChecksumError := bool(m.GetEEChecksumError())
		_EEChecksumErrorErr := writeBuffer.WriteBit("EEChecksumError", (EEChecksumError))
		if _EEChecksumErrorErr != nil {
			return errors.Wrap(_EEChecksumErrorErr, "Error serializing 'EEChecksumError' field")
		}

		// Simple Field (EEDataError)
		EEDataError := bool(m.GetEEDataError())
		_EEDataErrorErr := writeBuffer.WriteBit("EEDataError", (EEDataError))
		if _EEDataErrorErr != nil {
			return errors.Wrap(_EEDataErrorErr, "Error serializing 'EEDataError' field")
		}

		// Simple Field (microReset)
		microReset := bool(m.GetMicroReset())
		_microResetErr := writeBuffer.WriteBit("microReset", (microReset))
		if _microResetErr != nil {
			return errors.Wrap(_microResetErr, "Error serializing 'microReset' field")
		}

		// Simple Field (commsTxError)
		commsTxError := bool(m.GetCommsTxError())
		_commsTxErrorErr := writeBuffer.WriteBit("commsTxError", (commsTxError))
		if _commsTxErrorErr != nil {
			return errors.Wrap(_commsTxErrorErr, "Error serializing 'commsTxError' field")
		}

		// Simple Field (internalStackOverflow)
		internalStackOverflow := bool(m.GetInternalStackOverflow())
		_internalStackOverflowErr := writeBuffer.WriteBit("internalStackOverflow", (internalStackOverflow))
		if _internalStackOverflowErr != nil {
			return errors.Wrap(_internalStackOverflowErr, "Error serializing 'internalStackOverflow' field")
		}

		// Simple Field (microPowerReset)
		microPowerReset := bool(m.GetMicroPowerReset())
		_microPowerResetErr := writeBuffer.WriteBit("microPowerReset", (microPowerReset))
		if _microPowerResetErr != nil {
			return errors.Wrap(_microPowerResetErr, "Error serializing 'microPowerReset' field")
		}

		if popErr := writeBuffer.PopContext("IdentifyReplyCommandExtendedDiagnosticSummary"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for IdentifyReplyCommandExtendedDiagnosticSummary")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_IdentifyReplyCommandExtendedDiagnosticSummary) isIdentifyReplyCommandExtendedDiagnosticSummary() bool {
	return true
}

func (m *_IdentifyReplyCommandExtendedDiagnosticSummary) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
