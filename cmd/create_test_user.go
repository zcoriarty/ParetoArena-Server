package cmd

import (
	"fmt"

	"github.com/zcoriarty/Backend/config"
	"github.com/zcoriarty/Backend/manager"
	"github.com/zcoriarty/Backend/repository"
	"github.com/zcoriarty/Backend/secret"

	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

var testEmail string
var testPassword string
var testAccountId string
var createTestUserCmd = &cobra.Command{
	Use:   "create_test_user",
	Short: "create_test_user creates a new user",
	Long:  `create_test_user creates a new user for testing the app`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("create_test_user called")

		testEmail, _ = cmd.Flags().GetString("testEmail")
		fmt.Println(testEmail)
		if !validateEmail(testEmail) {
			fmt.Println("Invalid email provided; test user not created")
			return
		}

		testPassword, _ = cmd.Flags().GetString("testPassword")
		if password == "" {
			password, _ = secret.GenerateRandomString(16)
		}

		testAccountId, _ = cmd.Flags().GetString("testAccountId")

		db := config.GetConnection()
		log, _ := zap.NewDevelopment()
		defer log.Sync()
		accountRepo := repository.NewAccountRepo(db, log, secret.New())
		roleRepo := repository.NewRoleRepo(db, log)

		m := manager.NewManager(accountRepo, roleRepo, db)
		m.CreateTestUser(testEmail, testPassword, testAccountId)
	},
}

func init() {
	localFlags := createTestUserCmd.Flags()
	localFlags.StringVarP(&testEmail, "testEmail", "e", "", "Test user's email")
	localFlags.StringVarP(&testPassword, "testPassword", "p", "", "Test user's password")
	localFlags.StringVarP(&testAccountId, "testAccountId", "a", "", "Test user's account id")
	createTestUserCmd.MarkFlagRequired("testEmail")
	rootCmd.AddCommand(createTestUserCmd)
}
