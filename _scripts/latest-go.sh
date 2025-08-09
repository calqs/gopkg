#!/bin/sh
set -eu

fetch() {
  if command -v curl >/dev/null 2>&1; then
    curl -fsSL "https://go.dev/VERSION?m=text"
  elif command -v wget >/dev/null 2>&1; then
    wget -qO- "https://go.dev/VERSION?m=text"
  else
    printf >&2 "error: need curl or wget\n"
    exit 1
  fi
}

line="$(fetch)"
set -- $line
ver="$1"
ver="${ver#go}"
case "$ver" in
  [0-9]*.[0-9]*|[0-9]*.[0-9]*.[0-9]*) printf "%s\n" "$ver" ;;
  *) printf >&2 "error: unexpected response: %s\n" "$line"; exit 1 ;;
esac
