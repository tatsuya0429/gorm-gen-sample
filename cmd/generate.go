/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
	"github.com/tatsuya0429/gorm-gen-sample/internal/models"
	"github.com/tatsuya0429/gorm-gen-sample/internal/sqlhandler"
	"gorm.io/gen"
)

type Querier interface {
	// SELECT * FROM @@table WHERE id = @id
	FindByID(id int64) (gen.T, error)
	// SELECT * FROM @@table WHERE user_id = @userId
	FindByUserId(userId int64) ([]gen.T, error)
	// SELECT * FROM @@table
	FindAll() ([]gen.T, error)
}

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		sqlHandler := sqlhandler.New()
		g := gen.NewGenerator(gen.Config{
			OutPath: "./internal/queries",
			Mode:    gen.WithQueryInterface,
		})
		g.UseDB(sqlHandler.Conn)
		g.ApplyBasic(models.User{}, models.Todo{})
		g.ApplyInterface(func(Querier) {}, models.User{}, models.Todo{})
		g.Execute()
	},
}

func init() {
	ormCmd.AddCommand(generateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// generateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// generateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
