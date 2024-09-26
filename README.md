# Songlib

## Entrypoints

* From Dockerfile: `docker build -t <app_name> . && docker run -d -p <loc_port>:<cont_port> --name <app_name> <cont_name> `
* From current machine: `rootdir $ cd cmd/server && go run main.go`

  **Attention**: Firstly you need to initialize DB & setup all your links to the .env file, without it service won`t work correctly(fatal).Also you need to create a network, if you want to use Dockerized version.

---

## API Routes

All methods described by the route `host:port/swagger/index.html`, but I`ll rewrite descriptions here:

* [POST] /song - Create a new song;
* [GET] /song/{id:string(uuid)} - Get song details;
* [PUT] /song/{id:string(uuid)} - Update song data (excluding date);
* [DELETE] /song/{id:string(uuid)} - Deletes a song;
* [POST] /songs?page=int,opt&size=int,opt - Gets a list of songs, either filtered or not.

More deep usage you can see by going to the ./swagger/index.html

---

## Stack usage

For this service, I preferred to use wide-supported technologies, as:

* Fiber/v2 - cool http.Server implementation works more than a 40-times faster than an ordinary one.
* Gorm - with him, usage of the DB is much simplier, let it`ll be the best decision for simplifying work with DB.
* Swagger - provides to generate cool and detailed documentation for the API`s.

I choosed those ones because of simplicity of setting-up a new project.

For example, creating DB using stdlib is much more complicated to handle, like this:

```go
type Method uint8

const (
	Connection Method = iota
	Transaction
)

type Storage interface {
	User() UserRepo
	Tool() ControllerRepo
	Camera() CameraRepo
	Create(timeout time.Duration, t Method) (DBContext, error)
	Commit(DBContext) error
	Rollback(DBContext) error
	Close() error
}

type DBContext interface {
	context.Context
	Connection() (DBConnection, error)
	Close() error
	Ctx() context.Context
}

type DBConnection interface {
	Exec(query string, args ...any) (sql.Result, error)
	QueryRow(query string, args ...any) *sql.Row
	Close() error
	Commit() error
	Rollback() error
}
```

And each of CRUD-ops with the models are really hard to handle in moment. The alternative is the **gorm.** 

```go
type Store interface {
	Migrate() error
	Song() SongRepository
	Close() error
}

type SongRepository interface {
	Create(song *model.Song) error
	GetList(filter map[string]interface{}, page, pageSize int) ([]*model.Song, error)
	GetByID(id string) (*model.Song, error)
	Update(id string, updatedData map[string]interface{}) error
	Delete(id string) error
}
```

The other directives handles as the internal methods of already realised structures of the package.

## Attention

Здесь уже буду на русском :)

На самом деле, не понял какая нужна пагинация на куплеты песни, т.к. полностью не видно ответа от API, и так выкручиваюсь эмуляциями получения тела запроса(./internal/httpclient). Есть ещё много, чего можно докрутить, с чем грех спорить, к примеру - таймауты при работе с БД, но так как в задаче этого не поставлено - выполняю всё, что требуется.

---

*@nemesidaa*
