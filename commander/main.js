//!/usr/bin/env node

const express = require('express')
const uuid = require('uuid')

const app = express()

app.get('/command', (req, res) => {
    res.send({
        id: uuid.v4(),
        cmd: 'ls',
        time: 0,
        signature: '-'
    })
})

app.listen(3000, () => console.log('Commander listening on port 3000!'))