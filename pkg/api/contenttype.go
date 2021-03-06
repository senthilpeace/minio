/*
 * Mini Object Storage, (C) 2015 Minio, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package api

import (
	"net/http"
)

type contentType int

const (
	xmlType contentType = iota
	jsonType
)

// content-type to human readable map
var typeToString = map[contentType]string{
	xmlType:  "application/xml",
	jsonType: "application/json",
}

// human readbale to content-type map
var acceptToType = map[string]contentType{
	"application/xml":  xmlType,
	"application/json": jsonType,
}

// Get content type requested from 'Accept' header
func getContentType(req *http.Request) contentType {
	if accept := req.Header.Get("Accept"); accept != "" {
		return acceptToType[accept]
	}
	return xmlType
}

// Content type to human readable string
func getContentString(content contentType) string {
	return typeToString[content]
}
