package cmd

import (
	"fmt"

	"github.com/zcoriarty/Backend/config"
	"github.com/zcoriarty/Backend/repository"

	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

// createAlgoSchemaCmd represents the createschema command
var createAlgoSchemaCmd = &cobra.Command{
	Use:   "create_algo_schema",
	Short: "create_algo_schema adds the schema for the trading algorithms to the database",
	Long:  `create_algo_schema adds the schema for the trading algorithms to the database`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("create algo schema called")

		db := config.GetConnection()
		log, _ := zap.NewDevelopment()
		p := config.GetPostgresConfig()
		defer log.Sync()
		repository.CreateAlgoSchema(db, p)
		
	},
}

func init() {
	rootCmd.AddCommand(createAlgoSchemaCmd)
}