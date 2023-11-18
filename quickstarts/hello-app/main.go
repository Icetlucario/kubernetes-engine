/**
 * Copyright 2021 Google Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

// [START gke_hello_app]
// [START container_hello_app]
package main

import (
	"html/template"
	"net/http"
)

// Product representa un producto en el catálogo.
type Product struct {
	Name  string
	Price float64
}

// Catalog representa el catálogo de productos.
var Catalog = []Product{
	{"Producto 1", 19.99},
	{"Producto 2", 29.99},
	{"Producto 3", 39.99},
}

// IndexHandler maneja las solicitudes a la página de inicio.
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	// Renderizar la página de inicio con la lista de productos.
	tmpl, err := template.New("index").Parse(`
		<!DOCTYPE html>
		<html>
		<head>
			<title>Web E-Commerce</title>
		</head>
		<body>
			<h1>Bienvenido a nuestra tienda en línea</h1>
			<h2>Productos disponibles:</h2>
			<ul>
				{{range .}}
					<li>{{.Name}} - ${{.Price}}</li>
				{{end}}
			</ul>
		</body>
		</html>
	`)
	if err != nil {
		http.Error(w, "Error al procesar la plantilla", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, Catalog)
	if err != nil {
		http.Error(w, "Error al renderizar la página", http.StatusInternalServerError)
	}
}

func main() {
	// Manejadores de ruta
	http.HandleFunc("/", IndexHandler)

	// Configuración del servidor y escucha
	port := ":8080"
	server := &http.Server{
		Addr:    port,
		Handler: nil,
	}

	// Iniciar el servidor
	println("Servidor web en ejecución en http://localhost" + port)
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

// [END container_hello_app]
// [END gke_hello_app]
