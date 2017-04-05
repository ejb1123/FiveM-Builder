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
package magic

import (
	//"golang.org/x/sys/windows/registry"

	"golang.org/x/sys/windows/registry"
	//"runtime"
	"path"
	"os"
	"path/filepath"
)

func FindBuildTools() map[string]string {
	//fmt.Println(runtime.GOOS, runtime.GOARCH)
	keys, err := registry.OpenKey(registry.LOCAL_MACHINE, `SOFTWARE\WOW6432Node\Microsoft\VisualStudio\SxS\VS7`, registry.READ|registry.QUERY_VALUE)
	if err != nil {
		panic(err)
	}
	k, err := keys.ReadValueNames(-1)
	if err != nil {
		panic(err)
	}
	fin2 := make(map[string]string)

	//println(k.SubKeyCount,k.ValueCount)
	for _, val := range k {
		i := val
		if err != nil {
			panic(err)
		}
		j, _, _ := keys.GetStringValue(val)
		fin2[i] = j

	}
	return fin2
}

func IsVSVersionSupported(m map[string]string) map[string]string {
	newMap := make(map[string]string)
	for key, value := range m {
		newPath := filepath.ToSlash(path.Join(value, "MSBuild\\"+key+`\Bin\MSBuild.exe`))
		if _, err := os.Stat(newPath); os.IsNotExist(err) {
			//fmt.Println("msbuild does nor exist for visual studio " + key)

			///fmt.Println(newPath)
		} else {
			newMap[key] = path.Clean(newPath)
		}
	}
	return newMap
}
