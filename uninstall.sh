#!/bin/bash

# Colors
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
RED='\033[0;31m'
NC='\033[0m' # No Color

# Check both possible locations
LOCAL_BIN="$HOME/.local/bin/ai"
SYSTEM_BIN="/usr/local/bin/ai"

if [ -f "$LOCAL_BIN" ]; then
    echo -e "${YELLOW}Removing ai from ~/.local/bin...${NC}"
    rm "$LOCAL_BIN"
    echo -e "${GREEN}Uninstallation complete. 'ai' binary has been removed from ~/.local/bin/${NC}"
elif [ -f "$SYSTEM_BIN" ]; then
    echo -e "${YELLOW}Removing ai from /usr/local/bin...${NC}"
    sudo rm "$SYSTEM_BIN"
    echo -e "${GREEN}Uninstallation complete. 'ai' binary has been removed from /usr/local/bin/${NC}"
else
    echo -e "${RED}Error: 'ai' binary not found in either ~/.local/bin or /usr/local/bin${NC}"
    exit 1
fi