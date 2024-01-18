FROM  ghcr.io/swiftwasm/swift:5.9-focal

RUN apt-get update && apt-get install -y binaryen curl


RUN mkdir -p /workspace && curl -L https://github.com/thomasvl/swift-protobuf/archive/57b7cb562fd889db7a4b5559cbc647d54e5b132c.tar.gz | tar -xz --strip-components 1 -C /workspace
WORKDIR /workspace

RUN swift build -c release --triple wasm32-unknown-wasi --product protoc-gen-swift

CMD ["cp", ".build/wasm32-unknown-wasi/release/protoc-gen-swift.wasm", "/out/"]
