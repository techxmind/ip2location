```
import (
    "fmt"

    "github.com/techxmind/ip2location"
)

ip := "123.15.145.134"
if loc, err := ip2location.Get(ip); err != nil {
    fmt.Print(err)
} else {
    fmt.Printf(
        "country: %s\nprovince: %s\ncity: %s\ngeo id:%s\nregion id:%s\n",
        loc.Country,
        loc.Province,
        loc.City,
        loc.GeoID,
        loc.ChinaRegionID,
    )
}
```
