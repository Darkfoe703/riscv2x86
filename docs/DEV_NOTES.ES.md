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

## 1. Instrucciones mínimas
Conjunto mínimo de instrucciones reales (por el momento no se tendrán en cuenta pseudo-instrucciones) RISC-V para traducir

Primera selección:
- Operaciones aritméticas básicas.
- Movimiento de datos.
- Control de flujo.
- Entrada/salida mínima (por syscall).

### 1. Operaciones aritméticas básicas:
Suma `add rd, rs1, rs2`

Resta `sub rd, rs1, rs2`

