package translator

// RiscvToX86 define una traducción estática entre directivas RISC-V y x86-64 (sintaxis GAS)
var RiscvToX86 = map [string]string {
	".text": ".text",
	".data": ".data",
	".globl": ".globl",
	".section": ".section",
	".type": ".type",
	".size": ".size",
	".zero": ".zero",
	".align": ".align",
	".bss": ".bss",
	".byte": ".byte",	// 1byte
	// Cambios
	".half": ".word",	// 2 bytes
	".word": ".long",	// 4 bytes
	".dword": ".quad",	// 8 bytes
	".string": ".asciz",	// cadena terminada en null (\o)
	".asciz": ".asciz",

}

// Devuelve el nombre de la directiva x86-64 correspondiente a su hmologo RISC-V.
// Si no existe, devuelve string vacío.
func GetX86Directive(riscvDirective string) string {
    if x86, ok := RiscvToX86[riscvDirective]; ok {
        return x86
    }
    return ""
}