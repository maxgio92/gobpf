// Copyright 2017 Kinvolk
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"fmt"
	"os"

	"github.com/maxgio92/gobpf/pkg/tracepipe"
)

func main() {
	tp, err := tracepipe.New()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
	defer tp.Close()

	channel, errorChannel := tp.Channel()

	for {
		select {
		case event := <-channel:
			fmt.Printf("%+v\n", event)
		case err := <-errorChannel:
			fmt.Printf("%+v\n", err)
		}
	}
}
