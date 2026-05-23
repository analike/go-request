/**
 * @package go-request (2026)
 * @author Emmanuel Analike <emmanuel@analike.dev>
 * @created Feb 18, 2026; 3:41 PM
 */

package data

import (
	"encoding/json"
	"errors"

	"github.com/gin-gonic/gin"
)

type Data struct {
	data    *[]byte
	context any
	init    bool
	empty   bool
	err     error
}

func (d *Data) initiate() {
	if (d.err == nil) && (d.init == false) {
		d.loadData()
	}
}

func (d *Data) loadData() {
	if d.init == true {
		return
	}
	ctx := d.context
	switch t := ctx.(type) {
	case *gin.Context:
		gData, gErr := (*t).GetRawData()
		if gErr != nil {
			d.err = gErr
		} else {
			d.empty = len(gData) == 0
			d.data = &gData
		}
	}
	d.init = true
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
	d.initiate()
	if d.err != nil {
		return d.err
	}
	if d.data == nil {
		return errors.New("no request data")
	}
	return json.Unmarshal(*d.data, destInterface)
}

func (d *Data) ToString() string {
	d.initiate()
	if d.err != nil && d.data != nil {
		return string(*d.data)
	}
	return ""
}

func (d *Data) ToBytes() *[]byte {
	d.initiate()
	return d.data
}

func FromGin(c *gin.Context) *Data {
	return &Data{context: c}
}
