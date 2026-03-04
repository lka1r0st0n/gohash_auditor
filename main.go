package main

import (
	"encoding/csv"
	"fmt"
	"io"
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
	license.SetMeteredKey(os.Getenv("UNIPDF_LICENSE_KEY"))
}

func main() {
	// --- 1. PROCESSAMENTO E LÓGICA ---
	inputPath := filepath.Join("data", "apks.csv")
	outputDir := filepath.Join("data", "reports")
	os.MkdirAll(outputDir, 0755)

	blacklist := map[string]string{
		"dc5e71c0d9c13f0482858c9f48267154": "Trojan.Android.Agent",
		"57338b2881eaac606733b95b1f920698": "Adware.AirPush.B",
		"03dccc10cca691849ae982fadb13a1e0": "Spyware.Stealer.X",
	}

	file, _ := os.Open(inputPath)
	defer file.Close()
	reader := csv.NewReader(file)
	total, infectados := 0, 0
	malwares := [][]string{}

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		hash := strings.ToLower(strings.TrimSpace(record[0]))
		total++
		if nome, existe := blacklist[hash]; existe {
			infectados++
			malwares = append(malwares, []string{hash, nome})
		}
	}

	// --- 2. GERAÇÃO DO PDF (ESTILO DASHBOARD) ---
	c := creator.New()
	c.SetPageSize(creator.PageSizeA4)
	c.NewPage()

	fontBold, _ := model.NewStandard14Font(model.HelveticaBoldName)
	fontRegular, _ := model.NewStandard14Font(model.HelveticaName)

	// Cabeçalho Escuro
	headerRect := c.NewRectangle(0, 0, 595, 80)
	headerRect.SetFillColor(creator.ColorRGBFrom8bit(18, 26, 44))
	c.Draw(headerRect)

	title := c.NewParagraph("GOHASH-AUDITOR")
	title.SetFont(fontBold)
	title.SetFontSize(22)
	title.SetColor(creator.ColorWhite)
	title.SetMargins(180, 0, 35, 0)
	c.Draw(title)

	// Dashboard (Cards)
	statsTable := c.NewTable(4)
	statsTable.SetMargins(30, 30, 40, 0)

	headers := []string{"Ameaças Identificadas", "Arquivos Seguros", "Total Analisado", "Data da Análise"}
	for _, h := range headers {
		cell := statsTable.NewCell()
		cell.SetBackgroundColor(creator.ColorRGBFrom8bit(240, 240, 240))
		cell.SetBorder(creator.CellBorderSideAll, creator.CellBorderStyleSingle, 0.5)
		p := c.NewParagraph(h)
		p.SetFont(fontBold)
		p.SetFontSize(8)
		p.SetMargins(5, 5, 5, 5)
		cell.SetContent(p)
	}

	addStat := func(val string, col creator.Color) {
		cell := statsTable.NewCell()
		cell.SetBorder(creator.CellBorderSideAll, creator.CellBorderStyleSingle, 0.5)
		p := c.NewParagraph(val)
		p.SetFont(fontBold)
		p.SetFontSize(16)
		p.SetColor(col)
		p.SetMargins(5, 5, 10, 10)
		cell.SetContent(p)
	}

	addStat(fmt.Sprintf("%d", infectados), creator.ColorRed)
	addStat(fmt.Sprintf("%d", total-infectados), creator.ColorRGBFrom8bit(39, 174, 96))
	addStat(fmt.Sprintf("%d", total), creator.ColorRGBFrom8bit(127, 140, 141))
	addStat(time.Now().Format("02/01/2006"), creator.ColorBlack)
	c.Draw(statsTable)

	// Blacklist Ativa
	lbl := c.NewParagraph("Blacklist Ativa")
	lbl.SetFont(fontBold)
	lbl.SetFontSize(12)
	lbl.SetMargins(30, 0, 30, 5)
	c.Draw(lbl)

	for _, m := range malwares {
		p := c.NewParagraph(fmt.Sprintf("Hash: %s  |  Tipo: %s  |  Status: Ativa", m[0], m[1]))
		p.SetFont(fontRegular)
		p.SetFontSize(8)
		p.SetMargins(35, 0, 5, 0)
		c.Draw(p)
	}

	// --- 3. EXPORTAÇÃO COM DATA E HORA NO NOME ---
	agora := time.Now()
	// Formato: relatorio_dashboard_DD-MM-YYYY_HH-MM-SS.pdf
	nomeArquivo := fmt.Sprintf("relatorio_dashboard_%s.pdf", agora.Format("02-01-2006_15-04-05"))
	c.WriteToFile(filepath.Join(outputDir, nomeArquivo))

	fmt.Printf("✅ Relatório gerado com sucesso: %s\n", nomeArquivo)
}
