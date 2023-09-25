import express from 'express'

const app = express()
const port = 8080

app.get('/ping', (_req, res) => {
  res.send({ message: 'pong' })
})

app.listen(port, () => {
  console.log(`Listening on port  ${port}`)
})
