#!/bin/sh
set -e

REPO="0xmanny/lockin"
INSTALL_DIR="/usr/local/bin"
BINARY_NAME="lockin"

main() {
    OS=$(uname -s | tr '[:upper:]' '[:lower:]')
    ARCH=$(uname -m)

    if [ "$OS" != "darwin" ]; then
        echo "Error: lockin currently only supports macOS."
        exit 1
    fi

    case "$ARCH" in
        x86_64)  ARCH="amd64" ;;
        arm64)   ARCH="arm64" ;;
        aarch64) ARCH="arm64" ;;
        *)
            echo "Error: unsupported architecture: $ARCH"
            exit 1
            ;;
    esac

    LATEST=$(curl -fsSL "https://api.github.com/repos/${REPO}/releases/latest" | grep '"tag_name"' | sed -E 's/.*"([^"]+)".*/\1/')

    if [ -z "$LATEST" ]; then
        echo "Error: could not determine latest release."
        exit 1
    fi

    DOWNLOAD_URL="https://github.com/${REPO}/releases/download/${LATEST}/lockin-${OS}-${ARCH}"

    echo "Installing lockin ${LATEST} (${OS}/${ARCH})..."

    TMP=$(mktemp -d)
    trap 'rm -rf "$TMP"' EXIT

    curl -fsSL "$DOWNLOAD_URL" -o "${TMP}/${BINARY_NAME}"
    chmod +x "${TMP}/${BINARY_NAME}"

    if [ -w "$INSTALL_DIR" ]; then
        mv "${TMP}/${BINARY_NAME}" "${INSTALL_DIR}/${BINARY_NAME}"
    else
        echo "Installing to ${INSTALL_DIR} (requires sudo)..."
        sudo mv "${TMP}/${BINARY_NAME}" "${INSTALL_DIR}/${BINARY_NAME}"
    fi

    PLIST_PATH="/Library/LaunchDaemons/sh.lockin.cleanup.plist"
    if [ ! -f "$PLIST_PATH" ]; then
        cat > "${TMP}/sh.lockin.cleanup.plist" <<'PLIST'
<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
<plist version="1.0">
<dict>
    <key>Label</key>
    <string>sh.lockin.cleanup</string>
    <key>ProgramArguments</key>
    <array>
        <string>/usr/local/bin/lockin</string>
        <string>stop</string>
    </array>
    <key>RunAtLoad</key>
    <true/>
</dict>
</plist>
PLIST
        sudo cp "${TMP}/sh.lockin.cleanup.plist" "$PLIST_PATH"
        sudo launchctl load "$PLIST_PATH" 2>/dev/null
    fi

    echo ""
    echo "  lockin installed successfully!"
    echo ""
    echo "  Get started:"
    echo "    lockin             # block distractions"
    echo "    lockin status      # see what's blocked"
    echo "    lockin config      # change the blocklist"
    echo "    lockin stop        # unblock everything"
    echo ""
    echo "  Config: ~/.lockin/config.yaml"
    echo ""
}

main
