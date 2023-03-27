package testserv

import (
	"context"
	"flag"
	"log"
	"os"
	"strconv"
	"testing"

	roc "github.com/fronomenal/go_rpcket/protos/v1"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	client roc.RocketServiceClient
)

type RocketTestSuite struct {
	suite.Suite
}

func TestMain(m *testing.M) {
	//setup
	addr := flag.String("addr", "localhost:51515", "the listening address of the client")
	flag.Parse()
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("could not connect to %s: %v", "51515", err)
	}
	client = roc.NewRocketServiceClient(conn)

	// run tests
	code := m.Run()

	//teardown
	conn.Close()

	//exit
	os.Exit(code)
}

func (s *RocketTestSuite) TestAddRocket() {

	var sent int32
	s.T().Run("adds a new rocket successfully", func(t *testing.T) {
		wantName := "Test Rocket"
		r, err := client.SetRocket(context.Background(), &roc.SetReq{Rocket: &roc.Rocket{Id: 99, Name: wantName, Type: "Interstellar Ship", Flights: 1}})
		assert.NoError(s.T(), err)
		assert.Equal(s.T(), wantName, r.Rocket.Name)
		sent = r.Rocket.Id
	})

	s.T().Run("gets a rocket successfully", func(t *testing.T) {
		var wantId int32 = sent
		r, err := client.GetRocket(context.Background(), &roc.GetReq{Id: int32(wantId)})
		assert.NoError(s.T(), err)
		assert.Equal(s.T(), wantId, r.Rocket.Id)
	})

	s.T().Run("removes a rocket successfully", func(t *testing.T) {
		wantId := sent
		r, err := client.RemRocket(context.Background(), &roc.RemReq{Id: int32(wantId)})
		assert.NoError(s.T(), err)
		assert.Contains(s.T(), r.Status, strconv.Itoa(int(wantId)))
	})

}

func TestRocketService(t *testing.T) {
	suite.Run(t, new(RocketTestSuite))
}
