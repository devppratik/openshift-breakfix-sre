package main

import (
	"fmt"
	"openshift-breakfix/challenges"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Set the start value",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Welcome to Openshift Breakfix")
		prompt := promptui.Prompt{
			Label: "Cluster Id",
		}

		clusterId, err := prompt.Run()

		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return
		}
		getKubeConfigAndClient(clusterId)
		DisplayPrompt()
	},
}

var gradeCmd = &cobra.Command{
	Use:   "grade",
	Short: "Set the grade",
	Run: func(cmd *cobra.Command, args []string) {
		challenges.ValidateCluster()
	},
}

var endCmd = &cobra.Command{
	Use:   "end",
	Short: "Set the end value",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("[SUCCESS] Removed Break Fix Items")
		fmt.Println("Thank you for using the tool")
	},
}

var rootCmd = &cobra.Command{
	Use:   "",
	Short: "A simple CLI based SRE breakfix",
}

func main() {
	rootCmd.AddCommand(startCmd)
	rootCmd.AddCommand(gradeCmd)
	rootCmd.AddCommand(endCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
	}
}
