# Meli Challenge - Never Gonna Give You Up

Este repositorio contiene el código del **Meli Challenge Never Gonna Give You Up**. 

## Requisitos

Para ejecutar este proyecto, es necesario contar con:

- **Go** 1.24.1 

### Instalación
1. Clona el repositorio:
   ```bash
   git clone https://github.com/jisantillan/meli-challenge.git
   ```
2. Navega al directorio del proyecto:
   ```bash
   cd meli-challenge
   ```

   
# Validador de Melodías 

El **Validador de Melodías** toma como input una melodía representada en formato de texto plano a través de la **línea de comandos**, analizando si sigue una serie de reglas. Si el texto es válido, el programa confirmará que la melodía es correcta. En caso de que haya un error, el programa indicará la posición del primer error encontrado en la melodía.


## Ejecución

Para ejecutar el validador de melodías, usar el siguiente comando en la línea de comandos. El argumento `-melody` debe contener la melodía que se desea validar:
  ```bash
go run ./cmd/cli -melody="60 A{d=7/4;o=3;a=#} B{o=2;d=1/2} S A{d=2;a=n} G{a=b} B S{d=1/3}"
 ```
Ejemplo de salida:

Si la melodía es válida:

  ```bash
valid melody
 ```
 Si la melodía tiene un error:

  ```bash
error at position n
 ```

