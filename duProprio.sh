# Get all properties (clicked)
curl -s 'http://duproprio.com/webservice/search/get/' \
 -H'Accept: text/html, */*; q=0.01' \
 -H'Accept-Encoding: gzip, deflate' \
 -H'Accept-Language: en-US,en;q=0.5' \
 -H'Connection: keep-alive' \
 -H'Content-Type: application/x-www-form-urlencoded' \
 -H'Host: duproprio.com' \
 -H'X-Requested-With: XMLHttpRequest' \
 --data 'g-re%5B%5D=6&g-ci%5B%5D=6-1887&g-ci%5B%5D=6-1897&g-ci%5B%5D=6-1889&g-ci%5B%5D=6-1892&m-pack=false&s-pmin=0&s-pmax=500000&m-opts=false&s-days=0&s-bmin=2&s-bamin=0&pa-ge=1&m-ty=false&p-ord=date&p-dir=DESC&s-filter=forsale&s-build=check&s-parent=check&reload-featured-homes=true&s-lotunt=FT&s-lotmin=&s-lotmax=&saveSession=true' | gunzip 2>&1  | grep data-url | awk 'FS="=" {print $2}'

 # Get single property details (clicked)
 curl 'http://duproprio.com/condo-a-vendre-rosemont-petite-patrie-quebec-698776' \
 -H'Accept: text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8' \
 -H'Accept-Encoding: gzip, deflate' \
 -H'Accept-Language: en-US,en;q=0.5' \
 -H'Connection: keep-alive' \
 -H'Cookie: dp_shared_session=n42sf1k68phjjvvbbgslmlq0k7; dp_session2[uuid]=21394214e936c1f10122596ab95085e5f966bb7cf0849b77a7c80d93afb8a88f; _ga=GA1.2.908931465.1464007553; __gads=ID=2ddf10ae685be13c:T=1464007547:S=ALNI_MYH93OEqlo6rutUpRjjXH0C0RGapw; _sp_id.c781=839ca8011581645f.1464007554.1.1464010948.1464007554; _sp_ses.c781=*' \
 -H'Host: duproprio.com' \
 -H'Referer: http://duproprio.com/search/?hash=/g-ci=1887-1897-1889-1892/s-pmin=0/s-pmax=450000/s-bmin=2/s-build=1/s-parent=1/s-filter=forsale/s-days=0/p-con=main/p-ord=date/p-dir=DESC/pa-ge=1/' \
 -H'User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10.11; rv:46.0) Gecko/20100101 Firefox/46.0'

 # Get single property details (scripted)
 curl -s 'http://duproprio.com//condo-a-vendre-cote-des-neiges-quebec-699804' \
 -H'Accept: text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8' \
 -H'Accept-Encoding: gzip, deflate' \
 -H'Accept-Language: en-US,en;q=0.5' \
 -H'Connection: keep-alive' \
 -H'Host: duproprio.com' | gunzip | egrep "Aire habitable \(s-sol exclu\)|Nombre de chambres|Prix demand|*tages \(s-sol exclu|<title>|Situ?"

# Get single property details (scripted)
 curl -s 'http://duproprio.com/maison-a-un-etage-et-demi-a-vendre-ahuntsic-cartierville-quebec-692839' \
 -H'Accept: text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8' \
 -H'Accept-Encoding: gzip, deflate' \
 -H'Accept-Language: en-US,en;q=0.5' \
 -H'Connection: keep-alive' \
 -H'Host: duproprio.com' | gunzip | grep "Situ*tage"

curl -s 'http://duproprio.com/condo-a-vendre-st-leonard-quebec-625331' \
 -H'Accept: text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8' \
 -H'Accept-Encoding: gzip, deflate' \
 -H'Accept-Language: en-US,en;q=0.5' \
 -H'Connection: keep-alive' \
 -H'Host: duproprio.com' | gunzip | egrep "Aire habitable \(s-sol exclu\)|Nombre de chambres|Prix demand|*tages \(s-sol exclu|<title>|Situ*tage"

 # Aire habitable (s-sol exclu)
 # Adresse
 # Nombre de chambres :
 # Prix demandé :
 # Nombre d'étages (s-sol exclu) :
 # Situé à quel étage? (si condo) :

 egrep "Aire habitable \(s-sol exclu\)|Nombre de chambres|Prix demand|*tages \(s-sol exclu|<title>"

"http://duproprio.com/condo-a-vendre-st-leonard-quebec-625331",
"http://duproprio.com/condo-a-vendre-montreal-sud-ouest-quebec-696084",
"http://duproprio.com/maison-en-rangee-de-ville-a-vendre-ile-des-soeurs-quebec-695867",
"http://duproprio.com/condo-a-vendre-griffintown-quebec-689905",
"http://duproprio.com/condo-a-vendre-montreal-centre-ville-ville-marie-quebec-694600",
"http://duproprio.com/maison-a-vendre-montreal-centre-ville-ville-marie-quebec-700149",
"http://duproprio.com/maison-en-rangee-de-ville-a-vendre-rivieres-des-prairies-quebec-694269",
"http://duproprio.com/maison-a-vendre-beaconsfield-quebec-615947",
"http://duproprio.com/condo-a-vendre-montreal-sud-ouest-quebec-636666",
"http://duproprio.com/jumele-a-vendre-lachine-quebec-696319"