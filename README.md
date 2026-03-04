# 🛡️ GoHash-Auditor

>[!WARNING]
> **STATUS DO PROJETO: FASE BETA** 🧪

O **GoHash-Auditor** é uma ferramenta de auditoria de segurança de alto desempenho, escrita em Go, projetada para identificar rapidamente ameaças conhecidas em conjuntos de hashes MD5 de ficheiros APKs.

---

## 📍 Índice

1. [Sobre o Projeto](#-gohash-auditor)
2. [Tecnologias Utilizadas](#-tecnologias-utilizadas)
3. [Funcionalidades Principais](#-funcionalidades-principais)
4. [Pré-requisitos](#-pré-requisitos)
5. [Instalação](#-instalação)
6. [Como Utilizar](#-como-utilizar)
7. [Configuração](#-configuração)
8. [Licença](#-licença)

---

## 🚀 Funcionalidades Principais

* **Alta Performance:** Processamento eficiente de arquivos CSV para auditoria em escala.
* **Inteligência de Ameaças:** Comparação automatizada contra uma base de assinaturas de malwares.
* **Relatórios Profissionais:** Geração de documentos PDF contendo gráficos de pizza (Pie Chart) vetoriais nativos.
* **Zero Dependências de Imagem:** O gráfico é desenhado diretamente no PDF via código, resultando em arquivos menores e mais nítidos.
* **Arquitetura Moderna:** Construído sobre o UniPDF, garantindo suporte e atualizações de segurança.

---

## 🛠️ Pré-requisitos

* **Go 1.25+**
* **Chave UniPDF:** Este projeto utiliza a biblioteca `UniPDF`. Para remover a marca d'água de estudante, obtenha uma chave gratuita (Community) em [unidoc.io](https://cloud.unidoc.io/).

---

## 📦 Instalação

### 1. Clone o repositório
```bash
git clone https://github.com/seu-usuario/gohash_auditor.git
cd gohash_auditor
```

### 2. Inicialize o módulo
```bash
go mod init gohash-auditor 
```

### 3. Instale as dependências específicas do UniPDF
```bash
go get github.com/unidoc/unipdf/v3/common/license
go get github.com/unidoc/unipdf/v3/creator
go get github.com/unidoc/unipdf/v3/model
```

### 4. Organize as dependências
```bash
go mod tidy
```

### 5. Compile o binário
```bash
go build -o gohash-auditor main.go
```

---

## 📋 Como Utilizar

1.  **Prepare seus dados:** Coloque seu arquivo `apks.csv` na raiz do projeto (formato: um hash MD5 por linha).
2.  **Execute a ferramenta:**
    ```bash
    ./gohash-auditor
    ```
3.  **Resultado:** O arquivo `Relatorio_Final.pdf` será gerado instantaneamente.

---

## ⚙️ Configuração

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
