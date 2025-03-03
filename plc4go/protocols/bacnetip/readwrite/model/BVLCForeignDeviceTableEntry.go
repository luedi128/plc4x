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
	"context"
	"fmt"
	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
)

// Code generated by code-generation. DO NOT EDIT.

// BVLCForeignDeviceTableEntry is the corresponding interface of BVLCForeignDeviceTableEntry
type BVLCForeignDeviceTableEntry interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetIp returns Ip (property field)
	GetIp() []uint8
	// GetPort returns Port (property field)
	GetPort() uint16
	// GetTtl returns Ttl (property field)
	GetTtl() uint16
	// GetSecondRemainingBeforePurge returns SecondRemainingBeforePurge (property field)
	GetSecondRemainingBeforePurge() uint16
}

// BVLCForeignDeviceTableEntryExactly can be used when we want exactly this type and not a type which fulfills BVLCForeignDeviceTableEntry.
// This is useful for switch cases.
type BVLCForeignDeviceTableEntryExactly interface {
	BVLCForeignDeviceTableEntry
	isBVLCForeignDeviceTableEntry() bool
}

// _BVLCForeignDeviceTableEntry is the data-structure of this message
type _BVLCForeignDeviceTableEntry struct {
	Ip                         []uint8
	Port                       uint16
	Ttl                        uint16
	SecondRemainingBeforePurge uint16
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BVLCForeignDeviceTableEntry) GetIp() []uint8 {
	return m.Ip
}

func (m *_BVLCForeignDeviceTableEntry) GetPort() uint16 {
	return m.Port
}

func (m *_BVLCForeignDeviceTableEntry) GetTtl() uint16 {
	return m.Ttl
}

func (m *_BVLCForeignDeviceTableEntry) GetSecondRemainingBeforePurge() uint16 {
	return m.SecondRemainingBeforePurge
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBVLCForeignDeviceTableEntry factory function for _BVLCForeignDeviceTableEntry
func NewBVLCForeignDeviceTableEntry(ip []uint8, port uint16, ttl uint16, secondRemainingBeforePurge uint16) *_BVLCForeignDeviceTableEntry {
	return &_BVLCForeignDeviceTableEntry{Ip: ip, Port: port, Ttl: ttl, SecondRemainingBeforePurge: secondRemainingBeforePurge}
}

// Deprecated: use the interface for direct cast
func CastBVLCForeignDeviceTableEntry(structType any) BVLCForeignDeviceTableEntry {
	if casted, ok := structType.(BVLCForeignDeviceTableEntry); ok {
		return casted
	}
	if casted, ok := structType.(*BVLCForeignDeviceTableEntry); ok {
		return *casted
	}
	return nil
}

func (m *_BVLCForeignDeviceTableEntry) GetTypeName() string {
	return "BVLCForeignDeviceTableEntry"
}

func (m *_BVLCForeignDeviceTableEntry) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Array field
	if len(m.Ip) > 0 {
		lengthInBits += 8 * uint16(len(m.Ip))
	}

	// Simple field (port)
	lengthInBits += 16

	// Simple field (ttl)
	lengthInBits += 16

	// Simple field (secondRemainingBeforePurge)
	lengthInBits += 16

	return lengthInBits
}

func (m *_BVLCForeignDeviceTableEntry) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BVLCForeignDeviceTableEntryParse(ctx context.Context, theBytes []byte) (BVLCForeignDeviceTableEntry, error) {
	return BVLCForeignDeviceTableEntryParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func BVLCForeignDeviceTableEntryParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (BVLCForeignDeviceTableEntry, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("BVLCForeignDeviceTableEntry"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BVLCForeignDeviceTableEntry")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Array field (ip)
	if pullErr := readBuffer.PullContext("ip", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ip")
	}
	// Count array
	ip := make([]uint8, max(uint16(4), 0))
	// This happens when the size is set conditional to 0
	if len(ip) == 0 {
		ip = nil
	}
	{
		_numItems := uint16(max(uint16(4), 0))
		for _curItem := uint16(0); _curItem < _numItems; _curItem++ {
			arrayCtx := utils.CreateArrayContext(ctx, int(_numItems), int(_curItem))
			_ = arrayCtx
			_ = _curItem
			_item, _err := readBuffer.ReadUint8("", 8)
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'ip' field of BVLCForeignDeviceTableEntry")
			}
			ip[_curItem] = _item
		}
	}
	if closeErr := readBuffer.CloseContext("ip", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ip")
	}

	// Simple Field (port)
	_port, _portErr := readBuffer.ReadUint16("port", 16)
	if _portErr != nil {
		return nil, errors.Wrap(_portErr, "Error parsing 'port' field of BVLCForeignDeviceTableEntry")
	}
	port := _port

	// Simple Field (ttl)
	_ttl, _ttlErr := readBuffer.ReadUint16("ttl", 16)
	if _ttlErr != nil {
		return nil, errors.Wrap(_ttlErr, "Error parsing 'ttl' field of BVLCForeignDeviceTableEntry")
	}
	ttl := _ttl

	// Simple Field (secondRemainingBeforePurge)
	_secondRemainingBeforePurge, _secondRemainingBeforePurgeErr := readBuffer.ReadUint16("secondRemainingBeforePurge", 16)
	if _secondRemainingBeforePurgeErr != nil {
		return nil, errors.Wrap(_secondRemainingBeforePurgeErr, "Error parsing 'secondRemainingBeforePurge' field of BVLCForeignDeviceTableEntry")
	}
	secondRemainingBeforePurge := _secondRemainingBeforePurge

	if closeErr := readBuffer.CloseContext("BVLCForeignDeviceTableEntry"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BVLCForeignDeviceTableEntry")
	}

	// Create the instance
	return &_BVLCForeignDeviceTableEntry{
		Ip:                         ip,
		Port:                       port,
		Ttl:                        ttl,
		SecondRemainingBeforePurge: secondRemainingBeforePurge,
	}, nil
}

func (m *_BVLCForeignDeviceTableEntry) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BVLCForeignDeviceTableEntry) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("BVLCForeignDeviceTableEntry"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BVLCForeignDeviceTableEntry")
	}

	// Array Field (ip)
	if pushErr := writeBuffer.PushContext("ip", utils.WithRenderAsList(true)); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for ip")
	}
	for _curItem, _element := range m.GetIp() {
		_ = _curItem
		_elementErr := writeBuffer.WriteUint8("", 8, uint8(_element))
		if _elementErr != nil {
			return errors.Wrap(_elementErr, "Error serializing 'ip' field")
		}
	}
	if popErr := writeBuffer.PopContext("ip", utils.WithRenderAsList(true)); popErr != nil {
		return errors.Wrap(popErr, "Error popping for ip")
	}

	// Simple Field (port)
	port := uint16(m.GetPort())
	_portErr := writeBuffer.WriteUint16("port", 16, uint16((port)))
	if _portErr != nil {
		return errors.Wrap(_portErr, "Error serializing 'port' field")
	}

	// Simple Field (ttl)
	ttl := uint16(m.GetTtl())
	_ttlErr := writeBuffer.WriteUint16("ttl", 16, uint16((ttl)))
	if _ttlErr != nil {
		return errors.Wrap(_ttlErr, "Error serializing 'ttl' field")
	}

	// Simple Field (secondRemainingBeforePurge)
	secondRemainingBeforePurge := uint16(m.GetSecondRemainingBeforePurge())
	_secondRemainingBeforePurgeErr := writeBuffer.WriteUint16("secondRemainingBeforePurge", 16, uint16((secondRemainingBeforePurge)))
	if _secondRemainingBeforePurgeErr != nil {
		return errors.Wrap(_secondRemainingBeforePurgeErr, "Error serializing 'secondRemainingBeforePurge' field")
	}

	if popErr := writeBuffer.PopContext("BVLCForeignDeviceTableEntry"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BVLCForeignDeviceTableEntry")
	}
	return nil
}

func (m *_BVLCForeignDeviceTableEntry) isBVLCForeignDeviceTableEntry() bool {
	return true
}

func (m *_BVLCForeignDeviceTableEntry) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
