package main

/*
Реализовать паттерн «адаптер» на любом примере.
*/

/*Пусть внешний сервис (analytical-data-service) работает только с форматом xml.

Предположим, что мой сервис работает только с json. Задача состоит, чтобы транслировать свои json объекты
в объекты xml используя только AnalyticalDataService interface.

Решением этой задачи - адаптер.
*/