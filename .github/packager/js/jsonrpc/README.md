# aerium-jsonrpc

JavaScript client for interacting with the [Aerium](https://aerium.network) blockchain via JSON-RPC.

## Installation

```bash
npm install aerium-jsonrpc
```

## Usage

```javascript
import AeriumOpenRPC from "aerium-jsonrpc";

const jsonrpcClient = new AeriumOpenRPC({
  transport: {
    type: "http",
    host: "127.0.0.1",
    port: 8545
  },
});

const blockchainInfo = await jsonrpcClient.aeriumBlockchainGetBlockchainInfo();
console.log(JSON.stringify(blockchainInfo, null, 2));
```
