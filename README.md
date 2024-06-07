# PokéFetch

<img src="demo.png" style="width: 50%">

Displays a random Pokémon from the PokéAPI in your terminal.

You can also get shiny Pokémon with a 1/2 chance; you can change the probability in the `display.go` file.

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

4. Optionally, you can pass an `id` argument to display a specific Pokémon:

```sh
./pokefetch -id 1 # or pokefetch.exe -id 1 on Windows (displays Bulbasaur)
```

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
