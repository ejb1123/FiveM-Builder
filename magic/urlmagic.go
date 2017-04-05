package magic

import (
	"encoding/xml"
	"net/http"
	"os"
	"log"
	"io/ioutil"
)

type Cache_info struct {
	XMLName xml.Name `xml:"CacheInfo"`
	Content []ContentFile `xml:"ContentFile"`
}
type ContentFile struct {
	XMLName xml.Name        `xml:"ContentFile"`
	CompressedSize int `xml:"CompressedSize,attr"`
	SHA1Hash       string`xml:"SHA1Hash,attr"`
	Size           int`xml:"Size,attr"`
	Name string`xml:"Name,attr"`
}

func Parsexml() *Cache_info{
	res, err := http.Get(`https://runtime.fivem.net/client/prod/content/fivereborn/info.xml`)
	defer res.Body.Close()
	if err != nil {
		log.Fatal("failed to get data")
		os.Exit(1)
	}
	bytesl, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	intt := Cache_info{}
	err = xml.Unmarshal(bytesl, &intt)
	if err != nil {
		panic(err)
	}
	return &intt
}
