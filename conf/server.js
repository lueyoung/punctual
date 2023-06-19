var express = require('express'); 
var http = require('http'); 
var app = express();
var url = require("url");
var port = 8080 
var fs = require('fs');
var exec = require('child_process').execFile;

http.createServer(app).listen(port); 
app.get('/get',function(req,res){ 
	 var K = '-k';
         var url_parts = url.parse(req.url,true); 
        // console.log(url_parts);
         var query = url_parts.query; 
        // console.log(query);
         var key = query.key;
         var cmd = '/usr/local/bin/get'
         res.send('key: ' + key + '\n');
	 exec(cmd,['-k',key],null,function(err,stdout,stderr){
              if (err !== null) {
                  console.log('error: ' + error);
               }
           });
}); 
app.get('/post',function(req,res){ 
	 var K = '-k';
	 var V = '-v';
         var url_parts = url.parse(req.url,true); 
        // console.log(url_parts);
         var query = url_parts.query; 
        // console.log(query);
         var key = query.key;
         var value = query.value;
         res.send('the input key: ' + key );
         res.send('the input value: ' + value );
	 exec('/usr/local/bin/put'+' '+K+' '+key+' '+V+' '+value,function(err,stdout,stderr){
              if (err !== null) {
                  console.log('exec error: ' + error);
               }
           });
         res.send("the input process is completed!");
});
