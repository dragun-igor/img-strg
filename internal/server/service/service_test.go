package service

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/dragun-igor/img-strg/internal/server/service/mocks"
	strg "github.com/dragun-igor/img-strg/proto/api"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"
)

type ServiceSuite struct {
	suite.Suite

	ctrl    *gomock.Controller
	repo    *mocks.MockStorage
	service *Service
	path    string
}

func (s *ServiceSuite) SetupTest() {
	s.ctrl = gomock.NewController(s.T())
	s.repo = mocks.NewMockStorage(s.ctrl)
	s.path = "./"
	service, err := New(s.repo, s.path)
	s.Require().NoError(err)
	s.service = service
}

func (s *ServiceSuite) TearDownTest() {
	s.ctrl.Finish()
}

func TestServiceSuite(t *testing.T) {
	suite.Run(t, new(ServiceSuite))
}

func (s *ServiceSuite) TestSendImage() {
	ctx := context.Background()
	fileName := "test.name"
	r := &strg.SendImageRequest{
		Name:  fileName,
		Image: []byte{},
	}
	testErr := errors.New("test")

	s.repo.EXPECT().SetBirthTimeFile(fileName, gomock.Any()).Return(testErr)
	_, err := s.service.SendImage(ctx, r)
	s.Require().EqualError(err, testErr.Error())
	err = os.Remove(s.path + fileName)
	s.Require().NoError(err)

	s.repo.EXPECT().SetBirthTimeFile(fileName, gomock.Any()).Return(nil)
	_, err = s.service.SendImage(ctx, r)
	s.Require().NoError(err)

	_, err = s.service.SendImage(ctx, r)
	s.Require().NoError(err)
	err = os.Remove(s.path + fileName)
	s.Require().NoError(err)
}

func (s *ServiceSuite) TestGetImage() {
	// todo
}

func (s *ServiceSuite) TestGetImagesList() {
	// todo
}
