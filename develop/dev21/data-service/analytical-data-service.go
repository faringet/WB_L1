package data_service

import "fmt"

// Внешний пакет, который должен реализовывать эту структуру и этот интерфейс
// Отправляет только xml

type AnalyticalDataService interface {
	SendXmlData()
}

type XmlDocument struct {
}

func (doc XmlDocument) SendXmlData() {
	fmt.Println("Отправка xml документа!")
}
