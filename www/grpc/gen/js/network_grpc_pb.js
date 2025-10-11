// GENERATED CODE -- DO NOT EDIT!

'use strict';
var grpc = require('@grpc/grpc-js');
var network_pb = require('./network_pb.js');

function serialize_aerium_GetNetworkInfoRequest(arg) {
  if (!(arg instanceof network_pb.GetNetworkInfoRequest)) {
    throw new Error('Expected argument of type aerium.GetNetworkInfoRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_aerium_GetNetworkInfoRequest(buffer_arg) {
  return network_pb.GetNetworkInfoRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_aerium_GetNetworkInfoResponse(arg) {
  if (!(arg instanceof network_pb.GetNetworkInfoResponse)) {
    throw new Error('Expected argument of type aerium.GetNetworkInfoResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_aerium_GetNetworkInfoResponse(buffer_arg) {
  return network_pb.GetNetworkInfoResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_aerium_GetNodeInfoRequest(arg) {
  if (!(arg instanceof network_pb.GetNodeInfoRequest)) {
    throw new Error('Expected argument of type aerium.GetNodeInfoRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_aerium_GetNodeInfoRequest(buffer_arg) {
  return network_pb.GetNodeInfoRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_aerium_GetNodeInfoResponse(arg) {
  if (!(arg instanceof network_pb.GetNodeInfoResponse)) {
    throw new Error('Expected argument of type aerium.GetNodeInfoResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_aerium_GetNodeInfoResponse(buffer_arg) {
  return network_pb.GetNodeInfoResponse.deserializeBinary(new Uint8Array(buffer_arg));
}


// Network service provides RPCs for retrieving information about the network.
var NetworkService = exports.NetworkService = {
  // GetNetworkInfo retrieves information about the overall network.
getNetworkInfo: {
    path: '/aerium.Network/GetNetworkInfo',
    requestStream: false,
    responseStream: false,
    requestType: network_pb.GetNetworkInfoRequest,
    responseType: network_pb.GetNetworkInfoResponse,
    requestSerialize: serialize_aerium_GetNetworkInfoRequest,
    requestDeserialize: deserialize_aerium_GetNetworkInfoRequest,
    responseSerialize: serialize_aerium_GetNetworkInfoResponse,
    responseDeserialize: deserialize_aerium_GetNetworkInfoResponse,
  },
  // GetNodeInfo retrieves information about a specific node in the network.
getNodeInfo: {
    path: '/aerium.Network/GetNodeInfo',
    requestStream: false,
    responseStream: false,
    requestType: network_pb.GetNodeInfoRequest,
    responseType: network_pb.GetNodeInfoResponse,
    requestSerialize: serialize_aerium_GetNodeInfoRequest,
    requestDeserialize: deserialize_aerium_GetNodeInfoRequest,
    responseSerialize: serialize_aerium_GetNodeInfoResponse,
    responseDeserialize: deserialize_aerium_GetNodeInfoResponse,
  },
};

exports.NetworkClient = grpc.makeGenericClientConstructor(NetworkService, 'Network');
