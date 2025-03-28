/**
 * @fileoverview gRPC-Web generated client stub for viam.component.arm.v1
 * @enhanceable
 * @public
 */

// Code generated by protoc-gen-grpc-web. DO NOT EDIT.
// versions:
// 	protoc-gen-grpc-web v1.5.0
// 	protoc              v0.0.0
// source: component/arm/v1/arm.proto


/* eslint-disable */
// @ts-nocheck



const grpc = {};
grpc.web = require('grpc-web');


var common_v1_common_pb = require('../../../common/v1/common_pb.js')

var google_api_annotations_pb = require('../../../google/api/annotations_pb.js')

var google_protobuf_struct_pb = require('google-protobuf/google/protobuf/struct_pb.js')
const proto = {};
proto.viam = {};
proto.viam.component = {};
proto.viam.component.arm = {};
proto.viam.component.arm.v1 = require('./arm_pb.js');

/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?grpc.web.ClientOptions} options
 * @constructor
 * @struct
 * @final
 */
proto.viam.component.arm.v1.ArmServiceClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options.format = 'text';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname.replace(/\/+$/, '');

};


/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?grpc.web.ClientOptions} options
 * @constructor
 * @struct
 * @final
 */
proto.viam.component.arm.v1.ArmServicePromiseClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options.format = 'text';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname.replace(/\/+$/, '');

};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.viam.component.arm.v1.GetEndPositionRequest,
 *   !proto.viam.component.arm.v1.GetEndPositionResponse>}
 */
const methodDescriptor_ArmService_GetEndPosition = new grpc.web.MethodDescriptor(
  '/viam.component.arm.v1.ArmService/GetEndPosition',
  grpc.web.MethodType.UNARY,
  proto.viam.component.arm.v1.GetEndPositionRequest,
  proto.viam.component.arm.v1.GetEndPositionResponse,
  /**
   * @param {!proto.viam.component.arm.v1.GetEndPositionRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.viam.component.arm.v1.GetEndPositionResponse.deserializeBinary
);


/**
 * @param {!proto.viam.component.arm.v1.GetEndPositionRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.viam.component.arm.v1.GetEndPositionResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.viam.component.arm.v1.GetEndPositionResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.viam.component.arm.v1.ArmServiceClient.prototype.getEndPosition =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/viam.component.arm.v1.ArmService/GetEndPosition',
      request,
      metadata || {},
      methodDescriptor_ArmService_GetEndPosition,
      callback);
};


/**
 * @param {!proto.viam.component.arm.v1.GetEndPositionRequest} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.viam.component.arm.v1.GetEndPositionResponse>}
 *     Promise that resolves to the response
 */
proto.viam.component.arm.v1.ArmServicePromiseClient.prototype.getEndPosition =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/viam.component.arm.v1.ArmService/GetEndPosition',
      request,
      metadata || {},
      methodDescriptor_ArmService_GetEndPosition);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.viam.component.arm.v1.MoveToPositionRequest,
 *   !proto.viam.component.arm.v1.MoveToPositionResponse>}
 */
const methodDescriptor_ArmService_MoveToPosition = new grpc.web.MethodDescriptor(
  '/viam.component.arm.v1.ArmService/MoveToPosition',
  grpc.web.MethodType.UNARY,
  proto.viam.component.arm.v1.MoveToPositionRequest,
  proto.viam.component.arm.v1.MoveToPositionResponse,
  /**
   * @param {!proto.viam.component.arm.v1.MoveToPositionRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.viam.component.arm.v1.MoveToPositionResponse.deserializeBinary
);


/**
 * @param {!proto.viam.component.arm.v1.MoveToPositionRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.viam.component.arm.v1.MoveToPositionResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.viam.component.arm.v1.MoveToPositionResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.viam.component.arm.v1.ArmServiceClient.prototype.moveToPosition =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/viam.component.arm.v1.ArmService/MoveToPosition',
      request,
      metadata || {},
      methodDescriptor_ArmService_MoveToPosition,
      callback);
};


/**
 * @param {!proto.viam.component.arm.v1.MoveToPositionRequest} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.viam.component.arm.v1.MoveToPositionResponse>}
 *     Promise that resolves to the response
 */
proto.viam.component.arm.v1.ArmServicePromiseClient.prototype.moveToPosition =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/viam.component.arm.v1.ArmService/MoveToPosition',
      request,
      metadata || {},
      methodDescriptor_ArmService_MoveToPosition);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.viam.component.arm.v1.GetJointPositionsRequest,
 *   !proto.viam.component.arm.v1.GetJointPositionsResponse>}
 */
const methodDescriptor_ArmService_GetJointPositions = new grpc.web.MethodDescriptor(
  '/viam.component.arm.v1.ArmService/GetJointPositions',
  grpc.web.MethodType.UNARY,
  proto.viam.component.arm.v1.GetJointPositionsRequest,
  proto.viam.component.arm.v1.GetJointPositionsResponse,
  /**
   * @param {!proto.viam.component.arm.v1.GetJointPositionsRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.viam.component.arm.v1.GetJointPositionsResponse.deserializeBinary
);


/**
 * @param {!proto.viam.component.arm.v1.GetJointPositionsRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.viam.component.arm.v1.GetJointPositionsResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.viam.component.arm.v1.GetJointPositionsResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.viam.component.arm.v1.ArmServiceClient.prototype.getJointPositions =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/viam.component.arm.v1.ArmService/GetJointPositions',
      request,
      metadata || {},
      methodDescriptor_ArmService_GetJointPositions,
      callback);
};


/**
 * @param {!proto.viam.component.arm.v1.GetJointPositionsRequest} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.viam.component.arm.v1.GetJointPositionsResponse>}
 *     Promise that resolves to the response
 */
proto.viam.component.arm.v1.ArmServicePromiseClient.prototype.getJointPositions =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/viam.component.arm.v1.ArmService/GetJointPositions',
      request,
      metadata || {},
      methodDescriptor_ArmService_GetJointPositions);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.viam.component.arm.v1.MoveToJointPositionsRequest,
 *   !proto.viam.component.arm.v1.MoveToJointPositionsResponse>}
 */
const methodDescriptor_ArmService_MoveToJointPositions = new grpc.web.MethodDescriptor(
  '/viam.component.arm.v1.ArmService/MoveToJointPositions',
  grpc.web.MethodType.UNARY,
  proto.viam.component.arm.v1.MoveToJointPositionsRequest,
  proto.viam.component.arm.v1.MoveToJointPositionsResponse,
  /**
   * @param {!proto.viam.component.arm.v1.MoveToJointPositionsRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.viam.component.arm.v1.MoveToJointPositionsResponse.deserializeBinary
);


/**
 * @param {!proto.viam.component.arm.v1.MoveToJointPositionsRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.viam.component.arm.v1.MoveToJointPositionsResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.viam.component.arm.v1.MoveToJointPositionsResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.viam.component.arm.v1.ArmServiceClient.prototype.moveToJointPositions =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/viam.component.arm.v1.ArmService/MoveToJointPositions',
      request,
      metadata || {},
      methodDescriptor_ArmService_MoveToJointPositions,
      callback);
};


/**
 * @param {!proto.viam.component.arm.v1.MoveToJointPositionsRequest} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.viam.component.arm.v1.MoveToJointPositionsResponse>}
 *     Promise that resolves to the response
 */
proto.viam.component.arm.v1.ArmServicePromiseClient.prototype.moveToJointPositions =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/viam.component.arm.v1.ArmService/MoveToJointPositions',
      request,
      metadata || {},
      methodDescriptor_ArmService_MoveToJointPositions);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.viam.component.arm.v1.MoveThroughJointPositionsRequest,
 *   !proto.viam.component.arm.v1.MoveThroughJointPositionsResponse>}
 */
const methodDescriptor_ArmService_MoveThroughJointPositions = new grpc.web.MethodDescriptor(
  '/viam.component.arm.v1.ArmService/MoveThroughJointPositions',
  grpc.web.MethodType.UNARY,
  proto.viam.component.arm.v1.MoveThroughJointPositionsRequest,
  proto.viam.component.arm.v1.MoveThroughJointPositionsResponse,
  /**
   * @param {!proto.viam.component.arm.v1.MoveThroughJointPositionsRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.viam.component.arm.v1.MoveThroughJointPositionsResponse.deserializeBinary
);


/**
 * @param {!proto.viam.component.arm.v1.MoveThroughJointPositionsRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.viam.component.arm.v1.MoveThroughJointPositionsResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.viam.component.arm.v1.MoveThroughJointPositionsResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.viam.component.arm.v1.ArmServiceClient.prototype.moveThroughJointPositions =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/viam.component.arm.v1.ArmService/MoveThroughJointPositions',
      request,
      metadata || {},
      methodDescriptor_ArmService_MoveThroughJointPositions,
      callback);
};


/**
 * @param {!proto.viam.component.arm.v1.MoveThroughJointPositionsRequest} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.viam.component.arm.v1.MoveThroughJointPositionsResponse>}
 *     Promise that resolves to the response
 */
proto.viam.component.arm.v1.ArmServicePromiseClient.prototype.moveThroughJointPositions =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/viam.component.arm.v1.ArmService/MoveThroughJointPositions',
      request,
      metadata || {},
      methodDescriptor_ArmService_MoveThroughJointPositions);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.viam.component.arm.v1.StopRequest,
 *   !proto.viam.component.arm.v1.StopResponse>}
 */
const methodDescriptor_ArmService_Stop = new grpc.web.MethodDescriptor(
  '/viam.component.arm.v1.ArmService/Stop',
  grpc.web.MethodType.UNARY,
  proto.viam.component.arm.v1.StopRequest,
  proto.viam.component.arm.v1.StopResponse,
  /**
   * @param {!proto.viam.component.arm.v1.StopRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.viam.component.arm.v1.StopResponse.deserializeBinary
);


/**
 * @param {!proto.viam.component.arm.v1.StopRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.viam.component.arm.v1.StopResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.viam.component.arm.v1.StopResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.viam.component.arm.v1.ArmServiceClient.prototype.stop =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/viam.component.arm.v1.ArmService/Stop',
      request,
      metadata || {},
      methodDescriptor_ArmService_Stop,
      callback);
};


/**
 * @param {!proto.viam.component.arm.v1.StopRequest} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.viam.component.arm.v1.StopResponse>}
 *     Promise that resolves to the response
 */
proto.viam.component.arm.v1.ArmServicePromiseClient.prototype.stop =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/viam.component.arm.v1.ArmService/Stop',
      request,
      metadata || {},
      methodDescriptor_ArmService_Stop);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.viam.component.arm.v1.IsMovingRequest,
 *   !proto.viam.component.arm.v1.IsMovingResponse>}
 */
const methodDescriptor_ArmService_IsMoving = new grpc.web.MethodDescriptor(
  '/viam.component.arm.v1.ArmService/IsMoving',
  grpc.web.MethodType.UNARY,
  proto.viam.component.arm.v1.IsMovingRequest,
  proto.viam.component.arm.v1.IsMovingResponse,
  /**
   * @param {!proto.viam.component.arm.v1.IsMovingRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.viam.component.arm.v1.IsMovingResponse.deserializeBinary
);


/**
 * @param {!proto.viam.component.arm.v1.IsMovingRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.viam.component.arm.v1.IsMovingResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.viam.component.arm.v1.IsMovingResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.viam.component.arm.v1.ArmServiceClient.prototype.isMoving =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/viam.component.arm.v1.ArmService/IsMoving',
      request,
      metadata || {},
      methodDescriptor_ArmService_IsMoving,
      callback);
};


/**
 * @param {!proto.viam.component.arm.v1.IsMovingRequest} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.viam.component.arm.v1.IsMovingResponse>}
 *     Promise that resolves to the response
 */
proto.viam.component.arm.v1.ArmServicePromiseClient.prototype.isMoving =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/viam.component.arm.v1.ArmService/IsMoving',
      request,
      metadata || {},
      methodDescriptor_ArmService_IsMoving);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.viam.common.v1.DoCommandRequest,
 *   !proto.viam.common.v1.DoCommandResponse>}
 */
const methodDescriptor_ArmService_DoCommand = new grpc.web.MethodDescriptor(
  '/viam.component.arm.v1.ArmService/DoCommand',
  grpc.web.MethodType.UNARY,
  common_v1_common_pb.DoCommandRequest,
  common_v1_common_pb.DoCommandResponse,
  /**
   * @param {!proto.viam.common.v1.DoCommandRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  common_v1_common_pb.DoCommandResponse.deserializeBinary
);


/**
 * @param {!proto.viam.common.v1.DoCommandRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.viam.common.v1.DoCommandResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.viam.common.v1.DoCommandResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.viam.component.arm.v1.ArmServiceClient.prototype.doCommand =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/viam.component.arm.v1.ArmService/DoCommand',
      request,
      metadata || {},
      methodDescriptor_ArmService_DoCommand,
      callback);
};


/**
 * @param {!proto.viam.common.v1.DoCommandRequest} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.viam.common.v1.DoCommandResponse>}
 *     Promise that resolves to the response
 */
proto.viam.component.arm.v1.ArmServicePromiseClient.prototype.doCommand =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/viam.component.arm.v1.ArmService/DoCommand',
      request,
      metadata || {},
      methodDescriptor_ArmService_DoCommand);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.viam.common.v1.GetKinematicsRequest,
 *   !proto.viam.common.v1.GetKinematicsResponse>}
 */
const methodDescriptor_ArmService_GetKinematics = new grpc.web.MethodDescriptor(
  '/viam.component.arm.v1.ArmService/GetKinematics',
  grpc.web.MethodType.UNARY,
  common_v1_common_pb.GetKinematicsRequest,
  common_v1_common_pb.GetKinematicsResponse,
  /**
   * @param {!proto.viam.common.v1.GetKinematicsRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  common_v1_common_pb.GetKinematicsResponse.deserializeBinary
);


/**
 * @param {!proto.viam.common.v1.GetKinematicsRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.viam.common.v1.GetKinematicsResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.viam.common.v1.GetKinematicsResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.viam.component.arm.v1.ArmServiceClient.prototype.getKinematics =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/viam.component.arm.v1.ArmService/GetKinematics',
      request,
      metadata || {},
      methodDescriptor_ArmService_GetKinematics,
      callback);
};


/**
 * @param {!proto.viam.common.v1.GetKinematicsRequest} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.viam.common.v1.GetKinematicsResponse>}
 *     Promise that resolves to the response
 */
proto.viam.component.arm.v1.ArmServicePromiseClient.prototype.getKinematics =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/viam.component.arm.v1.ArmService/GetKinematics',
      request,
      metadata || {},
      methodDescriptor_ArmService_GetKinematics);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.viam.common.v1.GetGeometriesRequest,
 *   !proto.viam.common.v1.GetGeometriesResponse>}
 */
const methodDescriptor_ArmService_GetGeometries = new grpc.web.MethodDescriptor(
  '/viam.component.arm.v1.ArmService/GetGeometries',
  grpc.web.MethodType.UNARY,
  common_v1_common_pb.GetGeometriesRequest,
  common_v1_common_pb.GetGeometriesResponse,
  /**
   * @param {!proto.viam.common.v1.GetGeometriesRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  common_v1_common_pb.GetGeometriesResponse.deserializeBinary
);


/**
 * @param {!proto.viam.common.v1.GetGeometriesRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.viam.common.v1.GetGeometriesResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.viam.common.v1.GetGeometriesResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.viam.component.arm.v1.ArmServiceClient.prototype.getGeometries =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/viam.component.arm.v1.ArmService/GetGeometries',
      request,
      metadata || {},
      methodDescriptor_ArmService_GetGeometries,
      callback);
};


/**
 * @param {!proto.viam.common.v1.GetGeometriesRequest} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.viam.common.v1.GetGeometriesResponse>}
 *     Promise that resolves to the response
 */
proto.viam.component.arm.v1.ArmServicePromiseClient.prototype.getGeometries =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/viam.component.arm.v1.ArmService/GetGeometries',
      request,
      metadata || {},
      methodDescriptor_ArmService_GetGeometries);
};


module.exports = proto.viam.component.arm.v1;

