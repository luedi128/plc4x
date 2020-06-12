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
#ifndef PLC4C_S7_READ_WRITE_S7_PARAMETER_H_
#define PLC4C_S7_READ_WRITE_S7_PARAMETER_H_
#ifdef __cplusplus
extern "C" {
#endif

#include <stdbool.h>
#include <stdint.h>
#include <plc4c/utils/list.h>
#include "s7_var_request_parameter_item.h"
#include "s7_parameter.h"
#include "s7_parameter.h"
#include "s7_parameter.h"
#include "s7_var_request_parameter_item.h"
#include "s7_parameter.h"
#include "s7_parameter.h"
#include "s7_parameter_user_data_item.h"
#include "s7_parameter.h"

// Structure used to contain the discriminator values for discriminated types using this as a parent
struct plc4c_s7_read_write_s7_parameter_discriminator {
  uint8_t messageType;
  uint8_t parameterType;
};
typedef struct plc4c_s7_read_write_s7_parameter_discriminator plc4c_s7_read_write_s7_parameter_discriminator;

// Enum assigning each sub-type an individual id.
enum plc4c_s7_read_write_s7_parameter_type {
  plc4c_s7_read_write_s7_parameter_type_s7_read_write_s7_parameter_setup_communication = 0,
  plc4c_s7_read_write_s7_parameter_type_s7_read_write_s7_parameter_read_var_request = 1,
  plc4c_s7_read_write_s7_parameter_type_s7_read_write_s7_parameter_read_var_response = 2,
  plc4c_s7_read_write_s7_parameter_type_s7_read_write_s7_parameter_write_var_request = 3,
  plc4c_s7_read_write_s7_parameter_type_s7_read_write_s7_parameter_write_var_response = 4,
  plc4c_s7_read_write_s7_parameter_type_s7_read_write_s7_parameter_user_data = 5};
typedef enum plc4c_s7_read_write_s7_parameter_type plc4c_s7_read_write_s7_parameter_type;

// Function to get the discriminator values for a given type.
plc4c_s7_read_write_s7_parameter_discriminator plc4c_s7_read_write_s7_parameter_get_discriminator(plc4c_s7_read_write_s7_parameter_type type);

struct plc4c_s7_read_write_s7_parameter {
  /* This is an abstract type so this property saves the type of this typed union */
  plc4c_s7_read_write_s7_parameter_type _type;
  /* Properties */
  union {
    struct { /* S7ParameterSetupCommunication */
      uint16_t s7_parameter_setup_communication_max_amq_caller;
      uint16_t s7_parameter_setup_communication_max_amq_callee;
      uint16_t s7_parameter_setup_communication_pdu_length;
    };
    struct { /* S7ParameterReadVarRequest */
      plc4c_list* s7_parameter_read_var_request_items;
    };
    struct { /* S7ParameterReadVarResponse */
      uint8_t s7_parameter_read_var_response_num_items;
    };
    struct { /* S7ParameterWriteVarRequest */
      plc4c_list* s7_parameter_write_var_request_items;
    };
    struct { /* S7ParameterWriteVarResponse */
      uint8_t s7_parameter_write_var_response_num_items;
    };
    struct { /* S7ParameterUserData */
      plc4c_list* s7_parameter_user_data_items;
    };
  };
};
typedef struct plc4c_s7_read_write_s7_parameter plc4c_s7_read_write_s7_parameter;

plc4c_return_code plc4c_s7_read_write_s7_parameter_parse(plc4c_spi_read_buffer* buf, uint8_t messageType, plc4c_s7_read_write_s7_parameter** message);

plc4c_return_code plc4c_s7_read_write_s7_parameter_serialize(plc4c_spi_write_buffer* buf, plc4c_s7_read_write_s7_parameter* message);

#ifdef __cplusplus
}
#endif
#endif  // PLC4C_S7_READ_WRITE_S7_PARAMETER_H_
