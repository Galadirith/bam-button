const express = require('express');
const app = express();
const expressWs = require('express-ws')(app);
const events = require('events');
const event = new events.EventEmitter();

app.use(express.static('public'));

app.get('/', function(request, response) {
  response.sendFile(__dirname + '/views/index.html');
});

app.get('/bam', function(request, response) {
  event.emit('bam');
  response.sendStatus(204);
});

app.ws('/echo', function(ws, req) {
  console.log('connection to ws://bam-button.glitch.me/echo opened');
  var handler = () => {
    console.log('connection to ws://bam-button.glitch.me/echo sent bam');
    ws.send('bam');
  }
  event.on('bam', handler);
  ws.on('close', () => {
    console.log('connection to ws://bam-button.glitch.me/echo closed');
    event.removeListener('bam', handler)
  });
});

const listener = app.listen(process.env.PORT, function() {
  console.log('Your app is listening on port ' + listener.address().port);
});
