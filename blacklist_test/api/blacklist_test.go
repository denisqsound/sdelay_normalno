package api

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

var AddBlacklistParams = []Blacklist{{ClientID: 132,PoolID: -1,ExpireAt: },
}


func TestBlacklist(t *testing.T) {

	api := new(BlacklistApi)

	Convey("Given some integer with a starting value", t, func() {
		Blacklist := api.AddToBlacklist()

		Convey("When the integer is incremented", func() {


			Convey("The value should be greater by one", func() {
				So(x, ShouldEqual, 2)
			})
		})
	})
}
