## Агрегатор данных stageSystem  

### Описание системы
Система представляет собой микросервис, агрегирующий данные с различных систем.  

### Цель системы
Минимизация ресурсов со стороны службы поддержки компании. Информирование клиентов о текущем состоянии систем посредством веб-интерфейса. 

### Задачи системы
- Сбор данных о системе SMS.
- Сбор данных о системе MMS.
- Сбор данных о системе VoiceCall.
- Сбор данных о системе email.
- Сбор данных о системе Billing.
- Сбор данных о системе Support.
- Сбор данных о системе истории инцидентов.
- Визуализация и вывод данных на веб-страницу

### Описание каталогов
- config - содержит конфигурационные настройки, а также файл .env
- pkg - содержит готовые пакеты для сбора и обработки данных, а также вспомогательные функции
- simulator - содержит файлы для эмуляции работы подсистем
- web - содержит файлы для отображения собранных системой данных на веб-странице 

### Запуск системы
Запуск системы можно разбить на два этапа: запуск симулятора данных и запуск системы-агрегатора данных.
1. Запустите симулятор  
Перейдите в каталог simulator и запустите main.go (чтобы файлы сгенерировались в каталог simulator, запускать нужно именно находясь в каталоге simulator):    
`$ cd stageSystem/simulator`  
`$ go run main.go`  
Симулятор обеспечивает получение данных по API. Для корректной работы системы, симулятор должен быть запущен.   

2. Откроейте новый терминал, перейдите в корневой каталог проекта и запустите main.go  
 `$ go run main.go`  
Перейдите по адресу http://127.0.0.1:8585 (может быть изменен в .env), если выводится сообщение "OK", то сервис работает. Теперь вы можете открыть файл status_page.html (каталог web) в браузере и визуально отследить состояния всех систем, по которым были собраны данные.

### Особенности работы 
При каждом запуске симулятор генерирует новые данные, что позволяет произвести отладку системы на разных данных.
Часть данных доступна в виде файлов. Пути к файлам указаны в файле .env (каталог config).
Часть данных доступна через API. По-умолчанию симулятор открывает соединение по адресу **127.0.0.1:8383**. 
По-умолчанию система отдаёт данные в формате JSON под аресу: http://127.0.0.1:8585/api. Данный адрес также прописан в файле main.js (каталог web) для визуального представления полученных данных.

### Возможные ошибки
По мере сбора данных, в терминале могут появляться ошибки считывания и обработки данных. Например, такая ошибка:
`error: invalid string format at 49 row: <FRProtonil;404> | required delimiters not found`
говорит о том, что при сборе данных произошла ошибка разбора строки № 195. Поскольку данные некорректны, данная строка будет просто пропущена и это не приведет к завершению работы всей системы. 