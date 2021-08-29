# Тестовое задание для GORA Studio
Реализовать HTTP API фотогалереи(загрузка, просмотр, удаление фото).
## Возможности API:
* Загрузка фото
* Просмотр списка фото
* Удаление фото
## Используемые фреймворки:
* https://github.com/sirupsen/logrus - логирование
* https://github.com/spf13/viper - конфиги
* https://github.com/mattn/go-sqlite3 - SQLite
* https://github.com/gin-gonic/gin - web
* https://github.com/swaggo/swag - swagger документация
## Запуск:
По пути ```..\gora-studio-test-task\``` запустите:
```sh 
go run .\cmd\app\main.go
```
Возможно, придётся установить gcc
## Скриншоты
### Swagger
![image](https://user-images.githubusercontent.com/61324182/131260533-cedd1ace-df9a-4e8b-ae2a-47d14c185a63.png)
![image](https://user-images.githubusercontent.com/61324182/131260702-417a4394-6f76-4d8c-b033-15fcc48ae156.png)
### Загрузка фото
![image](https://user-images.githubusercontent.com/61324182/131260569-f1d4dbb6-acf5-437e-8543-730774fdabbd.png)
![image](https://user-images.githubusercontent.com/61324182/131260590-b330bb2d-79f9-4272-a7c9-029338674abb.png)
### Просмотр списка фото
![image](https://user-images.githubusercontent.com/61324182/131260627-b3ad0562-11d6-4934-97a1-1e645b3a0c55.png)
### Удаление фото
![image](https://user-images.githubusercontent.com/61324182/131260654-107858d3-0734-4b7c-a927-5bcaca302e97.png)
![image](https://user-images.githubusercontent.com/61324182/131260669-12dba156-7387-4e84-9537-35a694552718.png)
### Таблица из БД
![image](https://user-images.githubusercontent.com/61324182/131262205-d33a406b-1949-4dba-8d7b-fa11071ccf26.png)
### Логи
![image](https://user-images.githubusercontent.com/61324182/131262215-d3061ca9-67ca-4efd-9ddf-64c1b173875d.png)
