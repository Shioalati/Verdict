package cmd

import (
	"encoding/hex"
	"fmt"
	"os"
	"verdict/internal/domain/root"
	"verdict/internal/infrastructure/file"
	"verdict/internal/infrastructure/hash"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.Flags().StringVarP(&root.Algo, "algo", "a", "", "algorithm")
}

var rootCmd = &cobra.Command{
	Use:  "vec [path]",
	Args: cobra.MaximumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) > 0 {
			root.Path = args[0]
		}

		l := file.NewLocal()

		h, err := hash.New(root.Algo)
		if err != nil {
			return err
		}

		r, err := root.Handle(l, h)
		if err != nil {
			return err
		}

		fmt.Println(hex.EncodeToString(r))

		return nil
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
