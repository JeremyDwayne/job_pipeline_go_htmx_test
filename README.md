# Job Pipeline

- Track your job applications and interviews in one place.
- Record notes during or after interviews
- Keep track of recruiter and the last time you spoke with them
- Keep track of the interviewers and their contact information for thank you emails and networking
- Track job offers and salary negotiations

## Technologies

- Go - Backend
- Tailwind - CSS
- Templ - Templating
- HTMX - Interactivity

### Tailwind

```bash
brew install tailwindcss
```

### Templ

https://templ.guide/

### Air

Air is required for hot reloading used in `make dev`
https://github.com/cosmtrek/air

## Makefile

This Makefile is designed to simplify common development tasks for your project. It includes targets for building your Go application, watching and building Tailwind CSS, generating templates, and running your development server using Air.

### Targets:

```bash
make tailwind-watch
```

This target watches the ./static/css/input.css file and automatically rebuilds the Tailwind CSS styles whenever changes are detected.

```
make tailwind-build
```

This target minifies the Tailwind CSS styles by running the tailwindcss command.

```
make templ-watch
```

This target watches for changes to \*.templ files and automatically generates them.

```
make templ-generate
```

This target generates templates using the templ command.

```
make dev
```

This target runs the development server using Air, which helps in hot-reloading your Go application during development.

```
make build
```

This target orchestrates the building process by executing the tailwind-build, templ-generate, and go build commands sequentially. It creates the binary output in the ./bin/ directory.
