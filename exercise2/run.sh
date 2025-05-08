#!/bin/sh

# Скрипт для сборки и запуска конвертера валют (exercise2)
echo "Сборка и запуск конвертера валют..."
go build -o converter
./converter