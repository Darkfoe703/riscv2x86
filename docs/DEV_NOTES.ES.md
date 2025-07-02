# Notas de Desarrollo
> ###### Dev Notes

## Consideraciones sobre las arquitecturas
Las arquitecturas CISC (Computadora con Conjunto de Instrucciones Complejas) y RISC (Computadora con Conjunto de Instrucciones Reducidas) difieren principalmente en el número y complejidad de sus instrucciones. CISC utiliza un conjunto de instrucciones amplio y complejo, mientras que RISC se enfoca en un conjunto más pequeño y simple, lo que afecta la forma en que se diseñan los procesadores y se escriben los programas. 
|Base de comparación| Arquitectura RISC | Arquitectura CISC |
|--|--|--|
| Conjunto de instrucciones | Reducido | Extenso |
|Uso de registros|Requiere más|Requiere menos|
|Pipelining|Fácil|Difícil|
|Direccionamiento|Se requiere un modo de direccionamiento limitado|Se requieren más modos de direccionamiento|

> [Bibliografía](https://www.stromasys.com/resources/decoding-risc-vs-cisc-architecture/)
---

## 1. Los registros en ambas arquitecturas
> En este punto me doy cuenta que, antes de planear qué operaciones traducir y cómo hacerlo, es primordial encontrar como trabajar con los registros de estas arquitecturas.
> ¿Qué registros tengo disponibles? ¿Es posible mapearlos 1 a 1 completamente? quizá de forma parcial? de forma dinámica?

### Registros en **RISC-V**
RISC-V tiene **32 registros enteros de propósito general**, de 64 bits:

| Nombre | Alias | Convención de uso |
|--------|-------|-------------------|
| `x0` | zero | Siempre 0 |
| `x1` | ra | Return address |
| `x2` | sp | Stack pointer |
| `x3` | gp | Global pointer |
| `x4` | tp | Thread pointer |
| `x5-x7` | t0-t2 | Temporales |
| `x8` | s0/fp | Saved register / frame pointer |
| `x9` | s1 | Saved register |
| `x10-x17` | a0-a7 | Argumentos / retorno |
| `x18-x27` | s2-s11 | Saved registers |
| `x28-x31` | t3-t6 | Temporales |

### Registros **x86-64**
x86-64 tiene registros generales de 64 bits como:

| Registro | Convención de uso |
|----------|------------|
| `rax` | Retorno / acumulador |
| `rbx` | Base |
| `rcx` | Contador |
| `rdx` | Datos / argumento |
| `rsi` | Fuente de string |
| `rdi` | Destino de string |
| `rsp` | Stack pointer |
| `rbp` | Base pointer |
| `r8–r15` | Argumentos / uso general |

También hay 16 registros de propósito general en total.

---

## 2. Instrucciones mínimas
Conjunto mínimo de instrucciones reales (por el momento no se tendrán en cuenta pseudo-instrucciones) RISC-V para traducir

Primera selección:
- Operaciones aritméticas básicas.
- Movimiento de datos.
- Control de flujo.
- Entrada/salida mínima (por syscall).

### 1. Operaciones aritméticas básicas:
Suma `add rd, rs1, rs2`

Resta `sub rd, rs1, rs2`

TODO:
#### NOTA: Los registros en ambas arquitecturas
En este punto me doy cuenta que, antes de planear qué operaciones traducir y cómo hacerlo, es primordial encontrar como trabajar con los registros de estas arquitecturas.
¿Qué registros tengo disponibles? ¿Es posible mapearlos 1 a 1 completamente? quizá de forma parcial? de forma dinámica?

RISC-V tiene **32 registros enteros de propósito general**, de 64 bits:

| Nombre | Alias | Convención de uso |
|--------|-------|-------------------|
| x0 | zero | Siempre 0 |
| x1 | ra | Return address |
| x2 | sp | Stack pointer |
| x3 | gp | Global pointer |
| x4 | tp | Thread pointer |
| x5-x7 | t0-t2 | Temporales |
| x8 | s0/fp | Saved register / frame pointer |
| x9 | s1 | Saved register |
| x10-x17 | a0-a7 | Argumentos / retorno |
| x18-x27 | s2-s11 | Saved registers |
| x28-x31 | t3-t6 | Temporales |