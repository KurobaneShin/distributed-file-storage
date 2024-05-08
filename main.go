package main

import (
	"fmt"
	"io"
	"log"
	"time"

	"github.com/KurobaneShin/distributed-file-storage/p2p"
)

func makeServer(listenAddr string, nodes ...string) *FileServer {
	tcptransformOpts := p2p.TCPTransportOpts{
		ListenAddr:    listenAddr,
		HandshakeFunc: p2p.NOPHandshakeFunc,
		Decoder:       p2p.DefaultDecoder{},
	}

	tcpTransport := p2p.NewTCPTransport(tcptransformOpts)

	FileServerOpts := FileServerOpts{
		StorageRoot:       listenAddr + "_network",
		PathTransformFunc: CASPathTransformFunc,
		Transport:         tcpTransport,
		BoostrapNodes:     nodes,
	}
	s := NewFileServer(FileServerOpts)

	tcpTransport.OnPeer = s.OnPeer

	return s
}

func main() {
	s1 := makeServer(":3000", "")
	s2 := makeServer(":4000", ":3000")
	go func() {
		log.Fatal(s1.Start())
	}()
	// time.Sleep(1 * time.Second)

	go s2.Start()

	time.Sleep(1 * time.Second)

	// for i := 0; i < 10; i++ {
	// data := bytes.NewReader([]byte("my big data file here!"))
	// s2.Store(fmt.Sprintf("key"), data)
	// time.Sleep(5 * time.Millisecond)
	// }

	r, err := s2.Get("key")
	if err != nil {
		log.Fatal(err)
	}

	b, err := io.ReadAll(r)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(b))
}
