# PokéFetch

<img src="demo.png" style="width: 50%">

Displays a random Pokémon from the PokéAPI in your terminal.

## Usage

1. Clone the repository.
2. Run the following command in the root directory of the repository to build for your system:

```sh
go build
```

3. Run the following command to display a random Pokémon:

```sh
./pokefetch # or pokefetch.exe on Windows
```

## Flags

- `-id`: The ID of the Pokémon to display (default: random).
  - If this is provided, it takes precedence over the `-name` flag.
- `-name`: The name of the Pokémon to display (default: random).
  - If both `-id` and `-name` are provided, `-id` takes precedence.
- `-shiny`: The chance of the Pokémon being shiny (default: 0.5).
  - This is a float between 0 and 1.

## Terminal Startup

You can add this to your shell's startup script to display a random Pokémon every time you open a new terminal window.

For example, for `.zshrc`:

```sh
# Display a random Pokémon on terminal startup (but not in VSCode)
if [ "$TERM_PROGRAM" != "vscode" ]
then
  .local/bin/pokefetch
fi
```
