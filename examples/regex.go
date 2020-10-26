package main

import (
    "fmt"
    "regexp"
)

func main() {
   expr := "(?:((?i)kernel:.*sd.*|(?i)kernel: EXT.*))((?i)error)|( ata(\\d).*(: exception Emask))"
   fmt.Println(expr)
   reg := regexp.MustCompile(expr)
   str := "[Wed Sep  4 20:09:07 2019] audit: kernel:asd.asderror exception Emask type=1327 audit(1567653048.832:2956463): proctitle=69707461626C65732D726573746F7265002D770035002D2D6E6F666C757368002D2D636F756E74657273\n" +
"    [Wed Sep  4 20:09:36 2019] ata3.00: exception Emask 0x0 SAct 0x0 SErr 0x0 action 0x6 frozen\n" +
"    [Wed Sep  4 20:09:36 2019] ata3.00: failed command: FLUSH CACHE EXT\n" +
"    [Wed Sep  4 20:09:36 2019] ata3.00: cmd ea/00:00:00:00:00/00:00:00:00:00/a0 tag 18\n" +
"    res 40/00:00:00:4f:c2/00:00:00:00:00/00 Emask 0x4 (timeout)\n" +
"    [Wed Sep  4 20:09:36 2019] ata3.00: status: { DRDY }\n"+
"    [Wed Sep  4 20:09:36 2019] ata3: hard resetting link\n" +
"    [Wed Sep  4 20:09:37 2019] ata3: SATA link up 1.5 Gbps (SStatus 113 SControl 310)\n" +
"    [Wed Sep  4 20:09:37 2019] ata3.00: configured for UDMA/33\n" +
"    [Wed Sep  4 20:09:37 2019] ata3.00: retrying FLUSH 0xea Emask 0x4\n" +
"    [Wed Sep  4 20:09:37 2019] ata3: EH complete\n"
   loc := reg.FindStringIndex(str)
   fmt.Println(loc)
   fmt.Println(str[loc[0]:loc[1]])
}
