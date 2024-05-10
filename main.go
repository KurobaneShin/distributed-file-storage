package main

import (
	"bytes"
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
		EncKey:            newEncryptionKey(),
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
	s3 := makeServer(":5000", ":3000", ":4000")

	go func() {
		log.Fatal(s1.Start())
		time.Sleep(1 * time.Second)
	}()

	go s2.Start()
	time.Sleep(1 * time.Second)

	go s3.Start()

	time.Sleep(1 * time.Second)

	key := "key"
	data := bytes.NewReader([]byte("my big data file here!"))
	s3.Store(key, data)

	if err := s3.store.Delete(key); err != nil {
		log.Fatal(err)
	}
	r, err := s3.Get(key)
	if err != nil {
		log.Fatal(err)
	}

	b, err := io.ReadAll(r)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(b))
}
