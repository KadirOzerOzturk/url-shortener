
# URL Shortener API








## API Reference
Before sending the API request, make sure to check the base URL in the .env file.
#### Get all url
```http
  GET /url
```
Retrieves all shortened URL information without requiring any parameters or JSON body.
#### Shorten 

Converts a long URL into a short URL and returns the details of the shortened URL.
```http
  POST /url
```
```json
{
  "original_url": "https://www.examplelongurl.com"
}
```
**Response**
```json
{
    "id": "3c3d7aae-0f7c-483a-9d98-9b89eaf7635e",
    "created_at": "2024-10-15T09:06:40.32248376+03:00",
    "updated_at": "2024-10-15T09:06:40.32248376+03:00",
    "deleted_at": null,
    "original_url": "https://www.examplelongurl.com",
    "shortened_url": "http://baseurl/url/PapHcdtG8",
    "usage_count": 0,
    "expires_at": "2024-10-16T09:06:40.308709496+03:00"
}
```


#### Redirect 
Redirects to the original long URL using the provided shortened string.
```http
  GET /url/:shorten_str
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `shorten_str`      | `string` | **Required**. shorten_str to redirect |

#### Delete

Deletes the shortened URL using the provided shortened string.
```http
  DELETE /url/:shorten_str
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `shorten_str`      | `string` | **Required**. shorten_str to delete |

Response

``` json
{
    "message": "Item deleted successfully"
} 
```

Error
``` json
{
    "error": "Record not found"
}
```

## Features

- **Redis Limiter**: Utilizes Redis for rate limiting to control the number of requests.
- **Unique Shortened URLs**: Ensures that each shortened URL is unique.
- **IP Tracking**: Tracks how many times each IP address accesses a specific URL.

