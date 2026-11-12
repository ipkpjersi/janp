/*
   Copyright 2014 Franc[e]sco (lolisamurai@tfwno.gf)
   This file is part of sharenix.
   sharenix is free software: you can redistribute it and/or modify
   it under the terms of the GNU General Public License as published by
   the Free Software Foundation, either version 3 of the License, or
   (at your option) any later version.
   sharenix is distributed in the hope that it will be useful,
   but WITHOUT ANY WARRANTY; without even the implied warranty of
   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
   GNU General Public License for more details.
   You should have received a copy of the GNU General Public License
   along with sharenix. If not, see <http://www.gnu.org/licenses/>.
*/

package janplib

import (
	"encoding/json"
	"io/ioutil"
	"path"
	"os"
)
var debug = true
// A Config holds the janp settings loaded from janp.json
type Config struct {
	XineramaHead     uint32  `json:",omitempty"`
	NotificationTime float64 `json:",omitempty"`
	SaveFolder       string  `json:",omitempty"`
}

func LoadConfig() (cfg *Config, err error) {
	cfg = &Config{}
	cfg.XineramaHead = 0
	cfg.NotificationTime = 3

	exeFolder, err := GetExeDir()
	if err != nil {
		return
	}

	cfgName := "janp.json"

	cfgPath := os.Getenv("XDG_CONFIG_HOME")
	if cfgPath == "" {
		cfgPath = path.Join(GetHome(), ".config")
	}
	cfgPaths := [...]string{
		path.Join(GetHome(), "."+cfgName),
		path.Join(exeFolder, cfgName),
		path.Join("/etc/", cfgName),
		path.Join("/usr/local/etc/", cfgName),
		path.Join(cfgPath, "janp", cfgName),
	}

	var file []byte

	for _, path := range cfgPaths {
		file, err = ioutil.ReadFile(path)
		if err == nil {
			break
		}
	}

	if err != nil {
		return
	}

	err = json.Unmarshal(file, &cfg)
	return
}
