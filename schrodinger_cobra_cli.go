package main

import (
	"database/sql"
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

type SchrodingerCLI struct {
	DB                     *sql.DB
	RandomNumbersGenerator RandomNumbersGenerator
}

func InitCobraCLI(db *sql.DB) {
	cli := &SchrodingerCLI{
		DB:                     db,
		RandomNumbersGenerator: &RandomNumbersGeneratorImpl{},
	}

	rootCmd := &cobra.Command{
		Use:   "schrodinger",
		Short: "A key-value store that randomly breaks",
		Long:  `Schrödinger's Database - A key-value store in Go that randomly "breaks" — sometimes returning the wrong data or failing mysteriously.`,
	}

	putCmd := &cobra.Command{
		Use:   "put [key] [value]",
		Short: "Store a key-value pair",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			key := args[0]
			value := args[1]
			return storeSchrodingerData(key, value, cli.DB, cli.RandomNumbersGenerator)
		},
	}

	getCmd := &cobra.Command{
		Use:   "get [key]",
		Short: "Retrieve a value by key",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			key := args[0]
			err, value := retrieveSchrodingerData(cli.DB, cli.RandomNumbersGenerator, key)
			if err != nil {
				return err
			}
			fmt.Println(value)
			return nil
		},
	}

	delCmd := &cobra.Command{
		Use:   "del [key]",
		Short: "Delete a key-value pair",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			key := args[0]
			return removeSchrodingerData(cli.DB, key, cli.RandomNumbersGenerator)
		},
	}

	dumpCmd := &cobra.Command{
		Use:   "dump",
		Short: "List all stored key-value pairs",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			return dump(cli.DB)
		},
	}

	rootCmd.AddCommand(putCmd)
	rootCmd.AddCommand(getCmd)
	rootCmd.AddCommand(delCmd)
	rootCmd.AddCommand(dumpCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
