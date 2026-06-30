#!/bin/bash
set -e

DOCLIENT_DIR="$(cd "$(dirname "$0")/../doclient" && pwd)"
SCRIPT_DIR="$(cd "$(dirname "$0")" && pwd)"

cd "$SCRIPT_DIR"

# Build doclient packages if needed
(cd "$DOCLIENT_DIR" && pnpm build)

# Install tsx if not present
if [ ! -d "node_modules/tsx" ]; then
  npm install --no-save --no-package-lock tsx 2>/dev/null
fi

# Ensure @doclient packages are resolvable via node_modules
mkdir -p node_modules/@doclient
for pkg in core renderer-go cli; do
  if [ ! -L "node_modules/@doclient/$pkg" ]; then
    ln -snf "$DOCLIENT_DIR/packages/$pkg" "node_modules/@doclient/$pkg"
  fi
done

# Run the doclient CLI
node "$DOCLIENT_DIR/packages/cli/dist/cli.js" --config "$SCRIPT_DIR/shopee-go.config.ts"
