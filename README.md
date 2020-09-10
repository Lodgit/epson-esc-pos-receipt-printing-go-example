# Lodgit / Epson ESC-POS receipt printing example

> [Go](https://golang.org/) receipt printing example using [EPSON ESC/POS](https://reference.epson-biz.com/modules/ref_escpos/index.php?content_id=2) command system.

<img width="600" alt="Epson ESC-POS receipt printing example" src="https://user-images.githubusercontent.com/1700322/74636645-b5fe1e00-5168-11ea-994c-ee3404536355.png">

*Screenshot of a receipt printed by the [EPSON TM-T20II POS Receipt Printer](https://epson.com/For-Work/Printers/POS/TM-T20II-POS-Receipt-Printer/p/C31CD52062)*.

### Usage

```sh
env PRINTER_ADDRESS=127.0.0.1:9100 \
    make run
```

***NOTE:** Change the corresponding printer address.*

### Resources

- [seer-robotics/escpos — PR #1 Western European encoding support](https://github.com/seer-robotics/escpos/pull/1)
- [EPSON ESC/POS - POS Printer Command System Reference](https://reference.epson-biz.com/modules/ref_escpos/index.php?content_id=2)
- [ESC/P - EPSON Standard Code for Printers](https://en.wikipedia.org/wiki/ESC/P)
