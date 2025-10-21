#!/bin/bash

# Colors
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

# URL of the binary in your "ai" release
BINARY_URL="https://github.com/JeelRajodiya/fast-ai/releases/latest/download/ai"

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
    echo -e "${YELLOW}Downloading ai binary version ${VERSION} (${SIZE_MB} MB) from GitHub release...${NC}"
else
    echo -e "${YELLOW}Downloading ai binary version ${VERSION} from GitHub release...${NC}"
fi

# Create destination directory if it doesn't exist
mkdir -p "$(dirname "$DESTINATION")"

$USE_SUDO curl -fsSL  $BINARY_URL -o $DESTINATION

echo -e "${YELLOW}Setting executable permissions...${NC}"
$USE_SUDO chmod +x $DESTINATION

echo -e "${GREEN}Installation complete.${NC}"
echo -e "You can run it using the command: ${GREEN}ai${NC}"