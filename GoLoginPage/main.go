package main


import {
	"net/http"
	"fmt"
}


func Home(w http.ResponseWriter, r *http.Request) {

	//w.Write([]byte(fmt.Sprintf("Generating QR code\n")))

	// generate a random string - preferbly 6 or 8 characters
	randomStr := randStr(6, "alphanum")

   // For Google Authenticator purpose
   // for more details see
   // https://github.com/google/google-authenticator/wiki/Key-Uri-Format
   secret = base32.StdEncoding.EncodeToString([]byte(randomStr))
   //w.Write([]byte(fmt.Sprintf("Secret : %s !\n", secret)))

   // authentication link. Remember to replace SocketLoop with yours.
   // for more details see
   // https://github.com/google/google-authenticator/wiki/Key-Uri-Format
   authLink := "otpauth://totp/SocketLoop?secret=" + secret + "&issuer=SocketLoop"

   // Encode authLink to QR codes
   // qr.H = 65% redundant level
   // see https://godoc.org/code.google.com/p/rsc/qr#Level

   code, err := qr.Encode(authLink, qr.H)

   if err != nil {
		   fmt.Println(err)
		   os.Exit(1)
   }

   imgByte := code.PNG()

   // convert byte to image for saving to file
   img, _, _ := image.Decode(bytes.NewReader(imgByte))

   err = imaging.Save(img, "./QRImgGA.png")

   if err != nil {
		   fmt.Println(err)
		   os.Exit(1)
   }

   // in real world application, the QRImgGA.png file should
   // be a temporary file with dynamic name.
   // for this tutorial sake, we keep it as static name.

   w.Write([]byte(fmt.Sprintf("<html><body><h1>QR code for : %s</h1><img src='http://localhost:8080/QRImgGA.png'>", authLink)))
   w.Write([]byte(fmt.Sprintf("<form action='http://localhost:8080/verify' method='post'>Token : <input name='token' id='token'><input type='submit' value='Verify Token'></form></body></html>")))
}

func main() {
   http.HandleFunc("/", Home)
   http.HandleFunc("/verify", Verify)

   // this is for displaying the QRImgGA.png from the source directory
   http.Handle("/QRImgGA.png", http.FileServer(http.Dir("./"))) //<---------------- here!

   http.ListenAndServe(":8080", nil)
}