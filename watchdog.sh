#!/bin/bash

cd "$(dirname "$0")"  # Esto fuerza a que el script se mueva a su propia carpeta
PROCESS_NAME="sentinel_pro.exe"

echo "📡 Sentinel Windows activo. Vigilando: $PROCESS_NAME"

while true; do
    # Usamos tasklist para Windows
    if ! tasklist | grep -q "$PROCESS_NAME"; then
        echo "⚠️ Servidor caído. Reiniciando..."
        ./$PROCESS_NAME &
        sleep 2
    fi
    sleep 5
done