# pokedex

A command-line Pokedex built in Go using REPL interface. This project interacts with the PokeAPI to fetch data 

## Features
Interactive REPL-based command system.

Fetches live data from the PokéAPI.

Explore Pokémon locations and areas.

Catch and inspect Pokémon.

In-memory caching to reduce redundant API calls.

## Commands
help - Displays all commands

map - Get next page of locations 

mapb - Get previous page of locations

explore <location> - List of all pokemon in current location

catch <pokemon> - Try to catch specified pokemon

inspect <pokemon> - Inspect specified pokemon

pokedex - List of caught pokemon

exit - Exit the podedex
