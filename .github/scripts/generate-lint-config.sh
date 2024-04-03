#!/bin/bash

# Путь к файлу конфигурации
CONFIG_FILE=".github/workflows/lint-config.yml"

# Удаляем старый файл конфигурации, если он существует
rm -f $CONFIG_FILE

# Начало файла конфигурации
echo "name: golangci-lint" > $CONFIG_FILE
echo "on:" >> $CONFIG_FILE
echo " push:" >> $CONFIG_FILE
echo " pull_request:" >> $CONFIG_FILE
echo "" >> $CONFIG_FILE
echo "permissions:" >> $CONFIG_FILE
echo " contents: read" >> $CONFIG_FILE
echo " pull-requests: read" >> $CONFIG_FILE
echo "" >> $CONFIG_FILE
echo "jobs:" >> $CONFIG_FILE

# Ищем все директории с go.mod
for dir in $(find . -name "go.mod" -exec dirname {} \; | sort -u); do
 # Добавляем задачу для каждой директории
 echo " lint-${dir#./}:" >> $CONFIG_FILE
 echo "    name: lint in $dir" >> $CONFIG_FILE
 echo "    runs-on: ubuntu-latest" >> $CONFIG_FILE
 echo "    steps:" >> $CONFIG_FILE
 echo "      - uses: actions/checkout@v4" >> $CONFIG_FILE
 echo "      - uses: actions/setup-go@v5" >> $CONFIG_FILE
 echo "        with:" >> $CONFIG_FILE
 echo "          go-version: '1.22'" >> $CONFIG_FILE
 echo "          cache: false" >> $CONFIG_FILE
 echo "      - name: golangci-lint" >> $CONFIG_FILE
 echo "        uses: golangci/golangci-lint-action@v4" >> $CONFIG_FILE
 echo "        with:" >> $CONFIG_FILE
 echo "          version: v1.57.1" >> $CONFIG_FILE
 echo "          working-directory: $dir" >> $CONFIG_FILE
 echo "" >> $CONFIG_FILE
done