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

### Análisis de los registros de RISC-V

#### x0, **`zero`**
```
Siempre contiene el valor 0. No se puede modificar.
En x86-64 no hay equivalente directo. Se usa mov $0, %reg.
Nota: En x86-64, si necesitas 0, lo cargas desde un inmediato.
Posible traducción: x0 se traduce a 0 literal.
```
#### x1, **`ra`**, Return address
```
Guarda la dirección de retorno al hacer una llamada a función (jal).
En x86-64 call y ret usan implícitamente rsp, no rax. Se podría almacenar la dirección en %rax.
TODO: x1
Posible traducción: explícitamente??, usar %rax o guardar en stack.
```
#### x2, **`sp`**, Stack pointer
```
Puntero a la pila. Se decrementa para hacer espacio, se incrementa para liberar.
En x86-64 %rsp es el equivalente exacto.
Posible traducción: x2 a %rsp
```
#### x3, **`gp`**, Global pointer
```
Acceso a datos globales en tiempo de ejecución.
En x86-64 no hay equivalente directo.
TODO: x3
Posible traducción: puede omitirse en una primera iteración??
```
#### x4, **`tp`**, Thread pointer
```
Usado en hilos Thread-Local Storage (TLS). TODO: ver concepto
x86-64: Equivalente conceptual: %fs o %gs (segmentos de hilo).
Traducción? lo mismo que x3
```
#### x5 a x7, **`t0`,`t1`,`t2`**, Temporales
```
Registros volátiles para uso temporal.
En x86-64 se usan %r10, %r11, %r12
Traducción: directa
    t0 (x5) a %r10
    t1 (x6) a %r11
    t2 (x7) a %r12
```
#### x8, **`s0` o `fp`**, Save / Frame pointer
> Es un registro que apunta al inicio del "marco de pila" (stack frame) de una función en ejecución. Marca una posición fija en la pila desde donde se pueden acceder a parámetros, variables locales y valores guardados.
> El **stack pointer** cambia constantemente mientras la función ejecuta, pero el **frame pointer** se mantiene constante durante toda la ejecución de una función.
```
Puede actuar como frame pointer o registro salvado.
En x86-64 %rbp se usa tradicionalmente como frame pointer.
Traducción directa de x8 a %rbp
```
#### x9, **`s1`**, Saved register
```
Debe preservarse entre llamadas a función.
En x86-64 el registro %rbx cumple el mismo rol.
Traducción directa de x9 a %rbx
```
#### x10 a x17, **`a0` a `a7`**, Argumentos o Retornos
> Es decir, con estos registros se pasan datos a una función par trabajar con ellos y se reciben los resultados de las funciones una vez finalizadas. Los pongo en una tabla porque hay algunos detalles que revisar

| DIreccion | Nombre | Trad. x86 | Uso |
|--|--|--|--|
|x10|a0|%rdi|1er arg.


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
