package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {
    // Login
    token, err := login("http://localhost:8081/login", "dany", "dany13")
    if err != nil {
        log.Fatalf("Login error: %v", err)
    }
    fmt.Println("Token:", token)

    // Get all Bajus
    bajus, err := getAllBajus("http://localhost:8081/baju", token)
    if err != nil {
        log.Fatalf("Error getting bajus: %v", err)
    }
    fmt.Println("Bajus:", bajus)

    // Create a new Baju
    newBaju := Baju{Name: "Baju Baru", Size: "M", Price: 100000}
    err = createBaju("http://localhost:8081/baju", token, newBaju)
    if err != nil {
        log.Fatalf("Error creating baju: %v", err)
    }
    fmt.Println("Baju berhasil dibuat")

    // Get all Bajus again to see the new one
    bajus, err = getAllBajus("http://localhost:8081/baju", token)
    if err != nil {
        log.Fatalf("Error getting bajus: %v", err)
    }
    fmt.Println("Bajus:", bajus)
}

func login(url, username, password string) (string, error) {
    payload := map[string]string{
        "username": username,
        "password": password,
    }
    payloadBytes, err := json.Marshal(payload)
    if err != nil {
        return "", err
    }

    resp, err := http.Post(url, "application/json", bytes.NewBuffer(payloadBytes))
    if err != nil {
        return "", err
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        return "", fmt.Errorf("login gagal dengan status: %d", resp.StatusCode)
    }

    var result map[string]string
    err = json.NewDecoder(resp.Body).Decode(&result)
    if err != nil {
        return "", err
    }

    token, ok := result["token"]
    if !ok {
        return "", fmt.Errorf("tidak ada token dalam respons")
    }

    return token, nil
}

func getAllBajus(url, token string) ([]Baju, error) {
    req, err := http.NewRequest("GET", url, nil)
    if err != nil {
        return nil, err
    }
    req.Header.Set("Authorization", "Bearer "+token)

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        return nil, fmt.Errorf("gagal mendapatkan bajus dengan status: %d", resp.StatusCode)
    }

    var bajus []Baju
    err = json.NewDecoder(resp.Body).Decode(&bajus)
    if err != nil {
        return nil, err
    }

    return bajus, nil
}

func createBaju(url, token string, newBaju Baju) error {
    payloadBytes, err := json.Marshal(newBaju)
    if err != nil {
        return err
    }

    req, err := http.NewRequest("POST", url, bytes.NewBuffer(payloadBytes))
    if err != nil {
        return err
    }
    req.Header.Set("Authorization", "Bearer "+token)
    req.Header.Set("Content-Type", "application/json")

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        return err
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusCreated {
        return fmt.Errorf("gagal membuat baju dengan status: %d", resp.StatusCode)
    }

    return nil
}

type Baju struct {
    ID    int    `json:"id"`
    Name  string `json:"name"`
    Size  string `json:"size"`
    Price int    `json:"price"`
}
