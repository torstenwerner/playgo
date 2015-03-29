package main

import (
    "flag"
    "fmt"
    "os"
    "code.google.com/p/goauth2/oauth"
    "github.com/digitalocean/godo"
)

func main() {
    listFlag := flag.Bool("list", false, "list all droplets")
    createFlag := flag.Bool("create", false, "create a new droplet")
    deleteFlag := flag.Int("delete", -1, "delete a droplet by id")
    imagesFlag := flag.Bool("images", false, "list all images")
    keysFlag := flag.Bool("keys", false, "list all keys")
    flag.Parse()

    client := getClient()

    if (*listFlag) {
        allDroplets, err := listDroplets(client)
        if err != nil {
            panic(err)
        }
        fmt.Printf("droplet ids = %v\n", getDropletIds(allDroplets))
    }
    if (*createFlag) {
        newDroplet, err := createDroplet(client)
        if err != nil {
            panic(err)
        }
        fmt.Printf("new Droplet id = %d\n", newDroplet.Droplet.ID)
    }
    if (*deleteFlag > 0) {
        deleteDroplet(client, *deleteFlag)
    }
    if (*imagesFlag) {
        allImages, err := listImages(client)
        if err != nil {
            panic(err)
        }
        fmt.Printf("images = %v\n", allImages)
    }
    if (*keysFlag) {
        allKeys, err := listKeys(client)
        if err != nil {
            panic(err)
        }
        fmt.Printf("images = %v\n", convertKeys(allKeys))
    }
}

func getClient() *godo.Client {
    token := os.Getenv("DO_TOKEN")
    if (token == "") {
        panic("environment variable DO_TOKEN not set")
    }

    t := &oauth.Transport{
        Token: &oauth.Token{AccessToken: token},
    }

    return godo.NewClient(t.Client())
}

func listDroplets(client *godo.Client) ([]godo.Droplet, error) {
    allDroplets, _, err := client.Droplets.List(&godo.ListOptions{
        Page: 0,
        PerPage: 10,
    })
    return allDroplets, err
}

func getDropletIds(droplets []godo.Droplet) ([]int) {
    ids := make([]int, len(droplets))
    for i, droplet := range droplets {
        ids[i] = droplet.ID
    }
    return ids
}

func createDroplet(client *godo.Client) (*godo.DropletRoot, error) {
    sshKeys := make([]godo.DropletCreateSSHKey, 1)
    sshKeys[0] = godo.DropletCreateSSHKey{
        ID: 740966,
    }
    createRequest := &godo.DropletCreateRequest{
        Name:   "squid",
        Region: "lon1",
        Size:   "512mb",
        Image: godo.DropletCreateImage{
            ID: 11220117,
            //ssl: ID: 11218733,
        },
        SSHKeys: sshKeys,
    }

    newDroplet, _, err := client.Droplets.Create(createRequest)
    return newDroplet, err
}

func deleteDroplet(client *godo.Client, id int) {
    client.Droplets.Delete(id)
}

func listImages(client *godo.Client) ([]godo.Image, error) {
    allImages, _, err := client.Images.List(&godo.ListOptions{
        Page: 0,
        PerPage: 10,
    })
    return allImages, err
}

func listKeys(client *godo.Client) ([]godo.Key, error) {
    allKeys, _, err := client.Keys.List(&godo.ListOptions{
        Page: 0,
        PerPage: 10,
    })
    return allKeys, err
}

type keyType struct {
    ID int
    name string
}

func (k keyType) String() (string) {
    return fmt.Sprintf("{id: %d, name: \"%s\"}", k.ID, k.name)
}

func convertKeys(keys []godo.Key) ([]keyType) {
    k := make([]keyType, len(keys))
    for i, key := range keys {
        k[i] = keyType{
            ID: key.ID,
            name: key.Name,
        }
    }
    return k
}
