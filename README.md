# charette

Filter your [no-intro](http://www.no-intro.org) roms.

## Dependencies

**All no-into sets are archived with 7zip, so you have to install the `7z` tool first.**

On mac you can install it with [homebrew](http://brew.sh):

    `brew install p7zip`

## Usage

    $ charette

By default, the no-intro archives are searched into current directory, but you can set the `-input` flag instead:

    $ charette -input="/PATH/TO/NO-INTRO/ARCHIVES/"

Selected roms are then copied into a new `/roms/` sub directory in current directory. You can change the output directory thanks to the `-output` flag:

    $ charette -input="/PATH/TO/NO-INTRO/ARCHIVES/"  -output="/PATH/TO/ROMS/"

### Regions

Default preferred regions setting is `France,Europe,World,USA,Japan`.

It means that when a game has several roms, then `charette` selects the `France` one if it exists, otherwise the `Europe` one, etc. If the game has no rom with preferred region, it is still selected with a random region rom, unless you specify the `-strict` flag, and in that case the game is skipped.

You can change the `regions` setting with the `-regions` flag:

    $ charette -regions=USA,World,Europe,Japan

If you want to keep only specified regions, set the `-strict` flag. For example, to keep only `USA` roms:

    $ charette -regions=USA -strict

### Insane mode

By default, `charette` skips all roms tagged with `Proto`, `Demo`, `Pirate`, `Beta`, `Sample`...

If you want to keep all those roms, then use the `-insane` flag:

    $ charette -insane

### Scraper

Once `charette` ended, you can scrap roms images thanks to [scraper](https://github.com/sselph/scraper).

Launch `scraper` in a rom dictory:

    $ cd /PATH/TO/ROMS/snes/
    $ scraper -max_width=375 -no_thumb=true

Or, for all directories:

    $ cd /PATH/TO/ROMS/
    $ find . -maxdepth 1 -mindepth 1 -type d -exec bash -c 'cd "$0" && scraper -max_width=375 -no_thumb=true' {} \;

Adds `-append` flag when updating a rom directory.

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
