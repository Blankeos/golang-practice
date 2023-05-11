/*
THIS IS A WORK IN PROGRESS
I'M STILL LEARNING HOW TO USE CHROMEDP
*/
package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/chromedp/chromedp"
	"github.com/fatih/color"
)

func chromeDPStart() {
	// Create context
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	// Run tasks
	chromedp.Run(ctx, chromedp.Navigate("https://www.google.com"))

	// opts := append(
	// 	// select all the elements after the third element
	// 	chromedp.DefaultExecAllocatorOptions[3:],
	// 	chromedp.NoFirstRun,
	// 	chromedp.NoDefaultBrowserCheck,
	// )
	// // create chromedp's context
	// parentCtx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	// defer cancel()

	// ctx, cancel := chromedp.NewContext(parentCtx)
	// defer cancel()

	// if err := chromedp.Run(ctx, chromedp.Navigate("https://www.google.com")); err != nil {
	// 	panic(err)
	// }

	// fmt.Println("I've just saw Google!!!")
}

// function to format JSON data
func formatJSON(data []byte) string {
	var out bytes.Buffer
	err := json.Indent(&out, data, "", " ")

	if err != nil {
		fmt.Println(err)
	}

	d := out.Bytes()
	return string(d)
}

func startServer() {
	fmt.Printf("Starting %s...\n", color.YellowString("Aternos Server"))
	apiUrl := "https://aternos.org/ajax/server/start"

	// 1. Create new http request
	request, err := http.NewRequest("GET", apiUrl, nil)

	// 2. Add query strings
	query := request.URL.Query()
	query.Add("headstart", "0")
	query.Add("access-credits", "0")
	query.Add("SEC", "z3lix46d3nq00000%3Amcvtefyx3r000000")
	query.Add("TOKEN", "6atvSw2Ve0sw9uAQ4VRk")
	request.URL.RawQuery = query.Encode()

	// 3. Set Cookies
	request.Header.Set("Cookie", "ATERNOS_SEC_1dtpdmqiafc00000=c95cohi8xxj00000; ATERNOS_SEC_83ymdwi5zhf00000=wkc9fukge7800000; ATERNOS_SEC_zjwyimedef900000=8nld2c611q900000; ATERNOS_SEC_elhwo6yqhx000000=j6tlt0o2zqm00000; ATERNOS_SEC_2v1xu6b86cs00000=1kw74rdl9ky00000; ATERNOS_SEC_z3lix46d3nq00000=mcvtefyx3r000000; ATERNOS_SEC_76bcjwvbxlq00000=zhenwxk3vv000000; ATERNOS_LANGUAGE=en; _ga=GA1.2.340816199.1686579966; _gid=GA1.2.1960119343.1686579966; ATERNOS_GA=340816199.1686579966; ATERNOS_SESSION=GbjsRlHKHokIn2Kqmv0ithE8LiWhC2HaJ6do27UvBzkpJv230ZgMYxUGLw9Rqf7fxftM51HZbbq3V5GFSrOJf2CLGuWF8KfrGI1O; _pbjs_userid_consent_data=3524755945110770; __qca=P0-232787899-1686579990870; _cc_id=a6e6f07d515b8210ae03ce3e7fb56a28; panoramaId_expiry=1686666393177; panoramaId=25bf152e1ee3db6ccfba81b391e2a9fb927a87ad2455311cd1a19abd32a422b8; panoramaIdType=panoDevice; _au_1d=AU1D-0100-001686579993-I21PGVAS-770F; _lr_env_src_ats=false; _pubcid=e78b8aa9-2d0e-40d9-b454-ea489555d40d; __cf_bm=9jrlv007zdlnRtfYL2Pvecb11fiiPIUfdE8hWSc_c9Y-1686650035-0-AfjaZq4X31upRzHaXNrzdyEUiof/GOJshH3vH9YHBB3IERRwuBIBwU213kRgU4+s8Q==; ATERNOS_SERVER=1M3D4qM3LUtucvJD; _au_last_seen_pixels=eyJhcG4iOjE2ODY2NTAwMzksInR0ZCI6MTY4NjY1MDAzOSwicHViIjoxNjg2NjUwMDM5LCJydWIiOjE2ODY2NTAwMzksInRhcGFkIjoxNjg2NjUwMDM5LCJhZHgiOjE2ODY2NTAwMzksImdvbyI6MTY4NjY1MDAzOSwic21hcnQiOjE2ODY2NTAwMzksInBwbnQiOjE2ODY2NTAwNDIsIm1lZGlhbWF0aCI6MTY4NjY1MDA0Miwib3BlbngiOjE2ODY2NTAwNDIsInVucnVseSI6MTY4NjY1MDAzOSwiYmVlcyI6MTY4NjY1MDA0MiwiaW1wciI6MTY4NjY1MDA0Miwic29uIjoxNjg2NjUwMDM5LCJhZG8iOjE2ODY2NTAwNDIsInRhYm9vbGEiOjE2ODY2NTAwNDJ9; _lr_retry_request=true; __gads=ID=c30825366bc947b8:T=1686579992:RT=1686650724:S=ALNI_MZ71v8CpFEtLD3ZigNMBshPdwax7Q; __gpi=UID=00000c1281dc4d26:T=1686579992:RT=1686650724:S=ALNI_MaT4rdGsQrWzalmh47zM7vHwgc98Q; FCNEC=%5B%5B%22AKsRol_y8Ai0b6GA3ZmX37UXkbZFSb6rogzw4DRhKuY_Sif6nnTRslBWJilPvcQ3VgWlNcGTU-Hc577BRypiN9-gFJGXaBojTwfGbmI2XJ3UTxpaBFiV_Ulg9Kc4JetSu0rueAH7ayUzxg4p3m06PIs8fXqRSR-JlQ%3D%3D%22%5D%2Cnull%2C%5B%5D%5D; cto_bundle=wEVf5V85OVgxMmV1U25LQkM3ZEtvNHMzYk1GZjNwbWlUREclMkZXcU5XUEYlMkJFcXk4bTBuSm9rcm5YVUM2SFJkendENWlVWFMlMkIxOTl1WmhRdjZrdlBCaDNUYXZjOWd5VGQ5c05SRFJxcyUyRkExRE9RMVVUazRzWDJkTkpsJTJGT1RMWnB3S0RsZ2phQUxaWGM1MXclMkJVYzhRMklVTUliMEElM0QlM0Q; cto_bidid=7h6yFF9Rd2pOT2pSVmhRZHJWJTJGc2s1N0RSdFA4c213NUVKa1I1cmdWRmdKOGg5anB4ZEZJTUNlS2pKZ25IUjdYWEIzJTJCc1UySHViYWVOaEM5MkoyR3BQMjVZcmZvR05iTTR5QXg1bnlDMmdNVTdUejglM0Q")

	if err != nil {
		fmt.Printf("%s\n", color.RedString("Error when creating new Request"))
	}

	// 4. Send the request
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		fmt.Printf("%s\n", color.RedString("Error found when sending request"))
	}

	// 5. Print the response
	fmt.Printf("Response Status: %s\n", response.Status)
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("%s\n", color.RedString("Error found when parsssing responseBody"))
	}

	fmt.Printf("Response Body: %s\n", string(body))

	// responseBody, err := io.ReadAll(response.Body)

	// if err != nil {
	// 	fmt.Printf("%s\n", color.RedString("Error found when parsssing responseBody"))
	// }

	// formattedData := formatJSON(responseBody)
	// fmt.Println("Status: ", response.Status)
	// fmt.Println("Response Body: ", formattedData)

	// defer response.Body.Close()
}

func main() {
	chromeDPStart()
	// startServer()
}
