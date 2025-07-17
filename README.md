# Task-2

1.Проект представляет собой HTTP-API сервер, который обрабатывает запросы для прогнозирования медицинских показателей на основе входных данных. Сервер реализован на Go и предоставляет следующие эндпоинты:

HBA1C (гликированный гемоглобин)

TG (триглицериды)

HDL (липопротеины высокой плотности)

LDL (липопротеины низкой плотности)

FERR (ферритин)

LDLL (LDL-подобный показатель)

Каждый эндпоинт принимает GET-запросы, валидирует их и возвращает прогноз в формате JSON.

2. API Endpoints

Общий формат запроса: 

```
GET http://localhost:8080/predict/{indicator}?param1=value1&param2=value2...
```
Где {indicator} — один из поддерживаемых показателей (hba1c, tg, hdl, ldl, ferr, ldll).

3. Примеры вызовов

Тестирование происходит через Postman. 
Запрос прогноза HbA1C: 
```
GET http://localhost:8080/predict/hba1c?age=15&gender=1&rdw=10&wbc=10&rbc=10&hgb=10&hct=10&mcv=10&mch=10&mchc=10&plt=10&neu=10&eos=10&bas=10&lym=10&mon=10&soe=10&soe=10&chol=10&glu=10
```
Ответ:
```
{
    "uid": "web-client",
    "prediction": 7.83,
    "model": "hba1c"
}
```
Запрос прогноза TG:
```
GET http://localhost:8080/predict/tg?age=15&gender=1&rdw=10&wbc=10&rbc=10&hgb=10&hct=10&mcv=10&mch=10&mchc=10&plt=10&neu=10&eos=10&bas=10&lym=10&mon=10&soe=10&soe=10&chol=10&glu=10
```
Ответ:
```
{
    "uid": "web-client",
    "prediction": 4.59,
    "model": "tg"
}
```
Запрос прогноза FERR: 
```
GET http://localhost:8080/predict/ferr?age=20&gender=1&rdw=1&wbc=2&rbc=3&hgb=4&hct=5&mcv=6&mch=7&mchc=8&plt=9&neu=1&eos=2&bas=3&lym=4&mon=5&soe=6&crp=7
```
Ответ:
```
{
    "uid": "web-client",
    "prediction": 33.22,
    "model": "ferr"
}
```
Запрос прогноза LDLL: 
```
GET http://localhost:8080/predict/ldll?age=20&gender=1&chol=1&hdl=2&tg=3
```
Ответ:
```
{
    "uid": "web-client",
    "prediction": 0.55,
    "model": "ldll"
}
```
Запрос прогноза LDL: 
```
GET http://localhost:8080/predict/ldl?age=20&gender=1&rdw=1&wbc=2&rbc=3&hgb=4&hct=5&mcv=6&mch=7&mchc=8&plt=9&neu=1&eos=2&bas=3&lym=4&mon=5&soe=6&chol=7&glu=8
```
Ответ:
```
{
    "uid": "web-client",
    "prediction": 1.95,
    "model": "ldl"
}
```
Запрос прогноза HDL(отключен): 
```
GET http://localhost:8080/predict/hdl?age=20&gender=1&rdw=10&wbc=2&rbc=3&hgb=4&hct=5&mcv=6&mch=7&mchc=8&plt=9&neu=1&eos=2&bas=3&lym=4&mon=5&soe=6&chol=7&glu=8
```
Ответ:
```
{
    "details": "api error: {\"detail\":\"Файл модели не найден: Модель ЛПВП по данным ПМО не найдена ни локально, ни в S3\"}",
    "error": "Prediction failed"
}
```
Токен авторизации указан неверно. Ответ:
```
{
    "details": "Check your authorization token",
    "error": "Invalid token"
}
```
Токен авторизации не указан. Ответ:
```
{
    "details": "Format: Bearer <token>",
    "error": "Authorization header required"
}
```

4. Обработка ошибок

Сервер возвращает следующие HTTP-статусы:

200 OK — успешный прогноз

405 Method Not Allowed — неверный HTTP-метод

500 Internal Server Error — ошибка предсказания








