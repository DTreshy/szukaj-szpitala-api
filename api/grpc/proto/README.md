# gRPC protocol contract

## Prerequisities

1. Install unzip and wget

    ```console
    sudo apt install unzip wget
    ```

2. Install protoc compiler. Check if output of last command is libprotoc 3.18.1

    ```console
    wget https://github.com/protocolbuffers/protobuf/releases/download/v3.18.1/protoc-3.18.1-linux-x86_64.zip
    unzip protoc-3.18.1-linux-x86_64.zip -d $HOME/.local
    rm protoc-3.18.1-linux-x86_64.zip
    export PATH="$PATH:$HOME/.local/bin"
    protoc --version
    ```

3. Install go specific plugins

    ```console
    go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28.1
    go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
    go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@v2.11.0
    ```

4. Ensure all three tools are in your path:

    ```console
    protoc-gen-go-grpc[.exe]
    protoc-gen-go[.exe]
    protoc[.exe]
    protoc-gen-grpc-gateway[.exe]
    ```

5. If something goes wrong try:

    ```console
    export PATH="$PATH:$(go env GOPATH)/bin"
    ```

6. There may be also some errors, connected with ownership of catalogoues where

    some files are installed. If you see something like:
    ```get the error: google/protobuf/descriptor.proto: File not found```

    try:
    ```sudo chown [user] /usr/local/bin/protoc```
    ```sudo chown -R [user] /usr/local/include/google```

    If you still struggle with such error, there is issue connected with it:

    <https://github.com/grpc-ecosystem/grpc-gateway/issues/422>

7. Execute following command to rebuild protocol stubs (only needed if protocol described in `*.proto` files has changed):

    ```console
    task proto
    ```