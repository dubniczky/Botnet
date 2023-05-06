//!/usr/bin/env node

const express = require('express')
const uuid = require('uuid')

const app = express()

let count = 0
let testCommand = {
    id: uuid.v4(),
    cmd: 'ls',
    time: 0,
    signature: '-',
    once: true
}

app.get('/command', (req, res) => {
    count++
    console.log('Commands sent: ' + count)
    res.send(testCommand)
})

app.listen(3000, () => console.log('Commander listening on port 3000!'))