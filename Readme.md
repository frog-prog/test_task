##### My ToDo App

Код приложения разделен на несколько слоев. <br>
Перечислять слои буду от внешних.

1. Слой маршрутов. <br>
Маршруты перечисляются при инициализации приложения в пакете ___App___, в функции ___Start()___.<br>
В них происходит парсинг строки запроса и тела запроса, а также передача полученных данных в контроллеры.
2. Слой контроллеров. <br>
Контроллеры располагаются в папке ___Controllers___. <br>
Они предназначены исключительно для того, чтобы "распаковать" данные из контекста запроса, <br>
передать их в сервис, получить результат его работы и вернуть клиенту ответ.
3. Слой сервисов. <br>
Сервисы располагаются в папке ___Services___. <br>
Они предназначены для реализации бизнес-логики и принимают в качестве аргументов<br>
относительно примитивные значения, возвращая настолько же простые значения ошибки или результата.<br> 
Это позволяет сделать бизнес-логику тестируемой.
4. Слой моделей. <br>
Модели располагаются в папке ___Database/Models___ и реализуют методы взаимодействия с СУБД.<br> 
Принимают и возвращают еще более простые значения чем сервисы<br>

Запуск приложения происходит одной командой:<br>
___docker compose up -d___<br>
При этом массивы данных из бд не остаются на хост-машине<br>
Описание API в формате openapi v3 расположено в файле api_description.yaml

