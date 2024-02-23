package services

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/Setsu548/trival-assassement/models"
)

type Users []models.User

func GetUsers() (*[]Users, error) {
	// obtenemos los datos desde el end-point de la api: https://randomuser.me/api
	responseAPI, err := http.Get("https://randomuser.me/api")
	if err != nil {
		log.Fatal(err)
	}

	// agarramos todo el response y parseamos el body
	body, err := io.ReadAll(responseAPI.Body)
	if err != nil {
		log.Fatal(err)
	}

	// cerramos la conexion esperando a que termine todo
	defer responseAPI.Body.Close()

	// empezamos a parsear la estructura del del body en un interface creando una variable
	var data map[string]interface{}
	err = json.Unmarshal(body, &data) // parseamos el body en nuestra variable "data"
	if err != nil {
		log.Fatal(err) // si hay algun error, salta el parse
	}

	// creamos un slice para ir agregando de nuestro data a nuestro User
	users := make(Users, 0)

	// como el json results:[] tiene sub-estructuras, avanzamos de nivel
	// asignando a nuestro result todo lo que hay en dentro de data-results, en results
	results := data["results"].([]interface{})

	// empezamos a parsear a la estructura que queremos
	for _, result := range results {
		user := models.User{} // creamos un objeto de nuestro modelo

		// asignamos el genero a nuestro usuario parseando a travez de "gender"
		user.Genero = result.(map[string]interface{})["gender"].(string)

		// como los demas datos estan en una sub-estructura paseamos otro nivel
		name := result.(map[string]interface{})["name"].(map[string]interface{})
		user.Nombre = fmt.Sprintf(name["first"].(string))  // asignamos el nombre de name.first
		user.Apellido = fmt.Sprintf(name["last"].(string)) // asignamos el apellido de name.last

		user.Correo = result.(map[string]interface{})["email"].(string) // asignamos el email

		// asignamos nuestro uuid
		user.UUID = result.(map[string]interface{})["login"].(map[string]interface{})["uuid"].(string)

		// todos los datos que se han asignado, vamos agregando a nuestro slice
		users = append(users, user)
	}

	// quitamos los duplicados
	removedDuplicated := removeDuplicates(&users)

	return &removedDuplicated, nil
}

func removeDuplicates(users *Users) []Users {
	seen := make(map[string]bool)
	result := make([]Users, 0, len(*users))

	for _, user := range *users {
		key := user.UUID
		if _, ok := seen[key]; !ok {
			seen[key] = true
			result = append(result, *users)
		}

	}
	return result
}
