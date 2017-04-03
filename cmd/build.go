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
	//"fmt"

	"github.com/spf13/cobra"
	"ejb1123.tk/FiveMBuilder/magic"
	"fmt"
	"path/filepath"
	"os/exec"
	"os"
	"log"
)

// buildCmd represents the build command
var buildCmd = &cobra.Command{
	Use:   "build",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Work your own magic here
		fmt.Println("finding build tools")
		versions := magic.FindBuildTools()
		h := magic.IsVSVersionSupported(versions)
		proj, _ := (cmd.Flags().GetString("proj"))

		for key, val := range h {
			fmt.Println(key, val)
		}
		fmt.Println(filepath.Base(proj))
		cmdd := exec.Command(h["15.0"],filepath.Base(proj))
		cmdd.Dir=filepath.Dir(proj)
		_,err:= cmdd.StdoutPipe()
		if(err!=nil){
			log.Fatal(err)
		}
		cmdd.Stdout = os.Stdout
		cmdd.Stderr = os.Stderr
		if err:=cmdd.Run(); err!=nil{
			log.Fatal(err)
		}
		fmt.Println(cmdd.Stdout)
		//magic.IsVSVersionSupported(versions)
	},
}

func init() {
	//RootCmd.AddCommand(buildCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// buildCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	buildCmd.Flags().StringP("proj", "p", "", "project file (.csproj)")
}
