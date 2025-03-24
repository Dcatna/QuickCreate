/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

var projectName string

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create [template]",
	Short: "Create a new project using a predefined template",
	Long: `Scaffold a project from a built in template like react-app-js or react-app-ts`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("create called")
		template := args[0]

		if projectName == "" {
			fmt.Println("Please provide a project name with --name")
			return
		}

		fmt.Printf("Creating %s using template %s\n", projectName, template)

		switch template {
		case "react-app-js":
			runCommand("npm", "create", "vite@latest", projectName, "--", "--template", "react")
		case "react-app-ts":
			runCommand("npm", "create", "vite@latest", projectName, "--", "--template", "react-ts")

		default:
			fmt.Println("Unknown template, do --help for prefdefined templates")
		}
	},
}

func runCommand(name string, args ...string) {
	cmd := exec.Command(name, args...)
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout

	err := cmd.Run()
	if err != nil {
		fmt.Printf("Command failed %v\n", err)
	}
}

func init() {
	rootCmd.AddCommand(createCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	createCmd.Flags().StringVarP(&projectName, "name", "n", "", "Name of the project folder")
}
