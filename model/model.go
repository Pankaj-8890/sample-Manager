package model

type Segment struct {
	ID        uint64 `gorm:"primaryKey"`
	Segment   string
	MappingID uint64
}

type Mapping struct {
	ID          uint64 `gorm:"primaryKey"`
	SampleItemID string
	ItemID       string
	Segments     []Segment `gorm:"foreignKey:MappingID"`
}


// package main

// import (
// 	"context"
// 	"log"
// 	"net"
// 	"testing"
// 	"sampleManager/model"
// 	"github.com/stretchr/testify/assert"
// 	"google.golang.org/grpc"
// 	"gorm.io/driver/sqlite"
// 	"gorm.io/gorm"

// 	pb "sampleManager/proto"
// )

// func TestCreateMapping(t *testing.T) {
// 	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
// 	if err != nil {
// 		log.Fatalf("Failed to connect to test database: %v", err)
// 	}
	
// 	err = db.AutoMigrate(&model.Mapping{}, &model.Segment{})
// 	if err != nil {
// 		log.Fatalf("Failed to migrate database schema: %v", err)
// 	}

// 	srv := grpc.NewServer()
// 	pb.RegisterSampleManagerServer(srv, &Server{DB: db})
	
// 	lis, err := net.Listen("tcp", ":0")
// 	if err != nil {
// 		t.Fatalf("Failed to listen: %v", err)
// 	}
// 	go func() {
// 		if err := srv.Serve(lis); err != nil {
// 			log.Fatalf("Failed to serve: %v", err)
// 		}
// 	}()
// 	defer srv.Stop()
	
// 	conn, err := grpc.Dial(lis.Addr().String(), grpc.WithInsecure())
// 	if err != nil {
// 		t.Fatalf("Failed to dial server: %v", err)
// 	}
// 	defer conn.Close()
// 	client := pb.NewSampleManagerClient(conn)
	
// 	itemID := "testItemId"
// 	createReq := &pb.CreateRequest{
// 		ItemId:      itemID,
// 		SampleItemId: "sampleItemId123",
// 		Segments: []string{"segment1", "segment2"},
// 	}
	
// 	resp, err := client.CreateMapping(context.Background(), createReq)
	

// 	assert.NoError(t, err)
// 	assert.NotNil(t, resp)
// 	assert.Equal(t, "Mapping creation successfully", resp.Message)
	
	
// 	var mapping model.Mapping
// 	result := db.Where("item_id = ?", itemID).First(&mapping)
// 	assert.NoError(t, result.Error)
// 	assert.Equal(t, createReq.ItemId, mapping.ItemID)
// 	assert.Equal(t, createReq.SampleItemId, mapping.SampleItemID)
// }

