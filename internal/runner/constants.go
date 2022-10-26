package runner

const (
  version = "0.1.1"
  author  = "ZyperX"
  banner  = `
_________
\____    \_ |__ ___  ___ ______ ______
  /     / | __ \\  \/  //  ___//  ___/
 /     /_ | \_\ \>    < \___ \ \___ \ 
/_______ \|___  /__/\_ /____  /____  >
        \/    \/      \/    \/     \/ 
      v` + version + ` - @` + author
  usage = `
  [buffers] | zbxss [options]
  `
  options = `
  -sql                      blind SQL check
  -X, --method <METHOD>     Specify request method to use (default: GET)
`
)
