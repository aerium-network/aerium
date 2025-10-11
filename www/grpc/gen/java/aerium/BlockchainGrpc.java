package aerium;

import static io.grpc.MethodDescriptor.generateFullMethodName;

/**
 * <pre>
 * Blockchain service defines RPC methods for interacting with the blockchain.
 * </pre>
 */
@javax.annotation.Generated(
    value = "by gRPC proto compiler (version 1.71.0)",
    comments = "Source: blockchain.proto")
@io.grpc.stub.annotations.GrpcGenerated
public final class BlockchainGrpc {

  private BlockchainGrpc() {}

  public static final java.lang.String SERVICE_NAME = "aerium.Blockchain";

  // Static method descriptors that strictly reflect the proto.
  private static volatile io.grpc.MethodDescriptor<aerium.BlockchainOuterClass.GetBlockRequest,
      aerium.BlockchainOuterClass.GetBlockResponse> getGetBlockMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "GetBlock",
      requestType = aerium.BlockchainOuterClass.GetBlockRequest.class,
      responseType = aerium.BlockchainOuterClass.GetBlockResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<aerium.BlockchainOuterClass.GetBlockRequest,
      aerium.BlockchainOuterClass.GetBlockResponse> getGetBlockMethod() {
    io.grpc.MethodDescriptor<aerium.BlockchainOuterClass.GetBlockRequest, aerium.BlockchainOuterClass.GetBlockResponse> getGetBlockMethod;
    if ((getGetBlockMethod = BlockchainGrpc.getGetBlockMethod) == null) {
      synchronized (BlockchainGrpc.class) {
        if ((getGetBlockMethod = BlockchainGrpc.getGetBlockMethod) == null) {
          BlockchainGrpc.getGetBlockMethod = getGetBlockMethod =
              io.grpc.MethodDescriptor.<aerium.BlockchainOuterClass.GetBlockRequest, aerium.BlockchainOuterClass.GetBlockResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "GetBlock"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  aerium.BlockchainOuterClass.GetBlockRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  aerium.BlockchainOuterClass.GetBlockResponse.getDefaultInstance()))
              .setSchemaDescriptor(new BlockchainMethodDescriptorSupplier("GetBlock"))
              .build();
        }
      }
    }
    return getGetBlockMethod;
  }

  private static volatile io.grpc.MethodDescriptor<aerium.BlockchainOuterClass.GetBlockHashRequest,
      aerium.BlockchainOuterClass.GetBlockHashResponse> getGetBlockHashMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "GetBlockHash",
      requestType = aerium.BlockchainOuterClass.GetBlockHashRequest.class,
      responseType = aerium.BlockchainOuterClass.GetBlockHashResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<aerium.BlockchainOuterClass.GetBlockHashRequest,
      aerium.BlockchainOuterClass.GetBlockHashResponse> getGetBlockHashMethod() {
    io.grpc.MethodDescriptor<aerium.BlockchainOuterClass.GetBlockHashRequest, aerium.BlockchainOuterClass.GetBlockHashResponse> getGetBlockHashMethod;
    if ((getGetBlockHashMethod = BlockchainGrpc.getGetBlockHashMethod) == null) {
      synchronized (BlockchainGrpc.class) {
        if ((getGetBlockHashMethod = BlockchainGrpc.getGetBlockHashMethod) == null) {
          BlockchainGrpc.getGetBlockHashMethod = getGetBlockHashMethod =
              io.grpc.MethodDescriptor.<aerium.BlockchainOuterClass.GetBlockHashRequest, aerium.BlockchainOuterClass.GetBlockHashResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "GetBlockHash"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  aerium.BlockchainOuterClass.GetBlockHashRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  aerium.BlockchainOuterClass.GetBlockHashResponse.getDefaultInstance()))
              .setSchemaDescriptor(new BlockchainMethodDescriptorSupplier("GetBlockHash"))
              .build();
        }
      }
    }
    return getGetBlockHashMethod;
  }

  private static volatile io.grpc.MethodDescriptor<aerium.BlockchainOuterClass.GetBlockHeightRequest,
      aerium.BlockchainOuterClass.GetBlockHeightResponse> getGetBlockHeightMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "GetBlockHeight",
      requestType = aerium.BlockchainOuterClass.GetBlockHeightRequest.class,
      responseType = aerium.BlockchainOuterClass.GetBlockHeightResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<aerium.BlockchainOuterClass.GetBlockHeightRequest,
      aerium.BlockchainOuterClass.GetBlockHeightResponse> getGetBlockHeightMethod() {
    io.grpc.MethodDescriptor<aerium.BlockchainOuterClass.GetBlockHeightRequest, aerium.BlockchainOuterClass.GetBlockHeightResponse> getGetBlockHeightMethod;
    if ((getGetBlockHeightMethod = BlockchainGrpc.getGetBlockHeightMethod) == null) {
      synchronized (BlockchainGrpc.class) {
        if ((getGetBlockHeightMethod = BlockchainGrpc.getGetBlockHeightMethod) == null) {
          BlockchainGrpc.getGetBlockHeightMethod = getGetBlockHeightMethod =
              io.grpc.MethodDescriptor.<aerium.BlockchainOuterClass.GetBlockHeightRequest, aerium.BlockchainOuterClass.GetBlockHeightResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "GetBlockHeight"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  aerium.BlockchainOuterClass.GetBlockHeightRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  aerium.BlockchainOuterClass.GetBlockHeightResponse.getDefaultInstance()))
              .setSchemaDescriptor(new BlockchainMethodDescriptorSupplier("GetBlockHeight"))
              .build();
        }
      }
    }
    return getGetBlockHeightMethod;
  }

  private static volatile io.grpc.MethodDescriptor<aerium.BlockchainOuterClass.GetBlockchainInfoRequest,
      aerium.BlockchainOuterClass.GetBlockchainInfoResponse> getGetBlockchainInfoMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "GetBlockchainInfo",
      requestType = aerium.BlockchainOuterClass.GetBlockchainInfoRequest.class,
      responseType = aerium.BlockchainOuterClass.GetBlockchainInfoResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<aerium.BlockchainOuterClass.GetBlockchainInfoRequest,
      aerium.BlockchainOuterClass.GetBlockchainInfoResponse> getGetBlockchainInfoMethod() {
    io.grpc.MethodDescriptor<aerium.BlockchainOuterClass.GetBlockchainInfoRequest, aerium.BlockchainOuterClass.GetBlockchainInfoResponse> getGetBlockchainInfoMethod;
    if ((getGetBlockchainInfoMethod = BlockchainGrpc.getGetBlockchainInfoMethod) == null) {
      synchronized (BlockchainGrpc.class) {
        if ((getGetBlockchainInfoMethod = BlockchainGrpc.getGetBlockchainInfoMethod) == null) {
          BlockchainGrpc.getGetBlockchainInfoMethod = getGetBlockchainInfoMethod =
              io.grpc.MethodDescriptor.<aerium.BlockchainOuterClass.GetBlockchainInfoRequest, aerium.BlockchainOuterClass.GetBlockchainInfoResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "GetBlockchainInfo"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  aerium.BlockchainOuterClass.GetBlockchainInfoRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  aerium.BlockchainOuterClass.GetBlockchainInfoResponse.getDefaultInstance()))
              .setSchemaDescriptor(new BlockchainMethodDescriptorSupplier("GetBlockchainInfo"))
              .build();
        }
      }
    }
    return getGetBlockchainInfoMethod;
  }

  private static volatile io.grpc.MethodDescriptor<aerium.BlockchainOuterClass.GetConsensusInfoRequest,
      aerium.BlockchainOuterClass.GetConsensusInfoResponse> getGetConsensusInfoMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "GetConsensusInfo",
      requestType = aerium.BlockchainOuterClass.GetConsensusInfoRequest.class,
      responseType = aerium.BlockchainOuterClass.GetConsensusInfoResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<aerium.BlockchainOuterClass.GetConsensusInfoRequest,
      aerium.BlockchainOuterClass.GetConsensusInfoResponse> getGetConsensusInfoMethod() {
    io.grpc.MethodDescriptor<aerium.BlockchainOuterClass.GetConsensusInfoRequest, aerium.BlockchainOuterClass.GetConsensusInfoResponse> getGetConsensusInfoMethod;
    if ((getGetConsensusInfoMethod = BlockchainGrpc.getGetConsensusInfoMethod) == null) {
      synchronized (BlockchainGrpc.class) {
        if ((getGetConsensusInfoMethod = BlockchainGrpc.getGetConsensusInfoMethod) == null) {
          BlockchainGrpc.getGetConsensusInfoMethod = getGetConsensusInfoMethod =
              io.grpc.MethodDescriptor.<aerium.BlockchainOuterClass.GetConsensusInfoRequest, aerium.BlockchainOuterClass.GetConsensusInfoResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "GetConsensusInfo"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  aerium.BlockchainOuterClass.GetConsensusInfoRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  aerium.BlockchainOuterClass.GetConsensusInfoResponse.getDefaultInstance()))
              .setSchemaDescriptor(new BlockchainMethodDescriptorSupplier("GetConsensusInfo"))
              .build();
        }
      }
    }
    return getGetConsensusInfoMethod;
  }

  private static volatile io.grpc.MethodDescriptor<aerium.BlockchainOuterClass.GetAccountRequest,
      aerium.BlockchainOuterClass.GetAccountResponse> getGetAccountMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "GetAccount",
      requestType = aerium.BlockchainOuterClass.GetAccountRequest.class,
      responseType = aerium.BlockchainOuterClass.GetAccountResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<aerium.BlockchainOuterClass.GetAccountRequest,
      aerium.BlockchainOuterClass.GetAccountResponse> getGetAccountMethod() {
    io.grpc.MethodDescriptor<aerium.BlockchainOuterClass.GetAccountRequest, aerium.BlockchainOuterClass.GetAccountResponse> getGetAccountMethod;
    if ((getGetAccountMethod = BlockchainGrpc.getGetAccountMethod) == null) {
      synchronized (BlockchainGrpc.class) {
        if ((getGetAccountMethod = BlockchainGrpc.getGetAccountMethod) == null) {
          BlockchainGrpc.getGetAccountMethod = getGetAccountMethod =
              io.grpc.MethodDescriptor.<aerium.BlockchainOuterClass.GetAccountRequest, aerium.BlockchainOuterClass.GetAccountResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "GetAccount"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  aerium.BlockchainOuterClass.GetAccountRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  aerium.BlockchainOuterClass.GetAccountResponse.getDefaultInstance()))
              .setSchemaDescriptor(new BlockchainMethodDescriptorSupplier("GetAccount"))
              .build();
        }
      }
    }
    return getGetAccountMethod;
  }

  private static volatile io.grpc.MethodDescriptor<aerium.BlockchainOuterClass.GetValidatorRequest,
      aerium.BlockchainOuterClass.GetValidatorResponse> getGetValidatorMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "GetValidator",
      requestType = aerium.BlockchainOuterClass.GetValidatorRequest.class,
      responseType = aerium.BlockchainOuterClass.GetValidatorResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<aerium.BlockchainOuterClass.GetValidatorRequest,
      aerium.BlockchainOuterClass.GetValidatorResponse> getGetValidatorMethod() {
    io.grpc.MethodDescriptor<aerium.BlockchainOuterClass.GetValidatorRequest, aerium.BlockchainOuterClass.GetValidatorResponse> getGetValidatorMethod;
    if ((getGetValidatorMethod = BlockchainGrpc.getGetValidatorMethod) == null) {
      synchronized (BlockchainGrpc.class) {
        if ((getGetValidatorMethod = BlockchainGrpc.getGetValidatorMethod) == null) {
          BlockchainGrpc.getGetValidatorMethod = getGetValidatorMethod =
              io.grpc.MethodDescriptor.<aerium.BlockchainOuterClass.GetValidatorRequest, aerium.BlockchainOuterClass.GetValidatorResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "GetValidator"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  aerium.BlockchainOuterClass.GetValidatorRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  aerium.BlockchainOuterClass.GetValidatorResponse.getDefaultInstance()))
              .setSchemaDescriptor(new BlockchainMethodDescriptorSupplier("GetValidator"))
              .build();
        }
      }
    }
    return getGetValidatorMethod;
  }

  private static volatile io.grpc.MethodDescriptor<aerium.BlockchainOuterClass.GetValidatorByNumberRequest,
      aerium.BlockchainOuterClass.GetValidatorResponse> getGetValidatorByNumberMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "GetValidatorByNumber",
      requestType = aerium.BlockchainOuterClass.GetValidatorByNumberRequest.class,
      responseType = aerium.BlockchainOuterClass.GetValidatorResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<aerium.BlockchainOuterClass.GetValidatorByNumberRequest,
      aerium.BlockchainOuterClass.GetValidatorResponse> getGetValidatorByNumberMethod() {
    io.grpc.MethodDescriptor<aerium.BlockchainOuterClass.GetValidatorByNumberRequest, aerium.BlockchainOuterClass.GetValidatorResponse> getGetValidatorByNumberMethod;
    if ((getGetValidatorByNumberMethod = BlockchainGrpc.getGetValidatorByNumberMethod) == null) {
      synchronized (BlockchainGrpc.class) {
        if ((getGetValidatorByNumberMethod = BlockchainGrpc.getGetValidatorByNumberMethod) == null) {
          BlockchainGrpc.getGetValidatorByNumberMethod = getGetValidatorByNumberMethod =
              io.grpc.MethodDescriptor.<aerium.BlockchainOuterClass.GetValidatorByNumberRequest, aerium.BlockchainOuterClass.GetValidatorResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "GetValidatorByNumber"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  aerium.BlockchainOuterClass.GetValidatorByNumberRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  aerium.BlockchainOuterClass.GetValidatorResponse.getDefaultInstance()))
              .setSchemaDescriptor(new BlockchainMethodDescriptorSupplier("GetValidatorByNumber"))
              .build();
        }
      }
    }
    return getGetValidatorByNumberMethod;
  }

  private static volatile io.grpc.MethodDescriptor<aerium.BlockchainOuterClass.GetValidatorAddressesRequest,
      aerium.BlockchainOuterClass.GetValidatorAddressesResponse> getGetValidatorAddressesMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "GetValidatorAddresses",
      requestType = aerium.BlockchainOuterClass.GetValidatorAddressesRequest.class,
      responseType = aerium.BlockchainOuterClass.GetValidatorAddressesResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<aerium.BlockchainOuterClass.GetValidatorAddressesRequest,
      aerium.BlockchainOuterClass.GetValidatorAddressesResponse> getGetValidatorAddressesMethod() {
    io.grpc.MethodDescriptor<aerium.BlockchainOuterClass.GetValidatorAddressesRequest, aerium.BlockchainOuterClass.GetValidatorAddressesResponse> getGetValidatorAddressesMethod;
    if ((getGetValidatorAddressesMethod = BlockchainGrpc.getGetValidatorAddressesMethod) == null) {
      synchronized (BlockchainGrpc.class) {
        if ((getGetValidatorAddressesMethod = BlockchainGrpc.getGetValidatorAddressesMethod) == null) {
          BlockchainGrpc.getGetValidatorAddressesMethod = getGetValidatorAddressesMethod =
              io.grpc.MethodDescriptor.<aerium.BlockchainOuterClass.GetValidatorAddressesRequest, aerium.BlockchainOuterClass.GetValidatorAddressesResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "GetValidatorAddresses"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  aerium.BlockchainOuterClass.GetValidatorAddressesRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  aerium.BlockchainOuterClass.GetValidatorAddressesResponse.getDefaultInstance()))
              .setSchemaDescriptor(new BlockchainMethodDescriptorSupplier("GetValidatorAddresses"))
              .build();
        }
      }
    }
    return getGetValidatorAddressesMethod;
  }

  private static volatile io.grpc.MethodDescriptor<aerium.BlockchainOuterClass.GetPublicKeyRequest,
      aerium.BlockchainOuterClass.GetPublicKeyResponse> getGetPublicKeyMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "GetPublicKey",
      requestType = aerium.BlockchainOuterClass.GetPublicKeyRequest.class,
      responseType = aerium.BlockchainOuterClass.GetPublicKeyResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<aerium.BlockchainOuterClass.GetPublicKeyRequest,
      aerium.BlockchainOuterClass.GetPublicKeyResponse> getGetPublicKeyMethod() {
    io.grpc.MethodDescriptor<aerium.BlockchainOuterClass.GetPublicKeyRequest, aerium.BlockchainOuterClass.GetPublicKeyResponse> getGetPublicKeyMethod;
    if ((getGetPublicKeyMethod = BlockchainGrpc.getGetPublicKeyMethod) == null) {
      synchronized (BlockchainGrpc.class) {
        if ((getGetPublicKeyMethod = BlockchainGrpc.getGetPublicKeyMethod) == null) {
          BlockchainGrpc.getGetPublicKeyMethod = getGetPublicKeyMethod =
              io.grpc.MethodDescriptor.<aerium.BlockchainOuterClass.GetPublicKeyRequest, aerium.BlockchainOuterClass.GetPublicKeyResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "GetPublicKey"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  aerium.BlockchainOuterClass.GetPublicKeyRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  aerium.BlockchainOuterClass.GetPublicKeyResponse.getDefaultInstance()))
              .setSchemaDescriptor(new BlockchainMethodDescriptorSupplier("GetPublicKey"))
              .build();
        }
      }
    }
    return getGetPublicKeyMethod;
  }

  private static volatile io.grpc.MethodDescriptor<aerium.BlockchainOuterClass.GetTxPoolContentRequest,
      aerium.BlockchainOuterClass.GetTxPoolContentResponse> getGetTxPoolContentMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "GetTxPoolContent",
      requestType = aerium.BlockchainOuterClass.GetTxPoolContentRequest.class,
      responseType = aerium.BlockchainOuterClass.GetTxPoolContentResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<aerium.BlockchainOuterClass.GetTxPoolContentRequest,
      aerium.BlockchainOuterClass.GetTxPoolContentResponse> getGetTxPoolContentMethod() {
    io.grpc.MethodDescriptor<aerium.BlockchainOuterClass.GetTxPoolContentRequest, aerium.BlockchainOuterClass.GetTxPoolContentResponse> getGetTxPoolContentMethod;
    if ((getGetTxPoolContentMethod = BlockchainGrpc.getGetTxPoolContentMethod) == null) {
      synchronized (BlockchainGrpc.class) {
        if ((getGetTxPoolContentMethod = BlockchainGrpc.getGetTxPoolContentMethod) == null) {
          BlockchainGrpc.getGetTxPoolContentMethod = getGetTxPoolContentMethod =
              io.grpc.MethodDescriptor.<aerium.BlockchainOuterClass.GetTxPoolContentRequest, aerium.BlockchainOuterClass.GetTxPoolContentResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "GetTxPoolContent"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  aerium.BlockchainOuterClass.GetTxPoolContentRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  aerium.BlockchainOuterClass.GetTxPoolContentResponse.getDefaultInstance()))
              .setSchemaDescriptor(new BlockchainMethodDescriptorSupplier("GetTxPoolContent"))
              .build();
        }
      }
    }
    return getGetTxPoolContentMethod;
  }

  /**
   * Creates a new async stub that supports all call types for the service
   */
  public static BlockchainStub newStub(io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<BlockchainStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<BlockchainStub>() {
        @java.lang.Override
        public BlockchainStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new BlockchainStub(channel, callOptions);
        }
      };
    return BlockchainStub.newStub(factory, channel);
  }

  /**
   * Creates a new blocking-style stub that supports all types of calls on the service
   */
  public static BlockchainBlockingV2Stub newBlockingV2Stub(
      io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<BlockchainBlockingV2Stub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<BlockchainBlockingV2Stub>() {
        @java.lang.Override
        public BlockchainBlockingV2Stub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new BlockchainBlockingV2Stub(channel, callOptions);
        }
      };
    return BlockchainBlockingV2Stub.newStub(factory, channel);
  }

  /**
   * Creates a new blocking-style stub that supports unary and streaming output calls on the service
   */
  public static BlockchainBlockingStub newBlockingStub(
      io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<BlockchainBlockingStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<BlockchainBlockingStub>() {
        @java.lang.Override
        public BlockchainBlockingStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new BlockchainBlockingStub(channel, callOptions);
        }
      };
    return BlockchainBlockingStub.newStub(factory, channel);
  }

  /**
   * Creates a new ListenableFuture-style stub that supports unary calls on the service
   */
  public static BlockchainFutureStub newFutureStub(
      io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<BlockchainFutureStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<BlockchainFutureStub>() {
        @java.lang.Override
        public BlockchainFutureStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new BlockchainFutureStub(channel, callOptions);
        }
      };
    return BlockchainFutureStub.newStub(factory, channel);
  }

  /**
   * <pre>
   * Blockchain service defines RPC methods for interacting with the blockchain.
   * </pre>
   */
  public interface AsyncService {

    /**
     * <pre>
     * GetBlock retrieves information about a block based on the provided request parameters.
     * </pre>
     */
    default void getBlock(aerium.BlockchainOuterClass.GetBlockRequest request,
        io.grpc.stub.StreamObserver<aerium.BlockchainOuterClass.GetBlockResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getGetBlockMethod(), responseObserver);
    }

    /**
     * <pre>
     * GetBlockHash retrieves the hash of a block at the specified height.
     * </pre>
     */
    default void getBlockHash(aerium.BlockchainOuterClass.GetBlockHashRequest request,
        io.grpc.stub.StreamObserver<aerium.BlockchainOuterClass.GetBlockHashResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getGetBlockHashMethod(), responseObserver);
    }

    /**
     * <pre>
     * GetBlockHeight retrieves the height of a block with the specified hash.
     * </pre>
     */
    default void getBlockHeight(aerium.BlockchainOuterClass.GetBlockHeightRequest request,
        io.grpc.stub.StreamObserver<aerium.BlockchainOuterClass.GetBlockHeightResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getGetBlockHeightMethod(), responseObserver);
    }

    /**
     * <pre>
     * GetBlockchainInfo retrieves general information about the blockchain.
     * </pre>
     */
    default void getBlockchainInfo(aerium.BlockchainOuterClass.GetBlockchainInfoRequest request,
        io.grpc.stub.StreamObserver<aerium.BlockchainOuterClass.GetBlockchainInfoResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getGetBlockchainInfoMethod(), responseObserver);
    }

    /**
     * <pre>
     * GetConsensusInfo retrieves information about the consensus instances.
     * </pre>
     */
    default void getConsensusInfo(aerium.BlockchainOuterClass.GetConsensusInfoRequest request,
        io.grpc.stub.StreamObserver<aerium.BlockchainOuterClass.GetConsensusInfoResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getGetConsensusInfoMethod(), responseObserver);
    }

    /**
     * <pre>
     * GetAccount retrieves information about an account based on the provided address.
     * </pre>
     */
    default void getAccount(aerium.BlockchainOuterClass.GetAccountRequest request,
        io.grpc.stub.StreamObserver<aerium.BlockchainOuterClass.GetAccountResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getGetAccountMethod(), responseObserver);
    }

    /**
     * <pre>
     * GetValidator retrieves information about a validator based on the provided address.
     * </pre>
     */
    default void getValidator(aerium.BlockchainOuterClass.GetValidatorRequest request,
        io.grpc.stub.StreamObserver<aerium.BlockchainOuterClass.GetValidatorResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getGetValidatorMethod(), responseObserver);
    }

    /**
     * <pre>
     * GetValidatorByNumber retrieves information about a validator based on the provided number.
     * </pre>
     */
    default void getValidatorByNumber(aerium.BlockchainOuterClass.GetValidatorByNumberRequest request,
        io.grpc.stub.StreamObserver<aerium.BlockchainOuterClass.GetValidatorResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getGetValidatorByNumberMethod(), responseObserver);
    }

    /**
     * <pre>
     * GetValidatorAddresses retrieves a list of all validator addresses.
     * </pre>
     */
    default void getValidatorAddresses(aerium.BlockchainOuterClass.GetValidatorAddressesRequest request,
        io.grpc.stub.StreamObserver<aerium.BlockchainOuterClass.GetValidatorAddressesResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getGetValidatorAddressesMethod(), responseObserver);
    }

    /**
     * <pre>
     * GetPublicKey retrieves the public key of an account based on the provided address.
     * </pre>
     */
    default void getPublicKey(aerium.BlockchainOuterClass.GetPublicKeyRequest request,
        io.grpc.stub.StreamObserver<aerium.BlockchainOuterClass.GetPublicKeyResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getGetPublicKeyMethod(), responseObserver);
    }

    /**
     * <pre>
     * GetTxPoolContent retrieves current transactions in the transaction pool.
     * </pre>
     */
    default void getTxPoolContent(aerium.BlockchainOuterClass.GetTxPoolContentRequest request,
        io.grpc.stub.StreamObserver<aerium.BlockchainOuterClass.GetTxPoolContentResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getGetTxPoolContentMethod(), responseObserver);
    }
  }

  /**
   * Base class for the server implementation of the service Blockchain.
   * <pre>
   * Blockchain service defines RPC methods for interacting with the blockchain.
   * </pre>
   */
  public static abstract class BlockchainImplBase
      implements io.grpc.BindableService, AsyncService {

    @java.lang.Override public final io.grpc.ServerServiceDefinition bindService() {
      return BlockchainGrpc.bindService(this);
    }
  }

  /**
   * A stub to allow clients to do asynchronous rpc calls to service Blockchain.
   * <pre>
   * Blockchain service defines RPC methods for interacting with the blockchain.
   * </pre>
   */
  public static final class BlockchainStub
      extends io.grpc.stub.AbstractAsyncStub<BlockchainStub> {
    private BlockchainStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected BlockchainStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new BlockchainStub(channel, callOptions);
    }

    /**
     * <pre>
     * GetBlock retrieves information about a block based on the provided request parameters.
     * </pre>
     */
    public void getBlock(aerium.BlockchainOuterClass.GetBlockRequest request,
        io.grpc.stub.StreamObserver<aerium.BlockchainOuterClass.GetBlockResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getGetBlockMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     * <pre>
     * GetBlockHash retrieves the hash of a block at the specified height.
     * </pre>
     */
    public void getBlockHash(aerium.BlockchainOuterClass.GetBlockHashRequest request,
        io.grpc.stub.StreamObserver<aerium.BlockchainOuterClass.GetBlockHashResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getGetBlockHashMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     * <pre>
     * GetBlockHeight retrieves the height of a block with the specified hash.
     * </pre>
     */
    public void getBlockHeight(aerium.BlockchainOuterClass.GetBlockHeightRequest request,
        io.grpc.stub.StreamObserver<aerium.BlockchainOuterClass.GetBlockHeightResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getGetBlockHeightMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     * <pre>
     * GetBlockchainInfo retrieves general information about the blockchain.
     * </pre>
     */
    public void getBlockchainInfo(aerium.BlockchainOuterClass.GetBlockchainInfoRequest request,
        io.grpc.stub.StreamObserver<aerium.BlockchainOuterClass.GetBlockchainInfoResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getGetBlockchainInfoMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     * <pre>
     * GetConsensusInfo retrieves information about the consensus instances.
     * </pre>
     */
    public void getConsensusInfo(aerium.BlockchainOuterClass.GetConsensusInfoRequest request,
        io.grpc.stub.StreamObserver<aerium.BlockchainOuterClass.GetConsensusInfoResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getGetConsensusInfoMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     * <pre>
     * GetAccount retrieves information about an account based on the provided address.
     * </pre>
     */
    public void getAccount(aerium.BlockchainOuterClass.GetAccountRequest request,
        io.grpc.stub.StreamObserver<aerium.BlockchainOuterClass.GetAccountResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getGetAccountMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     * <pre>
     * GetValidator retrieves information about a validator based on the provided address.
     * </pre>
     */
    public void getValidator(aerium.BlockchainOuterClass.GetValidatorRequest request,
        io.grpc.stub.StreamObserver<aerium.BlockchainOuterClass.GetValidatorResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getGetValidatorMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     * <pre>
     * GetValidatorByNumber retrieves information about a validator based on the provided number.
     * </pre>
     */
    public void getValidatorByNumber(aerium.BlockchainOuterClass.GetValidatorByNumberRequest request,
        io.grpc.stub.StreamObserver<aerium.BlockchainOuterClass.GetValidatorResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getGetValidatorByNumberMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     * <pre>
     * GetValidatorAddresses retrieves a list of all validator addresses.
     * </pre>
     */
    public void getValidatorAddresses(aerium.BlockchainOuterClass.GetValidatorAddressesRequest request,
        io.grpc.stub.StreamObserver<aerium.BlockchainOuterClass.GetValidatorAddressesResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getGetValidatorAddressesMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     * <pre>
     * GetPublicKey retrieves the public key of an account based on the provided address.
     * </pre>
     */
    public void getPublicKey(aerium.BlockchainOuterClass.GetPublicKeyRequest request,
        io.grpc.stub.StreamObserver<aerium.BlockchainOuterClass.GetPublicKeyResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getGetPublicKeyMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     * <pre>
     * GetTxPoolContent retrieves current transactions in the transaction pool.
     * </pre>
     */
    public void getTxPoolContent(aerium.BlockchainOuterClass.GetTxPoolContentRequest request,
        io.grpc.stub.StreamObserver<aerium.BlockchainOuterClass.GetTxPoolContentResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getGetTxPoolContentMethod(), getCallOptions()), request, responseObserver);
    }
  }

  /**
   * A stub to allow clients to do synchronous rpc calls to service Blockchain.
   * <pre>
   * Blockchain service defines RPC methods for interacting with the blockchain.
   * </pre>
   */
  public static final class BlockchainBlockingV2Stub
      extends io.grpc.stub.AbstractBlockingStub<BlockchainBlockingV2Stub> {
    private BlockchainBlockingV2Stub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected BlockchainBlockingV2Stub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new BlockchainBlockingV2Stub(channel, callOptions);
    }

    /**
     * <pre>
     * GetBlock retrieves information about a block based on the provided request parameters.
     * </pre>
     */
    public aerium.BlockchainOuterClass.GetBlockResponse getBlock(aerium.BlockchainOuterClass.GetBlockRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getGetBlockMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * GetBlockHash retrieves the hash of a block at the specified height.
     * </pre>
     */
    public aerium.BlockchainOuterClass.GetBlockHashResponse getBlockHash(aerium.BlockchainOuterClass.GetBlockHashRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getGetBlockHashMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * GetBlockHeight retrieves the height of a block with the specified hash.
     * </pre>
     */
    public aerium.BlockchainOuterClass.GetBlockHeightResponse getBlockHeight(aerium.BlockchainOuterClass.GetBlockHeightRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getGetBlockHeightMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * GetBlockchainInfo retrieves general information about the blockchain.
     * </pre>
     */
    public aerium.BlockchainOuterClass.GetBlockchainInfoResponse getBlockchainInfo(aerium.BlockchainOuterClass.GetBlockchainInfoRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getGetBlockchainInfoMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * GetConsensusInfo retrieves information about the consensus instances.
     * </pre>
     */
    public aerium.BlockchainOuterClass.GetConsensusInfoResponse getConsensusInfo(aerium.BlockchainOuterClass.GetConsensusInfoRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getGetConsensusInfoMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * GetAccount retrieves information about an account based on the provided address.
     * </pre>
     */
    public aerium.BlockchainOuterClass.GetAccountResponse getAccount(aerium.BlockchainOuterClass.GetAccountRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getGetAccountMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * GetValidator retrieves information about a validator based on the provided address.
     * </pre>
     */
    public aerium.BlockchainOuterClass.GetValidatorResponse getValidator(aerium.BlockchainOuterClass.GetValidatorRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getGetValidatorMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * GetValidatorByNumber retrieves information about a validator based on the provided number.
     * </pre>
     */
    public aerium.BlockchainOuterClass.GetValidatorResponse getValidatorByNumber(aerium.BlockchainOuterClass.GetValidatorByNumberRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getGetValidatorByNumberMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * GetValidatorAddresses retrieves a list of all validator addresses.
     * </pre>
     */
    public aerium.BlockchainOuterClass.GetValidatorAddressesResponse getValidatorAddresses(aerium.BlockchainOuterClass.GetValidatorAddressesRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getGetValidatorAddressesMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * GetPublicKey retrieves the public key of an account based on the provided address.
     * </pre>
     */
    public aerium.BlockchainOuterClass.GetPublicKeyResponse getPublicKey(aerium.BlockchainOuterClass.GetPublicKeyRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getGetPublicKeyMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * GetTxPoolContent retrieves current transactions in the transaction pool.
     * </pre>
     */
    public aerium.BlockchainOuterClass.GetTxPoolContentResponse getTxPoolContent(aerium.BlockchainOuterClass.GetTxPoolContentRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getGetTxPoolContentMethod(), getCallOptions(), request);
    }
  }

  /**
   * A stub to allow clients to do limited synchronous rpc calls to service Blockchain.
   * <pre>
   * Blockchain service defines RPC methods for interacting with the blockchain.
   * </pre>
   */
  public static final class BlockchainBlockingStub
      extends io.grpc.stub.AbstractBlockingStub<BlockchainBlockingStub> {
    private BlockchainBlockingStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected BlockchainBlockingStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new BlockchainBlockingStub(channel, callOptions);
    }

    /**
     * <pre>
     * GetBlock retrieves information about a block based on the provided request parameters.
     * </pre>
     */
    public aerium.BlockchainOuterClass.GetBlockResponse getBlock(aerium.BlockchainOuterClass.GetBlockRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getGetBlockMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * GetBlockHash retrieves the hash of a block at the specified height.
     * </pre>
     */
    public aerium.BlockchainOuterClass.GetBlockHashResponse getBlockHash(aerium.BlockchainOuterClass.GetBlockHashRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getGetBlockHashMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * GetBlockHeight retrieves the height of a block with the specified hash.
     * </pre>
     */
    public aerium.BlockchainOuterClass.GetBlockHeightResponse getBlockHeight(aerium.BlockchainOuterClass.GetBlockHeightRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getGetBlockHeightMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * GetBlockchainInfo retrieves general information about the blockchain.
     * </pre>
     */
    public aerium.BlockchainOuterClass.GetBlockchainInfoResponse getBlockchainInfo(aerium.BlockchainOuterClass.GetBlockchainInfoRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getGetBlockchainInfoMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * GetConsensusInfo retrieves information about the consensus instances.
     * </pre>
     */
    public aerium.BlockchainOuterClass.GetConsensusInfoResponse getConsensusInfo(aerium.BlockchainOuterClass.GetConsensusInfoRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getGetConsensusInfoMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * GetAccount retrieves information about an account based on the provided address.
     * </pre>
     */
    public aerium.BlockchainOuterClass.GetAccountResponse getAccount(aerium.BlockchainOuterClass.GetAccountRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getGetAccountMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * GetValidator retrieves information about a validator based on the provided address.
     * </pre>
     */
    public aerium.BlockchainOuterClass.GetValidatorResponse getValidator(aerium.BlockchainOuterClass.GetValidatorRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getGetValidatorMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * GetValidatorByNumber retrieves information about a validator based on the provided number.
     * </pre>
     */
    public aerium.BlockchainOuterClass.GetValidatorResponse getValidatorByNumber(aerium.BlockchainOuterClass.GetValidatorByNumberRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getGetValidatorByNumberMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * GetValidatorAddresses retrieves a list of all validator addresses.
     * </pre>
     */
    public aerium.BlockchainOuterClass.GetValidatorAddressesResponse getValidatorAddresses(aerium.BlockchainOuterClass.GetValidatorAddressesRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getGetValidatorAddressesMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * GetPublicKey retrieves the public key of an account based on the provided address.
     * </pre>
     */
    public aerium.BlockchainOuterClass.GetPublicKeyResponse getPublicKey(aerium.BlockchainOuterClass.GetPublicKeyRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getGetPublicKeyMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * GetTxPoolContent retrieves current transactions in the transaction pool.
     * </pre>
     */
    public aerium.BlockchainOuterClass.GetTxPoolContentResponse getTxPoolContent(aerium.BlockchainOuterClass.GetTxPoolContentRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getGetTxPoolContentMethod(), getCallOptions(), request);
    }
  }

  /**
   * A stub to allow clients to do ListenableFuture-style rpc calls to service Blockchain.
   * <pre>
   * Blockchain service defines RPC methods for interacting with the blockchain.
   * </pre>
   */
  public static final class BlockchainFutureStub
      extends io.grpc.stub.AbstractFutureStub<BlockchainFutureStub> {
    private BlockchainFutureStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected BlockchainFutureStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new BlockchainFutureStub(channel, callOptions);
    }

    /**
     * <pre>
     * GetBlock retrieves information about a block based on the provided request parameters.
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<aerium.BlockchainOuterClass.GetBlockResponse> getBlock(
        aerium.BlockchainOuterClass.GetBlockRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getGetBlockMethod(), getCallOptions()), request);
    }

    /**
     * <pre>
     * GetBlockHash retrieves the hash of a block at the specified height.
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<aerium.BlockchainOuterClass.GetBlockHashResponse> getBlockHash(
        aerium.BlockchainOuterClass.GetBlockHashRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getGetBlockHashMethod(), getCallOptions()), request);
    }

    /**
     * <pre>
     * GetBlockHeight retrieves the height of a block with the specified hash.
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<aerium.BlockchainOuterClass.GetBlockHeightResponse> getBlockHeight(
        aerium.BlockchainOuterClass.GetBlockHeightRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getGetBlockHeightMethod(), getCallOptions()), request);
    }

    /**
     * <pre>
     * GetBlockchainInfo retrieves general information about the blockchain.
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<aerium.BlockchainOuterClass.GetBlockchainInfoResponse> getBlockchainInfo(
        aerium.BlockchainOuterClass.GetBlockchainInfoRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getGetBlockchainInfoMethod(), getCallOptions()), request);
    }

    /**
     * <pre>
     * GetConsensusInfo retrieves information about the consensus instances.
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<aerium.BlockchainOuterClass.GetConsensusInfoResponse> getConsensusInfo(
        aerium.BlockchainOuterClass.GetConsensusInfoRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getGetConsensusInfoMethod(), getCallOptions()), request);
    }

    /**
     * <pre>
     * GetAccount retrieves information about an account based on the provided address.
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<aerium.BlockchainOuterClass.GetAccountResponse> getAccount(
        aerium.BlockchainOuterClass.GetAccountRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getGetAccountMethod(), getCallOptions()), request);
    }

    /**
     * <pre>
     * GetValidator retrieves information about a validator based on the provided address.
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<aerium.BlockchainOuterClass.GetValidatorResponse> getValidator(
        aerium.BlockchainOuterClass.GetValidatorRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getGetValidatorMethod(), getCallOptions()), request);
    }

    /**
     * <pre>
     * GetValidatorByNumber retrieves information about a validator based on the provided number.
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<aerium.BlockchainOuterClass.GetValidatorResponse> getValidatorByNumber(
        aerium.BlockchainOuterClass.GetValidatorByNumberRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getGetValidatorByNumberMethod(), getCallOptions()), request);
    }

    /**
     * <pre>
     * GetValidatorAddresses retrieves a list of all validator addresses.
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<aerium.BlockchainOuterClass.GetValidatorAddressesResponse> getValidatorAddresses(
        aerium.BlockchainOuterClass.GetValidatorAddressesRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getGetValidatorAddressesMethod(), getCallOptions()), request);
    }

    /**
     * <pre>
     * GetPublicKey retrieves the public key of an account based on the provided address.
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<aerium.BlockchainOuterClass.GetPublicKeyResponse> getPublicKey(
        aerium.BlockchainOuterClass.GetPublicKeyRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getGetPublicKeyMethod(), getCallOptions()), request);
    }

    /**
     * <pre>
     * GetTxPoolContent retrieves current transactions in the transaction pool.
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<aerium.BlockchainOuterClass.GetTxPoolContentResponse> getTxPoolContent(
        aerium.BlockchainOuterClass.GetTxPoolContentRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getGetTxPoolContentMethod(), getCallOptions()), request);
    }
  }

  private static final int METHODID_GET_BLOCK = 0;
  private static final int METHODID_GET_BLOCK_HASH = 1;
  private static final int METHODID_GET_BLOCK_HEIGHT = 2;
  private static final int METHODID_GET_BLOCKCHAIN_INFO = 3;
  private static final int METHODID_GET_CONSENSUS_INFO = 4;
  private static final int METHODID_GET_ACCOUNT = 5;
  private static final int METHODID_GET_VALIDATOR = 6;
  private static final int METHODID_GET_VALIDATOR_BY_NUMBER = 7;
  private static final int METHODID_GET_VALIDATOR_ADDRESSES = 8;
  private static final int METHODID_GET_PUBLIC_KEY = 9;
  private static final int METHODID_GET_TX_POOL_CONTENT = 10;

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
        case METHODID_GET_BLOCK:
          serviceImpl.getBlock((aerium.BlockchainOuterClass.GetBlockRequest) request,
              (io.grpc.stub.StreamObserver<aerium.BlockchainOuterClass.GetBlockResponse>) responseObserver);
          break;
        case METHODID_GET_BLOCK_HASH:
          serviceImpl.getBlockHash((aerium.BlockchainOuterClass.GetBlockHashRequest) request,
              (io.grpc.stub.StreamObserver<aerium.BlockchainOuterClass.GetBlockHashResponse>) responseObserver);
          break;
        case METHODID_GET_BLOCK_HEIGHT:
          serviceImpl.getBlockHeight((aerium.BlockchainOuterClass.GetBlockHeightRequest) request,
              (io.grpc.stub.StreamObserver<aerium.BlockchainOuterClass.GetBlockHeightResponse>) responseObserver);
          break;
        case METHODID_GET_BLOCKCHAIN_INFO:
          serviceImpl.getBlockchainInfo((aerium.BlockchainOuterClass.GetBlockchainInfoRequest) request,
              (io.grpc.stub.StreamObserver<aerium.BlockchainOuterClass.GetBlockchainInfoResponse>) responseObserver);
          break;
        case METHODID_GET_CONSENSUS_INFO:
          serviceImpl.getConsensusInfo((aerium.BlockchainOuterClass.GetConsensusInfoRequest) request,
              (io.grpc.stub.StreamObserver<aerium.BlockchainOuterClass.GetConsensusInfoResponse>) responseObserver);
          break;
        case METHODID_GET_ACCOUNT:
          serviceImpl.getAccount((aerium.BlockchainOuterClass.GetAccountRequest) request,
              (io.grpc.stub.StreamObserver<aerium.BlockchainOuterClass.GetAccountResponse>) responseObserver);
          break;
        case METHODID_GET_VALIDATOR:
          serviceImpl.getValidator((aerium.BlockchainOuterClass.GetValidatorRequest) request,
              (io.grpc.stub.StreamObserver<aerium.BlockchainOuterClass.GetValidatorResponse>) responseObserver);
          break;
        case METHODID_GET_VALIDATOR_BY_NUMBER:
          serviceImpl.getValidatorByNumber((aerium.BlockchainOuterClass.GetValidatorByNumberRequest) request,
              (io.grpc.stub.StreamObserver<aerium.BlockchainOuterClass.GetValidatorResponse>) responseObserver);
          break;
        case METHODID_GET_VALIDATOR_ADDRESSES:
          serviceImpl.getValidatorAddresses((aerium.BlockchainOuterClass.GetValidatorAddressesRequest) request,
              (io.grpc.stub.StreamObserver<aerium.BlockchainOuterClass.GetValidatorAddressesResponse>) responseObserver);
          break;
        case METHODID_GET_PUBLIC_KEY:
          serviceImpl.getPublicKey((aerium.BlockchainOuterClass.GetPublicKeyRequest) request,
              (io.grpc.stub.StreamObserver<aerium.BlockchainOuterClass.GetPublicKeyResponse>) responseObserver);
          break;
        case METHODID_GET_TX_POOL_CONTENT:
          serviceImpl.getTxPoolContent((aerium.BlockchainOuterClass.GetTxPoolContentRequest) request,
              (io.grpc.stub.StreamObserver<aerium.BlockchainOuterClass.GetTxPoolContentResponse>) responseObserver);
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
          getGetBlockMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              aerium.BlockchainOuterClass.GetBlockRequest,
              aerium.BlockchainOuterClass.GetBlockResponse>(
                service, METHODID_GET_BLOCK)))
        .addMethod(
          getGetBlockHashMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              aerium.BlockchainOuterClass.GetBlockHashRequest,
              aerium.BlockchainOuterClass.GetBlockHashResponse>(
                service, METHODID_GET_BLOCK_HASH)))
        .addMethod(
          getGetBlockHeightMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              aerium.BlockchainOuterClass.GetBlockHeightRequest,
              aerium.BlockchainOuterClass.GetBlockHeightResponse>(
                service, METHODID_GET_BLOCK_HEIGHT)))
        .addMethod(
          getGetBlockchainInfoMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              aerium.BlockchainOuterClass.GetBlockchainInfoRequest,
              aerium.BlockchainOuterClass.GetBlockchainInfoResponse>(
                service, METHODID_GET_BLOCKCHAIN_INFO)))
        .addMethod(
          getGetConsensusInfoMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              aerium.BlockchainOuterClass.GetConsensusInfoRequest,
              aerium.BlockchainOuterClass.GetConsensusInfoResponse>(
                service, METHODID_GET_CONSENSUS_INFO)))
        .addMethod(
          getGetAccountMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              aerium.BlockchainOuterClass.GetAccountRequest,
              aerium.BlockchainOuterClass.GetAccountResponse>(
                service, METHODID_GET_ACCOUNT)))
        .addMethod(
          getGetValidatorMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              aerium.BlockchainOuterClass.GetValidatorRequest,
              aerium.BlockchainOuterClass.GetValidatorResponse>(
                service, METHODID_GET_VALIDATOR)))
        .addMethod(
          getGetValidatorByNumberMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              aerium.BlockchainOuterClass.GetValidatorByNumberRequest,
              aerium.BlockchainOuterClass.GetValidatorResponse>(
                service, METHODID_GET_VALIDATOR_BY_NUMBER)))
        .addMethod(
          getGetValidatorAddressesMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              aerium.BlockchainOuterClass.GetValidatorAddressesRequest,
              aerium.BlockchainOuterClass.GetValidatorAddressesResponse>(
                service, METHODID_GET_VALIDATOR_ADDRESSES)))
        .addMethod(
          getGetPublicKeyMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              aerium.BlockchainOuterClass.GetPublicKeyRequest,
              aerium.BlockchainOuterClass.GetPublicKeyResponse>(
                service, METHODID_GET_PUBLIC_KEY)))
        .addMethod(
          getGetTxPoolContentMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              aerium.BlockchainOuterClass.GetTxPoolContentRequest,
              aerium.BlockchainOuterClass.GetTxPoolContentResponse>(
                service, METHODID_GET_TX_POOL_CONTENT)))
        .build();
  }

  private static abstract class BlockchainBaseDescriptorSupplier
      implements io.grpc.protobuf.ProtoFileDescriptorSupplier, io.grpc.protobuf.ProtoServiceDescriptorSupplier {
    BlockchainBaseDescriptorSupplier() {}

    @java.lang.Override
    public com.google.protobuf.Descriptors.FileDescriptor getFileDescriptor() {
      return aerium.BlockchainOuterClass.getDescriptor();
    }

    @java.lang.Override
    public com.google.protobuf.Descriptors.ServiceDescriptor getServiceDescriptor() {
      return getFileDescriptor().findServiceByName("Blockchain");
    }
  }

  private static final class BlockchainFileDescriptorSupplier
      extends BlockchainBaseDescriptorSupplier {
    BlockchainFileDescriptorSupplier() {}
  }

  private static final class BlockchainMethodDescriptorSupplier
      extends BlockchainBaseDescriptorSupplier
      implements io.grpc.protobuf.ProtoMethodDescriptorSupplier {
    private final java.lang.String methodName;

    BlockchainMethodDescriptorSupplier(java.lang.String methodName) {
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
      synchronized (BlockchainGrpc.class) {
        result = serviceDescriptor;
        if (result == null) {
          serviceDescriptor = result = io.grpc.ServiceDescriptor.newBuilder(SERVICE_NAME)
              .setSchemaDescriptor(new BlockchainFileDescriptorSupplier())
              .addMethod(getGetBlockMethod())
              .addMethod(getGetBlockHashMethod())
              .addMethod(getGetBlockHeightMethod())
              .addMethod(getGetBlockchainInfoMethod())
              .addMethod(getGetConsensusInfoMethod())
              .addMethod(getGetAccountMethod())
              .addMethod(getGetValidatorMethod())
              .addMethod(getGetValidatorByNumberMethod())
              .addMethod(getGetValidatorAddressesMethod())
              .addMethod(getGetPublicKeyMethod())
              .addMethod(getGetTxPoolContentMethod())
              .build();
        }
      }
    }
    return result;
  }
}
