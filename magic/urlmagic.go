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
