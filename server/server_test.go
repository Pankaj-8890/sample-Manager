package main

import (
	"context"
	mocks "sampleManager/mocks"
	pb "sampleManager/proto"
	"testing"

	gomock "github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)


func TestCreateMapping(t *testing.T) {

    ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    mockServer := mocks.NewMockSampleManagerServer(ctrl)

    testCases := []struct {
        name           string
        req            *pb.CreateRequest
        expectedErr    error
        expectedErrMsg string
    }{
        {
            name: "Successful mapping creation",
            req: &pb.CreateRequest{
                Segments:     []string{"segment1", "segment2"},
                SampleItemId: "sampleItem1",
                ItemId:       "item1",
            },
            expectedErr:    nil,
            expectedErrMsg: "",
        },
    }

    for _, tc := range testCases {
        t.Run(tc.name, func(t *testing.T) {
    
            mockServer.EXPECT().CreateMapping(gomock.Any(), tc.req).Return(&pb.CreateResponse{Message: "Mapping creation successfully"}, nil)

            resp, err := mockServer.CreateMapping(context.Background(), tc.req)

            if tc.expectedErr != nil {
                assert.Error(t, err)
                assert.Contains(t, err.Error(), tc.expectedErrMsg)
            } else {
                assert.NoError(t, err)
                assert.NotNil(t, resp)
                assert.Equal(t, "Mapping creation successfully", resp.Message)
            }
        })
    }
}

func TestGetSampleId(t *testing.T) {
    ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    mockServer := mocks.NewMockSampleManagerServer(ctrl)

    testCases := []struct {
        name            string
        req             *pb.GetRequest
        mockServerBehavior func()
        expectedError   error
        expectedResult  *pb.GetResponse
    }{
        {
            name: "Mapping found",
            req: &pb.GetRequest{ItemId: "validItemId"},
            mockServerBehavior: func() {
                mockServer.EXPECT().GetSampleId(gomock.Any(), &pb.GetRequest{ItemId: "validItemId"}).Return(&pb.GetResponse{SampleItemId: "sampleItemID"}, nil)
            },
            expectedError:  nil,
            expectedResult: &pb.GetResponse{SampleItemId: "sampleItemID"},
        },
        {
            name: "No mapping found",
            req: &pb.GetRequest{ItemId: "nonexistentItemId"},
            mockServerBehavior: func() {
                mockServer.EXPECT().GetSampleId(gomock.Any(), &pb.GetRequest{ItemId: "nonexistentItemId"}).Return(nil, status.Error(codes.Unavailable, "No mapping found: record not found"))
            },
            expectedError: status.Error(codes.Unavailable, "No mapping found: record not found"),
            expectedResult: nil,
        },
    
    }

    for _, tc := range testCases {
        t.Run(tc.name, func(t *testing.T) {
    
            tc.mockServerBehavior()

            resp, err := mockServer.GetSampleId(context.Background(), tc.req)
            if tc.expectedError != nil {
                assert.Equal(t, tc.expectedError, err)
            } else {
                assert.NoError(t, err)
                assert.Equal(t, tc.expectedResult, resp)
            }
        })
    }
}
