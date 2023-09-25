import type { Config } from 'drizzle-kit'

export default {
    schema: './server.ts',
    driver: 'libsql',
    dbCredentials: {
        url: 'mydb.sqlite'
    },
    verbose: true,
    strict: true,
} satisfies Config
