#!/bin/bash
echo "🛡️  Configurando GoHash-Auditor com UniPDF e go-echarts..."

# Inicializa o módulo
if [ ! -f "go.mod" ]; then
    go mod init gohash-auditor
fi

# Instalação das novas dependências
echo "📥 Instalando bibliotecas profissionais..."
go get github.com/unidoc/unipdf/v3
go get github.com/go-echarts/go-echarts/v2/charts
go get github.com/go-echarts/go-echarts/v2/opts

go mod tidy

# Compila o binário
echo "🛠️  Compilando..."
go build -o gohash-auditor main.go

echo "✅ Ambiente pronto!"
