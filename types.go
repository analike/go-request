/**
 * @package go-request (2026)
 * @author Emmanuel Analike <emmanuel@analike.dev>
 * @created Feb 18, 2026; 2:20 PM
 */

package request

import "net/http"

type Values map[string][]string

func (v *Values) Has(key string) bool {
	_, ok := (*v)[key]
	return ok
}

func (v *Values) GetSingle(key string) string {
	val := (*v)[key]
	if len(val) > 0 {
		return val[0]
	}
	return ""
}

func (v *Values) GetMultiple(key string) []string {
	return (*v)[key]
}

type Header http.Header

type File struct {
	/**
	 * Field name in the form
	 */
	FieldName string
	/**
	 * Name of the file on the user's computer
	 */
	OriginalName string
	/**
	 * Encoding of the file
	 */
	Encoding string
	/**
	 * Mime type of the file
	 */
	Mimetype string
	/**
	 * Size of the file in bytes
	 */
	Size int64
	/**
	 * The folder to which the file has been saved
	 */
	Destination string
	/**
	 * The name of the file within the destination
	 */
	FileName string
	/**
	 * The full path to the uploaded file
	 */
	Path string
	/**
	 * Bytes of the file
	 */
	Bytes *[]byte
}
