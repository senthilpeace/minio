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

package file

import (
	"os"
	"sync"
)

// Start filesystem channel
func Start(root string) (chan<- string, <-chan error, *Storage) {
	ctrlChannel := make(chan string)
	errorChannel := make(chan error)
	s := Storage{}
	s.root = root
	s.lock = new(sync.Mutex)
	go start(ctrlChannel, errorChannel, &s)
	return ctrlChannel, errorChannel, &s
}

func start(ctrlChannel <-chan string, errorChannel chan<- error, s *Storage) {
	err := os.MkdirAll(s.root, 0700)
	errorChannel <- err
	close(errorChannel)
}
