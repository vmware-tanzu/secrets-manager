#!/usr/bin/env bash

CAST="669681.cast"

asciinema-edit quantize --range 0.35 $CAST > "$CAST-v1"
asciinema-edit quantize --range 0.1,0.34 "$CAST-v1" > "$CAST-v2"
asciinema play "$CAST-v2"
