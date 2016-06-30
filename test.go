package main

import (
  "io"
  "fmt"
  "time"
  "bufio"
  "net/http"
  "regexp"
  "bytes"
  "strings"
  "strconv"
  "compress/gzip"
)

 
var baseurl string = "http://duproprio.com"
var baseuri string = "/search/?hash=/g-re=6/s-pmin=200000/s-pmax=500000/s-bmin=2/s-build=1/s-parent=1/s-filter=forsale/s-days=0/m-pack=house-condo-multiplex/p-con=main/p-ord=date/p-dir=DESC/pa-ge="

var testUrls = []string {
  "http://duproprio.com/search/?hash=/g-re=6/s-pmin=200000/s-pmax=500000/s-bmin=2/s-build=1/s-parent=1/s-filter=forsale/s-days=0/s-gfpack=1/s-hide-sold=/s-mode=list/m-pack=house-condo-multiplex/m-ty=1-15-13-11-2-4-17-21-9-7-20-10-19-6-5-18-100-97-98-99-3-12-14-52-53-55-56-54/p-con=main/p-ord=date/p-dir=DESC/pa-ge=1/",
  "http://duproprio.com/search/?hash=/g-re=6/s-pmin=200000/s-pmax=500000/s-bmin=2/s-build=1/s-parent=1/s-filter=forsale/s-days=0/s-gfpack=1/s-hide-sold=/s-mode=list/m-pack=house-condo-multiplex/m-ty=1-15-13-11-2-4-17-21-9-7-20-10-19-6-5-18-100-97-98-99-3-12-14-52-53-55-56-54/p-con=main/p-ord=date/p-dir=DESC/pa-ge=2/",
  "http://duproprio.com/search/?hash=/g-re=6/s-pmin=200000/s-pmax=500000/s-bmin=2/s-build=1/s-parent=1/s-filter=forsale/s-days=0/s-gfpack=1/s-hide-sold=/s-mode=list/m-pack=house-condo-multiplex/m-ty=1-15-13-11-2-4-17-21-9-7-20-10-19-6-5-18-100-97-98-99-3-12-14-52-53-55-56-54/p-con=main/p-ord=date/p-dir=DESC/pa-ge=3/",
  "http://duproprio.com/search/?hash=/g-re=6/s-pmin=200000/s-pmax=500000/s-bmin=2/s-build=1/s-parent=1/s-filter=forsale/s-days=0/s-gfpack=1/s-hide-sold=/s-mode=list/m-pack=house-condo-multiplex/m-ty=1-15-13-11-2-4-17-21-9-7-20-10-19-6-5-18-100-97-98-99-3-12-14-52-53-55-56-54/p-con=main/p-ord=date/p-dir=DESC/pa-ge=4/",
  "http://duproprio.com/search/?hash=/g-re=6/s-pmin=200000/s-pmax=500000/s-bmin=2/s-build=1/s-parent=1/s-filter=forsale/s-days=0/s-gfpack=1/s-hide-sold=/s-mode=list/m-pack=house-condo-multiplex/m-ty=1-15-13-11-2-4-17-21-9-7-20-10-19-6-5-18-100-97-98-99-3-12-14-52-53-55-56-54/p-con=main/p-ord=date/p-dir=DESC/pa-ge=5/",
  "http://duproprio.com/search/?hash=/g-re=6/s-pmin=200000/s-pmax=500000/s-bmin=2/s-build=1/s-parent=1/s-filter=forsale/s-days=0/s-gfpack=1/s-hide-sold=/s-mode=list/m-pack=house-condo-multiplex/m-ty=1-15-13-11-2-4-17-21-9-7-20-10-19-6-5-18-100-97-98-99-3-12-14-52-53-55-56-54/p-con=main/p-ord=date/p-dir=DESC/pa-ge=6/",
  "http://duproprio.com/search/?hash=/g-re=6/s-pmin=200000/s-pmax=500000/s-bmin=2/s-build=1/s-parent=1/s-filter=forsale/s-days=0/s-gfpack=1/s-hide-sold=/s-mode=list/m-pack=house-condo-multiplex/m-ty=1-15-13-11-2-4-17-21-9-7-20-10-19-6-5-18-100-97-98-99-3-12-14-52-53-55-56-54/p-con=main/p-ord=date/p-dir=DESC/pa-ge=7/",
  "http://duproprio.com/search/?hash=/g-re=6/s-pmin=200000/s-pmax=500000/s-bmin=2/s-build=1/s-parent=1/s-filter=forsale/s-days=0/s-gfpack=1/s-hide-sold=/s-mode=list/m-pack=house-condo-multiplex/m-ty=1-15-13-11-2-4-17-21-9-7-20-10-19-6-5-18-100-97-98-99-3-12-14-52-53-55-56-54/p-con=main/p-ord=date/p-dir=DESC/pa-ge=8/",

}

type HttpResponse struct {
  url      string
  response *http.Response
  err      error
}

func getPageUrls() []string {
  var buf bytes.Buffer
  var pageUrls []string
  buf.WriteString(baseurl)
  buf.WriteString(baseuri)
  buf.WriteString("1/")
  url := buf.String()
  buf.Reset()
  
  client := http.Client{}
  resp, err := client.Get(url)

  re := regexp.MustCompile("search.+pa-ge=[0-9]{1,3}\" rel")
  
  if err == nil && resp != nil && resp.StatusCode == http.StatusOK {
    reader := bufio.NewReader(resp.Body)
    scanner := bufio.NewScanner(reader)
    scanner.Split(bufio.ScanLines)
    for scanner.Scan() {
      switch {
      case re.FindString(scanner.Text()) != "":
        buf.WriteString(baseurl)
        buf.WriteString("/")
        buf.WriteString(strings.Split(re.FindString(scanner.Text()), "\"")[0])
        buf.WriteString("/")
        pageUrls = append(pageUrls, buf.String())
        buf.Reset()
      }
    }
    resp.Body.Close()
  }
  return pageUrls
}

// func getSearchResultsCount() int {
//   re_results := regexp.MustCompile("<strong id=\"searchResultsCounter\">[0-9]\\w+</strong>")
//   client := http.Client{}
//   request, err := http.NewRequest("GET", url, nil)
//   request.Header.Add("Accept-Encoding", "gzip, deflate")
//   resp, err := client.Do(request)
  
// }

func getPropertyUrls(urls []string) []string {
  var propertyUrls []string
  var propertyCount int = 0
  var reader io.ReadCloser

  ch := make(chan string)
  client := http.Client{}
  re_url := regexp.MustCompile("\"\\/.+\" title.+showimage")
  re_results := regexp.MustCompile("<strong id=\"searchResultsCounter\">[0-9]\\w+</strong>")

  for _, url := range urls {
    go func(url string) {
      request, err := http.NewRequest("GET", url, nil)
      request.Header.Add("Accept-Encoding", "gzip, deflate")
      resp, err := client.Do(request)
      
      if err == nil && resp != nil && resp.StatusCode == http.StatusOK {
        // fmt.Println(resp.Header.Get("Content-Encoding")) // DEBUG
        switch resp.Header.Get("Content-Encoding") {
        case "gzip":
          reader, err = gzip.NewReader(resp.Body)
          defer reader.Close()
        default:
          reader = resp.Body
        }
        if err == nil {
            scanner := bufio.NewScanner(reader)
            scanner.Split(bufio.ScanLines)
            for scanner.Scan() {
              switch {
              case re_results.FindString(scanner.Text()) != "":
                if propertyCount == 0 {
                  re := regexp.MustCompile("[0-9]\\w+")
                  propertyCount, err = strconv.Atoi(re.FindString(re_results.FindString(scanner.Text())))
                }
              case re_url.FindString(scanner.Text()) != "":
                var buf bytes.Buffer
                buf.WriteString(baseurl)
                buf.WriteString(strings.Split(scanner.Text(), "\"")[1])
                ch <- buf.String()
                // ch <- url // DEBUG
                // fmt.Println(buf.String()) // DEBUG
              }
            }
          }
          // close(ch)
          resp.Body.Close()
        }
    }(url)
  }
  // FIRST APPROACH:
  fmt.Printf("Getting property URLs")
  for {
    select {
    case r := <-ch:
        propertyUrls = append(propertyUrls, r)
        // TODO: Fix Limitation, does not take into account changing number of items per page
        // Also, may fail if issue with regex
        // if len(propertyUrls) == propertyCount - 1 {
        if len(propertyUrls) == 1540 {
          return propertyUrls
        }
    case <-time.After(50 * time.Millisecond):
        fmt.Printf(".")
    }
  }
  // SECOND APPROACH
  // for r := range ch {
  //   fmt.Printf("Got %s\n", r)
  //     // propertyUrls = append(propertyUrls, r)
  // }
  fmt.Printf("%d properties found", propertyCount)
  return propertyUrls
}

func main() {
  pageUrls := getPageUrls()
  // for _,url := range getPageUrls() {
  //    fmt.Println(url)
  //  }
  for _,url := range getPropertyUrls(pageUrls) {
    fmt.Println(url)
  }
  // for _,url := range getPropertyUrls(testUrls) {
  //   fmt.Println(url)
  // }
}