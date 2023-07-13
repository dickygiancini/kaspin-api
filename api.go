package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/gorilla/mux"
)

// var iMids string = "IONPAYTEST"
// var merchantKeys string = "d05ba94a6f9a1411e5e45f429be9b09b138c4c81a67ceee363bea73d5f91c55b"

var iMids string = goDotEnvVariable("IMID_KEY")
var merchantKeys string = goDotEnvVariable("MERCHANT_KEY")
var timeStamp string = time.Now().Format("20060102150405")
var refNo string = "ord" + timeStamp

var devUrl string = "https://dev.nicepay.co.id/"
var logFile string = "log.json"

func (s *APIServer) Run() {
	router := mux.NewRouter()

	router.HandleFunc("/register", makeHTTPHandleFunc(s.handleRegister))
	router.HandleFunc("/payment", makeHTTPHandleFunc(s.handlePayment))
	router.HandleFunc("/status-inquiry", makeHTTPHandleFunc(s.handleInquiry))

	log.Println("JSON API server running on port ", s.listenAddr)
	http.ListenAndServe(s.listenAddr, router)
}

// Register API
func (s *APIServer) handleRegister(w http.ResponseWriter, r *http.Request) error {

	// var amount string = "2000"

	// Randomize amount berdasarkan length, ambil dari utils
	amt := generateRandomNumber(4)

	// Generate Token (utils.go)
	merchantToken := generateSHA256Token(timeStamp + iMids + refNo + amt + merchantKeys)

	// Looks good man! pepega
	// log.Println("timestamps : " + timeStamp)
	// log.Println("iMids : " + iMids)
	// log.Println("refNo : " + refNo)
	// log.Println("amount : " + amount)
	// log.Println("merchantKeys : " + merchantKeys)
	// log.Println("merchantToken : " + merchantToken)

	// Builder account
	account := NewAccount(iMids, merchantToken, timeStamp, refNo, amt)
	jsonPayload, err := json.Marshal(account)
	if err != nil {
		return err
	}

	// Simpan request ke log.json
	err = saveLogFile(logFile, "Request Register", string(jsonPayload))
	if err != nil {
		log.Println("Error saving request to log file:", err)
	}

	url := devUrl + "/nicepay/direct/v2/registration"
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonPayload))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Cek respon status
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("request error status %s : ", resp.Status)
	}

	var responsePayload Credentials
	err = json.NewDecoder(resp.Body).Decode(&responsePayload)
	if err != nil {
		return err
	}

	// Sementara simpan credentials response setelah register
	s.setCredentials(&responsePayload)

	responseJSON, err := json.Marshal(responsePayload)
	if err != nil {
		log.Println("Error encoding response JSON: ", err)
	} else {
		err = saveLogFile(logFile, "Response Register", string(responseJSON))
		if err != nil {
			log.Println("Error saving response to log file:", err)
		}
	}

	// log.Println("Credentials:", fmt.Sprintf("%+v", responsePayload))
	response := WriteJSON(w, http.StatusOK, responsePayload)

	return response
}

// StatusInquiry
func (s *APIServer) handlePayment(w http.ResponseWriter, r *http.Request) error {
	// account := NewAccount(iMids, merchantKeys)
	// request := WriteJSON(w, http.StatusOK, account)
	// currentTime := time.Now()
	// timestamp := currentTime.Format("20060102150405")

	if s.credentials == nil {
		return fmt.Errorf("credentials not found")
	}

	// Randomize amount berdasarkan length, ambil dari utils
	amt := generateRandomNumber(4)

	// Generate Token (utils.go)
	merchantToken := generateSHA256Token(timeStamp + iMids + refNo + amt + merchantKeys)

	queryParams := url.Values{}
	queryParams.Set("timeStamp", timeStamp)
	queryParams.Set("tXid", s.credentials.TXid)
	queryParams.Set("cardNo", generateRandomNumber(16))
	queryParams.Set("cardExpYymm", generate5Years())
	queryParams.Set("cardCvv", "123")
	queryParams.Set("cardHolderNm", "DICKY GIANCINI")
	queryParams.Set("recurringToken", "")
	queryParams.Set("preauthToken", "")
	queryParams.Set("clickPayNo", "")
	queryParams.Set("dataField3", "")
	queryParams.Set("clickPayToken", "")
	queryParams.Set("callBackUrl", "")
	queryParams.Set("merchantToken", merchantToken)

	// Simpan request ke log.json
	logData := queryParams.Encode()
	err := saveLogFile(logFile, "Request Inquiry", devUrl+"nicepay/direct/v2/payment?"+logData)
	if err != nil {
		log.Println("Error saving request to log file:", err)
	}

	url := devUrl + "nicepay/direct/v2/payment?" + queryParams.Encode()
	resp, err := http.Post(url, "application/x-www-form-urlencoded", nil)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Cek respon status
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("request error status %s : ", resp.Status)
	}

	// Karena responsenya berupa HTML
	if strings.HasPrefix(resp.Header.Get("Content-Type"), "text/html") {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return err
		}

		// log.Println(body)

		formValues, err := extractForm(string(body))
		if err != nil {
			return err
		}

		redirectURL := formValues.Get("callbackUrl")
		redirectURL = strings.TrimPrefix(redirectURL, "/")
		redirectURL = devUrl + redirectURL
		formValues.Del("callbackUrl")

		req, err := http.NewRequest("POST", redirectURL, strings.NewReader(formValues.Encode()))
		if err != nil {
			return err
		}
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

		// Submit the new request
		client := &http.Client{}
		resp, err = client.Do(req)
		if err != nil {
			return err
		}
		defer resp.Body.Close()

		// Read the response body
		responseBody, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return err
		}

		var payment Payment
		err = json.Unmarshal(responseBody, &payment)
		if err != nil {
			return err
		}

		s.setPayment(&payment)

		// Simpan response ke log.json
		responseJSON, err := json.Marshal(payment)
		if err != nil {
			return err
		}
		err = saveLogFile(logFile, "Response Payment", string(responseJSON))
		if err != nil {
			log.Println("Error saving response to log file:", err)
		}
	}

	// var responsePayload interface{}
	// err = json.NewDecoder(resp.Body).Decode(&responsePayload)
	// if err != nil {
	// 	return err
	// }

	// responseJSON, err := json.Marshal(responsePayload)
	// if err != nil {
	// 	log.Println("Error encoding response JSON: ", err)
	// } else {
	// 	err = saveLogFile(logFile, "Response Inquiry", string(responseJSON))
	// 	if err != nil {
	// 		log.Println("Error saving response to log file:", err)
	// 	}
	// }

	// response := WriteJSON(w, http.StatusOK, responsePayload)

	return nil
}

func (s *APIServer) handleInquiry(w http.ResponseWriter, r *http.Request) error {
	// Randomize amount berdasarkan length, ambil dari utils
	amt := generateRandomNumber(4)

	referenceNo := s.credentials.ReferenceNo
	tXid := s.credentials.TXid

	if s.credentials != nil {
		amt = s.credentials.Amt
	}

	if s.payment != nil {
		referenceNo = s.payment.ReferenceNo
		tXid = s.payment.TXid
	}

	// Generate Token (utils.go)
	merchantToken := generateSHA256Token(timeStamp + iMids + refNo + amt + merchantKeys)

	inquiry := NewStatusInquiry(iMids, merchantToken, timeStamp, tXid, referenceNo, amt)
	jsonPayload, err := json.Marshal(inquiry)
	if err != nil {
		return err
	}

	// Simpan request ke log.json
	err = saveLogFile(logFile, "Request Inquiry", string(jsonPayload))
	if err != nil {
		log.Println("Error saving request to log file:", err)
	}

	url := devUrl + "/nicepay/direct/v2/inquiry"
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonPayload))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Cek respon status
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("request error status %s : ", resp.Status)
	}

	var responsePayload interface{}
	err = json.NewDecoder(resp.Body).Decode(&responsePayload)
	if err != nil {
		return err
	}

	responseJSON, err := json.Marshal(responsePayload)
	if err != nil {
		log.Println("Error encoding response JSON: ", err)
	} else {
		err = saveLogFile(logFile, "Response Inquiry", string(responseJSON))
		if err != nil {
			log.Println("Error saving response to log file:", err)
		}
	}

	// log.Println("Credentials:", fmt.Sprintf("%+v", responsePayload))
	response := WriteJSON(w, http.StatusOK, responsePayload)

	return response
}

func (s *APIServer) setCredentials(credentials *Credentials) {
	s.credentials = credentials
}

func (s *APIServer) setPayment(payment *Payment) {
	s.payment = payment
}

// Save File
func saveLogFile(filename, section, data string) error {
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	logging := fmt.Sprintf("[%s]\n%s\n\n", section, data)

	_, err = file.WriteString(logging)
	if err != nil {
		return err
	}

	return nil
}

func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)

	return json.NewEncoder(w).Encode(v)
}

type apiFunc func(http.ResponseWriter, *http.Request) error

type ApiError struct {
	Error string
}

func makeHTTPHandleFunc(f apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			// Handle Error
			WriteJSON(w, http.StatusBadRequest, ApiError{Error: err.Error()})
		}
	}
}

type APIServer struct {
	listenAddr  string
	credentials *Credentials
	payment     *Payment
}

func NewAPIServer(listenAddr string) *APIServer {
	return &APIServer{
		listenAddr:  listenAddr,
		credentials: nil,
		payment:     nil,
	}
}
