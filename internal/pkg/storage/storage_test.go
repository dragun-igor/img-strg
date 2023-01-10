package storage

import (
	"errors"
	"strconv"
	"testing"
	"time"

	"github.com/dragun-igor/img-strg/internal/pkg/storage/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"
)

type StorageSuite struct {
	suite.Suite

	ctrl    *gomock.Controller
	redis   *mocks.MockRedis
	storage *Storage
}

func (s *StorageSuite) SetupTest() {
	s.ctrl = gomock.NewController(s.T())
	s.redis = mocks.NewMockRedis(s.ctrl)
	s.storage = New(s.redis)
}

func (s *StorageSuite) TearDownTest() {
	s.ctrl.Finish()
}

func TestServiceSuite(t *testing.T) {
	suite.Run(t, new(StorageSuite))
}

func (s *StorageSuite) TestStorage() {
	key := "key"
	value := time.Now()
	testErr := errors.New("test")

	s.redis.EXPECT().Set(key, value.Unix()).Return(nil)
	err := s.storage.SetBirthTimeFile(key, value)
	s.Require().NoError(err)

	s.redis.EXPECT().Get(key).Return([]byte(strconv.Itoa(int(value.Unix()))), nil)
	res, err := s.storage.GetBirthTimeFile(key)
	s.Require().NoError(err)
	s.Require().Equal(time.Unix(value.Unix(), 0), res)

	s.redis.EXPECT().Get(key).Return([]byte{}, testErr)
	res, err = s.storage.GetBirthTimeFile(key)
	s.Require().EqualError(err, testErr.Error())
	s.Require().Equal(time.Time{}, res)

	s.redis.EXPECT().Get(key).Return([]byte("invalid response"), nil)
	res, err = s.storage.GetBirthTimeFile(key)
	s.Require().Error(err)
	s.Require().Equal(time.Time{}, res)
}
