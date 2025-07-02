package translator

// RiscvToX86 define una traducción estática entre registros RISC-V y x86-64 (sintaxis GAS)
var RiscvToX86 = map[string]string{
    // Registros especiales
    "zero": "$0",       // constante cero
    "ra":   "%rax",     // return address
    "sp":   "%rsp",     // stack pointer
    "gp":   "",         // no mapeado
    "tp":   "",         // no mapeado (TLS)

    // Temporales
    "t0": "%r10",
    "t1": "%r11",
    "t2": "%r12",
    "t3": "%r13",
    "t4": "%r14",
    "t5": "%r15",
    "t6": "", // TODO: mapear a stack si es necesario

    // Saved registers
    "s0": "%rbp", // también se usa como frame pointer
    "s1": "%rbx",
    "s2": "%r13",
    "s3": "%r14",
    "s4": "%r15",
    "s5": "", // disponibles si no colisionan
    "s6": "",
    "s7": "",
    "s8": "",
    "s9": "",
    "s10": "",
    "s11": "",

    // Argumentos y retorno
    "a0": "%rdi", // 1er argumento / retorno
    "a1": "%rsi",
    "a2": "%rdx",
    "a3": "%rcx",
    "a4": "%r8",
    "a5": "%r9",
    "a6": "", // no disponible en ABI estándar
    "a7": "%rax", // syscall number
}

// GetX86Register devuelve el nombre del registro x86-64 correspondiente a un alias RISC-V.
// Si no existe, devuelve string vacío.
func GetX86Register(riscv string) string {
    if reg, ok := RiscvToX86[riscv]; ok {
        return reg
    }
    return ""
}
