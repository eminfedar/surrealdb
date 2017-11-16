// Copyright © 2016 Abcum Ltd
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

package fncs

import (
	"context"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestIf(t *testing.T) {

	var res interface{}

	Convey("if(a, b, c) works properly", t, func() {
		res, _ = Run(context.Background(), "if", true, "yes", "no")
		So(res, ShouldEqual, "yes")
		res, _ = Run(context.Background(), "if", "true", "yes", "no")
		So(res, ShouldEqual, "yes")
		res, _ = Run(context.Background(), "if", false, "yes", "no")
		So(res, ShouldEqual, "no")
		res, _ = Run(context.Background(), "if", "false", "yes", "no")
		So(res, ShouldEqual, "no")
		res, _ = Run(context.Background(), "if", "something", "yes", "no")
		So(res, ShouldEqual, "no")
	})

}
