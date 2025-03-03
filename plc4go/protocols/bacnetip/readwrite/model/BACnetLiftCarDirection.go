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

// BACnetLiftCarDirection is an enum
type BACnetLiftCarDirection uint16

type IBACnetLiftCarDirection interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

const (
	BACnetLiftCarDirection_UNKNOWN                  BACnetLiftCarDirection = 0
	BACnetLiftCarDirection_NONE                     BACnetLiftCarDirection = 1
	BACnetLiftCarDirection_STOPPED                  BACnetLiftCarDirection = 2
	BACnetLiftCarDirection_UP                       BACnetLiftCarDirection = 3
	BACnetLiftCarDirection_DOWN                     BACnetLiftCarDirection = 4
	BACnetLiftCarDirection_UP_AND_DOWN              BACnetLiftCarDirection = 5
	BACnetLiftCarDirection_VENDOR_PROPRIETARY_VALUE BACnetLiftCarDirection = 0xFFFF
)

var BACnetLiftCarDirectionValues []BACnetLiftCarDirection

func init() {
	_ = errors.New
	BACnetLiftCarDirectionValues = []BACnetLiftCarDirection{
		BACnetLiftCarDirection_UNKNOWN,
		BACnetLiftCarDirection_NONE,
		BACnetLiftCarDirection_STOPPED,
		BACnetLiftCarDirection_UP,
		BACnetLiftCarDirection_DOWN,
		BACnetLiftCarDirection_UP_AND_DOWN,
		BACnetLiftCarDirection_VENDOR_PROPRIETARY_VALUE,
	}
}

func BACnetLiftCarDirectionByValue(value uint16) (enum BACnetLiftCarDirection, ok bool) {
	switch value {
	case 0:
		return BACnetLiftCarDirection_UNKNOWN, true
	case 0xFFFF:
		return BACnetLiftCarDirection_VENDOR_PROPRIETARY_VALUE, true
	case 1:
		return BACnetLiftCarDirection_NONE, true
	case 2:
		return BACnetLiftCarDirection_STOPPED, true
	case 3:
		return BACnetLiftCarDirection_UP, true
	case 4:
		return BACnetLiftCarDirection_DOWN, true
	case 5:
		return BACnetLiftCarDirection_UP_AND_DOWN, true
	}
	return 0, false
}

func BACnetLiftCarDirectionByName(value string) (enum BACnetLiftCarDirection, ok bool) {
	switch value {
	case "UNKNOWN":
		return BACnetLiftCarDirection_UNKNOWN, true
	case "VENDOR_PROPRIETARY_VALUE":
		return BACnetLiftCarDirection_VENDOR_PROPRIETARY_VALUE, true
	case "NONE":
		return BACnetLiftCarDirection_NONE, true
	case "STOPPED":
		return BACnetLiftCarDirection_STOPPED, true
	case "UP":
		return BACnetLiftCarDirection_UP, true
	case "DOWN":
		return BACnetLiftCarDirection_DOWN, true
	case "UP_AND_DOWN":
		return BACnetLiftCarDirection_UP_AND_DOWN, true
	}
	return 0, false
}

func BACnetLiftCarDirectionKnows(value uint16) bool {
	for _, typeValue := range BACnetLiftCarDirectionValues {
		if uint16(typeValue) == value {
			return true
		}
	}
	return false
}

func CastBACnetLiftCarDirection(structType any) BACnetLiftCarDirection {
	castFunc := func(typ any) BACnetLiftCarDirection {
		if sBACnetLiftCarDirection, ok := typ.(BACnetLiftCarDirection); ok {
			return sBACnetLiftCarDirection
		}
		return 0
	}
	return castFunc(structType)
}

func (m BACnetLiftCarDirection) GetLengthInBits(ctx context.Context) uint16 {
	return 16
}

func (m BACnetLiftCarDirection) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetLiftCarDirectionParse(ctx context.Context, theBytes []byte) (BACnetLiftCarDirection, error) {
	return BACnetLiftCarDirectionParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func BACnetLiftCarDirectionParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetLiftCarDirection, error) {
	log := zerolog.Ctx(ctx)
	_ = log
	val, err := readBuffer.ReadUint16("BACnetLiftCarDirection", 16)
	if err != nil {
		return 0, errors.Wrap(err, "error reading BACnetLiftCarDirection")
	}
	if enum, ok := BACnetLiftCarDirectionByValue(val); !ok {
		log.Debug().Interface("val", val).Msg("no value val found for BACnetLiftCarDirection")
		return BACnetLiftCarDirection(val), nil
	} else {
		return enum, nil
	}
}

func (e BACnetLiftCarDirection) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e BACnetLiftCarDirection) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	log := zerolog.Ctx(ctx)
	_ = log
	return writeBuffer.WriteUint16("BACnetLiftCarDirection", 16, uint16(uint16(e)), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e BACnetLiftCarDirection) PLC4XEnumName() string {
	switch e {
	case BACnetLiftCarDirection_UNKNOWN:
		return "UNKNOWN"
	case BACnetLiftCarDirection_VENDOR_PROPRIETARY_VALUE:
		return "VENDOR_PROPRIETARY_VALUE"
	case BACnetLiftCarDirection_NONE:
		return "NONE"
	case BACnetLiftCarDirection_STOPPED:
		return "STOPPED"
	case BACnetLiftCarDirection_UP:
		return "UP"
	case BACnetLiftCarDirection_DOWN:
		return "DOWN"
	case BACnetLiftCarDirection_UP_AND_DOWN:
		return "UP_AND_DOWN"
	}
	return fmt.Sprintf("Unknown(%v)", uint16(e))
}

func (e BACnetLiftCarDirection) String() string {
	return e.PLC4XEnumName()
}
