package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/yuchanns/go-tour/internal/word"
	"log"
	"strings"
)

const (
	ModeUpper = iota + 1
	ModeLower
	ModeUnderscoreToUpperCamelcase
	ModeUnderscoreToLowerCamelcase
	ModeCamelcaseToUndersore
)

var str string
var mode int8

var desc = strings.Join([]string{
	"该子命令支持各种单次格式转换，模式如下：",
	"1：全部单词转为大写",
	"2：全部单词转为小写",
	"3：下划线单词转为大写驼峰单词",
	"4：下划线单词转为小写驼峰单词",
	"5：驼峰单词转为下划线单词",
}, "\n")

func init() {
	wordCmd.Flags().StringVarP(&str, "str", "s", "", "请输入单词内容")
	wordCmd.Flags().Int8VarP(&mode, "mode", "m", 0, "请输入单词转换模式")
}

var wordCmd = &cobra.Command{
	Use:   "word",
	Short: "单词格式转换",
	Long:  desc,
	Run: func(cmd *cobra.Command, args []string) {
		var content string
		switch mode {
		case ModeUpper:
			content = word.ToUpper(str)
		case ModeLower:
			content = word.ToLower(str)
		case ModeUnderscoreToUpperCamelcase:
			content = word.UnderscoreToUpperCamelCase(str)
		case ModeUnderscoreToLowerCamelcase:
			content = word.UnderscoreToLowerCamelCase(str)
		case ModeCamelcaseToUndersore:
			content = word.CamelCaseToUnderscore(str)
		default:
			log.Fatalln("暂不支持该模式转换，执行help word查看帮助文档")
		}
		fmt.Printf("输出结果： %s\n", content)
	},
}
