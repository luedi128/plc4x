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

// AdsInvalidResponse is the corresponding interface of AdsInvalidResponse
type AdsInvalidResponse interface {
	utils.LengthAware
	utils.Serializable
	AdsData
}

// _AdsInvalidResponse is the data-structure of this message
type _AdsInvalidResponse struct {
	*_AdsData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_AdsInvalidResponse) GetCommandId() CommandId {
	return CommandId_INVALID
}

func (m *_AdsInvalidResponse) GetResponse() bool {
	return bool(true)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_AdsInvalidResponse) InitializeParent(parent AdsData) {}

func (m *_AdsInvalidResponse) GetParent() AdsData {
	return m._AdsData
}

// NewAdsInvalidResponse factory function for _AdsInvalidResponse
func NewAdsInvalidResponse() *_AdsInvalidResponse {
	_result := &_AdsInvalidResponse{
		_AdsData: NewAdsData(),
	}
	_result._AdsData._AdsDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastAdsInvalidResponse(structType interface{}) AdsInvalidResponse {
	if casted, ok := structType.(AdsInvalidResponse); ok {
		return casted
	}
	if casted, ok := structType.(*AdsInvalidResponse); ok {
		return *casted
	}
	return nil
}

func (m *_AdsInvalidResponse) GetTypeName() string {
	return "AdsInvalidResponse"
}

func (m *_AdsInvalidResponse) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_AdsInvalidResponse) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	return lengthInBits
}

func (m *_AdsInvalidResponse) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func AdsInvalidResponseParse(readBuffer utils.ReadBuffer, commandId CommandId, response bool) (AdsInvalidResponse, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("AdsInvalidResponse"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for AdsInvalidResponse")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("AdsInvalidResponse"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for AdsInvalidResponse")
	}

	// Create a partially initialized instance
	_child := &_AdsInvalidResponse{
		_AdsData: &_AdsData{},
	}
	_child._AdsData._AdsDataChildRequirements = _child
	return _child, nil
}

func (m *_AdsInvalidResponse) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("AdsInvalidResponse"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for AdsInvalidResponse")
		}

		if popErr := writeBuffer.PopContext("AdsInvalidResponse"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for AdsInvalidResponse")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_AdsInvalidResponse) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
