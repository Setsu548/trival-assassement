package models

type User struct {
	Genero   string `json:"genero"`
	Nombre   string `json:"nombre"`
	Apellido string `json:"apellido"`
	Correo   string `json:"correo"`
	UUID     string `json:"uuid"`
}
