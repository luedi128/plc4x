/*
  Licensed to the Apache Software Foundation (ASF) under one
  or more contributor license agreements.  See the NOTICE file
  distributed with this work for additional information
  regarding copyright ownership.  The ASF licenses this file
  to you under the Apache License, Version 2.0 (the
  "License"); you may not use this file except in compliance
  with the License.  You may obtain a copy of the License at

      http://www.apache.org/licenses/LICENSE-2.0

  Unless required by applicable law or agreed to in writing,
  software distributed under the License is distributed on an
  "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
  KIND, either express or implied.  See the License for the
  specific language governing permissions and limitations
  under the License.
*/

#include <stdio.h>
#include <plc4c/spi/read_buffer.h>
#include <plc4c/spi/write_buffer.h>
#include <plc4c/spi/evaluation_helper.h>
#include "modbus_pdu_write_file_record_request_item.h"

// Parse function.
plc4c_return_code plc4c_modbus_read_write_modbus_pdu_write_file_record_request_item_parse(plc4c_spi_read_buffer* buf, plc4c_modbus_read_write_modbus_pdu_write_file_record_request_item** message) {
  uint16_t startPos = plc4c_spi_read_get_pos(buf);
  uint16_t curPos;

  // Pointer to the parsed data structure.
  plc4c_modbus_read_write_modbus_pdu_write_file_record_request_item* msg = malloc(sizeof(plc4c_modbus_read_write_modbus_pdu_write_file_record_request_item));

  // Simple Field (referenceType)
  uint8_t referenceType = plc4c_spi_read_unsigned_short(buf, 8);
  msg->reference_type = referenceType;

  // Simple Field (fileNumber)
  uint16_t fileNumber = plc4c_spi_read_unsigned_int(buf, 16);
  msg->file_number = fileNumber;

  // Simple Field (recordNumber)
  uint16_t recordNumber = plc4c_spi_read_unsigned_int(buf, 16);
  msg->record_number = recordNumber;

  // Implicit Field (recordLength) (Used for parsing, but it's value is not stored as it's implicitly given by the objects content)
  uint16_t recordLength = plc4c_spi_read_unsigned_int(buf, 16);

  // Array field (recordData)
  plc4c_list recordData;
  {
    // Length array
    uint8_t _recordDataLength = (recordLength) * (2);
    uint8_t recordDataEndPos = plc4c_spi_read_get_pos(buf) + _recordDataLength;
    while(plc4c_spi_read_get_pos(buf) < recordDataEndPos) {
      plc4c_utils_list_insert_head_value(&recordData, plc4c_spi_read_unsigned_int(buf, 16));
    }
  }


  return OK;
}

plc4c_return_code plc4c_modbus_read_write_modbus_pdu_write_file_record_request_item_serialize(plc4c_spi_write_buffer* buf, plc4c_modbus_read_write_modbus_pdu_write_file_record_request_item* message) {
  return OK;
}
