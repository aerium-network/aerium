// GENERATED CODE -- DO NOT EDIT!

'use strict';
var grpc = require('@grpc/grpc-js');
var utils_pb = require('./utils_pb.js');

function serialize_aerium_PublicKeyAggregationRequest(arg) {
  if (!(arg instanceof utils_pb.PublicKeyAggregationRequest)) {
    throw new Error('Expected argument of type aerium.PublicKeyAggregationRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_aerium_PublicKeyAggregationRequest(buffer_arg) {
  return utils_pb.PublicKeyAggregationRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_aerium_PublicKeyAggregationResponse(arg) {
  if (!(arg instanceof utils_pb.PublicKeyAggregationResponse)) {
    throw new Error('Expected argument of type aerium.PublicKeyAggregationResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_aerium_PublicKeyAggregationResponse(buffer_arg) {
  return utils_pb.PublicKeyAggregationResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_aerium_SignMessageWithPrivateKeyRequest(arg) {
  if (!(arg instanceof utils_pb.SignMessageWithPrivateKeyRequest)) {
    throw new Error('Expected argument of type aerium.SignMessageWithPrivateKeyRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_aerium_SignMessageWithPrivateKeyRequest(buffer_arg) {
  return utils_pb.SignMessageWithPrivateKeyRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_aerium_SignMessageWithPrivateKeyResponse(arg) {
  if (!(arg instanceof utils_pb.SignMessageWithPrivateKeyResponse)) {
    throw new Error('Expected argument of type aerium.SignMessageWithPrivateKeyResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_aerium_SignMessageWithPrivateKeyResponse(buffer_arg) {
  return utils_pb.SignMessageWithPrivateKeyResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_aerium_SignatureAggregationRequest(arg) {
  if (!(arg instanceof utils_pb.SignatureAggregationRequest)) {
    throw new Error('Expected argument of type aerium.SignatureAggregationRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_aerium_SignatureAggregationRequest(buffer_arg) {
  return utils_pb.SignatureAggregationRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_aerium_SignatureAggregationResponse(arg) {
  if (!(arg instanceof utils_pb.SignatureAggregationResponse)) {
    throw new Error('Expected argument of type aerium.SignatureAggregationResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_aerium_SignatureAggregationResponse(buffer_arg) {
  return utils_pb.SignatureAggregationResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_aerium_VerifyMessageRequest(arg) {
  if (!(arg instanceof utils_pb.VerifyMessageRequest)) {
    throw new Error('Expected argument of type aerium.VerifyMessageRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_aerium_VerifyMessageRequest(buffer_arg) {
  return utils_pb.VerifyMessageRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_aerium_VerifyMessageResponse(arg) {
  if (!(arg instanceof utils_pb.VerifyMessageResponse)) {
    throw new Error('Expected argument of type aerium.VerifyMessageResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_aerium_VerifyMessageResponse(buffer_arg) {
  return utils_pb.VerifyMessageResponse.deserializeBinary(new Uint8Array(buffer_arg));
}


// The Utils service defines RPC methods for utility functions such as message signing, verification, and more.
var UtilsService = exports.UtilsService = {
  // Signs a message with the provided private key.
signMessageWithPrivateKey: {
    path: '/aerium.Utils/SignMessageWithPrivateKey',
    requestStream: false,
    responseStream: false,
    requestType: utils_pb.SignMessageWithPrivateKeyRequest,
    responseType: utils_pb.SignMessageWithPrivateKeyResponse,
    requestSerialize: serialize_aerium_SignMessageWithPrivateKeyRequest,
    requestDeserialize: deserialize_aerium_SignMessageWithPrivateKeyRequest,
    responseSerialize: serialize_aerium_SignMessageWithPrivateKeyResponse,
    responseDeserialize: deserialize_aerium_SignMessageWithPrivateKeyResponse,
  },
  // Verifies a signature against the public key and message.
verifyMessage: {
    path: '/aerium.Utils/VerifyMessage',
    requestStream: false,
    responseStream: false,
    requestType: utils_pb.VerifyMessageRequest,
    responseType: utils_pb.VerifyMessageResponse,
    requestSerialize: serialize_aerium_VerifyMessageRequest,
    requestDeserialize: deserialize_aerium_VerifyMessageRequest,
    responseSerialize: serialize_aerium_VerifyMessageResponse,
    responseDeserialize: deserialize_aerium_VerifyMessageResponse,
  },
  // Aggregates multiple BLS public keys into a single key.
publicKeyAggregation: {
    path: '/aerium.Utils/PublicKeyAggregation',
    requestStream: false,
    responseStream: false,
    requestType: utils_pb.PublicKeyAggregationRequest,
    responseType: utils_pb.PublicKeyAggregationResponse,
    requestSerialize: serialize_aerium_PublicKeyAggregationRequest,
    requestDeserialize: deserialize_aerium_PublicKeyAggregationRequest,
    responseSerialize: serialize_aerium_PublicKeyAggregationResponse,
    responseDeserialize: deserialize_aerium_PublicKeyAggregationResponse,
  },
  // Aggregates multiple BLS signatures into a single signature.
signatureAggregation: {
    path: '/aerium.Utils/SignatureAggregation',
    requestStream: false,
    responseStream: false,
    requestType: utils_pb.SignatureAggregationRequest,
    responseType: utils_pb.SignatureAggregationResponse,
    requestSerialize: serialize_aerium_SignatureAggregationRequest,
    requestDeserialize: deserialize_aerium_SignatureAggregationRequest,
    responseSerialize: serialize_aerium_SignatureAggregationResponse,
    responseDeserialize: deserialize_aerium_SignatureAggregationResponse,
  },
};

exports.UtilsClient = grpc.makeGenericClientConstructor(UtilsService, 'Utils');
