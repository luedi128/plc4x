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

// SubscribeCOVPropertyMultipleError is the corresponding interface of SubscribeCOVPropertyMultipleError
type SubscribeCOVPropertyMultipleError interface {
	utils.LengthAware
	utils.Serializable
	BACnetError
	// GetErrorType returns ErrorType (property field)
	GetErrorType() ErrorEnclosed
	// GetFirstFailedSubscription returns FirstFailedSubscription (property field)
	GetFirstFailedSubscription() SubscribeCOVPropertyMultipleErrorFirstFailedSubscription
}

// _SubscribeCOVPropertyMultipleError is the data-structure of this message
type _SubscribeCOVPropertyMultipleError struct {
	*_BACnetError
	ErrorType               ErrorEnclosed
	FirstFailedSubscription SubscribeCOVPropertyMultipleErrorFirstFailedSubscription
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_SubscribeCOVPropertyMultipleError) GetErrorChoice() BACnetConfirmedServiceChoice {
	return BACnetConfirmedServiceChoice_SUBSCRIBE_COV_PROPERTY_MULTIPLE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_SubscribeCOVPropertyMultipleError) InitializeParent(parent BACnetError) {}

func (m *_SubscribeCOVPropertyMultipleError) GetParent() BACnetError {
	return m._BACnetError
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_SubscribeCOVPropertyMultipleError) GetErrorType() ErrorEnclosed {
	return m.ErrorType
}

func (m *_SubscribeCOVPropertyMultipleError) GetFirstFailedSubscription() SubscribeCOVPropertyMultipleErrorFirstFailedSubscription {
	return m.FirstFailedSubscription
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewSubscribeCOVPropertyMultipleError factory function for _SubscribeCOVPropertyMultipleError
func NewSubscribeCOVPropertyMultipleError(errorType ErrorEnclosed, firstFailedSubscription SubscribeCOVPropertyMultipleErrorFirstFailedSubscription) *_SubscribeCOVPropertyMultipleError {
	_result := &_SubscribeCOVPropertyMultipleError{
		ErrorType:               errorType,
		FirstFailedSubscription: firstFailedSubscription,
		_BACnetError:            NewBACnetError(),
	}
	_result._BACnetError._BACnetErrorChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastSubscribeCOVPropertyMultipleError(structType interface{}) SubscribeCOVPropertyMultipleError {
	if casted, ok := structType.(SubscribeCOVPropertyMultipleError); ok {
		return casted
	}
	if casted, ok := structType.(*SubscribeCOVPropertyMultipleError); ok {
		return *casted
	}
	return nil
}

func (m *_SubscribeCOVPropertyMultipleError) GetTypeName() string {
	return "SubscribeCOVPropertyMultipleError"
}

func (m *_SubscribeCOVPropertyMultipleError) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_SubscribeCOVPropertyMultipleError) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (errorType)
	lengthInBits += m.ErrorType.GetLengthInBits()

	// Simple field (firstFailedSubscription)
	lengthInBits += m.FirstFailedSubscription.GetLengthInBits()

	return lengthInBits
}

func (m *_SubscribeCOVPropertyMultipleError) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func SubscribeCOVPropertyMultipleErrorParse(readBuffer utils.ReadBuffer, errorChoice BACnetConfirmedServiceChoice) (SubscribeCOVPropertyMultipleError, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("SubscribeCOVPropertyMultipleError"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for SubscribeCOVPropertyMultipleError")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (errorType)
	if pullErr := readBuffer.PullContext("errorType"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for errorType")
	}
	_errorType, _errorTypeErr := ErrorEnclosedParse(readBuffer, uint8(uint8(0)))
	if _errorTypeErr != nil {
		return nil, errors.Wrap(_errorTypeErr, "Error parsing 'errorType' field")
	}
	errorType := _errorType.(ErrorEnclosed)
	if closeErr := readBuffer.CloseContext("errorType"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for errorType")
	}

	// Simple Field (firstFailedSubscription)
	if pullErr := readBuffer.PullContext("firstFailedSubscription"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for firstFailedSubscription")
	}
	_firstFailedSubscription, _firstFailedSubscriptionErr := SubscribeCOVPropertyMultipleErrorFirstFailedSubscriptionParse(readBuffer, uint8(uint8(1)))
	if _firstFailedSubscriptionErr != nil {
		return nil, errors.Wrap(_firstFailedSubscriptionErr, "Error parsing 'firstFailedSubscription' field")
	}
	firstFailedSubscription := _firstFailedSubscription.(SubscribeCOVPropertyMultipleErrorFirstFailedSubscription)
	if closeErr := readBuffer.CloseContext("firstFailedSubscription"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for firstFailedSubscription")
	}

	if closeErr := readBuffer.CloseContext("SubscribeCOVPropertyMultipleError"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for SubscribeCOVPropertyMultipleError")
	}

	// Create a partially initialized instance
	_child := &_SubscribeCOVPropertyMultipleError{
		ErrorType:               errorType,
		FirstFailedSubscription: firstFailedSubscription,
		_BACnetError:            &_BACnetError{},
	}
	_child._BACnetError._BACnetErrorChildRequirements = _child
	return _child, nil
}

func (m *_SubscribeCOVPropertyMultipleError) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("SubscribeCOVPropertyMultipleError"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for SubscribeCOVPropertyMultipleError")
		}

		// Simple Field (errorType)
		if pushErr := writeBuffer.PushContext("errorType"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for errorType")
		}
		_errorTypeErr := writeBuffer.WriteSerializable(m.GetErrorType())
		if popErr := writeBuffer.PopContext("errorType"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for errorType")
		}
		if _errorTypeErr != nil {
			return errors.Wrap(_errorTypeErr, "Error serializing 'errorType' field")
		}

		// Simple Field (firstFailedSubscription)
		if pushErr := writeBuffer.PushContext("firstFailedSubscription"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for firstFailedSubscription")
		}
		_firstFailedSubscriptionErr := writeBuffer.WriteSerializable(m.GetFirstFailedSubscription())
		if popErr := writeBuffer.PopContext("firstFailedSubscription"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for firstFailedSubscription")
		}
		if _firstFailedSubscriptionErr != nil {
			return errors.Wrap(_firstFailedSubscriptionErr, "Error serializing 'firstFailedSubscription' field")
		}

		if popErr := writeBuffer.PopContext("SubscribeCOVPropertyMultipleError"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for SubscribeCOVPropertyMultipleError")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_SubscribeCOVPropertyMultipleError) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
