package testsuite

import (
	"time"

	"github.com/stretchr/testify/suite"
	db "github.com/frazercomputing/upper-io-db"
	"github.com/frazercomputing/upper-io-db/lib/sqlbuilder"
)

const TimeZone = "Canada/Eastern"

var TimeLocation, _ = time.LoadLocation(TimeZone)

type Helper interface {
	SQLBuilder() sqlbuilder.Database
	Session() db.Database

	Adapter() string

	TearUp() error
	TearDown() error
}

type Suite struct {
	suite.Suite

	Helper
}

func (s *Suite) AfterTest(suiteName, testName string) {
	err := s.TearDown()
	s.NoError(err)
}

func (s *Suite) BeforeTest(suiteName, testName string) {
	err := s.TearUp()
	s.NoError(err)
}
