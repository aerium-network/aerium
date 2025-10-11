// GENERATED CODE -- DO NOT EDIT!

'use strict';
var grpc = require('@grpc/grpc-js');
var wallet_pb = require('./wallet_pb.js');

function serialize_aerium_CreateWalletRequest(arg) {
  if (!(arg instanceof wallet_pb.CreateWalletRequest)) {
    throw new Error('Expected argument of type aerium.CreateWalletRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_aerium_CreateWalletRequest(buffer_arg) {
  return wallet_pb.CreateWalletRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_aerium_CreateWalletResponse(arg) {
  if (!(arg instanceof wallet_pb.CreateWalletResponse)) {
    throw new Error('Expected argument of type aerium.CreateWalletResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_aerium_CreateWalletResponse(buffer_arg) {
  return wallet_pb.CreateWalletResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_aerium_GetAddressHistoryRequest(arg) {
  if (!(arg instanceof wallet_pb.GetAddressHistoryRequest)) {
    throw new Error('Expected argument of type aerium.GetAddressHistoryRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_aerium_GetAddressHistoryRequest(buffer_arg) {
  return wallet_pb.GetAddressHistoryRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_aerium_GetAddressHistoryResponse(arg) {
  if (!(arg instanceof wallet_pb.GetAddressHistoryResponse)) {
    throw new Error('Expected argument of type aerium.GetAddressHistoryResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_aerium_GetAddressHistoryResponse(buffer_arg) {
  return wallet_pb.GetAddressHistoryResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_aerium_GetAddressInfoRequest(arg) {
  if (!(arg instanceof wallet_pb.GetAddressInfoRequest)) {
    throw new Error('Expected argument of type aerium.GetAddressInfoRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_aerium_GetAddressInfoRequest(buffer_arg) {
  return wallet_pb.GetAddressInfoRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_aerium_GetAddressInfoResponse(arg) {
  if (!(arg instanceof wallet_pb.GetAddressInfoResponse)) {
    throw new Error('Expected argument of type aerium.GetAddressInfoResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_aerium_GetAddressInfoResponse(buffer_arg) {
  return wallet_pb.GetAddressInfoResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_aerium_GetNewAddressRequest(arg) {
  if (!(arg instanceof wallet_pb.GetNewAddressRequest)) {
    throw new Error('Expected argument of type aerium.GetNewAddressRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_aerium_GetNewAddressRequest(buffer_arg) {
  return wallet_pb.GetNewAddressRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_aerium_GetNewAddressResponse(arg) {
  if (!(arg instanceof wallet_pb.GetNewAddressResponse)) {
    throw new Error('Expected argument of type aerium.GetNewAddressResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_aerium_GetNewAddressResponse(buffer_arg) {
  return wallet_pb.GetNewAddressResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_aerium_GetTotalBalanceRequest(arg) {
  if (!(arg instanceof wallet_pb.GetTotalBalanceRequest)) {
    throw new Error('Expected argument of type aerium.GetTotalBalanceRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_aerium_GetTotalBalanceRequest(buffer_arg) {
  return wallet_pb.GetTotalBalanceRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_aerium_GetTotalBalanceResponse(arg) {
  if (!(arg instanceof wallet_pb.GetTotalBalanceResponse)) {
    throw new Error('Expected argument of type aerium.GetTotalBalanceResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_aerium_GetTotalBalanceResponse(buffer_arg) {
  return wallet_pb.GetTotalBalanceResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_aerium_GetTotalStakeRequest(arg) {
  if (!(arg instanceof wallet_pb.GetTotalStakeRequest)) {
    throw new Error('Expected argument of type aerium.GetTotalStakeRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_aerium_GetTotalStakeRequest(buffer_arg) {
  return wallet_pb.GetTotalStakeRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_aerium_GetTotalStakeResponse(arg) {
  if (!(arg instanceof wallet_pb.GetTotalStakeResponse)) {
    throw new Error('Expected argument of type aerium.GetTotalStakeResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_aerium_GetTotalStakeResponse(buffer_arg) {
  return wallet_pb.GetTotalStakeResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_aerium_GetValidatorAddressRequest(arg) {
  if (!(arg instanceof wallet_pb.GetValidatorAddressRequest)) {
    throw new Error('Expected argument of type aerium.GetValidatorAddressRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_aerium_GetValidatorAddressRequest(buffer_arg) {
  return wallet_pb.GetValidatorAddressRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_aerium_GetValidatorAddressResponse(arg) {
  if (!(arg instanceof wallet_pb.GetValidatorAddressResponse)) {
    throw new Error('Expected argument of type aerium.GetValidatorAddressResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_aerium_GetValidatorAddressResponse(buffer_arg) {
  return wallet_pb.GetValidatorAddressResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_aerium_GetWalletInfoRequest(arg) {
  if (!(arg instanceof wallet_pb.GetWalletInfoRequest)) {
    throw new Error('Expected argument of type aerium.GetWalletInfoRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_aerium_GetWalletInfoRequest(buffer_arg) {
  return wallet_pb.GetWalletInfoRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_aerium_GetWalletInfoResponse(arg) {
  if (!(arg instanceof wallet_pb.GetWalletInfoResponse)) {
    throw new Error('Expected argument of type aerium.GetWalletInfoResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_aerium_GetWalletInfoResponse(buffer_arg) {
  return wallet_pb.GetWalletInfoResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_aerium_ListAddressRequest(arg) {
  if (!(arg instanceof wallet_pb.ListAddressRequest)) {
    throw new Error('Expected argument of type aerium.ListAddressRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_aerium_ListAddressRequest(buffer_arg) {
  return wallet_pb.ListAddressRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_aerium_ListAddressResponse(arg) {
  if (!(arg instanceof wallet_pb.ListAddressResponse)) {
    throw new Error('Expected argument of type aerium.ListAddressResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_aerium_ListAddressResponse(buffer_arg) {
  return wallet_pb.ListAddressResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_aerium_ListWalletRequest(arg) {
  if (!(arg instanceof wallet_pb.ListWalletRequest)) {
    throw new Error('Expected argument of type aerium.ListWalletRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_aerium_ListWalletRequest(buffer_arg) {
  return wallet_pb.ListWalletRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_aerium_ListWalletResponse(arg) {
  if (!(arg instanceof wallet_pb.ListWalletResponse)) {
    throw new Error('Expected argument of type aerium.ListWalletResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_aerium_ListWalletResponse(buffer_arg) {
  return wallet_pb.ListWalletResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_aerium_LoadWalletRequest(arg) {
  if (!(arg instanceof wallet_pb.LoadWalletRequest)) {
    throw new Error('Expected argument of type aerium.LoadWalletRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_aerium_LoadWalletRequest(buffer_arg) {
  return wallet_pb.LoadWalletRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_aerium_LoadWalletResponse(arg) {
  if (!(arg instanceof wallet_pb.LoadWalletResponse)) {
    throw new Error('Expected argument of type aerium.LoadWalletResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_aerium_LoadWalletResponse(buffer_arg) {
  return wallet_pb.LoadWalletResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_aerium_RestoreWalletRequest(arg) {
  if (!(arg instanceof wallet_pb.RestoreWalletRequest)) {
    throw new Error('Expected argument of type aerium.RestoreWalletRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_aerium_RestoreWalletRequest(buffer_arg) {
  return wallet_pb.RestoreWalletRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_aerium_RestoreWalletResponse(arg) {
  if (!(arg instanceof wallet_pb.RestoreWalletResponse)) {
    throw new Error('Expected argument of type aerium.RestoreWalletResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_aerium_RestoreWalletResponse(buffer_arg) {
  return wallet_pb.RestoreWalletResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_aerium_SetAddressLabelRequest(arg) {
  if (!(arg instanceof wallet_pb.SetAddressLabelRequest)) {
    throw new Error('Expected argument of type aerium.SetAddressLabelRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_aerium_SetAddressLabelRequest(buffer_arg) {
  return wallet_pb.SetAddressLabelRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_aerium_SetAddressLabelResponse(arg) {
  if (!(arg instanceof wallet_pb.SetAddressLabelResponse)) {
    throw new Error('Expected argument of type aerium.SetAddressLabelResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_aerium_SetAddressLabelResponse(buffer_arg) {
  return wallet_pb.SetAddressLabelResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_aerium_SignMessageRequest(arg) {
  if (!(arg instanceof wallet_pb.SignMessageRequest)) {
    throw new Error('Expected argument of type aerium.SignMessageRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_aerium_SignMessageRequest(buffer_arg) {
  return wallet_pb.SignMessageRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_aerium_SignMessageResponse(arg) {
  if (!(arg instanceof wallet_pb.SignMessageResponse)) {
    throw new Error('Expected argument of type aerium.SignMessageResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_aerium_SignMessageResponse(buffer_arg) {
  return wallet_pb.SignMessageResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_aerium_SignRawTransactionRequest(arg) {
  if (!(arg instanceof wallet_pb.SignRawTransactionRequest)) {
    throw new Error('Expected argument of type aerium.SignRawTransactionRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_aerium_SignRawTransactionRequest(buffer_arg) {
  return wallet_pb.SignRawTransactionRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_aerium_SignRawTransactionResponse(arg) {
  if (!(arg instanceof wallet_pb.SignRawTransactionResponse)) {
    throw new Error('Expected argument of type aerium.SignRawTransactionResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_aerium_SignRawTransactionResponse(buffer_arg) {
  return wallet_pb.SignRawTransactionResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_aerium_UnloadWalletRequest(arg) {
  if (!(arg instanceof wallet_pb.UnloadWalletRequest)) {
    throw new Error('Expected argument of type aerium.UnloadWalletRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_aerium_UnloadWalletRequest(buffer_arg) {
  return wallet_pb.UnloadWalletRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_aerium_UnloadWalletResponse(arg) {
  if (!(arg instanceof wallet_pb.UnloadWalletResponse)) {
    throw new Error('Expected argument of type aerium.UnloadWalletResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_aerium_UnloadWalletResponse(buffer_arg) {
  return wallet_pb.UnloadWalletResponse.deserializeBinary(new Uint8Array(buffer_arg));
}


// Wallet service provides RPC methods for wallet management operations.
var WalletService = exports.WalletService = {
  // CreateWallet creates a new wallet with the specified parameters.
createWallet: {
    path: '/aerium.Wallet/CreateWallet',
    requestStream: false,
    responseStream: false,
    requestType: wallet_pb.CreateWalletRequest,
    responseType: wallet_pb.CreateWalletResponse,
    requestSerialize: serialize_aerium_CreateWalletRequest,
    requestDeserialize: deserialize_aerium_CreateWalletRequest,
    responseSerialize: serialize_aerium_CreateWalletResponse,
    responseDeserialize: deserialize_aerium_CreateWalletResponse,
  },
  // RestoreWallet restores an existing wallet with the given mnemonic.
restoreWallet: {
    path: '/aerium.Wallet/RestoreWallet',
    requestStream: false,
    responseStream: false,
    requestType: wallet_pb.RestoreWalletRequest,
    responseType: wallet_pb.RestoreWalletResponse,
    requestSerialize: serialize_aerium_RestoreWalletRequest,
    requestDeserialize: deserialize_aerium_RestoreWalletRequest,
    responseSerialize: serialize_aerium_RestoreWalletResponse,
    responseDeserialize: deserialize_aerium_RestoreWalletResponse,
  },
  // LoadWallet loads an existing wallet with the given name.
loadWallet: {
    path: '/aerium.Wallet/LoadWallet',
    requestStream: false,
    responseStream: false,
    requestType: wallet_pb.LoadWalletRequest,
    responseType: wallet_pb.LoadWalletResponse,
    requestSerialize: serialize_aerium_LoadWalletRequest,
    requestDeserialize: deserialize_aerium_LoadWalletRequest,
    responseSerialize: serialize_aerium_LoadWalletResponse,
    responseDeserialize: deserialize_aerium_LoadWalletResponse,
  },
  // UnloadWallet unloads a currently loaded wallet with the specified name.
unloadWallet: {
    path: '/aerium.Wallet/UnloadWallet',
    requestStream: false,
    responseStream: false,
    requestType: wallet_pb.UnloadWalletRequest,
    responseType: wallet_pb.UnloadWalletResponse,
    requestSerialize: serialize_aerium_UnloadWalletRequest,
    requestDeserialize: deserialize_aerium_UnloadWalletRequest,
    responseSerialize: serialize_aerium_UnloadWalletResponse,
    responseDeserialize: deserialize_aerium_UnloadWalletResponse,
  },
  // GetTotalBalance returns the total available balance of the wallet.
getTotalBalance: {
    path: '/aerium.Wallet/GetTotalBalance',
    requestStream: false,
    responseStream: false,
    requestType: wallet_pb.GetTotalBalanceRequest,
    responseType: wallet_pb.GetTotalBalanceResponse,
    requestSerialize: serialize_aerium_GetTotalBalanceRequest,
    requestDeserialize: deserialize_aerium_GetTotalBalanceRequest,
    responseSerialize: serialize_aerium_GetTotalBalanceResponse,
    responseDeserialize: deserialize_aerium_GetTotalBalanceResponse,
  },
  // SignRawTransaction signs a raw transaction for a specified wallet.
signRawTransaction: {
    path: '/aerium.Wallet/SignRawTransaction',
    requestStream: false,
    responseStream: false,
    requestType: wallet_pb.SignRawTransactionRequest,
    responseType: wallet_pb.SignRawTransactionResponse,
    requestSerialize: serialize_aerium_SignRawTransactionRequest,
    requestDeserialize: deserialize_aerium_SignRawTransactionRequest,
    responseSerialize: serialize_aerium_SignRawTransactionResponse,
    responseDeserialize: deserialize_aerium_SignRawTransactionResponse,
  },
  // GetValidatorAddress retrieves the validator address associated with a public key.
getValidatorAddress: {
    path: '/aerium.Wallet/GetValidatorAddress',
    requestStream: false,
    responseStream: false,
    requestType: wallet_pb.GetValidatorAddressRequest,
    responseType: wallet_pb.GetValidatorAddressResponse,
    requestSerialize: serialize_aerium_GetValidatorAddressRequest,
    requestDeserialize: deserialize_aerium_GetValidatorAddressRequest,
    responseSerialize: serialize_aerium_GetValidatorAddressResponse,
    responseDeserialize: deserialize_aerium_GetValidatorAddressResponse,
  },
  // GetNewAddress generates a new address for the specified wallet.
getNewAddress: {
    path: '/aerium.Wallet/GetNewAddress',
    requestStream: false,
    responseStream: false,
    requestType: wallet_pb.GetNewAddressRequest,
    responseType: wallet_pb.GetNewAddressResponse,
    requestSerialize: serialize_aerium_GetNewAddressRequest,
    requestDeserialize: deserialize_aerium_GetNewAddressRequest,
    responseSerialize: serialize_aerium_GetNewAddressResponse,
    responseDeserialize: deserialize_aerium_GetNewAddressResponse,
  },
  // GetAddressHistory retrieves the transaction history of an address.
getAddressHistory: {
    path: '/aerium.Wallet/GetAddressHistory',
    requestStream: false,
    responseStream: false,
    requestType: wallet_pb.GetAddressHistoryRequest,
    responseType: wallet_pb.GetAddressHistoryResponse,
    requestSerialize: serialize_aerium_GetAddressHistoryRequest,
    requestDeserialize: deserialize_aerium_GetAddressHistoryRequest,
    responseSerialize: serialize_aerium_GetAddressHistoryResponse,
    responseDeserialize: deserialize_aerium_GetAddressHistoryResponse,
  },
  // SignMessage signs an arbitrary message using a wallet's private key.
signMessage: {
    path: '/aerium.Wallet/SignMessage',
    requestStream: false,
    responseStream: false,
    requestType: wallet_pb.SignMessageRequest,
    responseType: wallet_pb.SignMessageResponse,
    requestSerialize: serialize_aerium_SignMessageRequest,
    requestDeserialize: deserialize_aerium_SignMessageRequest,
    responseSerialize: serialize_aerium_SignMessageResponse,
    responseDeserialize: deserialize_aerium_SignMessageResponse,
  },
  // GetTotalStake returns the total stake amount in the wallet.
getTotalStake: {
    path: '/aerium.Wallet/GetTotalStake',
    requestStream: false,
    responseStream: false,
    requestType: wallet_pb.GetTotalStakeRequest,
    responseType: wallet_pb.GetTotalStakeResponse,
    requestSerialize: serialize_aerium_GetTotalStakeRequest,
    requestDeserialize: deserialize_aerium_GetTotalStakeRequest,
    responseSerialize: serialize_aerium_GetTotalStakeResponse,
    responseDeserialize: deserialize_aerium_GetTotalStakeResponse,
  },
  // GetAddressInfo returns detailed information about a specific address.
getAddressInfo: {
    path: '/aerium.Wallet/GetAddressInfo',
    requestStream: false,
    responseStream: false,
    requestType: wallet_pb.GetAddressInfoRequest,
    responseType: wallet_pb.GetAddressInfoResponse,
    requestSerialize: serialize_aerium_GetAddressInfoRequest,
    requestDeserialize: deserialize_aerium_GetAddressInfoRequest,
    responseSerialize: serialize_aerium_GetAddressInfoResponse,
    responseDeserialize: deserialize_aerium_GetAddressInfoResponse,
  },
  // SetAddressLabel sets or updates the label for a given address.
setAddressLabel: {
    path: '/aerium.Wallet/SetAddressLabel',
    requestStream: false,
    responseStream: false,
    requestType: wallet_pb.SetAddressLabelRequest,
    responseType: wallet_pb.SetAddressLabelResponse,
    requestSerialize: serialize_aerium_SetAddressLabelRequest,
    requestDeserialize: deserialize_aerium_SetAddressLabelRequest,
    responseSerialize: serialize_aerium_SetAddressLabelResponse,
    responseDeserialize: deserialize_aerium_SetAddressLabelResponse,
  },
  // ListWallet returns list of all available wallets.
listWallet: {
    path: '/aerium.Wallet/ListWallet',
    requestStream: false,
    responseStream: false,
    requestType: wallet_pb.ListWalletRequest,
    responseType: wallet_pb.ListWalletResponse,
    requestSerialize: serialize_aerium_ListWalletRequest,
    requestDeserialize: deserialize_aerium_ListWalletRequest,
    responseSerialize: serialize_aerium_ListWalletResponse,
    responseDeserialize: deserialize_aerium_ListWalletResponse,
  },
  // GetWalletInfo returns detailed information about a specific wallet.
getWalletInfo: {
    path: '/aerium.Wallet/GetWalletInfo',
    requestStream: false,
    responseStream: false,
    requestType: wallet_pb.GetWalletInfoRequest,
    responseType: wallet_pb.GetWalletInfoResponse,
    requestSerialize: serialize_aerium_GetWalletInfoRequest,
    requestDeserialize: deserialize_aerium_GetWalletInfoRequest,
    responseSerialize: serialize_aerium_GetWalletInfoResponse,
    responseDeserialize: deserialize_aerium_GetWalletInfoResponse,
  },
  // ListAddress returns all addresses in the specified wallet.
listAddress: {
    path: '/aerium.Wallet/ListAddress',
    requestStream: false,
    responseStream: false,
    requestType: wallet_pb.ListAddressRequest,
    responseType: wallet_pb.ListAddressResponse,
    requestSerialize: serialize_aerium_ListAddressRequest,
    requestDeserialize: deserialize_aerium_ListAddressRequest,
    responseSerialize: serialize_aerium_ListAddressResponse,
    responseDeserialize: deserialize_aerium_ListAddressResponse,
  },
};

exports.WalletClient = grpc.makeGenericClientConstructor(WalletService, 'Wallet');
