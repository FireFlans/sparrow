package utils

import (
	"bytes"
	"fmt"

	svg "github.com/ajstarks/svgo"
)

// https://www.sgdsn.gouv.fr/files/files/Nos_missions/igi-1300-20210809.pdf

/*
<svg width="200" height="50" xmlns="http://www.w3.org/2000/svg">

	<!-- Rectangle with red border -->
	<rect x="5" y="5" width="190" height="40" fill="white" stroke="red" stroke-width="2.5"/>

	<!-- Centered text -->
	<text x="50%" y="50%" font-family="Arial" font-size="18" font-weight="bold" fill="red"
	      text-anchor="middle" dominant-baseline="central">
	    TRES SECRET
	</text>

</svg>

<svg width="200" height="50" xmlns="http://www.w3.org/2000/svg">

	<!-- Rectangle with red border -->
	<rect x="5" y="5" width="190" height="40" fill="white" stroke="blue" stroke-width="2.5"/>

	<!-- Centered text -->
	<text x="50%" y="50%" font-family="Arial" font-size="18" font-weight="bold" fill="blue"
	      text-anchor="middle" dominant-baseline="central">
	    SPECIAL FRANCE
	</text>

</svg>
*/
func GenerateSVG() []byte {
	buf := new(bytes.Buffer)
	canvas := svg.New(buf)
	canvas.Start(200, 200)
	canvas.Rect(0, 0, 200, 200, "fill:lightgray")
	canvas.Circle(100, 100, 50, fmt.Sprintf("fill:%s", "red"))
	canvas.Text(50, 50, fmt.Sprintf("%s - %s", "test_policy", "test_category"), "font-size:14px;fill:black")
	canvas.End()

	return buf.Bytes()
}
