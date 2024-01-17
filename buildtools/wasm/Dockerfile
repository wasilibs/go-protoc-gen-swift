FROM  ghcr.io/swiftwasm/swift:5.9-focal

RUN apt-get update && apt-get install -y binaryen curl


RUN mkdir -p /workspace && curl -L https://github.com/apple/swift-protobuf/archive/0e7015750b9bd7b80f28cf997ae77ac0e6d37f0b.tar.gz | tar -xz --strip-components 1 -C /workspace
WORKDIR /workspace

# Currently doesn't build but removing seems to not cause any problem
RUN rm Sources/SwiftProtobuf/AsyncMessageSequence.swift
RUN swift build -c release --triple wasm32-unknown-wasi --product protoc-gen-swift

CMD ["cp", ".build/wasm32-unknown-wasi/release/protoc-gen-swift.wasm", "/out/"]
