Верютин Михаил

Тестовое задание GO: 
https://docs.google.com/document/d/1-laS0wKfca9m3r0FOBkMI1GuZ6HSyC73/edit

Описание проекта: 
Учебный проект для летней производственной практики в компании MediaSoft.

Подготовительные действия (установки,
настройки и т.д.) для успешной работы:
- компилятор Go версии 1.22,
- Docker version: '3.7',
- Postman (для тестирования)

Описание, как запустить проект:
1. В папке проекта выполнить скрипт для создания Docker контейнера (Docker должен быть запущен)
  docker-compose up --build -d
2. Запустить скрипт main.go 
  cmd\rest\main.go
3. Посылать запросы по адресу запущенного сервера
  http://localhost:12342
