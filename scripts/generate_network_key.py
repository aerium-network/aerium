#!/usr/bin/env python3
"""
Network Key Generator for Aerium

This script generates a libp2p-compatible Ed25519 private key in the same format
as the Go implementation in network/network.go (loadOrCreateKey function).

The key is:
1. Generated as an Ed25519 key pair
2. Marshaled in libp2p protobuf format
3. Hex-encoded
4. Saved to a file

Usage:
    python3 generate_network_key.py [output_file]

If no output file is specified, the key is printed to stdout.
"""

import sys
import os
from cryptography.hazmat.primitives.asymmetric import ed25519
from cryptography.hazmat.primitives import serialization


def marshal_libp2p_private_key(private_key: ed25519.Ed25519PrivateKey) -> bytes:
    """
    Marshal an Ed25519 private key in libp2p protobuf format.

    The libp2p format is:
    - 1 byte: key type (0x01 for Ed25519)
    - remaining bytes: the raw private key (64 bytes for Ed25519)

    Ed25519 private key in libp2p includes both the 32-byte seed and 32-byte public key.
    """
    # Get the raw private key bytes (32 bytes seed)
    private_bytes = private_key.private_bytes(
        encoding=serialization.Encoding.Raw,
        format=serialization.PrivateFormat.Raw,
        encryption_algorithm=serialization.NoEncryption()
    )

    # Get the public key bytes (32 bytes)
    public_key = private_key.public_key()
    public_bytes = public_key.public_bytes(
        encoding=serialization.Encoding.Raw,
        format=serialization.PublicFormat.Raw
    )

    # Combine: private_seed (32 bytes) + public_key (32 bytes) = 64 bytes
    ed25519_key = private_bytes + public_bytes

    # Protobuf encoding for libp2p:
    # Field 1 (Type): varint, value 1 (Ed25519)
    # Field 2 (Data): length-delimited, value is the 64-byte key

    # Protobuf field encoding:
    # Field 1: tag = (1 << 3) | 0 = 0x08, value = 1
    # Field 2: tag = (2 << 3) | 2 = 0x12, length = 64 (0x40), data = 64 bytes

    key_type = 1  # Ed25519
    protobuf_data = bytes([0x08, key_type, 0x12, len(ed25519_key)]) + ed25519_key

    return protobuf_data


def generate_network_key() -> str:
    """
    Generate a new Ed25519 private key and return it as a hex-encoded string
    in libp2p format.
    """
    # Generate Ed25519 key pair
    private_key = ed25519.Ed25519PrivateKey.generate()

    # Marshal in libp2p format
    marshaled_key = marshal_libp2p_private_key(private_key)

    # Hex encode
    hex_key = marshaled_key.hex()

    return hex_key


def load_key_from_file(path: str) -> str:
    """Load a hex-encoded key from a file."""
    if not os.path.exists(path):
        return None

    with open(path, 'r') as f:
        return f.read().strip()


def save_key_to_file(path: str, hex_key: str):
    """Save a hex-encoded key to a file without any trailing newline."""
    # Write in binary mode to ensure no newline is added
    with open(path, 'wb') as f:
        f.write(hex_key.encode('ascii'))


def main():
    if len(sys.argv) > 1:
        output_file = sys.argv[1]

        # Check if file already exists
        if os.path.exists(output_file):
            print(f"Key file already exists: {output_file}")
            existing_key = load_key_from_file(output_file)
            print(f"Existing key: {existing_key}")
            response = input("Generate new key and overwrite? (y/N): ")
            if response.lower() != 'y':
                print("Aborted.")
                return

        # Generate new key
        hex_key = generate_network_key()

        # Save to file
        save_key_to_file(output_file, hex_key)
        print(f"Network key generated and saved to: {output_file}")
        print(f"Key: {hex_key}")
    else:
        # No file specified, print to stdout
        hex_key = generate_network_key()
        print(f"Generated network key (hex-encoded):")
        print(hex_key)
        print("\nTo save to a file, run:")
        print(f"  python3 {sys.argv[0]} <output_file>")


if __name__ == "__main__":
    main()
