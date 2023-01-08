package test

import (
	"bytes"
	"context"
	"fmt"
	"image/jpeg"
	"io"
	"os"
	"testing"

	strg "github.com/dragun-igor/img-strg/proto/api"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/emptypb"
)

func TestClientFlow(t *testing.T) {
	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	require.NoError(t, err)

	ctx := context.Background()
	service := strg.NewImageStorageClient(conn)

	fileImage1, err := os.Open("./images/bimage1")
	if err != nil {
		fmt.Println(err)
	}
	defer fileImage1.Close()
	bimage1, err := io.ReadAll(fileImage1)
	require.NoError(t, err)
	_, err = service.SendImage(ctx, &strg.SendImageRequest{
		Name:  "photo.jpg",
		Image: bimage1,
	})
	require.NoError(t, err)
	list1, err := service.GetImagesList(ctx, &emptypb.Empty{})
	require.NoError(t, err)

	fileImage2, err := os.Open("./images/bimage2")
	if err != nil {
		fmt.Println(err)
	}
	defer fileImage2.Close()
	bimage2, err := io.ReadAll(fileImage2)
	require.NoError(t, err)
	_, err = service.SendImage(ctx, &strg.SendImageRequest{
		Name:  "photo.jpg",
		Image: bimage2,
	})
	require.NoError(t, err)
	list2, err := service.GetImagesList(ctx, &emptypb.Empty{})
	require.NoError(t, err)
	require.Equal(t, list1.Images[0].CreationTime, list2.Images[0].CreationTime)
	require.NotEqual(t, list1.Images[0].ModificationTime, list2.Images[0].ModificationTime)
	res, err := service.GetImage(ctx, &strg.GetImageRequest{Name: "photo.jpg"})
	require.NoError(t, err)
	decoded1, err := jpeg.Decode(bytes.NewReader(bimage1))
	require.NoError(t, err)
	decoded2, err := jpeg.Decode(bytes.NewReader(bimage2))
	require.NoError(t, err)
	decodedActual, err := jpeg.Decode(bytes.NewReader(res.Image))
	require.NoError(t, err)
	require.NotEqual(t, decoded1, decodedActual)
	require.Equal(t, decoded2, decodedActual)
}
