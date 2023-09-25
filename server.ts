import express from 'express'

const app = express()
const port = 8080

app.get('/ping', (_req, res) => {
  res.send({ message: 'pong' })
})

app.listen(port, () => {
  console.log(`Listening on port  ${port}`)
})

import { drizzle, BunSQLiteDatabase } from 'drizzle-orm/bun-sqlite'
import { integer, text, sqliteTable } from 'drizzle-orm/sqlite-core'
import { Database } from 'bun:sqlite'

const sqlite = new Database('mydb.sqlite', { create: true })
const db: BunSQLiteDatabase = drizzle(sqlite)

export const users = sqliteTable('users', {
  id: integer('id').primaryKey({ autoIncrement: true }),
  fullName: text('full_name'),
  phone: text('phone'),
})

export type User = typeof users.$inferSelect // return type when queried
export type InsertUser = typeof users.$inferInsert // insert type

const result: User[] = db.select().from(users).all()

const insertUser = (user: InsertUser) => {
  return db.insert(users).values(user).run()
}
