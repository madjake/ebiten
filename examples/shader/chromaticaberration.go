// Copyright 2020 The Ebiten Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// +build ignore

package main

var Time float
var Cursor vec2
var ImageSize vec2

func Fragment(position vec4, texCoord vec2, color vec4) vec4 {
	center := ImageSize / 2
	amount := (center - Cursor) / texture2Size() / 10
	var clr vec3
	clr.r = texture2At(vec2(texCoord.x+amount.x, texCoord.y)).r
	clr.g = texture2At(texCoord).g
	clr.b = texture2At(vec2(texCoord.x-amount.x, texCoord.y)).b
	return vec4(clr, 1.0)
}