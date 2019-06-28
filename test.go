package main

import (
  "fmt"
  "net/http"
  "time"
  "bufio"
  "compress/gzip"
  "strings"
  "io"
  "regexp"
  "strconv"
  "bytes"
)

var baseurl string = "http://url.com"
var baseuri string = "/search/?hash=/g-re=6/s-pmin=200000/s-pmax=500000/s-bmin=2/s-build=1/s-parent=1/s-filter=forsale/s-days=0/m-pack=house-condo-multiplex/p-con=main/p-ord=date/p-dir=DESC/pa-ge="

var urls = []string{
  "http://url.com/condo-a-vendre-rosemont-petite-patrie-quebec-700071",
 
}

type HttpResponse struct {
  url      string
  response *http.Response
  err      error
}

type PropertyData struct {
  living_space    string
  address         string
  neighborhood    string
  asking_price    string
  n_rooms         string
  n_floors        string
  on_floor        string
  rights          string
  ptype           string
}

func asyncHttpGets(urls []string) []*HttpResponse {
  ch := make(chan *HttpResponse)
  responses := []*HttpResponse{}
  client := http.Client{}
  for _, url := range urls {
      go func(url string) {
          fmt.Printf("Fetching %s \n", url)
          request, err := http.NewRequest("GET", url, nil)
          request.Header.Add("Accept-Encoding", "gzip, deflate")
          resp, err := client.Do(request)
          ch <- &HttpResponse{url, resp, err}
          if err != nil && resp != nil && resp.StatusCode == http.StatusOK {
              resp.Body.Close()
          }
      }(url)
  }
  for {
      select {
      case r := <-ch:
          fmt.Printf("%s was fetched\n", r.url)
          if r.err != nil {
              fmt.Println("with an error", r.err)
          }
          responses = append(responses, r)
          if len(responses) == len(urls) {
            // fmt.Printf("Got everything!\n")
            return responses
          }
      case <-time.After(50 * time.Millisecond):
          fmt.Printf(".")
      }
  }
  return responses
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


func getPropertyData(results []*HttpResponse) []*PropertyData {
  channel := make(chan *PropertyData)
  p_data := []*PropertyData{}
  for _, result := range results {
    go func(result *HttpResponse){
      address, living_space, neighborhood, asking_price, n_rooms, n_floors, on_floor, rights := extractPropertyInfo(result.response.Body)
      ptype := strings.Split(strings.Split(result.url,"-a-vendre-")[0], "http://url.com/")[1]
      channel <- &PropertyData{living_space, address, neighborhood, asking_price, n_rooms, n_floors, on_floor, rights, ptype }
      result.response.Body.Close()
    }(result)
  }
  // fmt.Printf("Extracting Property Data...\n")
  for {
      select {
      case r := <- channel:
          p_data = append(p_data, r)
          if len(p_data) == len(results) {
              return p_data
          }
      case <-time.After(100 * time.Millisecond):
          // fmt.Printf(".")
      }
  }
  return p_data
}

func extractPropertyInfo(body io.Reader) (string, string, string, string, string, string, string, string) {

  // TODO: 
  // Stats prix: min, max, mediane, moyenne

  address := "address"
  living_space := "living_space"
  neighborhood := "neighborhood"
  asking_price := "asking_price"
  n_rooms := "n_rooms"
  n_floors := "n_floors"
  on_floor := "on_floor"
  rights := "rights"
  reader, err := gzip.NewReader(body)
  scanner := bufio.NewScanner(reader)
  if err == nil {
    scanner.Split(bufio.ScanLines)
    for scanner.Scan() {
      switch {
      case strings.Contains(scanner.Text(), "<title>"):
        fmt.Println(scanner.Text())
      //   if strings.Contains(scanner.Text(), "vendu") {
      //     address = "Vendu"
      //   } else {
      //     address = strings.Split(scanner.Text(),", ")[1]
      //     address = strings.Replace(address,";","_",-1)
      //     address = strings.Replace(address,"&eacute_","e",-1)
      //     address = strings.Replace(address,"&egrave_","e",-1)
      //     address = strings.Replace(address,"&frac12_","1/2",-1)
      //     address = strings.Replace(address,"&#039_","'",-1)
      //     address = strings.Replace(address,"&ocirc_","o",-1)
      //   }
      case strings.Contains(scanner.Text(), "Aire habitable"):
        re := regexp.MustCompile("[0-9]\\s?[0-9]+\\.?[0-9]+.pi")
        living_space = strings.Replace(strings.Split(scanner.Text(), "</strong> ")[1], "</li>", "", -1)
        fmt.Println(living_space)  
        if !strings.Contains(living_space,"x") {
          living_space = strings.Replace(strings.TrimSpace(strings.Replace(re.FindString(living_space),"pi","",-1))," ","",-1)
          living_space = strings.Replace(living_space,".",",",-1)
        }
      case strings.Contains(scanner.Text(), "addressLocality"):
        // neighborhood = strings.Replace(strings.Split(scanner.Text(), " content=\"")[1], "\" />", "", -1)
        neighborhood = "neighborhood" // DEBUG
      case strings.Contains(scanner.Text(), " <li><strong>Prix demand"):
        // asking_price = strings.Replace(strings.Split(scanner.Text(), "</strong> ")[1], "</li>", "", -1)
        // asking_price = strings.Replace(asking_price, " ", "", -1)
        // asking_price = strings.Replace(asking_price, "$", "", -1)
        asking_price = "asking_price" // DEBUG
      case strings.Contains(scanner.Text(), "Nombre de chambres"):
        // n_rooms = strings.Replace(strings.Split(scanner.Text(), "</strong> ")[1], "</li>", "", -1)
        n_rooms = "n_rooms" // DEBUG
      case strings.Contains(scanner.Text(), "tages (s-sol exclu)"):
        // n_floors = strings.Replace(strings.Split(scanner.Text(), "</strong> ")[1], "</li>", "", -1)
        n_floors = "n_rooms" // DEBUG
      case strings.Contains(scanner.Text(), " (si condo) "):
        // on_floor = strings.Replace(strings.Split(scanner.Text(), "</strong> ")[1], "</li>", "", -1)
        on_floor = "on_floor" // DEBUG
      case strings.Contains(scanner.Text(), "  <li><strong>Droit de propri"):
        // rights = strings.Replace(strings.Split(scanner.Text(), "</strong> ")[1], "</li>", "", -1)
        rights = "on_floor" // DEBUG
      }
    }
  }
  return address, living_space, neighborhood, asking_price, n_rooms, n_floors, on_floor, rights
}

func main() {
  // TODO:  1. Get number of pages, 4 lines before the line with keyword "Suivants"
  //        2. Generate page urls
  //        3. Contact page urls to get property urls
  //           => Pattern is "^                    <a href="|grep title|grep itemprop|awk 'BEGIN {FS="\""} {print $2}'
  results := asyncHttpGets(urls)
  propertyData := getPropertyData(results)
  // n_pages := getNumberOfPages()
  fmt.Printf("Number of pages: %d", getNumberOfPages())



  fmt.Printf("%s|%s|%s|%s|%s|%s|%s|%s|%s\n",
      "Quartier",
      "Adresse",
      "Aire Habitable",
      "Prix Demande",
      "Nombre de chambres",
      "Nombre d'etages",
      "Situe sur quel etage",
      "Acces a la propriete",
      "Type de propriete")
  for _, property := range propertyData {
    fmt.Printf("%s|%s|%s|%s|%s|%s|%s|%s|%s\n",
      property.neighborhood,
      property.address,
      property.living_space,
      property.asking_price,
      property.n_rooms,
      property.n_floors,
      property.on_floor,
      property.rights,
      property.ptype)
  // }
}
