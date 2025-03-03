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
package org.apache.plc4x.java.opcua.readwrite;

import static org.apache.plc4x.java.spi.codegen.fields.FieldReaderFactory.*;
import static org.apache.plc4x.java.spi.codegen.fields.FieldWriterFactory.*;
import static org.apache.plc4x.java.spi.codegen.io.DataReaderFactory.*;
import static org.apache.plc4x.java.spi.codegen.io.DataWriterFactory.*;
import static org.apache.plc4x.java.spi.generation.StaticHelper.*;

import java.time.*;
import java.util.*;
import org.apache.plc4x.java.api.exceptions.*;
import org.apache.plc4x.java.api.value.*;
import org.apache.plc4x.java.spi.codegen.*;
import org.apache.plc4x.java.spi.codegen.fields.*;
import org.apache.plc4x.java.spi.codegen.io.*;
import org.apache.plc4x.java.spi.generation.*;

// Code generated by code-generation. DO NOT EDIT.

public class OpcuaProtocolLimits implements Message {

  // Properties.
  protected final long receiveBufferSize;
  protected final long sendBufferSize;
  protected final long maxMessageSize;
  protected final long maxChunkCount;

  public OpcuaProtocolLimits(
      long receiveBufferSize, long sendBufferSize, long maxMessageSize, long maxChunkCount) {
    super();
    this.receiveBufferSize = receiveBufferSize;
    this.sendBufferSize = sendBufferSize;
    this.maxMessageSize = maxMessageSize;
    this.maxChunkCount = maxChunkCount;
  }

  public long getReceiveBufferSize() {
    return receiveBufferSize;
  }

  public long getSendBufferSize() {
    return sendBufferSize;
  }

  public long getMaxMessageSize() {
    return maxMessageSize;
  }

  public long getMaxChunkCount() {
    return maxChunkCount;
  }

  public void serialize(WriteBuffer writeBuffer) throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("OpcuaProtocolLimits");

    // Simple Field (receiveBufferSize)
    writeSimpleField("receiveBufferSize", receiveBufferSize, writeUnsignedLong(writeBuffer, 32));

    // Simple Field (sendBufferSize)
    writeSimpleField("sendBufferSize", sendBufferSize, writeUnsignedLong(writeBuffer, 32));

    // Simple Field (maxMessageSize)
    writeSimpleField("maxMessageSize", maxMessageSize, writeUnsignedLong(writeBuffer, 32));

    // Simple Field (maxChunkCount)
    writeSimpleField("maxChunkCount", maxChunkCount, writeUnsignedLong(writeBuffer, 32));

    writeBuffer.popContext("OpcuaProtocolLimits");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = 0;
    OpcuaProtocolLimits _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (receiveBufferSize)
    lengthInBits += 32;

    // Simple field (sendBufferSize)
    lengthInBits += 32;

    // Simple field (maxMessageSize)
    lengthInBits += 32;

    // Simple field (maxChunkCount)
    lengthInBits += 32;

    return lengthInBits;
  }

  public static OpcuaProtocolLimits staticParse(ReadBuffer readBuffer, Object... args)
      throws ParseException {
    PositionAware positionAware = readBuffer;
    return staticParse(readBuffer);
  }

  public static OpcuaProtocolLimits staticParse(ReadBuffer readBuffer) throws ParseException {
    readBuffer.pullContext("OpcuaProtocolLimits");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    long receiveBufferSize = readSimpleField("receiveBufferSize", readUnsignedLong(readBuffer, 32));

    long sendBufferSize = readSimpleField("sendBufferSize", readUnsignedLong(readBuffer, 32));

    long maxMessageSize = readSimpleField("maxMessageSize", readUnsignedLong(readBuffer, 32));

    long maxChunkCount = readSimpleField("maxChunkCount", readUnsignedLong(readBuffer, 32));

    readBuffer.closeContext("OpcuaProtocolLimits");
    // Create the instance
    OpcuaProtocolLimits _opcuaProtocolLimits;
    _opcuaProtocolLimits =
        new OpcuaProtocolLimits(receiveBufferSize, sendBufferSize, maxMessageSize, maxChunkCount);
    return _opcuaProtocolLimits;
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof OpcuaProtocolLimits)) {
      return false;
    }
    OpcuaProtocolLimits that = (OpcuaProtocolLimits) o;
    return (getReceiveBufferSize() == that.getReceiveBufferSize())
        && (getSendBufferSize() == that.getSendBufferSize())
        && (getMaxMessageSize() == that.getMaxMessageSize())
        && (getMaxChunkCount() == that.getMaxChunkCount())
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(
        getReceiveBufferSize(), getSendBufferSize(), getMaxMessageSize(), getMaxChunkCount());
  }

  @Override
  public String toString() {
    WriteBufferBoxBased writeBufferBoxBased = new WriteBufferBoxBased(true, true);
    try {
      writeBufferBoxBased.writeSerializable(this);
    } catch (SerializationException e) {
      throw new RuntimeException(e);
    }
    return "\n" + writeBufferBoxBased.getBox().toString() + "\n";
  }
}
