// GENERATED CODE -- DO NOT EDIT!

'use strict';
var grpc = require('@grpc/grpc-js');
var blockchain_pb = require('./blockchain_pb.js');
var transaction_pb = require('./transaction_pb.js');

function serialize_aerium_GetAccountRequest(arg) {
  if (!(arg instanceof blockchain_pb.GetAccountRequest)) {
    throw new Error('Expected argument of type aerium.GetAccountRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_aerium_GetAccountRequest(buffer_arg) {
  return blockchain_pb.GetAccountRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_aerium_GetAccountResponse(arg) {
  if (!(arg instanceof blockchain_pb.GetAccountResponse)) {
    throw new Error('Expected argument of type aerium.GetAccountResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_aerium_GetAccountResponse(buffer_arg) {
  return blockchain_pb.GetAccountResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_aerium_GetBlockHashRequest(arg) {
  if (!(arg instanceof blockchain_pb.GetBlockHashRequest)) {
    throw new Error('Expected argument of type aerium.GetBlockHashRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_aerium_GetBlockHashRequest(buffer_arg) {
  return blockchain_pb.GetBlockHashRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_aerium_GetBlockHashResponse(arg) {
  if (!(arg instanceof blockchain_pb.GetBlockHashResponse)) {
    throw new Error('Expected argument of type aerium.GetBlockHashResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_aerium_GetBlockHashResponse(buffer_arg) {
  return blockchain_pb.GetBlockHashResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_aerium_GetBlockHeightRequest(arg) {
  if (!(arg instanceof blockchain_pb.GetBlockHeightRequest)) {
    throw new Error('Expected argument of type aerium.GetBlockHeightRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_aerium_GetBlockHeightRequest(buffer_arg) {
  return blockchain_pb.GetBlockHeightRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_aerium_GetBlockHeightResponse(arg) {
  if (!(arg instanceof blockchain_pb.GetBlockHeightResponse)) {
    throw new Error('Expected argument of type aerium.GetBlockHeightResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_aerium_GetBlockHeightResponse(buffer_arg) {
  return blockchain_pb.GetBlockHeightResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_aerium_GetBlockRequest(arg) {
  if (!(arg instanceof blockchain_pb.GetBlockRequest)) {
    throw new Error('Expected argument of type aerium.GetBlockRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_aerium_GetBlockRequest(buffer_arg) {
  return blockchain_pb.GetBlockRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_aerium_GetBlockResponse(arg) {
  if (!(arg instanceof blockchain_pb.GetBlockResponse)) {
    throw new Error('Expected argument of type aerium.GetBlockResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_aerium_GetBlockResponse(buffer_arg) {
  return blockchain_pb.GetBlockResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_aerium_GetBlockchainInfoRequest(arg) {
  if (!(arg instanceof blockchain_pb.GetBlockchainInfoRequest)) {
    throw new Error('Expected argument of type aerium.GetBlockchainInfoRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_aerium_GetBlockchainInfoRequest(buffer_arg) {
  return blockchain_pb.GetBlockchainInfoRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_aerium_GetBlockchainInfoResponse(arg) {
  if (!(arg instanceof blockchain_pb.GetBlockchainInfoResponse)) {
    throw new Error('Expected argument of type aerium.GetBlockchainInfoResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_aerium_GetBlockchainInfoResponse(buffer_arg) {
  return blockchain_pb.GetBlockchainInfoResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_aerium_GetConsensusInfoRequest(arg) {
  if (!(arg instanceof blockchain_pb.GetConsensusInfoRequest)) {
    throw new Error('Expected argument of type aerium.GetConsensusInfoRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_aerium_GetConsensusInfoRequest(buffer_arg) {
  return blockchain_pb.GetConsensusInfoRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_aerium_GetConsensusInfoResponse(arg) {
  if (!(arg instanceof blockchain_pb.GetConsensusInfoResponse)) {
    throw new Error('Expected argument of type aerium.GetConsensusInfoResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_aerium_GetConsensusInfoResponse(buffer_arg) {
  return blockchain_pb.GetConsensusInfoResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_aerium_GetPublicKeyRequest(arg) {
  if (!(arg instanceof blockchain_pb.GetPublicKeyRequest)) {
    throw new Error('Expected argument of type aerium.GetPublicKeyRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_aerium_GetPublicKeyRequest(buffer_arg) {
  return blockchain_pb.GetPublicKeyRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_aerium_GetPublicKeyResponse(arg) {
  if (!(arg instanceof blockchain_pb.GetPublicKeyResponse)) {
    throw new Error('Expected argument of type aerium.GetPublicKeyResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_aerium_GetPublicKeyResponse(buffer_arg) {
  return blockchain_pb.GetPublicKeyResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_aerium_GetTxPoolContentRequest(arg) {
  if (!(arg instanceof blockchain_pb.GetTxPoolContentRequest)) {
    throw new Error('Expected argument of type aerium.GetTxPoolContentRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_aerium_GetTxPoolContentRequest(buffer_arg) {
  return blockchain_pb.GetTxPoolContentRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_aerium_GetTxPoolContentResponse(arg) {
  if (!(arg instanceof blockchain_pb.GetTxPoolContentResponse)) {
    throw new Error('Expected argument of type aerium.GetTxPoolContentResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_aerium_GetTxPoolContentResponse(buffer_arg) {
  return blockchain_pb.GetTxPoolContentResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_aerium_GetValidatorAddressesRequest(arg) {
  if (!(arg instanceof blockchain_pb.GetValidatorAddressesRequest)) {
    throw new Error('Expected argument of type aerium.GetValidatorAddressesRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_aerium_GetValidatorAddressesRequest(buffer_arg) {
  return blockchain_pb.GetValidatorAddressesRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_aerium_GetValidatorAddressesResponse(arg) {
  if (!(arg instanceof blockchain_pb.GetValidatorAddressesResponse)) {
    throw new Error('Expected argument of type aerium.GetValidatorAddressesResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_aerium_GetValidatorAddressesResponse(buffer_arg) {
  return blockchain_pb.GetValidatorAddressesResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_aerium_GetValidatorByNumberRequest(arg) {
  if (!(arg instanceof blockchain_pb.GetValidatorByNumberRequest)) {
    throw new Error('Expected argument of type aerium.GetValidatorByNumberRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_aerium_GetValidatorByNumberRequest(buffer_arg) {
  return blockchain_pb.GetValidatorByNumberRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_aerium_GetValidatorRequest(arg) {
  if (!(arg instanceof blockchain_pb.GetValidatorRequest)) {
    throw new Error('Expected argument of type aerium.GetValidatorRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_aerium_GetValidatorRequest(buffer_arg) {
  return blockchain_pb.GetValidatorRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_aerium_GetValidatorResponse(arg) {
  if (!(arg instanceof blockchain_pb.GetValidatorResponse)) {
    throw new Error('Expected argument of type aerium.GetValidatorResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_aerium_GetValidatorResponse(buffer_arg) {
  return blockchain_pb.GetValidatorResponse.deserializeBinary(new Uint8Array(buffer_arg));
}


// The Blockchain service defines RPC methods for interacting with the blockchain.
var BlockchainService = exports.BlockchainService = {
  // Retrieves information about a block based on the provided request parameters.
getBlock: {
    path: '/aerium.Blockchain/GetBlock',
    requestStream: false,
    responseStream: false,
    requestType: blockchain_pb.GetBlockRequest,
    responseType: blockchain_pb.GetBlockResponse,
    requestSerialize: serialize_aerium_GetBlockRequest,
    requestDeserialize: deserialize_aerium_GetBlockRequest,
    responseSerialize: serialize_aerium_GetBlockResponse,
    responseDeserialize: deserialize_aerium_GetBlockResponse,
  },
  // Retrieves the hash of a block at the specified height.
getBlockHash: {
    path: '/aerium.Blockchain/GetBlockHash',
    requestStream: false,
    responseStream: false,
    requestType: blockchain_pb.GetBlockHashRequest,
    responseType: blockchain_pb.GetBlockHashResponse,
    requestSerialize: serialize_aerium_GetBlockHashRequest,
    requestDeserialize: deserialize_aerium_GetBlockHashRequest,
    responseSerialize: serialize_aerium_GetBlockHashResponse,
    responseDeserialize: deserialize_aerium_GetBlockHashResponse,
  },
  // Retrieves the height of a block with the specified hash.
getBlockHeight: {
    path: '/aerium.Blockchain/GetBlockHeight',
    requestStream: false,
    responseStream: false,
    requestType: blockchain_pb.GetBlockHeightRequest,
    responseType: blockchain_pb.GetBlockHeightResponse,
    requestSerialize: serialize_aerium_GetBlockHeightRequest,
    requestDeserialize: deserialize_aerium_GetBlockHeightRequest,
    responseSerialize: serialize_aerium_GetBlockHeightResponse,
    responseDeserialize: deserialize_aerium_GetBlockHeightResponse,
  },
  // Retrieves general information about the blockchain.
getBlockchainInfo: {
    path: '/aerium.Blockchain/GetBlockchainInfo',
    requestStream: false,
    responseStream: false,
    requestType: blockchain_pb.GetBlockchainInfoRequest,
    responseType: blockchain_pb.GetBlockchainInfoResponse,
    requestSerialize: serialize_aerium_GetBlockchainInfoRequest,
    requestDeserialize: deserialize_aerium_GetBlockchainInfoRequest,
    responseSerialize: serialize_aerium_GetBlockchainInfoResponse,
    responseDeserialize: deserialize_aerium_GetBlockchainInfoResponse,
  },
  // Retrieves information about consensus instances.
getConsensusInfo: {
    path: '/aerium.Blockchain/GetConsensusInfo',
    requestStream: false,
    responseStream: false,
    requestType: blockchain_pb.GetConsensusInfoRequest,
    responseType: blockchain_pb.GetConsensusInfoResponse,
    requestSerialize: serialize_aerium_GetConsensusInfoRequest,
    requestDeserialize: deserialize_aerium_GetConsensusInfoRequest,
    responseSerialize: serialize_aerium_GetConsensusInfoResponse,
    responseDeserialize: deserialize_aerium_GetConsensusInfoResponse,
  },
  // Retrieves information about an account for the provided address.
getAccount: {
    path: '/aerium.Blockchain/GetAccount',
    requestStream: false,
    responseStream: false,
    requestType: blockchain_pb.GetAccountRequest,
    responseType: blockchain_pb.GetAccountResponse,
    requestSerialize: serialize_aerium_GetAccountRequest,
    requestDeserialize: deserialize_aerium_GetAccountRequest,
    responseSerialize: serialize_aerium_GetAccountResponse,
    responseDeserialize: deserialize_aerium_GetAccountResponse,
  },
  // Retrieves information about a validator for the provided address.
getValidator: {
    path: '/aerium.Blockchain/GetValidator',
    requestStream: false,
    responseStream: false,
    requestType: blockchain_pb.GetValidatorRequest,
    responseType: blockchain_pb.GetValidatorResponse,
    requestSerialize: serialize_aerium_GetValidatorRequest,
    requestDeserialize: deserialize_aerium_GetValidatorRequest,
    responseSerialize: serialize_aerium_GetValidatorResponse,
    responseDeserialize: deserialize_aerium_GetValidatorResponse,
  },
  // Retrieves information about a validator by its number.
getValidatorByNumber: {
    path: '/aerium.Blockchain/GetValidatorByNumber',
    requestStream: false,
    responseStream: false,
    requestType: blockchain_pb.GetValidatorByNumberRequest,
    responseType: blockchain_pb.GetValidatorResponse,
    requestSerialize: serialize_aerium_GetValidatorByNumberRequest,
    requestDeserialize: deserialize_aerium_GetValidatorByNumberRequest,
    responseSerialize: serialize_aerium_GetValidatorResponse,
    responseDeserialize: deserialize_aerium_GetValidatorResponse,
  },
  // Retrieves a list of all validator addresses.
getValidatorAddresses: {
    path: '/aerium.Blockchain/GetValidatorAddresses',
    requestStream: false,
    responseStream: false,
    requestType: blockchain_pb.GetValidatorAddressesRequest,
    responseType: blockchain_pb.GetValidatorAddressesResponse,
    requestSerialize: serialize_aerium_GetValidatorAddressesRequest,
    requestDeserialize: deserialize_aerium_GetValidatorAddressesRequest,
    responseSerialize: serialize_aerium_GetValidatorAddressesResponse,
    responseDeserialize: deserialize_aerium_GetValidatorAddressesResponse,
  },
  // Retrieves the public key of an account for the provided address.
getPublicKey: {
    path: '/aerium.Blockchain/GetPublicKey',
    requestStream: false,
    responseStream: false,
    requestType: blockchain_pb.GetPublicKeyRequest,
    responseType: blockchain_pb.GetPublicKeyResponse,
    requestSerialize: serialize_aerium_GetPublicKeyRequest,
    requestDeserialize: deserialize_aerium_GetPublicKeyRequest,
    responseSerialize: serialize_aerium_GetPublicKeyResponse,
    responseDeserialize: deserialize_aerium_GetPublicKeyResponse,
  },
  // Retrieves current transactions in the transaction pool.
getTxPoolContent: {
    path: '/aerium.Blockchain/GetTxPoolContent',
    requestStream: false,
    responseStream: false,
    requestType: blockchain_pb.GetTxPoolContentRequest,
    responseType: blockchain_pb.GetTxPoolContentResponse,
    requestSerialize: serialize_aerium_GetTxPoolContentRequest,
    requestDeserialize: deserialize_aerium_GetTxPoolContentRequest,
    responseSerialize: serialize_aerium_GetTxPoolContentResponse,
    responseDeserialize: deserialize_aerium_GetTxPoolContentResponse,
  },
};

exports.BlockchainClient = grpc.makeGenericClientConstructor(BlockchainService, 'Blockchain');
