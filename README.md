# po-async-http-task

### Asynchroniczne schematy komunikacji przez HTTP

* W zdaniach będziemy symulowali pobieranie maili z serwera poczty elektronicznej.

#### Zadanie 1: Short polling
1. W wybranym frameworku webowym (np. Flask lub FastAPI) stwórz aplikację w pliku app.py.
2. Stwórz endpoint POST /task, na który wyślesz zadanie do serwera. Body zdania powinno wyglądać następująco:
```json
{
"email": "nEY9R@example.com",
"count": 5,
}
```
Adres do endpointu gdzie będzie sprawdzany status powinien zostać przekazany w headerze Location odpowiedzi.

3. Stwórz endpoint GET /task/{id}, na który pobierasz status zadania z serwera. Body odpowiedzi, powinien wyglądać następująco:
```json
{
"email": "nEY9R@example.com",
"status": "done" | "pending" | "error",
}
```
Statusy mogą być ustawiane losowo.
Jeśli status jest ustawiony na "done" w headerze Location powinien zostać umieszczony URL do kolejnego endpointu.

4. Stwórz endpoint GET /task_result/{id}, który zwróci json o następującej strukturze:
```json
{
"email": "nEY9R@example.com",
"emails": [<dowolne dane>],
}
```
5. W pythonie stwórz klienta, który będzie odpytywał stworzonych endpointów w schemacie short polling.

#### Zadanie 2: Long polling
1. Proszę wykonać endpointy jak w punktach 1-4, z tą różnicą, że w punkcie 3 w endpoincie będzie funkcjonowała pętla, która zwróci status "done" dopiero po 3 minutach.
2. W pythonie stwórz klienta, który będzie odpytywał stworzonych endpointów w schemacie long polling.
