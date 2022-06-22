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

// APDUUnconfirmedRequest is the corresponding interface of APDUUnconfirmedRequest
type APDUUnconfirmedRequest interface {
	utils.LengthAware
	utils.Serializable
	APDU
	// GetServiceRequest returns ServiceRequest (property field)
	GetServiceRequest() BACnetUnconfirmedServiceRequest
}

// _APDUUnconfirmedRequest is the data-structure of this message
type _APDUUnconfirmedRequest struct {
	*_APDU
	ServiceRequest BACnetUnconfirmedServiceRequest

	// Arguments.
	ApduLength uint16
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_APDUUnconfirmedRequest) GetApduType() ApduType {
	return ApduType_UNCONFIRMED_REQUEST_PDU
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_APDUUnconfirmedRequest) InitializeParent(parent APDU) {}

func (m *_APDUUnconfirmedRequest) GetParent() APDU {
	return m._APDU
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_APDUUnconfirmedRequest) GetServiceRequest() BACnetUnconfirmedServiceRequest {
	return m.ServiceRequest
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewAPDUUnconfirmedRequest factory function for _APDUUnconfirmedRequest
func NewAPDUUnconfirmedRequest(serviceRequest BACnetUnconfirmedServiceRequest, apduLength uint16) *_APDUUnconfirmedRequest {
	_result := &_APDUUnconfirmedRequest{
		ServiceRequest: serviceRequest,
		_APDU:          NewAPDU(apduLength),
	}
	_result._APDU._APDUChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastAPDUUnconfirmedRequest(structType interface{}) APDUUnconfirmedRequest {
	if casted, ok := structType.(APDUUnconfirmedRequest); ok {
		return casted
	}
	if casted, ok := structType.(*APDUUnconfirmedRequest); ok {
		return *casted
	}
	return nil
}

func (m *_APDUUnconfirmedRequest) GetTypeName() string {
	return "APDUUnconfirmedRequest"
}

func (m *_APDUUnconfirmedRequest) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_APDUUnconfirmedRequest) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Reserved Field (reserved)
	lengthInBits += 4

	// Simple field (serviceRequest)
	lengthInBits += m.ServiceRequest.GetLengthInBits()

	return lengthInBits
}

func (m *_APDUUnconfirmedRequest) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func APDUUnconfirmedRequestParse(readBuffer utils.ReadBuffer, apduLength uint16) (APDUUnconfirmedRequest, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("APDUUnconfirmedRequest"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for APDUUnconfirmedRequest")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Reserved Field (Compartmentalized so the "reserved" variable can't leak)
	{
		reserved, _err := readBuffer.ReadUint8("reserved", 4)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'reserved' field")
		}
		if reserved != uint8(0) {
			log.Info().Fields(map[string]interface{}{
				"expected value": uint8(0),
				"got value":      reserved,
			}).Msg("Got unexpected response.")
		}
	}

	// Simple Field (serviceRequest)
	if pullErr := readBuffer.PullContext("serviceRequest"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for serviceRequest")
	}
	_serviceRequest, _serviceRequestErr := BACnetUnconfirmedServiceRequestParse(readBuffer, uint16(uint16(apduLength)-uint16(uint16(1))))
	if _serviceRequestErr != nil {
		return nil, errors.Wrap(_serviceRequestErr, "Error parsing 'serviceRequest' field")
	}
	serviceRequest := _serviceRequest.(BACnetUnconfirmedServiceRequest)
	if closeErr := readBuffer.CloseContext("serviceRequest"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for serviceRequest")
	}

	if closeErr := readBuffer.CloseContext("APDUUnconfirmedRequest"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for APDUUnconfirmedRequest")
	}

	// Create a partially initialized instance
	_child := &_APDUUnconfirmedRequest{
		ServiceRequest: serviceRequest,
		_APDU:          &_APDU{},
	}
	_child._APDU._APDUChildRequirements = _child
	return _child, nil
}

func (m *_APDUUnconfirmedRequest) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("APDUUnconfirmedRequest"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for APDUUnconfirmedRequest")
		}

		// Reserved Field (reserved)
		{
			_err := writeBuffer.WriteUint8("reserved", 4, uint8(0))
			if _err != nil {
				return errors.Wrap(_err, "Error serializing 'reserved' field")
			}
		}

		// Simple Field (serviceRequest)
		if pushErr := writeBuffer.PushContext("serviceRequest"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for serviceRequest")
		}
		_serviceRequestErr := writeBuffer.WriteSerializable(m.GetServiceRequest())
		if popErr := writeBuffer.PopContext("serviceRequest"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for serviceRequest")
		}
		if _serviceRequestErr != nil {
			return errors.Wrap(_serviceRequestErr, "Error serializing 'serviceRequest' field")
		}

		if popErr := writeBuffer.PopContext("APDUUnconfirmedRequest"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for APDUUnconfirmedRequest")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_APDUUnconfirmedRequest) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
