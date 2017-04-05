package magic

import (
	"os"
	"io/ioutil"
	"path"
	"path/filepath"
	"net/http"
	"github.com/ulikunitz/xz"
	"io"
	"fmt"
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
	os.MkdirAll(path.Join(*folder, "/libs"), os.ModeDir)
	allfiles := Parsexml()
	ffiles := reduceFiles(allfiles)
	for _, v := range *ffiles {
		extra:=""
		if v.CompressedSize!=v.Size{
			extra=".xz"
		}
		res, err := http.Get(`http://runtime.fivem.net/client/prod/content/fivereborn/` + filepath.ToSlash(v.Name) + extra)
		if err != nil {
			panic(err)
		}
		var r io.Reader
		if v.CompressedSize!=v.Size{
			r, err = xz.NewReader(res.Body)
			if err != nil {
				panic(err)
			}
		}else {
			r=res.Body
		}

		t := filepath.Join(*folder, filepath.Dir(v.Name))
		f := filepath.Join(*folder, v.Name)
		os.MkdirAll(t, os.ModeDir)
		filef, err := os.OpenFile(f, os.O_APPEND|os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0)
		if err!=nil{
			panic(err)
		}
		io.Copy(filef, r)
	}
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

type Directory struct {
	path string;
}

func (d Directory) f() {
	d.path = "k"
}
func reduceFiles(info *Cache_info) *[]ContentFile {
	filesretuned := []ContentFile{}
	for _, v := range info.Content {
		jj := filepath.Dir(v.Name)
		if filepath.ToSlash(filepath.Dir(v.Name)) == filepath.ToSlash(`citizen/clr2/lib/mono/4.5`) {
			filesretuned = append(filesretuned, v)
		}
	}
	return &filesretuned
}
