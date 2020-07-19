package cmd

import (
	"github.com/spf13/cobra"
	"github.com/yuchanns/go-tour/internal/sql2struct"
	"log"
)

func init() {
	sqlCmd.AddCommand(sql2structCmd)
	sql2structCmd.Flags().StringVarP(&username, "username", "", "", "请输入数据库的账号")
	sql2structCmd.Flags().StringVarP(&password, "password", "", "", "请输入数据库的密码")
	sql2structCmd.Flags().StringVarP(&host, "host", "", "127.0.0.1:3306", "请输入数据库的HOST")
	sql2structCmd.Flags().StringVarP(&charset, "charset", "", "utf8mb4", "请输入数据库的编码")
	sql2structCmd.Flags().StringVarP(&dbType, "dbType", "", "mysql", "请输入数据库的类型")
	sql2structCmd.Flags().StringVarP(&dbName, "dbName", "", "", "请输入数据库的名称")
	sql2structCmd.Flags().StringVarP(&tableName, "tableName", "", "", "请输入数据库的表名")
}

var (
	username  string
	password  string
	host      string
	charset   string
	dbType    string
	dbName    string
	tableName string
)

var sqlCmd = &cobra.Command{
	Use:   "sql",
	Short: "sql转换和处理",
	Long:  "sql转换和处理",
	Run: func(cmd *cobra.Command, args []string) {
		//
	},
}

var sql2structCmd = &cobra.Command{
	Use:   "struct",
	Short: "sql转换",
	Long:  "sql转换",
	Run: func(cmd *cobra.Command, args []string) {
		dbInfo := &sql2struct.DBInfo{
			DBType:   dbType,
			Host:     host,
			UserName: username,
			Password: password,
			Charset:  charset,
		}
		dbModel := sql2struct.NewDBModel(dbInfo)
		if err := dbModel.Connect(); err != nil {
			log.Fatalf("dbModel.Connect err: %v", tableName)
		}
		columns, err := dbModel.GetColumns(dbName, tableName)
		if err != nil {
			log.Fatalf("dbModel.GetColumns err: %v", err)
		}
		tpl := sql2struct.NewStructTemplate()
		tplColumns := tpl.AssemblyColumns(columns)
		if err := tpl.Generate(tableName, tplColumns); err != nil {
			log.Fatalf("template.Generate err: %v", err)
		}
	},
}
