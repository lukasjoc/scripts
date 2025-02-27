// fuzzypac provides a simole utility to search and install in the result of a
// `pacman -Qi` search using fzy fuzzy searching
package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

// PacmanInfo literally describes a single package info result from the `pacman -Qi` command
type PacmanInfo struct {
	// Name            : zxing-cpp
	Name string
	// Version         : 2.1.0-1
	Version string
	// Description     : A C++ library to decode QRCode
	Description string
	// Architecture    : x86_64
	Architecture string
	// URL             : https://github.com/nu-book/zxing-cpp
	URL string
	// Licenses        : Apache
	Licenses []string
	// Groups          : None
	Groups []string
	// Provides        : None
	Provides []string
	// Depends On      : gcc-libs
	DependsOn []string
	// Optional Deps   : None
	OptionalDeps []string
	// Required By     : prison
	RequiredBy []string
	// Optional For    : None
	OptionalFor []string
	// Conflicts With  : None
	ConflictsWith []string
	// Replaces        : None
	Replaces []string
	// Installed Size  : 1275.73 KiB
	InstalledSizeKiB float64
	// Packager        : Antonio Rojas <arojas@archlinux.org>
	Packager string
	// Build Date      : Wed 05 Jul 2023 09:59:12 PM CEST
	BuildDate time.Time
	// Install Date    : Fri 07 Jul 2023 08:16:21 PM CEST
	InstallDate time.Time
	// Install Reason  : Installed as a dependency for another package
	InstallReason string
	// Install Script  : No
	InstallScript string
	// Validated By    : Signature
	ValidatedBy string
}

// Name            : aalib
// Version         : 1.4rc5-17
// Description     : ASCII art graphic library
// Architecture    : x86_64
// URL             : https://aa-project.sourceforge.net/aalib/
// Licenses        : LGPL
// Groups          : None
// Provides        : None
// Depends On      : glibc  gpm  libx11  ncurses  slang
// Optional Deps   : xorg-fonts-misc: x11 driver
//                   xorg-mkfontscale: x11 driver
// Required By     : gst-plugins-good
// Optional For    : None
// Conflicts With  : None
// Replaces        : None
// Installed Size  : 287.53 KiB
// Packager        : Balló György <bgyorgy@archlinux.org>
// Build Date      : Thu 22 Jun 2023 01:57:48 AM CEST
// Install Date    : Sat 01 Jul 2023 09:52:51 AM CEST
// Install Reason  : Installed as a dependency for another package
// Install Script  : No
// Validated By    : Signature

const pacmanFormat = "Mon 2 Jan 2006 03:04:05 PM MST"

func parseFromLines(lines []string) (info PacmanInfo) {
	for _, l := range lines {
		parts := strings.SplitN(l, " : ", 2)
		if len(parts) < 2 {
			// fmt.Println("WARN: missing lines here. Fix this later.: ", parts)
			continue
		}

		field := strings.TrimSpace(parts[0])
		value := parts[1]
		switch field {
		case "Name":
			info.Name = value
			break
		case "Version":
			info.Version = value
			break
		case "Description":
			info.Description = value
			break
		case "Architecture":
			info.Architecture = value
			break
		case "URL":
			info.URL = value
			break
		case "Licenses":
			info.Licenses = strings.Split(value, " ")
			break
		case "Groups":
			info.Groups = strings.Split(value, " ")
			break
		case "Provides":
			info.Provides = strings.Split(value, " ")
			break
		case "Depends On":
			info.DependsOn = strings.Split(value, "  ")
			break
		case "Optional Deps":
			info.OptionalDeps = strings.Split(value, " ")
			break
		case "Required By":
			info.RequiredBy = strings.Split(value, " ")
			break
		case "Optional For":
			info.OptionalFor = strings.Split(value, " ")
			break
		case "Conflicts With":
			info.ConflictsWith = strings.Split(value, " ")
			break
		case "Replaces":
			info.Replaces = strings.Split(value, " ")
			break
		case "Installed Size":
			info.InstalledSizeKiB, _ = strconv.ParseFloat(strings.Split(value, " ")[0], 2)
			break
		case "Packager":
			info.Packager = value
			break
		case "Build Date":
            info.BuildDate, _ = time.Parse(pacmanFormat, value)
			break
		case "Install Date":
			info.InstallDate, _ = time.Parse(pacmanFormat, value)
			break
		case "Install Reason":
			info.InstallReason = value
			break
		case "Install Script":
			info.InstallScript = value
			break
		case "Validated By":
			info.ValidatedBy = value
			break
		}
	}
	return info
}

func main() {
	stdout, err := exec.Command("pacman", "-Q", "-i").Output()
	if err != nil {
		fmt.Fprintf(os.Stdout, "wow stdout not working for you huh??\n")
	}
	items := strings.Split(string(stdout), "\n\n")[0:1]
	for _, item := range items {
		info := parseFromLines(strings.Split(item, "\n"))
        pretty, _ := json.MarshalIndent(info, "", "  ")
		fmt.Println(string(pretty))
	}
}
