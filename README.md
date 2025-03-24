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
3. Instala las dependencias:
   ```bash
   go mod tidy
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


# API
## Levantar el servidor 
Para iniciar el servidor de la API, ejecutar el siguiente comando:

```bash
go run ./cmd/api
 ```

Al ejecutar el comando anterior, el servidor escuchará en el puerto 8080 y estará disponible para recibir solicitudes HTTP.
## Referencia
### Validar melodía
**`POST` `/melody/validate`**
### Descripción
Este endpoint acepta un body JSON con el campo **melody** que contiene la melodía a validar, si es válida retornará una estructura que represente dicha melodía, en caso contrario devolvera una respuesta con la primera posición donde se encuentre un error.

### Request
#### Headers

| Parámetro | Valor |
| :-------: | :---------- |
| `Content-Type` | `application/json` |

####  Request Body

| Campo | Tipo           | Descripción |
| :-------: | :-----------:  | :---------- |
| `melody`     | `string` | Melodia musical a ser validada |

#### Ejemplo
``` json
{
    "melody": "60 A{d=7/4;o=3;a=#} B{o=2;d=1/4} S G{d=2}"
}
```
### Response

#### Códigos HTTP

| Codigo | Descripción |
| :----: | :---------- |
| **200** | OK |
| **400** | Bad Request |

#### Headers

| Parámetro | Valor |
| :-------: | :---------- |
| `Content-Type` | `application/json` |

#### Ejemplo - Response Exitosa

``` json
{
    "tempo": {
        "value": 60,
        "unit": "bpm"
    },
    "notes": [
        {
            "type": "note",
            "name": "la",
            "octave": 3,
            "alteration": "#",
            "duration": 1.75,
            "frequency": 233.08
        },
        {
            "type": "note",
            "name": "si",
            "octave": 2,
            "alteration": "n",
            "duration": 0.25,
            "frequency": 123.47
        },
        {
            "type": "silence",
            "duration": 1
        },
        {
            "type": "note",
            "name": "sol",
            "octave": 4,
            "alteration": "n",
            "duration": 2,
            "frequency": 392.00
        }
    ]
}

 ```

 #### Ejemplo - Response con Error

``` json
 {
    "cause": "error at position 18"
 }
 ```

---

 ### Reproducir melodía
**`POST` `/melody/play`**
### Descripción
Este endpoint acepta un body JSON con una estructura que representa una melodía y la reproduce en el servidor.

### Request
#### Headers

| Parámetro | Valor |
| :-------: | :---------- |
| `Content-Type` | `application/json` |

####  Request Body

| Campo       | Tipo         | Descripción                          |
|------------|------------|----------------------------------|
| `tempo`    | `object`   | Contiene la información del tempo de la melodía. |
| `tempo.value` | `number` | Valor numérico del tempo. |
| `tempo.unit`  | `string` | Unidad del tempo. |
| `notes`    | `array`    | Lista de notas musicales en la melodía. |
| `notes[].type` | `string` | Tipo de elemento (`"note"` o `"silence"`). |
| `notes[].name` | `string` | Nombre de la nota musical (ej. `"do"`, `"re"`). |
| `notes[].octave` | `number` | Octava en la que se encuentra la nota. |
| `notes[].alteration` | `string` | Alteración de la nota (`"none"`, `"#"`, `"b"`). |
| `notes[].duration` | `number` | Duración de la nota. |
| `notes[].frequency` | `number` | Frecuencia de la nota. |


#### Ejemplo
``` json
{
  "tempo": {
    "value": 60,
    "unit": "bpm"
  },
  "notes": [
    {
      "type": "note",
      "name": "la",
      "octave": 3,
      "alteration": "#",
      "duration": 1.75,
      "frequency": 233.08
    },
    {
      "type": "note",
      "name": "si",
      "octave": 2,
      "alteration": "none",
      "duration": 0.25,
      "frequency": 123.94
    },
    {
      "type": "silence",
      "duration": 1
    },
    {
      "type": "note",
      "name": "sol",
      "octave": 4,
      "alteration": "none",
      "duration": 2,
      "frequency": 392.00
    }
  ]
}
```
### Response

#### Códigos HTTP

| Codigo | Descripción |
| :----: | :---------- |
| **202** | Accepted |

#### Body
No se espera cuerpo de respuesta.

## Postman Collection
Puedes descargar la colección de Postman [aquí](https://gist.github.com/jisantillan/01050b3b7b481ebef0ec8759d9ffe580) para probar los endpoints de la API directamente en Postman.




