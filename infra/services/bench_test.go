package services

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type BenchTestSuite struct {
	suite.Suite
}

func (s *BenchTestSuite) TestRequestUrl() {
	status, err := requestUrl("http://www.google.com")
	s.NoError(err)
	s.Equal(200, status)
}

func (s *BenchTestSuite) TestBenchUrl() {
	result, err := BenchUrl("https://httpbin.org/status/200", 1000, 10)
	s.NoError(err)
	s.NotNil(result)
}

func TestBenchRun(t *testing.T) {
	suite.Run(t, new(BenchTestSuite))
}
