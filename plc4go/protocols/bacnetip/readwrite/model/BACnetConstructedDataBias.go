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

// BACnetConstructedDataBias is the corresponding interface of BACnetConstructedDataBias
type BACnetConstructedDataBias interface {
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetBias returns Bias (property field)
	GetBias() BACnetApplicationTagReal
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagReal
}

// _BACnetConstructedDataBias is the data-structure of this message
type _BACnetConstructedDataBias struct {
	*_BACnetConstructedData
	Bias BACnetApplicationTagReal

	// Arguments.
	TagNumber          uint8
	ArrayIndexArgument BACnetTagPayloadUnsignedInteger
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataBias) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataBias) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_BIAS
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataBias) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataBias) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataBias) GetBias() BACnetApplicationTagReal {
	return m.Bias
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataBias) GetActualValue() BACnetApplicationTagReal {
	return CastBACnetApplicationTagReal(m.GetBias())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataBias factory function for _BACnetConstructedDataBias
func NewBACnetConstructedDataBias(bias BACnetApplicationTagReal, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataBias {
	_result := &_BACnetConstructedDataBias{
		Bias:                   bias,
		_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataBias(structType interface{}) BACnetConstructedDataBias {
	if casted, ok := structType.(BACnetConstructedDataBias); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataBias); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataBias) GetTypeName() string {
	return "BACnetConstructedDataBias"
}

func (m *_BACnetConstructedDataBias) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetConstructedDataBias) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (bias)
	lengthInBits += m.Bias.GetLengthInBits()

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataBias) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataBiasParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataBias, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataBias"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataBias")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (bias)
	if pullErr := readBuffer.PullContext("bias"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for bias")
	}
	_bias, _biasErr := BACnetApplicationTagParse(readBuffer)
	if _biasErr != nil {
		return nil, errors.Wrap(_biasErr, "Error parsing 'bias' field")
	}
	bias := _bias.(BACnetApplicationTagReal)
	if closeErr := readBuffer.CloseContext("bias"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for bias")
	}

	// Virtual field
	_actualValue := bias
	actualValue := _actualValue
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataBias"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataBias")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataBias{
		Bias:                   bias,
		_BACnetConstructedData: &_BACnetConstructedData{},
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataBias) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataBias"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataBias")
		}

		// Simple Field (bias)
		if pushErr := writeBuffer.PushContext("bias"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for bias")
		}
		_biasErr := writeBuffer.WriteSerializable(m.GetBias())
		if popErr := writeBuffer.PopContext("bias"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for bias")
		}
		if _biasErr != nil {
			return errors.Wrap(_biasErr, "Error serializing 'bias' field")
		}
		// Virtual field
		if _actualValueErr := writeBuffer.WriteVirtual("actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataBias"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataBias")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataBias) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
