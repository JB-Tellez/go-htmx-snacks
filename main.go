package main

import (
	"net/http"
	"os"

	"github.com/calvinmclean/babyapi"
	"github.com/calvinmclean/babyapi/extensions"
	"github.com/calvinmclean/babyapi/html"

	"github.com/go-chi/render"
)

const (
	allSNACKs         html.Template = "allSNACKs"
	allSNACKsTemplate string        = `<!doctype html>
<html>
	<head>
		<meta charset="UTF-8">
		<meta name="viewport" content="width=device-width, initial-scale=1.0">
		<title>Full Stack Snacks</title>
		<link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/uikit@3.17.11/dist/css/uikit.min.css" />
		<script src="https://unpkg.com/htmx.org@2.0.0"></script>
		<script src="https://unpkg.com/htmx-ext-sse@2.0.0/sse.js"></script>
	</head>

	<style>
	tr.htmx-swapping td {
		opacity: 0;
		transition: opacity 1s ease-out;
	}
	</style>

	<body>
		<div class="uk-container uk-margin-top text-center">
			<h1>Full Stack Snacks</h1>
		</div>
		<table class="uk-table uk-table-divider uk-margin-left uk-margin-right">
			<colgroup>
				<col>
				<col>
				<col style="width: 300px;">
			</colgroup>

			<thead>
				<tr>
					<th>Snack</th>
					<th>Description</th>
					<th></th>
				</tr>
			</thead>

			<tbody hx-ext="sse" sse-connect="/snacks/listen" sse-swap="newSNACK" hx-swap="beforeend">
				<form hx-post="/snacks" hx-swap="none" hx-on::after-request="this.reset()">
					<td>
						<input class="uk-input" name="Title" type="text">
					</td>
					<td>
						<input class="uk-input" name="Description" type="text">
					</td>
					<td>
						<button type="submit" class="uk-button uk-button-primary">Add SNACK</button>
					</td>
				</form>

				{{ range . }}
				{{ template "snackRow" . }}
				{{ end }}
			</tbody>
		</table>
	</body>
</html>`

	snackRow         html.Template = "snackRow"
	snackRowTemplate string        = `<tr hx-target="this" hx-swap="outerHTML">
	<td>{{ .Title }}</td>
	<td>{{ .Description }}</td>
	<td>
		{{- $color := "primary" }}
		{{- $disabled := "" }}
		{{- if .Eaten }}
			{{- $color = "secondary" }}
			{{- $disabled = "disabled" }}
		{{- end -}}

		<button class="uk-button uk-button-{{ $color }}"
			hx-put="/snacks/{{ .ID }}"
			hx-headers='{"Accept": "text/html"}'
			hx-include="this"
			{{ $disabled }}>

			<input type="hidden" name="Eaten" value="true">
			<input type="hidden" name="Title" value="{{ .Title }}">
			<input type="hidden" name="Description" value="{{ .Description }}">
			<input type="hidden" name="ID" value="{{ .ID }}">
			Eat
		</button>

		<button class="uk-button uk-button-danger" hx-delete="/snacks/{{ .ID }}" hx-swap="swap:1s">
			Delete
		</button>
	</td>
</tr>`
)

type SNACK struct {
	babyapi.DefaultResource

	Title       string
	Description string
	Eaten       *bool
}

func (t *SNACK) HTML(r *http.Request) string {
	return snackRow.Render(r, t)
}

type AllSNACKs struct {
	babyapi.ResourceList[*SNACK]
}

func (at AllSNACKs) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (at AllSNACKs) HTML(r *http.Request) string {
	return allSNACKs.Render(r, at.Items)
}

func createAPI() *babyapi.API[*SNACK] {
	api := babyapi.NewAPI("SNACKs", "/snacks", func() *SNACK { return &SNACK{} })

	api.AddCustomRootRoute(http.MethodGet, "/", http.RedirectHandler("/snacks", http.StatusFound))

	// Use AllSNACKs in the GetAll response since it implements HTMLer
	api.SetGetAllResponseWrapper(func(snacks []*SNACK) render.Renderer {
		return AllSNACKs{ResourceList: babyapi.ResourceList[*SNACK]{snacks}}
	})

	api.ApplyExtension(extensions.HTMX[*SNACK]{})

	// Add SSE handler endpoint which will receive events on the returned channel and write them to the front-end
	snackChan := api.AddServerSentEventHandler("/listen")

	// Push events onto the SSE channel when new SNACKs are created
	api.SetOnCreateOrUpdate(func(r *http.Request, t *SNACK) *babyapi.ErrResponse {
		if r.Method != http.MethodPost {
			return nil
		}

		select {
		case snackChan <- &babyapi.ServerSentEvent{Event: "newSNACK", Data: t.HTML(r)}:
		default:
			logger := babyapi.GetLoggerFromContext(r.Context())
			logger.Info("no listeners for server-sent event")
		}
		return nil
	})

	// Optionally setup redis storage if environment variables are defined
	api.ApplyExtension(extensions.KeyValueStorage[*SNACK]{
		KVConnectionConfig: extensions.KVConnectionConfig{
			RedisHost:     os.Getenv("REDIS_HOST"),
			RedisPassword: os.Getenv("REDIS_PASS"),
			Filename:      os.Getenv("STORAGE_FILE"),
			Optional:      true,
		},
	})

	html.SetMap(map[string]string{
		string(allSNACKs): allSNACKsTemplate,
		string(snackRow):  snackRowTemplate,
	})

	return api
}

func main() {
	api := createAPI()
	api.RunCLI()
}
