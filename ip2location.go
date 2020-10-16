// Package ip2location get geo location from given ip address
//
// ip data comes from www.ipcaa.org
package ip2location

import (
	"errors"
	"fmt"
)

// _locations data offset
const (
	offsetGeoID         = 0
	offsetChinaRegionID = 1
	offsetCountry       = 2
	offsetProvince      = 3
	offsetCity          = 4
)

var (
	ErrInvalidIp = errors.New("invalid ip address")
	ErrNotFound  = errors.New("ip location not found")
)

type Location struct {
	Country       string
	Province      string
	City          string
	GeoID         string
	ChinaRegionID string
}

// Get get geo location from ip address
//
func Get(ip string) (loc *Location, err error) {
	lowIdx, highIdx, ipNum, err := getDataIndex(ip)

	if err != nil {
		return
	}

	locationIdx, err := search(ipNum, lowIdx, highIdx)

	if err != nil {
		return
	}

	loc = &Location{
		Country:       _locations[locationIdx+offsetCountry],
		Province:      _locations[locationIdx+offsetProvince],
		City:          _locations[locationIdx+offsetCity],
		GeoID:         _locations[locationIdx+offsetGeoID],
		ChinaRegionID: _locations[locationIdx+offsetChinaRegionID],
	}

	return
}

// getDataIndex get search offset index
//
func getDataIndex(ip string) (lowIdx int, highIdx int, ipNum uint32, err error) {
	var byte1, byte2, byte3, byte4 uint32

	n, err := fmt.Sscanf(ip, "%d.%d.%d.%d", &byte1, &byte2, &byte3, &byte4)
	if err != nil || n != 4 {
		err = ErrInvalidIp
		return
	}

	if v, ok := _rangeIndex[int(byte1)]; ok {
		lowIdx = v[0] * ipDataGroupItems
		highIdx = v[1] * ipDataGroupItems
	} else {
		err = ErrInvalidIp
		return
	}

	ipNum = byte1<<24 + byte2<<16 + byte3<<8 + byte4

	return
}

// search ip location idx
//
func search(ipNum uint32, lowIdx int, highIdx int) (locationIdx int, err error) {
	if lowIdx == highIdx {
		if _ipData[lowIdx] <= ipNum && _ipData[lowIdx+1] >= ipNum {

			return int(_ipData[lowIdx+2]), nil
		}

		err = ErrNotFound

		return
	}

	// three in a group: start, end, locationIdx
	middleIdx := (lowIdx/ipDataGroupItems + (highIdx/ipDataGroupItems-lowIdx/ipDataGroupItems)/2) * ipDataGroupItems

	if _ipData[middleIdx] > ipNum {

		return search(ipNum, lowIdx, middleIdx)
	} else if _ipData[middleIdx+1] < ipNum {

		return search(ipNum, middleIdx+ipDataGroupItems, highIdx)
	}

	return search(ipNum, middleIdx, middleIdx)
}
