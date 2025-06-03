#!/bin/bash
RED=$(tput setaf 1)
GREEN=$(tput setaf 2)
RESET=$(tput sgr0)

echo "${RED}🛑 Остановка price-backend...${RESET}"
pkill -f "bin/price-backend" || echo "🔍 Процесс не найден."
echo "${GREEN}✅ Остановлено${RESET}"
