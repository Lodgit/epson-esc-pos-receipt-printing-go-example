package cmd

import (
	"bufio"
	"log"
	"net"
	"os"
	"strings"

	"github.com/seer-robotics/escpos"
)

// Execute starts application execution
func Execute() {
	// Usually "IP:9100"
	address := strings.TrimSpace(os.Getenv("PRINTER_ADDRESS"))

	if address == "" {
		log.Fatal("Please provides printer address (host and port)")
	}

	socket, err := net.Dial("tcp", address)

	if err != nil {
		println(err.Error())
	}

	defer socket.Close()

	w := bufio.NewWriter(socket)

	// Configure ESC/P control codes
	p := escpos.New(w)

	p.Verbose = true

	p.Init()
	p.SetFontSize(1, 1)
	p.SetAlign("center")
	p.Write("The quick brown fox jumps over the lazy dog")
	p.Formfeed()

	// Append a string to be printed using Western European encoding
	// which covers French, Spanish, Italian, Portuguese and German char-sets.
	// For more details take a look at https://github.com/seer-robotics/escpos/pull/1
	// a. Spanish sentence
	p.WriteWEU("Jovencillo emponzoñado de whisky: ¡qué figurota exhibe!")
	p.Formfeed()
	// b. German sentence
	p.WriteWEU("Zwölf Boxkämpfer jagten Victor quer über den großen Sylter Deich")
	p.Formfeed()

	// p.SetFontSize(2, 3)
	// p.SetFont("A")
	// p.SetFont("B")

	// p.SetEmphasize(1)
	// p.Write("hello")
	// p.Formfeed()

	// p.SetUnderline(1)
	// p.SetFontSize(4, 4)
	// p.Write("hello")

	// p.SetReverse(1)
	// p.SetFontSize(2, 4)
	// p.Write("hello")
	// p.FormfeedN(10)

	// p.SetAlign("center")
	// p.Write("test")
	// p.Linefeed()
	// p.Write("test")
	// p.Linefeed()
	// p.Write("test")
	// p.FormfeedD(200)

	p.Cut()

	w.Flush()

	log.Println("Printing data was delivered successfully.")
}
