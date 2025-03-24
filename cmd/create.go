/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/spf13/cobra"
)

var projectName string
var projectPath string

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create [template]",
	Short: "Create a new project using a predefined template",
	Long:  `Scaffold a project from a built in template like react-js or react-ts`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("create called")
		template := args[0]

		if projectName == "" {
			fmt.Println("Please provide a project name with --name")
			return
		}

		if projectPath == "" {
			cwd, _ := os.Getwd()
			projectPath = cwd
		}

		fullPath := filepath.Join(projectPath, projectName)

		fmt.Printf("Creating %s using template %s\n", projectName, template)

		switch template {
		case "react-js":
			fmt.Println("Creating React with JavaScript project")
			runCommand("npm", "create", "vite@latest", fullPath, "--", "--template", "react")

		case "react-ts":
			fmt.Println("Creating React with TypeScript project")
			runCommand("npm", "create", "vite@latest", fullPath, "--", "--template", "react-ts")

		case "react-js-tailwind":
			fmt.Println("Creating React + JavaScript + Tailwind project")
			runCommand("npm", "create", "vite@latest", fullPath, "--", "--template", "react")
			tailwindInstall := exec.Command("npm", "install", "-n", "postcss", "autoprefixer")
			tailwindInstall.Dir = projectName
			tailwindInstall.Stderr = os.Stderr
			tailwindInstall.Stdin = os.Stdin
			tailwindInstall.Stdout = os.Stdout
			installErr := tailwindInstall.Run()

			if installErr != nil {
				fmt.Printf("Command failed with %v\n", installErr)
			}

			tailwindInit := exec.Command("npx", "tailwindcss", "init", "-p")
			tailwindInit.Dir = projectName
			tailwindInit.Stderr = os.Stderr
			tailwindInit.Stdin = os.Stdin
			tailwindInit.Stdout = os.Stdout
			initErr := tailwindInit.Run()

			if initErr != nil {
				fmt.Printf("Command failed with %v\n", initErr)
			}
			updateTailwindFiles(fullPath)

		case "react-ts-tailwind":
			fmt.Println("Creating React + TypeScript + Tailwind project")
			runCommand("npm", "create", "vite@latest", fullPath, "--", "--template", "react-ts")
			tailwindInstall := exec.Command("npm", "install", "-n", "postcss", "autoprefixer")
			tailwindInstall.Dir = projectName
			tailwindInstall.Stderr = os.Stderr
			tailwindInstall.Stdin = os.Stdin
			tailwindInstall.Stdout = os.Stdout
			installErr := tailwindInstall.Run()

			if installErr != nil {
				fmt.Printf("Command failed with %v\n", installErr)
			}

			tailwindInit := exec.Command("npx", "tailwindcss", "init", "-p")
			tailwindInit.Dir = projectName
			tailwindInit.Stderr = os.Stderr
			tailwindInit.Stdin = os.Stdin
			tailwindInit.Stdout = os.Stdout
			initErr := tailwindInit.Run()

			if initErr != nil {
				fmt.Printf("Command failed with %v\n", initErr)
			}
			updateTailwindFiles(fullPath)

		case "go-api":
			fmt.Println("Creating go project... not done yet")

		case "next-js":
			fmt.Println("Creating Next.js project")

			if err := os.Mkdir(projectPath, os.ModePerm) ; err != nil {
				log.Fatalf("Error creating folder %e\n", err)

			}

			cmd := exec.Command("npx", "create-next-app@latest", projectName, "--ts")
			cmd.Dir = projectName
			cmd.Stderr = os.Stderr
			cmd.Stdin = os.Stdin
			cmd.Stdout = os.Stdout

			if err := cmd.Run() ; err != nil {
				log.Fatalf("Error creating next-js app %e\n", err)
			}


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

func updateTailwindFiles(projectName string) {
	tailwindConfig := filepath.Join(projectName, "tailwind.config.js")
	indexCSS := filepath.Join(projectName, "src", "index.css")

	tailwindConfigContent := `
	/** @type {import('tailwindcss').Config} */
	module.exports = {
	  content: ["./index.html", "./src/**/*.{js,ts,jsx,tsx}"],
	  theme: {
		extend: {},
	  },
	  plugins: [],
	}`

	os.WriteFile(tailwindConfig, []byte(tailwindConfigContent), 0644)

	indexCSSContent := `
	@tailwind base;
	@tailwind components;
	@tailwind utilities;`

	os.WriteFile(indexCSS, []byte(indexCSSContent), 0644)
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
	createCmd.Flags().StringVarP(&projectPath, "path", "p", "", "Optional path to create the project in")
}
