package aerium;

import static io.grpc.MethodDescriptor.generateFullMethodName;

/**
 * <pre>
 * The Wallet service provides RPC methods for wallet management operations.
 * </pre>
 */
@javax.annotation.Generated(
    value = "by gRPC proto compiler (version 1.71.0)",
    comments = "Source: wallet.proto")
@io.grpc.stub.annotations.GrpcGenerated
public final class WalletGrpc {

  private WalletGrpc() {}

  public static final java.lang.String SERVICE_NAME = "aerium.Wallet";

  // Static method descriptors that strictly reflect the proto.
  private static volatile io.grpc.MethodDescriptor<aerium.WalletOuterClass.CreateWalletRequest,
      aerium.WalletOuterClass.CreateWalletResponse> getCreateWalletMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "CreateWallet",
      requestType = aerium.WalletOuterClass.CreateWalletRequest.class,
      responseType = aerium.WalletOuterClass.CreateWalletResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<aerium.WalletOuterClass.CreateWalletRequest,
      aerium.WalletOuterClass.CreateWalletResponse> getCreateWalletMethod() {
    io.grpc.MethodDescriptor<aerium.WalletOuterClass.CreateWalletRequest, aerium.WalletOuterClass.CreateWalletResponse> getCreateWalletMethod;
    if ((getCreateWalletMethod = WalletGrpc.getCreateWalletMethod) == null) {
      synchronized (WalletGrpc.class) {
        if ((getCreateWalletMethod = WalletGrpc.getCreateWalletMethod) == null) {
          WalletGrpc.getCreateWalletMethod = getCreateWalletMethod =
              io.grpc.MethodDescriptor.<aerium.WalletOuterClass.CreateWalletRequest, aerium.WalletOuterClass.CreateWalletResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "CreateWallet"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  aerium.WalletOuterClass.CreateWalletRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  aerium.WalletOuterClass.CreateWalletResponse.getDefaultInstance()))
              .setSchemaDescriptor(new WalletMethodDescriptorSupplier("CreateWallet"))
              .build();
        }
      }
    }
    return getCreateWalletMethod;
  }

  private static volatile io.grpc.MethodDescriptor<aerium.WalletOuterClass.RestoreWalletRequest,
      aerium.WalletOuterClass.RestoreWalletResponse> getRestoreWalletMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "RestoreWallet",
      requestType = aerium.WalletOuterClass.RestoreWalletRequest.class,
      responseType = aerium.WalletOuterClass.RestoreWalletResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<aerium.WalletOuterClass.RestoreWalletRequest,
      aerium.WalletOuterClass.RestoreWalletResponse> getRestoreWalletMethod() {
    io.grpc.MethodDescriptor<aerium.WalletOuterClass.RestoreWalletRequest, aerium.WalletOuterClass.RestoreWalletResponse> getRestoreWalletMethod;
    if ((getRestoreWalletMethod = WalletGrpc.getRestoreWalletMethod) == null) {
      synchronized (WalletGrpc.class) {
        if ((getRestoreWalletMethod = WalletGrpc.getRestoreWalletMethod) == null) {
          WalletGrpc.getRestoreWalletMethod = getRestoreWalletMethod =
              io.grpc.MethodDescriptor.<aerium.WalletOuterClass.RestoreWalletRequest, aerium.WalletOuterClass.RestoreWalletResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "RestoreWallet"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  aerium.WalletOuterClass.RestoreWalletRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  aerium.WalletOuterClass.RestoreWalletResponse.getDefaultInstance()))
              .setSchemaDescriptor(new WalletMethodDescriptorSupplier("RestoreWallet"))
              .build();
        }
      }
    }
    return getRestoreWalletMethod;
  }

  private static volatile io.grpc.MethodDescriptor<aerium.WalletOuterClass.LoadWalletRequest,
      aerium.WalletOuterClass.LoadWalletResponse> getLoadWalletMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "LoadWallet",
      requestType = aerium.WalletOuterClass.LoadWalletRequest.class,
      responseType = aerium.WalletOuterClass.LoadWalletResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<aerium.WalletOuterClass.LoadWalletRequest,
      aerium.WalletOuterClass.LoadWalletResponse> getLoadWalletMethod() {
    io.grpc.MethodDescriptor<aerium.WalletOuterClass.LoadWalletRequest, aerium.WalletOuterClass.LoadWalletResponse> getLoadWalletMethod;
    if ((getLoadWalletMethod = WalletGrpc.getLoadWalletMethod) == null) {
      synchronized (WalletGrpc.class) {
        if ((getLoadWalletMethod = WalletGrpc.getLoadWalletMethod) == null) {
          WalletGrpc.getLoadWalletMethod = getLoadWalletMethod =
              io.grpc.MethodDescriptor.<aerium.WalletOuterClass.LoadWalletRequest, aerium.WalletOuterClass.LoadWalletResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "LoadWallet"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  aerium.WalletOuterClass.LoadWalletRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  aerium.WalletOuterClass.LoadWalletResponse.getDefaultInstance()))
              .setSchemaDescriptor(new WalletMethodDescriptorSupplier("LoadWallet"))
              .build();
        }
      }
    }
    return getLoadWalletMethod;
  }

  private static volatile io.grpc.MethodDescriptor<aerium.WalletOuterClass.UnloadWalletRequest,
      aerium.WalletOuterClass.UnloadWalletResponse> getUnloadWalletMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "UnloadWallet",
      requestType = aerium.WalletOuterClass.UnloadWalletRequest.class,
      responseType = aerium.WalletOuterClass.UnloadWalletResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<aerium.WalletOuterClass.UnloadWalletRequest,
      aerium.WalletOuterClass.UnloadWalletResponse> getUnloadWalletMethod() {
    io.grpc.MethodDescriptor<aerium.WalletOuterClass.UnloadWalletRequest, aerium.WalletOuterClass.UnloadWalletResponse> getUnloadWalletMethod;
    if ((getUnloadWalletMethod = WalletGrpc.getUnloadWalletMethod) == null) {
      synchronized (WalletGrpc.class) {
        if ((getUnloadWalletMethod = WalletGrpc.getUnloadWalletMethod) == null) {
          WalletGrpc.getUnloadWalletMethod = getUnloadWalletMethod =
              io.grpc.MethodDescriptor.<aerium.WalletOuterClass.UnloadWalletRequest, aerium.WalletOuterClass.UnloadWalletResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "UnloadWallet"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  aerium.WalletOuterClass.UnloadWalletRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  aerium.WalletOuterClass.UnloadWalletResponse.getDefaultInstance()))
              .setSchemaDescriptor(new WalletMethodDescriptorSupplier("UnloadWallet"))
              .build();
        }
      }
    }
    return getUnloadWalletMethod;
  }

  private static volatile io.grpc.MethodDescriptor<aerium.WalletOuterClass.GetTotalBalanceRequest,
      aerium.WalletOuterClass.GetTotalBalanceResponse> getGetTotalBalanceMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "GetTotalBalance",
      requestType = aerium.WalletOuterClass.GetTotalBalanceRequest.class,
      responseType = aerium.WalletOuterClass.GetTotalBalanceResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<aerium.WalletOuterClass.GetTotalBalanceRequest,
      aerium.WalletOuterClass.GetTotalBalanceResponse> getGetTotalBalanceMethod() {
    io.grpc.MethodDescriptor<aerium.WalletOuterClass.GetTotalBalanceRequest, aerium.WalletOuterClass.GetTotalBalanceResponse> getGetTotalBalanceMethod;
    if ((getGetTotalBalanceMethod = WalletGrpc.getGetTotalBalanceMethod) == null) {
      synchronized (WalletGrpc.class) {
        if ((getGetTotalBalanceMethod = WalletGrpc.getGetTotalBalanceMethod) == null) {
          WalletGrpc.getGetTotalBalanceMethod = getGetTotalBalanceMethod =
              io.grpc.MethodDescriptor.<aerium.WalletOuterClass.GetTotalBalanceRequest, aerium.WalletOuterClass.GetTotalBalanceResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "GetTotalBalance"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  aerium.WalletOuterClass.GetTotalBalanceRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  aerium.WalletOuterClass.GetTotalBalanceResponse.getDefaultInstance()))
              .setSchemaDescriptor(new WalletMethodDescriptorSupplier("GetTotalBalance"))
              .build();
        }
      }
    }
    return getGetTotalBalanceMethod;
  }

  private static volatile io.grpc.MethodDescriptor<aerium.WalletOuterClass.SignRawTransactionRequest,
      aerium.WalletOuterClass.SignRawTransactionResponse> getSignRawTransactionMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "SignRawTransaction",
      requestType = aerium.WalletOuterClass.SignRawTransactionRequest.class,
      responseType = aerium.WalletOuterClass.SignRawTransactionResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<aerium.WalletOuterClass.SignRawTransactionRequest,
      aerium.WalletOuterClass.SignRawTransactionResponse> getSignRawTransactionMethod() {
    io.grpc.MethodDescriptor<aerium.WalletOuterClass.SignRawTransactionRequest, aerium.WalletOuterClass.SignRawTransactionResponse> getSignRawTransactionMethod;
    if ((getSignRawTransactionMethod = WalletGrpc.getSignRawTransactionMethod) == null) {
      synchronized (WalletGrpc.class) {
        if ((getSignRawTransactionMethod = WalletGrpc.getSignRawTransactionMethod) == null) {
          WalletGrpc.getSignRawTransactionMethod = getSignRawTransactionMethod =
              io.grpc.MethodDescriptor.<aerium.WalletOuterClass.SignRawTransactionRequest, aerium.WalletOuterClass.SignRawTransactionResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "SignRawTransaction"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  aerium.WalletOuterClass.SignRawTransactionRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  aerium.WalletOuterClass.SignRawTransactionResponse.getDefaultInstance()))
              .setSchemaDescriptor(new WalletMethodDescriptorSupplier("SignRawTransaction"))
              .build();
        }
      }
    }
    return getSignRawTransactionMethod;
  }

  private static volatile io.grpc.MethodDescriptor<aerium.WalletOuterClass.GetValidatorAddressRequest,
      aerium.WalletOuterClass.GetValidatorAddressResponse> getGetValidatorAddressMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "GetValidatorAddress",
      requestType = aerium.WalletOuterClass.GetValidatorAddressRequest.class,
      responseType = aerium.WalletOuterClass.GetValidatorAddressResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<aerium.WalletOuterClass.GetValidatorAddressRequest,
      aerium.WalletOuterClass.GetValidatorAddressResponse> getGetValidatorAddressMethod() {
    io.grpc.MethodDescriptor<aerium.WalletOuterClass.GetValidatorAddressRequest, aerium.WalletOuterClass.GetValidatorAddressResponse> getGetValidatorAddressMethod;
    if ((getGetValidatorAddressMethod = WalletGrpc.getGetValidatorAddressMethod) == null) {
      synchronized (WalletGrpc.class) {
        if ((getGetValidatorAddressMethod = WalletGrpc.getGetValidatorAddressMethod) == null) {
          WalletGrpc.getGetValidatorAddressMethod = getGetValidatorAddressMethod =
              io.grpc.MethodDescriptor.<aerium.WalletOuterClass.GetValidatorAddressRequest, aerium.WalletOuterClass.GetValidatorAddressResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "GetValidatorAddress"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  aerium.WalletOuterClass.GetValidatorAddressRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  aerium.WalletOuterClass.GetValidatorAddressResponse.getDefaultInstance()))
              .setSchemaDescriptor(new WalletMethodDescriptorSupplier("GetValidatorAddress"))
              .build();
        }
      }
    }
    return getGetValidatorAddressMethod;
  }

  private static volatile io.grpc.MethodDescriptor<aerium.WalletOuterClass.GetNewAddressRequest,
      aerium.WalletOuterClass.GetNewAddressResponse> getGetNewAddressMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "GetNewAddress",
      requestType = aerium.WalletOuterClass.GetNewAddressRequest.class,
      responseType = aerium.WalletOuterClass.GetNewAddressResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<aerium.WalletOuterClass.GetNewAddressRequest,
      aerium.WalletOuterClass.GetNewAddressResponse> getGetNewAddressMethod() {
    io.grpc.MethodDescriptor<aerium.WalletOuterClass.GetNewAddressRequest, aerium.WalletOuterClass.GetNewAddressResponse> getGetNewAddressMethod;
    if ((getGetNewAddressMethod = WalletGrpc.getGetNewAddressMethod) == null) {
      synchronized (WalletGrpc.class) {
        if ((getGetNewAddressMethod = WalletGrpc.getGetNewAddressMethod) == null) {
          WalletGrpc.getGetNewAddressMethod = getGetNewAddressMethod =
              io.grpc.MethodDescriptor.<aerium.WalletOuterClass.GetNewAddressRequest, aerium.WalletOuterClass.GetNewAddressResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "GetNewAddress"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  aerium.WalletOuterClass.GetNewAddressRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  aerium.WalletOuterClass.GetNewAddressResponse.getDefaultInstance()))
              .setSchemaDescriptor(new WalletMethodDescriptorSupplier("GetNewAddress"))
              .build();
        }
      }
    }
    return getGetNewAddressMethod;
  }

  private static volatile io.grpc.MethodDescriptor<aerium.WalletOuterClass.GetAddressHistoryRequest,
      aerium.WalletOuterClass.GetAddressHistoryResponse> getGetAddressHistoryMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "GetAddressHistory",
      requestType = aerium.WalletOuterClass.GetAddressHistoryRequest.class,
      responseType = aerium.WalletOuterClass.GetAddressHistoryResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<aerium.WalletOuterClass.GetAddressHistoryRequest,
      aerium.WalletOuterClass.GetAddressHistoryResponse> getGetAddressHistoryMethod() {
    io.grpc.MethodDescriptor<aerium.WalletOuterClass.GetAddressHistoryRequest, aerium.WalletOuterClass.GetAddressHistoryResponse> getGetAddressHistoryMethod;
    if ((getGetAddressHistoryMethod = WalletGrpc.getGetAddressHistoryMethod) == null) {
      synchronized (WalletGrpc.class) {
        if ((getGetAddressHistoryMethod = WalletGrpc.getGetAddressHistoryMethod) == null) {
          WalletGrpc.getGetAddressHistoryMethod = getGetAddressHistoryMethod =
              io.grpc.MethodDescriptor.<aerium.WalletOuterClass.GetAddressHistoryRequest, aerium.WalletOuterClass.GetAddressHistoryResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "GetAddressHistory"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  aerium.WalletOuterClass.GetAddressHistoryRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  aerium.WalletOuterClass.GetAddressHistoryResponse.getDefaultInstance()))
              .setSchemaDescriptor(new WalletMethodDescriptorSupplier("GetAddressHistory"))
              .build();
        }
      }
    }
    return getGetAddressHistoryMethod;
  }

  private static volatile io.grpc.MethodDescriptor<aerium.WalletOuterClass.SignMessageRequest,
      aerium.WalletOuterClass.SignMessageResponse> getSignMessageMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "SignMessage",
      requestType = aerium.WalletOuterClass.SignMessageRequest.class,
      responseType = aerium.WalletOuterClass.SignMessageResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<aerium.WalletOuterClass.SignMessageRequest,
      aerium.WalletOuterClass.SignMessageResponse> getSignMessageMethod() {
    io.grpc.MethodDescriptor<aerium.WalletOuterClass.SignMessageRequest, aerium.WalletOuterClass.SignMessageResponse> getSignMessageMethod;
    if ((getSignMessageMethod = WalletGrpc.getSignMessageMethod) == null) {
      synchronized (WalletGrpc.class) {
        if ((getSignMessageMethod = WalletGrpc.getSignMessageMethod) == null) {
          WalletGrpc.getSignMessageMethod = getSignMessageMethod =
              io.grpc.MethodDescriptor.<aerium.WalletOuterClass.SignMessageRequest, aerium.WalletOuterClass.SignMessageResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "SignMessage"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  aerium.WalletOuterClass.SignMessageRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  aerium.WalletOuterClass.SignMessageResponse.getDefaultInstance()))
              .setSchemaDescriptor(new WalletMethodDescriptorSupplier("SignMessage"))
              .build();
        }
      }
    }
    return getSignMessageMethod;
  }

  private static volatile io.grpc.MethodDescriptor<aerium.WalletOuterClass.GetTotalStakeRequest,
      aerium.WalletOuterClass.GetTotalStakeResponse> getGetTotalStakeMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "GetTotalStake",
      requestType = aerium.WalletOuterClass.GetTotalStakeRequest.class,
      responseType = aerium.WalletOuterClass.GetTotalStakeResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<aerium.WalletOuterClass.GetTotalStakeRequest,
      aerium.WalletOuterClass.GetTotalStakeResponse> getGetTotalStakeMethod() {
    io.grpc.MethodDescriptor<aerium.WalletOuterClass.GetTotalStakeRequest, aerium.WalletOuterClass.GetTotalStakeResponse> getGetTotalStakeMethod;
    if ((getGetTotalStakeMethod = WalletGrpc.getGetTotalStakeMethod) == null) {
      synchronized (WalletGrpc.class) {
        if ((getGetTotalStakeMethod = WalletGrpc.getGetTotalStakeMethod) == null) {
          WalletGrpc.getGetTotalStakeMethod = getGetTotalStakeMethod =
              io.grpc.MethodDescriptor.<aerium.WalletOuterClass.GetTotalStakeRequest, aerium.WalletOuterClass.GetTotalStakeResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "GetTotalStake"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  aerium.WalletOuterClass.GetTotalStakeRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  aerium.WalletOuterClass.GetTotalStakeResponse.getDefaultInstance()))
              .setSchemaDescriptor(new WalletMethodDescriptorSupplier("GetTotalStake"))
              .build();
        }
      }
    }
    return getGetTotalStakeMethod;
  }

  private static volatile io.grpc.MethodDescriptor<aerium.WalletOuterClass.GetAddressInfoRequest,
      aerium.WalletOuterClass.GetAddressInfoResponse> getGetAddressInfoMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "GetAddressInfo",
      requestType = aerium.WalletOuterClass.GetAddressInfoRequest.class,
      responseType = aerium.WalletOuterClass.GetAddressInfoResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<aerium.WalletOuterClass.GetAddressInfoRequest,
      aerium.WalletOuterClass.GetAddressInfoResponse> getGetAddressInfoMethod() {
    io.grpc.MethodDescriptor<aerium.WalletOuterClass.GetAddressInfoRequest, aerium.WalletOuterClass.GetAddressInfoResponse> getGetAddressInfoMethod;
    if ((getGetAddressInfoMethod = WalletGrpc.getGetAddressInfoMethod) == null) {
      synchronized (WalletGrpc.class) {
        if ((getGetAddressInfoMethod = WalletGrpc.getGetAddressInfoMethod) == null) {
          WalletGrpc.getGetAddressInfoMethod = getGetAddressInfoMethod =
              io.grpc.MethodDescriptor.<aerium.WalletOuterClass.GetAddressInfoRequest, aerium.WalletOuterClass.GetAddressInfoResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "GetAddressInfo"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  aerium.WalletOuterClass.GetAddressInfoRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  aerium.WalletOuterClass.GetAddressInfoResponse.getDefaultInstance()))
              .setSchemaDescriptor(new WalletMethodDescriptorSupplier("GetAddressInfo"))
              .build();
        }
      }
    }
    return getGetAddressInfoMethod;
  }

  private static volatile io.grpc.MethodDescriptor<aerium.WalletOuterClass.SetAddressLabelRequest,
      aerium.WalletOuterClass.SetAddressLabelResponse> getSetAddressLabelMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "SetAddressLabel",
      requestType = aerium.WalletOuterClass.SetAddressLabelRequest.class,
      responseType = aerium.WalletOuterClass.SetAddressLabelResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<aerium.WalletOuterClass.SetAddressLabelRequest,
      aerium.WalletOuterClass.SetAddressLabelResponse> getSetAddressLabelMethod() {
    io.grpc.MethodDescriptor<aerium.WalletOuterClass.SetAddressLabelRequest, aerium.WalletOuterClass.SetAddressLabelResponse> getSetAddressLabelMethod;
    if ((getSetAddressLabelMethod = WalletGrpc.getSetAddressLabelMethod) == null) {
      synchronized (WalletGrpc.class) {
        if ((getSetAddressLabelMethod = WalletGrpc.getSetAddressLabelMethod) == null) {
          WalletGrpc.getSetAddressLabelMethod = getSetAddressLabelMethod =
              io.grpc.MethodDescriptor.<aerium.WalletOuterClass.SetAddressLabelRequest, aerium.WalletOuterClass.SetAddressLabelResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "SetAddressLabel"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  aerium.WalletOuterClass.SetAddressLabelRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  aerium.WalletOuterClass.SetAddressLabelResponse.getDefaultInstance()))
              .setSchemaDescriptor(new WalletMethodDescriptorSupplier("SetAddressLabel"))
              .build();
        }
      }
    }
    return getSetAddressLabelMethod;
  }

  private static volatile io.grpc.MethodDescriptor<aerium.WalletOuterClass.ListWalletRequest,
      aerium.WalletOuterClass.ListWalletResponse> getListWalletMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "ListWallet",
      requestType = aerium.WalletOuterClass.ListWalletRequest.class,
      responseType = aerium.WalletOuterClass.ListWalletResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<aerium.WalletOuterClass.ListWalletRequest,
      aerium.WalletOuterClass.ListWalletResponse> getListWalletMethod() {
    io.grpc.MethodDescriptor<aerium.WalletOuterClass.ListWalletRequest, aerium.WalletOuterClass.ListWalletResponse> getListWalletMethod;
    if ((getListWalletMethod = WalletGrpc.getListWalletMethod) == null) {
      synchronized (WalletGrpc.class) {
        if ((getListWalletMethod = WalletGrpc.getListWalletMethod) == null) {
          WalletGrpc.getListWalletMethod = getListWalletMethod =
              io.grpc.MethodDescriptor.<aerium.WalletOuterClass.ListWalletRequest, aerium.WalletOuterClass.ListWalletResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "ListWallet"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  aerium.WalletOuterClass.ListWalletRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  aerium.WalletOuterClass.ListWalletResponse.getDefaultInstance()))
              .setSchemaDescriptor(new WalletMethodDescriptorSupplier("ListWallet"))
              .build();
        }
      }
    }
    return getListWalletMethod;
  }

  private static volatile io.grpc.MethodDescriptor<aerium.WalletOuterClass.GetWalletInfoRequest,
      aerium.WalletOuterClass.GetWalletInfoResponse> getGetWalletInfoMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "GetWalletInfo",
      requestType = aerium.WalletOuterClass.GetWalletInfoRequest.class,
      responseType = aerium.WalletOuterClass.GetWalletInfoResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<aerium.WalletOuterClass.GetWalletInfoRequest,
      aerium.WalletOuterClass.GetWalletInfoResponse> getGetWalletInfoMethod() {
    io.grpc.MethodDescriptor<aerium.WalletOuterClass.GetWalletInfoRequest, aerium.WalletOuterClass.GetWalletInfoResponse> getGetWalletInfoMethod;
    if ((getGetWalletInfoMethod = WalletGrpc.getGetWalletInfoMethod) == null) {
      synchronized (WalletGrpc.class) {
        if ((getGetWalletInfoMethod = WalletGrpc.getGetWalletInfoMethod) == null) {
          WalletGrpc.getGetWalletInfoMethod = getGetWalletInfoMethod =
              io.grpc.MethodDescriptor.<aerium.WalletOuterClass.GetWalletInfoRequest, aerium.WalletOuterClass.GetWalletInfoResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "GetWalletInfo"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  aerium.WalletOuterClass.GetWalletInfoRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  aerium.WalletOuterClass.GetWalletInfoResponse.getDefaultInstance()))
              .setSchemaDescriptor(new WalletMethodDescriptorSupplier("GetWalletInfo"))
              .build();
        }
      }
    }
    return getGetWalletInfoMethod;
  }

  private static volatile io.grpc.MethodDescriptor<aerium.WalletOuterClass.ListAddressRequest,
      aerium.WalletOuterClass.ListAddressResponse> getListAddressMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "ListAddress",
      requestType = aerium.WalletOuterClass.ListAddressRequest.class,
      responseType = aerium.WalletOuterClass.ListAddressResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<aerium.WalletOuterClass.ListAddressRequest,
      aerium.WalletOuterClass.ListAddressResponse> getListAddressMethod() {
    io.grpc.MethodDescriptor<aerium.WalletOuterClass.ListAddressRequest, aerium.WalletOuterClass.ListAddressResponse> getListAddressMethod;
    if ((getListAddressMethod = WalletGrpc.getListAddressMethod) == null) {
      synchronized (WalletGrpc.class) {
        if ((getListAddressMethod = WalletGrpc.getListAddressMethod) == null) {
          WalletGrpc.getListAddressMethod = getListAddressMethod =
              io.grpc.MethodDescriptor.<aerium.WalletOuterClass.ListAddressRequest, aerium.WalletOuterClass.ListAddressResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "ListAddress"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  aerium.WalletOuterClass.ListAddressRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  aerium.WalletOuterClass.ListAddressResponse.getDefaultInstance()))
              .setSchemaDescriptor(new WalletMethodDescriptorSupplier("ListAddress"))
              .build();
        }
      }
    }
    return getListAddressMethod;
  }

  /**
   * Creates a new async stub that supports all call types for the service
   */
  public static WalletStub newStub(io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<WalletStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<WalletStub>() {
        @java.lang.Override
        public WalletStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new WalletStub(channel, callOptions);
        }
      };
    return WalletStub.newStub(factory, channel);
  }

  /**
   * Creates a new blocking-style stub that supports all types of calls on the service
   */
  public static WalletBlockingV2Stub newBlockingV2Stub(
      io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<WalletBlockingV2Stub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<WalletBlockingV2Stub>() {
        @java.lang.Override
        public WalletBlockingV2Stub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new WalletBlockingV2Stub(channel, callOptions);
        }
      };
    return WalletBlockingV2Stub.newStub(factory, channel);
  }

  /**
   * Creates a new blocking-style stub that supports unary and streaming output calls on the service
   */
  public static WalletBlockingStub newBlockingStub(
      io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<WalletBlockingStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<WalletBlockingStub>() {
        @java.lang.Override
        public WalletBlockingStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new WalletBlockingStub(channel, callOptions);
        }
      };
    return WalletBlockingStub.newStub(factory, channel);
  }

  /**
   * Creates a new ListenableFuture-style stub that supports unary calls on the service
   */
  public static WalletFutureStub newFutureStub(
      io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<WalletFutureStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<WalletFutureStub>() {
        @java.lang.Override
        public WalletFutureStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new WalletFutureStub(channel, callOptions);
        }
      };
    return WalletFutureStub.newStub(factory, channel);
  }

  /**
   * <pre>
   * The Wallet service provides RPC methods for wallet management operations.
   * </pre>
   */
  public interface AsyncService {

    /**
     * <pre>
     * Creates a new wallet with the specified parameters.
     * </pre>
     */
    default void createWallet(aerium.WalletOuterClass.CreateWalletRequest request,
        io.grpc.stub.StreamObserver<aerium.WalletOuterClass.CreateWalletResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getCreateWalletMethod(), responseObserver);
    }

    /**
     * <pre>
     * Restores an existing wallet with the given mnemonic.
     * </pre>
     */
    default void restoreWallet(aerium.WalletOuterClass.RestoreWalletRequest request,
        io.grpc.stub.StreamObserver<aerium.WalletOuterClass.RestoreWalletResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getRestoreWalletMethod(), responseObserver);
    }

    /**
     * <pre>
     * Loads an existing wallet with the given name.
     * </pre>
     */
    default void loadWallet(aerium.WalletOuterClass.LoadWalletRequest request,
        io.grpc.stub.StreamObserver<aerium.WalletOuterClass.LoadWalletResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getLoadWalletMethod(), responseObserver);
    }

    /**
     * <pre>
     * Unloads a currently loaded wallet with the specified name.
     * </pre>
     */
    default void unloadWallet(aerium.WalletOuterClass.UnloadWalletRequest request,
        io.grpc.stub.StreamObserver<aerium.WalletOuterClass.UnloadWalletResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getUnloadWalletMethod(), responseObserver);
    }

    /**
     * <pre>
     * Returns the total available balance of the wallet.
     * </pre>
     */
    default void getTotalBalance(aerium.WalletOuterClass.GetTotalBalanceRequest request,
        io.grpc.stub.StreamObserver<aerium.WalletOuterClass.GetTotalBalanceResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getGetTotalBalanceMethod(), responseObserver);
    }

    /**
     * <pre>
     * Signs a raw transaction for a specified wallet.
     * </pre>
     */
    default void signRawTransaction(aerium.WalletOuterClass.SignRawTransactionRequest request,
        io.grpc.stub.StreamObserver<aerium.WalletOuterClass.SignRawTransactionResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getSignRawTransactionMethod(), responseObserver);
    }

    /**
     * <pre>
     * Retrieves the validator address associated with a public key.
     * </pre>
     */
    default void getValidatorAddress(aerium.WalletOuterClass.GetValidatorAddressRequest request,
        io.grpc.stub.StreamObserver<aerium.WalletOuterClass.GetValidatorAddressResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getGetValidatorAddressMethod(), responseObserver);
    }

    /**
     * <pre>
     * Generates a new address for the specified wallet.
     * </pre>
     */
    default void getNewAddress(aerium.WalletOuterClass.GetNewAddressRequest request,
        io.grpc.stub.StreamObserver<aerium.WalletOuterClass.GetNewAddressResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getGetNewAddressMethod(), responseObserver);
    }

    /**
     * <pre>
     * Retrieves the transaction history of an address.
     * </pre>
     */
    default void getAddressHistory(aerium.WalletOuterClass.GetAddressHistoryRequest request,
        io.grpc.stub.StreamObserver<aerium.WalletOuterClass.GetAddressHistoryResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getGetAddressHistoryMethod(), responseObserver);
    }

    /**
     * <pre>
     * Signs an arbitrary message using a wallet's private key.
     * </pre>
     */
    default void signMessage(aerium.WalletOuterClass.SignMessageRequest request,
        io.grpc.stub.StreamObserver<aerium.WalletOuterClass.SignMessageResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getSignMessageMethod(), responseObserver);
    }

    /**
     * <pre>
     * Returns the total stake amount in the wallet.
     * </pre>
     */
    default void getTotalStake(aerium.WalletOuterClass.GetTotalStakeRequest request,
        io.grpc.stub.StreamObserver<aerium.WalletOuterClass.GetTotalStakeResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getGetTotalStakeMethod(), responseObserver);
    }

    /**
     * <pre>
     * Returns detailed information about a specific address.
     * </pre>
     */
    default void getAddressInfo(aerium.WalletOuterClass.GetAddressInfoRequest request,
        io.grpc.stub.StreamObserver<aerium.WalletOuterClass.GetAddressInfoResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getGetAddressInfoMethod(), responseObserver);
    }

    /**
     * <pre>
     * Sets or updates the label for a given address.
     * </pre>
     */
    default void setAddressLabel(aerium.WalletOuterClass.SetAddressLabelRequest request,
        io.grpc.stub.StreamObserver<aerium.WalletOuterClass.SetAddressLabelResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getSetAddressLabelMethod(), responseObserver);
    }

    /**
     * <pre>
     * Returns a list of all available wallets.
     * </pre>
     */
    default void listWallet(aerium.WalletOuterClass.ListWalletRequest request,
        io.grpc.stub.StreamObserver<aerium.WalletOuterClass.ListWalletResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getListWalletMethod(), responseObserver);
    }

    /**
     * <pre>
     * Returns detailed information about a specific wallet.
     * </pre>
     */
    default void getWalletInfo(aerium.WalletOuterClass.GetWalletInfoRequest request,
        io.grpc.stub.StreamObserver<aerium.WalletOuterClass.GetWalletInfoResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getGetWalletInfoMethod(), responseObserver);
    }

    /**
     * <pre>
     * Returns all addresses in the specified wallet.
     * </pre>
     */
    default void listAddress(aerium.WalletOuterClass.ListAddressRequest request,
        io.grpc.stub.StreamObserver<aerium.WalletOuterClass.ListAddressResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getListAddressMethod(), responseObserver);
    }
  }

  /**
   * Base class for the server implementation of the service Wallet.
   * <pre>
   * The Wallet service provides RPC methods for wallet management operations.
   * </pre>
   */
  public static abstract class WalletImplBase
      implements io.grpc.BindableService, AsyncService {

    @java.lang.Override public final io.grpc.ServerServiceDefinition bindService() {
      return WalletGrpc.bindService(this);
    }
  }

  /**
   * A stub to allow clients to do asynchronous rpc calls to service Wallet.
   * <pre>
   * The Wallet service provides RPC methods for wallet management operations.
   * </pre>
   */
  public static final class WalletStub
      extends io.grpc.stub.AbstractAsyncStub<WalletStub> {
    private WalletStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected WalletStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new WalletStub(channel, callOptions);
    }

    /**
     * <pre>
     * Creates a new wallet with the specified parameters.
     * </pre>
     */
    public void createWallet(aerium.WalletOuterClass.CreateWalletRequest request,
        io.grpc.stub.StreamObserver<aerium.WalletOuterClass.CreateWalletResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getCreateWalletMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     * <pre>
     * Restores an existing wallet with the given mnemonic.
     * </pre>
     */
    public void restoreWallet(aerium.WalletOuterClass.RestoreWalletRequest request,
        io.grpc.stub.StreamObserver<aerium.WalletOuterClass.RestoreWalletResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getRestoreWalletMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     * <pre>
     * Loads an existing wallet with the given name.
     * </pre>
     */
    public void loadWallet(aerium.WalletOuterClass.LoadWalletRequest request,
        io.grpc.stub.StreamObserver<aerium.WalletOuterClass.LoadWalletResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getLoadWalletMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     * <pre>
     * Unloads a currently loaded wallet with the specified name.
     * </pre>
     */
    public void unloadWallet(aerium.WalletOuterClass.UnloadWalletRequest request,
        io.grpc.stub.StreamObserver<aerium.WalletOuterClass.UnloadWalletResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getUnloadWalletMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     * <pre>
     * Returns the total available balance of the wallet.
     * </pre>
     */
    public void getTotalBalance(aerium.WalletOuterClass.GetTotalBalanceRequest request,
        io.grpc.stub.StreamObserver<aerium.WalletOuterClass.GetTotalBalanceResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getGetTotalBalanceMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     * <pre>
     * Signs a raw transaction for a specified wallet.
     * </pre>
     */
    public void signRawTransaction(aerium.WalletOuterClass.SignRawTransactionRequest request,
        io.grpc.stub.StreamObserver<aerium.WalletOuterClass.SignRawTransactionResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getSignRawTransactionMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     * <pre>
     * Retrieves the validator address associated with a public key.
     * </pre>
     */
    public void getValidatorAddress(aerium.WalletOuterClass.GetValidatorAddressRequest request,
        io.grpc.stub.StreamObserver<aerium.WalletOuterClass.GetValidatorAddressResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getGetValidatorAddressMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     * <pre>
     * Generates a new address for the specified wallet.
     * </pre>
     */
    public void getNewAddress(aerium.WalletOuterClass.GetNewAddressRequest request,
        io.grpc.stub.StreamObserver<aerium.WalletOuterClass.GetNewAddressResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getGetNewAddressMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     * <pre>
     * Retrieves the transaction history of an address.
     * </pre>
     */
    public void getAddressHistory(aerium.WalletOuterClass.GetAddressHistoryRequest request,
        io.grpc.stub.StreamObserver<aerium.WalletOuterClass.GetAddressHistoryResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getGetAddressHistoryMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     * <pre>
     * Signs an arbitrary message using a wallet's private key.
     * </pre>
     */
    public void signMessage(aerium.WalletOuterClass.SignMessageRequest request,
        io.grpc.stub.StreamObserver<aerium.WalletOuterClass.SignMessageResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getSignMessageMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     * <pre>
     * Returns the total stake amount in the wallet.
     * </pre>
     */
    public void getTotalStake(aerium.WalletOuterClass.GetTotalStakeRequest request,
        io.grpc.stub.StreamObserver<aerium.WalletOuterClass.GetTotalStakeResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getGetTotalStakeMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     * <pre>
     * Returns detailed information about a specific address.
     * </pre>
     */
    public void getAddressInfo(aerium.WalletOuterClass.GetAddressInfoRequest request,
        io.grpc.stub.StreamObserver<aerium.WalletOuterClass.GetAddressInfoResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getGetAddressInfoMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     * <pre>
     * Sets or updates the label for a given address.
     * </pre>
     */
    public void setAddressLabel(aerium.WalletOuterClass.SetAddressLabelRequest request,
        io.grpc.stub.StreamObserver<aerium.WalletOuterClass.SetAddressLabelResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getSetAddressLabelMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     * <pre>
     * Returns a list of all available wallets.
     * </pre>
     */
    public void listWallet(aerium.WalletOuterClass.ListWalletRequest request,
        io.grpc.stub.StreamObserver<aerium.WalletOuterClass.ListWalletResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getListWalletMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     * <pre>
     * Returns detailed information about a specific wallet.
     * </pre>
     */
    public void getWalletInfo(aerium.WalletOuterClass.GetWalletInfoRequest request,
        io.grpc.stub.StreamObserver<aerium.WalletOuterClass.GetWalletInfoResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getGetWalletInfoMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     * <pre>
     * Returns all addresses in the specified wallet.
     * </pre>
     */
    public void listAddress(aerium.WalletOuterClass.ListAddressRequest request,
        io.grpc.stub.StreamObserver<aerium.WalletOuterClass.ListAddressResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getListAddressMethod(), getCallOptions()), request, responseObserver);
    }
  }

  /**
   * A stub to allow clients to do synchronous rpc calls to service Wallet.
   * <pre>
   * The Wallet service provides RPC methods for wallet management operations.
   * </pre>
   */
  public static final class WalletBlockingV2Stub
      extends io.grpc.stub.AbstractBlockingStub<WalletBlockingV2Stub> {
    private WalletBlockingV2Stub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected WalletBlockingV2Stub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new WalletBlockingV2Stub(channel, callOptions);
    }

    /**
     * <pre>
     * Creates a new wallet with the specified parameters.
     * </pre>
     */
    public aerium.WalletOuterClass.CreateWalletResponse createWallet(aerium.WalletOuterClass.CreateWalletRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getCreateWalletMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * Restores an existing wallet with the given mnemonic.
     * </pre>
     */
    public aerium.WalletOuterClass.RestoreWalletResponse restoreWallet(aerium.WalletOuterClass.RestoreWalletRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getRestoreWalletMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * Loads an existing wallet with the given name.
     * </pre>
     */
    public aerium.WalletOuterClass.LoadWalletResponse loadWallet(aerium.WalletOuterClass.LoadWalletRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getLoadWalletMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * Unloads a currently loaded wallet with the specified name.
     * </pre>
     */
    public aerium.WalletOuterClass.UnloadWalletResponse unloadWallet(aerium.WalletOuterClass.UnloadWalletRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getUnloadWalletMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * Returns the total available balance of the wallet.
     * </pre>
     */
    public aerium.WalletOuterClass.GetTotalBalanceResponse getTotalBalance(aerium.WalletOuterClass.GetTotalBalanceRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getGetTotalBalanceMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * Signs a raw transaction for a specified wallet.
     * </pre>
     */
    public aerium.WalletOuterClass.SignRawTransactionResponse signRawTransaction(aerium.WalletOuterClass.SignRawTransactionRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getSignRawTransactionMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * Retrieves the validator address associated with a public key.
     * </pre>
     */
    public aerium.WalletOuterClass.GetValidatorAddressResponse getValidatorAddress(aerium.WalletOuterClass.GetValidatorAddressRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getGetValidatorAddressMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * Generates a new address for the specified wallet.
     * </pre>
     */
    public aerium.WalletOuterClass.GetNewAddressResponse getNewAddress(aerium.WalletOuterClass.GetNewAddressRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getGetNewAddressMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * Retrieves the transaction history of an address.
     * </pre>
     */
    public aerium.WalletOuterClass.GetAddressHistoryResponse getAddressHistory(aerium.WalletOuterClass.GetAddressHistoryRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getGetAddressHistoryMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * Signs an arbitrary message using a wallet's private key.
     * </pre>
     */
    public aerium.WalletOuterClass.SignMessageResponse signMessage(aerium.WalletOuterClass.SignMessageRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getSignMessageMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * Returns the total stake amount in the wallet.
     * </pre>
     */
    public aerium.WalletOuterClass.GetTotalStakeResponse getTotalStake(aerium.WalletOuterClass.GetTotalStakeRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getGetTotalStakeMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * Returns detailed information about a specific address.
     * </pre>
     */
    public aerium.WalletOuterClass.GetAddressInfoResponse getAddressInfo(aerium.WalletOuterClass.GetAddressInfoRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getGetAddressInfoMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * Sets or updates the label for a given address.
     * </pre>
     */
    public aerium.WalletOuterClass.SetAddressLabelResponse setAddressLabel(aerium.WalletOuterClass.SetAddressLabelRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getSetAddressLabelMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * Returns a list of all available wallets.
     * </pre>
     */
    public aerium.WalletOuterClass.ListWalletResponse listWallet(aerium.WalletOuterClass.ListWalletRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getListWalletMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * Returns detailed information about a specific wallet.
     * </pre>
     */
    public aerium.WalletOuterClass.GetWalletInfoResponse getWalletInfo(aerium.WalletOuterClass.GetWalletInfoRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getGetWalletInfoMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * Returns all addresses in the specified wallet.
     * </pre>
     */
    public aerium.WalletOuterClass.ListAddressResponse listAddress(aerium.WalletOuterClass.ListAddressRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getListAddressMethod(), getCallOptions(), request);
    }
  }

  /**
   * A stub to allow clients to do limited synchronous rpc calls to service Wallet.
   * <pre>
   * The Wallet service provides RPC methods for wallet management operations.
   * </pre>
   */
  public static final class WalletBlockingStub
      extends io.grpc.stub.AbstractBlockingStub<WalletBlockingStub> {
    private WalletBlockingStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected WalletBlockingStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new WalletBlockingStub(channel, callOptions);
    }

    /**
     * <pre>
     * Creates a new wallet with the specified parameters.
     * </pre>
     */
    public aerium.WalletOuterClass.CreateWalletResponse createWallet(aerium.WalletOuterClass.CreateWalletRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getCreateWalletMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * Restores an existing wallet with the given mnemonic.
     * </pre>
     */
    public aerium.WalletOuterClass.RestoreWalletResponse restoreWallet(aerium.WalletOuterClass.RestoreWalletRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getRestoreWalletMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * Loads an existing wallet with the given name.
     * </pre>
     */
    public aerium.WalletOuterClass.LoadWalletResponse loadWallet(aerium.WalletOuterClass.LoadWalletRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getLoadWalletMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * Unloads a currently loaded wallet with the specified name.
     * </pre>
     */
    public aerium.WalletOuterClass.UnloadWalletResponse unloadWallet(aerium.WalletOuterClass.UnloadWalletRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getUnloadWalletMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * Returns the total available balance of the wallet.
     * </pre>
     */
    public aerium.WalletOuterClass.GetTotalBalanceResponse getTotalBalance(aerium.WalletOuterClass.GetTotalBalanceRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getGetTotalBalanceMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * Signs a raw transaction for a specified wallet.
     * </pre>
     */
    public aerium.WalletOuterClass.SignRawTransactionResponse signRawTransaction(aerium.WalletOuterClass.SignRawTransactionRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getSignRawTransactionMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * Retrieves the validator address associated with a public key.
     * </pre>
     */
    public aerium.WalletOuterClass.GetValidatorAddressResponse getValidatorAddress(aerium.WalletOuterClass.GetValidatorAddressRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getGetValidatorAddressMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * Generates a new address for the specified wallet.
     * </pre>
     */
    public aerium.WalletOuterClass.GetNewAddressResponse getNewAddress(aerium.WalletOuterClass.GetNewAddressRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getGetNewAddressMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * Retrieves the transaction history of an address.
     * </pre>
     */
    public aerium.WalletOuterClass.GetAddressHistoryResponse getAddressHistory(aerium.WalletOuterClass.GetAddressHistoryRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getGetAddressHistoryMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * Signs an arbitrary message using a wallet's private key.
     * </pre>
     */
    public aerium.WalletOuterClass.SignMessageResponse signMessage(aerium.WalletOuterClass.SignMessageRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getSignMessageMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * Returns the total stake amount in the wallet.
     * </pre>
     */
    public aerium.WalletOuterClass.GetTotalStakeResponse getTotalStake(aerium.WalletOuterClass.GetTotalStakeRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getGetTotalStakeMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * Returns detailed information about a specific address.
     * </pre>
     */
    public aerium.WalletOuterClass.GetAddressInfoResponse getAddressInfo(aerium.WalletOuterClass.GetAddressInfoRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getGetAddressInfoMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * Sets or updates the label for a given address.
     * </pre>
     */
    public aerium.WalletOuterClass.SetAddressLabelResponse setAddressLabel(aerium.WalletOuterClass.SetAddressLabelRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getSetAddressLabelMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * Returns a list of all available wallets.
     * </pre>
     */
    public aerium.WalletOuterClass.ListWalletResponse listWallet(aerium.WalletOuterClass.ListWalletRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getListWalletMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * Returns detailed information about a specific wallet.
     * </pre>
     */
    public aerium.WalletOuterClass.GetWalletInfoResponse getWalletInfo(aerium.WalletOuterClass.GetWalletInfoRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getGetWalletInfoMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * Returns all addresses in the specified wallet.
     * </pre>
     */
    public aerium.WalletOuterClass.ListAddressResponse listAddress(aerium.WalletOuterClass.ListAddressRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getListAddressMethod(), getCallOptions(), request);
    }
  }

  /**
   * A stub to allow clients to do ListenableFuture-style rpc calls to service Wallet.
   * <pre>
   * The Wallet service provides RPC methods for wallet management operations.
   * </pre>
   */
  public static final class WalletFutureStub
      extends io.grpc.stub.AbstractFutureStub<WalletFutureStub> {
    private WalletFutureStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected WalletFutureStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new WalletFutureStub(channel, callOptions);
    }

    /**
     * <pre>
     * Creates a new wallet with the specified parameters.
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<aerium.WalletOuterClass.CreateWalletResponse> createWallet(
        aerium.WalletOuterClass.CreateWalletRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getCreateWalletMethod(), getCallOptions()), request);
    }

    /**
     * <pre>
     * Restores an existing wallet with the given mnemonic.
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<aerium.WalletOuterClass.RestoreWalletResponse> restoreWallet(
        aerium.WalletOuterClass.RestoreWalletRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getRestoreWalletMethod(), getCallOptions()), request);
    }

    /**
     * <pre>
     * Loads an existing wallet with the given name.
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<aerium.WalletOuterClass.LoadWalletResponse> loadWallet(
        aerium.WalletOuterClass.LoadWalletRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getLoadWalletMethod(), getCallOptions()), request);
    }

    /**
     * <pre>
     * Unloads a currently loaded wallet with the specified name.
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<aerium.WalletOuterClass.UnloadWalletResponse> unloadWallet(
        aerium.WalletOuterClass.UnloadWalletRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getUnloadWalletMethod(), getCallOptions()), request);
    }

    /**
     * <pre>
     * Returns the total available balance of the wallet.
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<aerium.WalletOuterClass.GetTotalBalanceResponse> getTotalBalance(
        aerium.WalletOuterClass.GetTotalBalanceRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getGetTotalBalanceMethod(), getCallOptions()), request);
    }

    /**
     * <pre>
     * Signs a raw transaction for a specified wallet.
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<aerium.WalletOuterClass.SignRawTransactionResponse> signRawTransaction(
        aerium.WalletOuterClass.SignRawTransactionRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getSignRawTransactionMethod(), getCallOptions()), request);
    }

    /**
     * <pre>
     * Retrieves the validator address associated with a public key.
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<aerium.WalletOuterClass.GetValidatorAddressResponse> getValidatorAddress(
        aerium.WalletOuterClass.GetValidatorAddressRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getGetValidatorAddressMethod(), getCallOptions()), request);
    }

    /**
     * <pre>
     * Generates a new address for the specified wallet.
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<aerium.WalletOuterClass.GetNewAddressResponse> getNewAddress(
        aerium.WalletOuterClass.GetNewAddressRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getGetNewAddressMethod(), getCallOptions()), request);
    }

    /**
     * <pre>
     * Retrieves the transaction history of an address.
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<aerium.WalletOuterClass.GetAddressHistoryResponse> getAddressHistory(
        aerium.WalletOuterClass.GetAddressHistoryRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getGetAddressHistoryMethod(), getCallOptions()), request);
    }

    /**
     * <pre>
     * Signs an arbitrary message using a wallet's private key.
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<aerium.WalletOuterClass.SignMessageResponse> signMessage(
        aerium.WalletOuterClass.SignMessageRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getSignMessageMethod(), getCallOptions()), request);
    }

    /**
     * <pre>
     * Returns the total stake amount in the wallet.
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<aerium.WalletOuterClass.GetTotalStakeResponse> getTotalStake(
        aerium.WalletOuterClass.GetTotalStakeRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getGetTotalStakeMethod(), getCallOptions()), request);
    }

    /**
     * <pre>
     * Returns detailed information about a specific address.
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<aerium.WalletOuterClass.GetAddressInfoResponse> getAddressInfo(
        aerium.WalletOuterClass.GetAddressInfoRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getGetAddressInfoMethod(), getCallOptions()), request);
    }

    /**
     * <pre>
     * Sets or updates the label for a given address.
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<aerium.WalletOuterClass.SetAddressLabelResponse> setAddressLabel(
        aerium.WalletOuterClass.SetAddressLabelRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getSetAddressLabelMethod(), getCallOptions()), request);
    }

    /**
     * <pre>
     * Returns a list of all available wallets.
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<aerium.WalletOuterClass.ListWalletResponse> listWallet(
        aerium.WalletOuterClass.ListWalletRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getListWalletMethod(), getCallOptions()), request);
    }

    /**
     * <pre>
     * Returns detailed information about a specific wallet.
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<aerium.WalletOuterClass.GetWalletInfoResponse> getWalletInfo(
        aerium.WalletOuterClass.GetWalletInfoRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getGetWalletInfoMethod(), getCallOptions()), request);
    }

    /**
     * <pre>
     * Returns all addresses in the specified wallet.
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<aerium.WalletOuterClass.ListAddressResponse> listAddress(
        aerium.WalletOuterClass.ListAddressRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getListAddressMethod(), getCallOptions()), request);
    }
  }

  private static final int METHODID_CREATE_WALLET = 0;
  private static final int METHODID_RESTORE_WALLET = 1;
  private static final int METHODID_LOAD_WALLET = 2;
  private static final int METHODID_UNLOAD_WALLET = 3;
  private static final int METHODID_GET_TOTAL_BALANCE = 4;
  private static final int METHODID_SIGN_RAW_TRANSACTION = 5;
  private static final int METHODID_GET_VALIDATOR_ADDRESS = 6;
  private static final int METHODID_GET_NEW_ADDRESS = 7;
  private static final int METHODID_GET_ADDRESS_HISTORY = 8;
  private static final int METHODID_SIGN_MESSAGE = 9;
  private static final int METHODID_GET_TOTAL_STAKE = 10;
  private static final int METHODID_GET_ADDRESS_INFO = 11;
  private static final int METHODID_SET_ADDRESS_LABEL = 12;
  private static final int METHODID_LIST_WALLET = 13;
  private static final int METHODID_GET_WALLET_INFO = 14;
  private static final int METHODID_LIST_ADDRESS = 15;

  private static final class MethodHandlers<Req, Resp> implements
      io.grpc.stub.ServerCalls.UnaryMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.ServerStreamingMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.ClientStreamingMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.BidiStreamingMethod<Req, Resp> {
    private final AsyncService serviceImpl;
    private final int methodId;

    MethodHandlers(AsyncService serviceImpl, int methodId) {
      this.serviceImpl = serviceImpl;
      this.methodId = methodId;
    }

    @java.lang.Override
    @java.lang.SuppressWarnings("unchecked")
    public void invoke(Req request, io.grpc.stub.StreamObserver<Resp> responseObserver) {
      switch (methodId) {
        case METHODID_CREATE_WALLET:
          serviceImpl.createWallet((aerium.WalletOuterClass.CreateWalletRequest) request,
              (io.grpc.stub.StreamObserver<aerium.WalletOuterClass.CreateWalletResponse>) responseObserver);
          break;
        case METHODID_RESTORE_WALLET:
          serviceImpl.restoreWallet((aerium.WalletOuterClass.RestoreWalletRequest) request,
              (io.grpc.stub.StreamObserver<aerium.WalletOuterClass.RestoreWalletResponse>) responseObserver);
          break;
        case METHODID_LOAD_WALLET:
          serviceImpl.loadWallet((aerium.WalletOuterClass.LoadWalletRequest) request,
              (io.grpc.stub.StreamObserver<aerium.WalletOuterClass.LoadWalletResponse>) responseObserver);
          break;
        case METHODID_UNLOAD_WALLET:
          serviceImpl.unloadWallet((aerium.WalletOuterClass.UnloadWalletRequest) request,
              (io.grpc.stub.StreamObserver<aerium.WalletOuterClass.UnloadWalletResponse>) responseObserver);
          break;
        case METHODID_GET_TOTAL_BALANCE:
          serviceImpl.getTotalBalance((aerium.WalletOuterClass.GetTotalBalanceRequest) request,
              (io.grpc.stub.StreamObserver<aerium.WalletOuterClass.GetTotalBalanceResponse>) responseObserver);
          break;
        case METHODID_SIGN_RAW_TRANSACTION:
          serviceImpl.signRawTransaction((aerium.WalletOuterClass.SignRawTransactionRequest) request,
              (io.grpc.stub.StreamObserver<aerium.WalletOuterClass.SignRawTransactionResponse>) responseObserver);
          break;
        case METHODID_GET_VALIDATOR_ADDRESS:
          serviceImpl.getValidatorAddress((aerium.WalletOuterClass.GetValidatorAddressRequest) request,
              (io.grpc.stub.StreamObserver<aerium.WalletOuterClass.GetValidatorAddressResponse>) responseObserver);
          break;
        case METHODID_GET_NEW_ADDRESS:
          serviceImpl.getNewAddress((aerium.WalletOuterClass.GetNewAddressRequest) request,
              (io.grpc.stub.StreamObserver<aerium.WalletOuterClass.GetNewAddressResponse>) responseObserver);
          break;
        case METHODID_GET_ADDRESS_HISTORY:
          serviceImpl.getAddressHistory((aerium.WalletOuterClass.GetAddressHistoryRequest) request,
              (io.grpc.stub.StreamObserver<aerium.WalletOuterClass.GetAddressHistoryResponse>) responseObserver);
          break;
        case METHODID_SIGN_MESSAGE:
          serviceImpl.signMessage((aerium.WalletOuterClass.SignMessageRequest) request,
              (io.grpc.stub.StreamObserver<aerium.WalletOuterClass.SignMessageResponse>) responseObserver);
          break;
        case METHODID_GET_TOTAL_STAKE:
          serviceImpl.getTotalStake((aerium.WalletOuterClass.GetTotalStakeRequest) request,
              (io.grpc.stub.StreamObserver<aerium.WalletOuterClass.GetTotalStakeResponse>) responseObserver);
          break;
        case METHODID_GET_ADDRESS_INFO:
          serviceImpl.getAddressInfo((aerium.WalletOuterClass.GetAddressInfoRequest) request,
              (io.grpc.stub.StreamObserver<aerium.WalletOuterClass.GetAddressInfoResponse>) responseObserver);
          break;
        case METHODID_SET_ADDRESS_LABEL:
          serviceImpl.setAddressLabel((aerium.WalletOuterClass.SetAddressLabelRequest) request,
              (io.grpc.stub.StreamObserver<aerium.WalletOuterClass.SetAddressLabelResponse>) responseObserver);
          break;
        case METHODID_LIST_WALLET:
          serviceImpl.listWallet((aerium.WalletOuterClass.ListWalletRequest) request,
              (io.grpc.stub.StreamObserver<aerium.WalletOuterClass.ListWalletResponse>) responseObserver);
          break;
        case METHODID_GET_WALLET_INFO:
          serviceImpl.getWalletInfo((aerium.WalletOuterClass.GetWalletInfoRequest) request,
              (io.grpc.stub.StreamObserver<aerium.WalletOuterClass.GetWalletInfoResponse>) responseObserver);
          break;
        case METHODID_LIST_ADDRESS:
          serviceImpl.listAddress((aerium.WalletOuterClass.ListAddressRequest) request,
              (io.grpc.stub.StreamObserver<aerium.WalletOuterClass.ListAddressResponse>) responseObserver);
          break;
        default:
          throw new AssertionError();
      }
    }

    @java.lang.Override
    @java.lang.SuppressWarnings("unchecked")
    public io.grpc.stub.StreamObserver<Req> invoke(
        io.grpc.stub.StreamObserver<Resp> responseObserver) {
      switch (methodId) {
        default:
          throw new AssertionError();
      }
    }
  }

  public static final io.grpc.ServerServiceDefinition bindService(AsyncService service) {
    return io.grpc.ServerServiceDefinition.builder(getServiceDescriptor())
        .addMethod(
          getCreateWalletMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              aerium.WalletOuterClass.CreateWalletRequest,
              aerium.WalletOuterClass.CreateWalletResponse>(
                service, METHODID_CREATE_WALLET)))
        .addMethod(
          getRestoreWalletMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              aerium.WalletOuterClass.RestoreWalletRequest,
              aerium.WalletOuterClass.RestoreWalletResponse>(
                service, METHODID_RESTORE_WALLET)))
        .addMethod(
          getLoadWalletMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              aerium.WalletOuterClass.LoadWalletRequest,
              aerium.WalletOuterClass.LoadWalletResponse>(
                service, METHODID_LOAD_WALLET)))
        .addMethod(
          getUnloadWalletMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              aerium.WalletOuterClass.UnloadWalletRequest,
              aerium.WalletOuterClass.UnloadWalletResponse>(
                service, METHODID_UNLOAD_WALLET)))
        .addMethod(
          getGetTotalBalanceMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              aerium.WalletOuterClass.GetTotalBalanceRequest,
              aerium.WalletOuterClass.GetTotalBalanceResponse>(
                service, METHODID_GET_TOTAL_BALANCE)))
        .addMethod(
          getSignRawTransactionMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              aerium.WalletOuterClass.SignRawTransactionRequest,
              aerium.WalletOuterClass.SignRawTransactionResponse>(
                service, METHODID_SIGN_RAW_TRANSACTION)))
        .addMethod(
          getGetValidatorAddressMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              aerium.WalletOuterClass.GetValidatorAddressRequest,
              aerium.WalletOuterClass.GetValidatorAddressResponse>(
                service, METHODID_GET_VALIDATOR_ADDRESS)))
        .addMethod(
          getGetNewAddressMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              aerium.WalletOuterClass.GetNewAddressRequest,
              aerium.WalletOuterClass.GetNewAddressResponse>(
                service, METHODID_GET_NEW_ADDRESS)))
        .addMethod(
          getGetAddressHistoryMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              aerium.WalletOuterClass.GetAddressHistoryRequest,
              aerium.WalletOuterClass.GetAddressHistoryResponse>(
                service, METHODID_GET_ADDRESS_HISTORY)))
        .addMethod(
          getSignMessageMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              aerium.WalletOuterClass.SignMessageRequest,
              aerium.WalletOuterClass.SignMessageResponse>(
                service, METHODID_SIGN_MESSAGE)))
        .addMethod(
          getGetTotalStakeMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              aerium.WalletOuterClass.GetTotalStakeRequest,
              aerium.WalletOuterClass.GetTotalStakeResponse>(
                service, METHODID_GET_TOTAL_STAKE)))
        .addMethod(
          getGetAddressInfoMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              aerium.WalletOuterClass.GetAddressInfoRequest,
              aerium.WalletOuterClass.GetAddressInfoResponse>(
                service, METHODID_GET_ADDRESS_INFO)))
        .addMethod(
          getSetAddressLabelMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              aerium.WalletOuterClass.SetAddressLabelRequest,
              aerium.WalletOuterClass.SetAddressLabelResponse>(
                service, METHODID_SET_ADDRESS_LABEL)))
        .addMethod(
          getListWalletMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              aerium.WalletOuterClass.ListWalletRequest,
              aerium.WalletOuterClass.ListWalletResponse>(
                service, METHODID_LIST_WALLET)))
        .addMethod(
          getGetWalletInfoMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              aerium.WalletOuterClass.GetWalletInfoRequest,
              aerium.WalletOuterClass.GetWalletInfoResponse>(
                service, METHODID_GET_WALLET_INFO)))
        .addMethod(
          getListAddressMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              aerium.WalletOuterClass.ListAddressRequest,
              aerium.WalletOuterClass.ListAddressResponse>(
                service, METHODID_LIST_ADDRESS)))
        .build();
  }

  private static abstract class WalletBaseDescriptorSupplier
      implements io.grpc.protobuf.ProtoFileDescriptorSupplier, io.grpc.protobuf.ProtoServiceDescriptorSupplier {
    WalletBaseDescriptorSupplier() {}

    @java.lang.Override
    public com.google.protobuf.Descriptors.FileDescriptor getFileDescriptor() {
      return aerium.WalletOuterClass.getDescriptor();
    }

    @java.lang.Override
    public com.google.protobuf.Descriptors.ServiceDescriptor getServiceDescriptor() {
      return getFileDescriptor().findServiceByName("Wallet");
    }
  }

  private static final class WalletFileDescriptorSupplier
      extends WalletBaseDescriptorSupplier {
    WalletFileDescriptorSupplier() {}
  }

  private static final class WalletMethodDescriptorSupplier
      extends WalletBaseDescriptorSupplier
      implements io.grpc.protobuf.ProtoMethodDescriptorSupplier {
    private final java.lang.String methodName;

    WalletMethodDescriptorSupplier(java.lang.String methodName) {
      this.methodName = methodName;
    }

    @java.lang.Override
    public com.google.protobuf.Descriptors.MethodDescriptor getMethodDescriptor() {
      return getServiceDescriptor().findMethodByName(methodName);
    }
  }

  private static volatile io.grpc.ServiceDescriptor serviceDescriptor;

  public static io.grpc.ServiceDescriptor getServiceDescriptor() {
    io.grpc.ServiceDescriptor result = serviceDescriptor;
    if (result == null) {
      synchronized (WalletGrpc.class) {
        result = serviceDescriptor;
        if (result == null) {
          serviceDescriptor = result = io.grpc.ServiceDescriptor.newBuilder(SERVICE_NAME)
              .setSchemaDescriptor(new WalletFileDescriptorSupplier())
              .addMethod(getCreateWalletMethod())
              .addMethod(getRestoreWalletMethod())
              .addMethod(getLoadWalletMethod())
              .addMethod(getUnloadWalletMethod())
              .addMethod(getGetTotalBalanceMethod())
              .addMethod(getSignRawTransactionMethod())
              .addMethod(getGetValidatorAddressMethod())
              .addMethod(getGetNewAddressMethod())
              .addMethod(getGetAddressHistoryMethod())
              .addMethod(getSignMessageMethod())
              .addMethod(getGetTotalStakeMethod())
              .addMethod(getGetAddressInfoMethod())
              .addMethod(getSetAddressLabelMethod())
              .addMethod(getListWalletMethod())
              .addMethod(getGetWalletInfoMethod())
              .addMethod(getListAddressMethod())
              .build();
        }
      }
    }
    return result;
  }
}
