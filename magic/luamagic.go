package magic

import (
	"github.com/Shopify/go-lua"
	"log"
	"fmt"
)

var tempfiles Tempfiles

func DoFile(l *lua.State, fileName string) error {
	if err := lua.LoadFile(l, fileName, ""); err != nil {
		log.Fatal(err)
	}

	return l.ProtectedCall(0, lua.MultipleReturns, 1)
}

func Parselua(file *string) {
	l := lua.NewState()
	lua.OpenLibraries(l)

	l.PushGoFunction(client_script)
	l.SetGlobal("client_script")
	l.PushGoFunction(client_script)
	l.SetGlobal("client_scripts")
	l.PushGoFunction(client_script)
	l.SetGlobal("server_script")
	l.PushGoFunction(client_script)
	l.SetGlobal("server_scripts")
	l.PushGoFunction(client_script)
	l.SetGlobal("files")
	l.PushGoFunction(error_script)
	//l.SetGlobal("error_script")
	if err := DoFile(l, *file); err != nil {
		log.Fatal(err)
	}
	for _, v := range tempfiles.files {
		fmt.Println(v.src, v.isFile)
	}
}
func client_script(state *lua.State) int {
	numofArgs := state.Top()
	for i := 1; i <= numofArgs; i++ {
		if state.IsString(1) {
			lsting, ok := state.ToString(1)
			if ok != true {
				panic(ok)
			}
			tempfiles.files = append(tempfiles.files, File{src: lsting})
			state.Pop(-1)
		} else {
			nonnil := state.IsNoneOrNil(1)
			funcq := state.IsFunction(1)
			fmt.Println(nonnil, funcq)
			state.Pop(-1)
		}

	}
	return 0
}
func error_script(state *lua.State) int {
	numOfArgs := state.Top()
	for i := 0; i < numOfArgs; i++ {

	}
	/*nonnil:=state.IsNoneOrNil(1)
	funcq:=state.IsFunction(1)
	fmt.Println(nonnil,funcq)*/
	fmt.Println(state)
	//state.SetTop(0)
	return 0
}
