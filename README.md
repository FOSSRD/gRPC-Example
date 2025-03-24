# Ejemplo de gRPC y Protocol Buffers

Este repositorio demuestra la facilidad de construir aplicaciones utilizando gRPC y Protocol Buffers como métodos de comunicación, resaltando su eficiencia y capacidad de interoperabilidad entre diferentes lenguajes de programación.

## ¿Qué es gRPC?

gRPC es un framework de código abierto para llamadas a procedimientos remotos (RPC) de alto rendimiento y de propósito general. Desarrollado originalmente por Google, gRPC permite que una aplicación cliente invoque directamente métodos en una aplicación de servidor en otra máquina, como si fuera un objeto local. gRPC utiliza HTTP/2 como protocolo de transporte (Aunque, en verdad, esta es una definicion un tanto vieja, gRPC puede ser usado sobre HTTP/3 o QUIC desde 2024 para mas velocidad), lo que proporciona características como multiplexación y transmisión bidireccional.

## ¿Qué son los Protocol Buffers?

Los Protocol Buffers son un formato de serialización de datos agnóstico al lenguaje y a la plataforma. Permiten definir estructuras de datos una vez y luego generar código fuente en varios lenguajes para leer y escribir esos datos hacia y desde una variedad de flujos de datos o formatos. Su diseño se centra en la eficiencia y la simplicidad.

## Ventajas de usar gRPC y Protocol Buffers

* *Alto Rendimiento:* Protocol Buffers ofrece una serialización más rápida y eficiente en comparación con formatos basados en texto como JSON o XML, lo que se traduce en menor tamaño de los mensajes y menor latencia.
* *Eficiencia de Red:* gRPC, al utilizar HTTP/2, optimiza el uso de la red a través de la multiplexación de conexiones y la compresión de encabezados.
* *Independencia de Lenguaje:* Tanto gRPC como Protocol Buffers son compatibles con una amplia gama de lenguajes de programación, lo que facilita la creación de sistemas distribuidos donde los servicios pueden estar implementados en diferentes lenguajes.
* *Generación de Código:* La definición de los servicios y los mensajes mediante archivos .proto permite la generación automática de código de cliente y servidor en los lenguajes soportados. Esto reduce significativamente la cantidad de código repetitivo necesario.
* *Contratos Claros:* El uso de Protocol Buffers define un contrato estricto entre el cliente y el servidor, lo que ayuda a prevenir errores de comunicación y facilita la comprensión de la API.

## Demostración de Comunicación Entre Lenguajes

Este repositorio contiene un ejemplo práctico de un servicio de conversión de imágenes construido utilizando gRPC y Protocol Buffers.

* El directorio imager alberga un servidor implementado en *Go*. Este servidor ofrece funcionalidades para realizar transformaciones comunes en imágenes, como la conversión a blanco y negro, la aplicación de un filtro sepia y el desenfoque gaussiano.
* El directorio servit contiene un cliente implementado en *Python*. Este cliente se comunica con el servidor Go a través de gRPC para solicitar el procesamiento de una imagen y luego muestra los resultados utilizando la librería Streamlit.

Este ejemplo ilustra cómo un cliente desarrollado en Python puede interactuar sin problemas con un servidor desarrollado en Go mediante gRPC y Protocol Buffers, lo que subraya la capacidad de comunicación entre diferentes tecnologías y lenguajes.

## Cómo Construir y Ejecutar el Ejemplo

### Prerrequisitos

1.  Asegúrese de tener instaladas las herramientas de desarrollo necesarias para *Go* y *Python* en su sistema. Esto incluye los compiladores, los gestores de paquetes (como go mod para Go y pip para Python) y las bibliotecas gRPC y Protocol Buffers para ambos lenguajes. Puede encontrar guías detalladas de instalación en los siguientes enlaces:
    * [gRPC Quickstart en Python](https://grpc.io/docs/languages/python/quickstart/)
    * [gRPC Quickstart en Go](https://grpc.io/docs/languages/go/quickstart/)
2. Asegurate de tener instalados los tools mencionados en el QuickStart de go en especial, ya que para python casi todos se instalan como parte del script de build.

### Construcción

Ejecute el script _build.sh_ ubicado en la raíz del proyecto. Este script automatiza los siguientes pasos:

1.  *Construcción del cliente Python:* Genera el código Python necesario a partir del archivo de definición de Protocol Buffers (converter.proto).
2.  *Construcción del servidor Go:* Gestiona las dependencias de Go y genera el código Go a partir del mismo archivo .proto.

### Ejecución

1.  Abra una terminal, navegue al directorio imager y ejecute el servidor Go:
    ```bash
    go run .
    ```

2.  Abra una segunda terminal, navegue al directorio servit y ejecute el cliente Python utilizando Streamlit:
    ```bash
    streamlit run servit.py
    ```

Una vez que la aplicación Streamlit se inicie en su navegador web, mostrará la imagen original junto con las versiones procesadas (blanco y negro, borrosa y sepia) generadas por el servidor Go a través de la comunicación gRPC.

Este ejemplo práctico demuestra la eficiencia y la facilidad con la que se pueden construir aplicaciones distribuidas y que interoperan entre diferentes lenguajes de programación utilizando gRPC y Protocol Buffers.
