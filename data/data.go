/**
 * @package go-request (2026)
 * @author Emmanuel Analike <emmanuel@analike.dev>
 * @created Feb 18, 2026; 3:41 PM
 */

package data

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
)

type Data struct {
	data    *[]byte
	context any
	init    bool
	empty   bool
	err     error
}

func (d *Data) loadData() {
	dd := *d
	if dd.init == true {
		return
	}
	ctx := dd.context
	switch t := ctx.(type) {
	case *gin.Context:
		gData, gErr := (*t).GetRawData()
		if gErr != nil {
			dd.err = gErr
		} else {
			dd.data = &gData
		}
		// *dd.data, dd.err = (*t).GetRawData()
	}
	dd.init = true
	dd.empty = dd.data == nil
	// return d.err
}

func (d *Data) IsEmpty() bool {
	return d.empty
}

func (d *Data) Clear() {
	if !d.IsEmpty() {
		d.data = nil
	}
}

func (d *Data) ToJSON(destInterface any) error {
	if (d.err == nil) && (d.init == false) {
		d.loadData()
	}
	if d.err != nil {
		return d.err
	}
	return json.Unmarshal(*d.data, destInterface)
}

func FromGin(c *gin.Context) *Data {
	return &Data{context: c}
}
