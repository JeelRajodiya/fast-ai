#!/bin/bash

# Colors
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

# URL of the binary in your "ai" release
BINARY_URL="https://github.com/JeelRajodiya/fast-ai/releases/latest/download/ai"

# Destination path to place the binary
DESTINATION="/usr/local/bin/ai"

echo -e "${YELLOW}Downloading ai binary from GitHub release...${NC}"
sudo curl -L $BINARY_URL -o $DESTINATION

echo -e "${YELLOW}Setting executable permissions...${NC}"
sudo chmod +x $DESTINATION

echo -e "${GREEN}Installation complete.${NC}"
echo -e "You can run it using the command: ${GREEN}ai${NC}"