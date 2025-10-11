# aerium-grpc

JavaScript client for interacting with the [Aerium](https://aerium.network) blockchain via gRPC.

## Installation

```bash
npm install aerium-grpc
```

## Usage

```javascript
import grpc from '@grpc/grpc-js';
import blockchain_pb from "aerium-grpc/blockchain_pb.js";
import blockchain_grpc_pb from "aerium-grpc/blockchain_grpc_pb.js";

const client = new blockchain_grpc_pb.BlockchainClient(
  "127.0.0.1:50051",
  grpc.credentials.createInsecure()
);

const request = new blockchain_pb.GetBlockchainInfoRequest();
client.getBlockchainInfo(request, (err, response) => {
  if (err) {
    console.log(err);
  } else {
    console.log(response.toObject());
  }
});

```
