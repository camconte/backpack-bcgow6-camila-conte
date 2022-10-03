/*Una empresa de redes sociales requiere implementar una estructura usuario con funciones que vayan agregando información a la estructura. Para optimizar y ahorrar memoria requieren que la estructura de usuarios ocupe el mismo lugar en memoria para el main del programa y para las funciones.
La estructura debe tener los campos: Nombre, Apellido, Edad, Correo y Contraseña
Y deben implementarse las funciones:
- Cambiar nombre: me permite cambiar el nombre y apellido.
- Cambiar edad: me permite cambiar la edad.
- Cambiar correo: me permite cambiar el correo.
- Cambiar contraseña: me permite cambiar la contraseña.
*/

package main

import "fmt"

type user struct {
	Name string
	Lastname string
	Age int
	Mail string
	Password string
}

func (u *user) changeAllName(newName string, newLastname string) {
	u.Name = newName
	u.Lastname = newLastname
}
func (u *user) changeAge(newAge int) {
	u.Age = newAge
}
func (u *user) changeMail(newMail string) {
	u.Mail = newMail
}
func (u *user) changePassword(newPass string) {
	u.Password = newPass
}

func main()  {
	user := user{}

	fmt.Printf("El usuario es: %+v", user)
	
	user.changeAllName("Camila", "Conte")
	
	fmt.Printf("El usuario es: %+v", user)

}