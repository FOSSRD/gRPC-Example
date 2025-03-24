#!/bin/bash

set -e # Exit immediately if a command exits with a non-zero status

PROTO_FILE="converter.proto"
IMAGER_DIR="imager"
SERVIT_DIR="servit"

echo "Building Python client..."
cd "$SERVIT_DIR"
if [ ! -d "venv" ]; then
  python3 -m venv venv
fi
source venv/bin/activate
pip install --upgrade pip
pip install -r requirements.txt
python -m grpc_tools.protoc -I=.. --python_out=. --pyi_out=. --grpc_python_out=. ../"$PROTO_FILE"
deactivate
echo "Python client build finished."

echo "Building Go server..."
cd "../$IMAGER_DIR"
go mod tidy
protoc -I=.. --go_out=. --go-grpc_out=. ../"$PROTO_FILE"
echo "Go server build finished."

echo "Build completed successfully!"
