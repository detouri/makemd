#!/usr/bin/env sh

set -eu

coverage="${1:-0}"
output="${2:-.github/badges/coverage.svg}"

coverage="$(printf '%s' "$coverage" | tr -d '%')"

color="#e05d44"
awk "BEGIN { exit !($coverage >= 80) }" && color="#4c1" || true
if [ "$color" = "#e05d44" ]; then
    awk "BEGIN { exit !($coverage >= 60) }" && color="#dfb317" || true
fi

mkdir -p "$(dirname "$output")"

cat > "$output" <<EOF
<svg xmlns="http://www.w3.org/2000/svg" width="118" height="20" role="img" aria-label="coverage: ${coverage}%">
  <linearGradient id="s" x2="0" y2="100%">
    <stop offset="0" stop-color="#bbb" stop-opacity=".1"/>
    <stop offset="1" stop-opacity=".1"/>
  </linearGradient>
  <clipPath id="r">
    <rect width="118" height="20" rx="3" fill="#fff"/>
  </clipPath>
  <g clip-path="url(#r)">
    <rect width="63" height="20" fill="#555"/>
    <rect x="63" width="55" height="20" fill="${color}"/>
    <rect width="118" height="20" fill="url(#s)"/>
  </g>
  <g fill="#fff" text-anchor="middle" font-family="Verdana,Geneva,DejaVu Sans,sans-serif" font-size="11">
    <text x="32.5" y="15" fill="#010101" fill-opacity=".3">coverage</text>
    <text x="32.5" y="14">coverage</text>
    <text x="89.5" y="15" fill="#010101" fill-opacity=".3">${coverage}%</text>
    <text x="89.5" y="14">${coverage}%</text>
  </g>
</svg>
EOF
