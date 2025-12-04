#!/bin/bash

# go-projo installation script

set -e

echo "🚀 Installing go-projo..."
echo ""

# Check if Go is installed
if ! command -v go &> /dev/null; then
    echo "❌ Error: Go is not installed"
    echo "Please install Go from https://go.dev/dl/"
    exit 1
fi

echo "✓ Go version: $(go version)"
echo ""

# Build the binary
echo "📦 Building go-projo..."
go build -o go-projo .

if [ $? -ne 0 ]; then
    echo "❌ Build failed"
    exit 1
fi

echo "✓ Build successful"
echo ""

# Determine installation directory
GOPATH_BIN="$(go env GOPATH)/bin"

# Create GOPATH/bin if it doesn't exist
if [ ! -d "$GOPATH_BIN" ]; then
    echo "📁 Creating $GOPATH_BIN..."
    mkdir -p "$GOPATH_BIN"
fi

# Copy binary
echo "📥 Installing to $GOPATH_BIN..."
cp go-projo "$GOPATH_BIN/"
chmod +x "$GOPATH_BIN/go-projo"

echo "✓ Installed successfully"
echo ""

# Check if GOPATH/bin is in PATH
if [[ ":$PATH:" != *":$GOPATH_BIN:"* ]]; then
    echo "⚠️  Warning: $GOPATH_BIN is not in your PATH"
    echo ""
    echo "Add it to your shell configuration file (~/.bashrc, ~/.zshrc, etc.):"
    echo ""
    echo "    export PATH=\"\$PATH:\$(go env GOPATH)/bin\""
    echo ""
    echo "Then reload your shell:"
    echo "    source ~/.bashrc  # or ~/.zshrc"
    echo ""
else
    echo "✓ $GOPATH_BIN is in your PATH"
    echo ""
fi

# Verify installation
if command -v go-projo &> /dev/null; then
    echo "✅ Installation complete!"
    echo ""
    echo "Version: $(go-projo version)"
    echo ""
    echo "Try it out:"
    echo "  go-projo help"
    echo "  go-projo gen -name demo -module github.com/user/demo -type api"
else
    echo "⚠️  Installation complete, but go-projo is not in PATH yet"
    echo "You may need to reload your shell or add GOPATH/bin to PATH"
fi

echo ""
echo "📚 Documentation:"
echo "  README.md - Full documentation"
echo "  INSTALL.md - Installation guide"
echo ""
