#!/bin/sh

set -e

REPO="soft4dev/clonei"  
BIN_NAME="clonei"             

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

# Detect OS and architecture
get_os() {
    os=$(uname -s)
    case "$os" in
        Darwin) echo "Darwin" ;;
        Linux) echo "Linux" ;;
        *) 
            printf "${RED}Unsupported OS: $os${NC}\n" >&2
            printf "This script only supports macOS and Linux.\n" >&2
            printf "For Windows, use: irm https://raw.githubusercontent.com/${REPO}/main/install.ps1 | iex\n" >&2
            exit 1
            ;;
    esac
}

get_arch() {
    arch=$(uname -m)
    case "$arch" in
        x86_64|amd64) echo "x86_64" ;;
        aarch64|arm64) echo "arm64" ;;
        i386|i686) echo "i386" ;;
        armv7l) echo "armv7" ;;
        *) 
            printf "${RED}Unsupported architecture: $arch${NC}\n" >&2
            exit 1
            ;;
    esac
}

# Get the latest release version
get_latest_version() {
    echo "$(curl -s "https://api.github.com/repos/$REPO/releases/latest" | grep '"tag_name":' | sed -E 's/.*"([^"]+)".*/\1/')"
}

# Download and extract the release
install_binary() {
    os=$(get_os)
    arch=$(get_arch)
    version=$(get_latest_version)

    if [ -z "$version" ]; then
        printf "${RED}Error: Could not fetch latest version${NC}\n" >&2
        exit 1
    fi

    printf "${GREEN}Installing ${BIN_NAME} ${version}...${NC}\n"

    # Construct download URL
    archive_name="${BIN_NAME}_${os}_${arch}.tar.gz"
    download_url="https://github.com/${REPO}/releases/download/${version}/${archive_name}"
    
    # Create temporary directory
    tmp_dir=$(mktemp -d)
    trap "rm -rf $tmp_dir" EXIT

    printf "${YELLOW}Downloading from ${download_url}...${NC}\n"

    # Download the archive
    if command -v curl >/dev/null 2>&1; then
        curl -fsSL "$download_url" -o "$tmp_dir/$archive_name"
    elif command -v wget >/dev/null 2>&1; then
        wget -q "$download_url" -O "$tmp_dir/$archive_name"
    else
        printf "${RED}Error: curl or wget is required${NC}\n" >&2
        exit 1
    fi

    # Extract the archive
    printf "${YELLOW}Extracting...${NC}\n"
    tar -xzf "$tmp_dir/$archive_name" -C "$tmp_dir"

    # Determine install directory
    bin_dir="${BIN_DIR:-$HOME/.local/bin}"

    # Create bin directory if it doesn't exist
    mkdir -p "$bin_dir"

    # Move binary to install directory
    printf "${YELLOW}Installing to ${bin_dir}...${NC}\n"
    mv "$tmp_dir/$BIN_NAME" "$bin_dir/$BIN_NAME"
    chmod +x "$bin_dir/$BIN_NAME"

    printf "${GREEN}✓ ${BIN_NAME} ${version} installed successfully!${NC}\n\n"

    # Check if bin_dir is in PATH
    case ":$PATH:" in
        *":$bin_dir:"*) 
            printf "${GREEN}Run '${BIN_NAME} --help' to get started${NC}\n"
            ;;
        *)
            printf "${YELLOW}Adding ${bin_dir} to PATH...${NC}\n"
            
            # Detect shell and profile file
            shell_profile=""
            if [ -n "$BASH_VERSION" ]; then
                shell_profile="$HOME/.bashrc"
            elif [ -n "$ZSH_VERSION" ]; then
                shell_profile="$HOME/.zshrc"
            elif [ -n "$FISH_VERSION" ]; then
                shell_profile="$HOME/.config/fish/config.fish"
            else
                # Try to detect from SHELL environment variable
                case "$SHELL" in
                    */bash) shell_profile="$HOME/.bashrc" ;;
                    */zsh) shell_profile="$HOME/.zshrc" ;;
                    */fish) shell_profile="$HOME/.config/fish/config.fish" ;;
                    *) shell_profile="$HOME/.profile" ;;
                esac
            fi

            # Add to PATH in shell profile
            if [ -n "$shell_profile" ]; then
                # Check if PATH export already exists in profile
                if ! grep -q "export PATH=.*${bin_dir}" "$shell_profile" 2>/dev/null; then
                    echo "" >> "$shell_profile"
                    echo "# Added by ${BIN_NAME} installer" >> "$shell_profile"
                    echo "export PATH=\"\$PATH:${bin_dir}\"" >> "$shell_profile"
                    printf "${GREEN}✓ Added to ${shell_profile}${NC}\n"
                else
                    printf "${YELLOW}PATH already configured in ${shell_profile}${NC}\n"
                fi
            fi

            printf "\n${YELLOW}Please restart your terminal or run:${NC}\n"
            printf "  source ${shell_profile}\n\n"
            printf "Then run '${BIN_NAME} --help' to get started\n"
            ;;
    esac
}

# Main execution
install_binary