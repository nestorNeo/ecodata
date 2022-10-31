package models

const (
	INDEXCONTENT = `
	<!DOCTYPE html>
<html>
  <head>
    <meta charset="utf-8">
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <title>ecodata api</title>
    <link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/bulma@0.9.4/css/bulma.min.css">
  </head>
  <body>

    <section class="hero is-primary">
        <div class="hero-body">
          <p class="title">
            Welcome to ecodata api V1
          </p>
          <p class="subtitle">
            Work in progress
            <div class="buttons">
                <button class="button is-info"  onclick="window.location.href='https://www.gdl.cinvestav.mx';">source</button>              
            </div>
          </p>
        </div>
      </section>
  </body>
</html>
	`

	NOTFOUND = `
	<!DOCTYPE html>
<html>
  <head>
    <meta charset="utf-8">
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <title>Not found</title>
    <link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/bulma@0.9.4/css/bulma.min.css">
  </head>
  <body>
  <section class="section">
    <div class="container">
      <h1 class="title">
        Not found
      </h1>
      <p class="subtitle">
        please contact developers
      </p>
    </div>
  </section>
  </body>
</html>
	`
)
