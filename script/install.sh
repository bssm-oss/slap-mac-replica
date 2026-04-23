#!/usr/bin/env bash
set -euo pipefail

repo="bssm-oss/slap-mac-replica"
asset="slap-mac-replica_darwin_arm64.tar.gz"
download_url="https://github.com/${repo}/releases/latest/download/${asset}"
install_dir="${SLAP_MAC_INSTALL_DIR:-/usr/local/bin}"
binary_name="slap-mac-replica"

if [[ "$(uname -s)" != "Darwin" ]]; then
  echo "slap-mac-replica is macOS-only." >&2
  exit 1
fi

if [[ "$(uname -m)" != "arm64" ]]; then
  echo "slap-mac-replica requires Apple Silicon (arm64)." >&2
  exit 1
fi

tmpdir="$(mktemp -d)"
trap 'rm -rf "$tmpdir"' EXIT

echo "Downloading ${download_url}"
curl -fsSL "$download_url" -o "$tmpdir/$asset"
tar -xzf "$tmpdir/$asset" -C "$tmpdir"
chmod +x "$tmpdir/$binary_name"

mkdir -p "$install_dir" 2>/dev/null || true
target="$install_dir/$binary_name"

if [[ -w "$install_dir" ]]; then
  install "$tmpdir/$binary_name" "$target"
else
  echo "Installing to ${target} with sudo."
  sudo install "$tmpdir/$binary_name" "$target"
fi

echo "Installed: $target"
"$target" version
echo
echo "Next:"
echo "  $target doctor"
echo "  sudo $target run"
