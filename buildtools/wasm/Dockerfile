FROM  ghcr.io/swiftwasm/swift:5.9-focal

RUN apt-get update && apt-get install -y binaryen curl

ADD buildtools/wasm/version.txt version.txt
RUN mkdir -p /workspace && curl -L https://github.com/apple/swift-protobuf/archive/$(cat version.txt | awk '{$1=$1};1').tar.gz | tar -xz --strip-components 1 -C /workspace
WORKDIR /workspace

RUN swift build -c release --triple wasm32-unknown-wasi --product protoc-gen-swift

CMD ["cp", ".build/wasm32-unknown-wasi/release/protoc-gen-swift.wasm", "/out/"]
