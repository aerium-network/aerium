from pathlib import Path

from setuptools import find_packages, setup

setup(
    name="aerium-grpc",
    version="{{ VERSION }}",
    author="Aerium Development Team",
    author_email="info@aerium.network",
    url="https://aerium.network",
    description="Python client for interacting with the Aerium blockchain via gRPC",
    long_description=Path("README.md").read_text(encoding="utf-8"),
    long_description_content_type="text/markdown",
    license="MIT",
    packages=find_packages(),
    keywords=["aerium", "blockchain", "grpc"],
    install_requires=[
        "grpcio",
        "protobuf",
    ],
    classifiers=[
        "Development Status :: 5 - Production/Stable",
        "Intended Audience :: Developers",
        "Topic :: Software Development :: Build Tools",
        "Operating System :: OS Independent",
    ],
    python_requires=">=3.6",
)
