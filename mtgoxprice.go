package main

import (
   "code.google.com/p/go.net/html" 
   "net/http"
   "log"
   "fmt"
)

func main() {
    resp, err := http.Get("https://www.mtgox.com")

    if err != nil {
        log.Fatal(err)
    }

    doc, err := html.Parse(resp.Body)

    if err != nil {
        log.Fatal(err)
    }

    fmt.Println(getWeightedAverage(doc))
}

func getWeightedAverage(n *html.Node) (price string) {
    price = ""

    if getAttribute(n, "id") == "lastPrice" {
        span := findFirstByTagName(n, "span")
        price = getText(span)
        if price != "" {
            return
        }
    }

    for c := n.FirstChild; c != nil; c = c.NextSibling {
        if price = getWeightedAverage(c); price != "" {
            return
        }
    }

    return
}

func findFirstByTagName(root *html.Node, tagName string) *html.Node {
    if root.Type == html.ElementNode && root.Data == tagName {
        return root
    }

    for c := root.FirstChild; c != nil; c = c.NextSibling {
        if foundNode := findFirstByTagName(c, tagName); foundNode != nil {
            return foundNode
        }
    }

    return nil
}

func getText(root *html.Node) string {

    if root.FirstChild != nil && root.FirstChild.Type == html.TextNode {
        return root.FirstChild.Data
    }

    return "" 
}

func getAttribute(n *html.Node, attr string) (value string) {
    for _, val := range(n.Attr) {
        if (val.Key == attr) {
            value = val.Val
            return
        }
    }

    return
}
