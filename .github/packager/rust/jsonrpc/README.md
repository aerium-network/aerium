# aerium-jsonrpc

Rust client for interacting with the [Aerium](https://aerium.network) blockchain via JSON-RPC.

## Installation

```bash
cargo add aerium-jsonrpc
```

## Usage

```rust
use jsonrpsee::http_client::HttpClient;
use aerium_jsonrpc::aerium::AeriumOpenRPC;

#[tokio::main]
async fn main() {
    let client = HttpClient::builder().build("http://127.0.0.1:8545").unwrap();
    let rpc: AeriumOpenRPC<HttpClient> = AeriumOpenRPC::new(client);

    let info = rpc.aerium_blockchain_get_blockchain_info().await.unwrap();
    println!("get_blockchain_info Response: {:?}", info);
}
```
