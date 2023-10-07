# Color [![Static Badge](https://img.shields.io/badge/Color%20Shades-3A2BE2)](https://github.com/bedirhandogan/color#color-shades) 

You can use the color pack to colorize the output in the terminal with custom formatters or create custom colors with RGB and [ANSI 256](https://robotmoon.com/256-colors/) index. 
All formatters are parsed as ANSI Escape Code.

![code sample](https://github.com/bedirhandogan/color/assets/59766658/1cd9292b-f6da-43d1-a247-5bd65609d96f)

## Install
```bash
$ go get github.com/bedirhandogan/color
```

## Uses
### Single Colorize
```golang
// Print all text in one color
fmt.Println(color.Colorize("%Red Print this text in color."))

// Print all text in one color, specifying tone
fmt.Println(color.Colorize("%Magenta80 Print this text in magenta 80 shades."))

// Color and print the background of all text by adding bg in front of it
fmt.Println(color.Colorize("%BgCyan Color this text with a cyan background and print it."))
fmt.Println(color.Colorize("%BgYellow50 Color this text with a yellow 50 shades background and print it"))
```

### Multiple Colorize
```golang
// Print all text with multiple colors
fmt.Println(color.Colorize("%Blue Print this text with %Green multiple colors."))

// Print all text with multiple color, specifying tone
fmt.Println(color.Colorize("%Blue70 Print this text with %Green20 multiple colors."))

// Print all text with multiple background color
fmt.Println(color.Colorize("%Yellow20 Print this text with %Blue multiple colors."))
``` 

### Use [SGR](https://en.wikipedia.org/wiki/ANSI_escape_code#SGR_(Select_Graphic_Rendition)_parameters) Parameters
```golang
// Normalize text
fmt.Println(color.Colorize("%Blue Print this text in color with reset. %Reset"))

// Bold text
fmt.Println(color.Colorize("%Bold %Yellow80 Print this text in color with bold."))
fmt.Println(color.Colorize("%Bold Print this text in bold."))

// Italic text
fmt.Println(color.Colorize("%Italic %Cyan50 Print this text in color with italicized."))
fmt.Println(color.Colorize("%Italic Print this text in italicized."))

// You can also format as Blink, Underline, Overline, Invert & Strike.
``` 

### Native Formatters
```golang
// This way you can still use native formatters
fmt.Printf(color.Colorize("%Cyan Print this text with %s"), "native formatter")
```

### Create New Color
```golang
// Create and use new solid color with ANSI 256 index
NewColor("FaintRed", 124)
fmt.Println(color.Colorize("%FaintRed Print this text faint red color."))

// Create and use new background color with ANSI 256 index
NewColor("BgFaintRed", 124)
fmt.Println(color.Colorize("%BgFaintRed Print this text faint red background color."))

// Create and use new solid color with RGB
NewColor("LightRed", 255, 192, 203)
fmt.Println(color.Colorize("%FaintRed Print this text light red color."))

// Create and use new background color with RGB
NewColor("BgLightRed", 255, 192, 203)
fmt.Println(color.Colorize("%BgFaintRed Print this text light red background color."))
```

## [Colors](#color-shades)
Below are the default color shades you can use:
- [Red Shades](#red-shades)
- [Green Shades](#green-shades)
- [Yellow Shades](#yellow-shades)
- [Cyan Shades](#cyan-shades)
- [Blue Shades](#blue-shades)
- [Magenta Shades](#magenta-shades)
- [White Shades](#white-shades)
- [Black Shades](#black-shades)

### [Red Shades](#red-shades)

| Color Name | RGB           | Hex       | Preview                                                            |
|------------|---------------|-----------|--------------------------------------------------------------------|
| Red        | `(255, 0, 0)` | `#FF0000` | ![Loading...](https://via.placeholder.com/15/FF0000/000000?text=+) |
| Red10      | `(240, 0, 0)` | `#F00000` | ![Loading...](https://via.placeholder.com/15/F00000/000000?text=+) |
| Red20      | `(224, 0, 0)` | `#E00000` | ![Loading...](https://via.placeholder.com/15/E00000/000000?text=+) |
| Red30      | `(208, 0, 0)` | `#D00000` | ![Loading...](https://via.placeholder.com/15/D00000/000000?text=+) |
| Red40      | `(192, 0, 0)` | `#C00000` | ![Loading...](https://via.placeholder.com/15/C00000/000000?text=+) |
| Red50      | `(176, 0, 0)` | `#B00000` | ![Loading...](https://via.placeholder.com/15/B00000/000000?text=+) |
| Red60      | `(160, 0, 0)` | `#A00000` | ![Loading...](https://via.placeholder.com/15/A00000/000000?text=+) |
| Red70      | `(144, 0, 0)` | `#900000` | ![Loading...](https://via.placeholder.com/15/900000/000000?text=+) |
| Red80      | `(128, 0, 0)` | `#800000` | ![Loading...](https://via.placeholder.com/15/800000/000000?text=+) |
| Red90      | `(112, 0, 0)` | `#700000` | ![Loading...](https://via.placeholder.com/15/700000/000000?text=+) |
| Red100     | `(96, 0, 0)`  | `#600000` | ![Loading...](https://via.placeholder.com/15/600000/000000?text=+) |

### [Green Shades](#green-shades)

| Color Name | RGB           | Hex       | Preview                                                            |
|------------|---------------|-----------|--------------------------------------------------------------------|
| Green      | `(0, 255, 0)` | `#00FF00` | ![Loading...](https://via.placeholder.com/15/00FF00/000000?text=+) |
| Green10    | `(0, 240, 0)` | `#00F000` | ![Loading...](https://via.placeholder.com/15/00F000/000000?text=+) |
| Green20    | `(0, 224, 0)` | `#00E000` | ![Loading...](https://via.placeholder.com/15/00E000/000000?text=+) |
| Green30    | `(0, 208, 0)` | `#00D000` | ![Loading...](https://via.placeholder.com/15/00D000/000000?text=+) |
| Green40    | `(0, 192, 0)` | `#00C000` | ![Loading...](https://via.placeholder.com/15/00C000/000000?text=+) |
| Green50    | `(0, 176, 0)` | `#00B000` | ![Loading...](https://via.placeholder.com/15/00B000/000000?text=+) |
| Green60    | `(0, 160, 0)` | `#00A000` | ![Loading...](https://via.placeholder.com/15/00A000/000000?text=+) |
| Green70    | `(0, 144, 0)` | `#009000` | ![Loading...](https://via.placeholder.com/15/009000/000000?text=+) |
| Green80    | `(0, 128, 0)` | `#008000` | ![Loading...](https://via.placeholder.com/15/008000/000000?text=+) |
| Green90    | `(0, 112, 0)` | `#007000` | ![Loading...](https://via.placeholder.com/15/007000/000000?text=+) |
| Green100   | `(0, 96, 0)`  | `#006000` | ![Loading...](https://via.placeholder.com/15/006000/000000?text=+) |

### [Yellow Shades](#yellow-shades)

| Color Name | RGB             | Hex       | Preview                                                            |
|------------|-----------------|-----------|--------------------------------------------------------------------|
| Yellow     | `(255, 255, 0)` | `#FFFF00` | ![Loading...](https://via.placeholder.com/15/FFFF00/000000?text=+) |
| Yellow10   | `(240, 240, 0)` | `#F0F000` | ![Loading...](https://via.placeholder.com/15/F0F000/000000?text=+) |
| Yellow20   | `(224, 224, 0)` | `#E0E000` | ![Loading...](https://via.placeholder.com/15/E0E000/000000?text=+) |
| Yellow30   | `(208, 208, 0)` | `#D0D000` | ![Loading...](https://via.placeholder.com/15/D0D000/000000?text=+) |
| Yellow40   | `(192, 192, 0)` | `#C0C000` | ![Loading...](https://via.placeholder.com/15/C0C000/000000?text=+) |
| Yellow50   | `(176, 176, 0)` | `#B0B000` | ![Loading...](https://via.placeholder.com/15/B0B000/000000?text=+) |
| Yellow60   | `(160, 160, 0)` | `#A0A000` | ![Loading...](https://via.placeholder.com/15/A0A000/000000?text=+) |
| Yellow70   | `(144, 144, 0)` | `#909000` | ![Loading...](https://via.placeholder.com/15/909000/000000?text=+) |
| Yellow80   | `(128, 128, 0)` | `#808000` | ![Loading...](https://via.placeholder.com/15/808000/000000?text=+) |
| Yellow90   | `(112, 112, 0)` | `#707000` | ![Loading...](https://via.placeholder.com/15/707000/000000?text=+) |
| Yellow100  | `(96, 96, 0)`   | `#606000` | ![Loading...](https://via.placeholder.com/15/606000/000000?text=+) |


### [Cyan Shades](#cyan-shades)

| Color Name | RGB             | Hex       | Preview                                                            |
|------------|-----------------|-----------|--------------------------------------------------------------------|
| Cyan       | `(0, 255, 255)` | `#00FFFF` | ![Loading...](https://via.placeholder.com/15/00FFFF/000000?text=+) |
| Cyan10     | `(0, 240, 240)` | `#00F0F0` | ![Loading...](https://via.placeholder.com/15/00F0F0/000000?text=+) |
| Cyan20     | `(0, 224, 224)` | `#00E0E0` | ![Loading...](https://via.placeholder.com/15/00E0E0/000000?text=+) |
| Cyan30     | `(0, 208, 208)` | `#00D0D0` | ![Loading...](https://via.placeholder.com/15/00D0D0/000000?text=+) |
| Cyan40     | `(0, 192, 192)` | `#00C0C0` | ![Loading...](https://via.placeholder.com/15/00C0C0/000000?text=+) |
| Cyan50     | `(0, 176, 176)` | `#00B0B0` | ![Loading...](https://via.placeholder.com/15/00B0B0/000000?text=+) |
| Cyan60     | `(0, 160, 160)` | `#00A0A0` | ![Loading...](https://via.placeholder.com/15/00A0A0/000000?text=+) |
| Cyan70     | `(0, 144, 144)` | `#009090` | ![Loading...](https://via.placeholder.com/15/009090/000000?text=+) |
| Cyan80     | `(0, 128, 128)` | `#008080` | ![Loading...](https://via.placeholder.com/15/008080/000000?text=+) |
| Cyan90     | `(0, 112, 112)` | `#007070` | ![Loading...](https://via.placeholder.com/15/007070/000000?text=+) |
| Cyan100    | `(0, 96, 96)`   | `#006060` | ![Loading...](https://via.placeholder.com/15/006060/000000?text=+) |

### [Blue Shades](#blue-shades)

| Color Name | RGB           | Hex       | Preview                                                            |
|------------|---------------|-----------|--------------------------------------------------------------------|
| Blue       | `(0, 0, 255)` | `#0000FF` | ![Loading...](https://via.placeholder.com/15/0000FF/000000?text=+) |
| Blue10     | `(0, 0, 240)` | `#0000F0` | ![Loading...](https://via.placeholder.com/15/0000F0/000000?text=+) |
| Blue20     | `(0, 0, 224)` | `#0000E0` | ![Loading...](https://via.placeholder.com/15/0000E0/000000?text=+) |
| Blue30     | `(0, 0, 208)` | `#0000D0` | ![Loading...](https://via.placeholder.com/15/0000D0/000000?text=+) |
| Blue40     | `(0, 0, 192)` | `#0000C0` | ![Loading...](https://via.placeholder.com/15/0000C0/000000?text=+) |
| Blue50     | `(0, 0, 176)` | `#0000B0` | ![Loading...](https://via.placeholder.com/15/0000B0/000000?text=+) |
| Blue60     | `(0, 0, 160)` | `#0000A0` | ![Loading...](https://via.placeholder.com/15/0000A0/000000?text=+) |
| Blue70     | `(0, 0, 144)` | `#000090` | ![Loading...](https://via.placeholder.com/15/000090/000000?text=+) |
| Blue80     | `(0, 0, 128)` | `#000080` | ![Loading...](https://via.placeholder.com/15/000080/000000?text=+) |
| Blue90     | `(0, 0, 112)` | `#000070` | ![Loading...](https://via.placeholder.com/15/000070/000000?text=+) |
| Blue100    | `(0, 0, 96)`  | `#000060` | ![Loading...](https://via.placeholder.com/15/000060/000000?text=+) |

### [Magenta Shades](#magenta-shades)

| Color Name | RGB             | Hex       | Preview                                                            |
|------------|-----------------|-----------|--------------------------------------------------------------------|
| Magenta    | `(255, 0, 255)` | `#FF00FF` | ![Loading...](https://via.placeholder.com/15/FF00FF/000000?text=+) |
| Magenta10  | `(240, 0, 240)` | `#F000F0` | ![Loading...](https://via.placeholder.com/15/F000F0/000000?text=+) |
| Magenta20  | `(224, 0, 224)` | `#E000E0` | ![Loading...](https://via.placeholder.com/15/E000E0/000000?text=+) |
| Magenta30  | `(208, 0, 208)` | `#D000D0` | ![Loading...](https://via.placeholder.com/15/D000D0/000000?text=+) |
| Magenta40  | `(192, 0, 192)` | `#C000C0` | ![Loading...](https://via.placeholder.com/15/C000C0/000000?text=+) |
| Magenta50  | `(176, 0, 176)` | `#B000B0` | ![Loading...](https://via.placeholder.com/15/B000B0/000000?text=+) |
| Magenta60  | `(160, 0, 160)` | `#A000A0` | ![Loading...](https://via.placeholder.com/15/A000A0/000000?text=+) |
| Magenta70  | `(144, 0, 144)` | `#900090` | ![Loading...](https://via.placeholder.com/15/900090/000000?text=+) |
| Magenta80  | `(128, 0, 128)` | `#800080` | ![Loading...](https://via.placeholder.com/15/800080/000000?text=+) |
| Magenta90  | `(112, 0, 112)` | `#700070` | ![Loading...](https://via.placeholder.com/15/700070/000000?text=+) |
| Magenta100 | `(96, 0, 96)`   | `#600060` | ![Loading...](https://via.placeholder.com/15/600060/000000?text=+) |

### [White Shades](#white-shades)

| Color Name | RGB               | Hex       | Preview                                                            |
|------------|-------------------|-----------|--------------------------------------------------------------------|
| White      | `(255, 255, 255)` | `#FFFFFF` | ![Loading...](https://via.placeholder.com/15/FFFFFF/000000?text=+) |
| White10    | `(240, 240, 240)` | `#F0F0F0` | ![Loading...](https://via.placeholder.com/15/F0F0F0/000000?text=+) |
| White20    | `(224, 224, 224)` | `#E0E0E0` | ![Loading...](https://via.placeholder.com/15/E0E0E0/000000?text=+) |
| White30    | `(208, 208, 208)` | `#D0D0D0` | ![Loading...](https://via.placeholder.com/15/D0D0D0/000000?text=+) |
| White40    | `(192, 192, 192)` | `#C0C0C0` | ![Loading...](https://via.placeholder.com/15/C0C0C0/000000?text=+) |
| White50    | `(176, 176, 176)` | `#B0B0B0` | ![Loading...](https://via.placeholder.com/15/B0B0B0/000000?text=+) |
| White60    | `(160, 160, 160)` | `#A0A0A0` | ![Loading...](https://via.placeholder.com/15/A0A0A0/000000?text=+) |
| White70    | `(144, 144, 144)` | `#909090` | ![Loading...](https://via.placeholder.com/15/909090/000000?text=+) |
| White80    | `(128, 128, 128)` | `#808080` | ![Loading...](https://via.placeholder.com/15/808080/000000?text=+) |
| White90    | `(112, 112, 112)` | `#707070` | ![Loading...](https://via.placeholder.com/15/707070/000000?text=+) |
| White100   | `(96, 96, 96)`    | `#606060` | ![Loading...](https://via.placeholder.com/15/606060/000000?text=+) |

### [Black Shades](#black-shades)

| Color Name | RGB               | Hex       | Preview                                                            |
|------------|-------------------|-----------|--------------------------------------------------------------------|
| Black      | `(0, 0, 0)`       | `#000000` | ![Loading...](https://via.placeholder.com/15/000000/000000?text=+) |
| Black10    | `(16, 16, 16)`    | `#101010` | ![Loading...](https://via.placeholder.com/15/101010/000000?text=+) |
| Black20    | `(32, 32, 32)`    | `#202020` | ![Loading...](https://via.placeholder.com/15/202020/000000?text=+) |
| Black30    | `(48, 48, 48)`    | `#303030` | ![Loading...](https://via.placeholder.com/15/303030/000000?text=+) |
| Black40    | `(64, 64, 64)`    | `#404040` | ![Loading...](https://via.placeholder.com/15/404040/000000?text=+) |
| Black50    | `(80, 80, 80)`    | `#505050` | ![Loading...](https://via.placeholder.com/15/505050/000000?text=+) |
| Black60    | `(96, 96, 96)`    | `#606060` | ![Loading...](https://via.placeholder.com/15/606060/000000?text=+) |
| Black70    | `(112, 112, 112)` | `#707070` | ![Loading...](https://via.placeholder.com/15/707070/000000?text=+) |
| Black80    | `(128, 128, 128)` | `#808080` | ![Loading...](https://via.placeholder.com/15/808080/000000?text=+) |
| Black90    | `(144, 144, 144)` | `#909090` | ![Loading...](https://via.placeholder.com/15/909090/000000?text=+) |
| Black100   | `(160, 160, 160)` | `#A0A0A0` | ![Loading...](https://via.placeholder.com/15/A0A0A0/000000?text=+) |

## Ups!!!
You scrolled down too much 😂