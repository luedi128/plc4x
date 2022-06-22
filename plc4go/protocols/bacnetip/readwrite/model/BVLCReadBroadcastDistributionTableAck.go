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

// BVLCReadBroadcastDistributionTableAck is the corresponding interface of BVLCReadBroadcastDistributionTableAck
type BVLCReadBroadcastDistributionTableAck interface {
	utils.LengthAware
	utils.Serializable
	BVLC
	// GetTable returns Table (property field)
	GetTable() []BVLCBroadcastDistributionTableEntry
}

// _BVLCReadBroadcastDistributionTableAck is the data-structure of this message
type _BVLCReadBroadcastDistributionTableAck struct {
	*_BVLC
	Table []BVLCBroadcastDistributionTableEntry

	// Arguments.
	BvlcPayloadLength uint16
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BVLCReadBroadcastDistributionTableAck) GetBvlcFunction() uint8 {
	return 0x03
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BVLCReadBroadcastDistributionTableAck) InitializeParent(parent BVLC) {}

func (m *_BVLCReadBroadcastDistributionTableAck) GetParent() BVLC {
	return m._BVLC
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BVLCReadBroadcastDistributionTableAck) GetTable() []BVLCBroadcastDistributionTableEntry {
	return m.Table
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBVLCReadBroadcastDistributionTableAck factory function for _BVLCReadBroadcastDistributionTableAck
func NewBVLCReadBroadcastDistributionTableAck(table []BVLCBroadcastDistributionTableEntry, bvlcPayloadLength uint16) *_BVLCReadBroadcastDistributionTableAck {
	_result := &_BVLCReadBroadcastDistributionTableAck{
		Table: table,
		_BVLC: NewBVLC(),
	}
	_result._BVLC._BVLCChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBVLCReadBroadcastDistributionTableAck(structType interface{}) BVLCReadBroadcastDistributionTableAck {
	if casted, ok := structType.(BVLCReadBroadcastDistributionTableAck); ok {
		return casted
	}
	if casted, ok := structType.(*BVLCReadBroadcastDistributionTableAck); ok {
		return *casted
	}
	return nil
}

func (m *_BVLCReadBroadcastDistributionTableAck) GetTypeName() string {
	return "BVLCReadBroadcastDistributionTableAck"
}

func (m *_BVLCReadBroadcastDistributionTableAck) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BVLCReadBroadcastDistributionTableAck) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Array field
	if len(m.Table) > 0 {
		for _, element := range m.Table {
			lengthInBits += element.GetLengthInBits()
		}
	}

	return lengthInBits
}

func (m *_BVLCReadBroadcastDistributionTableAck) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BVLCReadBroadcastDistributionTableAckParse(readBuffer utils.ReadBuffer, bvlcPayloadLength uint16) (BVLCReadBroadcastDistributionTableAck, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BVLCReadBroadcastDistributionTableAck"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BVLCReadBroadcastDistributionTableAck")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Array field (table)
	if pullErr := readBuffer.PullContext("table", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for table")
	}
	// Length array
	table := make([]BVLCBroadcastDistributionTableEntry, 0)
	{
		_tableLength := bvlcPayloadLength
		_tableEndPos := positionAware.GetPos() + uint16(_tableLength)
		for positionAware.GetPos() < _tableEndPos {
			_item, _err := BVLCBroadcastDistributionTableEntryParse(readBuffer)
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'table' field")
			}
			table = append(table, _item.(BVLCBroadcastDistributionTableEntry))
		}
	}
	if closeErr := readBuffer.CloseContext("table", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for table")
	}

	if closeErr := readBuffer.CloseContext("BVLCReadBroadcastDistributionTableAck"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BVLCReadBroadcastDistributionTableAck")
	}

	// Create a partially initialized instance
	_child := &_BVLCReadBroadcastDistributionTableAck{
		Table: table,
		_BVLC: &_BVLC{},
	}
	_child._BVLC._BVLCChildRequirements = _child
	return _child, nil
}

func (m *_BVLCReadBroadcastDistributionTableAck) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BVLCReadBroadcastDistributionTableAck"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BVLCReadBroadcastDistributionTableAck")
		}

		// Array Field (table)
		if m.GetTable() != nil {
			if pushErr := writeBuffer.PushContext("table", utils.WithRenderAsList(true)); pushErr != nil {
				return errors.Wrap(pushErr, "Error pushing for table")
			}
			for _, _element := range m.GetTable() {
				_elementErr := writeBuffer.WriteSerializable(_element)
				if _elementErr != nil {
					return errors.Wrap(_elementErr, "Error serializing 'table' field")
				}
			}
			if popErr := writeBuffer.PopContext("table", utils.WithRenderAsList(true)); popErr != nil {
				return errors.Wrap(popErr, "Error popping for table")
			}
		}

		if popErr := writeBuffer.PopContext("BVLCReadBroadcastDistributionTableAck"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BVLCReadBroadcastDistributionTableAck")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BVLCReadBroadcastDistributionTableAck) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
