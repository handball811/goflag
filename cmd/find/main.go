// +build ignore

package main

import (
	. "github.com/mmcloughlin/avo/build"
	. "github.com/mmcloughlin/avo/operand"
)

func main() {

	TEXT(
		"Find",
		NOSPLIT,
		"func(t uint64) byte",
	)
	Doc("Find the lowest down flag")
	t64 := Load(Param("t"), GP64())
	tmp := GP64()
	XORQ(tmp, tmp)
	ans := GP8()
	XORB(ans, ans)

	Comment("Check lower 16bit")
	MOVQ(t64, tmp)
	ANDQ(U32(0xffff), tmp)
	CMPQ(tmp, U32(0xffff))
	JNE(LabelRef("point"))
	ADDB(Imm(16), ans)
	SHRQ(Imm(16), t64)
	Label("point")

	Comment("Check lower 16bit")
	MOVQ(t64, tmp)
	ANDQ(U32(0xffff), tmp)
	CMPQ(tmp, U32(0xffff))
	JNE(LabelRef("point0"))
	ADDB(Imm(16), ans)
	SHRQ(Imm(16), t64)
	Label("point0")

	Comment("Check lower 16bit")
	MOVQ(t64, tmp)
	ANDQ(U32(0xffff), tmp)
	CMPQ(tmp, U32(0xffff))
	JNE(LabelRef("point1"))
	ADDB(Imm(16), ans)
	SHRQ(Imm(16), t64)
	Label("point1")

	Comment("Check lower 8bit")
	MOVQ(t64, tmp)
	ANDQ(U32(0xff), tmp)
	CMPQ(tmp, U32(0xff))
	JNE(LabelRef("point2"))
	ADDB(Imm(8), ans)
	SHRQ(Imm(8), t64)
	Label("point2")

	Comment("Check lower 4bit")
	MOVQ(t64, tmp)
	ANDQ(U32(0xf), tmp)
	CMPQ(tmp, U32(0xf))
	JNE(LabelRef("point3"))
	ADDB(Imm(4), ans)
	SHRQ(Imm(4), t64)
	Label("point3")

	Comment("Check lower 2bit")
	MOVQ(t64, tmp)
	ANDQ(U32(0x3), tmp)
	CMPQ(tmp, U32(0x3))
	JNE(LabelRef("point4"))
	ADDB(Imm(2), ans)
	SHRQ(Imm(2), t64)
	Label("point4")

	Comment("Check lower 1bit")
	MOVQ(t64, tmp)
	ANDQ(U32(0x1), tmp)
	CMPQ(tmp, U32(0x1))
	JNE(LabelRef("point5"))
	ADDB(Imm(1), ans)
	Label("point5")

	Store(ans, ReturnIndex(0))

	RET()
	Generate()
}
