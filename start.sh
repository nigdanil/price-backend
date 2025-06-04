#!/bin/bash
set -e

RED=$(tput setaf 1)
GREEN=$(tput setaf 2)
BLUE=$(tput setaf 4)
RESET=$(tput sgr0)

log() {
    echo "${BLUE}[$(date +"%H:%M:%S")]${RESET} $1"
}

success() {
    echo "${GREEN}✅ $1${RESET}"
}

log "📦 Обновляем зависимости..."
go mod tidy

log "🔨 Сборка проекта..."
mkdir -p bin
go build -o bin/price-backend main.go

log "🛑 Завершаем старый процесс (если есть)..."
pkill -f "bin/price-backend" || true

log "🚀 Запускаем приложение..."
nohup ./bin/price-backend > backend.log 2>&1 &

success "Проект price-backend запущен!"
