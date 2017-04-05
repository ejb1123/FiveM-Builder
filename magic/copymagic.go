package magic

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"gopkg.in/yaml.v2"
	"path/filepath"
	"path"
	"os/exec"
	"time"
	"io"
)

func GetFiles(src *string) *Tempfiles {
	files := Tempfiles{}
	/*if hasLuaResourceFile(src) {
		luaFile := path.Join(*src, "__resource.lua")
		Parselua(&luaFile)
	} else {*/
		_, err := os.Stat(*src)
		if err != nil {
			fmt.Println("err")
		}

		err = filepath.Walk(*src, func(path string, info os.FileInfo, err error) error {
			res, _ := os.Stat(path)
			if filepath.IsAbs(path){
				path,err = filepath.Rel(*src,path)
				if(err!=nil){
					log.Fatal(err)
					return nil
				}
			}
			files.files = append(files.files, File{src: path, isFile: !res.IsDir()})
			return nil
		})
		if err != nil {
			panic(err)
		}
	//}

	return &files
}

func DoCopy(tempfiles *Tempfiles, src *string, root *string, projectName *string) {
	for _, q := range tempfiles.files {
		pathn := filepath.FromSlash(path.Join(filepath.FromSlash(path.Join(*root, "resources", *projectName)), filepath.FromSlash(q.src)))
		if (q.isFile) {
			dstFile, err := os.OpenFile(pathn, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0)
			if (err != nil) {
				log.Fatal(err)
				continue
			}
			srcFile, err := os.OpenFile(pathn, os.O_RDONLY, 0)
			if (err != nil) {
				log.Fatal(err)
				continue
			}
			io.Copy(dstFile, srcFile)

		} else {
			if _, err := os.Stat(pathn); os.IsNotExist(err) {
				os.MkdirAll(pathn, os.ModeDir)
			} else {
				if v, _ := os.Stat(pathn); v.Mode().IsRegular() {
					panic(v.Name() + "is a file and should be a directory\n Please manualy fix this and rerun.")
				}
			}
		}
	}
}

type File struct {
	src    string
	isFile bool
}

type Tempfiles struct {
	files []File
}
type T struct {
	Server struct {
		Enabled     bool
		Url         string `yaml:"url"`
		Password    string
		Src         string
		Root        string
		ProjectName string `yaml:"name"`
		IceCon      string        `yaml:"iceconpath"`
	}`yaml:"server"`
}

func ReadConfig(config string) *T {
	result, err := ioutil.ReadFile(config)
	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	g := T{}
	yaml.Unmarshal(result, &g)
	if g.Server.Enabled == false {
		os.Exit(0)
	}
	return &g
}

func hasLuaResourceFile(src *string) bool {
	luaFile := (path.Join(*src, "__resource.lua"))
	if v, _ := os.Stat(luaFile);!v.IsDir() && v.Name() == "__resource.lua" {
		return true
	} else {
		return false
	}
}

func RestartServer(url *string, password *string, projectName *string, iceconPath *string) {
	time.Sleep(1000)
	cmdd := exec.Command(*iceconPath, "-c restart " + *projectName, *url, *password)
	cmdd.Stdout = os.Stdout
	//hhh,_:=cmdd.Output()
	cmdd.Stderr = os.Stderr
	cmdd.Run()
}
