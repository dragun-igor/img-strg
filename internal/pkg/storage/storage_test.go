package storage

import (
	"errors"
	"strconv"
	"testing"
	"time"

	"github.com/dragun-igor/img-strg/internal/pkg/storage/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
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
	require.NoError(s.T(), err)

	s.redis.EXPECT().Get(key).Return([]byte(strconv.Itoa(int(value.Unix()))), nil)
	res, err := s.storage.GetBirthTimeFile(key)
	require.NoError(s.T(), err)
	require.Equal(s.T(), time.Unix(value.Unix(), 0), res)

	s.redis.EXPECT().Get(key).Return([]byte{}, testErr)
	res, err = s.storage.GetBirthTimeFile(key)
	require.EqualError(s.T(), err, testErr.Error())
	require.Equal(s.T(), time.Time{}, res)

	s.redis.EXPECT().Get(key).Return([]byte("invalid response"), nil)
	res, err = s.storage.GetBirthTimeFile(key)
	require.Error(s.T(), err)
	require.Equal(s.T(), time.Time{}, res)
}
