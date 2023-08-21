package main

import "fmt"

func main() {
	fmt.Print("hello")
}

/*
输出编译过程
$ go build -n

#
# github.com/miiy/go-examples/go-build/build3
#

mkdir -p $WORK/b001/
cat >$WORK/b001/importcfg << 'EOF' # internal
# import config
packagefile fmt=/usr/local/go/pkg/darwin_amd64/fmt.a
packagefile runtime=/usr/local/go/pkg/darwin_amd64/runtime.a
EOF
cd /go-examples/go_build/build3
/usr/local/go/pkg/tool/darwin_amd64/compile -o $WORK/b001/_pkg_.a -trimpath "$WORK/b001=>" -p main -lang=go1.18 -complete -buildid HQStidYd9gSaHV3HPzBW/HQStidYd9gSaHV3HPzBW -goversion go1.18.3 -c=4 -nolocalimports -importcfg $WORK/b001/importcfg -pack ./main.go
/usr/local/go/pkg/tool/darwin_amd64/buildid -w $WORK/b001/_pkg_.a # internal
cat >$WORK/b001/importcfg.link << 'EOF' # internal
packagefile github.com/miiy/go-examples/go-build/build3=$WORK/b001/_pkg_.a
packagefile fmt=/usr/local/go/pkg/darwin_amd64/fmt.a
packagefile runtime=/usr/local/go/pkg/darwin_amd64/runtime.a
packagefile errors=/usr/local/go/pkg/darwin_amd64/errors.a
packagefile internal/fmtsort=/usr/local/go/pkg/darwin_amd64/internal/fmtsort.a
packagefile io=/usr/local/go/pkg/darwin_amd64/io.a
packagefile math=/usr/local/go/pkg/darwin_amd64/math.a
packagefile os=/usr/local/go/pkg/darwin_amd64/os.a
packagefile reflect=/usr/local/go/pkg/darwin_amd64/reflect.a
packagefile strconv=/usr/local/go/pkg/darwin_amd64/strconv.a
packagefile sync=/usr/local/go/pkg/darwin_amd64/sync.a
packagefile unicode/utf8=/usr/local/go/pkg/darwin_amd64/unicode/utf8.a
packagefile internal/abi=/usr/local/go/pkg/darwin_amd64/internal/abi.a
packagefile internal/bytealg=/usr/local/go/pkg/darwin_amd64/internal/bytealg.a
packagefile internal/cpu=/usr/local/go/pkg/darwin_amd64/internal/cpu.a
packagefile internal/goarch=/usr/local/go/pkg/darwin_amd64/internal/goarch.a
packagefile internal/goexperiment=/usr/local/go/pkg/darwin_amd64/internal/goexperiment.a
packagefile internal/goos=/usr/local/go/pkg/darwin_amd64/internal/goos.a
packagefile runtime/internal/atomic=/usr/local/go/pkg/darwin_amd64/runtime/internal/atomic.a
packagefile runtime/internal/math=/usr/local/go/pkg/darwin_amd64/runtime/internal/math.a
packagefile runtime/internal/sys=/usr/local/go/pkg/darwin_amd64/runtime/internal/sys.a
packagefile internal/reflectlite=/usr/local/go/pkg/darwin_amd64/internal/reflectlite.a
packagefile sort=/usr/local/go/pkg/darwin_amd64/sort.a
packagefile math/bits=/usr/local/go/pkg/darwin_amd64/math/bits.a
packagefile internal/itoa=/usr/local/go/pkg/darwin_amd64/internal/itoa.a
packagefile internal/oserror=/usr/local/go/pkg/darwin_amd64/internal/oserror.a
packagefile internal/poll=/usr/local/go/pkg/darwin_amd64/internal/poll.a
packagefile internal/syscall/execenv=/usr/local/go/pkg/darwin_amd64/internal/syscall/execenv.a
packagefile internal/syscall/unix=/usr/local/go/pkg/darwin_amd64/internal/syscall/unix.a
packagefile internal/testlog=/usr/local/go/pkg/darwin_amd64/internal/testlog.a
packagefile internal/unsafeheader=/usr/local/go/pkg/darwin_amd64/internal/unsafeheader.a
packagefile io/fs=/usr/local/go/pkg/darwin_amd64/io/fs.a
packagefile sync/atomic=/usr/local/go/pkg/darwin_amd64/sync/atomic.a
packagefile syscall=/usr/local/go/pkg/darwin_amd64/syscall.a
packagefile time=/usr/local/go/pkg/darwin_amd64/time.a
packagefile unicode=/usr/local/go/pkg/darwin_amd64/unicode.a
packagefile internal/race=/usr/local/go/pkg/darwin_amd64/internal/race.a
packagefile path=/usr/local/go/pkg/darwin_amd64/path.a
modinfo "0w\xaf\f\x92t\b\x02A\xe1\xc1\a\xe6\xd6\x18\xe6path\tgithub.com/miiy/go-examples/go-build/build3\nmod\tgithub.com/miiy/go-examples/go-build/build3\t(devel)\t\nbuild\t-compiler=gc\nbuild\tCGO_ENABLED=1\nbuild\tCGO_CFLAGS=\nbuild\tCGO_CPPFLAGS=\nbuild\tCGO_CXXFLAGS=\nbuild\tCGO_LDFLAGS=\nbuild\tGOARCH=amd64\nbuild\tGOOS=darwin\nbuild\tGOAMD64=v1\nbuild\tvcs=git\nbuild\tvcs.revision=fa305af521bd179130074698019177d7de3ccb73\nbuild\tvcs.time=2022-06-17T09:08:28Z\nbuild\tvcs.modified=true\n\xf92C1\x86\x18 r\x00\x82B\x10A\x16\xd8\xf2"
EOF
mkdir -p $WORK/b001/exe/
cd .
/usr/local/go/pkg/tool/darwin_amd64/link -o $WORK/b001/exe/a.out -importcfg $WORK/b001/importcfg.link -buildmode=exe -buildid=wpZOntJk7ADpNg98ufhL/HQStidYd9gSaHV3HPzBW/HQStidYd9gSaHV3HPzBW/wpZOntJk7ADpNg98ufhL -extld=clang $WORK/b001/_pkg_.a
/usr/local/go/pkg/tool/darwin_amd64/buildid -w $WORK/b001/exe/a.out # internal
mv $WORK/b001/exe/a.out build3
*/

/*
中间码生成 SSA
查看从代码到SSA中间码的整个过程
$ GOSSAFUNC=main go build

# runtime
dumped SSA to /go-examples/go_build/build3/ssa.html
# github.com/miiy/go-examples/go-build/build3
dumped SSA to ./ssa.html
*/

/*
机器码生成
查看 Pllan9 汇编代码
$ go build -gcflags -S

# github.com/miiy/go-examples/go-build/build3
"".main STEXT size=103 args=0x0 locals=0x40 funcid=0x0 align=0x0
0x0000 00000 (/go-examples/go_build/build3/main.go:5)    TEXT    "".main(SB), ABIInternal, $64-0
0x0000 00000 (/go-examples/go_build/build3/main.go:5)    CMPQ    SP, 16(R14)
0x0004 00004 (/go-examples/go_build/build3/main.go:5)    PCDATA  $0, $-2
0x0004 00004 (/go-examples/go_build/build3/main.go:5)    JLS     92
0x0006 00006 (/go-examples/go_build/build3/main.go:5)    PCDATA  $0, $-1
0x0006 00006 (/go-examples/go_build/build3/main.go:5)    SUBQ    $64, SP
0x000a 00010 (/go-examples/go_build/build3/main.go:5)    MOVQ    BP, 56(SP)
0x000f 00015 (/go-examples/go_build/build3/main.go:5)    LEAQ    56(SP), BP
0x0014 00020 (/go-examples/go_build/build3/main.go:5)    FUNCDATA        $0, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
0x0014 00020 (/go-examples/go_build/build3/main.go:5)    FUNCDATA        $1, gclocals·f207267fbf96a0178e8758c6e3e0ce28(SB)
0x0014 00020 (/go-examples/go_build/build3/main.go:5)    FUNCDATA        $2, "".main.stkobj(SB)
0x0014 00020 (/go-examples/go_build/build3/main.go:6)    MOVUPS  X15, ""..autotmp_8+40(SP)
0x001a 00026 (/go-examples/go_build/build3/main.go:6)    LEAQ    type.string(SB), DX
0x0021 00033 (/go-examples/go_build/build3/main.go:6)    MOVQ    DX, ""..autotmp_8+40(SP)
0x0026 00038 (/go-examples/go_build/build3/main.go:6)    LEAQ    ""..stmp_0(SB), DX
0x002d 00045 (/go-examples/go_build/build3/main.go:6)    MOVQ    DX, ""..autotmp_8+48(SP)
0x0032 00050 (<unknown line number>)    NOP
0x0032 00050 ($GOROOT/src/fmt/print.go:242)     MOVQ    os.Stdout(SB), BX
0x0039 00057 ($GOROOT/src/fmt/print.go:242)     LEAQ    go.itab.*os.File,io.Writer(SB), AX
0x0040 00064 ($GOROOT/src/fmt/print.go:242)     LEAQ    ""..autotmp_8+40(SP), CX
0x0045 00069 ($GOROOT/src/fmt/print.go:242)     MOVL    $1, DI
0x004a 00074 ($GOROOT/src/fmt/print.go:242)     MOVQ    DI, SI
0x004d 00077 ($GOROOT/src/fmt/print.go:242)     PCDATA  $1, $0
0x004d 00077 ($GOROOT/src/fmt/print.go:242)     CALL    fmt.Fprint(SB)
0x0052 00082 (/go-examples/go_build/build3/main.go:7)    MOVQ    56(SP), BP
0x0057 00087 (/go-examples/go_build/build3/main.go:7)    ADDQ    $64, SP
0x005b 00091 (/go-examples/go_build/build3/main.go:7)    RET
0x005c 00092 (/go-examples/go_build/build3/main.go:7)    NOP
0x005c 00092 (/go-examples/go_build/build3/main.go:5)    PCDATA  $1, $-1
0x005c 00092 (/go-examples/go_build/build3/main.go:5)    PCDATA  $0, $-2
0x005c 00092 (/go-examples/go_build/build3/main.go:5)    NOP
0x0060 00096 (/go-examples/go_build/build3/main.go:5)    CALL    runtime.morestack_noctxt(SB)
0x0065 00101 (/go-examples/go_build/build3/main.go:5)    PCDATA  $0, $-1
0x0065 00101 (/go-examples/go_build/build3/main.go:5)    JMP     0
0x0000 49 3b 66 10 76 56 48 83 ec 40 48 89 6c 24 38 48  I;f.vVH..@H.l$8H
0x0010 8d 6c 24 38 44 0f 11 7c 24 28 48 8d 15 00 00 00  .l$8D..|$(H.....
0x0020 00 48 89 54 24 28 48 8d 15 00 00 00 00 48 89 54  .H.T$(H......H.T
0x0030 24 30 48 8b 1d 00 00 00 00 48 8d 05 00 00 00 00  $0H......H......
0x0040 48 8d 4c 24 28 bf 01 00 00 00 48 89 fe e8 00 00  H.L$(.....H.....
0x0050 00 00 48 8b 6c 24 38 48 83 c4 40 c3 0f 1f 40 00  ..H.l$8H..@...@.
0x0060 e8 00 00 00 00 eb 99                             .......
rel 2+0 t=23 type.string+0
rel 2+0 t=23 type.*os.File+0
rel 29+4 t=14 type.string+0
rel 41+4 t=14 ""..stmp_0+0
rel 53+4 t=14 os.Stdout+0
rel 60+4 t=14 go.itab.*os.File,io.Writer+0
rel 78+4 t=7 fmt.Fprint+0
rel 97+4 t=7 runtime.morestack_noctxt+0
go.cuinfo.producer.main SDWARFCUINFO dupok size=0
0x0000 72 65 67 61 62 69                                regabi
go.cuinfo.packagename.main SDWARFCUINFO dupok size=0
0x0000 6d 61 69 6e                                      main
go.info.fmt.Print$abstract SDWARFABSFCN dupok size=40
0x0000 05 66 6d 74 2e 50 72 69 6e 74 00 01 01 13 61 00  .fmt.Print....a.
0x0010 00 00 00 00 00 13 6e 00 01 00 00 00 00 13 65 72  ......n.......er
0x0020 72 00 01 00 00 00 00 00                          r.......
rel 0+0 t=22 type.[]interface {}+0
rel 0+0 t=22 type.error+0
rel 0+0 t=22 type.int+0
rel 17+4 t=31 go.info.[]interface {}+0
rel 25+4 t=31 go.info.int+0
rel 35+4 t=31 go.info.error+0
""..inittask SNOPTRDATA size=32
0x0000 00 00 00 00 00 00 00 00 01 00 00 00 00 00 00 00  ................
0x0010 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
rel 24+8 t=1 fmt..inittask+0
go.string."hello" SRODATA dupok size=5
0x0000 68 65 6c 6c 6f                                   hello
go.itab.*os.File,io.Writer SRODATA dupok size=32
0x0000 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
0x0010 44 b5 f3 33 00 00 00 00 00 00 00 00 00 00 00 00  D..3............
rel 0+8 t=1 type.io.Writer+0
rel 8+8 t=1 type.*os.File+0
rel 24+8 t=-32767 os.(*File).Write+0
""..stmp_0 SRODATA static size=16
0x0000 00 00 00 00 00 00 00 00 05 00 00 00 00 00 00 00  ................
rel 0+8 t=1 go.string."hello"+0
runtime.nilinterequal·f SRODATA dupok size=8
0x0000 00 00 00 00 00 00 00 00                          ........
rel 0+8 t=1 runtime.nilinterequal+0
runtime.memequal64·f SRODATA dupok size=8
0x0000 00 00 00 00 00 00 00 00                          ........
rel 0+8 t=1 runtime.memequal64+0
runtime.gcbits.01 SRODATA dupok size=1
0x0000 01                                               .
type..namedata.*interface {}- SRODATA dupok size=15
0x0000 00 0d 2a 69 6e 74 65 72 66 61 63 65 20 7b 7d     ..*interface {}
type.*interface {} SRODATA dupok size=56
0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
0x0010 4f 0f 96 9d 08 08 08 36 00 00 00 00 00 00 00 00  O......6........
0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
0x0030 00 00 00 00 00 00 00 00                          ........
rel 24+8 t=1 runtime.memequal64·f+0
rel 32+8 t=1 runtime.gcbits.01+0
rel 40+4 t=5 type..namedata.*interface {}-+0
rel 48+8 t=1 type.interface {}+0
runtime.gcbits.02 SRODATA dupok size=1
0x0000 02                                               .
type.interface {} SRODATA dupok size=80
0x0000 10 00 00 00 00 00 00 00 10 00 00 00 00 00 00 00  ................
0x0010 e7 57 a0 18 02 08 08 14 00 00 00 00 00 00 00 00  .W..............
0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
0x0030 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
0x0040 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
rel 24+8 t=1 runtime.nilinterequal·f+0
rel 32+8 t=1 runtime.gcbits.02+0
rel 40+4 t=5 type..namedata.*interface {}-+0
rel 44+4 t=-32763 type.*interface {}+0
rel 56+8 t=1 type.interface {}+80
type..namedata.*[]interface {}- SRODATA dupok size=17
0x0000 00 0f 2a 5b 5d 69 6e 74 65 72 66 61 63 65 20 7b  ..*[]interface {
0x0010 7d                                               }
type.*[]interface {} SRODATA dupok size=56
0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
0x0010 f3 04 9a e7 08 08 08 36 00 00 00 00 00 00 00 00  .......6........
0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
0x0030 00 00 00 00 00 00 00 00                          ........
rel 24+8 t=1 runtime.memequal64·f+0
rel 32+8 t=1 runtime.gcbits.01+0
rel 40+4 t=5 type..namedata.*[]interface {}-+0
rel 48+8 t=1 type.[]interface {}+0
type.[]interface {} SRODATA dupok size=56
0x0000 18 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
0x0010 70 93 ea 2f 02 08 08 17 00 00 00 00 00 00 00 00  p../............
0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
0x0030 00 00 00 00 00 00 00 00                          ........
rel 32+8 t=1 runtime.gcbits.01+0
rel 40+4 t=5 type..namedata.*[]interface {}-+0
rel 44+4 t=-32763 type.*[]interface {}+0
rel 48+8 t=1 type.interface {}+0
type..namedata.*[1]interface {}- SRODATA dupok size=18
0x0000 00 10 2a 5b 31 5d 69 6e 74 65 72 66 61 63 65 20  ..*[1]interface
0x0010 7b 7d                                            {}
type.*[1]interface {} SRODATA dupok size=56
0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
0x0010 bf 03 a8 35 08 08 08 36 00 00 00 00 00 00 00 00  ...5...6........
0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
0x0030 00 00 00 00 00 00 00 00                          ........
rel 24+8 t=1 runtime.memequal64·f+0
rel 32+8 t=1 runtime.gcbits.01+0
rel 40+4 t=5 type..namedata.*[1]interface {}-+0
rel 48+8 t=1 type.[1]interface {}+0
type.[1]interface {} SRODATA dupok size=72
0x0000 10 00 00 00 00 00 00 00 10 00 00 00 00 00 00 00  ................
0x0010 50 91 5b fa 02 08 08 11 00 00 00 00 00 00 00 00  P.[.............
0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
0x0030 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
0x0040 01 00 00 00 00 00 00 00                          ........
rel 24+8 t=1 runtime.nilinterequal·f+0
rel 32+8 t=1 runtime.gcbits.02+0
rel 40+4 t=5 type..namedata.*[1]interface {}-+0
rel 44+4 t=-32763 type.*[1]interface {}+0
rel 48+8 t=1 type.interface {}+0
rel 56+8 t=1 type.[]interface {}+0
type..importpath.fmt. SRODATA dupok size=5
0x0000 00 03 66 6d 74                                   ..fmt
gclocals·33cdeccccebe80329f1fdbee7f5874cb SRODATA dupok size=8
0x0000 01 00 00 00 00 00 00 00                          ........
gclocals·f207267fbf96a0178e8758c6e3e0ce28 SRODATA dupok size=9
0x0000 01 00 00 00 02 00 00 00 00                       .........
"".main.stkobj SRODATA static size=24
0x0000 01 00 00 00 00 00 00 00 f0 ff ff ff 10 00 00 00  ................
0x0010 10 00 00 00 00 00 00 00                          ........
rel 20+4 t=5 runtime.gcbits.02+0
*/
