# 🛡️ GoHash-Auditor

>[!WARNING]
> **STATUS DO PROJETO: FASE BETA** 🧪

O **GoHash-Auditor** é uma ferramenta de auditoria de segurança de alto desempenho, escrita em Go, projetada para identificar rapidamente ameaças conhecidas em conjuntos de hashes MD5 de ficheiros APKs.

---

## 📍 Índice

1. [🛡️ Sobre o GoHash-Auditor](#-gohash-auditor)
2. [🚀 Funcionalidades Principais](#-funcionalidades)
3. [🛠️ Tecnologias Utilizadas](#-tecnologias-utilizadas)
4. [📋 Pré-requisitos](#-pré-requisitos)
5. [📦 Instalação e Módulos](#-instalação-e-configuração)
6. [🎮 Guia de Uso (Pasta Data)](#-como-utilizar)
7. [📂 Estrutura do Repositório](#-estrutura-de-pastas)
8. [⚙️ Configuração da Blacklist](#-configuração-da-blacklist)
9. [⚖️ Licença](#-licença)

---

## 🚀 Funcionalidades

* **Análise Automatizada:** Identificação instantânea de malwares (Trojans, Adwares, Spywares) via blacklist.
* **Organização Inteligente:** Leitura de dados e exportação de resultados centralizada na pasta `data`.
* **Relatórios Temporais:** Geração de PDFs com carimbo de data e hora no nome do arquivo (ex: `Relatorio_04-03-2026_08-15.pdf`).
* **Performance:** Processamento ultra-rápido utilizando as capacidades nativas de concorrência do Go.

## 🛠️ Tecnologias Utilizadas

* **Linguagem:** [Go (Golang)](https://go.dev/)
* **PDF Engine:** [UniPDF v3](https://github.com/unidoc/unipdf) - Geração de documentos profissionais.
* **Data Format:** CSV para entrada de grandes conjuntos de dados.

## 📋 Pré-requisitos

* **Go 1.25** ou superior instalado.
* **Chave de Licença UniDoc:** Necessária para o funcionamento da biblioteca UniPDF (obtenha uma conta gratuita em [unidoc.io](https://cloud.unidoc.io)).

## 📦 Instalação e Configuração

Para configurar o ambiente e garantir que todas as dependências estejam na versão correta, execute:

```bash
# 1. Clone o repositório
git clone [https://github.com/seu-usuario/gohash_auditor.git](https://github.com/seu-usuario/gohash_auditor.git)
cd gohash_auditor

# 2. Inicialize o módulo Go (caso não exista)
go mod init gohash-auditor 

# 3. Instale as dependências estáveis
go get [github.com/unidoc/unipdf/v3@v3.60.0](https://github.com/unidoc/unipdf/v3@v3.60.0)
go mod tidy

# 4. Configure sua licença no código
# Abra o arquivo main.go e insira sua chave na função init()
# license.SetMeteredKey("SUA_CHAVE_AQUI")
```

---

🎮 Como Utilizar
Prepare seus dados: Coloque o seu arquivo apks.csv dentro da pasta data/ (se a pasta não existir, o programa a criará).
Formato do CSV: um hash MD5 por linha.

Execute a ferramenta:

```bash
go run main.go
```

Resultado: O relatório será salvo em data/Relatorio_DD-MM-YYYY_HH-mm.pdf.

---

## 📂 Estrutura de Pastas

```plaintext
gohash_auditor/
├── data/               # Contém o input (apks.csv) e os relatórios gerados
├── main.go             # Código-fonte principal com a lógica de auditoria
├── go.mod              # Definição do módulo e dependências
├── go.sum              # Checksums das dependências
└── README.md           # Documentação do projeto
```

---

## ⚙️ Configuração da Blacklist

Para atualizar a base de malwares, edite o mapa `blacklist` no `main.go`:

```go
blacklist := map[string]string{
    "hash_md5_aqui": "Nome_da_Ameaça",
}
```

---

## 🛠️ Tecnologias Utilizadas

- UniPDF: Engine profissional para criação de documentos e gráficos vetoriais.

-  Go Standard Library: encoding/csv, os, io.

---

⚖️ Licença
Distribuído sob a licença MIT. Veja o arquivo LICENSE para mais detalhes.

---

*Desenvolvido por [lka1r0st0n](https://github.com/lka1r0st0n) para fins de estudo em Segurança da Informação.*

---
