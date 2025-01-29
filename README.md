## API エンドポイント

### タスク一覧取得

GET /tasks
登録されているすべてのタスクを取得します。

```sh
curl -X GET http://localhost:8080/tasks
```

レスポンス

```json
{
  "data": [
    {
      "id": 1,
      "title": "タスクタイトル",
      "description": "タスクの説明",
      "status": "progress",
      "created_at": "2025-01-29T12:00:00Z"
    }
  ]
}

```

### タスク作成

POST /tasks
新しいタスクを登録します。

```sh
curl -X POST http://localhost:8080/tasks \
-H "Content-Type: application/json" \
-d '{"title": "新しいタスク", "description": "説明", "status": "pending"}'
```

レスポンス

```
{
  "id": 2,
  "title": "新しいタスク",
  "description": "説明",
  "status": "pending",
  "created_at": "2025-01-29T12:05:00Z"
}
```
