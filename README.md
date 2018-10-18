# wol-go
A simple go implementation of Wake On LAN.

## Usage
```
wol-go v1.0
Usage: wol [--ip IP] [--port PORT] [HWADDR [HWADDR ...]]

Positional arguments:
  HWADDR

Options:
  --ip IP, -i IP         set the destination IP address [default: 255.255.255.255]
  --port PORT, -p PORT   set the destination port [default: 9]
  --help, -h             display this help and exit
  --version              display version and exit
```

## Build

```bash
git clone https://github.com/hcl/wol
cd wol
make deps
make
```

## Acknoledgment 
[Perl version Wakeonlan - José Pedro Oliveira](http://gsd.di.uminho.pt/jpo/software/wakeonlan.html)

## License
```
	  wol - A simple go implementation of Wake On LAN.
    Copyright (C) 2018  hcl(HydroChLorica)

    This program is free software: you can redistribute it and/or modify
    it under the terms of the GNU General Public License as published by
    the Free Software Foundation, either version 3 of the License, or
    (at your option) any later version.

    This program is distributed in the hope that it will be useful,
    but WITHOUT ANY WARRANTY; without even the implied warranty of
    MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
    GNU General Public License for more details.

    You should have received a copy of the GNU General Public License
    along with this program.  If not, see <https://www.gnu.org/licenses/>.
```