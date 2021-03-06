package sql2struct

import (
	"bytes"
	"fmt"
	"github.com/yuchanns/go-tour/internal/word"
	"go/format"
	"html/template"
	"os"
	"strings"
)

const StructTpl = `
type {{.TableName | ToCamelCase}} struct {
{{range .Columns}}{{$typeLen := len .Type}}{{if gt $typeLen 0}}{{.Name | ToCamelCase}} {{.Type}}{{.Tag}}{{else}}{{.Name}}{{end}} {{$length := len .Comment}}{{if gt $length 0}}// {{.Comment}}{{else}}// {{.Name}}{{end}}
{{end}}}

func (model {{.TableName | ToCamelCase}}) TableName() string {
    return "{{.TableName}}"
}
`

type StructTemplate struct {
	structTpl string
}

type StructColumn struct {
	Name    string
	Type    string
	Tag     string
	Comment string
}

type StructTemplateDB struct {
	TableName string
	Columns   []*StructColumn
}

func NewStructTemplate() *StructTemplate {
	return &StructTemplate{structTpl: StructTpl}
}

func (t *StructTemplate) AssemblyColumns(tbColumns []*TableColumn) []*StructColumn {
	tplColumns := make([]*StructColumn, 0, len(tbColumns))
	for _, column := range tbColumns {
		tplColumns = append(tplColumns, &StructColumn{
			Name:    column.ColumnName,
			Type:    DBTypeToStructType[column.DataType],
			Tag:     fmt.Sprintf("`json:%q`", column.ColumnName),
			Comment: column.ColumnComment,
		})
	}

	return tplColumns
}

func (t *StructTemplate) Generate(tableName string, tplColumns []*StructColumn) error {
	tpl := template.Must(template.New("sql2struct").Funcs(template.FuncMap{
		"ToCamelCase": word.UnderscoreToUpperCamelCase,
	}).Parse(t.structTpl))

	tplDB := StructTemplateDB{
		TableName: tableName,
		Columns:   tplColumns,
	}
	source := new(bytes.Buffer)
	if err := tpl.Execute(source, tplDB); err != nil {
		return err
	}

	content, err := format.Source(source.Bytes())
	if err != nil {
		return err
	}

	_, err = fmt.Fprint(os.Stdout, strings.Replace(string(content), "&#34;", "\"", -1))

	return err
}
