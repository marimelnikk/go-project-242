package main

import (
	"context"
	"fmt"
	"log"
	"os"

	code "code"

	"github.com/urfave/cli/v3"
)

func main() {
	cmd := &cli.Command{
		Name:      "hexlet-path-size",
		Usage:     "print size of a file or directory; supports -r (recursive), -H (human-readable), -a (include hidden)",
		ArgsUsage: "<path>",
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:    "human",
				Aliases: []string{"H"},
				Value:   false,
				Usage:   "human-readable sizes (auto-select unit) (default: false)",
			},
			&cli.BoolFlag{
				Name:    "all",
				Aliases: []string{"a"},
				Value:   false,
				Usage:   "include hidden files and directories (default: false)",
			},
			&cli.BoolFlag{
				Name: "recursive",
				Aliases: []string{"r"},
				Value: false,
				Usage: "recursive size of directories (default: false)",
			},
		},
		Action: func(ctx context.Context, cmd *cli.Command) error {
			path := cmd.Args().First()
			all := cmd.Bool("all")
			recursive := cmd.Bool("recursive")
			human := cmd.Bool("human")

			size, err := code.GetPathSize(path, recursive, human, all)
			if err != nil {
				return err
			}

			fmt.Printf("%s\t%s\n", size, path)
			return nil
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
