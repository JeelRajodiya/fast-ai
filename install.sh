#!/bin/bash

# Colors
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
RED='\033[0;31m'
NC='\033[0m' # No Color

# Detect architecture
ARCH=$(uname -m)
case $ARCH in
    x86_64)
        GOARCH="amd64"
        ;;
    aarch64|arm64)
        GOARCH="arm64"
        ;;
    *)
        echo -e "${RED}Unsupported architecture: $ARCH${NC}"
        echo "This script supports x86_64 (amd64) and aarch64/arm64 only."
        exit 1
        ;;
esac

# Binary name based on architecture
BINARY_NAME="ai-linux-${GOARCH}"

# URL of the binary in your "ai" release
BINARY_URL="https://github.com/JeelRajodiya/fast-ai/releases/latest/download/${BINARY_NAME}"

# Destination path to place the binary
# Use ~/.local/bin if it exists, otherwise fall back to /usr/local/bin (requires sudo)
if [ -d "$HOME/.local/bin" ]; then
    DESTINATION="$HOME/.local/bin/ai"
    USE_SUDO=""
else
    DESTINATION="/usr/local/bin/ai"
    USE_SUDO="sudo"
fi

# Get the version of the latest release
VERSION=$(curl -sI https://github.com/JeelRajodiya/fast-ai/releases/latest | grep "location:" | awk -F "/" '{ print $NF }' | tr -d '\r')

# Get the size of the binary (follow redirects to get actual file size)
SIZE=$(curl -sIL $BINARY_URL | grep -i "^content-length:" | tail -1 | awk '{print $2}' | tr -d '\r')
if [ -n "$SIZE" ] && [ "$SIZE" -gt 0 ] 2>/dev/null; then
    SIZE_MB=$(awk "BEGIN {printf \"%.2f\", $SIZE/1048576}")
    echo -e "${YELLOW}Downloading ai binary version ${VERSION} for Linux ${GOARCH} (${SIZE_MB} MB) from GitHub release...${NC}"
else
    echo -e "${YELLOW}Downloading ai binary version ${VERSION} for Linux ${GOARCH} from GitHub release...${NC}"
fi

# Create destination directory if it doesn't exist
mkdir -p "$(dirname "$DESTINATION")"

# Download the binary
if $USE_SUDO curl -fsSL $BINARY_URL -o $DESTINATION; then
    echo -e "${GREEN}Download successful.${NC}"
else
    echo -e "${RED}Download failed. Please check your internet connection and try again.${NC}"
    exit 1
fi

echo -e "${YELLOW}Setting executable permissions...${NC}"
$USE_SUDO chmod +x $DESTINATION

echo -e "${GREEN}Installation complete.${NC}"
echo -e "You can run it using the command: ${GREEN}ai${NC}"

# Check if the installation directory is in PATH
if [[ ":$PATH:" != *":$(dirname "$DESTINATION"):"* ]]; then
    echo -e "${YELLOW}Warning: $(dirname "$DESTINATION") is not in your PATH.${NC}"
    if [ "$DESTINATION" = "$HOME/.local/bin/ai" ]; then
        echo -e "Add the following line to your ~/.bashrc or ~/.zshrc file:"
        echo -e "${GREEN}export PATH=\"\$HOME/.local/bin:\$PATH\"${NC}"
    fi
fi