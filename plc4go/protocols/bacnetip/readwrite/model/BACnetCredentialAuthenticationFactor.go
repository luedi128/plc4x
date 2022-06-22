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

// BACnetCredentialAuthenticationFactor is the corresponding interface of BACnetCredentialAuthenticationFactor
type BACnetCredentialAuthenticationFactor interface {
	utils.LengthAware
	utils.Serializable
	// GetDisable returns Disable (property field)
	GetDisable() BACnetAccessAuthenticationFactorDisableTagged
	// GetAuthenticationFactor returns AuthenticationFactor (property field)
	GetAuthenticationFactor() BACnetAuthenticationFactorEnclosed
}

// _BACnetCredentialAuthenticationFactor is the data-structure of this message
type _BACnetCredentialAuthenticationFactor struct {
	Disable              BACnetAccessAuthenticationFactorDisableTagged
	AuthenticationFactor BACnetAuthenticationFactorEnclosed
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetCredentialAuthenticationFactor) GetDisable() BACnetAccessAuthenticationFactorDisableTagged {
	return m.Disable
}

func (m *_BACnetCredentialAuthenticationFactor) GetAuthenticationFactor() BACnetAuthenticationFactorEnclosed {
	return m.AuthenticationFactor
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetCredentialAuthenticationFactor factory function for _BACnetCredentialAuthenticationFactor
func NewBACnetCredentialAuthenticationFactor(disable BACnetAccessAuthenticationFactorDisableTagged, authenticationFactor BACnetAuthenticationFactorEnclosed) *_BACnetCredentialAuthenticationFactor {
	return &_BACnetCredentialAuthenticationFactor{Disable: disable, AuthenticationFactor: authenticationFactor}
}

// Deprecated: use the interface for direct cast
func CastBACnetCredentialAuthenticationFactor(structType interface{}) BACnetCredentialAuthenticationFactor {
	if casted, ok := structType.(BACnetCredentialAuthenticationFactor); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetCredentialAuthenticationFactor); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetCredentialAuthenticationFactor) GetTypeName() string {
	return "BACnetCredentialAuthenticationFactor"
}

func (m *_BACnetCredentialAuthenticationFactor) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetCredentialAuthenticationFactor) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(0)

	// Simple field (disable)
	lengthInBits += m.Disable.GetLengthInBits()

	// Simple field (authenticationFactor)
	lengthInBits += m.AuthenticationFactor.GetLengthInBits()

	return lengthInBits
}

func (m *_BACnetCredentialAuthenticationFactor) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetCredentialAuthenticationFactorParse(readBuffer utils.ReadBuffer) (BACnetCredentialAuthenticationFactor, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetCredentialAuthenticationFactor"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetCredentialAuthenticationFactor")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (disable)
	if pullErr := readBuffer.PullContext("disable"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for disable")
	}
	_disable, _disableErr := BACnetAccessAuthenticationFactorDisableTaggedParse(readBuffer, uint8(uint8(0)), TagClass(TagClass_CONTEXT_SPECIFIC_TAGS))
	if _disableErr != nil {
		return nil, errors.Wrap(_disableErr, "Error parsing 'disable' field")
	}
	disable := _disable.(BACnetAccessAuthenticationFactorDisableTagged)
	if closeErr := readBuffer.CloseContext("disable"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for disable")
	}

	// Simple Field (authenticationFactor)
	if pullErr := readBuffer.PullContext("authenticationFactor"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for authenticationFactor")
	}
	_authenticationFactor, _authenticationFactorErr := BACnetAuthenticationFactorEnclosedParse(readBuffer, uint8(uint8(1)))
	if _authenticationFactorErr != nil {
		return nil, errors.Wrap(_authenticationFactorErr, "Error parsing 'authenticationFactor' field")
	}
	authenticationFactor := _authenticationFactor.(BACnetAuthenticationFactorEnclosed)
	if closeErr := readBuffer.CloseContext("authenticationFactor"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for authenticationFactor")
	}

	if closeErr := readBuffer.CloseContext("BACnetCredentialAuthenticationFactor"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetCredentialAuthenticationFactor")
	}

	// Create the instance
	return NewBACnetCredentialAuthenticationFactor(disable, authenticationFactor), nil
}

func (m *_BACnetCredentialAuthenticationFactor) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("BACnetCredentialAuthenticationFactor"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetCredentialAuthenticationFactor")
	}

	// Simple Field (disable)
	if pushErr := writeBuffer.PushContext("disable"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for disable")
	}
	_disableErr := writeBuffer.WriteSerializable(m.GetDisable())
	if popErr := writeBuffer.PopContext("disable"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for disable")
	}
	if _disableErr != nil {
		return errors.Wrap(_disableErr, "Error serializing 'disable' field")
	}

	// Simple Field (authenticationFactor)
	if pushErr := writeBuffer.PushContext("authenticationFactor"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for authenticationFactor")
	}
	_authenticationFactorErr := writeBuffer.WriteSerializable(m.GetAuthenticationFactor())
	if popErr := writeBuffer.PopContext("authenticationFactor"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for authenticationFactor")
	}
	if _authenticationFactorErr != nil {
		return errors.Wrap(_authenticationFactorErr, "Error serializing 'authenticationFactor' field")
	}

	if popErr := writeBuffer.PopContext("BACnetCredentialAuthenticationFactor"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetCredentialAuthenticationFactor")
	}
	return nil
}

func (m *_BACnetCredentialAuthenticationFactor) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
