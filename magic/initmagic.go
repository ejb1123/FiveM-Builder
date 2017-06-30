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
    "os"
    "io/ioutil"
    "path"
    "path/filepath"
    //	"net/http"
    //"github.com/ulikunitz/xz"
    "io"
    "fmt"
    //"go/ast"
    "bytes"
    //	"archive/zip"
    //	"encoding/binary"
    "strings"
    "log"
    //	"xi2.org/x/xz"
    //	"github.com/docker/docker/pkg/discovery/file"
    "github.com/nu7hatch/gouuid"
    "text/template"
)

//var filesToDownload= []string{"s","s"};
func CreateDirectory(folder *string) {
    if _, err := os.Stat(*folder); os.IsExist(err) {
        fmt.Println("folder exist")
        os.Exit(0)
    }
    os.MkdirAll(*folder, os.ModeDir)
    data, err := Asset("data/config.yml")
    if err != nil {
        // Asset was not found.
    }
    ioutil.WriteFile(*folder+"/config.yml", data, 0)
    data2, _ := Asset("data/icecon.exe")
    ioutil.WriteFile(*folder+"/icecon.exe", data2, 0)

    extractCSproj(folder)
    setUpCsProj(folder)

}

type Tempvars struct {
    Safeprojectname string
    Guid1           string
}

func extractCSproj(folder *string) {

    data := AssetNames()

    for _, file := range data {



        fmt.Println(file, "extarcting ")


        newDstFileName := strings.Replace(file, "files", "src", -1)
        err := os.MkdirAll(filepath.Dir(filepath.Join(*folder, newDstFileName)), os.ModeDir)

        if err != nil {
            panic(err)
        }
        newFile, err := os.OpenFile(path.Join(*folder, newDstFileName), os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0)
        if err != nil {
            log.Fatal(err)
        }
        //dstFile, err := os.OpenFile(filepath.Join(*folder, "src", file.Name), os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0)
        if err != nil {
            panic(err)
        }

        io.Copy(newFile, bytes.NewReader(MustAsset(file)))
    }

}

func setUpCsProj(folder *string) {


    guidsalt, _ := filepath.Abs(*folder)

    u5, err := uuid.NewV5(uuid.NamespaceURL, []byte(guidsalt))
    if err != nil {
        fmt.Println("error:", err)
        return
    }
    Safeprojectname := Tempvars{
        Safeprojectname: *folder, Guid1: u5.String(),
    }
    csprojTemplate := template.New("")
}

type Directory struct {
    path string;
}

func (d Directory) f() {
    d.path = "k"
}
func
reduceFiles(info *Cache_info) []ContentFile {
    filesretuned := []ContentFile{}
    for _, v := range info.Content {
        if filepath.ToSlash(filepath.Dir(v.Name)) == filepath.ToSlash(`citizen/clr2/lib/mono/4.5`) {
            filesretuned = append(filesretuned, v)
        }
    }
    return filesretuned
}
