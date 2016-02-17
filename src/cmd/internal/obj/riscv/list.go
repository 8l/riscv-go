//	Copyright © 1994-1999 Lucent Technologies Inc.  All rights reserved.
//	Portions Copyright © 1995-1997 C H Forsyth (forsyth@terzarima.net)
//	Portions Copyright © 1997-1999 Vita Nuova Limited
//	Portions Copyright © 2000-2008 Vita Nuova Holdings Limited (www.vitanuova.com)
//	Portions Copyright © 2004,2006 Bruce Ellis
//	Portions Copyright © 2005-2007 C H Forsyth (forsyth@terzarima.net)
//	Revisions Copyright © 2000-2008 Lucent Technologies Inc. and others
//	Portions Copyright © 2009 The Go Authors.  All rights reserved.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.  IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package riscv

import (
	"fmt"

	"cmd/internal/obj"
)

var (
	// Instructions is a map of instruction names to integer IDs.
	Instructions = make(map[string]int)

	// Registers is a map of register names to integer IDs.
	Registers = make(map[string]int16)

	// regNames is a map for register IDs to names. We use
	// the ABI name.
	regNames = map[int16]string{
		0: "NONE",

		// General registers with ABI names.
		REG_ZERO: "ZERO",
		REG_RA:   "RA",
		REG_FP:   "FP",
		// REG_S0 is REG_FP.
		// REG_S1 is REG_SB.
		// REG_S2 is REG_RT1.
		// REG_S3 is REG_RT2.
		// REG_S4 is REG_CTXT.
		// REG_S5 is REG_G.
		REG_S6:  "S6",
		REG_S7:  "S7",
		REG_S8:  "S8",
		REG_S9:  "S9",
		REG_S10: "S10",
		REG_S11: "S11",
		REG_SP:  "SP",
		REG_TP:  "TP",
		REG_V0:  "V0",
		REG_V1:  "V1",
		REG_A0:  "A0",
		REG_A1:  "A1",
		REG_A2:  "A2",
		REG_A3:  "A3",
		REG_A4:  "A4",
		REG_A5:  "A5",
		REG_A6:  "A6",
		REG_A7:  "A7",
		REG_T0:  "T0",
		REG_T1:  "T1",
		REG_T2:  "T2",
		REG_T3:  "T3",
		REG_T4:  "T4",
		REG_GP:  "GP",

		// Golang runtime register names.
		REG_SB:   "SB",
		REG_RT1:  "RT1",
		REG_RT2:  "RT2",
		REG_CTXT: "CTXT",
		REG_G:    "G",

		// ABI names for floating point registers.
		REG_FS0:  "FS0",
		REG_FS1:  "FS1",
		REG_FS2:  "FS2",
		REG_FS3:  "FS3",
		REG_FS4:  "FS4",
		REG_FS5:  "FS5",
		REG_FS6:  "FS6",
		REG_FS7:  "FS7",
		REG_FS8:  "FS8",
		REG_FS9:  "FS9",
		REG_FS10: "FS10",
		REG_FS11: "FS11",
		REG_FS12: "FS12",
		REG_FS13: "FS13",
		REG_FS14: "FS14",
		REG_FS15: "FS15",
		REG_FV0:  "FV0",
		REG_FV1:  "FV1",
		REG_FA0:  "FA0",
		REG_FA1:  "FA1",
		REG_FA2:  "FA2",
		REG_FA3:  "FA3",
		REG_FA4:  "FA4",
		REG_FA5:  "FA5",
		REG_FA6:  "FA6",
		REG_FA7:  "FA7",
		REG_FT0:  "FT0",
		REG_FT1:  "FT1",
		REG_FT2:  "FT2",
		REG_FT3:  "FT3",
		REG_FT4:  "FT4",
		REG_FT5:  "FT5",
	}
)

func initRegisters() {
	// Standard register names.
	for i := REG_X0; i <= REG_X31; i++ {
		name := fmt.Sprintf("X%d", i-REG_X0)
		Registers[name] = int16(i)
	}
	for i := REG_F0; i <= REG_F31; i++ {
		name := fmt.Sprintf("F%d", i-REG_F0)
		Registers[name] = int16(i)
	}

	// General registers with ABI names.
	Registers["ZERO"] = REG_ZERO
	Registers["RA"] = REG_RA
	Registers["FP"] = REG_FP
	Registers["S0"] = REG_S0
	Registers["S1"] = REG_S1
	Registers["S2"] = REG_S2
	Registers["S3"] = REG_S3
	Registers["S4"] = REG_S4
	Registers["S5"] = REG_S5
	Registers["S6"] = REG_S6
	Registers["S7"] = REG_S7
	Registers["S8"] = REG_S8
	Registers["S9"] = REG_S9
	Registers["S10"] = REG_S10
	Registers["S11"] = REG_S11
	Registers["SP"] = REG_SP
	Registers["TP"] = REG_TP
	Registers["V0"] = REG_V0
	Registers["V1"] = REG_V1
	Registers["A0"] = REG_A0
	Registers["A1"] = REG_A1
	Registers["A2"] = REG_A2
	Registers["A3"] = REG_A3
	Registers["A4"] = REG_A4
	Registers["A5"] = REG_A5
	Registers["A6"] = REG_A6
	Registers["A7"] = REG_A7
	Registers["T0"] = REG_T0
	Registers["T1"] = REG_T1
	Registers["T2"] = REG_T2
	Registers["T3"] = REG_T3
	Registers["T4"] = REG_T4
	Registers["GP"] = REG_GP

	// Golang runtime register names.
	Registers["SB"] = REG_SB
	Registers["RT1"] = REG_RT1
	Registers["RT2"] = REG_RT2
	Registers["CTXT"] = REG_CTXT
	Registers["G"] = REG_G

	// ABI names for floating point registers.
	Registers["FS0"] = REG_FS0
	Registers["FS1"] = REG_FS1
	Registers["FS2"] = REG_FS2
	Registers["FS3"] = REG_FS3
	Registers["FS4"] = REG_FS4
	Registers["FS5"] = REG_FS5
	Registers["FS6"] = REG_FS6
	Registers["FS7"] = REG_FS7
	Registers["FS8"] = REG_FS8
	Registers["FS9"] = REG_FS9
	Registers["FS10"] = REG_FS10
	Registers["FS11"] = REG_FS11
	Registers["FS12"] = REG_FS12
	Registers["FS13"] = REG_FS13
	Registers["FS14"] = REG_FS14
	Registers["FS15"] = REG_FS15
	Registers["FV0"] = REG_FV0
	Registers["FV1"] = REG_FV1
	Registers["FA0"] = REG_FA0
	Registers["FA1"] = REG_FA1
	Registers["FA2"] = REG_FA2
	Registers["FA3"] = REG_FA3
	Registers["FA4"] = REG_FA4
	Registers["FA5"] = REG_FA5
	Registers["FA6"] = REG_FA6
	Registers["FA7"] = REG_FA7
	Registers["FT0"] = REG_FT0
	Registers["FT1"] = REG_FT1
	Registers["FT2"] = REG_FT2
	Registers["FT3"] = REG_FT3
	Registers["FT4"] = REG_FT4
	Registers["FT5"] = REG_FT5
}

// checkRegNames asserts that regNames includes all registers.
func checkRegNames() {
	for i := REG_X0; i <= REG_X31; i++ {
		if _, ok := regNames[int16(i)]; !ok {
			panic(fmt.Sprintf("REG_X%d missing from regNames", i))
		}
	}
	for i := REG_F0; i <= REG_F31; i++ {
		if _, ok := regNames[int16(i)]; !ok {
			panic(fmt.Sprintf("REG_F%d missing from regNames", i))
		}
	}
}

func initInstructions() {
	for i, s := range obj.Anames {
		Instructions[s] = i
	}
	for i, s := range Anames {
		if i >= obj.A_ARCHSPECIFIC {
			Instructions[s] = i + obj.ABaseRISCV
		}
	}
}

func init() {
	// initRegnames uses Registers during initialization,
	// and must be called after initRegisters.
	initRegisters()
	checkRegNames()
	initInstructions()
	obj.RegisterRegister(obj.RBaseRISCV, REG_END, PrettyPrintReg)
	obj.RegisterOpcode(obj.ABaseRISCV, Anames)
}

func PrettyPrintReg(r int) string {
	name, ok := regNames[int16(r)]
	if !ok {
		name = fmt.Sprintf("R???%d", r) // Similar format to Aconv.
	}

	return name
}