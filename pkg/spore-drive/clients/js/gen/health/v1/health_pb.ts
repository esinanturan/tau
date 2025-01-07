// @generated by protoc-gen-es v1.4.0 with parameter "target=ts"
// @generated from file health/v1/health.proto (package health.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";

/**
 * Data Structures
 *
 * @generated from message health.v1.SupportsRequest
 */
export class SupportsRequest extends Message<SupportsRequest> {
  /**
   * @generated from field: string version = 1;
   */
  version = "";

  constructor(data?: PartialMessage<SupportsRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "health.v1.SupportsRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "version", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): SupportsRequest {
    return new SupportsRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): SupportsRequest {
    return new SupportsRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): SupportsRequest {
    return new SupportsRequest().fromJsonString(jsonString, options);
  }

  static equals(a: SupportsRequest | PlainMessage<SupportsRequest> | undefined, b: SupportsRequest | PlainMessage<SupportsRequest> | undefined): boolean {
    return proto3.util.equals(SupportsRequest, a, b);
  }
}

/**
 * @generated from message health.v1.Empty
 */
export class Empty extends Message<Empty> {
  constructor(data?: PartialMessage<Empty>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "health.v1.Empty";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Empty {
    return new Empty().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Empty {
    return new Empty().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Empty {
    return new Empty().fromJsonString(jsonString, options);
  }

  static equals(a: Empty | PlainMessage<Empty> | undefined, b: Empty | PlainMessage<Empty> | undefined): boolean {
    return proto3.util.equals(Empty, a, b);
  }
}

