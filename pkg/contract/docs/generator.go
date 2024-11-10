package docs

import (
	"bytes"
	"encoding/json"
	"html/template"
	"path/filepath"

	"github.com/wanglilind/qqq/pkg/contract"
)

// ææ¡£çæå?
type DocumentGenerator struct {
	templates    map[string]*template.Template
	outputPath   string
	contractInfo map[string]ContractInfo
}

type ContractInfo struct {
	Name        string
	Version     string
	Author      string
	Description string
	ABI         []contract.ABIMethod
	Source      string
	Examples    []Example
}

type Example struct {
	Name        string
	Description string
	Code        string
	Output      string
}

func NewDocumentGenerator(outputPath string) (*DocumentGenerator, error) {
	dg := &DocumentGenerator{
		templates:    make(map[string]*template.Template),
		outputPath:   outputPath,
		contractInfo: make(map[string]ContractInfo),
	}

	// å è½½ææ¡£æ¨¡æ¿
	if err := dg.loadTemplates(); err != nil {
		return nil, err
	}

	return dg, nil
}

// çæåçº¦ææ¡£
func (dg *DocumentGenerator) GenerateContractDocs(contractAddr string, info ContractInfo) error {
	// çæMarkdownææ¡£
	if err := dg.generateMarkdown(contractAddr, info); err != nil {
		return err
	}

	// çæHTMLææ¡£
	if err := dg.generateHTML(contractAddr, info); err != nil {
		return err
	}

	// çæAPIææ¡£
	if err := dg.generateAPIDoc(contractAddr, info); err != nil {
		return err
	}

	return nil
}

// çæMarkdownææ¡£
func (dg *DocumentGenerator) generateMarkdown(contractAddr string, info ContractInfo) error {
	tmpl := dg.templates["markdown"]
	var buf bytes.Buffer

	if err := tmpl.Execute(&buf, info); err != nil {
		return err
	}

	outputFile := filepath.Join(dg.outputPath, contractAddr, "README.md")
	return writeFile(outputFile, buf.Bytes())
}

// çæHTMLææ¡£
func (dg *DocumentGenerator) generateHTML(contractAddr string, info ContractInfo) error {
	tmpl := dg.templates["html"]
	var buf bytes.Buffer

	if err := tmpl.Execute(&buf, info); err != nil {
		return err
	}

	outputFile := filepath.Join(dg.outputPath, contractAddr, "index.html")
	return writeFile(outputFile, buf.Bytes())
}

// çæAPIææ¡£
func (dg *DocumentGenerator) generateAPIDoc(contractAddr string, info ContractInfo) error {
	// çæOpenAPI/Swaggerææ¡£
	apiDoc := generateOpenAPISpec(info)
	
	data, err := json.MarshalIndent(apiDoc, "", "  ")
	if err != nil {
		return err
	}

	outputFile := filepath.Join(dg.outputPath, contractAddr, "api.json")
	return writeFile(outputFile, data)
}

// å è½½ææ¡£æ¨¡æ¿
func (dg *DocumentGenerator) loadTemplates() error {
	// å è½½Markdownæ¨¡æ¿
	mdTmpl, err := template.ParseFiles("templates/contract.md")
	if err != nil {
		return err
	}
	dg.templates["markdown"] = mdTmpl

	// å è½½HTMLæ¨¡æ¿
	htmlTmpl, err := template.ParseFiles("templates/contract.html")
	if err != nil {
		return err
	}
	dg.templates["html"] = htmlTmpl

	return nil
}

// çæOpenAPIè§è
func generateOpenAPISpec(info ContractInfo) map[string]interface{} {
	// å®ç°OpenAPIè§èçæé»è¾
	return nil
}

// åå¥æä»¶
func writeFile(path string, data []byte) error {
	// å®ç°æä»¶åå¥é»è¾
	return nil
} 
