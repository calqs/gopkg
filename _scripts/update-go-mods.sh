#!/bin/sh
set -eu

LATEST_SCRIPT="${LATEST_SCRIPT:-./_scripts/latest-go.sh}"

if [ ! -x "$LATEST_SCRIPT" ]; then
  printf >&2 "error: %s not found or not executable\n" "$LATEST_SCRIPT"
  exit 1
fi

full_ver="$("$LATEST_SCRIPT")"
# major=${full_ver%%.*}
# rest=${full_ver#*.}
# minor=${rest%%.*}
gomod_ver="$full_ver"

update_file() {
  f="$1"
  tmp="$f.tmp.$$"

  if grep -qE '^go[[:space:]][0-9]+' "$f"; then
    awk -v v="$gomod_ver" '
      BEGIN{replaced=0}
      $1=="go" && $2 ~ /^[0-9]+\.[0-9]+(\.[0-9]+)?$/ && !replaced { print "go " v; replaced=1; next }
      { print }
    ' "$f" >"$tmp"
  else
    # Insert after the `module` line if no 'go' line present
    awk -v v="$gomod_ver" '
      BEGIN{inserted=0}
      {
        print
        if (!inserted && $1=="module") {
          print "go " v
          inserted=1
        }
      }
    ' "$f" >"$tmp"
  fi

  if cmp -s "$f" "$tmp"; then
    rm -f "$tmp"
  else
    mv "$tmp" "$f"
    printf "updated %s -> go %s\n" "$f" "$gomod_ver"
  fi
}

find . -maxdepth 2 -type f -name go.mod | while IFS= read -r mod; do
  update_file "$mod"
done
