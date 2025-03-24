/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// helpCmd represents the help command
var helpCmd = &cobra.Command{
	Use:   "help [command]",
	Short: "Help will give more specific details about a command",
	Args: cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		if (len(args) == 0) {
			printGeneralHelp()
			return 
		}

		switch args[0] {
			case "create":
				printCreateHelp()

			default:
				fmt.Printf("Unknwon function %s\n", args[0])
		}
			
	},
}

func printGeneralHelp() {
	fmt.Println(`
		QuickCreate CLI - Project Generator

		Commands:
			help	Show this help screen or details about a specific command
			create  Scaffold a new project from a template
		
		Examples:
			qc help create
			qc create react-js
	`)
}

func printCreateHelp() {
	fmt.Println(`
	QuickCreate create command
	
	Create a project from a built-in template
	
	Usage:
		qc create <template_name> --name or -n <project_name> [--path ./mydir]
		
	Templates:
		react-js       ->   React + Vite (JS)
		react-ts       ->   React + Vite (TS)
		go-api		   ->   Basic Golang Backend
		next-app	   ->   Next.js Fullstack App

		Examples:
			qc create react-ts  --name dashboard
			qc create go-api --name api --path ./apps
		`)
}

func init() {
	rootCmd.AddCommand(helpCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// helpCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// helpCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
