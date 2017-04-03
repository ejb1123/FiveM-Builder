package magic

import (
	"os"
	"fmt"
	"io/ioutil"
	"net/http"
	"path"

	"bytes"
	"log"
	"xi2.org/x/xz"
	"io"
	"gopkg.in/go-playground/validator.v8"
)

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

	res, err := http.Get(`http://runtime.fivem.net/client/prod/content/fivereborn/citizen/clr2/lib/mono/4.5/CitizenFX.Core.dll.xz`)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	os.MkdirAll(path.Join(*folder, "/libs"), os.ModeDir)
	filee, err := os.OpenFile(path.Join(*folder, "/libs/CitizenFX.Core.dll"), os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0)
	defer filee.Close()

	buf := new(bytes.Buffer)

	buf.ReadFrom(res.Body)
	dd,err:=ioutil.ReadAll(res.Body)
	r, err := xz.NewReader(dd, 0)
	if err != nil {
		log.Fatal(err)
	}
	j := r
	io.Copy(filee, j)

}
