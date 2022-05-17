package main

import (
	"log"

	"github.com/weisrc/nsink"
)

func main() {
	println("running")
	server := nsink.NewServer(nsink.Options{
		Ttl:   1,
		Proxy: "1.1.1.1:53",
		Open:  true,
	})

	server.SetClient("::1", []string{"base"})
	server.SetTreeBlock("base", "example.com.")

	tree := server.Trees["base"]
	println(tree.String("*", 0))

	log.Fatalln(server.Listen(":53", "udp"))
}
