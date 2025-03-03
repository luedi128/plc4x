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

package bacnetip

import (
	"context"
	"fmt"
	"github.com/apache/plc4x/plc4go/spi/options"
	"github.com/rs/zerolog"
	"net"
	"net/url"
	"time"

	"github.com/apache/plc4x/plc4go/protocols/bacnetip/readwrite/model"
	"github.com/apache/plc4x/plc4go/spi"
	"github.com/apache/plc4x/plc4go/spi/default"
	"github.com/apache/plc4x/plc4go/spi/transports"
	"github.com/apache/plc4x/plc4go/spi/transports/udp"
	"github.com/pkg/errors"
)

// ApplicationLayerMessageCodec is a wrapper for MessageCodec which takes care of segmentation, retries etc.
//
//go:generate go run ../../tools/plc4xgenerator/gen.go -type=ApplicationLayerMessageCodec
type ApplicationLayerMessageCodec struct {
	bipSimpleApplication *BIPSimpleApplication
	messageCode          *MessageCodec
	deviceInfoCache      DeviceInfoCache

	localAddress  *net.UDPAddr `stringer:"true"`
	remoteAddress *net.UDPAddr `stringer:"true"`

	log zerolog.Logger `ignore:"true"`
}

func NewApplicationLayerMessageCodec(localLog zerolog.Logger, udpTransport *udp.Transport, transportUrl url.URL, options map[string][]string, localAddress *net.UDPAddr, remoteAddress *net.UDPAddr) (*ApplicationLayerMessageCodec, error) {
	// TODO: currently this is done by the BIP down below
	// Have the transport create a new transport-instance.
	//transportInstance, err := udpTransport.CreateTransportInstanceForLocalAddress(transportUrl, options, localAddress)
	//if err != nil {
	//	return nil, errors.Wrap(err, "error creating transport instance")
	//}
	a := &ApplicationLayerMessageCodec{
		localAddress:  localAddress,
		remoteAddress: remoteAddress,

		log: localLog,
	}
	address, err := NewAddress(localLog, localAddress)
	if err != nil {
		return nil, errors.Wrap(err, "error creating address")
	}
	// TODO: workaround for strange address parsing
	at := AddressTuple[string, uint16]{fmt.Sprintf("%d.%d.%d.%d", address.AddrAddress[0], address.AddrAddress[1], address.AddrAddress[2], address.AddrAddress[3]), *address.AddrPort}
	address.AddrTuple = &at
	application, err := NewBIPSimpleApplication(localLog, &LocalDeviceObject{
		NumberOfAPDURetries: func() *uint { retries := uint(10); return &retries }(),
	}, *address, &a.deviceInfoCache, nil)
	if err != nil {
		return nil, err
	}
	a.bipSimpleApplication = application
	// TODO: this is currently done by the BIP
	//a.messageCode = NewMessageCodec(transportInstance)
	return a, nil
}

func (m *ApplicationLayerMessageCodec) GetCodec() spi.MessageCodec {
	return m
}

func (m *ApplicationLayerMessageCodec) Connect() error {
	// TODO: this is currently done by the BIP
	//	return m.messageCode.Connect()
	return nil
}

func (m *ApplicationLayerMessageCodec) ConnectWithContext(ctx context.Context) error {
	// TODO: this is currently done by the BIP
	//	return m.messageCode.ConnectWithContext(ctx)
	return nil
}

func (m *ApplicationLayerMessageCodec) Disconnect() error {
	if err := m.bipSimpleApplication.Close(); err != nil {
		m.log.Error().Err(err).Msg("error closing application")
	}
	return m.messageCode.Disconnect()
}

func (m *ApplicationLayerMessageCodec) IsRunning() bool {
	return m.messageCode.IsRunning()
}

func (m *ApplicationLayerMessageCodec) Send(message spi.Message) error {
	address, err := NewAddress(m.log, m.remoteAddress)
	if err != nil {
		return err
	}
	iocb, err := NewIOCB(m.log, NewPDU(message, WithPDUDestination(address)), address)
	if err != nil {
		return errors.Wrap(err, "error creating IOCB")
	}
	go func() {
		go func() {
			if err := m.bipSimpleApplication.RequestIO(iocb); err != nil {
				m.log.Debug().Err(err).Msg("errored")
			}
		}()
		iocb.Wait()
		if iocb.ioError != nil {
			// TODO: handle error
			fmt.Printf("Err: %v\n", iocb.ioError)
		} else if iocb.ioResponse != nil {
			// TODO: response?
			fmt.Printf("Response: %v\n", iocb.ioResponse)
		} else {
			// TODO: what now?
		}
	}()
	return nil
}

func (m *ApplicationLayerMessageCodec) Expect(ctx context.Context, acceptsMessage spi.AcceptsMessage, handleMessage spi.HandleMessage, handleError spi.HandleError, ttl time.Duration) error {
	// TODO: implement me
	panic("not yet implemented")
}

func (m *ApplicationLayerMessageCodec) SendRequest(ctx context.Context, message spi.Message, acceptsMessage spi.AcceptsMessage, handleMessage spi.HandleMessage, handleError spi.HandleError, ttl time.Duration) error {
	address, err := NewAddress(m.log, m.remoteAddress)
	if err != nil {
		return err
	}
	iocb, err := NewIOCB(m.log, NewPDU(message, WithPDUDestination(address)), address)
	if err != nil {
		return errors.Wrap(err, "error creating IOCB")
	}
	go func() {
		go func() {
			if err := m.bipSimpleApplication.RequestIO(iocb); err != nil {

			}
		}()
		iocb.Wait()
		if err := iocb.ioError; err != nil {
			if err := handleError(err); err != nil {
				m.log.Debug().Err(err).Msg("error handling error")
				return
			}
		} else if response := iocb.ioResponse; response != nil {
			// TODO: we wrap it into a BVLC for now. Once we change the Readers etc. to accept apdus we can remove that
			tempBVLC := model.NewBVLCOriginalUnicastNPDU(
				model.NewNPDU(
					0,
					model.NewNPDUControl(
						false,
						false,
						false,
						false,
						model.NPDUNetworkPriority_NORMAL_MESSAGE,
					),
					nil,
					nil,
					nil,
					nil,
					nil,
					nil,
					nil,
					nil,
					response.GetMessage().(model.APDU),
					0,
				),
				0,
			)
			if acceptsMessage(tempBVLC) {
				if err := handleMessage(
					tempBVLC,
				); err != nil {
					m.log.Debug().Err(err).Msg("error handling message")
					return
				}
			}
		} else {
			// TODO: what now?
		}
	}()
	return nil
}

func (m *ApplicationLayerMessageCodec) GetDefaultIncomingMessageChannel() chan spi.Message {
	// TODO: this is currently done by the BIP
	//return m.messageCode.GetDefaultIncomingMessageChannel()
	return make(chan spi.Message)
}

//go:generate go run ../../tools/plc4xgenerator/gen.go -type=MessageCodec
type MessageCodec struct {
	_default.DefaultCodec

	passLogToModel bool
	log            zerolog.Logger `ignore:"true"`
}

func NewMessageCodec(transportInstance transports.TransportInstance) *MessageCodec {
	codec := &MessageCodec{}
	codec.DefaultCodec = _default.NewDefaultCodec(codec, transportInstance, _default.WithCustomMessageHandler(codec.handleCustomMessage))
	return codec
}

func (m *MessageCodec) GetCodec() spi.MessageCodec {
	return m
}

func (m *MessageCodec) Send(message spi.Message) error {
	m.log.Trace().Msg("Sending message")
	// Cast the message to the correct type of struct
	bvlcPacket := message.(model.BVLC)
	// Serialize the request
	theBytes, err := bvlcPacket.Serialize()
	if err != nil {
		return errors.Wrap(err, "error serializing request")
	}

	// Send it to the PLC
	err = m.GetTransportInstance().Write(theBytes)
	if err != nil {
		return errors.Wrap(err, "error sending request")
	}
	return nil
}

func (m *MessageCodec) Receive() (spi.Message, error) {
	// We need at least 6 bytes in order to know how big the packet is in total
	if num, err := m.GetTransportInstance().GetNumBytesAvailableInBuffer(); (err == nil) && (num >= 4) {
		m.log.Debug().Uint32("num", num).Msg("we got num readable bytes")
		data, err := m.GetTransportInstance().PeekReadableBytes(4)
		if err != nil {
			m.log.Warn().Err(err).Msg("error peeking")
			// TODO: Possibly clean up ...
			return nil, nil
		}
		packetSize := uint32((uint16(data[2]) << 8) + uint16(data[3]))
		if num < packetSize {
			m.log.Debug().
				Uint32("num", num).
				Uint32("packetSize", packetSize).
				Msg("Not enough bytes. Got: num Need: packetSize")
			return nil, nil
		}
		data, err = m.GetTransportInstance().Read(packetSize)
		if err != nil {
			m.log.Debug().Err(err).Msg("Error reading")
			// TODO: Possibly clean up ...
			return nil, nil
		}
		ctxForModel := options.GetLoggerContextForModel(context.TODO(), m.log, options.WithPassLoggerToModel(m.passLogToModel))
		bvlcPacket, err := model.BVLCParse(ctxForModel, data)
		if err != nil {
			m.log.Warn().Err(err).Msg("error parsing")
			// TODO: Possibly clean up ...
			return nil, nil
		}
		return bvlcPacket, nil
	} else if err != nil {
		m.log.Warn().Err(err).Msg("Got error reading")
		return nil, nil
	}
	// TODO: maybe we return here a not enough error error
	return nil, nil
}

func (m *MessageCodec) handleCustomMessage(_ _default.DefaultCodecRequirements, message spi.Message) bool {
	// For now, we just put them in the incoming channel
	m.GetDefaultIncomingMessageChannel() <- message
	return true
}
