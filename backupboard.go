package main

import (
	"github.com/chegaa/pb.v1"
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/pkg/browser"
	"io/ioutil"
	"io"
	"net/http"
	"log"
	"os"
	"strconv"
	"strings"
	"regexp"
	"os/exec"
)

//AedTacvaF1GN2rnSvZ9mF6enRZztFNURu1aeKwdEMbu62SA0OQAAAAA

// pinterest api structs
type PinterestAuthResponse struct {
	Auth string `json:"access_token"`
}
type PinterestBoards struct {
	Boards []PinterestBoard `json:"data"`
	Page   pinterestPage    `json:"page"`
}
type PinterestBoard struct {
	Pins []PinterestPin
	Url  string `json:"url"`
	Id   string `json:"id"`
	Name string `json:"name"`
}
type pinterestBoardPage struct {
	Pins []PinterestPin `json:"data"`
	Page pinterestPage  `json:"page"`
}

func (p *PinterestBoard) AddPins(auth string) {
	url := "https://api.pinterest.com/v1/boards/" + p.Id + "/pins?access_token=" + auth + "&limit=100&fields=image,note"
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	var j pinterestBoardPage
	json.Unmarshal(body, &j)
	pins := j.Pins
	page := j.Page
	for page != (pinterestPage{}) {
		resp, err := http.Get(page.Next)
		if err != nil {
			log.Fatal(err)
		}
		body, err := ioutil.ReadAll(resp.Body)
		var j pinterestBoardPage
		json.Unmarshal(body, &j)
		pins = append(pins, j.Pins...)
		page = j.Page
	}
	p.Pins = pins
}

type PinterestPin struct {
	Image pinterestPinImage `json:"image"`
	Note  string            `json:"note"`
}

func (p PinterestPin) GetSource() string {
	return p.Image.Original.Url
}

type pinterestPinImage struct {
	Original pinterestImage `json:"original"`
	id       string         `json:"id"`
}
type pinterestImage struct {
	Url    string `json:"url"`
	height int    `json:"height"`
	width  int    `json:"width"`
}
type pinterestPage struct {
	Next string `json:"next"`
}
// pinterest api functions
func GetToken() string {
	// app configurations and such
	client_id := "4913911707199357020"
	// not putting this on the internet....
	panic("WOOOOPS. Please use your own application WHEN HACKING.")
	client_secret := ""
	scope := "read_public"
	redirect_uri := "https://calderwhite.github.io/myGitFolio/register"
	code_uri := "https://api.pinterest.com/oauth?client_id=" + client_id + "&redirect_uri=" + redirect_uri + "&scope=" + scope + "&response_type=code&stat=calderwhiteapp"
	// open the web browser for the user, so they can enter the code (so we can get the access token)
	err := browser.OpenURL(code_uri)
	if err != nil {
		log.Fatal(err)
	}
	// the code
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter Code:")
	code, _ := reader.ReadString('\n')
	code = code[0 : len(code)-2]
	// get the auth token from the api
	access_uri := "https://api.pinterest.com/v1/oauth/token?grant_type=authorization_code&client_id=" + client_id + "&client_secret=" + client_secret + "&code=" + code
	req, err := http.NewRequest("POST", access_uri, nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Cookie", "_b=\"ASpC2pAwFwhAqoy+s+QIInAW4eFZKimueWU24XDU70oHglcFvRgoA2PeZTPrc2dDxD0=\"; _auth=1; _pinterest_sess=\"TWc9PSZ4MTVybndxYXVyRG16TE5wVFpkTklKVklOQ2RxZnloN3U3dFNtYmR2bFB1VXhTUENvcDRyd3BuQmhmSWZ6ZkcrbVBBdlc1cFJNNzFENkNNK2tveFBFWjBzNHoxWFNjc2w3Vk1lZ3NpM2N4VGorbWVxaUhpWXY3TzJ2L0JrUGxzVHVZQkN6bkNPWjc4M1Y4eGNOQVNKQWV5QlB5T3BzOHZMYzhLdTZZQmJpVmxMd29Pc01TR04relhlUnBDMk1qdGRGRFBBSmNTUEtneVJWcUlwSHRUcnNrNW8vb29jd29qQnF5VlRDVFRTTkkwZllETENJVkFIK0VDeCs3SXVUb09sV0psQUlHWmwzT0xISm95ZDNOdHhpeEwxamFicjBlbStyekZkTzhDbmppMTQ4eTg1SFV5WGs1bUJ2cFRlS0ZxbEJsNDBITnNMN2I0czkza3RCcVVwek1RZGZORi9ESTJvSCswWlV4MityM1B1UTAyWHNhWnRaK0dTZmlJeWY2T1Blc1lGSjZIQUc2dzUwU3ZZTEkxVkp0TVdGcHI1R080a1hRTVBSayswMzRmaExaNEVrWTFTcXIwdnNtamN4Z1JlUDU0dHZISmQ2d2x0bEllajhncExUOHhyckxBSjJ3MVoyUWdCOXUxbnBWaz0mbDlkdDBTZWVQT1BRZmg0a3FvM3JGUEd3NVRRPQ==\"; _pinterest_pfob=disabled; _ga=GA1.2.1463300846.1501109959; _gid=GA1.2.114243072.1501109959; csrftoken=QXaUKZfjlXulnZx1FRW3bBsg6e6F6Nef")
	req.Header.Set("Origin", "https://api.pinterest.com")
	req.Header.Set("Accept-Encoding", "utf-8")
	req.Header.Set("Accept-Language", "en-US,en;q=0.8")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 6.3; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/59.0.3071.115 Safari/537.36")
	req.Header.Set("Accept", "*/*")
	req.Header.Set("Referer", "https://api.pinterest.com/")
	req.Header.Set("X-Requested-With", "XMLHttpRequest")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Content-Length", "0")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	j := PinterestAuthResponse{}
	json.Unmarshal(body, &j)
	return j.Auth
}
func GetBoards(auth string) []PinterestBoard {
	fmt.Println("Getting boards...")
	url := "https://api.pinterest.com/v1/me/boards/?access_token=" + auth + "&limit=100"
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	if resp.StatusCode != 200 {
		fmt.Println("Error, bad status code: ", resp.StatusCode)
	}
	// getting the boards' meta data
	body, err := ioutil.ReadAll(resp.Body)
	var boards PinterestBoards
	json.Unmarshal(body, &boards)
	bList := boards.Boards
	for boards.Page != (pinterestPage{}) {
		fmt.Println(boards.Page,"ded")
		resp, err := http.Get(boards.Page.Next)
		if err != nil {
			log.Fatal(err)
		}
		body, err = ioutil.ReadAll(resp.Body)
		json.Unmarshal(body, &boards)
		bList = append(bList, boards.Boards...)
	}
	bar := pb.StartNew(len(bList))
	// adding the pins to each board
	for i := 0; i < len(bList); i++ {
		bList[i].AddPins(auth)
		/* CODE TO SAVE WITH LOCAL DATA
	    file, err := os.Create(bList[i].Name+".json")
	    if err != nil {
	        log.Fatal(err)
	    }
	    // Use io.Copy to just dump the response body to the file. This supports huge files
	    m,_ := json.Marshal(bList[i])
	    err = ioutil.WriteFile(bList[i].Name+".json", m,os.FileMode(int(777)))
	    if err != nil {
	        log.Fatal(err)
	    }
	    file.Close()
	    */
		bar.Increment()
	}
	bar.FinishPrint("")
	return bList
}
// saving functions
func pathExists(path string) bool {
    _, err := os.Stat(path)
    if err == nil { return true }
    if os.IsNotExist(err) { return false}
    if err != nil{
    	log.Fatal(err)
    }
    return true
}
func SaveBoard(parentDir string,board PinterestBoard,bar *pb.ProgressBar,done chan bool,iteration int){
	parentDir = parentDir+"/"
	reg := regexp.MustCompile("[^abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789 ]+")
	pinI := 0
	boardName := board.Name
	boardName = reg.ReplaceAllString(boardName,"")
	boardName = strings.Trim(boardName," ")
	if strings.Replace(boardName," ","",-1) == ""{
		boardName = "board" + strconv.Itoa(iteration)
	}
	if !pathExists(parentDir+boardName){
		/// 777 is rwx (7), for ugo
		os.Mkdir(parentDir+boardName,os.FileMode(int(777)))
	}
	for pin:=0;pin<len(board.Pins);pin++{
		source := board.Pins[pin].GetSource()
	    response, e := http.Get(source)
	    if e != nil {
	    	fmt.Println("FATAL!")
	        log.Fatal(e)
	    }

	    defer response.Body.Close()

	    //open a file for writing
	    name := board.Pins[pin].Note
	    name = reg.ReplaceAllString(name,"")
	    if strings.Replace(name," ","",-1) == ""{
	    	name ="pin" +  strconv.Itoa(pinI)
	    }
	    if len(name) > 30{
	    	name = name[0:31]
	    }
	    file, err := os.Create(parentDir+boardName+"/"+name+source[len(source)-4:len(source)])
	    if err != nil {
	    	fmt.Println("FATAL!")
	        log.Fatal(err)
	    }
	    // Use io.Copy to just dump the response body to the file. This supports huge files
	    _, err = io.Copy(file, response.Body)
	    if err != nil {
	    	fmt.Println("FATAL!")
	        log.Fatal(err)
	    }
	    file.Close()
		bar.Increment()
		pinI++
	}
	done <- true
}
func SaveBoards(boards []PinterestBoard){
	if !pathExists("backupBoard"){
		/// 777 is rwx (7), for ugo
		os.Mkdir("backupBoard",os.FileMode(int(777)))
	}
	total := 0
	for i:=0;i<len(boards);i++{
		total+=len(boards[i].Pins)
	}
	bar := pb.StartNew(total)
	// threading let's gooooo
	done:= make(chan bool)
	for board:=0;board<len(boards);board++{
		go SaveBoard("backupBoard",boards[board],bar,done,board)
	}
	for i:=0;i<len(boards);i++{
		<-done
	}
	bar.FinishPrint("")
}
func main() {

	auth := GetToken()
	//auth := "AedTacvaF1GN2rnSvZ9mF6enRZztFNURu1aeKwdEMbu62SA0OQAAAAA"
	boards := GetBoards(auth)
	/* CODE TO WORK WITH LOCAL DATA
	var boards []PinterestBoard
	d,err:= ioutil.ReadDir("savedBoards")
	if err != nil{
		panic(err)
	}
	for i:=0;i<len(d);i++{
		data, err := ioutil.ReadFile("savedBoards/"+d[i].Name())
		if err != nil {
			panic(err)
		}
		var j PinterestBoard
		json.Unmarshal(data, &j)
		boards = append(boards,j)
	}
	*/
	SaveBoards(boards)
	exec.Command("pause")
}
