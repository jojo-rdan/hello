package greet

// las variables declaradas afuera de una función
// son variables de paquete
// y se pueden utilizar en varias partes a nivel de paquete
const emoji = "happy-face"

// retorna un saludo en inglés
func English() string {
	return "Hi " + emoji
}

// retorna un saludo en italiano
func Italian() string {
	return "Ciao " + emoji
}

// Para identificar qué se puede exportar y que no
// sólo debemos declarar con la primera letra en mayúscula
// lo que que sea utilizado desde otros paquetes externos

// todo se puede compartir entre un mismo paquete
