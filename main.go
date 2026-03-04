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

	"github.com/joho/godotenv"
	"github.com/unidoc/unipdf/v3/common/license"
	"github.com/unidoc/unipdf/v3/creator"
	"github.com/unidoc/unipdf/v3/model"
)

func init() {
	godotenv.Load()
	key := os.Getenv("UNIPDF_LICENSE_KEY")
	license.SetMeteredKey(key)
}

func main() {
	// --- 1. DEFINIÇÃO DE CAMINHOS ---
	inputPath := filepath.Join("data", "apks.csv")
	outputDir := filepath.Join("data", "reports")

	if _, err := os.Stat(outputDir); os.IsNotExist(err) {
		os.MkdirAll(outputDir, 0755)
	}

	// --- 2. BASE DE CONHECIMENTO (BLACKLIST) ---
	blacklist := map[string]string{
		"dc5e71c0d9c13f0482858c9f48267154": "Trojan.Android.Agent",
		"57338b2881eaac606733b95b1f920698": "Adware.AirPush.B",
		"03dccc10cca691849ae982fadb13a1e0": "Spyware.Stealer.X",
	}

	// --- 3. PROCESSAMENTO DOS DADOS (CSV) ---
	file, err := os.Open(inputPath)
	if err != nil {
		log.Fatalf("❌ Erro: Ficheiro '%s' não encontrado.", inputPath)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	total, infectados := 0, 0
	malwaresEncontrados := [][]string{}

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil || len(record) == 0 {
			continue
		}

		hash := strings.ToLower(strings.TrimSpace(record[0]))
		total++
		if nome, existe := blacklist[hash]; existe {
			infectados++
			malwaresEncontrados = append(malwaresEncontrados, []string{hash, nome})
		}
	}
	limpos := total - infectados

	// --- 4. CONFIGURAÇÃO VISUAL DO PDF ---
	c := creator.New()
	c.SetPageSize(creator.PageSizeA4)
	c.NewPage()

	fontBold, _ := model.NewStandard14Font(model.HelveticaBoldName)
	fontRegular, _ := model.NewStandard14Font(model.HelveticaName)

	corPrimaria := creator.ColorRGBFrom8bit(44, 62, 80)
	corSucesso := creator.ColorRGBFrom8bit(39, 174, 96)
	corPerigo := creator.ColorRGBFrom8bit(192, 57, 43)

	headerBg := c.NewRectangle(0, 0, 595, 100)
	headerBg.SetFillColor(corPrimaria)
	c.Draw(headerBg)

	// Desenho do Título Atualizado
	pTitle := c.NewParagraph("GOHASH-AUDITOR")
	pTitle.SetFont(fontBold)
	pTitle.SetFontSize(24)
	pTitle.SetColor(creator.ColorWhite)
	pTitle.SetMargins(40, 0, 45, 0) // Margem superior alterada para 45
	c.Draw(pTitle)

	subtitle := c.NewParagraph("RELATÓRIO TÉCNICO DE INTEGRIDADE")
	subtitle.SetFont(fontRegular)
	subtitle.SetFontSize(12)
	subtitle.SetColor(creator.ColorRGBFrom8bit(189, 195, 199))
	subtitle.SetMargins(40, 0, 5, 0)
	c.Draw(subtitle)

	// --- 5. DASHBOARD DE INDICADORES (TABELA) ---
	table := c.NewTable(3)
	table.SetMargins(40, 40, 40, 0)

	addLabel := func(txt string) {
		cell := table.NewCell()
		cell.SetBackgroundColor(creator.ColorRGBFrom8bit(236, 240, 241))
		p := c.NewParagraph(txt)
		p.SetFont(fontBold)
		p.SetFontSize(10)
		p.SetMargins(10, 10, 10, 10)
		cell.SetContent(p)
	}

	addLabel("TOTAL ANALISADO")
	addLabel("AMEAÇAS IDENTIFICADAS")
	addLabel("ARQUIVOS SEGUROS")

	addValue := func(txt string, col creator.Color) {
		cell := table.NewCell()
		cell.SetBorder(creator.CellBorderSideBottom, creator.CellBorderStyleSingle, 2)
		p := c.NewParagraph(txt)
		p.SetFont(fontBold)
		p.SetFontSize(22)
		p.SetColor(col)
		p.SetMargins(10, 10, 10, 10)
		cell.SetContent(p)
	}

	addValue(fmt.Sprintf("%d", total), corPrimaria)
	addValue(fmt.Sprintf("%d", infectados), corPerigo)
	addValue(fmt.Sprintf("%d", limpos), corSucesso)

	c.Draw(table)

	if len(malwaresEncontrados) > 0 {
		sectionTitle := c.NewParagraph("DETALHES TÉCNICOS DA BLACKLIST")
		sectionTitle.SetFont(fontBold)
		sectionTitle.SetFontSize(14)
		sectionTitle.SetMargins(40, 0, 40, 10)
		c.Draw(sectionTitle)

		for _, m := range malwaresEncontrados {
			p := c.NewParagraph(fmt.Sprintf("• %s  →  DETECTADO: %s", m[0], m[1]))
			p.SetFont(fontRegular)
			p.SetFontSize(9)
			p.SetMargins(50, 0, 2, 0)
			c.Draw(p)
		}
	}

	// --- 6. FINALIZAÇÃO E EXPORTAÇÃO ---
	agora := time.Now()
	fileName := fmt.Sprintf("Relatorio_%s.pdf", agora.Format("02-01-2006_15-04-05"))
	finalPath := filepath.Join(outputDir, fileName)

	if err := c.WriteToFile(finalPath); err != nil {
		log.Fatalf("❌ Erro ao salvar PDF: %v", err)
	}

	fmt.Printf("✅ Auditoria finalizada!\n📁 Local: %s\n", finalPath)
}
