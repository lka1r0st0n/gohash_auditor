# 🛡️ GoHash-Auditor

>[!WARNING]
> **STATUS DO PROJETO: FASE BETA** 🧪

O **GoHash-Auditor** é uma ferramenta de linha de comando (CLI) desenvolvida em Go para auditoria técnica de integridade de arquivos APK. O sistema processa grandes volumes de dados via CSV, confronta hashes contra uma blacklist de malwares e gera um **Relatório Executivo em PDF** com visual de Dashboard.

---

- [🛡️ Sobre o Projeto](#-gohash-auditor)
- [🌟 Principais Funcionalidades](#-principais-funcionalidades)
- [📊 Estrutura do Relatório](#-estrutura-do-relatório)
- [🚀 Como Executar](#-como-executar)
  - [Pré-requisitos](#pré-requisitos)
  - [📦 Instalação e Configuração](#-instalação-e-configuração)
- [🎮 Como Utilizar](#-como-utilizar)
- [📂 Estrutura de Pastas](#-estrutura-de-pastas)
- [⚙️ Configuração da Blacklist](#-configuração-da-blacklist)
- [🛠️ Tecnologias Utilizadas](#-tecnologias-utilizadas)
- [⚖️ Licença](#-licença)

---

## 🌟 Principais Funcionalidades

- **Processamento de Alta Performance**: Leitura eficiente de arquivos CSV contendo hashes de aplicações.
- **Detecção de Ameaças**: Confronto automático com base de dados de malwares conhecidos (Trojan, Adware, Spyware).
- **Relatórios Dinâmicos**: Geração de PDFs profissionais com timestamp no nome do arquivo (ex: `relatorio_dashboard_04-03-2026_09-24-07.pdf`).
- **Interface Visual de Dashboard**: Relatórios com cabeçalho institucional azul marinho, cards de indicadores coloridos e lista detalhada de ameaças.

---

## 📊 Estrutura do Relatório

O relatório gerado é dividido em seções lógicas para facilitar a tomada de decisão:
1. **Cabeçalho Executivo**: Identificação clara do software de auditoria com título centralizado.
2. **Cards de Estatísticas**: 
   - **Ameaças**: Destacadas em vermelho para atenção imediata (Ex: 3 ameaças encontradas).
   - **Arquivos Seguros**: Destacados em verde, confirmando a integridade (Ex: 1212 seguros).
   - **Total Analisado**: Panorama geral do volume de dados processados (Ex: 1215 total).
   - **Data da Análise**: Registro temporal da execução da auditoria.
3. **Detalhamento Técnico**: Lista completa de hashes comprometidos, tipos de infecção e status da ameaça.

---

## 🚀 Como Executar

### Pré-requisitos
- Go 1.16+
- Chave de licença UniPDF (configurada no arquivo `.env` ou via `license.SetMeteredKey`)

---

## 📦 Instalação e Configuração

Para configurar o ambiente e garantir que todas as dependências estejam na versão correta (v3.60.0), execute:

```bash
# 1. Clone o repositório
git clone [https://github.com/lka1r0st0n/gohash_auditor.git](https://github.com/lka1r0st0n/gohash_auditor.git)
cd gohash_auditor

# 2. Instale as dependências estáveis (Compatíveis com o layout atual)
go get [github.com/unidoc/unipdf/v3@v3.60.0](https://github.com/unidoc/unipdf/v3@v3.60.0)
go mod tidy
```

---

## 🎮 Como Utilizar

Prepare seus dados: Coloque o seu arquivo apks.csv dentro da pasta data/. O formato deve ser um hash MD5 por linha.

Execute a ferramenta:

```Bash
go run main.go
```

Resultado: O relatório será salvo automaticamente em ```data/reports/relatorio_dashboard_DD-MM-YYYY_HH-MM-SS.pdf```.

---

## 📂 Estrutura de Pastas

```plaintext
gohash_auditor/
├── data/
│   ├── apks.csv          # Arquivo de entrada com os hashes
│   └── reports/          # Pasta de saída dos relatórios PDF dinâmicos
├── main.go               # Código-fonte com lógica de Dashboard e Auditoria
├── .env                  # Variáveis de ambiente (Chave de Licença)
├── go.mod                # Módulo Go (travado na UniPDF v3.60.0)
└── README.md             # Documentação do projeto
```

---

## ⚙️ Configuração da Blacklist
Para atualizar a base de malwares, edite o mapa blacklist dentro do main.go. Atualmente, o sistema identifica ameaças como:

- Trojan.Android.Agent
- Adware.AirPush.B
- Spyware.Stealer.X

```Go
blacklist := map[string]string{
    "hash_md5_aqui": "Nome_da_Ameaça",
}
```

---

## 🛠️ Tecnologias Utilizadas
UniPDF (v3.60.0): Engine profissional para criação de documentos e gráficos vetoriais.

Go Standard Library: encoding/csv, os, io, time.

Godotenv: Para carregamento seguro de chaves de API.

---

## ⚖️ Licença

Distribuído sob a licença MIT.

---

Desenvolvido por [lka1r0st0n](https://github.com/lka1r0st0n) para fins de estudo em Segurança da Informação.

---
