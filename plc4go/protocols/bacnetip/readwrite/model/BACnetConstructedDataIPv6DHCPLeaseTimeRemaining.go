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

// BACnetConstructedDataIPv6DHCPLeaseTimeRemaining is the corresponding interface of BACnetConstructedDataIPv6DHCPLeaseTimeRemaining
type BACnetConstructedDataIPv6DHCPLeaseTimeRemaining interface {
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetIpv6DhcpLeaseTimeRemaining returns Ipv6DhcpLeaseTimeRemaining (property field)
	GetIpv6DhcpLeaseTimeRemaining() BACnetApplicationTagUnsignedInteger
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagUnsignedInteger
}

// _BACnetConstructedDataIPv6DHCPLeaseTimeRemaining is the data-structure of this message
type _BACnetConstructedDataIPv6DHCPLeaseTimeRemaining struct {
	*_BACnetConstructedData
	Ipv6DhcpLeaseTimeRemaining BACnetApplicationTagUnsignedInteger

	// Arguments.
	TagNumber          uint8
	ArrayIndexArgument BACnetTagPayloadUnsignedInteger
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataIPv6DHCPLeaseTimeRemaining) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataIPv6DHCPLeaseTimeRemaining) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_IPV6_DHCP_LEASE_TIME_REMAINING
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataIPv6DHCPLeaseTimeRemaining) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataIPv6DHCPLeaseTimeRemaining) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataIPv6DHCPLeaseTimeRemaining) GetIpv6DhcpLeaseTimeRemaining() BACnetApplicationTagUnsignedInteger {
	return m.Ipv6DhcpLeaseTimeRemaining
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataIPv6DHCPLeaseTimeRemaining) GetActualValue() BACnetApplicationTagUnsignedInteger {
	return CastBACnetApplicationTagUnsignedInteger(m.GetIpv6DhcpLeaseTimeRemaining())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataIPv6DHCPLeaseTimeRemaining factory function for _BACnetConstructedDataIPv6DHCPLeaseTimeRemaining
func NewBACnetConstructedDataIPv6DHCPLeaseTimeRemaining(ipv6DhcpLeaseTimeRemaining BACnetApplicationTagUnsignedInteger, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataIPv6DHCPLeaseTimeRemaining {
	_result := &_BACnetConstructedDataIPv6DHCPLeaseTimeRemaining{
		Ipv6DhcpLeaseTimeRemaining: ipv6DhcpLeaseTimeRemaining,
		_BACnetConstructedData:     NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataIPv6DHCPLeaseTimeRemaining(structType interface{}) BACnetConstructedDataIPv6DHCPLeaseTimeRemaining {
	if casted, ok := structType.(BACnetConstructedDataIPv6DHCPLeaseTimeRemaining); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataIPv6DHCPLeaseTimeRemaining); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataIPv6DHCPLeaseTimeRemaining) GetTypeName() string {
	return "BACnetConstructedDataIPv6DHCPLeaseTimeRemaining"
}

func (m *_BACnetConstructedDataIPv6DHCPLeaseTimeRemaining) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetConstructedDataIPv6DHCPLeaseTimeRemaining) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (ipv6DhcpLeaseTimeRemaining)
	lengthInBits += m.Ipv6DhcpLeaseTimeRemaining.GetLengthInBits()

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataIPv6DHCPLeaseTimeRemaining) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataIPv6DHCPLeaseTimeRemainingParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataIPv6DHCPLeaseTimeRemaining, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataIPv6DHCPLeaseTimeRemaining"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataIPv6DHCPLeaseTimeRemaining")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (ipv6DhcpLeaseTimeRemaining)
	if pullErr := readBuffer.PullContext("ipv6DhcpLeaseTimeRemaining"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ipv6DhcpLeaseTimeRemaining")
	}
	_ipv6DhcpLeaseTimeRemaining, _ipv6DhcpLeaseTimeRemainingErr := BACnetApplicationTagParse(readBuffer)
	if _ipv6DhcpLeaseTimeRemainingErr != nil {
		return nil, errors.Wrap(_ipv6DhcpLeaseTimeRemainingErr, "Error parsing 'ipv6DhcpLeaseTimeRemaining' field")
	}
	ipv6DhcpLeaseTimeRemaining := _ipv6DhcpLeaseTimeRemaining.(BACnetApplicationTagUnsignedInteger)
	if closeErr := readBuffer.CloseContext("ipv6DhcpLeaseTimeRemaining"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ipv6DhcpLeaseTimeRemaining")
	}

	// Virtual field
	_actualValue := ipv6DhcpLeaseTimeRemaining
	actualValue := _actualValue
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataIPv6DHCPLeaseTimeRemaining"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataIPv6DHCPLeaseTimeRemaining")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataIPv6DHCPLeaseTimeRemaining{
		Ipv6DhcpLeaseTimeRemaining: ipv6DhcpLeaseTimeRemaining,
		_BACnetConstructedData:     &_BACnetConstructedData{},
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataIPv6DHCPLeaseTimeRemaining) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataIPv6DHCPLeaseTimeRemaining"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataIPv6DHCPLeaseTimeRemaining")
		}

		// Simple Field (ipv6DhcpLeaseTimeRemaining)
		if pushErr := writeBuffer.PushContext("ipv6DhcpLeaseTimeRemaining"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ipv6DhcpLeaseTimeRemaining")
		}
		_ipv6DhcpLeaseTimeRemainingErr := writeBuffer.WriteSerializable(m.GetIpv6DhcpLeaseTimeRemaining())
		if popErr := writeBuffer.PopContext("ipv6DhcpLeaseTimeRemaining"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ipv6DhcpLeaseTimeRemaining")
		}
		if _ipv6DhcpLeaseTimeRemainingErr != nil {
			return errors.Wrap(_ipv6DhcpLeaseTimeRemainingErr, "Error serializing 'ipv6DhcpLeaseTimeRemaining' field")
		}
		// Virtual field
		if _actualValueErr := writeBuffer.WriteVirtual("actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataIPv6DHCPLeaseTimeRemaining"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataIPv6DHCPLeaseTimeRemaining")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataIPv6DHCPLeaseTimeRemaining) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
