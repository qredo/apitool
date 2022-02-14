package webui

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const indexHTML string = `<html>
<head><title>Qredo API tool</title></head>
<script>
	function sign() {
		var form = document.getElementById('form')
		let req = {
			"api_key": form.apiKey.value,
			"secret" : form.apiSecret.value,
            "url"    : form.url.value,
			"method" : form.method_select.value,
            "body"   : form.body.value
        }
		fetch('http://localhost:4569/sign', {
			method: 'POST',
		    headers: {
    		  'Content-Type': 'application/json'
    		},
            body: JSON.stringify(req)
        })
  		.then(response => response.json())
  		.then(data => writeResponse(data))
    }

	function send() {
		var form = document.getElementById('form')
		let req = {
			"api_key": form.apiKey.value,
			"secret" : form.apiSecret.value,
            "url"    : form.url.value,
			"method" : form.method_select.value,
            "body"   : form.body.value
        }
		fetch('http://localhost:4569/send', {
			method: 'POST',
		    headers: {
    		  'Content-Type': 'application/json'
    		},
            body: JSON.stringify(req)
        })
  		.then(response => response.json())
  		.then(data => writeResponse(data))
    }
	
	function writeResponse(response) {
		var headerRespDiv = document.getElementById('headerResponse');
		headerRespDiv.style.visibility = 'visible'
		var respKey = document.getElementById('apiKeyResult');
		respKey.textContent = response.api_key
		var respTs = document.getElementById('apiTsResult');
		respTs.textContent = response.timestamp
		var respSig = document.getElementById('apiSigResult');
		respSig.textContent = response.signature
		var respResp = document.getElementById('apiResponse');
		respResp.textContent = response.response
    }

</script
<body>
	<center>
		<table style="width:80%">
		<tr>
		<td>
		<form id="form"">
  			<label for="apiKey">API Key</label><br>
  			<input type="text" id="apiKey" size="70" value=""><br><br>
  			<label for="apiSecret">API Secret</label><br>
  			<input type="text" id="apiSecret" size="70" value=""><br><br>
			<label for="url">URL</label><br>
  			<input type="text" id="url" size="70" value=""><br><br>
			<label for="method_select">Method</label>
			<select id="method_select">
    			<option value="GET">GET</option>
    			<option value="POST">POST</option>
    			<option value="PUT">PUT</option>
    			<option value="PATCH">PATCH</option>
    			<option value="DELETE">DELETE</option>
			</select><br><br>
			<label for="body">Body</label><br>
			<textarea id="body" name="body" rows="20" cols="60"></textarea><br><br>
  			<input type="button" value="Sign" onclick="sign()">
			<input type="button" value="Send" onclick="send()">
		</form>
		</td>
		<td>
		<div id="headerResponse" style="visibility: hidden">
			<table>
			<tr>
				<td><b>Qredo-API-Key:</td>
				<td><span id="apiKeyResult"></span></td>
			<tr>
				<td><b>Qredo-API-Ts:</b></td>
				<td><span id="apiTsResult"></span></td>
			</tr>
			<tr>
				<td><b>Qredo-API-Sig:</td>
				<td><span id="apiSigResult"></span></td>
			</tr>
			</table>
		</div>
		</td>
		<tr>
		</table>
	</center>
	<hr>
	<div id="serverResponse" style="width:80%; margin:auto">
		<pre id="apiResponse"></pre>
	</div
</body>
</html>
`

func GetIndex(c *gin.Context) {
	c.Header("Content-type", "text/html")
	c.String(http.StatusOK, indexHTML)
}
