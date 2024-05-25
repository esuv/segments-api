<img align="right" width="50%" src="./images/gopher.png">

# Books App

Basic application for managing segments

## API Reference

#### Get all books

```http
  GET /segments
```

| Parameter | Type     | Description                |
|:----------|:---------|:---------------------------|
| `api_key` | `string` | **Required**. Your API key |

#### Get segment by ID

```http
  GET /segments/${id}
```

| Parameter | Type     | Description                       |
|:----------|:---------|:----------------------------------|
| `id`      | `string` | **Required**. Id of item to fetch |

#### add(num1, num2)

Takes two numbers and returns the sum.

## Deployment

To deploy this project run

```bash
    $ git clone ...
    $ make build
    $ docker-compose up -d
    $ docker-compose run --rm app ./app
```

