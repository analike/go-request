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
	err     error
}

func (d *Data) loadData() error {
	dd := *d
	if dd.init == true {
		return d.err
	}
	ctx := dd.context
	switch t := ctx.(type) {
	case *gin.Context:
		*dd.data, dd.err = (*t).GetRawData()
	}
	return d.err
}

func (d *Data) ToJSON(destInterface *any) error {
	if (d.err != nil) && (d.init == false) {
	}
	return json.Unmarshal(*d.data, destInterface)
}

func FromGin(c *gin.Context) *Data {
	return &Data{context: c}
}
