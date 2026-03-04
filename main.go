package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/unidoc/unipdf/v3/common/license"
	"github.com/unidoc/unipdf/v3/creator"
	"github.com/unidoc/unipdf/v3/model"
)

func init() {
	// Agora o código procura uma chave no sistema, em vez de ter o texto fixo
	key := os.Getenv("UNIPDF_LICENSE_KEY")
	if key == "" {
		fmt.Println("⚠️ Aviso: Variável UNIPDF_LICENSE_KEY não configurada.")
	}
	license.SetMeteredKey(key)
}

func main() {
	// 1. CONFIGURAÇÃO DE CAMINHOS
	inputPath := filepath.Join("data", "apks.csv")
	outputDir := "data"

	// Garante que a pasta data existe
	if _, err := os.Stat(outputDir); os.IsNotExist(err) {
		os.Mkdir(outputDir, 0755)
	}

	blacklist := map[string]string{
		"dc5e71c0d9c13f0482858c9f48267154": "Trojan.Android.Agent",
		"57338b2881eaac606733b95b1f920698": "Adware.AirPush.B",
		"03dccc10cca691849ae982fadb13a1e0": "Spyware.Stealer.X",
	}

	// 2. LEITURA DOS DADOS (Agora puxando de data/apks.csv)
	file, err := os.Open(inputPath)
	if err != nil {
		log.Fatalf("❌ Erro: O arquivo '%s' não foi encontrado na pasta data!", inputPath)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	total, infectados := 0, 0
	listaAmeacas := [][]string{}

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			continue
		}
		hash := strings.ToLower(strings.TrimSpace(record[0]))
		total++

		if nomeMalware, existe := blacklist[hash]; existe {
			infectados++
			listaAmeacas = append(listaAmeacas, []string{hash, nomeMalware})
		}
	}
	limpos := total - infectados

	// 3. CONSTRUÇÃO DO PDF
	c := creator.New()
	c.SetPageSize(creator.PageSizeA4)
	c.NewPage()

	fontBold, _ := model.NewStandard14Font(model.HelveticaBoldName)

	title := c.NewParagraph("GoHash-Auditor: Relatório de Segurança")
	title.SetFont(fontBold)
	title.SetFontSize(20)
	title.SetColor(creator.ColorRGBFrom8bit(44, 62, 80))
	title.SetMargins(0, 0, 20, 10)
	c.Draw(title)

	agora := time.Now()
	ts := c.NewParagraph("Gerado em: " + agora.Format("02/01/2006 15:04:05"))
	ts.SetFontSize(10)
	c.Draw(ts)

	// Tabela de Resumo
	table := c.NewTable(2)
	table.SetMargins(0, 0, 20, 20)

	drawCell := func(text string, isHeader bool) {
		cell := table.NewCell()
		cell.SetBorder(creator.CellBorderSideAll, creator.CellBorderStyleSingle, 1)
		if isHeader {
			cell.SetBackgroundColor(creator.ColorRGBFrom8bit(236, 240, 241))
		}
		p := c.NewParagraph(text)
		if isHeader {
			p.SetFont(fontBold)
		}
		cell.SetContent(p)
	}

	drawCell("Total Analisado", true)
	drawCell(fmt.Sprintf("%d", total), false)
	drawCell("Ameaças", true)
	drawCell(fmt.Sprintf("%d", infectados), false)
	drawCell("Limpos", true)
	drawCell(fmt.Sprintf("%d", limpos), false)
	c.Draw(table)

	// 4. SALVAMENTO (Nome com Data e Hora na pasta data)
	fileName := fmt.Sprintf("Relatorio_%s.pdf", agora.Format("02-01-2006_15-04"))
	finalPath := filepath.Join(outputDir, fileName)

	err = c.WriteToFile(finalPath)
	if err != nil {
		log.Fatalf("❌ Erro ao salvar PDF: %v", err)
	}

	fmt.Printf("✨ Auditoria concluída!\n📥 Entrada: %s\n📄 Relatório: %s\n", inputPath, finalPath)
}
