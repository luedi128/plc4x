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

// KnxGroupAddress is the corresponding interface of KnxGroupAddress
type KnxGroupAddress interface {
	utils.LengthAware
	utils.Serializable
	// GetNumLevels returns NumLevels (discriminator field)
	GetNumLevels() uint8
}

// _KnxGroupAddress is the data-structure of this message
type _KnxGroupAddress struct {
	_KnxGroupAddressChildRequirements
}

type _KnxGroupAddressChildRequirements interface {
	GetLengthInBits() uint16
	GetLengthInBitsConditional(lastItem bool) uint16
	GetNumLevels() uint8
	Serialize(writeBuffer utils.WriteBuffer) error
}

type KnxGroupAddressParent interface {
	SerializeParent(writeBuffer utils.WriteBuffer, child KnxGroupAddress, serializeChildFunction func() error) error
	GetTypeName() string
}

type KnxGroupAddressChild interface {
	Serialize(writeBuffer utils.WriteBuffer) error
	InitializeParent(parent KnxGroupAddress)
	GetParent() *KnxGroupAddress

	GetTypeName() string
	KnxGroupAddress
}

// NewKnxGroupAddress factory function for _KnxGroupAddress
func NewKnxGroupAddress() *_KnxGroupAddress {
	return &_KnxGroupAddress{}
}

// Deprecated: use the interface for direct cast
func CastKnxGroupAddress(structType interface{}) KnxGroupAddress {
	if casted, ok := structType.(KnxGroupAddress); ok {
		return casted
	}
	if casted, ok := structType.(*KnxGroupAddress); ok {
		return *casted
	}
	return nil
}

func (m *_KnxGroupAddress) GetTypeName() string {
	return "KnxGroupAddress"
}

func (m *_KnxGroupAddress) GetParentLengthInBits() uint16 {
	lengthInBits := uint16(0)

	return lengthInBits
}

func (m *_KnxGroupAddress) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func KnxGroupAddressParse(readBuffer utils.ReadBuffer, numLevels uint8) (KnxGroupAddress, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("KnxGroupAddress"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for KnxGroupAddress")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	type KnxGroupAddressChildSerializeRequirement interface {
		KnxGroupAddress
		InitializeParent(KnxGroupAddress)
		GetParent() KnxGroupAddress
	}
	var _childTemp interface{}
	var _child KnxGroupAddressChildSerializeRequirement
	var typeSwitchError error
	switch {
	case numLevels == uint8(1): // KnxGroupAddressFreeLevel
		_childTemp, typeSwitchError = KnxGroupAddressFreeLevelParse(readBuffer, numLevels)
	case numLevels == uint8(2): // KnxGroupAddress2Level
		_childTemp, typeSwitchError = KnxGroupAddress2LevelParse(readBuffer, numLevels)
	case numLevels == uint8(3): // KnxGroupAddress3Level
		_childTemp, typeSwitchError = KnxGroupAddress3LevelParse(readBuffer, numLevels)
	default:
		// TODO: return actual type
		typeSwitchError = errors.New("Unmapped type")
	}
	if typeSwitchError != nil {
		return nil, errors.Wrap(typeSwitchError, "Error parsing sub-type for type-switch.")
	}
	_child = _childTemp.(KnxGroupAddressChildSerializeRequirement)

	if closeErr := readBuffer.CloseContext("KnxGroupAddress"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for KnxGroupAddress")
	}

	// Finish initializing
	_child.InitializeParent(_child)
	return _child, nil
}

func (pm *_KnxGroupAddress) SerializeParent(writeBuffer utils.WriteBuffer, child KnxGroupAddress, serializeChildFunction func() error) error {
	// We redirect all calls through client as some methods are only implemented there
	m := child
	_ = m
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("KnxGroupAddress"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for KnxGroupAddress")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	if _typeSwitchErr := serializeChildFunction(); _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	if popErr := writeBuffer.PopContext("KnxGroupAddress"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for KnxGroupAddress")
	}
	return nil
}

func (m *_KnxGroupAddress) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
