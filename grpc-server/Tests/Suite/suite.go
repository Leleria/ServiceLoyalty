package Suite

//import (
//	"context"
//	sl "github.com/Leleria/Contract/GeneratedFilesProtoBufGo"
//	"github.com/Leleria/ServiceLoyalty/grpc-server/Config"
//	db "github.com/Leleria/ServiceLoyalty/grpc-server/Internal/Storage/SQLite"
//	"google.golang.org/grpc"
//	"google.golang.org/grpc/credentials/insecure"
//	"net"
//	"strconv"
//	"testing"
//)

//type Suite struct {
//	*testing.T
//	Cfg                  *Config.Config
//	LoyaltyServiceClient sl.LoyaltyServiceClient
//	DB                   *db.Storage
//}
//
//const (
//	grpcHost = "localhost"
//)
//
//func New(t *testing.T) (context.Context, *Suite) {
//	t.Helper()
//	t.Parallel()
//
//	//cfg := Config.MustLoadByPath("E:/ServiceLoyalty/Config/local_tests.yaml")
//
//	//ctx, cancelCtx := context.WithTimeout(context.Background(), cfg.GRPC.Timeout)
//
//	t.Cleanup(func() {
//		t.Helper()
//		cancelCtx()
//	})
//
//	database, err := db.New("E:/ServiceLoyalty/Storage/ServiceLoyaltyDB.db")
//	if err != nil {
//		t.Fatalf("created database failed: %v", err)
//	}
//	cc, err := grpc.DialContext(context.Background(),
//		grpcAddress(cfg),
//		grpc.WithTransportCredentials(insecure.NewCredentials()))
//	if err != nil {
//		t.Fatalf("grpc server connection failed: %v", err)
//	}
//
//	return ctx, &Suite{
//		T:                    t,
//		Cfg:                  cfg,
//		LoyaltyServiceClient: sl.NewLoyaltyServiceClient(cc),
//		DB:                   database,
//	}
//}
//
//func grpcAddress(cfg *Config.Config) string {
//	return net.JoinHostPort(grpcHost, strconv.Itoa(cfg.GRPC.Port))
//}
