let express = require('express');
let serveStatic = require('serve-static');
let route = require('./routes/route');
let bodyParser = require('body-parser');
let logger = require('morgan');

app = express();

app.use(serveStatic(__dirname + "/dist"));
let port = process.env.PORT || 8217;
app.listen(port);

app.set('route',__dirname+'/routes');

app.use('/',route);
app.use(bodyParser.json());
app.use(logger('dev'));

console.log('server started '+ port);
module.exports=app;