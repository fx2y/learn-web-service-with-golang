package memory

import (
	"github.com/fx2y/learn-web-service-with-golang/internal/linkgraph/graph/graphtest"
	gc "gopkg.in/check.v1"
	"testing"
)

var _ = gc.Suite(new(InMemoryGraphTestSuite))

func Test(t *testing.T) { gc.TestingT(t) }

type InMemoryGraphTestSuite struct {
	graphtest.SuiteBase
}

func (s *InMemoryGraphTestSuite) SetUpTest(c *gc.C) {
	s.SetGraph(NewInMemoryGraph())
}
