const express = require('express')
const app = express()
const port = 8000

app.use(express.json()); 
app.use(express.urlencoded({extended: true}));

app.get('/', (req, res) => {
  res.status(200).send("Welcome")
})

app.get('/get', (req, res) => {
    res.status(200).json({message: "Hello from node server"})
  })


app.post('/g', (req, res) => {
    let myJson = req.body;
	
	res.status(200).send(myJson);
})

app.post('/postform', (req, res) => {
    res.status(200).send(JSON.stringify(req.body));
})
  

app.listen(port, () => {
  console.log(`server listening at http://localhost:${port}`)
})