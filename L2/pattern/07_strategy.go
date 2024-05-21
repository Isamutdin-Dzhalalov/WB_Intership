/*
	Паттерн Стратегия - это поведенческий паттерн проектирования,
	который позволяет определять семейство алгоритмов, 
	инкапсулировать каждый из них и делать их взаимозаменяемыми. 

	+:
		- Паттерн позволяет контексту не зависеть от конкретных реализаций алгоритмов. 
		- Легко добавлять новые алгоритмы, просто реализуя интерфейс стратегии. 
	
	-:
		- Паттерн увеличивает количество классов и объектов в системе, что может усложнить понимание и сопровождение кода.
		- Вызов метода через интерфейс может быть медленнее, чем прямой вызов метода, особенно если алгоритмы просты и вызываются часто.

	Примеры:
		Веб-приложения: Различные стратегии аутентификации пользователей.
		Платежные системы: Разные стратегии оплаты (например, кредитная карта, PayPal, банковский перевод).

*/

package main

import "fmt"

// Определяем интерфейс стратегии.
type PaymentStrategy interface {
	Pay(amount float64) string
}

// Реализация стратегии для оплаты кредитной картой.
type CreditCardPayment struct{}

func (c *CreditCardPayment) Pay(amount float64) string {
	return fmt.Sprintf("Paid %.2f using credit card", amount)
}

// Реализация стратегии для оплаты через PayPal.
type PayPalPayment struct{}

func (p *PayPalPayment) Pay(amount float64) string {
	return fmt.Sprintf("Paid %.2f using PayPal", amount)
}

// Контекст, который использует стратегию для оплаты.
type PaymentContext struct {
	paymentStrategy PaymentStrategy
}

func (pc *PaymentContext) SetPaymentStrategy(strategy PaymentStrategy) {
	pc.paymentStrategy = strategy
}

func (pc *PaymentContext) MakePayment(amount float64) string {
	return pc.paymentStrategy.Pay(amount)
}

func main() {
	// Создаем контекст оплаты.
	paymentContext := PaymentContext{}

	// Используем стратегию оплаты кредитной картой.
	paymentContext.SetPaymentStrategy(&CreditCardPayment{})
	fmt.Println(paymentContext.MakePayment(100.0))

	// Используем стратегию оплаты через PayPal.
	paymentContext.SetPaymentStrategy(&PayPalPayment{})
	fmt.Println(paymentContext.MakePayment(50.0))
}

