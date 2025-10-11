# aerium-jsonrpc

Python client for interacting with the [Aerium](https://aerium.network) blockchain via JSON-RPC.

## Installation

```bash
pip install aerium-jsonrpc
```

## Usage

```python
import asyncio
from aerium_jsonrpc.client import AeriumOpenRPCClient


async def main():
    client = AeriumOpenRPCClient(
        headers={},
        client_url="http://127.0.0.1:8545"
    )

    blockchain_info = await client.aerium.blockchain.get_blockchain_info()
    print(blockchain_info)


if __name__ == "__main__":
    asyncio.run(main())
```
