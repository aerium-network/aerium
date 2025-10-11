// GENERATED CODE -- DO NOT EDIT!

'use strict';
var grpc = require('@grpc/grpc-js');
var transaction_pb = require('./transaction_pb.js');

function serialize_aerium_BroadcastTransactionRequest(arg) {
  if (!(arg instanceof transaction_pb.BroadcastTransactionRequest)) {
    throw new Error('Expected argument of type aerium.BroadcastTransactionRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_aerium_BroadcastTransactionRequest(buffer_arg) {
  return transaction_pb.BroadcastTransactionRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_aerium_BroadcastTransactionResponse(arg) {
  if (!(arg instanceof transaction_pb.BroadcastTransactionResponse)) {
    throw new Error('Expected argument of type aerium.BroadcastTransactionResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_aerium_BroadcastTransactionResponse(buffer_arg) {
  return transaction_pb.BroadcastTransactionResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_aerium_CalculateFeeRequest(arg) {
  if (!(arg instanceof transaction_pb.CalculateFeeRequest)) {
    throw new Error('Expected argument of type aerium.CalculateFeeRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_aerium_CalculateFeeRequest(buffer_arg) {
  return transaction_pb.CalculateFeeRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_aerium_CalculateFeeResponse(arg) {
  if (!(arg instanceof transaction_pb.CalculateFeeResponse)) {
    throw new Error('Expected argument of type aerium.CalculateFeeResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_aerium_CalculateFeeResponse(buffer_arg) {
  return transaction_pb.CalculateFeeResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_aerium_DecodeRawTransactionRequest(arg) {
  if (!(arg instanceof transaction_pb.DecodeRawTransactionRequest)) {
    throw new Error('Expected argument of type aerium.DecodeRawTransactionRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_aerium_DecodeRawTransactionRequest(buffer_arg) {
  return transaction_pb.DecodeRawTransactionRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_aerium_DecodeRawTransactionResponse(arg) {
  if (!(arg instanceof transaction_pb.DecodeRawTransactionResponse)) {
    throw new Error('Expected argument of type aerium.DecodeRawTransactionResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_aerium_DecodeRawTransactionResponse(buffer_arg) {
  return transaction_pb.DecodeRawTransactionResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_aerium_GetRawBatchTransferTransactionRequest(arg) {
  if (!(arg instanceof transaction_pb.GetRawBatchTransferTransactionRequest)) {
    throw new Error('Expected argument of type aerium.GetRawBatchTransferTransactionRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_aerium_GetRawBatchTransferTransactionRequest(buffer_arg) {
  return transaction_pb.GetRawBatchTransferTransactionRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_aerium_GetRawBondTransactionRequest(arg) {
  if (!(arg instanceof transaction_pb.GetRawBondTransactionRequest)) {
    throw new Error('Expected argument of type aerium.GetRawBondTransactionRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_aerium_GetRawBondTransactionRequest(buffer_arg) {
  return transaction_pb.GetRawBondTransactionRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_aerium_GetRawTransactionResponse(arg) {
  if (!(arg instanceof transaction_pb.GetRawTransactionResponse)) {
    throw new Error('Expected argument of type aerium.GetRawTransactionResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_aerium_GetRawTransactionResponse(buffer_arg) {
  return transaction_pb.GetRawTransactionResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_aerium_GetRawTransferTransactionRequest(arg) {
  if (!(arg instanceof transaction_pb.GetRawTransferTransactionRequest)) {
    throw new Error('Expected argument of type aerium.GetRawTransferTransactionRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_aerium_GetRawTransferTransactionRequest(buffer_arg) {
  return transaction_pb.GetRawTransferTransactionRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_aerium_GetRawUnbondTransactionRequest(arg) {
  if (!(arg instanceof transaction_pb.GetRawUnbondTransactionRequest)) {
    throw new Error('Expected argument of type aerium.GetRawUnbondTransactionRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_aerium_GetRawUnbondTransactionRequest(buffer_arg) {
  return transaction_pb.GetRawUnbondTransactionRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_aerium_GetRawWithdrawTransactionRequest(arg) {
  if (!(arg instanceof transaction_pb.GetRawWithdrawTransactionRequest)) {
    throw new Error('Expected argument of type aerium.GetRawWithdrawTransactionRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_aerium_GetRawWithdrawTransactionRequest(buffer_arg) {
  return transaction_pb.GetRawWithdrawTransactionRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_aerium_GetTransactionRequest(arg) {
  if (!(arg instanceof transaction_pb.GetTransactionRequest)) {
    throw new Error('Expected argument of type aerium.GetTransactionRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_aerium_GetTransactionRequest(buffer_arg) {
  return transaction_pb.GetTransactionRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_aerium_GetTransactionResponse(arg) {
  if (!(arg instanceof transaction_pb.GetTransactionResponse)) {
    throw new Error('Expected argument of type aerium.GetTransactionResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_aerium_GetTransactionResponse(buffer_arg) {
  return transaction_pb.GetTransactionResponse.deserializeBinary(new Uint8Array(buffer_arg));
}


// Transaction service defines various RPC methods for interacting with transactions.
var TransactionService = exports.TransactionService = {
  // GetTransaction retrieves transaction details based on the provided request parameters.
getTransaction: {
    path: '/aerium.Transaction/GetTransaction',
    requestStream: false,
    responseStream: false,
    requestType: transaction_pb.GetTransactionRequest,
    responseType: transaction_pb.GetTransactionResponse,
    requestSerialize: serialize_aerium_GetTransactionRequest,
    requestDeserialize: deserialize_aerium_GetTransactionRequest,
    responseSerialize: serialize_aerium_GetTransactionResponse,
    responseDeserialize: deserialize_aerium_GetTransactionResponse,
  },
  // CalculateFee calculates the transaction fee based on the specified amount and payload type.
calculateFee: {
    path: '/aerium.Transaction/CalculateFee',
    requestStream: false,
    responseStream: false,
    requestType: transaction_pb.CalculateFeeRequest,
    responseType: transaction_pb.CalculateFeeResponse,
    requestSerialize: serialize_aerium_CalculateFeeRequest,
    requestDeserialize: deserialize_aerium_CalculateFeeRequest,
    responseSerialize: serialize_aerium_CalculateFeeResponse,
    responseDeserialize: deserialize_aerium_CalculateFeeResponse,
  },
  // BroadcastTransaction broadcasts a signed transaction to the network.
broadcastTransaction: {
    path: '/aerium.Transaction/BroadcastTransaction',
    requestStream: false,
    responseStream: false,
    requestType: transaction_pb.BroadcastTransactionRequest,
    responseType: transaction_pb.BroadcastTransactionResponse,
    requestSerialize: serialize_aerium_BroadcastTransactionRequest,
    requestDeserialize: deserialize_aerium_BroadcastTransactionRequest,
    responseSerialize: serialize_aerium_BroadcastTransactionResponse,
    responseDeserialize: deserialize_aerium_BroadcastTransactionResponse,
  },
  // GetRawTransferTransaction retrieves raw details of a transfer transaction.
getRawTransferTransaction: {
    path: '/aerium.Transaction/GetRawTransferTransaction',
    requestStream: false,
    responseStream: false,
    requestType: transaction_pb.GetRawTransferTransactionRequest,
    responseType: transaction_pb.GetRawTransactionResponse,
    requestSerialize: serialize_aerium_GetRawTransferTransactionRequest,
    requestDeserialize: deserialize_aerium_GetRawTransferTransactionRequest,
    responseSerialize: serialize_aerium_GetRawTransactionResponse,
    responseDeserialize: deserialize_aerium_GetRawTransactionResponse,
  },
  // GetRawBondTransaction retrieves raw details of a bond transaction.
getRawBondTransaction: {
    path: '/aerium.Transaction/GetRawBondTransaction',
    requestStream: false,
    responseStream: false,
    requestType: transaction_pb.GetRawBondTransactionRequest,
    responseType: transaction_pb.GetRawTransactionResponse,
    requestSerialize: serialize_aerium_GetRawBondTransactionRequest,
    requestDeserialize: deserialize_aerium_GetRawBondTransactionRequest,
    responseSerialize: serialize_aerium_GetRawTransactionResponse,
    responseDeserialize: deserialize_aerium_GetRawTransactionResponse,
  },
  // GetRawUnbondTransaction retrieves raw details of an unbond transaction.
getRawUnbondTransaction: {
    path: '/aerium.Transaction/GetRawUnbondTransaction',
    requestStream: false,
    responseStream: false,
    requestType: transaction_pb.GetRawUnbondTransactionRequest,
    responseType: transaction_pb.GetRawTransactionResponse,
    requestSerialize: serialize_aerium_GetRawUnbondTransactionRequest,
    requestDeserialize: deserialize_aerium_GetRawUnbondTransactionRequest,
    responseSerialize: serialize_aerium_GetRawTransactionResponse,
    responseDeserialize: deserialize_aerium_GetRawTransactionResponse,
  },
  // GetRawWithdrawTransaction retrieves raw details of a withdraw transaction.
getRawWithdrawTransaction: {
    path: '/aerium.Transaction/GetRawWithdrawTransaction',
    requestStream: false,
    responseStream: false,
    requestType: transaction_pb.GetRawWithdrawTransactionRequest,
    responseType: transaction_pb.GetRawTransactionResponse,
    requestSerialize: serialize_aerium_GetRawWithdrawTransactionRequest,
    requestDeserialize: deserialize_aerium_GetRawWithdrawTransactionRequest,
    responseSerialize: serialize_aerium_GetRawTransactionResponse,
    responseDeserialize: deserialize_aerium_GetRawTransactionResponse,
  },
  // GetRawBatchTransferTransaction retrieves raw details of batch transfer transaction.
getRawBatchTransferTransaction: {
    path: '/aerium.Transaction/GetRawBatchTransferTransaction',
    requestStream: false,
    responseStream: false,
    requestType: transaction_pb.GetRawBatchTransferTransactionRequest,
    responseType: transaction_pb.GetRawTransactionResponse,
    requestSerialize: serialize_aerium_GetRawBatchTransferTransactionRequest,
    requestDeserialize: deserialize_aerium_GetRawBatchTransferTransactionRequest,
    responseSerialize: serialize_aerium_GetRawTransactionResponse,
    responseDeserialize: deserialize_aerium_GetRawTransactionResponse,
  },
  // DecodeRawTransaction accepts raw transaction and returns decoded transaction.
decodeRawTransaction: {
    path: '/aerium.Transaction/DecodeRawTransaction',
    requestStream: false,
    responseStream: false,
    requestType: transaction_pb.DecodeRawTransactionRequest,
    responseType: transaction_pb.DecodeRawTransactionResponse,
    requestSerialize: serialize_aerium_DecodeRawTransactionRequest,
    requestDeserialize: deserialize_aerium_DecodeRawTransactionRequest,
    responseSerialize: serialize_aerium_DecodeRawTransactionResponse,
    responseDeserialize: deserialize_aerium_DecodeRawTransactionResponse,
  },
};

exports.TransactionClient = grpc.makeGenericClientConstructor(TransactionService, 'Transaction');
