# charette

Filter your no-intro roms

Usage:

    $ go run main.go -dir="/PATH/TO/snes/" -sane

All unwanted roms are moved to /PATH/TO/snes/_GARBAGE_/ directory.

## Allowed regions

    Asia
    Australia
    Brazil
    Canada
    China
    Denmark
    Europe
    Finland
    France
    Germany
    Hong Kong
    Italy
    Japan
    Korea
    Netherlands
    Russia
    Spain
    Sweden
    Taiwan
    Unknown
    USA
    World

## Test

To run all tests:

    $ go test ./...
