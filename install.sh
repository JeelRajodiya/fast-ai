#!/bin/bash

# Colors
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

# URL of the binary in your "ai" release
BINARY_URL="https://github.com/JeelRajodiya/fast-ai/releases/latest/download/ai"

# Destination path to place the binary
DESTINATION="/usr/local/bin/ai"

# Get the version of the latest release
VERSION=$(curl -sI https://github.com/JeelRajodiya/fast-ai/releases/latest | grep "location:" | awk -F "/" '{ print $NF }' | tr -d '\r')

echo -e "${YELLOW}Downloading ai binary version ${VERSION} from GitHub release...${NC}"
# Get terminal width and set progress bar width to 50%
WIDTH=$(tput cols)
PROGRESS_BAR_WIDTH=$((WIDTH / 2))
(stty cols $PROGRESS_BAR_WIDTH; sudo curl -L# $BINARY_URL -o $DESTINATION)

echo -e "${YELLOW}Setting executable permissions...${NC}"
sudo chmod +x $DESTINATION

echo -e "${GREEN}Installation complete.${NC}"
echo -e "You can run it using the command: ${GREEN}ai${NC}"