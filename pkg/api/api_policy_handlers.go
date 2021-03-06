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
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	mstorage "github.com/minio-io/minio/pkg/storage"
)

// PUT Bucket policy
// -----------------
// This implementation of the PUT operation uses the policy subresource
// to add to or replace a policy on a bucket
func (server *minioAPI) putBucketPolicyHandler(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	bucket := vars["bucket"]
	acceptsContentType := getContentType(req)

	policy, ok := mstorage.Parsepolicy(req.Body)
	if ok == false {
		error := errorCodeError(InvalidPolicyDocument)
		errorResponse := getErrorResponse(error, bucket)
		w.WriteHeader(error.HTTPStatusCode)
		w.Write(writeErrorResponse(w, errorResponse, acceptsContentType))
		return
	}

	err := server.storage.CreateBucketPolicy(bucket, policy)
	switch err := err.(type) {
	case nil:
		{
			w.WriteHeader(http.StatusNoContent)
			writeCommonHeaders(w, getContentString(acceptsContentType))
			w.Header().Set("Connection", "keep-alive")
		}
	case mstorage.BucketNameInvalid:
		{
			error := errorCodeError(InvalidBucketName)
			errorResponse := getErrorResponse(error, bucket)
			w.WriteHeader(error.HTTPStatusCode)
			w.Write(writeErrorResponse(w, errorResponse, acceptsContentType))
		}
	case mstorage.BucketNotFound:
		{
			error := errorCodeError(NoSuchBucket)
			errorResponse := getErrorResponse(error, bucket)
			w.WriteHeader(error.HTTPStatusCode)
			w.Write(writeErrorResponse(w, errorResponse, acceptsContentType))
		}
	case mstorage.BackendCorrupted:
		{
			log.Println(err)
			error := errorCodeError(InternalError)
			errorResponse := getErrorResponse(error, bucket)
			w.WriteHeader(error.HTTPStatusCode)
			w.Write(writeErrorResponse(w, errorResponse, acceptsContentType))
		}
	case mstorage.ImplementationError:
		{
			log.Println(err)
			error := errorCodeError(InternalError)
			errorResponse := getErrorResponse(error, bucket)
			w.WriteHeader(error.HTTPStatusCode)
			w.Write(writeErrorResponse(w, errorResponse, acceptsContentType))
		}
	}
}

// GET Bucket policy
// -----------------
// This implementation of the GET operation uses the policy subresource
// to return the policy of a specified bucket.
func (server *minioAPI) getBucketPolicyHandler(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	bucket := vars["bucket"]
	acceptsContentType := getContentType(req)

	p, err := server.storage.GetBucketPolicy(bucket)
	switch err := err.(type) {
	case nil:
		{
			responsePolicy, ret := json.Marshal(p)
			if ret != nil {
				error := errorCodeError(InternalError)
				errorResponse := getErrorResponse(error, bucket)
				w.WriteHeader(error.HTTPStatusCode)
				w.Write(writeErrorResponse(w, errorResponse, acceptsContentType))
			}
			writeCommonHeaders(w, getContentString(acceptsContentType))
			w.Header().Set("Connection", "keep-alive")
			w.Write(responsePolicy)
		}
	case mstorage.BucketNameInvalid:
		{
			error := errorCodeError(InvalidBucketName)
			errorResponse := getErrorResponse(error, bucket)
			w.WriteHeader(error.HTTPStatusCode)
			w.Write(writeErrorResponse(w, errorResponse, acceptsContentType))
		}
	case mstorage.BucketNotFound:
		{
			error := errorCodeError(NoSuchBucket)
			errorResponse := getErrorResponse(error, bucket)
			w.WriteHeader(error.HTTPStatusCode)
			w.Write(writeErrorResponse(w, errorResponse, acceptsContentType))
		}
	case mstorage.BucketPolicyNotFound:
		{
			error := errorCodeError(NoSuchBucketPolicy)
			errorResponse := getErrorResponse(error, bucket)
			w.WriteHeader(error.HTTPStatusCode)
			w.Write(writeErrorResponse(w, errorResponse, acceptsContentType))
		}
	case mstorage.BackendCorrupted:
		{
			log.Println(err)
			error := errorCodeError(InternalError)
			errorResponse := getErrorResponse(error, bucket)
			w.WriteHeader(error.HTTPStatusCode)
			w.Write(writeErrorResponse(w, errorResponse, acceptsContentType))
		}
	case mstorage.ImplementationError:
		{
			log.Println(err)
			error := errorCodeError(InternalError)
			errorResponse := getErrorResponse(error, bucket)
			w.WriteHeader(error.HTTPStatusCode)
			w.Write(writeErrorResponse(w, errorResponse, acceptsContentType))
		}
	}
}
