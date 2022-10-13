package zetaclient

import (
	. "gopkg.in/check.v1"
)

type DummyTssSignerSuite struct {
	signer *TestSigner
}

var _ = Suite(&DummyTssSignerSuite{})

func (s *DummyTssSignerSuite) SetUpTest(c *C) {
	signer, err := NewTestSigner("fad9c8855b740a0b7ed4c221dbad0f33a83a49cad6b3fe8d5817ac83d38b6a19")

	c.Assert(err, IsNil)
	s.signer = signer
}

func (s *DummyTssSignerSuite) TestSign(c *C) {
	ok := s.signer.TestKeysign()
	c.Assert(ok, Equals, true)
}
