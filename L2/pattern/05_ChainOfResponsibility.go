/* 
	Паттерн "Цепочка обязанностей" (Chain of Responsibility) — это поведенческий шаблон 
	проектирования, который позволяет передавать запросы последовательно по цепочке обработчиков. 
	Каждый обработчик решает, может ли он обработать запрос самостоятельно, или передать его следующему обработчику в цепочке.

	+:
		- Ослабление зависимости между отправителем и получателем запроса: Отправитель запроса не знает, какой обработчик его обработает.
		- Гибкость в добавлении новых обработчиков: Новые обработчики можно добавлять без изменения существующих классов.
		- Уменьшение количества подклассов: Паттерн позволяет избежать создания множества подклассов, обрабатывающих различные типы запросов.

	-: 
		-Запрос может остаться необработанным: Если ни один из обработчиков не сможет обработать запрос, он останется необработанным.
		- Отладка может быть сложной: При сложных цепочках обработчиков может быть трудно отследить, какой обработчик обрабатывает запрос.

	Примеры:
		- Для последовательной обработки запросов на аутентификацию и авторизацию. 
		- Паттерн часто используется для обработки HTTP-запросов.
		- При обработке финансовых транзакций может быть применена цепочка обязанностей для проверки различных условий: 
			проверка баланса, проверка лимитов, верификация получателя и так далее.
*/

package main

import "fmt"

// User структура пользователя
type User struct {
	Username string
	Password string
}

// Handler интерфейс обработчика авторизации.
type Handler interface {
	Handle(user User) bool // Метод обработки пользователя.
	SetNextHandler(next Handler) // Метод для установки следующего обработчика в цепочке.
}

// Структура хранит ссылку на след.обработчик.
type UsernameCheckHandler struct {
	next Handler
}

/* Проверяем имя пользователя. Если имя не пустое - 
   выводим информацию и если поле next структуры UsernameCheckHandler 
   указывает на след.обработчик, передаём ему пользователя и возвращаем
   результат вызова след.обработчика. */
func (u *UsernameCheckHandler) Handle(user User) bool {
	if user.Username != "" {
		fmt.Println("Username is valid")
		if u.next != nil {
			return u.next.Handle(user)
		}
		return true
	}
	fmt.Println("Username is invalid")
	return false
}

// Метод устанавливает следующий обработчик.
func (u *UsernameCheckHandler) SetNextHandler(next Handler) {
	u.next = next
}

// Обработчик проверки пароля, который хранит ссылку на след.обработчик.
type PasswordCheckHandler struct {
	next Handler
}

// Метод обработки пароля, аналогично методу обработки имени.
func (p *PasswordCheckHandler) Handle(user User) bool {
	if user.Password == "secret" {
		fmt.Println("Password is valid")
		if p.next != nil {
			return p.next.Handle(user)
		}
		return true
	}
	fmt.Println("Password is invalid")
	return false
}

// Метод устанавливает следующий обработчик.
func (p *PasswordCheckHandler) SetNextHandler(next Handler) {
	p.next = next
}

func main() {
	// Создаём пользователя.
	user := User{"user123", "secret"}

	// Создаем обработчики.
	usernameHandler := &UsernameCheckHandler{}
	passwordHandler := &PasswordCheckHandler{}

	// Строим цепочку обработчиков, пароль следует за именем.
	usernameHandler.SetNextHandler(passwordHandler)

	// Пытаемся авторизовать пользователя(в зависимости от ответа: true, false).
	if usernameHandler.Handle(user) {
		fmt.Println("User is authenticated")
	} else {
		fmt.Println("User authentication failed")
	}
}

