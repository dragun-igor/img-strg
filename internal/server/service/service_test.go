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
	"google.golang.org/protobuf/types/known/emptypb"
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
	fileName := "image.name"
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
	ctx := context.Background()
	fileName := "image.name"
	r := &strg.SendImageRequest{
		Name:  fileName,
		Image: []byte{1, 2, 3, 4, 5, 6},
	}

	s.repo.EXPECT().SetBirthTimeFile(fileName, gomock.Any()).Return(nil)
	_, err := s.service.SendImage(ctx, r)
	s.Require().NoError(err)

	res, err := s.service.GetImage(ctx, &strg.GetImageRequest{Name: "invalid name"})
	s.Require().Error(err)
	s.Require().Nil(res)

	res, err = s.service.GetImage(ctx, &strg.GetImageRequest{Name: fileName})
	s.Require().NoError(err)
	s.Require().Equal([]byte{1, 2, 3, 4, 5, 6}, res.Image)

	err = os.Remove(s.path + fileName)
	s.Require().NoError(err)
}

func (s *ServiceSuite) TestGetImagesList() {
	ctx := context.Background()
	fileName := "image.name"
	r := &strg.SendImageRequest{
		Name:  fileName,
		Image: []byte{1, 2, 3, 4, 5, 6},
	}
	testPath := "./test/"

	// Тест с несуществующей папкой
	oldPath := s.service.storagePath
	s.service.storagePath = "./invalidnamefolder"
	res, err := s.service.GetImagesList(ctx, &emptypb.Empty{})
	s.Require().Error(err)
	s.Require().Nil(res)
	s.service.storagePath = oldPath

	err = os.Mkdir(testPath, 0750)
	s.Require().NoError(err)

	// Тест с пустой папкой
	s.service.storagePath = testPath
	res, err = s.service.GetImagesList(ctx, &emptypb.Empty{})
	s.Require().NoError(err)
	s.Require().Equal([]*strg.Images{}, res.Images)

	s.repo.EXPECT().SetBirthTimeFile(fileName, gomock.Any()).Return(nil)
	_, err = s.service.SendImage(ctx, r)
	s.Require().NoError(err)

	s.repo.EXPECT().GetBirthTimeFile(fileName)
	res, err = s.service.GetImagesList(ctx, &emptypb.Empty{})
	s.Require().NoError(err)
 s.Require().Equal(1, len(res.Images))
	s.Require().Equal(fileName, res.Images[0].Name)

	err = os.RemoveAll(s.service.storagePath)
	s.Require().NoError(err)
}
