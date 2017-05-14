package main

import (
	"fmt"
)

func filterOutGhost(set map[string]BigApp) (map[string]BigApp, error) {
	retSet := make(map[string]BigApp)
	if err := deepCopy(&retSet, set); err != nil {
		return nil, err
	}

	for bigAppId, big := range retSet {
		if big.IsGhost {
			delete(retSet, bigAppId)
			continue
		}

		subs := big.SubApps
		for subAppId, sub := range subs {
			if sub.IsGhost {
				delete(subs, subAppId)
			}
		}
	}
	return retSet, nil
}

func deepCopy(dst, src interface{}) error {
	var buf bytes.Buffer
	if err := gob.NewEncoder(&buf).Encode(src); err != nil {
		return err
	}
	return gob.NewDecoder(bytes.NewBuffer(buf.Bytes())).Decode(dst)
}
