package lab2

import (
	. "gopkg.in/check.v1"
)

type PrefixComputeSuite struct{}

var _ = Suite(&PrefixComputeSuite{})

func (s *PrefixComputeSuite) TestCalculatePrefix(c *C) {
	res, err := EvaluatePrefix("* + 2 3 5")
	c.Assert(err, IsNil)
	c.Assert(res, Equals, "25")
}

func (s *PrefixComputeSuite) TestCalculatePrefixError(c *C) {
	res, err := EvaluatePrefix("+ 2 invalid")
	c.Assert(err, NotNil)
	c.Assert(res, Equals, "")
}
