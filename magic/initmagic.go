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
	"net/http"
	//"github.com/ulikunitz/xz"
	"io"
	"fmt"
	//"go/ast"
	"bytes"
	"archive/zip"
	"encoding/binary"
	"strings"
	"log"
	"text/template"
	"github.com/nu7hatch/gouuid"
	"xi2.org/x/xz"
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
	os.MkdirAll(path.Join(*folder, "/lib"), os.ModeDir)
	allfiles := Parsexml()
	ffiles := reduceFiles(allfiles)
	for _, v := range ffiles {
		extra := ""
		if v.CompressedSize != v.Size {
			extra = ".xz"
		}
		res, err := http.Get(`http://runtime.fivem.net/client/prod/content/fivereborn/` + filepath.ToSlash(v.Name) + extra)
		if err != nil {
			panic(err)
		}
		var r io.Reader
		if v.CompressedSize != v.Size {
			r, err = xz.NewReader(res.Body,0)
			if err != nil {
				panic(err)
			}
		} else {
			r = res.Body
		}

		t := filepath.Join(*folder, filepath.Dir(v.Name))

		//change path to /lib
		f := filepath.Join(*folder, "lib", path.Base(v.Name))
		os.MkdirAll(t, os.ModeDir)
		filef, err := os.OpenFile(f, os.O_APPEND|os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0)
		if err != nil {
			panic(err)
		}
		io.Copy(filef, r)
		fmt.Println(f, "downloaded")

	}
	extractCSproj(folder)
	setUpCsProj(folder)
	/*res, err := http.Get(`http://runtime.fivem.net/client/prod/content/fivereborn/citizen/clr2/lib/mono/4.5/`+fileel+`.xz`)
	if err != nil {
		panic(err)
	}

	filee, err := os.OpenFile(path.Join(*folder, "/libs/"+fileel), os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0)
	//defer filee.Close()

	buf := new(bytes.Buffer)

	buf.ReadFrom(res.Body)
	r, err := xz.NewReader(buf, 0)
	if err != nil {
		fmt.Println(fileel,res.StatusCode,res.Request)
		log.Fatal(err)
	}
	io.Copy(filee, r)*/

}

type Tempvars struct {
	Safeprojectname string
	Guid1           string
}

func extractCSproj(folder *string) {
	err := os.MkdirAll(path.Join(*folder, "src"), os.ModeDir)
	if err != nil {
		panic(err)
	}
	/*
		lAssetFiles, err := AssetDir("data/template")
		l2AssetFile, err := AssetDir("data/template/Properties")
		for _, v := range l2AssetFile {
			if v != "Properties" {
				lAssetFiles = append(lAssetFiles, v)
			}

		}

		if err != nil {
			panic(err)
		}
		for _, v := range lAssetFiles {
			if _, err := os.Stat(path.Join(*folder, "src", path.Dir(v))); os.IsNotExist(err) {
				os.MkdirAll(v, os.ModeDir)
			}
			Safeprojectname := Tempvars{
				Safeprojectname: *folder, Guid1: "hjhj",
			}
			fmt.Println(v)
			csprojTemplate := template.New(v)
			csprojTemplate.Parse(string(MustAsset("data/template/" + v)))
			if err != nil {
				panic(err)
			}
			buff := bytes.Buffer{}
			err = csprojTemplate.Execute(&buff, Safeprojectname)
			if err != nil {
				panic(err)
			}
			newName := strings.Replace(v, "template", *folder, -1)
			newFile, err := os.OpenFile(path.Join(*folder, "src", newName), os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0)
			if err != nil {
				log.Fatal(err)
				continue
			}
			io.Copy(newFile, &buff)
		}*/

	data, err := Asset("data/template.zip")
	if err != nil {
		panic(err)
	}
	zipReader, err := zip.NewReader(bytes.NewReader(data), int64(binary.Size(data)))
	if err != nil {
		panic(err)
	}
	for _, file := range zipReader.File {
		if file.Mode().IsDir() {
			err := os.MkdirAll(filepath.Join(*folder, "src", file.Name), os.ModeDir)
			if err != nil {
				panic(err)
			}
			continue
		}
		fmt.Println(file.Name, "extarcting ")
		lsrcreader, err := file.Open()
		if err != nil {
			panic(err)
		}

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
		bytesres, err := ioutil.ReadAll(lsrcreader)
		if err != nil {
			panic(err)
		}
		csprojTemplate.Parse(string(bytesres))
		if err != nil {
			panic(err)
		}
		buff := bytes.Buffer{}
		err = csprojTemplate.Execute(&buff, Safeprojectname)
		if err != nil {
			panic(err)
		}
		newDstFileName := strings.Replace(file.Name, "template", *folder, -1)
		newFile, err := os.OpenFile(path.Join(*folder, "src", newDstFileName), os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0)
		if err != nil {
			log.Fatal(err)
			continue
		}
		//dstFile, err := os.OpenFile(filepath.Join(*folder, "src", file.Name), os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0)
		if err != nil {
			panic(err)
		}
		io.Copy(newFile, &buff)
	}

}

func setUpCsProj(folder *string) {

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
