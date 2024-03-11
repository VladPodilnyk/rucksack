package main

import (
	"fmt"
	_ "net/http/pprof"

	"github.com/vladpodilnyk/rucksack/kv/config"
	"github.com/vladpodilnyk/rucksack/kv/storage/standalone"
)

// var (
// 	schedulerAddr = flag.String("scheduler", "", "scheduler address")
// 	storeAddr     = flag.String("addr", "", "store address")
// 	dbPath        = flag.String("path", "", "directory path of db")
// 	logLevel      = flag.String("loglevel", "", "the level of log")
// )

func main() {
	config := config.Config{
		StoreAddr: "some address",
	}

	storage := standalone.NewStandaloneStorage(&config)

	if err := storage.Start(); err != nil {
		fmt.Println("Error")
	}
	// server := server.NewServer(storage)

	// var alivePolicy = keepalive.EnforcementPolicy{
	// 	MinTime:             2 * time.Second, // If a client pings more than once every 2 seconds, terminate the connection
	// 	PermitWithoutStream: true,            // Allow pings even when there are no active streams
	// }

	// grpcServer := grpc.NewServer()
	// listenAddr := conf.StoreAddr[strings.IndexByte(conf.StoreAddr, ':'):]
	// l, err := net.Listen("tcp", "localhost:8080")
	// if err != nil {
	// 	fmt.Println("Error")
	// }
	// gracefull shutdown
	// handleSignal(grpcServer)

	// err = grpcServer.Serve(l)
	// if err != nil {
	// 	fmt.Println("Error")
	// }
	// fmt.Println("Error")
}

// func handleSignal(grpcServer *grpc.Server) {
// 	sigCh := make(chan os.Signal, 1)
// 	signal.Notify(sigCh,
// 		syscall.SIGHUP,
// 		syscall.SIGINT,
// 		syscall.SIGTERM,
// 		syscall.SIGQUIT)
// 	go func() {
// 		sig := <-sigCh
// 		log.Infof("Got signal [%s] to exit.", sig)
// 		grpcServer.Stop()
// 	}()
// }
