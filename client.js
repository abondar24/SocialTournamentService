var http = require('http');
var fs = require('fs');
var path =require('path');
var filePath = path.join(__dirname, 'index.html');


fs.readFile(filePath, function (err, index) {
    if (err) {
        throw err;
    }

    http.createServer(function (request, response) {
        response.writeHead(200,{"Content-Type":"text/html"});
        response.write(index);
        response.end();
    }).listen(8217);
});

