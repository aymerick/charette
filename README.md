# charette

Filter your [no-intro](http://www.no-intro.org) roms.

## Usage

    $ go run main.go -dir="/PATH/TO/snes/"

All unwanted roms are then moved to `/PATH/TO/snes/_GARBAGE_/` directory.

You probably want to use the `-saun` flag, that is equivalent to `-sane -unzip`:

    $ go run main.go -dir="/PATH/TO/snes/" -saun

With that flag, only sane roms are selected, then they are unziped.

### Regions

Default preferred regions setting is `France,Europe,World,USA,Japan`.

It means that when a game has several roms, then `charette` selects the `France` one if it exists, otherwise the `Europe` one, etc. If the game has no rom with preferred region, it is still selected with a random rom, unless you specify the `-leave-me-alone` flag, and in that case the game is skipped.

You can change set the `regions` setting with the `-regions` flag:

    $ go run main.go -dir="/PATH/TO/snes/" -regions=USA,World,Europe,Japan

If you want to keep only specified regions, set the `-leave-me-alone` flag. For example, to keep only `USA` roms:

    $ go run main.go -dir="/PATH/TO/snes/" -regions=USA -leave-me-alone

## MAME

When working on `mame` roms, you have to set the `-mame` flag:

    $ go run main.go -dir="/PATH/TO/mame/" -mame

### Sane

The `-sane` flag skips all roms marked as `Proto`, `Demo`, `Pirate`, `Beta`, `Sample`, etc.

    $ go run main.go -dir="/PATH/TO/snes/" -sane

### Unzip

If you want to unzip selected roms, use the `-unzip` flag:

    $ go run main.go -dir="/PATH/TO/snes/" -sane -unzip

The roms are extracted and the `.zip` files are deleted.

Note that `.7z` files are NOT supported for the moment.

### Scraper

If you want to scrap roms images, install [scraper](https://github.com/sselph/scraper).

First remove the `_GARBAGE_` subdirectory (otherwise it will be scraped), then launch `scraper`:

    $ cd /PATH/TO/snes/
    $ scraper -max_width=375 -no_thumb=true

For MAME roms:

    $ cd /PATH/TO/mame/
    $ scraper -max_width=375 -no_thumb=true -name -mame_img=m,t,s


## Allowed regions

Some [no-intro](http://www.no-intro.org) file names are buggy, so here is the hardcoded list of allowed regions:

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
