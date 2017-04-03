/*MIT License

Copyright (c) 2017 ejb1123

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.*/

package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	//	"net"
	"ejb1123.tk/FiveMBuilder/magic"
)

var configFile string

// copyCmd represents the copy command
var copyCmd = &cobra.Command{
	Use:   "copy",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Work your own magic here

		fmt.Println("copy called")
		h:=magic.ReadConfig(configFile)
		fmt.Println(*h)
		files := magic.GetFiles(&h.Server.Src)
		magic.DoCopy(files,&h.Server.Src,&h.Server.Root,&h.Server.ProjectName)
		magic.RestartServer(&h.Server.Url,&h.Server.Password,&h.Server.ProjectName,&h.Server.IceCon)
	},
}

func init() {
	RootCmd.AddCommand(copyCmd)

	copyCmd.Flags().StringVarP(&configFile, "configy", "c", "", "YAML sonfig file")
	/*buildCmd.Flags().StringP("source", "s", "", "the folder to copy files from")
	buildCmd.Flags().StringP("dest","d","","the folder to copy the files to")
	buildCmd.Flags().IntP("port","p",30120,"port for server")
	buildCmd.Flags().String("ip","127.0.0.1","server ip")
	buildCmd.Flags().StringP("password","pass","lovely","server rcon password")*/
}
