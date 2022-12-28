# brightness

A package that allows changing the screen brightness on linux machines.
It does so by modifiying the appropriate files in `/sys/class/backlight/intel_backlight`

It exposes a percentage based interface that can be either called either programatcally
via the bright package or via a cli in the cmd folder
