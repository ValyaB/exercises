"".RecursiveChannel STEXT size=158 args=0x18 locals=0x38
	0x0000 00000 (./test.go:5)	TEXT	"".RecursiveChannel(SB), $56-24
	0x0000 00000 (./test.go:5)	MOVQ	(TLS), CX
	0x0009 00009 (./test.go:5)	CMPQ	SP, 16(CX)
	0x000d 00013 (./test.go:5)	JLS	148
	0x0013 00019 (./test.go:5)	SUBQ	$56, SP
	0x0017 00023 (./test.go:5)	MOVQ	BP, 48(SP)
	0x001c 00028 (./test.go:5)	LEAQ	48(SP), BP
	0x0021 00033 (./test.go:5)	FUNCDATA	$0, gclocals·a35fa7d7e7129fc330c152d6236a3e5c(SB)
	0x0021 00033 (./test.go:5)	FUNCDATA	$1, gclocals·69c1753bd5f81501d95132d08af04464(SB)
	0x0021 00033 (./test.go:5)	FUNCDATA	$3, gclocals·9fb7f0986f647f17cb53dda1484e0f7a(SB)
	0x0021 00033 (./test.go:7)	PCDATA	$2, $0
	0x0021 00033 (./test.go:7)	PCDATA	$0, $0
	0x0021 00033 (./test.go:7)	MOVQ	"".number+64(SP), AX
	0x0026 00038 (./test.go:7)	MOVQ	"".product+72(SP), CX
	0x002b 00043 (./test.go:7)	ADDQ	AX, CX
	0x002e 00046 (./test.go:9)	CMPQ	AX, $1
	0x0032 00050 (./test.go:9)	JNE	91
	0x0034 00052 (./test.go:11)	MOVQ	CX, ""..autotmp_3+40(SP)
	0x0039 00057 (./test.go:11)	PCDATA	$2, $1
	0x0039 00057 (./test.go:11)	PCDATA	$0, $1
	0x0039 00057 (./test.go:11)	MOVQ	"".result+80(SP), AX
	0x003e 00062 (./test.go:11)	PCDATA	$2, $0
	0x003e 00062 (./test.go:11)	MOVQ	AX, (SP)
	0x0042 00066 (./test.go:11)	PCDATA	$2, $1
	0x0042 00066 (./test.go:11)	LEAQ	""..autotmp_3+40(SP), AX
	0x0047 00071 (./test.go:11)	PCDATA	$2, $0
	0x0047 00071 (./test.go:11)	MOVQ	AX, 8(SP)
	0x004c 00076 (./test.go:11)	CALL	runtime.chansend1(SB)
	0x0051 00081 (./test.go:12)	MOVQ	48(SP), BP
	0x0056 00086 (./test.go:12)	ADDQ	$56, SP
	0x005a 00090 (./test.go:12)	RET
	0x005b 00091 (./test.go:15)	PCDATA	$0, $0
	0x005b 00091 (./test.go:15)	DECQ	AX
	0x005e 00094 (./test.go:15)	MOVQ	AX, 16(SP)
	0x0063 00099 (./test.go:15)	MOVQ	CX, 24(SP)
	0x0068 00104 (./test.go:15)	PCDATA	$2, $1
	0x0068 00104 (./test.go:15)	PCDATA	$0, $1
	0x0068 00104 (./test.go:15)	MOVQ	"".result+80(SP), AX
	0x006d 00109 (./test.go:15)	PCDATA	$2, $0
	0x006d 00109 (./test.go:15)	MOVQ	AX, 32(SP)
	0x0072 00114 (./test.go:15)	MOVL	$24, (SP)
	0x0079 00121 (./test.go:15)	PCDATA	$2, $1
	0x0079 00121 (./test.go:15)	LEAQ	"".RecursiveChannel·f(SB), AX
	0x0080 00128 (./test.go:15)	PCDATA	$2, $0
	0x0080 00128 (./test.go:15)	MOVQ	AX, 8(SP)
	0x0085 00133 (./test.go:15)	CALL	runtime.newproc(SB)
	0x008a 00138 (./test.go:16)	MOVQ	48(SP), BP
	0x008f 00143 (./test.go:16)	ADDQ	$56, SP
	0x0093 00147 (./test.go:16)	RET
	0x0094 00148 (./test.go:16)	NOP
	0x0094 00148 (./test.go:5)	PCDATA	$0, $-1
	0x0094 00148 (./test.go:5)	PCDATA	$2, $-1
	0x0094 00148 (./test.go:5)	CALL	runtime.morestack_noctxt(SB)
	0x0099 00153 (./test.go:5)	JMP	0
	0x0000 64 48 8b 0c 25 00 00 00 00 48 3b 61 10 0f 86 81  dH..%....H;a....
	0x0010 00 00 00 48 83 ec 38 48 89 6c 24 30 48 8d 6c 24  ...H..8H.l$0H.l$
	0x0020 30 48 8b 44 24 40 48 8b 4c 24 48 48 01 c1 48 83  0H.D$@H.L$HH..H.
	0x0030 f8 01 75 27 48 89 4c 24 28 48 8b 44 24 50 48 89  ..u'H.L$(H.D$PH.
	0x0040 04 24 48 8d 44 24 28 48 89 44 24 08 e8 00 00 00  .$H.D$(H.D$.....
	0x0050 00 48 8b 6c 24 30 48 83 c4 38 c3 48 ff c8 48 89  .H.l$0H..8.H..H.
	0x0060 44 24 10 48 89 4c 24 18 48 8b 44 24 50 48 89 44  D$.H.L$.H.D$PH.D
	0x0070 24 20 c7 04 24 18 00 00 00 48 8d 05 00 00 00 00  $ ..$....H......
	0x0080 48 89 44 24 08 e8 00 00 00 00 48 8b 6c 24 30 48  H.D$......H.l$0H
	0x0090 83 c4 38 c3 e8 00 00 00 00 e9 62 ff ff ff        ..8.......b...
	rel 5+4 t=16 TLS+0
	rel 77+4 t=8 runtime.chansend1+0
	rel 124+4 t=15 "".RecursiveChannel·f+0
	rel 134+4 t=8 runtime.newproc+0
	rel 149+4 t=8 runtime.morestack_noctxt+0
"".main STEXT size=256 args=0x0 locals=0x68
	0x0000 00000 (./test.go:18)	TEXT	"".main(SB), $104-0
	0x0000 00000 (./test.go:18)	MOVQ	(TLS), CX
	0x0009 00009 (./test.go:18)	CMPQ	SP, 16(CX)
	0x000d 00013 (./test.go:18)	JLS	246
	0x0013 00019 (./test.go:18)	SUBQ	$104, SP
	0x0017 00023 (./test.go:18)	MOVQ	BP, 96(SP)
	0x001c 00028 (./test.go:18)	LEAQ	96(SP), BP
	0x0021 00033 (./test.go:18)	FUNCDATA	$0, gclocals·7d2d5fca80364273fb07d5820a76fef4(SB)
	0x0021 00033 (./test.go:18)	FUNCDATA	$1, gclocals·f921db378992a411b04a03b038505262(SB)
	0x0021 00033 (./test.go:18)	FUNCDATA	$3, gclocals·1cf923758aae2e428391d1783fe59973(SB)
	0x0021 00033 (./test.go:20)	PCDATA	$2, $1
	0x0021 00033 (./test.go:20)	PCDATA	$0, $0
	0x0021 00033 (./test.go:20)	LEAQ	type.chan int(SB), AX
	0x0028 00040 (./test.go:20)	PCDATA	$2, $0
	0x0028 00040 (./test.go:20)	MOVQ	AX, (SP)
	0x002c 00044 (./test.go:20)	MOVQ	$0, 8(SP)
	0x0035 00053 (./test.go:20)	CALL	runtime.makechan(SB)
	0x003a 00058 (./test.go:20)	PCDATA	$2, $1
	0x003a 00058 (./test.go:20)	MOVQ	16(SP), AX
	0x003f 00063 (./test.go:20)	PCDATA	$2, $0
	0x003f 00063 (./test.go:20)	PCDATA	$0, $1
	0x003f 00063 (./test.go:20)	MOVQ	AX, "".result+72(SP)
	0x0044 00068 (./test.go:22)	MOVQ	$10000000000000, CX
	0x004e 00078 (./test.go:22)	MOVQ	CX, (SP)
	0x0052 00082 (./test.go:22)	MOVQ	$0, 8(SP)
	0x005b 00091 (./test.go:22)	CALL	"".RecursiveChannel(SB)
	0x0060 00096 (./test.go:23)	MOVQ	$0, ""..autotmp_2+64(SP)
	0x0069 00105 (./test.go:23)	PCDATA	$2, $1
	0x0069 00105 (./test.go:23)	PCDATA	$0, $0
	0x0069 00105 (./test.go:23)	MOVQ	"".result+72(SP), AX
	0x006e 00110 (./test.go:23)	PCDATA	$2, $0
	0x006e 00110 (./test.go:23)	MOVQ	AX, (SP)
	0x0072 00114 (./test.go:23)	PCDATA	$2, $1
	0x0072 00114 (./test.go:23)	LEAQ	""..autotmp_2+64(SP), AX
	0x0077 00119 (./test.go:23)	PCDATA	$2, $0
	0x0077 00119 (./test.go:23)	MOVQ	AX, 8(SP)
	0x007c 00124 (./test.go:23)	CALL	runtime.chanrecv1(SB)
	0x0081 00129 (./test.go:23)	MOVQ	""..autotmp_2+64(SP), AX
	0x0086 00134 (./test.go:25)	PCDATA	$0, $2
	0x0086 00134 (./test.go:25)	XORPS	X0, X0
	0x0089 00137 (./test.go:25)	MOVUPS	X0, ""..autotmp_3+80(SP)
	0x008e 00142 (./test.go:25)	PCDATA	$2, $2
	0x008e 00142 (./test.go:25)	LEAQ	type.int(SB), CX
	0x0095 00149 (./test.go:25)	PCDATA	$2, $0
	0x0095 00149 (./test.go:25)	MOVQ	CX, (SP)
	0x0099 00153 (./test.go:25)	MOVQ	AX, 8(SP)
	0x009e 00158 (./test.go:25)	CALL	runtime.convT2E64(SB)
	0x00a3 00163 (./test.go:25)	MOVQ	16(SP), AX
	0x00a8 00168 (./test.go:25)	PCDATA	$2, $2
	0x00a8 00168 (./test.go:25)	MOVQ	24(SP), CX
	0x00ad 00173 (./test.go:25)	MOVQ	AX, ""..autotmp_3+80(SP)
	0x00b2 00178 (./test.go:25)	PCDATA	$2, $0
	0x00b2 00178 (./test.go:25)	MOVQ	CX, ""..autotmp_3+88(SP)
	0x00b7 00183 (./test.go:25)	PCDATA	$2, $1
	0x00b7 00183 (./test.go:25)	LEAQ	go.string."Recursive: %d\n"(SB), AX
	0x00be 00190 (./test.go:25)	PCDATA	$2, $0
	0x00be 00190 (./test.go:25)	MOVQ	AX, (SP)
	0x00c2 00194 (./test.go:25)	MOVQ	$14, 8(SP)
	0x00cb 00203 (./test.go:25)	PCDATA	$2, $1
	0x00cb 00203 (./test.go:25)	LEAQ	""..autotmp_3+80(SP), AX
	0x00d0 00208 (./test.go:25)	PCDATA	$2, $0
	0x00d0 00208 (./test.go:25)	MOVQ	AX, 16(SP)
	0x00d5 00213 (./test.go:25)	MOVQ	$1, 24(SP)
	0x00de 00222 (./test.go:25)	MOVQ	$1, 32(SP)
	0x00e7 00231 (./test.go:25)	CALL	fmt.Printf(SB)
	0x00ec 00236 (./test.go:26)	PCDATA	$0, $0
	0x00ec 00236 (./test.go:26)	MOVQ	96(SP), BP
	0x00f1 00241 (./test.go:26)	ADDQ	$104, SP
	0x00f5 00245 (./test.go:26)	RET
	0x00f6 00246 (./test.go:26)	NOP
	0x00f6 00246 (./test.go:18)	PCDATA	$0, $-1
	0x00f6 00246 (./test.go:18)	PCDATA	$2, $-1
	0x00f6 00246 (./test.go:18)	CALL	runtime.morestack_noctxt(SB)
	0x00fb 00251 (./test.go:18)	JMP	0
	0x0000 64 48 8b 0c 25 00 00 00 00 48 3b 61 10 0f 86 e3  dH..%....H;a....
	0x0010 00 00 00 48 83 ec 68 48 89 6c 24 60 48 8d 6c 24  ...H..hH.l$`H.l$
	0x0020 60 48 8d 05 00 00 00 00 48 89 04 24 48 c7 44 24  `H......H..$H.D$
	0x0030 08 00 00 00 00 e8 00 00 00 00 48 8b 44 24 10 48  ..........H.D$.H
	0x0040 89 44 24 48 48 b9 00 a0 72 4e 18 09 00 00 48 89  .D$HH...rN....H.
	0x0050 0c 24 48 c7 44 24 08 00 00 00 00 e8 00 00 00 00  .$H.D$..........
	0x0060 48 c7 44 24 40 00 00 00 00 48 8b 44 24 48 48 89  H.D$@....H.D$HH.
	0x0070 04 24 48 8d 44 24 40 48 89 44 24 08 e8 00 00 00  .$H.D$@H.D$.....
	0x0080 00 48 8b 44 24 40 0f 57 c0 0f 11 44 24 50 48 8d  .H.D$@.W...D$PH.
	0x0090 0d 00 00 00 00 48 89 0c 24 48 89 44 24 08 e8 00  .....H..$H.D$...
	0x00a0 00 00 00 48 8b 44 24 10 48 8b 4c 24 18 48 89 44  ...H.D$.H.L$.H.D
	0x00b0 24 50 48 89 4c 24 58 48 8d 05 00 00 00 00 48 89  $PH.L$XH......H.
	0x00c0 04 24 48 c7 44 24 08 0e 00 00 00 48 8d 44 24 50  .$H.D$.....H.D$P
	0x00d0 48 89 44 24 10 48 c7 44 24 18 01 00 00 00 48 c7  H.D$.H.D$.....H.
	0x00e0 44 24 20 01 00 00 00 e8 00 00 00 00 48 8b 6c 24  D$ .........H.l$
	0x00f0 60 48 83 c4 68 c3 e8 00 00 00 00 e9 00 ff ff ff  `H..h...........
	rel 5+4 t=16 TLS+0
	rel 36+4 t=15 type.chan int+0
	rel 54+4 t=8 runtime.makechan+0
	rel 92+4 t=8 "".RecursiveChannel+0
	rel 125+4 t=8 runtime.chanrecv1+0
	rel 145+4 t=15 type.int+0
	rel 159+4 t=8 runtime.convT2E64+0
	rel 186+4 t=15 go.string."Recursive: %d\n"+0
	rel 232+4 t=8 fmt.Printf+0
	rel 247+4 t=8 runtime.morestack_noctxt+0
"".init STEXT size=92 args=0x0 locals=0x8
	0x0000 00000 (<autogenerated>:1)	TEXT	"".init(SB), $8-0
	0x0000 00000 (<autogenerated>:1)	MOVQ	(TLS), CX
	0x0009 00009 (<autogenerated>:1)	CMPQ	SP, 16(CX)
	0x000d 00013 (<autogenerated>:1)	JLS	85
	0x000f 00015 (<autogenerated>:1)	SUBQ	$8, SP
	0x0013 00019 (<autogenerated>:1)	MOVQ	BP, (SP)
	0x0017 00023 (<autogenerated>:1)	LEAQ	(SP), BP
	0x001b 00027 (<autogenerated>:1)	FUNCDATA	$0, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x001b 00027 (<autogenerated>:1)	FUNCDATA	$1, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x001b 00027 (<autogenerated>:1)	FUNCDATA	$3, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x001b 00027 (<autogenerated>:1)	PCDATA	$2, $0
	0x001b 00027 (<autogenerated>:1)	PCDATA	$0, $0
	0x001b 00027 (<autogenerated>:1)	MOVBLZX	"".initdone·(SB), AX
	0x0022 00034 (<autogenerated>:1)	CMPB	AL, $1
	0x0025 00037 (<autogenerated>:1)	JLS	48
	0x0027 00039 (<autogenerated>:1)	PCDATA	$2, $-2
	0x0027 00039 (<autogenerated>:1)	PCDATA	$0, $-2
	0x0027 00039 (<autogenerated>:1)	MOVQ	(SP), BP
	0x002b 00043 (<autogenerated>:1)	ADDQ	$8, SP
	0x002f 00047 (<autogenerated>:1)	RET
	0x0030 00048 (<autogenerated>:1)	JNE	57
	0x0032 00050 (<autogenerated>:1)	PCDATA	$2, $0
	0x0032 00050 (<autogenerated>:1)	PCDATA	$0, $0
	0x0032 00050 (<autogenerated>:1)	CALL	runtime.throwinit(SB)
	0x0037 00055 (<autogenerated>:1)	UNDEF
	0x0039 00057 (<autogenerated>:1)	MOVB	$1, "".initdone·(SB)
	0x0040 00064 (<autogenerated>:1)	CALL	fmt.init(SB)
	0x0045 00069 (<autogenerated>:1)	MOVB	$2, "".initdone·(SB)
	0x004c 00076 (<autogenerated>:1)	MOVQ	(SP), BP
	0x0050 00080 (<autogenerated>:1)	ADDQ	$8, SP
	0x0054 00084 (<autogenerated>:1)	RET
	0x0055 00085 (<autogenerated>:1)	NOP
	0x0055 00085 (<autogenerated>:1)	PCDATA	$0, $-1
	0x0055 00085 (<autogenerated>:1)	PCDATA	$2, $-1
	0x0055 00085 (<autogenerated>:1)	CALL	runtime.morestack_noctxt(SB)
	0x005a 00090 (<autogenerated>:1)	JMP	0
	0x0000 64 48 8b 0c 25 00 00 00 00 48 3b 61 10 76 46 48  dH..%....H;a.vFH
	0x0010 83 ec 08 48 89 2c 24 48 8d 2c 24 0f b6 05 00 00  ...H.,$H.,$.....
	0x0020 00 00 80 f8 01 76 09 48 8b 2c 24 48 83 c4 08 c3  .....v.H.,$H....
	0x0030 75 07 e8 00 00 00 00 0f 0b c6 05 00 00 00 00 01  u...............
	0x0040 e8 00 00 00 00 c6 05 00 00 00 00 02 48 8b 2c 24  ............H.,$
	0x0050 48 83 c4 08 c3 e8 00 00 00 00 eb a4              H...........
	rel 5+4 t=16 TLS+0
	rel 30+4 t=15 "".initdone·+0
	rel 51+4 t=8 runtime.throwinit+0
	rel 59+4 t=15 "".initdone·+-1
	rel 65+4 t=8 fmt.init+0
	rel 71+4 t=15 "".initdone·+-1
	rel 86+4 t=8 runtime.morestack_noctxt+0
go.loc."".RecursiveChannel SDWARFLOC size=154
	0x0000 ff ff ff ff ff ff ff ff 00 00 00 00 00 00 00 00  ................
	0x0010 26 00 00 00 00 00 00 00 9e 00 00 00 00 00 00 00  &...............
	0x0020 01 00 9c 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 ff ff ff ff ff ff ff ff 00 00 00 00 00  ................
	0x0040 00 00 00 2e 00 00 00 00 00 00 00 51 00 00 00 00  ...........Q....
	0x0050 00 00 00 01 00 52 00 00 00 00 00 00 00 00 00 00  .....R..........
	0x0060 00 00 00 00 00 00 ff ff ff ff ff ff ff ff 00 00  ................
	0x0070 00 00 00 00 00 00 26 00 00 00 00 00 00 00 9e 00  ......&.........
	0x0080 00 00 00 00 00 00 02 00 91 10 00 00 00 00 00 00  ................
	0x0090 00 00 00 00 00 00 00 00 00 00                    ..........
	rel 8+8 t=1 "".RecursiveChannel+0
	rel 59+8 t=1 "".RecursiveChannel+0
	rel 110+8 t=1 "".RecursiveChannel+0
go.info."".RecursiveChannel SDWARFINFO size=100
	0x0000 02 22 22 2e 52 65 63 75 72 73 69 76 65 43 68 61  ."".RecursiveCha
	0x0010 6e 6e 65 6c 00 00 00 00 00 00 00 00 00 00 00 00  nnel............
	0x0020 00 00 00 00 00 01 9c 00 00 00 00 01 0f 6e 75 6d  .............num
	0x0030 62 65 72 00 00 05 00 00 00 00 00 00 00 00 0f 70  ber............p
	0x0040 72 6f 64 75 63 74 00 00 05 00 00 00 00 00 00 00  roduct..........
	0x0050 00 0f 72 65 73 75 6c 74 00 00 05 00 00 00 00 00  ..result........
	0x0060 00 00 00 00                                      ....
	rel 21+8 t=1 "".RecursiveChannel+0
	rel 29+8 t=1 "".RecursiveChannel+158
	rel 39+4 t=29 gofile../home/vbukhalova/go/src/exercises/recursionChannels/test.go+0
	rel 54+4 t=28 go.info.int+0
	rel 58+4 t=28 go.loc."".RecursiveChannel+0
	rel 73+4 t=28 go.info.int+0
	rel 77+4 t=28 go.loc."".RecursiveChannel+51
	rel 91+4 t=28 go.info.chan int+0
	rel 95+4 t=28 go.loc."".RecursiveChannel+102
go.range."".RecursiveChannel SDWARFRANGE size=0
go.isstmt."".RecursiveChannel SDWARFMISC size=0
	0x0000 04 13 04 0e 03 05 01 08 02 04 01 02 02 05 01 18  ................
	0x0010 02 0d 01 2c 02 14 00                             ...,...
go.string."Recursive: %d\n" SRODATA dupok size=14
	0x0000 52 65 63 75 72 73 69 76 65 3a 20 25 64 0a        Recursive: %d.
go.loc."".main SDWARFLOC size=122
	0x0000 ff ff ff ff ff ff ff ff 00 00 00 00 00 00 00 00  ................
	0x0010 86 00 00 00 00 00 00 00 a3 00 00 00 00 00 00 00  ................
	0x0020 01 00 50 00 00 00 00 00 00 00 00 00 00 00 00 00  ..P.............
	0x0030 00 00 00 ff ff ff ff ff ff ff ff 00 00 00 00 00  ................
	0x0040 00 00 00 3f 00 00 00 00 00 00 00 60 00 00 00 00  ...?.......`....
	0x0050 00 00 00 01 00 50 60 00 00 00 00 00 00 00 00 01  .....P`.........
	0x0060 00 00 00 00 00 00 02 00 91 58 00 00 00 00 00 00  .........X......
	0x0070 00 00 00 00 00 00 00 00 00 00                    ..........
	rel 8+8 t=1 "".main+0
	rel 59+8 t=1 "".main+0
go.info."".main SDWARFINFO size=67
	0x0000 02 22 22 2e 6d 61 69 6e 00 00 00 00 00 00 00 00  ."".main........
	0x0010 00 00 00 00 00 00 00 00 00 01 9c 00 00 00 00 01  ................
	0x0020 0a 61 6e 73 77 65 72 00 17 00 00 00 00 00 00 00  .answer.........
	0x0030 00 0a 72 65 73 75 6c 74 00 14 00 00 00 00 00 00  ..result........
	0x0040 00 00 00                                         ...
	rel 9+8 t=1 "".main+0
	rel 17+8 t=1 "".main+256
	rel 27+4 t=29 gofile../home/vbukhalova/go/src/exercises/recursionChannels/test.go+0
	rel 41+4 t=28 go.info.int+0
	rel 45+4 t=28 go.loc."".main+0
	rel 58+4 t=28 go.info.chan int+0
	rel 62+4 t=28 go.loc."".main+51
go.range."".main SDWARFRANGE size=0
go.isstmt."".main SDWARFMISC size=0
	0x0000 04 13 04 0e 03 07 01 1c 02 0a 01 12 02 09 01 1d  ................
	0x0010 02 03 01 63 02 14 00                             ...c...
go.loc."".init SDWARFLOC size=0
go.info."".init SDWARFINFO size=33
	0x0000 02 22 22 2e 69 6e 69 74 00 00 00 00 00 00 00 00  ."".init........
	0x0010 00 00 00 00 00 00 00 00 00 01 9c 00 00 00 00 01  ................
	0x0020 00                                               .
	rel 9+8 t=1 "".init+0
	rel 17+8 t=1 "".init+92
	rel 27+4 t=29 gofile..<autogenerated>+0
go.range."".init SDWARFRANGE size=0
go.isstmt."".init SDWARFMISC size=0
	0x0000 04 0f 04 0c 03 07 01 05 02 09 01 07 02 09 01 15  ................
	0x0010 02 07 00                                         ...
"".initdone· SNOPTRBSS size=1
"".RecursiveChannel·f SRODATA dupok size=8
	0x0000 00 00 00 00 00 00 00 00                          ........
	rel 0+8 t=1 "".RecursiveChannel+0
runtime.gcbits.01 SRODATA dupok size=1
	0x0000 01                                               .
type..namedata.*interface {}- SRODATA dupok size=16
	0x0000 00 00 0d 2a 69 6e 74 65 72 66 61 63 65 20 7b 7d  ...*interface {}
type.*interface {} SRODATA dupok size=56
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 4f 0f 96 9d 00 08 08 36 00 00 00 00 00 00 00 00  O......6........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.algarray+80
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
	rel 24+8 t=1 runtime.algarray+144
	rel 32+8 t=1 runtime.gcbits.02+0
	rel 40+4 t=5 type..namedata.*interface {}-+0
	rel 44+4 t=6 type.*interface {}+0
	rel 56+8 t=1 type.interface {}+80
type..namedata.*[]interface {}- SRODATA dupok size=18
	0x0000 00 00 0f 2a 5b 5d 69 6e 74 65 72 66 61 63 65 20  ...*[]interface 
	0x0010 7b 7d                                            {}
type.*[]interface {} SRODATA dupok size=56
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 f3 04 9a e7 00 08 08 36 00 00 00 00 00 00 00 00  .......6........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.algarray+80
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*[]interface {}-+0
	rel 48+8 t=1 type.[]interface {}+0
type.[]interface {} SRODATA dupok size=56
	0x0000 18 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 70 93 ea 2f 02 08 08 17 00 00 00 00 00 00 00 00  p../............
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.algarray+0
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*[]interface {}-+0
	rel 44+4 t=6 type.*[]interface {}+0
	rel 48+8 t=1 type.interface {}+0
type..namedata.*[1]interface {}- SRODATA dupok size=19
	0x0000 00 00 10 2a 5b 31 5d 69 6e 74 65 72 66 61 63 65  ...*[1]interface
	0x0010 20 7b 7d                                          {}
type.*[1]interface {} SRODATA dupok size=56
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 bf 03 a8 35 00 08 08 36 00 00 00 00 00 00 00 00  ...5...6........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.algarray+80
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*[1]interface {}-+0
	rel 48+8 t=1 type.[1]interface {}+0
type.[1]interface {} SRODATA dupok size=72
	0x0000 10 00 00 00 00 00 00 00 10 00 00 00 00 00 00 00  ................
	0x0010 50 91 5b fa 02 08 08 11 00 00 00 00 00 00 00 00  P.[.............
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0040 01 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.algarray+144
	rel 32+8 t=1 runtime.gcbits.02+0
	rel 40+4 t=5 type..namedata.*[1]interface {}-+0
	rel 44+4 t=6 type.*[1]interface {}+0
	rel 48+8 t=1 type.interface {}+0
	rel 56+8 t=1 type.[]interface {}+0
type..namedata.*chan int- SRODATA dupok size=12
	0x0000 00 00 09 2a 63 68 61 6e 20 69 6e 74              ...*chan int
type.*chan int SRODATA dupok size=56
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 ed 7b ed 3b 00 08 08 36 00 00 00 00 00 00 00 00  .{.;...6........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.algarray+80
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*chan int-+0
	rel 48+8 t=1 type.chan int+0
type.chan int SRODATA dupok size=64
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 91 55 cb 71 02 08 08 32 00 00 00 00 00 00 00 00  .U.q...2........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00 03 00 00 00 00 00 00 00  ................
	rel 24+8 t=1 runtime.algarray+80
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*chan int-+0
	rel 44+4 t=6 type.*chan int+0
	rel 48+8 t=1 type.int+0
type..importpath.fmt. SRODATA dupok size=6
	0x0000 00 00 03 66 6d 74                                ...fmt
gclocals·a35fa7d7e7129fc330c152d6236a3e5c SRODATA dupok size=10
	0x0000 02 00 00 00 03 00 00 00 04 00                    ..........
gclocals·69c1753bd5f81501d95132d08af04464 SRODATA dupok size=8
	0x0000 02 00 00 00 00 00 00 00                          ........
gclocals·9fb7f0986f647f17cb53dda1484e0f7a SRODATA dupok size=10
	0x0000 02 00 00 00 01 00 00 00 00 01                    ..........
gclocals·7d2d5fca80364273fb07d5820a76fef4 SRODATA dupok size=8
	0x0000 03 00 00 00 00 00 00 00                          ........
gclocals·f921db378992a411b04a03b038505262 SRODATA dupok size=11
	0x0000 03 00 00 00 03 00 00 00 00 01 04                 ...........
gclocals·1cf923758aae2e428391d1783fe59973 SRODATA dupok size=11
	0x0000 03 00 00 00 02 00 00 00 00 01 02                 ...........
gclocals·33cdeccccebe80329f1fdbee7f5874cb SRODATA dupok size=8
	0x0000 01 00 00 00 00 00 00 00                          ........
