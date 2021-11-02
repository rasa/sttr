// Code generated by github.com/abhimanyu003/sttr/cmd/generate.go. DO NOT EDIT

package cmd

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/abhimanyu003/sttr/processors"
	"github.com/spf13/cobra"
)

var base64Encode_flag_r bool

func init() {
	base64EncodeCmd.Flags().BoolVarP(&base64Encode_flag_r, "raw", "r", false, "unpadded base64 encoding")
	rootCmd.AddCommand(base64EncodeCmd)
}

var base64EncodeCmd = &cobra.Command{
	Use:     "base64-encode",
	Short:   "Encode your text to Base64",
	Aliases: []string{"b64-enc", "b64-encode"},
	Args:    cobra.MaximumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		var err error
		var in []byte
		var out string

		if len(args) == 0 {
			in, err = ioutil.ReadAll(cmd.InOrStdin())
			if err != nil {
				return err
			}
		} else {
			if fi, err := os.Stat(args[0]); err == nil && !fi.IsDir() {
				d, err := ioutil.ReadFile(args[0])
				if err != nil {
					return err
				}
				in = d
			} else {
				in = []byte(args[0])
			}
		}

		flags := make([]processors.Flag, 0)
		p := processors.Base64Encode{}
		flags = append(flags, processors.Flag{Short: "r", Value: base64Encode_flag_r})

		out, err = p.Transform(in, flags...)
		if err != nil {
			return err
		}

		_, err = fmt.Fprint(os.Stdout, out)
		return err
	},
}
