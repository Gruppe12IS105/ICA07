# ICA07

## Ukryptert    
Hensikten med denne mappen er å samle både den ukrypterte TCP og UDP server/klient relasjonen.    
For å få det til å fungere starter man først serveren, og deretter kjører man klienten slik at meldingen blir sendt over.      

## Kryptert
Essensen med denne mappen er det samme som Ukryptert, bare at meldingen i mellom skal bli kryptert.

## TCP client
Kjør to vinduer i terminal, en med "go run tcp-server.go", og en med "go run tcp-client.go"
Skriv tekst i clienten, så skal det vises på server vinduet.
